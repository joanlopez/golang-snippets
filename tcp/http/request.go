package http

type Request struct {
	Method     string
	RequestURI string
	Proto      string
	Headers    map[string][]string
	Body       []byte
}
