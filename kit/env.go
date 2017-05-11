package kit

import "os"

const (
	envAddress     = "ARACHNEADDRESS"
	defaultAddress = ":10000"
)

// GetAddressFromEnv gets the arachne server address
func GetAddressFromEnv() string {
	var address string

	if address = os.Getenv(envAddress); address == "" {
		address = defaultAddress
	}

	return address
}
