package smb

import (
	"bytes"
	"encoding/binary"
)

// SessionSetupRequest represents an SMB session setup request
type SessionSetupRequest struct {
	WordCount       uint8
	AndXCommand     Command
	AndXReserved    uint8
	AndXOffset      uint16
	MaxBufferSize   uint16
	MaxMpxCount     uint16
	VCNumber        uint16
	SessionKey      uint32
	SecurityBlobLen uint16
	Reserved        uint16
	Capabilities    uint32
	SecurityBlob    []byte
}

// SessionSetupRequestParse parses an SMB session setup request
func SessionSetupRequestParse(data []byte) (*SessionSetupRequest, error) {
	// Create a new bytes reader
	r := bytes.NewReader(data)

	// Read the word count
	var wordCount uint8
	if err := binary.Read(r, binary.LittleEndian, &wordCount); err != nil {
		return nil, err
	}

	// Read the data block size
	var dataBlockSize uint16
	if err := binary.Read(r, binary.LittleEndian, &dataBlockSize); err != nil {
		return nil, err
	}

	// Read the data block
	dataBlock := make([]byte, dataBlockSize)
	if err := binary.Read(r, binary.LittleEndian, dataBlock); err != nil {
		return nil, err
	}

	// Read the byte count
	var byteCount uint16
	if err := binary.Read(r, binary.LittleEndian, &byteCount); err != nil {
		return nil, err
	}

	// Read the padding
	padding := make([]byte, byteCount)
	if err := binary.Read(r, binary.LittleEndian, padding); err != nil {
		return nil, err
	}

	// FIXME Create the request
	request := &SessionSetupRequest{
		// DataBlockSize: dataBlockSize,
		// DataBlock:     dataBlock,
		// ByteCount:     byteCount,
		// Padding:       padding,
	}

	return request, nil
}
