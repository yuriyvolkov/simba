package smb

import (
	"bytes"
	"encoding/binary"
)

type TreeDisconnectRequest struct {
	WordCount uint8
	ByteCount uint16
}

func TreeDisconnectRequestParse(data []byte) (*TreeDisconnectRequest, error) {
	// Parse the request structure from the byte slice
	var request TreeDisconnectRequest
	err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &request)
	if err != nil {
		return nil, err
	}

	// Return the parsed request
	return &request, nil
}
