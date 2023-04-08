package main

import (
	"fmt"
	"math"
)

func main() {
	var index int8 = 15
	var bigIndex int32
	bigIndex = int32(index)

	fmt.Println(bigIndex)
	fmt.Printf("%T \n", bigIndex)

	var a int32 = 22
	var b uint64
	b = uint64(a)

	fmt.Println(b)
	fmt.Printf("%T \n", b)

	var big int64 = 129
	var little = int8(big)
	fmt.Println(little)

	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxUint8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxUint16)
}
