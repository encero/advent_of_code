package main

import (
  "fmt"
//  "bufio"
//  "os"
)

var effects []Spell = []Spell {
  Spell{"missile", 53, 0, false}, //0
  Spell{"drain", 73, 0, false}, //1
  Spell{"shield", 113, 6, true}, //2
  Spell{"poison", 173, 6, true}, //3
  Spell{"recharge", 229, 5, true}, //4
}

type Spell struct {
  name string
  cost,timer int
  effect bool
}

func turn(attacker *Entity, defender *Entity, spell *Spell) bool {
  //reset armor
  attacker.armor = 0
  defender.armor = 0

  attacker.turnStart(defender)
  defender.turnStart(attacker)

  if attacker.hitpoints > 0 && defender.hitpoints > 0 {
    // if no spell attack!
    if spell != nil {
      if !attacker.canCast(*spell) {
        return false
      }

      attacker.cast(*spell, defender)
    } else {
      attacker.attack(defender)
    }
  }

  return true
}

func show(player *Entity, boss *Entity) {
    //fmt.Println("- Player has", player.hitpoints, "hitpoints and", player.mana, "mana used", player.used)
    //fmt.Println("- Boss has", boss.hitpoints, "hitpoints")
}

func playerTurn(_player Entity, _boss Entity, minCost *int) {
  if _player.used > *minCost {
    return
  }

  for _, v := range effects {
    player := _player.clone()
    boss := _boss.clone()

    //fmt.Println("-- Player turn")
    show(&player, &boss)
    player.hitpoints -= 1
    if !turn(&player, &boss, &v) {
      //fmt.Println("Cannot cast", v.name)
      continue
    }

    if (boss.hitpoints > 0) {
      bossTurn(player, boss, minCost)
    } else {
      fmt.Println("--- Boss died ---")
      if *minCost > player.used {
        *minCost = player.used
      }
    }
  }
}

func bossTurn(player Entity, boss Entity, minCost *int) {
  //fmt.Println("-- Boss turn")
  show(&player, &boss)
  turn(&boss, &player, nil)

  if player.hitpoints > 0 {
    playerTurn(player, boss, minCost)
  } else {
    //fmt.Println("--- Player died ---\n\n")
    //bufio.NewReader(os.Stdin).ReadBytes('\n')
  }

}

func main()  {

  if true {
    cost := 100000000

    player := Entity{50, 0, 0, 500, 0, make([]*Spell, 0)}
    boss := Entity{71, 0, 10, 0, 0, make([]*Spell, 0)}

    playerTurn(player, boss, &cost)

    fmt.Println("minimal cost", cost)
  } else {
    test()
  }



  fmt.Println("end")
}

func test() {
  player := Entity{10, 0, 0, 250, 0, make([]*Spell, 0)}
  boss := Entity{14, 0, 8, 0, 0, make([]*Spell, 0)}

  playerTurn(player, boss, nil)

  return

  prepared := []Spell{effects[3], effects[0]}

  for _, v := range prepared {
    //fmt.Println("-- Player turn")
    show(&player, &boss)
    turn(&player, &boss, &v)

    //fmt.Println("-- Boss turn")
    show(&player, &boss)
    turn(&boss, &player, nil)
  }

  fmt.Println("-------")
  show(&player, &boss)
}



// --------------------

func max(a, b int) int {
  if (a > b) {
    return a
  } else {
    return b
  }
}
