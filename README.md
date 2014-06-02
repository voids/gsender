gsender
=======

a smtp sender written by golang

##install

`go get github.com/voids/gsender`

```go
	import (
    	"github.com/voids/gsender"
    )
```
##demo

```go
	sender := new(gsender.Sender)
	sender.Mail = new(gsender.Mail)        // alloc memory for the anonymous struct
	sender.Address = "voids@example.com"   // sender.Mail.Address = xxx
	sender.Name = "voids"                  // sender.Mail.Name = xxx
	sender.Password = "********"
	sender.Host = "smtp.example.com"
	sender.Port = 994
	sender.TLS = true

	msg := new(gsender.Message)
	msg.SetSubject("A test email")
	msg.SetBody("this is a test email which sent by golang.")
	if err := msg.AddAttachment(`/home/voids/images/icon_clockwork.png`); err != nil {
		panic(err)
	}

	receiver := new(gsender.Receiver)
	receiver.AddTo("John", "John@example.com")
    receiver.AddTo("Lily", "Lily@example.com")
	receiver.AddCc("Dog", "dog@example.com")
    receiver.AddBcc("hacker", "hacker@example.com")

	if err := sender.Send(msg, receiver); err != nil {
		panic(err)
	}
```