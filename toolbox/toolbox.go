// Package toolbox implements miscellaneous functions for operating on numbers, strings, and such
package toolbox

import (
	"time"
	"math"
	"math/rand"
	"fmt"
	"strings"
)

func Prod[Num int | float64](iterable []Num) Num {
	var a Num = 1
	
	for _, num := range iterable {
		a *= num
	}
	
	return a

}

func Log(p float64) float64 {
	if p < 0 {
		err := fmt.Errorf("error: p < 0: %f", p)
		panic(err)
	}
	
	if p == 0 {
		return -999
	} else {
		return math.Log(p)
	}
}

func Sumlog(v1 float64, v2 float64) float64 {
	s := math.Exp(v1) + math.Exp(v2)
	return math.Log(s)
}

func _kmers(letters []string, kmers *[]string, key string, n int, k int) {
	if k == 0 {
		*kmers = append(*kmers, key)
		return
	}
	
	for i := 0; i < n; i++ {
		t := key + letters[i]
		_kmers(letters, kmers, t, n, k-1)
	}
}

func Generate_kmers(seq_type string, k int) []string {
	kmers := make([]string, 0)
	if seq_type == "nt" {
		letters := []string{"A", "C", "G", "T"}
		_kmers(letters, &kmers, "", 4, k)
	} else if seq_type == "aa" {
		letters := []string{"G", "A", "L", "M", "F", "W", "K", "Q", "E", "S",
						"P", "V", "I", "C", "Y", "H", "R", "N", "D", "T"}
		_kmers(letters, &kmers, "", 20, k)
	} else {
		err := fmt.Errorf("error: seq_type received \"%s\"; expects \"nt\" or \"aa\"", seq_type)
		panic(err)
	}
	
	return kmers
}

var CODONS = map[string]string {
	"AAA" : "K",	"AAC" : "N",	"AAG" : "K",	"AAT" : "N",
	"AAR" : "K",	"AAY" : "N",	"ACA" : "T",	"ACC" : "T",
	"ACG" : "T",	"ACT" : "T",	"ACR" : "T",	"ACY" : "T",
	"ACK" : "T",	"ACM" : "T",	"ACW" : "T",	"ACS" : "T",
	"ACB" : "T",	"ACD" : "T",	"ACH" : "T",	"ACV" : "T",
	"ACN" : "T",	"AGA" : "R",	"AGC" : "S",	"AGG" : "R",
	"AGT" : "S",	"AGR" : "R",	"AGY" : "S",	"ATA" : "I",
	"ATC" : "I",	"ATG" : "M",	"ATT" : "I",	"ATY" : "I",
	"ATM" : "I",	"ATW" : "I",	"ATH" : "I",	"CAA" : "Q",
	"CAC" : "H",	"CAG" : "Q",	"CAT" : "H",	"CAR" : "Q",
	"CAY" : "H",	"CCA" : "P",	"CCC" : "P",	"CCG" : "P",
	"CCT" : "P",	"CCR" : "P",	"CCY" : "P",	"CCK" : "P",
	"CCM" : "P",	"CCW" : "P",	"CCS" : "P",	"CCB" : "P",
	"CCD" : "P",	"CCH" : "P",	"CCV" : "P",	"CCN" : "P",
	"CGA" : "R",	"CGC" : "R",	"CGG" : "R",	"CGT" : "R",
	"CGR" : "R",	"CGY" : "R",	"CGK" : "R",	"CGM" : "R",
	"CGW" : "R",	"CGS" : "R",	"CGB" : "R",	"CGD" : "R",
	"CGH" : "R",	"CGV" : "R",	"CGN" : "R",	"CTA" : "L",
	"CTC" : "L",	"CTG" : "L",	"CTT" : "L",	"CTR" : "L",
	"CTY" : "L",	"CTK" : "L",	"CTM" : "L",	"CTW" : "L",
	"CTS" : "L",	"CTB" : "L",	"CTD" : "L",	"CTH" : "L",
	"CTV" : "L",	"CTN" : "L",	"GAA" : "E",	"GAC" : "D",
	"GAG" : "E",	"GAT" : "D",	"GAR" : "E",	"GAY" : "D",
	"GCA" : "A",	"GCC" : "A",	"GCG" : "A",	"GCT" : "A",
	"GCR" : "A",	"GCY" : "A",	"GCK" : "A",	"GCM" : "A",
	"GCW" : "A",	"GCS" : "A",	"GCB" : "A",	"GCD" : "A",
	"GCH" : "A",	"GCV" : "A",	"GCN" : "A",	"GGA" : "G",
	"GGC" : "G",	"GGG" : "G",	"GGT" : "G",	"GGR" : "G",
	"GGY" : "G",	"GGK" : "G",	"GGM" : "G",	"GGW" : "G",
	"GGS" : "G",	"GGB" : "G",	"GGD" : "G",	"GGH" : "G",
	"GGV" : "G",	"GGN" : "G",	"GTA" : "V",	"GTC" : "V",
	"GTG" : "V",	"GTT" : "V",	"GTR" : "V",	"GTY" : "V",
	"GTK" : "V",	"GTM" : "V",	"GTW" : "V",	"GTS" : "V",
	"GTB" : "V",	"GTD" : "V",	"GTH" : "V",	"GTV" : "V",
	"GTN" : "V",	"TAA" : "*",	"TAC" : "Y",	"TAG" : "*",
	"TAT" : "Y",	"TAR" : "*",	"TAY" : "Y",	"TCA" : "S",
	"TCC" : "S",	"TCG" : "S",	"TCT" : "S",	"TCR" : "S",
	"TCY" : "S",	"TCK" : "S",	"TCM" : "S",	"TCW" : "S",
	"TCS" : "S",	"TCB" : "S",	"TCD" : "S",	"TCH" : "S",
	"TCV" : "S",	"TCN" : "S",	"TGA" : "*",	"TGC" : "C",
	"TGG" : "W",	"TGT" : "C",	"TGY" : "C",	"TTA" : "L",
	"TTC" : "F",	"TTG" : "L",	"TTT" : "F",	"TTR" : "L",
	"TTY" : "F",	"TRA" : "*",	"YTA" : "L",	"YTG" : "L",
	"YTR" : "L",	"MGA" : "R",	"MGG" : "R",	"MGR" : "R",
}

