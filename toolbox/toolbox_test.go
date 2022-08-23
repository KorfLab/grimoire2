package toolbox

import (
	"testing"
	"math"
)

func almost_equal(a float64, b float64) bool {
	return math.Abs(a - b) <= 0.00001
}

func Test_Prod(t *testing.T) {
	Prod_tests := map[string]struct {
		int_input []int
		float_input []float64
		output float64
	} {
		"List of integers": {
			int_input: []int{2, 4, 5},
			output: 40,
		},
		"List of floats": {
			float_input: []float64{2, 1.5, -1},
			output: -3,
		},
	}

	t.Run("List of integers", func(t *testing.T) {
		//t.Parallel()
		test := Prod_tests["List of integers"]
		if got, expected := Prod(test.int_input), test.output; !almost_equal(got, expected) {
			t.Errorf("Prod(%v) returned %f; expected %f\n", test.int_input, got, expected)
		}
	})
	
	t.Run("List of floats", func(t *testing.T) {
		//t.Parallel()
		test := Prod_tests["List of floats"]
		if got, expected := Prod(test.float_input), test.output; !almost_equal(got, expected) {
			t.Errorf("Prod(%v) returned %f; expected %f\n", test.float_input, got, expected)
		}
	})
	
}


func Test_Log(t *testing.T) {
	Log_tests := map[string]struct {
		input float64
		output float64
	} {
		"input 0": {
			input: 0,
			output: -999,
		},
		"input > 0": {
			input: 1,
			output: 0,
		},
	}
	
	for name, test := range Log_tests {
		test := test
		t.Run(name, func(t *testing.T){
			//t.Parallel()
			if got, expected := Log(test.input), test.output;
			!almost_equal(got, expected) {
				t.Errorf("Log(%f) returned %f; expected %f\n", test.input, got, expected)
			}
		})
	}
}


func Test_Sumlog(t *testing.T) {
	Sumlog_tests := map[string]struct {
		input1 float64
		input2 float64
		output float64
	} {
		"input instance a": {
			input1: -1,
			input2: -1,
			output: -0.3068528194400547,
		},
		"input instance b": {
			input1: -1.2,
			input2: -2.5,
			output: -0.9589915461670078,
		},
	}
	
	for name, test := range Sumlog_tests {
		test := test
		t.Run(name, func(t *testing.T){
			//t.Parallel()
			if got, expected := Sumlog(test.input1, test.input2), test.output;
			!almost_equal(got, expected) {
				t.Errorf("Sumlog(%f, %f) returned %f; expected %f\n", test.input1,
				test.input2, got, expected)
			}
		})
	}
}

/*
func Test_Generate_kmers(t *testing.T) {

}

func Test_Revcomp_str(t *testing.T) {

}

func Test_Translate_str(t *testing.T) {

}

func Test_Longest_orf(t *testing.T) {

}

func Test_Random_dna(t *testing.T) {

}
*/
