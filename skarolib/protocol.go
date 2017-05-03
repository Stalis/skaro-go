// protocol
package skaro

import (
	"log"

	"github.com/ugorji/go/codec"
)

type PacketChain struct {
	chainType string
	data      []Packet
}

type Packet map[string]interface{}

func (pack *Packet) Build() ([]byte, error) {
	log.Print("Building packet...")
	var (
		output []byte = make([]byte, 0, 64)
		err    error

		handler codec.Handle   = new(codec.CborHandle)
		encoder *codec.Encoder = codec.NewEncoderBytes(&output, handler)
	)

	err = encoder.Encode(pack)

	return output, err
}

func (pack *Packet) Read(input []byte) error {
	log.Println("Reading packet...")
    var err error
    handler codec.Handle = new(codec.CborHandle)
    decoder *codec.Encoder = codec.NewDecoderBytes(&pack, handler)

    err = decoder.Decode(input)
    return err
}
