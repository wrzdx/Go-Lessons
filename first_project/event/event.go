package event

import (
	"fmt"
	"time"
)

type Event struct {
	input     string
	msg       string
	createdAt time.Time
}

func (event *Event) CreatedAt() time.Time {
	return event.createdAt
}

func (event *Event) Input() string {
	return event.input
}

func (event *Event) Msg() string {
	return event.msg
}

func CreateEvent(input string, msg string) Event {
	return Event{
		input:     input,
		msg:       msg,
		createdAt: time.Now(),
	}
}

func (event Event) String() string {
	template := `Event{
    Input: %q, 
    Message: %q, 
    CreatedAt: %s, 
}`

	return fmt.Sprintf(
		template,
		event.input,
		event.msg,
		event.createdAt.Format(time.DateTime),
	)
}
