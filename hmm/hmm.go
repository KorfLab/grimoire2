package hmm

import (
	"github.com/korflab/grimoire2/toolbox"
	"fmt"
)

func emission_model_ho(seq_type string, order int) map[string]map[string]float64{	
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

func emission_model_zo(seq_type string) map[string]float64{	
	table := make(map[string]float64)
	kmers := toolbox.Generate_kmers(seq_type, 1)
	for _, kmer := range kmers {
		table[kmer] = 0
	}
	return table

}

/*
func Train_emission_ho(){

}
*/

func Train_emission_zo(seqs []map){

}
