package calculateFlex

import (
	"fmt"
)

func init(){
	fmt.Println("Flex calculation program initialized.")
}

func Calculate(weight int) (flex int){
	var lbs float64=float64(weight)*2.20462
	flex=int(lbs/2.20462)
	return
}
