package main

import (
    "fmt"
)

func main(){
    list := []int{1,2,4,8}
    fmt.Printf("The sum is %d\n", Sum(list))

}

func Sum(list []int) int{

    if len(list) == 0{
        return 0
    }
    return list[0] + Sum(list[1:])
    
}

