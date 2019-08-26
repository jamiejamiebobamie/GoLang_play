/*
https://exercism.io/my/solutions/bd69933304614652a21e3bd36bfe0a30

Introduction

Take a nested list and return a single flattened list
with all values except nil/null.

The challenge is to write a function that accepts an
arbitrarily-deep nested list-like structure and
returns a flattened structure without any nil/null values.

For Example

input: [1,[2,3,null,4],[null],5]

output: [1,2,3,4,5]
*/


// In Go the interior lists of a nested list
// all must be of the same type
// (example [...][3]int)
// So the example input would not be recognized in Go:
// input: [1,[2,3,null,4],[null],5]
// output: [1,2,3,4,5]
// the nil vale of an int is 0...

package main

import "fmt"

func flatten(nestedList [3][5]int) (result [15]int) {
    counter := 0
    for _, lists := range(nestedList) {
        for _, v := range(lists){
            result[counter] = v
            counter+=1
        }
    }
    return
}

func main() {
    // the downside of the function / solution has to do with the strongly-
    // typed nature of Go. The input and output are fixed based on the size.
    nestedList := [3][5]int{{1,2,3,0,0},{1,2,3,0,4},{1,2,3,0,4}}
    fmt.Println(nestedList)
    result := flatten(nestedList)
    fmt.Println(result)
}
