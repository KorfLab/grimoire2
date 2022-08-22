package classes

import (
	"errors"
	"math"
)

// By default, I am exporting all fields in Feature
type Feature struct {
	Dna       *DNA
	Beg       int
	End       int
	Len       int
	Strand    byte
	Type      string // Should be SO-compliant
	Phase     byte
	Score     float64 // need a representative of . scores, floatMAX?
	Source    string
	Issues    *Tracker
	Parent    *Feature
	Children  []*Feature //what type is this?
	validated bool
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
	feat.Phase = byte(".")
	feat.Score = math.maxFloat64
	feat.Source = ""
	feat.Issues = nil
	feat.Parent = nil
	feat.Children = nil
	feat.validated = false

	return feat
}

// validate() method checks internal values of Features and sets validated to
// true if the values are legal (incomplete)
func (f Feature) validate() {
	errors := []Issue
	if f.Beg < 1 {
		errors = append(errors, beg_oob)

	} else if f.Beg > f.End {
		errors = append(errors, bad_coordinates)

	} else if len == 0 { //see comment above
		errors = append(errors, zero_len)

	}

	// If these are the first issues, make a tracker for them
	if errors != nil && f.Issues == nil {
		f.Issues = NewTracker(f)
	}

	// Append our errors slice to the end of the existing Issues.Raised slice
	if errors != nil {
		f.Issues.Raised = append(f.Issues.Raised, errors)
	}
}

func (f Feature) revcomp() {
	if f.strand == byte("+") {
		f.strand = byte("-")
	} else if f.strand == byte("-") {
		f.strand = byte("+")
	}

	newbeg = len(f.Dna.seq) - f.End + 1
	newend = len(f.Dna.seq) - f.Beg + 1
	f.Beg = newbeg
	f.End = newend
}

type Tracker struct {
	Raised []Issue
	Feature *Feature
}

func NewTracker (f *Feature) *Tracker {
	t = new(Tracker)
	t.Raised = []Issue
	t.Feature = f

	return t
}

type Issue int64

const (
	beg_oob Issue = iota
	bad_coordinates
	zero_len
	nonCDS_phase
	incorrect_phase
	nonSO_type
	unexpected_value
	all_Ns
	feature_size
	exon_overlap
	child_oob
)

func (i Tracker) addissue(Issue) {
	i.Issues += Issue
}
