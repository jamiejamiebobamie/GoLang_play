/*
https://exercism.io/my/solutions/2fd16199ba04440f867be3a637ba6c7e

Introduction

Reverse a string

For example: input: "cool" output: "looc"

*/


package main

import (
    "fmt"
    "strings"
)

func reverse(str string) (reversed string) {

    var b =make([]string, len(str))

    for i, v := range str{
        index := len(str) - 1 - i
        b[index] = string(v)
    }

    reversed = strings.Join(b,"")

    return
}


func main(){
    a := "cool"
    b := reverse(a)
    fmt.Println(b)
}
