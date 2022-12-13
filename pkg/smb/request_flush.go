package smb

import (
	"bytes"
	"encoding/binary"
)

// FlushRequest structure represents a request to flush the buffers of a file in the Server Message Block (SMB) protocol.
// It has the following fields:
//     WordCount: an 8-bit integer indicating the number of 16-bit words in the request.
//     FID: a 16-bit integer identifying the file to flush.
//     Reserved: a 16-bit integer reserved for future use.
//     Flags: a 16-bit integer containing flags that modify the request.
//     NameLength: an 8-bit integer indicating the length of the FileName field in bytes.
//     FileName: a variable-length string containing the name of the file to flush.
//     CreateContextData: a variable-length byte slice containing additional data for the request.
type FlushRequest struct {
	WordCount         uint8
	FID               uint16
	Reserved          uint16
	Flags             uint16
	NameLength        uint8
	FileName          string
	CreateContextData []byte
}

func FlushRequestParse(data []byte) (*FlushRequest, error) {
	// Parse the request structure from the byte slice
	var request FlushRequest
	err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &request)
	if err != nil {
		return nil, err
	}

	// Read the variable-length fields from the request
	buf := bytes.NewReader(data[12:])

	// Read the file name
	fileName, err := readVarString(buf, int(request.NameLength), request.Flags&FlagUnicode != 0)
	if err != nil {
		return nil, err
	}
	request.FileName = fileName

	// Read the create context data
	createContextData, err := readVarBytes(buf)
	if err != nil {
		return nil, err
	}
	request.CreateContextData = createContextData

	// Return the parsed request
	return &request, nil
}
