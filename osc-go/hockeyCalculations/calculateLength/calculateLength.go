package calculateLength

import (
	"fmt"
)

func init(){
	fmt.Println("Stick length calculation program initialized.")
}

func Calculate(height int, position string) int{
	var inch float64=float64(height)*0.4
	var length int
	if(position=="forward"){
		length=int(inch)-8
	}else if(position=="defense"){
		length=int(inch)-6
	}
	return length
}
