package smb

import (
	"bytes"
	"encoding/binary"
)

// This structure represents a request to write data to a file in the Server Message Block (SMB) protocol.
// It has the following fields:
// 	WordCount: an 8-bit integer indicating the number of 16-bit words in the request.
// 	FID: a 16-bit integer identifying the file to write to.
// 	Offset: a 64-bit integer specifying the offset in the file to write to.
// 	Remaining: a 32-bit integer indicating the number of bytes remaining to be written to the file.
// 	DataLength: a 16-bit integer indicating the length of the DataToWrite field in bytes.
// 	DataOffset: a 16-bit integer specifying the offset in the packet where the DataToWrite field begins.
// 	HighOffset: a 32-bit integer specifying the high 32 bits of the file offset.
// 	Reserved: a 32-bit integer reserved for future use.
// 	Flags: a 16-bit integer containing flags that modify the request.
// 	NameLength: an 8-bit integer indicating the length of the FileName field in bytes.
// 	FileName: a variable-length string containing the name of the file to write to.
// 	WriteContextData: a variable-length byte slice containing additional data for the request.
// 	DataToWrite: a variable-length byte slice containing the data to write to the file.
type WriteRequest struct {
	WordCount        uint8
	FID              uint16
	Offset           uint64
	Remaining        uint32
	DataLength       uint16
	DataOffset       uint16
	HighOffset       uint32
	Reserved         uint32
	Flags            uint16
	NameLength       uint8
	FileName         string
	WriteContextData []byte
	DataToWrite      []byte
}

func WriteRequestParse(data []byte) (*WriteRequest, error) {
	// Parse the request structure from the byte slice
	var request WriteRequest
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

	// Read the write context data
	writeContextData, err := readVarBytes(buf)
	if err != nil {
		return nil, err
	}
	request.WriteContextData = writeContextData

	// Read the data to write
	dataToWrite, err := readVarBytes(buf)
	if err != nil {
		return nil, err
	}
	request.DataToWrite = dataToWrite

	// Return the parsed request
	return &request, nil
}
