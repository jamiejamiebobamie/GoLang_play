/*
https://exercism.io/my/solutions/e0d2fb529c7844f18056a0767e93597d

Introduction

There are 10 types of people in the world:
Those who understand binary, and those who don't.

You and your fellow cohort of those in the "know"
when it comes to binary decide to come up with a secret "handshake".

1 = wink
10 = double blink
100 = close your eyes
1000 = jump

10000 = Reverse the order of the operations in the secret handshake.
Given a decimal number, convert it to the
appropriate sequence of events for a secret handshake.

Here's a couple of examples:

Given the input 3, the function would return the array
["wink", "double blink"] because 3 is 11 in binary.

Given the input 19, the function would return
the array ["double blink", "wink"] because 19 is 10011 in binary.
Notice that the addition of 16 (10000 in binary)
has caused the array to be reversed.

*/

package main

import "fmt"

func reverseStringArray(inputArray []string)(returnArray []string){
    for i, v := range (inputArray){
        index := len(inputArray)-1-i
        fmt.Println(i,v, index, len(inputArray))
        returnArray[index] = v
    }
    return
}

func decimalToBinarySecret(num int) (binaryString []string) {
    mod := 1
    secret := ""
    remainder := 0
    counter := 0
    reversals := 0

    for {
        if num < 1 {
            break
        }
        if num >= 16 {
            // this should come after the string has been built
            // right now it's always coming before.
            // also the program should count the number of times
            // the number is divided by 16 and reverse the array
            // that many times...
            binaryString = reverseStringArray(binaryString)
            mod = 16
        } else {
             if num >= 8 {
                secret = "jump"
                mod = 8
            } else if num >= 4 {
                secret = "close your eyes"
                mod = 4
            } else if num >= 2 {
                secret = "double blink"
                mod = 2
            } else {
                secret = "wink"
                mod = 1
            }
        }

        remainder = num % mod
        counter = ( num - remainder ) / mod

        if mod != 16 {
            for i:=0; i < counter; i++{
                binaryString = append(binaryString, secret)
            }
        } else {
            reversals += counter
        }

        num = remainder
    }

    // future implementation for array reversal:
    // for i:=0; i < reversals; i++{
    //     binaryString = reverseStringArray(binaryString)
    // }

    return binaryString

}


func main(){
fmt.Println(decimalToBinarySecret(27))
}
