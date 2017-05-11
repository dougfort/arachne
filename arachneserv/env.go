package main

import "os"

const (
	envAddress     = "ARACHNEADDRESS"
	defaultAddress = ":10000"
)

func getAddressFromEnv() string {
	var address string

	if address = os.Getenv(envAddress); address == "" {
		address = defaultAddress
	}

	return address
}
