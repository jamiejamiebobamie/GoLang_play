/*

Introduction

Implement the accumulate operation, which,
given a collection and an operation to perform on each element
of the collection, returns a new collection containing the result
of applying that operation to each element of the input collection.

Given the collection of strings:

"cat", "Dog", "b4t", "gO"
And the operation:

upcase a string
Your code should be able to produce the collection of strings:

"CAT", "DOG", "B4T, "GO"

*/

package main


import (
    "fmt"
    "strings"
)

func accumulate(inputList []string,
    method func(string) string) (outputList []string){
    for i := 0; i < len(inputList); i++{
        outputList = append(outputList, method(inputList[i]))
    }
    return
}

func main(){
    strings_and_things := make([]string, 5)
    strings_and_things = []string{"dude", "Man", "123", "6dgs", "HOWDY!"}

    fmt.Println(accumulate(strings_and_things, strings.ToUpper))
    fmt.Println(accumulate(strings_and_things, strings.ToLower))
}
