package main

import (
    "fmt"
)

func fibMain(num int) (result int) {

    // func fibHelper(number int) (sum int){
    //
    // }

    for i, j := 1,0; i < num; i, j = i+j, i{
        result+=i
        fmt.Println(i)
    }
    return
}


func main() {
    fmt.Println(fibMain(30))
}
