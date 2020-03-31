# botWhatsappCovid
first you have to install [go-whatsapp](https://github.com/Rhymen/go-whatsapp) from ryhmen 

## Installation

Use this command to install [go-whatsapp](https://github.com/Rhymen/go-whatsapp).

```bash
go get github.com/Rhymen/go-whatsapp
```

## Usage

copy the botWhatsappCovid to your go-whatsapp folder

## API
this  work using api from [Novel Covid API](https://github.com/NovelCOVID/API)

```go
var stringCut = strings.TrimLeft(message.Text,"Covid19: ")
	url := "https://corona.lmao.ninja/countries/"+strings.ToUpper(stringCut)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
```

## Test
![screenshot](https://scontent-sin6-2.xx.fbcdn.net/v/t1.15752-9/92025924_281793502810950_4359468354642116608_n.jpg?_nc_cat=103&_nc_sid=b96e70&_nc_eui2=AeGHtQtpJ6ID8d2J_73Uju_JgEG5es5kxOFGWNOAnGJC32qy6RwLC5ae376FkL_lcDyVOoBOqPNKVlW7rNhMgMRgs_IpEVQelSsSHOzpnN7z-w&_nc_oc=AQnA7kCkHa4tIDiqeJNDL1-R_qiwtSyGc0n9x3wz0ji0NzxcHIubgeH00Homru-KRZI&_nc_ht=scontent-sin6-2.xx&oh=1bb4fe9a38ef769ae8d38645a19f7b8b&oe=5EAA0E82)


## Source
Big Thanks for Rhymen-[go-whatsapp](https://github.com/Rhymen/go-whatsapp) and [Novel Covid API](https://github.com/NovelCOVID/API)
