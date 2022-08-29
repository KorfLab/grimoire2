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

  flag.StringVar(&fname, "f", "", "path to GFF file")

  flag.Parse()

  gffreader := gffreader.Read_record(fname)

  featuretable := make([]classes.Feature, gffreader.Linecount)

  fmt.Println(featuretable)
}
