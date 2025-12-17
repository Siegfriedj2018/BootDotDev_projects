package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion		string
	RequestTarget	string
	Method 				string
}

func parseRequestLine(request []byte) (*RequestLine, error) {
	httpParsed := strings.Split(string(request), "\r\n")
	
	return httpParsed
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	req, err := io.ReadAll(reader)
	if (err != nil) {
		return &Request{}, fmt.Errorf("could not read bytes: %w", err)
	}

	reqLine, err := parseRequestLine(req)

}