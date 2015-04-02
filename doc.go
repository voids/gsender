/*
package gsender makes smtp easy to use.

Installation

	go get github.com/voids/gsender

Get started

	sender := new(gsender.Sender)
	sender.Address = "voids@example.com"
	sender.Name = "voids"
	sender.Password = "********"
	sender.Host = "smtp.example.com"
	sender.Port = 25
	// sender.TLS = true

	msg := new(gsender.Message)
	// msg.Html = true
	msg.SetSubject("A test email")
	msg.SetBody("this is a test email which sent by golang.")
	if err := msg.AddAttachment(`/home/voids/pic.jpg`); err != nil {
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


*/
package gsender
