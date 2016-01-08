package main

import (
  "fmt"
)

type Entity struct {
  damage, armor, hitpoints int
}

type Item struct {
  damage, armor, cost int
}

func (i Item) String() string {return fmt.Sprintf("d:%d a:%d c:%d", i.damage, i.armor, i.cost)}

func turn(one, two *Entity) {
  two.hitpoints -= max(1, one.damage - two.armor)
}

func figth(player, boss *Entity) bool{
  for  {
    turn(player, boss)
    if boss.hitpoints <= 0 {
      return true
    }

    turn(boss, player)
    if player.hitpoints <= 0 {
      return false
    }
  }
}

func equip() chan []Item {
  weapons := []Item {
    {4, 0, 8},
    {5, 0, 10},
    {6, 0, 25},
    {7, 0, 40},
    {8, 0, 74},
  }
  armors := []Item {
    {0, 1, 13},
    {0, 2, 31},
    {0, 3, 53},
    {0, 4, 75},
    {0, 5, 102},
    {0, 0, 0},
  }

  rings := []Item {
    {1, 0, 25},
    {2, 0, 50},
    {3, 0, 100},
    {0, 1, 20},
    {0, 2, 40},
    {0, 3, 80},
    {0, 0, 0},
  }

  ch := make(chan []Item)

  go func() {
    for _, weapon := range weapons {
      for _, armor := range armors {
        for _, ring := range rings {
          for _, ring2 := range rings {
            if (ring == ring2) { ring2 = Item{0,0,0}}
            ch <- []Item {weapon, armor, ring, ring2}
          }
        }
      }
    }

    ch <- nil
  }()

  return ch
}

func main()  {

  min := 0;
  best := []Item{}

  for eq := range equip() {
    if eq == nil {
      break
    }

    boss := Entity{8, 1, 104}
    player := Entity{0, 0, 100}

    cost := 0

    for _, item := range eq {
      player.damage += item.damage
      player.armor += item.armor

      cost += item.cost
    }

    if cost <= min {
      continue
    }

    if !figth(&player, &boss) {
      min = cost
      best = eq
    }
  }

  fmt.Println(min)
  fmt.Println(best)
}



// --------------------

func max(a, b int) int {
  if (a > b) {
    return a
  } else {
    return b
  }
}
