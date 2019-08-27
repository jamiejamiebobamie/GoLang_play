/*

https://exercism.io/my/solutions/dc4fe43763ac42ae80e133e289f80923

Introduction

Translate RNA sequences into proteins.

RNA can be broken into three nucleotide sequences called codons,
and then translated to a polypeptide like so:

RNA: "AUGUUUUCU" => translates to

Codons: "AUG", "UUU", "UCU" => which become a polypeptide
with the following sequence =>
Protein: "Methionine", "Phenylalanine", "Serine"

There are 64 codons which in turn correspond to 20 amino acids;
however, all of the codon sequences and resulting amino acids are
not important in this exercise. If it works for one codon,
the program should work for all of them. However,
feel free to expand the list in the test suite to include them all.

There are also three terminating codons (also known as 'STOP' codons);
if any of these codons are encountered (by the ribosome),
all translation ends and the protein is terminated.

All subsequent codons after are ignored, like this:

RNA: "AUGUUUUCUUAAAUG" =>

Codons: "AUG", "UUU", "UCU", "UAA", "AUG" =>

Protein: "Methionine", "Phenylalanine", "Serine"

Note the stop codon "UAA" terminates the translation and the
final methionine is not translated into the protein sequence.

Below are the codons and resulting Amino Acids needed for the exercise.

Codon	Protein
AUG	Methionine
UUU, UUC	Phenylalanine
UUA, UUG	Leucine
UCU, UCC, UCA, UCG	Serine
UAU, UAC	Tyrosine
UGU, UGC	Cysteine
UGG	Tryptophan
UAA, UAG, UGA	STOP

*/

package main

import "fmt"

    func mapCodons(RNA string) (codonList []string){
        var codons map[string]string = map[string]string{
                                            "AUG":"Methionine",
                                            "UUU":"Phenylalanine",
                                            "UUC":"Phenylalanine",
                                            "UUA":"Leucine",
                                            "UUG":"Leucine",
                                            "UCU":"Serine",
                                            "UCC":"Serine",
                                            "UCA":"Serine",
                                            "UCG":"Serine",
                                            "UAU":"Tyrosine",
                                            "UAC":"Tyrosine",
                                            "UGU":"Cysteine",
                                            "UGC":"Cysteine",
                                            "UGG":"Tryptophan",
                                            "UAA":"STOP",
                                            "UAG":"STOP",
                                            "UGA":"STOP",
                                        }
        for i := 0; i < len(RNA);{
            if i+3 <= len(RNA){
                codon := RNA[i:i+3]
                v, ok := codons[codon]
                if ok {
                    if v == "STOP"{
                        return
                    }
                    codonList = append(codonList, v) // how slow is this?
                }
            }
            i+=3
        }
        return
    }

func main(){
    RNA:= "AUGUUUUCUUAAAUG"
    fmt.Println(mapCodons(RNA))
}
