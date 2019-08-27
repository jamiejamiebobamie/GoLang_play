/*
https://exercism.io/my/solutions/ca5d698ba8404cfeab3a39c657595bfb

Introduction

Given a string of digits, calculate the largest product for a contiguous
substring of digits of length n.

For example, for the input '1027839564', the largest product for
a series of 3 digits is 270 (9 * 5 * 6),and the largest product
for a series of 5 digits is 7560 (7 * 8 * 3 * 9 * 5).

Note that these series are only required to occupy adjacent
positions in the input; the digits need not be numerically consecutive.

For the input '73167176531330624919225119674426574742355349194934',
the largest product for a series of 6 digits is 23520.

*/

package main

import (
    "fmt"
    "strconv"
    _"reflect"

    )

func largestSub(input string, n int) (maxInt, maxIndex int){
    var product int

    for i := 0; i < len(input); i++ {

        product = 1
        for j:= i; j < n+i; {

            if j < len(input){
            number := string(input[j])
            inte, ok := strconv.Atoi(number)

            if ok == nil { product = product * inte }
            // fmt.Println(inte, maxInt, i , j)
            }
            j+=1

        }

        if maxInt < product {
            maxInt = product
            maxIndex = i
        }
    }
    return
}

func main(){
    test := "73167176531330624919225119674426574742355349194934"
    fmt.Println(largestSub(test, 6))
}
