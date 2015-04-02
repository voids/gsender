package gsender

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
)

const BASE64_MAX_LEN = 76

type Mail mail.Address

func (m *Mail) String() string {
	return (*mail.Address)(m).String()
}

func addressListToString(list []*Mail) []string {
	addresses := make([]string, 0)
	for _, address := range list {
		addresses = append(addresses, address.String())
	}
	return addresses
}

func addressListEmails(lists ...[]*Mail) []string {
	emails := make([]string, 0)
	for _, list := range lists {
		for _, address := range list {
			emails = append(emails, address.Address)
		}
	}
	return emails
}

func base64EncodeRfc2045(src []byte) string {
	dsc := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dsc, src)
	buf := bytes.NewBuffer(dsc)
	result := ""
	for {
		tmp := buf.Next(BASE64_MAX_LEN)
		if len(tmp) == 0 {
			break
		}
		result += string(tmp) + "\r\n"
	}
	return strings.TrimSpace(result)
}

func dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func sendMail(addr string, auth smtp.Auth, from string, to []string, msg []byte, tls bool) (err error) {
	var c *smtp.Client
	if tls == true {
		c, err = dial(addr)
	} else {
		c, err = smtp.Dial(addr)
	}
	if err != nil {
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
