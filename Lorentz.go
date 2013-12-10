package main

import (
"fmt"
"math"
)



func recurse(start float64, val float64, pos int, pattern [6]int) {

	if val < 0.5 && pattern[pos%6] != 0 {


	}else if val > 0.5 && pattern[pos%6] != 1{



	}else if pos == 6 && math.Abs(start-val) > 0.0000001{


	}else if pos == 6 && math.Abs(start-val) < 0.0000001{

		fmt.Println("The value is: ", start)

	}else {

		//fmt.Println("Start value: ", start, " Current value at pos ", pos, " is: ", val)
		recurse(start, math.Mod((2*val),1), pos+1, pattern)
	}

	
} 

func main(){


	var pattern [6] int
	pattern[0] = 0
	pattern[1] = 0
	pattern[2] = 1
	pattern[3] = 0
	pattern[4] = 1
	pattern[5] = 1

	for i:=0.17; i < 0.19; i+=0.000000001{
		//fmt.Println("testing ", i)
		recurse(i, i, 0, pattern)
	}

}