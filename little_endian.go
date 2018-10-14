// +build !386, !x86, !amd64

package intdian

import (
	"encoding/binary"
)

func ByteOrder() binary.ByteOrder {
	return binary.LittleEndian
}
