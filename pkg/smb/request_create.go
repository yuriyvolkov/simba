package smb

import (
	"bytes"
	"encoding/binary"
)

type CreateRequest struct {
	WordCount          uint8
	NameLength         uint8
	Flags              uint32
	RootDirectoryFID   uint32
	DesiredAccess      uint32
	AllocationSize     uint64
	ExtFileAttributes  uint32
	ShareAccess        uint32
	CreateDisposition  uint32
	CreateOptions      uint32
	ImpersonationLevel uint32
	SecurityFlags      uint8
	ByteCount          uint16
	FileName           string
	CreateContextData  []byte
}

func CreateRequestParse(data []byte) (*CreateRequest, error) {
	// Parse the request structure from the byte slice
	var request CreateRequest
	err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &request)
	if err != nil {
		return nil, err
	}

	// Read the variable-length fields from the request
	buf := bytes.NewReader(data[24:])

	// Read the file name
	fileName, err := readVarString(buf, int(request.NameLength), request.Flags&uint32(FlagUnicode) != 0)
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
