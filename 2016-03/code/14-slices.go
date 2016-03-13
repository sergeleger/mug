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
	var ids = make([]int, 0, 10)
	for i := 0; i < 100; i++ {
		fmt.Println(len(ids), cap(ids))
		ids = append(ids, i)
	}
}
