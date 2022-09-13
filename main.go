package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

func main() {
	fmt.Printf("Here is a random %T ) : %d\n", big.NewInt(int64(rand.Int())), big.NewInt(int64(rand.Int())))
}
