package classes

import (
	"errors"
	"math"
)

// By default, I am exporting all fields in Feature
type Feature struct {
	Dna      *DNA
	Beg      int
	End      int
	Len      int
	Strand	 byte
	Type	   string // Should be SO-compliant
	Phase    byte
	Score    float64 // need a representative of . scores, floatMAX?
	Source   string
	Issues  *Issues
	Parent	*Feature
	Children []*Feature //what type is this?
	validated    bool
}

// NewFeature creates a pointer to an empty Feature object, with each fields
// initialized to default zero values. User must specify values their feature
// will have.
func NewFeature() *Feature {
	feat = new(Feature)
	feat.DNA = nil
	feat.Beg = -1
	feat.End = -1
	feat.Len = 0 //do we even need this? its easy to calculate
	feat.Strand = byte(".")
	feat.Type = ""
	feat.Phase = "."
	feat.Score = math.maxFloat64
	feat.Source = ""
	feat.Parent = nil
	feat.Children = nil
	feat.validated = false

	return feat
}

// validate() method checks internal values of Features and sets validated to
// true if the values are legal
func (f Feature) validate() {
	if f.Beg < 1 {
		f.Issues
		//panic("Error: beg<1")
	} else if f.Beg > f.End {
		f.Issues //do something
		panic("Error: beg>End")
	} else if len == 0 { //see comment above
		panic("Error: Len = 0")
	}
}

func (f Feature) revcomp() {
	if self.strand == byte("+") {
		self.strand = byte("-")
	} else if self.strand == byte("-") {
		self.strand = byte("+")
	}

//	newbeg =
}

type Issues struct {
	//this could be an empty slice of type issue, append to it for each Issues
}
