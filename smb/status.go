package smb

// Status represents an SMB status code
type Status uint32

const (
	// StatusSuccess indicates a successful operation
	StatusSuccess Status = 0x00000000

	// StatusInvalidParameter indicates an invalid parameter
	StatusInvalidParameter Status = 0x00000057

	// StatusAccessDenied indicates access was denied
	StatusAccessDenied Status = 0x00000005

	// StatusIncorrectPassword indicates an incorrect password
	StatusIncorrectPassword Status = 0x00000056

	// StatusBadNetworkName indicates a bad network name
	StatusBadNetworkName Status = 0x000004B3

	// StatusPathNotFound indicates a path was not found
	StatusPathNotFound Status = 0x00000003

	// StatusInvalidHandle indicates an invalid handle
	StatusInvalidHandle Status = 0x00000006

	// StatusFileExists indicates a file already exists
	StatusFileExists Status = 0x00000050

	// StatusInvalidDevice indicates an invalid device
	StatusInvalidDevice Status = 0x00000047

	// StatusInvalidNetworkResponse indicates an invalid network response
	StatusInvalidNetworkResponse Status = 0x000004B2

	// StatusNotSupported indicates a request is not supported
	StatusNotSupported Status = 0x00000032

	// StatusObjectNameInvalid indicates an invalid object name
	StatusObjectNameInvalid Status = 0x00000033

	// StatusObjectNameNotFound indicates an object name was not found
	StatusObjectNameNotFound Status = 0x00000034

	// StatusObjectNameCollision indicates an object name collision
	StatusObjectNameCollision Status = 0x00000035
)
