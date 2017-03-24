package push

import (
	"fmt"

	"github.com/anachronistic/apns"
)

const (
	// ApnsEndpointSandbox is the endpoint of the sandbox APNS service.
	ApnsEndpointSandbox = "gateway.sandbox.push.apple.com"

	// ApnsEndpointProduction is the endpoint of the production APNS service.
	ApnsEndpointProduction = "gateway.push.apple.com"

	// ApnsPort is the port of the APNS service.
	ApnsPort = 2195

	// IosDefaultSound is the name of the default sound to play for iOS notifications.
	IosDefaultSound = "default"
)

// IosPusher supports sending of APNS push notifications to iOS devices.
type IosPusher struct {
	apns *apns.Client
}

// NewIosPusher returns an IosPusher, initialized with the specified
// certificate and key files.
func NewIosPusher(isProd bool, certificateFile string, keyFile string) *IosPusher {
	endpoint := ApnsEndpointSandbox
	if isProd {
		endpoint = ApnsEndpointProduction
	}

	return &IosPusher{
		apns: apns.NewClient(fmt.Sprintf("%v:%v", endpoint, ApnsPort), certificateFile, keyFile),
	}
}

// SendMessage sends a JSON payload to the specified DeviceIds through the APNS service.
func (i *IosPusher) SendMessage(message *Message, deviceIds ...string) error {
	for _, deviceID := range deviceIds {
		// Construct the APNS payload...
		payload := apns.NewPayload()
		if len(message.Content) > 0 {
			payload.Alert = message.Content
		}
		if len(message.IosSound) > 0 {
			payload.Sound = message.IosSound
		}

		pn := apns.NewPushNotification()
		pn.DeviceToken = deviceID
		pn.AddPayload(payload)

		for key := range message.Data {
			pn.Set(key, message.Data[key])
		}

		res := i.apns.Send(pn)
		if res.Error != nil {
			// TODO: Continue attempting to send notifictions,
			// wrapping the error for each error that occurs.
			return res.Error
		}
	}

	return nil
}
