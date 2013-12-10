package main

import (
"fmt"
"math"
)

const(

	LRANGE = 0.17
	HRANGE = 0.19
	INCREMENT = 0.000000001
	DELTA = 0.0000001
)

func recurse(start float64, val float64, pos int, pattern [6]int, orbit [6] float64) {

	if ((val < 0.5 && pattern[pos%6] == 0) ||  (val > 0.5 && pattern[pos%6] == 1)){

		if pos == 6 && math.Abs(start-val) < DELTA{

			fmt.Println("One good initial value is: ", start)
			fmt.Println("The corresponding orbit is: ", orbit, ". The sequence repeats after this.")

		}else if pos == 6 && math.Abs(start-val) > DELTA{

			// Dead End

		}else {

			orbit[pos] = val
			recurse(start, math.Mod((2*val),1), pos+1, pattern, orbit)
		}

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

	var orbit [6] float64

	for i := LRANGE; i < HRANGE; i += INCREMENT{

		recurse(i, i, 0, pattern, orbit)
	}

}