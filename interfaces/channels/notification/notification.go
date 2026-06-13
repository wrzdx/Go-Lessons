package notification

import "fmt"

type Notification interface {
	Send()
}

type Email struct {
	To      string
	Content string
}

type SMS struct {
	To      string
	Content string
}
type Push struct {
	To      string
	Content string
}

func (e Email) Send() {
	fmt.Printf("Email sent to %s: %s\n", e.To, e.Content)
}

func (e SMS) Send() {
	fmt.Printf("SMS sent to %s: %s\n", e.To, e.Content)
}
func (e Push) Send() {
	fmt.Printf("Push sent to %s: %s\n", e.To, e.Content)
}