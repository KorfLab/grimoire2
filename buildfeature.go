package main

import (
  //"errors"
  "flag"
  "fmt"
  "grimoire2/feature"
  "grimoire2/gffreader"
  //"strings"
)


func main () {
  var fname string
  var i int

  flag.StringVar(&fname, "f", "", "path to GFF file")

  flag.Parse()

  gffr := gffreader.Read_record(fname)

  ftable := make([]classes.Feature, gffr.Linecount)

  for gffr.Next() {
    line := gffr.Record()
    if i < gffr.Linecount {
      ftable[i].Seqid = line.Seqid
      ftable[i].Type = line.Type
      ftable[i].Beg = line.Beg
      ftable[i].End = line.End
      ftable[i].Score = line.Score
      ftable[i].Strand = line.Strand
      ftable[i].Phase = line.Phase
      ftable[i].ID = line.ID

      if line.Parent != nil{
      ftable[i].Parent = line.Parent[0]
      }

      ftable[i].Len = ftable[i].End - ftable[i].Beg + 1
      i++

      if len(line.Parent) > 1 {
        for _, parent := range line.Parent[1:] {
          ftable[i] = ftable[i-1]
          ftable[i].Parent = parent
          i++
        }
      }
    }
  }

  fmt.Println(ftable[0])
}
