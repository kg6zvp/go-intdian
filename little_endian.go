// +build 386 x86 amd64

package intdian

import (
	"encoding/binary"
)

const (
	Little_Endian = true
	Big_Endian    = false
)

func ByteOrder() binary.ByteOrder {
	return binary.LittleEndian
}
