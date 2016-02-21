package main

import (
	"fmt"
	"strings"
)

const X = -1

// --- Arena ---

type Arena struct {
	id       int
	name     string
	trophies int
}

func (a *Arena) String() string {
	return fmt.Sprintf("Arena %d: %s", a.id, a.name)
}

var (
	ARENA_0 = &Arena{0, "Training Camp", 0}
	ARENA_1 = &Arena{1, "Goblin Stadium", 0}
	ARENA_2 = &Arena{2, "Bone Pit", 400}
	ARENA_3 = &Arena{3, "Barbarian Bowl", 800}
	ARENA_4 = &Arena{4, "P.E.K.K.A's Playhouse", 1100}
	ARENA_5 = &Arena{5, "Spell Valley", 1400}
	ARENA_6 = &Arena{6, "Royal Arena", 1700}
	ARENA_7 = &Arena{7, "Legendary Arena", 3000}
)

var ARENAS = [...]*Arena{
	ARENA_0,
	ARENA_1,
	ARENA_2,
	ARENA_3,
	ARENA_4,
	ARENA_5,
	ARENA_6,
	ARENA_7,
}

// --- Rarity ---

type RarityAttribute string

const (
	CARDS_REQ = RarityAttribute("Cards Required")
	GOLD_REQ  = RarityAttribute("Gold Required")
	EXP_GAIN  = RarityAttribute("Experience Gained")
)

var RARITY_ATTRIBUTES = [...]RarityAttribute{
	CARDS_REQ,
	GOLD_REQ,
	EXP_GAIN,
}

type Rarity struct {
	name  string
	cards []int
	gold  []int
	exp   []int
}

func (r *Rarity) String() string {
	return r.name
}

var (
	COMMON = &Rarity{
		name:  "Common",
		cards: []int{0, 2, 4, 10, 20, 50, 100, X, X, X, X, X},
		gold:  []int{0, 5, 20, 50, 150, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 4, 5, 6, 10, 25, 50, X, X, X, X, X},
	}
	RARE = &Rarity{
		name:  "Rare",
		cards: []int{0, 2, 4, 10, 20, 50, 100, X, X, X},
		gold:  []int{0, 50, 150, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 6, 10, 25, 50, X, X, X, X, X},
	}
	EPIC = &Rarity{
		name:  "Epic",
		cards: []int{0, 2, 4, 10, 20, 50, 100, X},
		gold:  []int{0, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 6, 10, 25, 50, X, X, X},
	}
)

var RARITIES = [...]*Rarity{
	COMMON,
	RARE,
	EPIC,
}

// --- Type ---

type Type string

const (
	TROOP    = Type("Troop")
	BUILDING = Type("Building")
	SPELL    = Type("Spell")
)

var TYPES = [...]Type{
	TROOP,
	BUILDING,
	SPELL,
}

// --- Targets ---

type Targets string

const (
	GROUND         = Targets("Ground")
	AIR_AND_GROUND = Targets("Air & Ground")
	BUILDINGS      = Targets("Buildings")
)

var TARGETSES = [...]Targets{
	GROUND,
	AIR_AND_GROUND,
	BUILDINGS,
}

// --- Speed ---

type Speed string

const (
	SLOW      = Speed("Slow")
	MEDIUM    = Speed("Medium")
	FAST      = Speed("Fast")
	VERY_FAST = Speed("Very Fast")
)

var SPEEDS = [...]Speed{
	SLOW,
	MEDIUM,
	FAST,
	VERY_FAST,
}

// --- Range ---

const MELEE = 0

// --- Card ---

type CardAttribute struct {
	name      string
	isMutable bool
	format    func(x interface{}) string
}

func (attr *CardAttribute) String() string {
	return attr.name
}

