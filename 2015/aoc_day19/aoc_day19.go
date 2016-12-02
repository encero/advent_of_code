package main

import (
  "fmt"
  "regexp"
  "bufio"
  "os"
  "io"
  "strings"
)

type Replacement struct {
  from, to string
}

func (i Replacement) String() string {
  return fmt.Sprintf("%s => %s", i.from, i.to)
}

func read() []*Replacement {
  in := bufio.NewReader(os.Stdin)
  reg := regexp.MustCompile("^([a-zA-Z]+) => ([a-zA-Z]+)\n$")

  var replacements []*Replacement

  for  {
    line, err := in.ReadString('\n')

    if (err == io.EOF || line == "") {
      break
    }

    subMatch := reg.FindStringSubmatch(line)

    replacements = append(replacements, &Replacement{subMatch[1], subMatch[2]})
  }

  return replacements
}

func findReplaces(r *Replacement, molecule string, hashes map[string]bool) {
  prefix := ""
  suffix := molecule

  for {
    index := strings.Index(suffix, r.from)
    if (index == -1) { break }

    replaced := prefix + suffix[:index] + r.to + suffix[index + len(r.from):]
    hashes[replaced] = true

    prefix += suffix[:index + len(r.from)]
    suffix = suffix[index + len(r.from):]
  }
}

func reduce(molecule string, r *Replacement) (string, int) {
  prefix := ""
  suffix := molecule

  steps := 0

  for {
    index := strings.Index(suffix, r.to)
    if (index == -1) { break }

    prefix += suffix[:index] + r.from

    suffix = suffix[index + len(r.to):]

    steps ++
  }

  return prefix + suffix, steps
}

func star1(molecule string, replacements []*Replacement) {
  hashes := make(map[string]bool, 500)

  for _, v := range replacements {
    findReplaces(v, molecule, hashes)
  }

  fmt.Println(len(hashes))
}

func star2(molecule string, replacements []*Replacement) {
  sumSteps := 0

  for len(molecule) > 1 {
    for _, v := range replacements {
      reduced, steps := reduce(molecule, v)

      fmt.Println(steps > 0, reduced)

      molecule = reduced

      sumSteps += steps
    }
  }

  fmt.Println(sumSteps)
}

func main()  {
  molecule := "CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl"
  //molecule := "HOHOHO"

  replacements := read()

  star2(molecule, replacements)

}
