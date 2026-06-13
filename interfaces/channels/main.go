package main

import "channels/notification"

// type NotificationService struct{}

// func (s NotificationService) SendAll(notifications []Notification) {
// 	for _, notification := range notifications {
// 		notification.Send()
// 	}
// }

func main() {
	notifications := []notification.Notification{
		notification.Email{
			To:      "test@mail.com",
			Content: "Hello",
		},
		notification.Email{
			To:      "test@mail.com",
			Content: "Hello",
		}, notification.Email{
			To:      "test@mail.com",
			Content: "Hello",
		},
	}

	for _, ntf := range notifications {
		ntf.Send()
	}
}
