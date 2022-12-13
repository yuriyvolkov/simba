package smb

import (
	"bytes"
	"encoding/binary"
)

// ReadRequest structure represents a request to read data from a file in the Server Message Block (SMB) protocol.
// It has the following fields:

//     WordCount: an 8-bit integer indicating the number of 16-bit words in the request.
//     FID: a 16-bit integer identifying the file to read from.
//     Offset: a 64-bit integer specifying the offset in the file to read from.
//     MaxCount: a 16-bit integer indicating the maximum number of bytes to read from the file.
//     MinCount: a 16-bit integer indicating the minimum number of bytes to read from the file.
//     Remaining: a 32-bit integer indicating the number of bytes remaining to be read from the file.
//     HighOffset: a 32-bit integer specifying the high 32 bits of the file offset.
//     Reserved: a 32-bit integer reserved for future use.
//     Flags: a 16-bit integer containing flags that modify the request.
//     NameLength: an 8-bit integer indicating the length of the FileName field in bytes.
//     FileName: a variable-length string containing the name of the file to read from.
//     ReadContextData: a variable-length byte slice containing additional data for the request.
type ReadRequest struct {
	WordCount       uint8
	FID             uint16
	Offset          uint64
	MaxCount        uint16
	MinCount        uint16
	Remaining       uint32
	HighOffset      uint32
	Reserved        uint32
	Flags           uint16
	NameLength      uint8
	FileName        string
	ReadContextData []byte
}

func ReadRequestParse(data []byte) (*ReadRequest, error) {
	// Parse the request structure from the byte slice
	var request ReadRequest
	err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &request)
	if err != nil {
		return nil, err
	}

	// Read the variable-length fields from the request
	buf := bytes.NewReader(data[17:])

	// Read the file name
	fileName, err := readVarString(buf, int(request.NameLength), request.Flags&FlagUnicode != 0)
	if err != nil {
		return nil, err
	}
	request.FileName = fileName

	// Read the read context data
	readContextData, err := readVarBytes(buf)
	if err != nil {
		return nil, err
	}
	request.ReadContextData = readContextData

	// Return the parsed request
	return &request, nil
}
