package main

import (
  "flag"
  "fmt"
  "strings"
  )

var seq_len int
var seq_num int
var motif_lens_str string
var mutation_num int
var min_motifs int
var active_len int

func main() {
	flag.IntVar(&seq_len, "sequence-length", 500, "the length of a sequence")
  flag.IntVar(&seq_num, "sequence-number", 100, "the number of sequences to be generated")
  flag.StringVar(&motif_lens_str, "motif-lengths", "15,16,17,18,19,20", "a comma-separated list of motif lengths")
  flag.IntVar(&mutation_num, "mutation-number", 2, "the number of mutations per motif")
  flag.IntVar(&min_motifs, "min-motifs", 3, "The minimum number of motifs in a sequence")
  flag.IntVar(&active_len, "active-length", 150, "length of the active sub-region in the sequence")

  var motif_lens = strings.Split(motif_lens_str, ",")

  fmt.Printf("Sequence Length: %d\n", seq_len)
  fmt.Printf("Sequence Number: %d\n", seq_num)
  fmt.Printf("Motif Lengths: %s\n", motif_lens)
  fmt.Printf("Mutation Number: %d\n", mutation_num)
  fmt.Printf("Minimum Motifs: %d\n", min_motifs)
  fmt.Printf("Active Sub-region length: %d\n", active_len)
}