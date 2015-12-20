package entities

//SocketResponse is a wrapper to be used in every websocket response
type SocketResponse struct {
	Response interface{}
	Action   string
}
