package main 

import (
	"fmt"
)

type myInt int

type Inccer  interface{
	inc()
}

func (i *myInt) inc(){
	*i = *i + 1
}

func BenchmarkMethod(){
	i := new(myInt)
	incrementMethod(i)
	fmt.Println(*i)
}

func incrementMethod(any Inccer){
	for i:=0; i<10; i++{
		any.inc()
	}
}

func main() {
	BenchmarkMethod()
}

