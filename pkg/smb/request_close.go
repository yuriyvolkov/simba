package smb

import (
	"bytes"
	"encoding/binary"
)

// This structure represents a request to close a file in the Server Message Block (SMB) protocol.
// It has the following fields:
//     WordCount: an 8-bit integer indicating the number of 16-bit words in the request.
//     FID: a 16-bit integer identifying the file to close.
//     LastWriteTime: a 64-bit integer representing the last write time of the file.
//     Reserved: a 32-bit integer reserved for future use.
//     ByteCount: a 16-bit integer that indicates the total number of bytes in the CloseRequest structure. This includes the length of both the fixed-length and variable-length fields of the structure. The ByteCount field is used by the SMB protocol to determine the end of the CloseRequest structure in a packet. It allows the receiver of the packet to parse the CloseRequest structure correctly, even if the structure includes variable-length fields with different lengths in different packets.
//     NameLength: an 8-bit integer indicating the length of the FileName field in bytes.
//     Flags: a 16-bit integer that contains various flags that control the behavior of the CloseRequest command. The exact meaning and use of these flags varies depending on the specific version of the SMB protocol in use. In the CloseRequestParse function, the Flags field is used to determine if the FileName field is encoded in Unicode or ASCII format.
//     FileName: a variable-length string containing the name of the file to close.
//     CreateContextData: a variable-length byte slice containing additional data for the request.
type CloseRequest struct {
	WordCount        uint8
	FID              uint16
	LastWriteTime    uint64
	Reserved         uint16
	NameLength       uint16
	ByteCount        uint16
	Flags            uint16
	FileName         string
	CloseContextData []byte
}

func (c *CloseRequest) Marshal() ([]byte, error) {
	// Create a byte buffer for the request
	buf := bytes.NewBuffer([]byte{})

	// Write the fixed-length fields to the buffer
	binary.Write(buf, binary.LittleEndian, c.WordCount)
	binary.Write(buf, binary.LittleEndian, c.FID)
	binary.Write(buf, binary.LittleEndian, c.LastWriteTime)
	binary.Write(buf, binary.LittleEndian, c.Reserved)
	binary.Write(buf, binary.LittleEndian, uint16(len(c.FileName)))

	// Write the variable-length fields to the buffer
	binary.Write(buf, binary.LittleEndian, c.ByteCount)
	binary.Write(buf, binary.LittleEndian, c.Flags)
	writeVarString(buf, c.FileName)
	writeVarBytes(buf, c.CloseContextData)

	// Return the serialized request
	return buf.Bytes(), nil
}

func CloseRequestParse(data []byte) (*CloseRequest, error) {
	// Parse the request structure from the byte slice
	var request CloseRequest
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
	closeContextData, err := readVarBytes(buf)
	if err != nil {
		return nil, err
	}
	request.CloseContextData = closeContextData

	// Return the parsed request
	return &request, nil
}
