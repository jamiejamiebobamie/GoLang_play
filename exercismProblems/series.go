/*
https://exercism.io/my/solutions/c4d45148458a46eebc794c82e05a2bac

Introduction

Given a string of digits, output all the contiguous substrings
of length n in that string in the order that they appear.

For example, the string "49142" has the following 3-digit series:

"491"
"914"
"142"
And the following 4-digit series:

"4914"
"9142"
And if you ask for a 6-digit series from a 5-digit string,
you deserve whatever you get.

Note that these series are only required to occupy adjacent
positions in the input; the digits need not be numerically consecutive.

*/


package main

import "fmt"

func series(inputString string, outputLength int) (outputArray []string) {

    lengthOfInputString := len(inputString)
    // if the subsets are supposed to
    // be of a larger length than the input string
    // or if the outputLength is less than or equal to 0
    if lengthOfInputString < outputLength ||  outputLength <= 0 {
        return
    }

    // if the subset length is the length of the array
    if lengthOfInputString == outputLength {
        outputArray = append(outputArray, inputString)
        return
    }

    // otherwise iterate through the string and append the subsets
    for i := 0; i < lengthOfInputString; i++{
        if i+outputLength < lengthOfInputString+1 {
            outputArray = append(outputArray, inputString[i:i+outputLength])
        }
    }
    return
}

func main(){
    fmt.Println(series("334", 2))
}
