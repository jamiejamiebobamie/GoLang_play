package main

import (
    "fmt"
    _"reflect"
    _"math"
)

func main(){

    a := [...]int{1,2,3,4,5,6,7,8}

    /*
    // uses O(n) space to copy the array:
            var b [len(a)]int
            fmt.Println(a,b)
            for i, v := range a{
                index := len(a) - 1 - i
                b[index] = v
            }
            fmt.Println(a,b)
    */

    // O(1) space, O(n) time
            r := len(a) / 2
            for i:= 0 ; i < r; i++{
                a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
            }
}
