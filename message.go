package gsender

import (
	"encoding/base64"
	"io/ioutil"
	"path"
)

type Attachment struct {
	name    string
	content string
}

type Message struct {
	subject     string
	body        string
	attachments []*Attachment
	// set the type to "html" if Html is true.
	Html bool
}

func (m *Message) SetSubject(subject string) {
	if subject != "" {
		m.subject = "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(subject)) + "?="
	} else {
		m.subject = ""
	}
}

func (m *Message) SetBody(body string) {
	m.body = base64EncodeRfc2045([]byte(body))
}

func (m *Message) AddAttachment(filename string) error {
	attachment := new(Attachment)
	attachment.name = "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(path.Base(filename))) + "?="
	if attach, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else {
		attachment.content = base64EncodeRfc2045(attach)
	}
	m.attachments = append(m.attachments, attachment)
	return nil
}
