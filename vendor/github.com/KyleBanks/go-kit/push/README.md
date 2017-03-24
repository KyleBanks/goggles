# push
--
    import "github.com/KyleBanks/go-kit/push/"

Package push provides GCM and APN push notification functionality.

## Usage

```go
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
```

```go
const (
	// MaxRetries defines the number of times to retry message sender if an error occurs.
	MaxRetries = 2
)
```

#### type AndroidPusher

```go
type AndroidPusher struct {
}
```

AndroidPusher supports sending of GCM push notifications to Android devices.

#### func  NewAndroidPusher

```go
func NewAndroidPusher(apiKey string) *AndroidPusher
```
NewAndroidPusher returns an AndroidPusher, initialized with the specified GCM
API key.

#### func (*AndroidPusher) SendMessage

```go
func (a *AndroidPusher) SendMessage(message *Message, deviceIds ...string) error
```
SendMessage sends a JSON payload to the specified DeviceIds through the GCM
service.

#### type IosPusher

```go
type IosPusher struct {
}
```

IosPusher supports sending of APNS push notifications to iOS devices.

#### func  NewIosPusher

```go
func NewIosPusher(isProd bool, certificateFile string, keyFile string) *IosPusher
```
NewIosPusher returns an IosPusher, initialized with the specified certificate
and key files.

#### func (*IosPusher) SendMessage

```go
func (i *IosPusher) SendMessage(message *Message, deviceIds ...string) error
```
SendMessage sends a JSON payload to the specified DeviceIds through the APNS
service.

#### type Message

```go
type Message struct {

	// Content is the title or content to be displayed to the user.
	Content string

	// Data contains any additional data to be sent with the notification.
	Data map[string]interface{}

	// IosSound is the sound to play for iOS devices specifically.
	IosSound string
}
```

Message defines a push message payload to be send to the client.

#### type Pusher

```go
type Pusher interface {

	// SendMessage sends the payload provided as a JSON string to the deviceIds.
	SendMessage(message *Message, deviceIds string) error
}
```

Pusher defines the interface for a push notification Pusher.
