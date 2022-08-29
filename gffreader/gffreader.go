package gffreader

import (
	"bufio"
	//"bytes"
	"compress/gzip"
	//"fmt"
	//"grimoire2/feature"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	Seqid  string
	Source string
	Type   string
	Beg    int64
	End    int64
	Score  float64
	Strand byte
	Phase  byte
	ID     string
	Parent []string
}

type Iterator interface {
	Record() Record
	Next() bool
}

type State_Iterator struct {
	scanner    *bufio.Scanner
	gzipfh     *gzip.Reader
	filehandle *os.File
	Current    *Record
	Linecount  int
	done       bool
}

func (si *State_Iterator) Record() *Record {
	return si.Current
}

func (si *State_Iterator) Next() bool {
	if si.done == true {
		si.filehandle.Close()

		if si.gzipfh != nil {
			si.gzipfh.Close()
		}
		return false

	} else {

		if si.scanner.Scan() {
			line := si.scanner.Text()
			c := strings.Fields(line)

			si.Current.Seqid = c[1]
			si.Current.Source = c[1]
			si.Current.Type = c[2]
			si.Current.Beg, _ = strconv.ParseInt(c[3], 10, 32)
			si.Current.End, _ = strconv.ParseInt(c[4], 10, 32)
			si.Current.Score, _ = strconv.ParseFloat(c[5], 64)
			strand := []byte(c[6])
			si.Current.Strand = strand[0]
			phase := []byte(c[7])
			si.Current.Phase = phase[0]
			attributes := strings.Split(c[8], ";")

			for _, att := range attributes {
				if strings.HasPrefix(att, "Parent=") {
					si.Current.Parent = strings.Split(att, ",")

				} else if strings.HasPrefix(att, "ID=") {
					si.Current.ID = att
				}
			}
		} else {
			si.done = true

		}

	}
	return true

}

func Read_record(fname string) *State_Iterator {
	filehandle, err := os.Open(fname)

	if err != nil {
		panic(err)
	}

	si := new(State_Iterator)
	si.Current = new(Record)
	si.done = false
	si.filehandle = filehandle

	scanner := bufio.NewScanner(filehandle)
	if strings.HasSuffix(fname, ".gz") {

		gzipfh, err := gzip.NewReader(filehandle)
		si.gzipfh = gzipfh

		if err != nil {
			panic(err)
		}

		scanner = bufio.NewScanner(gzipfh)

	}

	si.scanner = scanner


	si.Linecount = countlines(fname) + 1

	return si
}


func countlines(fname string) int {
	newfh, err := os.Open(fname)
	var linecount int
	var newgzipfh *gzip.Reader

	if err != nil {
		panic(err)
	}

	newscanner := bufio.NewScanner(newfh)
	if strings.HasSuffix(fname, ".gz") {

		newgzipfh, err := gzip.NewReader(newfh)

		if err != nil {
			panic(err)
		}

		newscanner = bufio.NewScanner(newgzipfh)

	}

	for newscanner.Scan() {
		linecount++
	}

	newfh.Close()

	if newgzipfh != nil {
		newgzipfh.Close()
	}
	return linecount
}
