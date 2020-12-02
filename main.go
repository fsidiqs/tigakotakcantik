package main

import (
	"fmt"
	"kotak/tigakotakcantik"
)

func main() {
	mykotak1 := tigakotakcantik.NewTigaKotakCantik(6, 1, 3000)
	permenEaten := mykotak1.PerformAdjustment()
	fmt.Println(permenEaten)
}
