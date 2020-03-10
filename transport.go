package hc

import (
	"hc/accessory"
)

// Transport provides accessories over a network.
type Transport interface {
	AddAccessory(a *accessory.Accessory) error
	RemoveAccessory(a *accessory.Accessory)
	// Start starts the transport
	Start()

	// Stop stops the transport
	// Use the returned channel to wait until the transport is fully stopped.
	Stop() <-chan struct{}
}
