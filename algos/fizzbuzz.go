package main

import (
    "fmt"
)

func main(){

    FizzBuzz(50)

}

func FizzBuzz(number int){
    for i := 1; i <=number; i++ {
        if (i % 3 == 0){
            if (i % 5 == 0){
                fmt.Printf("Fizz Buzz, ")
                continue
            }
            fmt.Printf("Fizz")
        }else if(i % 5 == 0){
            if (i % 3 == 0){
                fmt.Printf("Fizz Buzz, ")
                continue
            }
            fmt.Printf("Buzz")
        }else{
            fmt.Printf("%d", i)
        }
        fmt.Printf(", ")
    }
    fmt.Printf("\n")
}


