/*

Introduction

Calculate the Hamming difference between two DNA strands.

A mutation is simply a mistake that occurs during
the creation or copying of a nucleic acid, in particular DNA.
Because nucleic acids are vital to cellular functions, mutations
tend to cause a ripple effect throughout the cell.
Although mutations are technically mistakes, a very rare
mutation may equip the cell with a beneficial attribute. In fact,
the macro effects of evolution are attributable by the accumulated result of
beneficial microscopic mutations over many generations.

The simplest and most common type of nucleic acid mutation is a point mutation,
which replaces one base with another at a single nucleotide.

By counting the number of differences between two homologous DNA strands taken
from different genomes with a common ancestor, we get a measure of the minimum
number of point mutations that could have occurred on the evolutionary path
between the two strands.

This is called the 'Hamming distance'.

It is found by comparing two DNA strands and counting how many of the
nucleotides are different from their equivalent in the other string.

GAGCCTACTAACGGGAT
CATCGTAATGACGGCCT
^ ^ ^  ^ ^    ^^
The Hamming distance between these two DNA strands is 7.

*/

package main

import (
    "fmt"
)

func compare(string1, string2 string) (hamming int) {
    for i := 0; i < len(string1); i++ {
        if string1[i] != string2[i]{
            hamming += 1
        }
    }
    return
}

func main(){

    string1:="GAGCCTACTAACGGGAT"
    string2:="CATCGTAATGACGGCCT"

    fmt.Println(compare(string1, string2))
}
