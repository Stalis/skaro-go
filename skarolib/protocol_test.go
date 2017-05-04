package skaro

import (
	"reflect"
	"testing"
)

var (
	packs = []Packet{
		Packet(map[string]interface{}{
			"a": "A",
			"b": "B",
			"c": "C",
			"d": "D",
			"e": "E",
		}),
	}

	cbors = [][]byte{
		[]byte(`pWFhYUFhYmFCYWNhQ2FkYURhZWFF`),
	}
)

func TestPacket_Read(t *testing.T) {
	type args struct {
		input []byte
	}

	inputs := make([]args, len(cbors))
	empties := make([]Packet, len(cbors))
	for i := 0; i < len(empties); i++ {
		empties[i] = Packet{}
		inputs[i] = args{cbors[i]}
	}

	tests := []struct {
		name    string
		pack    *Packet
		args    args
		want    Packet
		wantErr bool
	}{
		{
			"First",
			&empties[0],
			inputs[0],
			packs[0],
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pack.Read(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Packet.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Packet.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPacket_Build(t *testing.T) {
	packs := []Packet{
		Packet(map[string]interface{}{
			"a": "A",
			"b": "B",
			"c": "C",
			"d": "D",
			"e": "E",
		}),
		Packet(map[string]interface{}{"string": "skaro", "number": 11, "bool": true}),
	}
	wants := [][]byte{
		[]byte(`pWFhYUFhYmFCYWNhQ2FkYURhZWFF`),
	}

	tests := []struct {
		name    string
		pack    *Packet
		want    []byte
		wantErr bool
	}{
		{
			"First",
			&packs[0],
			wants[0],
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pack.Build()
			if (err != nil) != tt.wantErr {
				t.Errorf("Packet.Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Packet.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
