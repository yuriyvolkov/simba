package smb

import (
	"bytes"
	"encoding/binary"
)

// SetInfoRequest structure represents a request to set information about a file in the Server Message Block (SMB) protocol.
// It has the following fields:

//     WordCount: an 8-bit integer indicating the number of 16-bit words in the request.
//     InformationLevel: a 16-bit integer specifying the type of information to set.
//     Reserved: an 8-bit integer reserved for future use.
//     NameLength: an 8-bit integer indicating the length of the FileName field in bytes.
//     Flags: a 16-bit integer containing flags that modify the request.
//     FileName: a variable-length string containing the name of the file to set.
//     SetContextData: a variable-length byte slice containing additional data for the request.
type SetInfoRequest struct {
	WordCount        uint8
	InformationLevel uint16
	Reserved         uint8
	NameLength       uint8
	Flags            uint16
	FileName         string
	SetContextData   []byte
}

func SetInfoRequestParse(data []byte) (*SetInfoRequest, error) {
	// Parse the request structure from the byte slice
	var request SetInfoRequest
	err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &request)
	if err != nil {
		return nil, err
	}

	// Read the variable-length fields from the request
	buf := bytes.NewReader(data[18:])

	// Read the file name
	fileName, err := readVarString(buf, int(request.NameLength), request.Flags&FlagUnicode != 0)
	if err != nil {
		return nil, err
	}
	request.FileName = fileName

	// Read the set context data
	setContextData, err := readVarBytes(buf)
	if err != nil {
		return nil, err
	}
	request.SetContextData = setContextData

	// Return the parsed request
	return &request, nil
}
