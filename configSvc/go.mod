module github.com/pramineni01/madr/configsvc

go 1.14

replace github.com/pramineni01/madr/configstore => ../configstore

replace github.com/pramineni01/madr/mockgrpc => ../mockgrpc

require (
	github.com/labstack/echo/v4 v4.1.16
	github.com/labstack/gommon v0.3.0 // indirect
	golang.org/x/crypto v0.0.0-20200414173820-0848c9571904 // indirect
	github.com/pramineni01/madr/configstore v1.0.0 // indirect
	github.com/pramineni01/madr/mockgrpc v1.0.0 // indirect
)
