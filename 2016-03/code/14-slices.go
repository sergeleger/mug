package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func slicePtr(s []int) string {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	return fmt.Sprintf("0x%X", hdr.Data)
}

func main() {
	var ids = make([]int, 0, 10) // HL
	for i := 0; i < 100; i++ {
		fmt.Println(len(ids), cap(ids)) // HL
		ids = append(ids, i)
	}

	// The in-memory location of the array changes as Go allocates more memory for the
	// growing slice.
	// slicePtr(ids)
}
