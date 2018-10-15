// +build 386 x86 amd64

package intdian

import (
	"encoding/binary"
)

func BigEndian() bool {
	return false
}

func LittleEndian() bool {
	return true
}

func ByteOrder() binary.ByteOrder {
	return binary.LittleEndian
}
