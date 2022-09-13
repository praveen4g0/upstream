package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Here is a random %T ) : %d\n", big.NewInt(int64(rand.Int())), big.NewInt(int64(rand.Int()))) //nolint:gosec
}
