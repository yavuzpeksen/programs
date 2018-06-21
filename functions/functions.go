package main 

import (
	"fmt"
	"citybydistance"
)

var x int = 5

func returnTriple(value int) (int,int,int){
	return value, value+1, value+2
	
}

func acceptMultipleParams(args ...int) int{
	var total int
	for _, v:= range args{
		total += v
	}
	return total
	
}

func acceptSliceParam(arr []int){
	for i:=0; i<len(arr); i++{
		fmt.Print(arr[i],",")
	}
	fmt.Println()
	
	for _, v:= range arr{
		fmt.Print(v, ",")
	}
}

func main() {
	
	//This has to be dynamic array
	intArr := []int{5,2,3,4}
	citybydistance.FindClosestCity("Atlanta")
	fmt.Println(returnTriple(10))
	fmt.Println(acceptMultipleParams(1,2,3,4,5,6,7,8))
	fmt.Println(acceptMultipleParams(intArr...))

	var xPtr *[]int = &intArr
	yPtr := new(int)
	*yPtr = 5
	var zPtr *int
	//*zPtr = *yPtr
	if((*xPtr)[0] ==  *yPtr){
		fmt.Println("Pointing the same value")
	}
	intArr = append(intArr, 6,7)
	fmt.Println(len(intArr))
	fmt.Println("yPtr, zPtr", &yPtr, &zPtr)
	
	acceptSliceParam(intArr)
	
}


