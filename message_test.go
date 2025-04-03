package message

import (
	"fmt"
	"testing"
)

const VERSION byte = 0x01

func TestDecode(t *testing.T) {
	data := []byte("2")

	//{STX:2 Version:1 Pin:5 IO:1 RW:0 Length:6 Data:[50 102 97 108 115 101] Checksum:0 ETX:3}
	message, err := CreateMessage(&Builder{
		Version: VERSION,
		Pin:     "D13",
		IO:      Output,
		RW:      Read,
		Data:    data,
	})

	if err != nil {
		t.Errorf("Error creating message: %s", err)
	}

	encoded := message.Encode()

	decoded, err := Decode(encoded)

	if err != nil {
		t.Error(err)
	}

	if decoded.Checksum != message.Checksum {
		t.Errorf("Checksum mismatch. Expected: %d, got: %d", message.Checksum, decoded.Checksum)
	}

	fmt.Println(len(encoded))

	fmt.Printf("Message: %+v\n", message)
	fmt.Printf("Encoded binary: %b\n", encoded)
	fmt.Printf("Encoded hex: %x\n", encoded)
	fmt.Printf("Decoded: %+v\n", decoded)
}
