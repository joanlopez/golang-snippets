package http

import (
	"bufio"
	"errors"
	"net"
	"net/textproto"
	"strings"
)

func ReadRequest(conn net.Conn) (Request, error) {
	tp := textproto.NewReader(bufio.NewReader(conn))

	var first string
	var err error

	if first, err = tp.ReadLine(); err != nil {
		return Request{}, err
	}

	reqMethod, reqURI, reqProto, ok := parseRequestLine(first)
	if !ok {
		return Request{}, errors.New("invalid first line")
	}

	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		return Request{}, err
	}

	var reqBytes []byte

	for {
		recvBuf := make([]byte, 1024)

		// Non-blocking,
		// reads as many bytes as possible
		n, err := tp.R.Read(recvBuf[:])
		if err != nil {
			return Request{}, err
		}

		// Only append the first n bytes, the value of
		// the other ones are corresponds to zero value
		reqBytes = append(reqBytes, recvBuf[:n]...)

		// If we didn't read as many byte as the size
		// of the slice, it means there's no more data
		if n < len(recvBuf) {
			break
		}
	}

	// Remember to call `conn.Close()` if you no longer
	// need to perform additional operations on connection.
	return Request{
		Method:     reqMethod,
		RequestURI: reqURI,
		Proto:      reqProto,
		Headers:    mimeHeader,
		Body:       reqBytes,
	}, nil
}

func parseRequestLine(line string) (method, requestURI, proto string, ok bool) {
	s1 := strings.Index(line, " ")
	s2 := strings.Index(line[s1+1:], " ")

	if s1 < 0 || s2 < 0 {
		return
	}

	s2 += s1 + 1

	return line[:s1], line[s1+1 : s2], line[s2+1:], true
}
