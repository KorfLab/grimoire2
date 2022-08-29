package gffreader

import(
  "bufio"
  "compress/gzip"
  "fmt"
  "grimoire2/feature"
  "os"
)


type Record struct {
  Seqid string
  Source string
  Type string
  Beg int
  End int
  Score float64
  Strand byte
  Phase byte
  ID string
  Parent []string
}

type Iterator interface {
  Record() Record
  Next() bool
}

type state_Iterator struct {
  scanner *bufio.Scanner
  gzipfh *gzip.Reader
  filehandle *os.File
  current *Record
  done bool
}


func (si *state_iterator) Record() *Record {
  return si.current
}

func (si *state_iterator) Next() bool {
  if si.done {
    si.filehandle.Close()

    if si.gzipfh != nil {
      it.gzipfh.Close()
    }

  return false
  }

  if si.scanner.Scan() == true {
    line := scanner.Text()
    c := strings.Fields(line)

    si.current.Seqid = c[0]
    si.current.Source = c[1]
    si.current.Type = c[2]
    si.current.Beg = strconv.ParseInt(c[3], 10, 32)
    si.current.End = strconv.ParseInt(c[4], 10, 32)
    si.current.Score = strconv.ParseFloat(c[5], 64)
    si.current.Strand = byte(c[6])
    si.current.Phase = byte(c[7])
    attributes := strings.Split(c[8], ";")

    for _, att i*-+n range attributes {
      if strings.HasPrefix(att, "Parent=") {
        si.current.Parent = strings.Split(att, ",")

      } else if strings.HasPrefix(att, "ID=") {
        si.current.ID = att
      }
    }
  } else {
    si.done = true

    return true
  }
}

func Read_record(fname string) *state_Iterator {
  filehandle, err := os.Open(fname)
  if err != nil {
    panic(err)
  }

  si = new(state_Iterator)
  si.current = new(Record)
  si.done = false
  si.filehandle = filehandle

  scanner := bufio.NewScanner(filehandle)
  if strings.HasSuffix(fname, ".gz") {
    gzipfh, err := gzip.NewReader(filehandle)
    si.gzipfh = gzipfh

    if err != nil {
      panic(err)
    }

    scanner = bufio.Newscanner(gzipfh)
  }

  si.scanner = scanner
  return si
}