func Reverse_str(str string) string {
	rev := []rune(str)
	for i, j := 0, len(rev) - 1; i < j; i, j = i + 1, j - 1 {
		rev[i], rev[j] = rev[j], rev[i]
	}	
	
	return string(rev)	
}

func Revcomp_str(seq string) string {
	from := "ACGTRYMKWSBDHVN"
	to   := "tgcayrkmwsvhdbn"
	
	revcomp := strings.ToUpper(seq)
	for i := 0; i < len(from); i++ {
		revcomp = strings.ReplaceAll(revcomp, string(from[i]), string(to[i]))
	}
	revcomp = strings.ToUpper(revcomp)
	
	return Reverse_str(revcomp)
}

func Translate_str(seq string) string {
	pro := make([]string,0)
	
	for i := 0; i < len(seq)-2; i+=3 {
		codon := seq[i:i+3]
		aa, is_codon := CODONS[codon]
		if !is_codon {
			pro = append(pro, "X")
		} else {
			pro = append(pro, aa)
		}
	}
	
	return strings.Join(pro, "")
}

func Longest_orf(seq string) (string, int) {
	longest_orf := ""
	longest_orf_start := -1
	for f := 0; f < 3; f++ {
		pro := Translate_str(seq[f:])
		start := 0
		for start < len(pro) {
			end := 0
			if string(pro[start]) == "M" {
				for i, aa := range pro[start+1:] {
					if string(aa) == "*" {
						end = i + start + 1
						break
					}
				}
			}
			if end != 0 && len(pro[start:end]) > len(longest_orf) {
				longest_orf = pro[start:end]
				longest_orf_start = start * 3 + f
			}
			start += 1
		}
	}
	
	return longest_orf, longest_orf_start
}

func Random_dna(length int, a float64, c float64, g float64, t float64) string {
	if a + c + g + t - 1.0 > 0.00001 {
		err := fmt.Errorf("error: sequence probabilities must add up to 1.0")
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	seq := make([]string, length)
	for i := 0; i < length; i++ {
		rf := rand.Float64()
		if rf < a {
			seq[i] = "A"
		} else if rf < a + c {
			seq[i] = "C"
		} else if rf < a + c + g {
			seq[i] = "G"
		} else {
			seq[i] = "T"
		}
	}
	return strings.Join(seq,"")
}
