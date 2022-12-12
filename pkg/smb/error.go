package smb

// ErrorClass represents an SMB error class
type ErrorClass uint8

const (
	// ErrorClassSuccess indicates a successful operation
	ErrorClassSuccess ErrorClass = 0

	// ErrorClassDOS indicates a DOS error
	ErrorClassDOS ErrorClass = 0x01

	// ErrorClassServer indicates a server error
	ErrorClassServer ErrorClass = 0x02

	// ErrorClassHardware indicates a hardware error
	ErrorClassHardware ErrorClass = 0x03

	// ErrorClassProto indicates a protocol error
	ErrorClassProto ErrorClass = 0x04
)

// ErrorResponse represents an SMB error response
type ErrorResponse struct {
	ErrorClass    ErrorClass
	Reserved      uint8
	ErrorCode     uint32
	ErrorReserved uint32
	ErrorString   string
}
