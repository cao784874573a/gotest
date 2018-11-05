package main

import (
	
	"fmt"
	
	
	
	
)



func main() {

	//
	//
	//
	//

	 var a1 string= "1101010000100110011010101100011000100111100101001011011000110010"
	 
	 var a2 string= "1100100010000111110110101110100010011011101110001101001101101010"
	
	 var n int=0
	

	for i := 0; i < len(a1); i++ {

		if a1[i]^a2[i]==1{
			
			n++
		}
		
	}

	
	fmt.Println(n)

	

	 	
	 	

	 

	
	 


}
