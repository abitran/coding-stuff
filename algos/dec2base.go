package main

import (
	"fmt"
)

func main(){

    dec := 255
    fmt.Printf("%s \n", DecToBase(dec, 2))
}

func DecToBase(dec, base int) string{

    s := ""
    for  dec > 0{
        s = fmt.Sprintf("%X%s", dec%base, s)
        dec = dec/base
    }

    return s

}

