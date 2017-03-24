package push

import (
	"github.com/alexjlockwood/gcm"
)

const (
	// MaxRetries defines the number of times to retry message sender if an error occurs.
	MaxRetries = 2
)

// AndroidPusher supports sending of GCM push notifications to Android devices.
type AndroidPusher struct {
	gcm gcm.Sender
}

// NewAndroidPusher returns an AndroidPusher, initialized with the specified
// GCM API key.
func NewAndroidPusher(apiKey string) *AndroidPusher {
	return &AndroidPusher{
		gcm: gcm.Sender{
			ApiKey: apiKey,
		},
	}
}

// SendMessage sends a JSON payload to the specified DeviceIds through the GCM service.
func (a *AndroidPusher) SendMessage(message *Message, deviceIds ...string) error {
	notif := map[string]interface{}{
		"data": message.Data,
	}
	if len(message.Content) > 0 {
		notif["message"] = message.Content
	}

	msg := gcm.NewMessage(notif, deviceIds...)
	if _, err := a.gcm.Send(msg, MaxRetries); err != nil {
		return err
	}

	return nil
}
