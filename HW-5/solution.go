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

        for n, i := range motif_lens  {
                var int_len, _ = strconv.Atoi(i)
                go gen_motif(int_len, n, ch)
        }

        motif_templates := make([]string, len(motif_lens))

        for i := 0; i < len(motif_lens); i++ {
                motif_templates[i] = <-ch
                fmt.Println(motif_templates[i]) // TODO: Write to file
        }

        for i := 0; i < seq_num; i++ {
                go gen_sequence(seq_len, active_len, min_motifs, mutation_num, motif_templates, ch)
        }

        for i := 0; i < seq_num; i++ {
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

func gen_motif(len int, motif_num int, c chan<- string) {
        // var result string = "m" + strconv.Itoa(motif_num) + " "
        var result string
        for i := 0; i < len; i++ {
                result += get_rand_char()
        }
        c <- result
}

func gen_sequence(length int, active_len int, min_motifs int, mutations int, motifs_ref []string, c chan<- string) {
        var motifs = motifs_ref
        var header string = ">"
        var sequence string
        starting_point := rand.Int() % (length - active_len)
        current_pos := starting_point
        for i := 0; i < min_motifs; i++ {
                chosen_motif := rand.Int() % len(motifs)
                motif := motifs[chosen_motif]
                motifs = append(motifs[:chosen_motif], motifs[chosen_motif+1:]...)
                header += " m" + strconv.Itoa(chosen_motif) + " " + strconv.Itoa(len(motif)) + " " + strconv.Itoa(current_pos)
                sequence += gen_mutation(motif, mutations)
                current_pos += len(motif)
        }
        c <- header + "\n" + sequence
}

// Deterministic algorithm
func gen_mutation(motif string, num_mut int) string {
        var mutation string = motif
        for i := 0; i < num_mut; i++ {
                prev := mutation[0:i]
                post := mutation[i+1:len(mutation)]
                newchar := get_next_char(mutation[i])
                mutation = prev + newchar + post
        }
        return mutation
}

func get_next_char(char byte) string {
        switch (char) {
        case 64:
                return "C"
        case 67:
                return "G"
        case 71:
                return "T"
        case 84:
                return "A"
        default:
                return "A"
        }
}
