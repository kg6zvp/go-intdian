// +build !386, !x86, !amd64

package intdian

import (
	"encoding/binary"
)

func BigEndian() bool {
	return true
}

func LittleEndian() bool {
	return false
}

func ByteOrder() binary.ByteOrder {
	return binary.BigEndian
}
