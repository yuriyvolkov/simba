package smb

import (
	"bytes"
	"encoding/binary"
)

// DeleteRequest structure represents a request to delete a file in the Server Message Block (SMB) protocol.
// It has the following fields:

//     WordCount: an 8-bit integer indicating the number of 16-bit words in the request.
//     SearchAttributes: a 16-bit integer containing flags specifying the search attributes for the file to delete.
//     NameLength: an 8-bit integer indicating the length of the FileName field in bytes.
//     Flags: a 16-bit integer containing flags that modify the request.
//     FileName: a variable-length string containing the name of the file to delete.
//     DeleteContextData: a variable-length byte slice containing additional data for the request.
type DeleteRequest struct {
	WordCount         uint8
	SearchAttributes  uint16
	NameLength        uint8
	Flags             uint16
	FileName          string
	DeleteContextData []byte
}

func DeleteRequestParse(data []byte) (*DeleteRequest, error) {
	// Parse the request structure from the byte slice
	var request DeleteRequest
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

	// Read the delete context data
	deleteContextData, err := readVarBytes(buf)
	if err != nil {
		return nil, err
	}
	request.DeleteContextData = deleteContextData

	// Return the parsed request
	return &request, nil
}
