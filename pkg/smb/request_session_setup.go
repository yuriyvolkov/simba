package smb

import (
	"bytes"
	"encoding/binary"
)

// SessionSetupRequest represents an SMB session setup request
// The WordCount field specifies the number of words in the request, where a word is a 16-bit field.
// The AndXCommand field is used to chain multiple SMB commands together into a single request.
// The AndXReserved field is reserved for future use and should be set to zero.
// The AndXOffset field specifies the offset, in bytes, from the start of the SMB header to the next SMB command in the chain.
// The MaxBufferSize field specifies the maximum size, in bytes, of the buffer that the client can use for the request.
// The MaxMpxCount field specifies the maximum number of pending multiplexed requests that the client can have.
// The VCNumber field specifies the virtual circuit (VC) number that the client is using for the request.
// The SessionKey field is a 32-bit value that is used to identify the session.
// The SecurityBlobLen field specifies the length, in bytes, of the security blob.
// The Reserved field is reserved for future use and should be set to zero.
// The Capabilities field specifies the capabilities of the client, such as support for Unicode strings and large files.
// The SecurityBlob field contains the security blob, which is used to authenticate the client.
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
