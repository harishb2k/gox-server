package server

import (
	"fmt"
	"github.com/harishb2k/gox-base"
	"github.com/harishb2k/gox-errors"
	"io/ioutil"
	"net/http"
)

type DefaultApplicationServer struct {
	Port int
}

func (ap *DefaultApplicationServer) Start() (err error) {
	e := http.ListenAndServe(fmt.Sprintf(":%d", ap.Port), nil)
	return errors.Wrap(e, "Failed to start http server")
}

func (*DefaultApplicationServer) Stop() (err error) {
	return nil
}

func (ap *DefaultApplicationServer) Register(path string, handler RequestHandlerFunc) {
	http.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {

		// Build http request which is sent to client for processing
		request := &Request{HttpRequest: req}

		var err error
		if request.ByteBody, err = ioutil.ReadAll(req.Body); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer req.Body.Close()

		if request.ByteBody == nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Get a extractor which can be used by client handle function to get data from http request
		extractor, err := NewDefaultRequestExtractor(request)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Call registered handler and get response
		response, err := handler(request, extractor)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Set content type in response
		if response.ContentType == "" {
			w.Header().Add("Content-Type", "application/json")
		} else {
			w.Header().Add("Content-Type", response.ContentType)
		}

		// Set http status code
		if response.StatusCode == 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(response.StatusCode)
		}

		// Write response to client
		if response.Body != nil {
			if body, err := gox.BytesWithError(response.Body); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			} else {
				w.Write(body)
			}
		} else if response.ByteBody != nil {
			w.Write(response.ByteBody)
		}
	})
}
