package smb

import (
	"bytes"
	"encoding/binary"
)

// NegotiateRequest represents an SMB negotiate request
type NegotiateRequest struct {
	Dialects Dialects
}

// NegotiateRequestParse parses an SMB negotiate request
func NegotiateRequestParse(data []byte) (*NegotiateRequest, error) {
	// Create a new bytes reader
	r := bytes.NewReader(data)

	// Read the word count
	var wordCount uint8
	if err := binary.Read(r, binary.LittleEndian, &wordCount); err != nil {
		return nil, err
	}

	// Read the dialects
	dialects := make([]Dialect, wordCount)
	if err := binary.Read(r, binary.LittleEndian, &dialects); err != nil {
		return nil, err
	}

	// Create the request
	request := &NegotiateRequest{
		Dialects: dialects,
	}

	return request, nil
}
