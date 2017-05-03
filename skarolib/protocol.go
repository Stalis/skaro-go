// protocol
package skaro

import (
	"log"

	"errors"

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
	var (
		decoded interface{}
		err     error = codec.NewDecoderBytes(input, new(codec.CborHandle)).Decode(&decoded)
	)
	if err != nil {
		return err
	}
	switch packetTyped := decoded.(type) {
	case map[string]interface{}:
		*pack = packetTyped
	default:
		log.Fatalln("Wrong packet recieved:", input)
		return errors.New("Wrong packet")
	}
	return err
}
