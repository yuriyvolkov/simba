package smb

import (
	"bytes"
	"encoding/binary"
)

// RenameRequest structure represents a request to rename a file in the Server Message Block (SMB) protocol.
// It has the following fields:

//     WordCount: an 8-bit integer indicating the number of 16-bit words in the request.
//     SearchAttributes: a 16-bit integer containing flags specifying the search attributes for the file to rename.
//     OldNameLength: an 8-bit integer indicating the length of the OldFileName field in bytes.
//     NewNameLength: an 8-bit integer indicating the length of the NewFileName field in bytes.
//     Flags: a 16-bit integer containing flags that modify the request.
//     OldFileName: a variable-length string containing the current name of the file to rename.
//     NewFileName: a variable-length string containing the new name of the file.
//     RenameContextData: a variable-length byte slice containing additional data for the request.
type RenameRequest struct {
	WordCount         uint8
	SearchAttributes  uint16
	OldNameLength     uint8
	NewNameLength     uint8
	Flags             uint16
	OldFileName       string
	NewFileName       string
	RenameContextData []byte
}

func RenameRequestParse(data []byte) (*RenameRequest, error) {
	// Parse the request structure from the byte slice
	var request RenameRequest
	err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &request)
	if err != nil {
		return nil, err
	}

	// Read the variable-length fields from the request
	buf := bytes.NewReader(data[20:])

	// Read the old file name
	oldFileName, err := readVarString(buf, int(request.OldNameLength), request.Flags&FlagUnicode != 0)
	if err != nil {
		return nil, err
	}
	request.OldFileName = oldFileName

	// Read the new file name
	newFileName, err := readVarString(buf, int(request.NewNameLength), request.Flags&FlagUnicode != 0)
	if err != nil {
		return nil, err
	}
	request.NewFileName = newFileName

	// Read the rename context data
	renameContextData, err := readVarBytes(buf)
	if err != nil {
		return nil, err
	}
	request.RenameContextData = renameContextData

	// Return the parsed request
	return &request, nil
}
