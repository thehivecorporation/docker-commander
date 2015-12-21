package communications

//ConnectionResponse is a wrapper to be used in every websocket response
type ConnectionResponse struct {
	Response interface{}
	Action   string
}
