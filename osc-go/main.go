package main

import (
	"fmt"
	"math"
)

func inputArray(numberOfItems int, section string) []float64{
	fmt.Printf("Enter your %v: ", section)
	var ans=make([]float64, numberOfItems)
	for i:=0; i<numberOfItems; i++{
		fmt.Scan(&ans[i])
	}
	return ans
}

func getGradeLevel(mark int) string{
	var gradeLevel string
	switch {
		case mark>=80:
			gradeLevel="4"
		case mark>=70:
			gradeLevel="3"
		case mark>=60:
			gradeLevel="2"
		case mark>=50:
			gradeLevel="1"
		default:
			gradeLevel="R"
	}
	return gradeLevel
}

func main() {
	var numberOfItems int;
	fmt.Printf("Enter the number of items: ")
	fmt.Scanln(&numberOfItems)

	var marks=inputArray(numberOfItems, "marks")
	var denominators=inputArray(numberOfItems, "denominators")
	var weights=inputArray(numberOfItems, "weights")

	var totalWeight float64
	for _, v:=range weights{
		totalWeight+=v
	}

	var exactMark float64
	for i:=0; i<numberOfItems; i++{
		var mark=marks[i]/denominators[i]
		var weight=weights[i]/totalWeight

		exactMark += mark * weight
	}

	exactMark*=100
	var finalMark int=int(math.Round(exactMark))
	var gradeLevel=getGradeLevel(finalMark)

	fmt.Printf("Your final mark is %v%%, which is a level %v.\n", finalMark, gradeLevel)
}
