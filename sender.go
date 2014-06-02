package gsender

import (
	"fmt"
	"crypto/md5"
	"net/smtp"
	"bytes"
	"strings"
	"time"
	"path/filepath"
	"io"
	"mime"
)

type Sender struct {
	*Mail
	Password string
	Host     string
	Port     uint
	TLS      bool
}

func (s *Sender) Send(m *Message, r *Receiver) error {
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%s", time.Now().Nanosecond()))
	boundary := fmt.Sprintf("%x", h.Sum(nil))
	to := addressListToString(r.to)
	cc := addressListToString(r.cc)
	bcc := addressListToString(r.bcc)
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "From: %s\r\n", s.String())
	if len(to) > 0 {
		fmt.Fprintf(&buf, "To: %s\r\n", strings.Join(to, ","))
	}
	if len(cc) > 0 {
		fmt.Fprintf(&buf, "Cc: %s\r\n", strings.Join(cc, ","))
	}
	if len(bcc) > 0 {
		fmt.Fprintf(&buf, "Bcc: %s\r\n", strings.Join(bcc, ","))
	}
	fmt.Fprintf(&buf, "Subject: %s\r\n", m.subject)
	fmt.Fprintf(&buf, "Message-ID: <%v.%s>\r\n", time.Now().UnixNano(), s.Address)
	fmt.Fprintf(&buf, "Date: %s\r\n", time.Now().UTC().Format(time.RFC822))
	fmt.Fprintf(&buf, "MIME-Version: 1.0\r\n")
	fmt.Fprintf(&buf, "Content-Type: multipart/mixed; boundary=%s\r\n\r\n", boundary)
	fmt.Fprintf(&buf, "--%s\r\n", boundary)
	textType := "plain"
	if m.htmlEnable == true {
		textType = "html"
	}
	fmt.Fprintf(&buf, "Content-Type: text/%s; charset=UTF-8\r\n", textType)
	fmt.Fprintf(&buf, "Content-Transfer-Encoding: base64\r\n\r\n")
	fmt.Fprintf(&buf, "%s\r\n\r\n", m.body)
	for _, attach := range m.attachments {
		fname := filepath.Clean(attach.name)
		name := filepath.Base(fname)
		mimetype := mime.TypeByExtension(filepath.Ext(name))
		if mimetype == "" {
			mimetype = "application/octet-stream"
		}
		fmt.Fprintf(&buf, "--%s\r\n", boundary)
		fmt.Fprintf(&buf, "Content-Type: %s; name=%s\r\n", mimetype, name)
		fmt.Fprintf(&buf, "Content-Disposition: attachment; filename=%s\r\n", name)
		fmt.Fprintf(&buf, "Content-Transfer-Encoding: base64\r\n\r\n")
		fmt.Fprintf(&buf, "%s\r\n\r\n", attach.content)
	}
	fmt.Fprintf(&buf, "--%s--\r\n", boundary)
	auth := smtp.PlainAuth("", s.Address, s.Password, s.Host)
	return SendMail(s.Host+fmt.Sprintf(":%v", s.Port), auth, s.Address, addressListEmails(r.to, r.cc, r.bcc), buf.Bytes(), s.TLS)
}
