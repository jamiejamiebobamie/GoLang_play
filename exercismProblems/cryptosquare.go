/*
https://exercism.io/my/solutions/3241c108ac374413b5cf774e01c2861e

Introduction

Implement the classic method for composing secret messages called a square code.

Given an English text, output the encoded version of that text.

First, the input is normalized: the spaces and
punctuation are removed from the English text and the message is downcased.

Then, the normalized characters are broken into rows.
These rows can be regarded as forming a rectangle
when printed with intervening newlines.

For example, the sentence

"If man was meant to stay on the ground, god would have given us roots."
is normalized to:

"ifmanwasmeanttostayonthegroundgodwouldhavegivenusroots"
The plaintext should be organized in to a rectangle.
The size of the rectangle (r x c) should be decided by the
length of the message, such that c >= r and c - r <= 1,
where c is the number of columns and r is the number of rows.

Our normalized text is 54 characters long,
dictating a rectangle with c = 8 and r = 7:

"ifmanwas"
"meanttos"
"tayonthe"
"groundgo"
"dwouldha"
"vegivenu"
"sroots  "
The coded message is obtained by reading down the columns going left to right.

The message above is coded as:

"imtgdvsfearwermayoogoanouuiontnnlvtwttddesaohghnsseoau"
Output the encoded text in chunks that fill perfect
rectangles (r X c), with c chunks of r length, separated by spaces.
For phrases that are n characters short of the perfect rectangle,
pad each of the last n chunks with a single trailing space.

"imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn  sseoau "
Notice that were we to stack these, we could visually decode the
cyphertext back in to the original message:

"imtgdvs"
"fearwer"
"mayoogo"
"anouuio"
"ntnnlvt"
"wttddes"
"aohghn "
"sseoau "

*/

package main

import (
    "fmt"
    _"strconv"
)

func normalize(input string) (output string) {
    var LOOKUP map[string]int = map[string]int{"a":0,"b":0,"c":0,"d":0,
        "e":0,"f":0,"g":0, "h":0,"i":0,"j":0,"k":0,"l":0,"m":0,
        "n":0,"o":0,"p":0,"q":0,"r":0, "s":0,"t":0,"u":0,"v":0,
        "w":0,"x":0,"y":0,"z":0}

    for i,_ := range(input){
        test := string(input[i])
        _, ok := LOOKUP[test]
        if ok {
            output += test
        }
    }
    
    return
}

func findPair(product int) (c,r int){
    for ; r < product; r++{
        for c = 0; c < product; c++{
            if c*r == product && c - r <= 1 && c >= r{
                return
            }
        }
    }
    return
}

func main() {

    // normalize text by removing all non a-z chars
    normalized := normalize("asdaqdsadw1112!!@@ccc dsdds d!!!e")
    fmt.Println(normalized)

    // find c X r so that c >= r and c - r <= 1
    var (c, r int)
    c, r = findPair(20)
    fmt.Println(c,r)

    // create a new multidimensional array of length [c][r]
    normalizedList := make([][]string, c)
    for i := range normalizedList {
        normalizedList[i] = make([]string, r)
    }
    fmt.Println(normalizedList)

    // iterate through the normalized string and put each character into
    // the array going left to right, top to bottom
    counter := 0
    for i := 0; i < len(normalizedList); i++{
        for j := 0; j < len(normalizedList[i]); j++{
            normalizedList[i][j] = string(normalized[counter])
            counter += 1
        }
    }
    fmt.Println(normalizedList)

    // create a new array of length [r][c] for the scrambled list
    scrambledList := make([][]string, r)
    for i := range scrambledList {
        scrambledList[i] = make([]string, c)
    }
    fmt.Println(scrambledList)

    // iterate through the multidimensional array and swap the characters
    // so that normalizedList[c][r] == scrambledList[r][c]
    counter = 0
    for i := 0; i < len(normalizedList); i++{
        for j := 0; j < len(normalizedList[i]); j++{
            scrambledList[j][i] = string(normalized[counter])
            counter += 1
        }
    }
    fmt.Println(scrambledList)

    // print the scrambled list line by line
    for line := range(scrambledList){
        fmt.Println(scrambledList[line])
    }

}
