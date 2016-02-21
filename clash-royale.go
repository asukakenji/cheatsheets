package main

import (
	"fmt"
)

const X = -1

type Arena struct {
	id       int
	name     string
	trophies int
}

var ARENAS = [...]Arena{
	Arena{id: 0, name: "Training Camp", trophies: 0},
	Arena{id: 1, name: "Goblin Stadium", trophies: 0},
	Arena{id: 2, name: "Bone Pit", trophies: 0},
	Arena{id: 3, name: "Barbarian Bowl", trophies: 0},
	Arena{id: 4, name: "P.E.K.K.A's Playhouse", trophies: 0},
	Arena{id: 5, name: "Spell Valley", trophies: 0},
	Arena{id: 6, name: "Royal Arena", trophies: 0},
	Arena{id: 7, name: "Legendary Arena", trophies: 0},
}

var (
	ARENA_0 = &ARENAS[0]
	ARENA_1 = &ARENAS[1]
	ARENA_2 = &ARENAS[2]
	ARENA_3 = &ARENAS[3]
	ARENA_4 = &ARENAS[4]
	ARENA_5 = &ARENAS[5]
	ARENA_6 = &ARENAS[6]
	ARENA_7 = &ARENAS[7]
)

//type RarityAttribute string
//
//const (
//	CARDS RarityAttribute = "Cards Required"
//	GOLD                  = "Gold Required"
//	EXP                   = "Experience Gained"
//)

type Rarity struct {
	id    int
	name  string
	cards []int
	gold  []int
	exp   []int
}

var RARITIES = [...]Rarity{
	Rarity{
		id:    0,
		name:  "Common",
		cards: []int{0, 2, 4, 10, 20, 50, 100, X, X, X, X, X},
		gold:  []int{0, 5, 20, 50, 150, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 4, 5, 6, 10, 25, 50, X, X, X, X, X},
	},
	Rarity{
		id:    1,
		name:  "Rare",
		cards: []int{0, 2, 4, 10, 20, 50, 100, X, X, X},
		gold:  []int{0, 50, 150, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 6, 10, 25, 50, X, X, X, X, X},
	},
	Rarity{
		id:    2,
		name:  "Epic",
		cards: []int{0, 2, 4, 10, 20, 50, 100, X},
		gold:  []int{0, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 6, 10, 25, 50, X, X, X},
	},
}

var (
	COMMON = &RARITIES[0]
	RARE   = &RARITIES[1]
	EPIC   = &RARITIES[2]
)

type Type string

const (
	TROOP    Type = "Troop"
	BUILDING      = "Building"
	SPELL         = "Spell"
)

type Targets string

const (
	GROUND         Targets = "Ground"
	AIR_AND_GROUND         = "Air & Ground"
	BUILDINGS              = "Buildings"
)

type Speed string

const (
	SLOW      Speed = "Slow"
	MEDIUM          = "Medium"
	FAST            = "Fast"
	VERY_FAST       = "Very Fast"
)

const MELEE = 0

type CardAttribute string

const (
	NAME   CardAttribute = "Name"
	ARENA                = "Arena"
	RARITY               = "Rarity"
	TYPE                 = "Type"
	DESC                 = "Description"
	COST                 = "Elixir Cost"
	HP                   = "Hitpoints"
	DPS                  = "Damage per Second"
	DAM                  = "Damage"
	ADAM                 = "Area Damage"
	HSPD                 = "Hit Speed"
	TGTS                 = "Targets"
	SPD                  = "Speed"
	RNG                  = "Range"
	DTIME                = "Deploy Time"
	COUNT                = "Count"
)

type Card map[CardAttribute]interface{}

var CARDS = map[string]Card{
	"KNIGHT": Card{
		NAME:   "Knight",
		ARENA:  ARENA_0,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   "A tough melee fighter. The Barbarian's handsome, cultured cousin. Rumor has it that he was knighted based on the sheer awesomeness of his mustache alone.",
		COST:   3,
		HP:     []int{600, 660, 726, 798, 876, X, 1056, X, 1272, 1398, X, 1686},
		DPS:    []int{68, 74, 81, 90, 99, X, 120, X, 144, 158, X, 190},
		DAM:    []int{75, 82, 90, 99, 109, X, 132, X, 159, 174, X, 210},
		HSPD:   1.1,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    MELEE,
		DTIME:  1,
	},
	"ARCHERS": Card{
		NAME:   "Archers",
		ARENA:  ARENA_0,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   "A pair of unarmored ranged attackers. They'll help you with ground and air unit attacks, but you're on your own with coloring your hair.",
		COST:   3,
		HP:     []int{125, 137, 151, 166, X, 200, 220, X, 265, 291, 320, 351},
		DPS:    []int{33, 36, 40, 44, X, 53, 58, X, 70, 77, 85, 93},
		DAM:    []int{40, 44, 48, 53, X, 64, 70, X, 84, 93, 102, 112},
		HSPD:   1.2,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RANGE:  5.5,
		DTIME:  1,
		COUNT:  2,
	},
}

var (
	KNIGHT = CARDS["KNIGHT"]
)

func main() {
	fmt.Println("Hello")
}
