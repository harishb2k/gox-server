package server

import "net/http"

type Request struct {
    ByteBody    []byte
    Body        interface{}
    HttpRequest *http.Request
    Extractor   RequestExtractor
}

type Response struct {
    ByteBody   []byte
    Body       interface{}
    StatusCode int
    Err        error
    Headers    map[string]string
}

type RequestHandlerFunc func(request *Request, extractor RequestExtractor) (response *Response, err error)

type Application interface {
    Start() (err error)
    Register(path string, handler RequestHandlerFunc)
    Stop() (err error)
}
