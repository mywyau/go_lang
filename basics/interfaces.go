package basics

import "fmt"

type Notifier interface {
	Notify(msg string)
}

type Emailer struct{ Address string }

func (e Emailer) Notify(msg string) {
	fmt.Println("sending email:", msg)
}

func Send(n Notifier) {
	n.Notify("hello!")
}
