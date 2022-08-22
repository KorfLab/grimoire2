package hmm

import (
	"github.com/korflab/grimoire2/toolbox"
	"fmt"
)

func Emission_model_ho(seq_type string, order int) map[string]map[string]float64{	
	var letters []string
	if seq_type == "nt" {
		letters = []string{"A", "C", "G", "T"}
	} else if seq_type == "aa" {
		letters = []string{"G", "A", "L", "M", "F", "W", "K", "Q", "E", "S",
						"P", "V", "I", "C", "Y", "H", "R", "N", "D", "T"}
	} else {
		err := fmt.Errorf("error: seq_type received \"%s\"; expects \"nt\" or \"aa\"", seq_type)
		panic(err)
	}
	
	table := make(map[string]map[string]float64)
	if order > 0 {
		kmers := toolbox.Generate_kmers(seq_type, order)
		for _, kmer := range kmers {
			table[kmer] = make(map[string]float64)
			for _, letter := range letters {
				table[kmer][letter] = 0
			}
		}
	} else {
		err := fmt.Errorf("error: order must be greater than 0; received %d", order)
		panic(err)
	}
	
	return table
}

func Emission_model_zo(seq_type string) map[string]float64{	
	table := make(map[string]float64)
	kmers := toolbox.Generate_kmers(seq_type, 1)
	for _, kmer := range kmers {
		table[kmer] = 0
	}
	return table

}


func Train_emission_ho(seqs []map[string]interface{}, order int) map[string]map[string]float64 {
	count := Emission_model_ho("nt", order)
	freq := make(map[string]map[string]float64)
	
	for _, dict := range seqs {
		seq := dict["seq"].(string)
		weight := dict["weight"].(float64)
		for i := 0; i < len(seq) - order; i++ {
			ctx := seq[i:i+order]
			nt := seq[i+order:i+order+1]
			if val, ok := count[ctx][nt]; ok {
				_ = val
				count[ctx][nt] += weight
			}
		}
	}
	
	for ctx, _ := range count {
		total := 0.0
		freq[ctx] = make(map[string]float64)
		for nt, _ := range count[ctx] {
			total += count[ctx][nt]
		}
		if total > 0 {
			for nt := range count[ctx] {
				freq[ctx][nt] = count[ctx][nt] / total
			}
		} else {
			for nt := range count[ctx] {
				freq[ctx][nt] = 0
			}
		}
	}
	
	return freq
}


func Train_emission_zo(seqs []map[string]interface{}) map[string]float64{
	count := Emission_model_zo("nt")
	freq := make(map[string]float64)
	
	total := 0.0
	for _, dict := range seqs {
		seq := dict["seq"].(string)
		weight := dict["weight"].(float64)
		for i := 0; i < len(seq); i++ {
			nt := seq[i:i+1]
			if val, ok := count[nt]; ok {
				_ = val
				count[nt] += weight
				total += weight
			}
		}
	}
	
	for nt, _ := range count {
		freq[nt] = count[nt] / total
	}
	
	return freq
}


