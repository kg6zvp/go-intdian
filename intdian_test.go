package intdian

import (
	"testing"
	"unsafe"
)

const INT_SIZE int = int(unsafe.Sizeof(0))

func TestEndianness(t *testing.T) {
	var i int = 0x1
	bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
	if bs[0] == 0 { // big endian
		if LittleEndian() {
			t.Fatalf("")
		}
		if !BigEndian() {
			t.Fatalf("")
		}
	} else { // little endian
		if BigEndian() {
			t.Fatalf("")
		}
		if !LittleEndian() {
			t.Fatalf("")
		}
	}
}
