package main

import (
        "flag"
        "fmt"
        "strings"
        "math/rand"
        "strconv"
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

        ch := make(chan string)

        for _, i := range motif_lens  {
                var int_len, _ = strconv.Atoi(i)
                go gen_motif(int_len, ch)
        }

        for i := 0; i < len(motif_lens); i++ {
                fmt.Println(<-ch)
        }
}

func get_rand_char() string {
        switch (rand.Int() % 4) {
        case 0:
                return "A"
        case 1:
                return "C"
        case 2:
                return "G"
        case 3:
                return "T"
        default:
                return "X"
        }
}

func gen_motif(len int, c chan<- string) {
        var result string
        for i := 0; i < len; i++ {
                result += get_rand_char()
        }
        c <- result
}