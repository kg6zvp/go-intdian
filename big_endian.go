// +build !386, !x86, !amd64

package intdian

import (
	"encoding/binary"
)

const (
	Big_Endian    = true
	Little_Endian = false
)

func ByteOrder() binary.ByteOrder {
	return binary.BigEndian
}
