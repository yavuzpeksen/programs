package main 

import (
	"fmt"
)

func main() {
	
	exerciseDefer()
}


func exerciseDefer(){
	
	var x int = 10
	defer executeSecond(x)
	executeFirst()
	
	defer func(){
		
		str := recover()	
		fmt.Println(str)
	}()
	x = alterOnDefer()
	fmt.Println(x)
	panic("Error Message")

}

func executeFirst(){
	fmt.Println("Executing first")
}
func executeSecond(x int){
	fmt.Println("Executing second, x:", x)
}

func alterOnDefer() (result int){
	defer func(){
		result = -1
	}()
	return result
}