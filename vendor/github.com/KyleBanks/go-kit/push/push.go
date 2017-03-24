// Package push provides GCM and APN push notification functionality.
package push

// Pusher defines the interface for a push notification Pusher.
type Pusher interface {

	// SendMessage sends the payload provided as a JSON string to the deviceIds.
	SendMessage(message *Message, deviceIds string) error
}

// Message defines a push message payload to be send to the client.
type Message struct {

	// Content is the title or content to be displayed to the user.
	Content string

	// Data contains any additional data to be sent with the notification.
	Data map[string]interface{}

	// IosSound is the sound to play for iOS devices specifically.
	IosSound string
}
