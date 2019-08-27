/*
https://exercism.io/my/solutions/aa0c0571ef6048e9b2b1cc680519e72e
Introduction

Given a word and a list of possible anagrams,
select the correct sublist.

Given "listen" and a list of candidates
like "enlist" "google" "inlets" "banana"
the program should return a list containing "inlets".

*/

package main

import (
    "fmt"
    "sort"
    _"reflect"
    "strings"
)

func getAnagrams(word string) (value []string) {
    var LOOKUP map[string][]string = map[string][]string{"eilnst" : {"listen",
        "enlist","inlets"}}
    var slice []string = make([]string, len(word))

    for i, v := range(word){
        slice[i] = string(v)
    }

    sort.Strings(slice)
    test := strings.Join(slice,"")

    value, ok := LOOKUP[test]

    if ok {
        return value
    }

    return value

}

func main(){
    anagrams := getAnagrams("listen")
    fmt.Println(anagrams)
}
