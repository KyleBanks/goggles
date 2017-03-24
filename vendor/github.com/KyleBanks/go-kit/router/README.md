# router
--
    import "github.com/KyleBanks/go-kit/router/"

Package router defines the Route interface, and registers routes to an http
server.

## Usage

#### func  HasParam

```go
func HasParam(r *http.Request, key string) bool
```
HasParam returns a boolean indicating if the request has a particular parameter.

#### func  Param

```go
func Param(r *http.Request, key string) string
```
Param returns a POST/GET parameter from the request.

If the parameter is found in the POST and the GET parameter set, the POST
parameter will be given priority.

#### func  Register

```go
func Register(s Server, routes []Route)
```
Register registers each Route with the Server provided.

Each Route will be wrapped in a middleware function that adds trace logging.

#### type Handler

```go
type Handler func(http.ResponseWriter, *http.Request)
```

Handler defines a function that accepts an HTTP request and returns a Response.

#### type Route

```go
type Route struct {
	Path   string // The URL path to listen for (i.e. "/api")
	Handle Handler
}
```

Route defines a URL path and function to execute when the URL is accessed.

#### type Server

```go
type Server interface {
	HandleFunc(string, func(http.ResponseWriter, *http.Request))
}
```

Server defines an interface for the provided server to comply with.
