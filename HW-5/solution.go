/**
 * CPL project 5
 * Copywrite Sam Beckmann
 * All Rights Reserved
 */
package main

import (
        "flag"
        "os"
        "strings"
        "math/rand"
        "strconv"
)

func main() {

        var seq_len int
        var seq_num int
        var motif_lens_str string
        var mutation_nums string
        var min_motifs int
        var active_len int

        flag.IntVar(&seq_len, "sequence-length", 500, "the length of a sequence")
        flag.IntVar(&seq_num, "sequence-number", 100, "the number of sequences to be generated")
        flag.StringVar(&motif_lens_str, "motif-lengths", "15,16,17,18,19,20", "a comma-separated list of motif lengths")
        flag.StringVar(&mutation_nums, "mutation-numbers", "1,2,3,3,2,1", "the number of mutations per motif")
        flag.IntVar(&min_motifs, "min-motifs", 3, "The minimum number of motifs in a sequence")
        flag.IntVar(&active_len, "active-length", 150, "length of the active sub-region in the sequence")

        var motif_lens = strings.Split(motif_lens_str, ",")
        var mut_nums_str = strings.Split(mutation_nums, ",")

        mut_nums := make([]int, len(mut_nums_str))
        for n, i := range mut_nums_str {
                mut_nums[n], _ = strconv.Atoi(i)
        }


        motif_f, _ := os.Create("motifs.txt")
        defer motif_f.Close()

        seq_f, _ := os.Create("sequences.txt")
        defer seq_f.Close()

        motif_templates := make([]string, len(motif_lens))

        ch := make(chan string)
        for n, i := range motif_lens {
                var value, _ = strconv.Atoi(i)
                go gen_template(value, n, ch)
        }

        for i := 0; i < len(motif_lens); i++ {
                motif_templates[i] = <-ch
                motif_f.WriteString(motif_templates[i] + "\n")
        }

        rand_chan := make(chan string)
        motif_chan := make(chan string)

        go gen_motif(motif_templates, mut_nums, motif_chan)
        go gen_rand_seq(rand_chan)

        for i := 0; i < seq_num; i++ {
                seq := gen_sequence(seq_len, i, active_len, min_motifs, motif_templates, mut_nums, motif_chan, rand_chan)
                seq_f.WriteString(seq + "\n")
        }

        motif_f.Sync()
        seq_f.Sync()
}

func get_rand_char() string {
        switch (rand.Intn(4)) {
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

func gen_template(len int, motif_num int, c chan<- string) {
        var result string
        for i := 0; i < len; i++ {
                result += get_rand_char()
        }
        c <- result
}

func gen_sequence(length int, seq_num int, active_len int,
                         min_motifs int, motifs []string, mut_nums []int,
                         motif_chan <-chan string, rand_chan <-chan string) string {

        var start_loc int
        var sequence string
        var header string
        not_valid := true
        for not_valid {
                sequence = ""
                start_loc = rand.Intn(length - active_len)
                current_loc := start_loc
                header = ">seq" + strconv.Itoa(seq_num)
                num_motifs := 0

                for num_motifs < min_motifs {
                        var selection string
                        select {
                        case selection = <- rand_chan:
                                sequence += selection
                                current_loc += len(selection)
                        case selection = <- motif_chan:
                                split := strings.Split(selection, ",")
                                header += split[0] + " " + strconv.Itoa(current_loc) + " " + strconv.Itoa(len(split[1]) - 1 + current_loc)
                                sequence += split[1]
                                num_motifs++
                                current_loc += len(split[1])
                        }
                }

                not_valid = len(sequence) > active_len
        }

        var pre string
        for i := 0; i < start_loc; i++ {
                pre += get_rand_char()
        }

        sequence = pre + sequence

        for len(sequence) < length {
                sequence += get_rand_char()
        }

        return header + "\n" + sequence
}

func gen_motif(motifs []string, muts []int, c chan<- string) {
        for {
                motif := rand.Intn(len(motifs))
                c <- " m" + strconv.Itoa(motif) + "," + gen_mutation(motifs[motif], muts[motif])
        }
}

// Deterministic algorithm TODO
func gen_mutation(motif string, num_mut int) string {
        mut_chance := float32(num_mut) / float32(len(motif))
        var mutation string = motif
        num_muts := 0
        for i := 0; i < num_mut; i++ {
                if (rand.Float32() < mut_chance) {
                    newchar := get_rand_char()
                    for newchar == string(mutation[i]) {
                        newchar = get_rand_char()
                    }
                    mutation = mutation[:i] + newchar + mutation[i+1:]
                    num_muts++
                }
        }
        if (num_muts == num_mut) {
            return mutation
        } else {
            return gen_mutation(motif, num_mut)
        }
}

func gen_rand_seq(c chan<- string) {
        var result string = ""
        for {
                if rand.Intn(10) == 0 {
                        c <- result
                        result = ""
                } else {
                        result += get_rand_char()
                }
        }
}
