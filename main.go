package main

import "fmt"


func main (){
	var marks int
	fmt.Println("Please enter your marks")
	fmt.Scanf("%d\n", &marks)

	if upper := 70; marks >= upper {
		fmt.Printf("You scored %v, exellent job", marks)
	}else if average := 40; marks >= average{
		fmt.Printf("You scored %v, Good work", marks)
	}else{
		fmt.Printf("You scored %v, Failed, try again", marks)
	}
}