var (
	NAME = &CardAttribute{
		"Name",
		false,
		func(x interface{}) string {
			return x.(string)
		},
	}
	ARENA = &CardAttribute{
		"Arena",
		false,
		func(x interface{}) string {
			return x.(*Arena).String()
		},
	}
	RARITY = &CardAttribute{
		"Rarity",
		false,
		func(x interface{}) string {
			return x.(*Rarity).String()
		},
	}
	TYPE = &CardAttribute{
		"Type",
		false,
		func(x interface{}) string {
			return string(x.(Type))
		},
	}
	DESC = &CardAttribute{
		"Description",
		false,
		func(x interface{}) string {
			return x.(string)
		},
	}
	COST = &CardAttribute{
		"Elixir Cost",
		false,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	HP = &CardAttribute{
		"Hitpoints",
		true,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	DPS = &CardAttribute{
		"Damage per Second",
		true,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	DAM = &CardAttribute{
		"Damage",
		true,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	ADAM = &CardAttribute{
		"Area Damage",
		true,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	DDAM = &CardAttribute{
		"Death Damage",
		true,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	SKE_LV = &CardAttribute{
		"Skeleton Level",
		true,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	SGO_LV = &CardAttribute{
		"Spear Goblin Level",
		true,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
	HSPD = &CardAttribute{
		"Hit Speed",
		false,
		func(x interface{}) string {
			switch x.(type) {
			case int:
				return fmt.Sprintf("%dsec", x)
			case float64:
				return fmt.Sprintf("%.1fsec", x)
			default:
				return "???"
			}
		},
	}
	TGTS = &CardAttribute{
		"Targets",
		false,
		func(x interface{}) string {
			return string(x.(Targets))
		},
	}
	SPD = &CardAttribute{
		"Speed",
		false,
		func(x interface{}) string {
			return string(x.(Speed))
		},
	}
	RNG = &CardAttribute{
		"Range",
		false,
		func(x interface{}) string {
			switch x.(type) {
			case int:
				if x.(int) == MELEE {
					return "Melee"
				}
				return fmt.Sprintf("%d", x)
			case float64:
				return fmt.Sprintf("%.1f", x)
			default:
				return "???"
			}
		},
	}
	DTIME = &CardAttribute{
		"Deploy Time",
		false,
		func(x interface{}) string {
			switch x.(type) {
			case int:
				return fmt.Sprintf("%dsec", x)
			case float64:
				return fmt.Sprintf("%.1fsec", x)
			default:
				return "???"
			}
		},
	}
	COUNT = &CardAttribute{
		"Count",
		false,
		func(x interface{}) string {
			return fmt.Sprintf("%d", x.(int))
		},
	}
)

var CARD_ATTRIBUTES = [...]*CardAttribute{
	NAME,
	ARENA,
	RARITY,
	TYPE,
	DESC,
	COST,
	HP,
	DPS,
	DAM,
	ADAM,
	DDAM,
	SKE_LV,
	SGO_LV,
	HSPD,
	TGTS,
	SPD,
	RNG,
	DTIME,
	COUNT,
}

type Card map[*CardAttribute]interface{}

var (
	KNIGHT = Card{
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
	}
	ARCHERS = Card{
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
		RNG:    5.5,
		DTIME:  1,
		COUNT:  2,
	}
	BOMBER = Card{
		NAME:   "Bomber",
		ARENA:  ARENA_0,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   "Small, lightly protected skeleton that throws bombs. Deals damage that can wipe out a swarm of enemies.",
		COST:   3,
		HP:     []int{150, 165, 181, 199, 219, 240, 264, X, 318, 349, 384, 421},
		DPS:    []int{52, 57, 63, 70, 76, 84, 92, X, 111, 122, 134, 147},
		ADAM:   []int{100, 110, 121, 133, 146, 160, 176, X, 212, 233, 256, 281},
		HSPD:   1.9,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    5,
		DTIME:  1,
	}
	GIANT = Card{
		NAME:   "Giant",
		ARENA:  ARENA_0,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   "Slow but durable, only attacks buildings. A real one-man wrecking crew!",
		COST:   5,
		HP:     []int{2000, 2200, 2420, 2660, X, X, X, X, X, X},
		DPS:    []int{80, 88, 96, 106, X, X, X, 154, X, X},
		DAM:    []int{120, 132, 145, 159, X, X, X, 231, X, X},
		HSPD:   1.5,
		TGTS:   BUILDINGS,
		SPD:    SLOW,
		RNG:    MELEE,
		DTIME:  1,
	}
	MUSKETEER = Card{
		NAME:   "Musketeer",
		ARENA:  ARENA_0,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   "Don't be fooled by her delicately coiffed hair, the musketeer is a mean shot with her trusty boomstick.",
		COST:   4,
		HP:     []int{340, 374, 411, 452, 496, X, X, 656, X, X},
		DPS:    []int{81, 90, 98, 108, 119, X, X, X, X, X},
		DAM:    []int{90, 99, 108, 119, 131, X, X, X, X, X, X},
		HSPD:   1.1,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RNG:    6.5,
		DTIME:  1,
	}
)

var CARDS = [...]Card{
	KNIGHT,
	ARCHERS,
	BOMBER,
	GIANT,
	MUSKETEER,
}

// --- Main ---

const (
	attrTitle     = "Attribute"
	valueTitle    = "Value"
	attrTitleLen  = len(attrTitle)
	valueTitleLen = len(valueTitle)
	attrValueLen  = 4
)

func main() {
	fixedAttrNameLen := attrTitleLen
	mutableAttrNameLen := attrTitleLen
	for _, attr := range CARD_ATTRIBUTES {
		attrNameLen := len(attr.name)
		if attr.isMutable {
			if attrNameLen > mutableAttrNameLen {
				mutableAttrNameLen = attrNameLen
			}
		} else {
			if attrNameLen > fixedAttrNameLen {
				fixedAttrNameLen = attrNameLen
			}
		}
	}
	for _, attr := range RARITY_ATTRIBUTES {
		attrNameLen := len(attr)
		if attrNameLen > mutableAttrNameLen {
			mutableAttrNameLen = attrNameLen
		}
	}

	for _, card := range CARDS {
		// Header
		fmt.Printf("### %s\n", card[NAME])
		fmt.Println()

		// Fixed Attribute Table
		fmt.Printf("%*s | %s\n", -fixedAttrNameLen, attrTitle, valueTitle)
		fmt.Printf("%*s | %s\n", fixedAttrNameLen, strings.Repeat("-", fixedAttrNameLen), strings.Repeat("-", valueTitleLen))
		for _, attr := range CARD_ATTRIBUTES {
			if attr.isMutable {
				continue
			}
			if value, ok := card[attr]; ok {
				fmt.Printf("%*s | %s\n", -fixedAttrNameLen, attr, attr.format(value))
			}
		}
		fmt.Println()

		// Mutable Attribute Table
		fmt.Printf("%*s ", -mutableAttrNameLen, attrTitle)
		// Any field will do, not just "cards".
		for level := range card[RARITY].(*Rarity).cards {
			fmt.Printf("| %*s ", -attrValueLen, fmt.Sprintf("LV%d", level+1))
		}
		fmt.Println()

		fmt.Printf("%*s ", mutableAttrNameLen, strings.Repeat("-", mutableAttrNameLen))
		// Any field will do, not just "cards".
		for range card[RARITY].(*Rarity).cards {
			fmt.Printf("| %*s ", attrValueLen, strings.Repeat("-", attrValueLen))
		}
		fmt.Println()

		for _, attr := range CARD_ATTRIBUTES {
			if !attr.isMutable {
				continue
			}
			if _, ok := card[attr]; ok {
				fmt.Printf("%*s |\n", -mutableAttrNameLen, attr)
			}
		}

		fmt.Printf("%*s ", -mutableAttrNameLen, CARDS_REQ)
		for _, cardsReq := range card[RARITY].(*Rarity).cards {
			fmt.Printf("| %*d ", attrValueLen, cardsReq)
		}
		fmt.Println()

		fmt.Printf("%*s ", -mutableAttrNameLen, GOLD_REQ)
		for _, goldReq := range card[RARITY].(*Rarity).gold {
			fmt.Printf("| %*d ", attrValueLen, goldReq)
		}
		fmt.Println()

		fmt.Printf("%*s ", -mutableAttrNameLen, EXP_GAIN)
		for _, expGain := range card[RARITY].(*Rarity).exp {
			fmt.Printf("| %*d ", attrValueLen, expGain)
		}
		fmt.Println()

		fmt.Println()
	}
}
