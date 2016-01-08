package main

import (
  //"fmt"
)

type Entity struct {
  hitpoints, armor, damage, mana, used int
  effects []*Spell
}

func (p *Entity) cast(e Spell, opponent *Entity) {
  if p.canCast(e) {
    //fmt.Println("Casting", e.name, "for", e.cost,"mana")
    if e.effect {
      p.effects = append(p.effects, &e)
    } else {
      p.applySpell(&e, opponent)
    }

    p.mana -= e.cost
    p.used += e.cost
  } else {
    //fmt.Println("!!!! Cannot cast", e.name , "already in effects")
  }
}

func (p *Entity) canCast(e Spell) bool {
  if p.mana < e.cost {
    return false
  }

  for _, v := range p.effects {
    if v.name == e.name {
      return false
    }
  }

  return true
}

func (e *Entity) applySpell(s *Spell, opponent *Entity) bool {
  switch s.name {
    case "drain":
      opponent.hitpoints -= 2
      e.hitpoints += 2
    case "missile":
      opponent.hitpoints -= 4
    case "shield":
      e.armor = 7
    case "poison":
      opponent.hitpoints -= 3
    case "recharge":
      e.mana += 101
  }

  if s.effect {
    s.timer --
    //fmt.Println("Applying effect", s.name, "times is", s.timer)
  } else {
    //fmt.Println("Applying spell", s.name)
  }

  return !(s.effect && s.timer > 0)
}

func (e *Entity) attack(opponent *Entity) {
  damage := max(1, e.damage - opponent.armor)

  //fmt.Println("Attacking for", damage, "damage")

  opponent.hitpoints -= damage
}

func (p *Entity) turnStart(opponent *Entity) {
  for i := len(p.effects) - 1; i >= 0; i-- {
    eff := p.effects[i]

    if eff.effect {
      if p.applySpell(eff, opponent) {
        //fmt.Println("Effect", eff.name, "worn off")
        p.effects = append(p.effects[:i], p.effects[i + 1:]...)
      }
    }
  }
}

func (entity *Entity) clone() Entity {
  new := *entity
  new.effects = make([]*Spell, len(entity.effects))

  for i := 0; i < len(entity.effects); i++ {
    copy := *entity.effects[i]
    new.effects[i] = &copy
  }

  return new
}
