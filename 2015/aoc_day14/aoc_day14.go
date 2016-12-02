package main

import (
  "fmt"
  "regexp"
  "bufio"
  "os"
  "io"
  "strconv"
  "math"
//"strings"
)

type Deer struct {
  name string
  speed, flyTime, restTime, points, distance int
}

func (d Deer) String() string {
  return fmt.Sprintf("d:%d p:%d \n", d.distance, d.points)
}

func toInt(number string) int {
  value, _ := strconv.Atoi(number)
  return value
}

func CreateDeer(spec []string) *Deer {
  return &Deer{spec[0], toInt(spec[1]), toInt(spec[2]), toInt(spec[3]), 0, 0}
}

func read() map[string]*Deer {
  in := bufio.NewReader(os.Stdin)

//Rudolph can fly 22 km/s for 8 seconds, but then must rest for 165 seconds.
  reg := regexp.MustCompile(`^(\w+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds.\n$`)

  deers := make(map[string]*Deer)

  for  {
    line, err := in.ReadString('\n')
    if err == io.EOF || line == "" {
      break
    }

    match := reg.FindStringSubmatch(line)

    deers[match[1]] = CreateDeer(match[1:])
  }

  return deers
}

func step(seconds int, deers map[string]*Deer) {
  for _, deer := range deers {
    fullCycle := float64(seconds/(deer.flyTime + deer.restTime))
    reminderTime := float64(math.Mod(float64(seconds), float64(deer.flyTime + deer.restTime)))

    deer.distance = int(fullCycle * float64(deer.speed) * float64(deer.flyTime) + math.Min(float64(deer.flyTime), reminderTime) * float64(deer.speed))
  }
}

func awardBestDeer(deers map[string]*Deer) {
  best := 0

  for _, v := range deers {
    if v.distance > best {
      best = v.distance
    }
  }

  for _, v := range deers {
    if v.distance == best {
      v.points += 1
    }
  }

}

func main()  {
  deers := read()

  for i := 1; i <= 2503; i++ {
    step(i, deers)
    awardBestDeer(deers)
  }

  fmt.Println(deers)

  fmt.Printf("end\n");
}
