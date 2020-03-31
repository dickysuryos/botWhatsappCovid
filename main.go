package main

import (
	"encoding/gob"
	"fmt"
	"github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp/binary/proto"
	"github.com/Rhymen/go-whatsapp"
	"os"
	"strings"
	"time"
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
	"strconv"
	
)

type waHandler struct {
	wac       *whatsapp.Conn
	startTime uint64
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func (wh *waHandler) HandleError(err error) {
	fmt.Fprintf(os.Stderr, "error caught in handler: %v\n", err)
}

type GoCorona struct {
    Deaths int
    Country string
    Recovered int 
    Cases int
    TodayCases int
    Updated int
}


// HandleTextMessage receives whatsapp text messages and checks if the message was send by the current
// user, if it does not contain the keyword '@echo' or if it is from before the program start and then returns.
// Otherwise the message is echoed back to the original author.
func (wh *waHandler) HandleTextMessage(message whatsapp.TextMessage) {
	if !strings.Contains(strings.ToLower(message.Text), "covid19: ") || message.Info.Timestamp < wh.startTime {
		return
	}
	var stringCut = strings.TrimLeft(message.Text,"Covid19: ")
	url := "https://corona.lmao.ninja/countries/"+strings.ToUpper(stringCut)


	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	

	res, getErr := myClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	if body != nil {
    	defer res.Body.Close()
    }
    fmt.Printf("data '%v'",res.Body)
 
    var corona GoCorona
	jsonErr := json.Unmarshal([]byte(body), &corona)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
previousMessage := message.Text
	quotedMessage := proto.Message{
		Conversation: &previousMessage,
	}

ContextInfo := whatsapp.ContextInfo{
		QuotedMessage:   &quotedMessage,
		QuotedMessageID: message.Text,
		Participant:message.Info.RemoteJid, //Whot sent the original message
	}


	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: message.Info.RemoteJid,

		},
		ContextInfo: ContextInfo,
		Text: "Data Covid19 untuk negara " + corona.Country +" meninggal : "+ strconv.Itoa(corona.Deaths) +", sembuh :" + strconv.Itoa(corona.Recovered) + ", kasus hari ini :  "+ strconv.Itoa(corona.TodayCases) +" , total kasus : "+ strconv.Itoa(corona.Cases) +", Stay Healthy Stay at Home" ,
	}

	if _, err := wh.wac.Send(msg); err != nil {
		fmt.Fprintf(os.Stderr, "error sending message: %v\n", err)
	}

}


func getJson(url string, target interface{}) error {
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }

    defer r.Body.Close()
	  fmt.Printf("data : '%v' ", r.Body)
    return json.NewDecoder(r.Body).Decode(target)
}

func login(wac *whatsapp.Conn) error {
	session, err := readSession()
	if err == nil {
		session, err = wac.RestoreWithSession(session)
		if err != nil {
			return fmt.Errorf("restoring session failed: %v", err)
		}
	} else {
		qr := make(chan string)

		go func() {
			terminal := qrcodeTerminal.New()
			terminal.Get(<-qr).Print()
		}()

		session, err = wac.Login(qr)
		if err != nil {
			return fmt.Errorf("error during login: %v", err)
		}
	}

	if err = writeSession(session); err != nil {
		return fmt.Errorf("error saving session: %v", err)
	}

	return nil
}

func readSession() (whatsapp.Session, error) {
	session := whatsapp.Session{}

	file, err := os.Open(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return session, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err = decoder.Decode(&session); err != nil {
		return session, err
	}

	return session, nil
}

func writeSession(session whatsapp.Session) error {
	file, err := os.Create(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err = encoder.Encode(session); err != nil {
		return err
	}

	return nil
}

func main() {
	wac, err := whatsapp.NewConn(5 * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
		return
	}

	wac.AddHandler(&waHandler{wac, uint64(time.Now().Unix())})

	if err = login(wac); err != nil {
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		return
	}

	<-time.After(60 * time.Minute)
}
