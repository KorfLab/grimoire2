package toolbox

import (
	"errors"
	"strings"
)


func _kmers(seq []string, table *map[string]int, key string, n int, k int, v int) {
	if k == 0 {
		(*table)[key] = v
		return
	}
	
	for i := 0; i < n; i++ {
		t := key + seq[i]
		_kmers(seq, table, t, n, k-1,v)
	}
}

func Generate_kmers(seq_type string, k int, pseudo int) map[string]int{
	table := make(map[string]int)
	if seq_type == "nt" {
		seq := []string{"A", "C", "G", "T"}
		_kmers(seq, &table, "", 4, k, pseudo)
	} else if seq_type == "aa" {
		seq := []string{"G", "A", "L", "M", "F", "W", "K", "Q", "E", "S",
						"P", "V", "I", "C", "Y", "H", "R", "N", "D", "T"}
		_kmers(seq, &table, "", 20, k, pseudo)
	} else {
		err := errors.New("error: seq_type expects only \"nt\" or \"aa\"")
		panic(err)
	}
	
	return table
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

func Longest_orf(seq string) string {
	
	for f := 0; f < 3; f++ {
		pro := Translate_str(seq)
		start := 0
		stop := 0
		for start < len(pro) {
			
		}
	}
}