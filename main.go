package main

import "fmt"

// word maligned
// go-lint will complain "struct of size 48 bytes could be of size 40 bytes (maligned)"
type CarA struct {
	Doors int32  // 4 bytes
	Brand string // 16 bytes
	IsNew bool   // 1 byte
	Model string // 16 bytes
}

// word better aligned
// go-lint won't complain anymore
type CarB struct {
	IsNew bool   // 1 byte
	Doors int32  // 4 bytes
	Brand string // 16 bytes
	Model string // 16 bytes
}

func main() {

	carA := CarA{5, "Tesla", true, "Model 3"}
	carB := CarB{true, 5, "Tesla", "Model S"}

	fmt.Println(carA)
	fmt.Println(carB)
}
