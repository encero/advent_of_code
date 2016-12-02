package main

import (
  "fmt"
  "regexp"
  "bufio"
  "os"
  "io"
//  "strings"
  "strconv"
)

type Person struct {
  name string
  happiness map[string]int
}

func CreatePerson(persons map[string]*Person, name string, nextTo string, happiness int) {
  person, ok := persons[name]
  if !ok {
    person = &Person{name, make(map[string]int)}
    persons[name] = person
  }

  person.happiness[nextTo] = happiness
}

func read() []*Person {
  persons := make(map[string]*Person)
  in := bufio.NewReader(os.Stdin)

  //Alice would gain 54 happiness units by sitting next to Bob.
  reg := regexp.MustCompile("^([a-zA-Z]+) would (gain|lose) ([0-9]+) happiness units by sitting next to ([a-zA-Z]+)\\.\n$")

  for  {
    line, err := in.ReadString('\n')
    if (err == io.EOF || line == "") {
      break
    }

    match := reg.FindStringSubmatch(line)

    happiness, _ := strconv.Atoi(match[3])

    if (match[2] == "lose") {
      happiness = -happiness
    }

    CreatePerson(persons, match[1], match[4], happiness )
  }

  list := make([]*Person, len(persons))
  index := 0
  for _, v := range persons {
    list[index] = v
    index++
  }

  return list
}

func happiness(setup []*Person) int {
  prev := setup[len(setup) - 1]
  sum := 0

  for _, v := range setup {
    sum += v.happiness[prev.name]
    sum += prev.happiness[v.name]

    prev = v
  }

  return sum
}

func main()  {
  persons := read()

  me := Person{"me", make(map[string]int)}

  for _, v := range persons {
    v.happiness["me"] = 0
    me.happiness[v.name] = 0
  }

  persons = append(persons, &me)

  max := 0

  ch := permutations(persons, len(persons))

  for permu := range ch {
    if permu == nil { break }
    happ := happiness(permu)
    if (happ > max) {
      max = happ
    }
  }

  fmt.Printf("max: %d\n", max);
}
