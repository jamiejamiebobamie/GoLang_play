/*
https://exercism.io/my/solutions/e35b66a36ba84c8b90861b1a9c35e43f
Introduction

Determine if a sentence is a pangram.
A pangram (Greek: παν γράμμα, pan gramma, "every letter")
is a sentence using every letter of the alphabet at least once.
The best known English pangram is:

The quick brown fox jumps over the lazy dog.

The alphabet used consists of ASCII letters a to z, inclusive,
and is case insensitive. Input will not contain non-ASCII symbols.

*/

package main

import "fmt"

func pangramTest(input string) bool {
    var LOOKUP map[string]int = map[string]int{"a":0,"b":0,"c":0,"d":0,
        "e":0,"f":0,"g":0, "h":0,"i":0,"j":0,"k":0,"l":0,"m":0,
        "n":0,"o":0,"p":0,"q":0,"r":0, "s":0,"t":0,"u":0,"v":0,
        "w":0,"x":0,"y":0,"z":0}

    for _,v := range(input){
        _, ok := LOOKUP[string(v)]
        if ok {
            LOOKUP[string(v)] += 1
        }
    }

    for _, value := range LOOKUP{
        if value == 0 {
            return false
        }
    }

    return true

}

func main() {
    // false
    test_input1 := "a man a plan can a henry hello"
    test1 := pangramTest(test_input1)
    fmt.Println(test_input1,test1)

    // true
    test_input2 := "The quick brown fox jumps over the lazy dog."
    test2 := pangramTest(test_input2)
    fmt.Println(test_input2,test2)
}
