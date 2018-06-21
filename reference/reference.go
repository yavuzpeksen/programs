package main 

import (
	"fmt"
)

type Company struct {
    Name string
    Workers [](*Worker)
}

type Worker struct {
    Name string
    Other []int
}

func (cmp *Company) AddWorker(name string) *Worker {
    wrk := Worker{Name: name}
    cmp.Workers = append(cmp.Workers, &wrk)
    return &wrk
}

func main() {

	cmp := Company{}
    cmp.Name = "Google"
    
    //wrk and wrk2 are pointers
    wrk := cmp.AddWorker("Joe")
    wrk2 := cmp.AddWorker("Doe")
    for i := 1; i <= 10; i++ {
        wrk.Other = append(wrk.Other, i)
        wrk2.Other = append(wrk2.Other, i)

    }
    
    for _,ptr:= range cmp.Workers{
    	fmt.Println(*ptr)
    }
    
    for i:= range cmp.Workers{
    	fmt.Println(*(cmp.Workers[i]))
    }

}

