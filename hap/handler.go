package hap

import (
	"io"
	"net"
	"net/url"

	"github.com/brutella/hc/util"
)

// A ContainerHandler abstracts request/response communication
type ContainerHandler interface {
	Handle(util.Container) (util.Container, error)
}

// A PairVerifyHandler is a ContainerHandler which negotations a shared key.
type PairVerifyHandler interface {
	ContainerHandler
	SharedKey() [32]byte
}

// A AccessoriesHandler returns a list of accessories as json.
type AccessoriesHandler interface {
	HandleGetAccessories(r io.Reader) (io.Reader, error)
}

// A CharacteristicsHandler handles get and update characteristic.
type CharacteristicsHandler interface {
	HandleGetCharacteristics(url.Values, net.Conn) (io.Reader, error)
	HandleUpdateCharacteristics(io.Reader, net.Conn) error
}

// IdentifyHandler calls Identify() on accessories.
type IdentifyHandler interface {
	IdentifyAccessory()
}
