package main

import (
	"fmt"

	"github.com/marcosalvalaggio-bip/lab/src/power"
)

func main() {
	fmt.Println("Hello world")
	Tool()
	fmt.Println(Pub)
	fmt.Println(power.Power)
	power.PowTest()
	df := Dataset()
	fmt.Println("test returing:")
	fmt.Println(df)
}
