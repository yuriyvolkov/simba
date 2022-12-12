package smb

// Command represents an SMB command
type Command uint16

const (
	// CommandNegotiate indicates a negotiate command
	CommandNegotiate Command = 0x72

	// CommandSessionSetup indicates a session setup command
	CommandSessionSetup Command = 0x73

	// CommandTreeConnect indicates a tree connect command
	CommandTreeConnect Command = 0x75

	// CommandCreate indicates a create command
	CommandCreate Command = 0x6D

	// CommandClose indicates a close command
	CommandClose Command = 0x6E

	// CommandWrite indicates a write command
	CommandWrite Command = 0x2B

	// CommandRead indicates a read command
	CommandRead Command = 0x2E

	// CommandTreeDisconnect indicates a tree disconnect command
	CommandTreeDisconnect Command = 0x71

	// CommandLogoff indicates a logoff command
	CommandLogoff Command = 0x74
)
