package main

import (
  "fmt"
  "regexp"
  "bufio"
  "os"
  "io"
  "strconv"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)

type City struct {
  distances map[string]int
}

var cities map[string]*City
var max int

func init() {
  cities = make(map[string]*City)
  max = 0
}

func getCity(name string) *City {
  if city, ok := cities[name]; ok {
    return city
  } else {
    city := City{make(map[string]int)}
    cities[name] = &city

    return &city
  }
}

func load(from string, to string, distance int) {
  fromCity := getCity(from)
  toCity := getCity(to)

  fromCity.distances[to] = distance
  toCity.distances[from] = distance
}

func getNames() []string {
  names := make([]string, len(cities))
  i := 0
  for k := range cities {
    names[i] = k
    i++
  }

  return names
}

func dist(from string, to string) int {
  if from == "" {
    return 0
  }

  return getCity(from).distances[to]
}

func find(name string, names []string, visited []string, distance int) {
  if len(names) == 0 {
    fmt.Printf("%v = %d\n", visited, distance)
    if (distance > max) {
      max = distance
    }
    return
  }

  for k, v := range names {
    find(v, append(append([]string(nil), names[:k]...), names[k+1:]...), append(append([]string(nil), visited...), v), distance + dist(name, v))
  }
}

func main()  {
  in := bufio.NewReader(os.Stdin)

  reg := regexp.MustCompile("^(.+) to (.+) = ([0-9]+)\n$")

  for  {
    line, err := in.ReadString('\n')
    if (err == io.EOF) { break }

    match := reg.FindStringSubmatch(line)

    distance, _ := strconv.Atoi(match[3])

    load(match[1], match[2], distance)
  }

  names := getNames()

  fmt.Printf("--- %v ----\n", names)

  find("", names, []string(nil), 0)

  fmt.Printf("max: %d end\n", max);
}
