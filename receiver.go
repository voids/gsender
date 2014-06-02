package gsender

type Receiver struct {
	to  []*Mail
	cc  []*Mail
	bcc []*Mail
}

func (r *Receiver) AddTo(name, email string) {
	r.to = append(r.to, &Mail{Name: name, Address: email})
}

func (r *Receiver) AddCc(name, email string) {
	r.cc = append(r.cc, &Mail{Name: name, Address: email})
}

func (r *Receiver) AddBcc(name, email string) {
	r.bcc = append(r.bcc, &Mail{Name: name, Address: email})
}
