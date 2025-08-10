package request

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	
	bytes, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("ERROR- io.ReadAll: %v\n", err.Error())
	}

	headers := string(bytes)
	httpParts := strings.Split(headers, "\r\n")
	if len(httpParts) == 0 {
		return nil, fmt.Errorf("request header too short")
	}

	reqLineParts := strings.Split(httpParts[0], " ")
	if len(reqLineParts) < 3 {
		return nil, fmt.Errorf("not enough sections in request line")
	}

	httpVer := strings.Split(reqLineParts[2], "/")
	if len(httpVer) < 2 {
		return nil, fmt.Errorf("invalid http version")
	}

	reqLine := RequestLine{
		Method:       reqLineParts[0],
		RequestTarget: reqLineParts[1],
		HttpVersion:   httpVer[1],
	}

	return &Request{RequestLine: reqLine}, nil


}

func parseRequestLine() {

}