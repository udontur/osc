package main

import (
	"fmt"
	"osc-go/hockeyCalculations/calculateFlex"
	"osc-go/hockeyCalculations/calculateLength"
	// "math"
)

func main() {
	var weight int=47;
	var height int=162;
	var position string="forward"
	flex:=calculateFlex.Calculate(weight)
	length:=calculateLength.Calculate(height, position)
	fmt.Println("Your ideal flex should be", flex, "flex.")
	fmt.Printf("Your stick length should be %v\".\n", length)
}
