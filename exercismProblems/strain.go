/*

Introduction

Implement the keep and discard operation on collections.
Given a collection and a predicate on the collection's elements,
keep returns a new collection containing those elements where
the predicate is true, while discard returns a new collection
containing those elements where the predicate is false.

For example, given the collection of numbers:

1, 2, 3, 4, 5
And the predicate:

is the number even?
Then your keep operation should produce:

2, 4
While your discard operation should produce:

1, 3, 5
Note that the union of keep and discard is all the elements.

The functions may be called keep and discard, or they may
need different names in order to not clash with existing
functions or concepts in your language.

*/

package main

import (
    "fmt"
)

func keep(inputList []int, condition func(int) bool) (outputList []int){
    for i:= 0; i < len(inputList); i++{
        if condition(inputList[i]) == true {
            outputList = append(outputList, inputList[i])
        }
    }
        return
}

func discard(inputList []int, condition func(int) bool) (outputList []int){
    for i:= 0; i < len(inputList); i++{
        if condition(inputList[i]) == false {
            outputList = append(outputList, inputList[i])
        }
    }
        return
}

func main(){
    var testList []int = []int{1,2,3,4}
    // the condition must be passed in
    // as an anonymous function that returns a bool... D,:
    fmt.Println(keep(testList, func(num int)bool{ return num % 2 == 0}))
    fmt.Println(discard(testList, func(num int)bool{ return num % 2 == 0}))
}
