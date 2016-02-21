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
		cards: []int{0, 2, 4, 10, 20, 50, 100, 200, X, X, X, X},
		gold:  []int{0, 5, 20, 50, 150, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 4, 5, 6, 10, 25, 50, X, X, X, X, X},
	}
	RARE = &Rarity{
		name:  "Rare",
		cards: []int{0, 2, 4, 10, 20, 50, 100, 200, X, X},
		gold:  []int{0, 50, 150, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 6, 10, 25, 50, X, X, X, X, X},
	}
	EPIC = &Rarity{
		name:  "Epic",
		cards: []int{0, 2, 4, 10, 20, 50, 100, 200},
		gold:  []int{0, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 25, 50, X, X, X, X, X},
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
	format    func(value interface{}) string
}

func (attr *CardAttribute) String() string {
	return attr.name
}

// --- Format Functions ---

func formatString(value interface{}) string {
	return value.(string)
}

func formatInt(value interface{}) string {
	if X == value.(int) {
		return ""
	}
	return fmt.Sprintf("%d", value.(int))
}

func formatTime(value interface{}) string {
	switch value.(type) {
	case int:
		return fmt.Sprintf("%dsec", value)
	case float64:
		return fmt.Sprintf("%.1fsec", value)
	default:
		panic("Unknown value type")
	}
}

var (
	NAME = &CardAttribute{
		"Name",
		false,
		formatString,
	}
	ARENA = &CardAttribute{
		"Arena",
		false,
		func(value interface{}) string {
			return value.(*Arena).String()
		},
	}
	RARITY = &CardAttribute{
		"Rarity",
		false,
		func(value interface{}) string {
			return value.(*Rarity).String()
		},
	}
	TYPE = &CardAttribute{
		"Type",
		false,
		func(value interface{}) string {
			return string(value.(Type))
		},
	}
	DESC = &CardAttribute{
		"Description",
		false,
		formatString,
	}
	COST = &CardAttribute{
		"Elixir Cost",
		false,
		formatInt,
	}
	HP = &CardAttribute{
		"Hitpoints",
		true,
		formatInt,
	}
	DPS = &CardAttribute{
		"Damage per Second",
		true,
		formatInt,
	}
	DAM = &CardAttribute{
		"Damage",
		true,
		formatInt,
	}
	ADAM = &CardAttribute{
		"Area Damage",
		true,
		formatInt,
	}
	DDAM = &CardAttribute{
		"Death Damage",
		true,
		formatInt,
	}
	SKE_LV = &CardAttribute{
		"Skeleton Level",
		true,
		formatInt,
	}
	SGO_LV = &CardAttribute{
		"Spear Goblin Level",
		true,
		formatInt,
	}
	HSPD = &CardAttribute{
		"Hit Speed",
		false,
		formatTime,
	}
	TGTS = &CardAttribute{
		"Targets",
		false,
		func(value interface{}) string {
			return string(value.(Targets))
		},
	}
	SPD = &CardAttribute{
		"Speed",
		false,
		func(value interface{}) string {
			return string(value.(Speed))
		},
	}
	RNG = &CardAttribute{
		"Range",
		false,
		func(value interface{}) string {
			switch value.(type) {
			case int:
				if value.(int) == MELEE {
					return "Melee"
				}
				return fmt.Sprintf("%d", value)
			case float64:
				return fmt.Sprintf("%.1f", value)
			default:
				panic("Unknown value type")
			}
		},
	}
	DTIME = &CardAttribute{
		"Deploy Time",
		false,
		formatTime,
	}
	COUNT = &CardAttribute{
		"Count",
		false,
		func(value interface{}) string {
			return fmt.Sprintf("x %d", value.(int))
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
		HP:     []int{125, 137, 151, 166, 182, 200, 220, X, 265, 291, 320, 351},
		DPS:    []int{33, 36, 40, 44, 48, 53, 58, X, 70, 77, 85, 93},
		DAM:    []int{40, 44, 48, 53, 58, 64, 70, X, 84, 93, 102, 112},
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
	GOBLINS = Card{
		NAME:   "Goblins",
		ARENA:  ARENA_1,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   "Three fast, unarmored melee attackers. Small, fast, green and mean!",
		COST:   2,
		HP:     []int{80, 88, 96, 106, 116, 128, X, X, 169, 186, 204, 224},
		DPS:    []int{45, 50, 54, 60, 66, 72, X, X, 96, 105, 116, 127},
		DAM:    []int{50, 55, 60, 66, 73, 80, X, X, 106, 116, 128, 140},
		HSPD:   1.1,
		TGTS:   GROUND,
		SPD:    VERY_FAST,
		RNG:    MELEE,
		DTIME:  1,
		COUNT:  3,
	}
	SPEAR_GOBLINS = Card{
		NAME:   "Spear Goblins",
		ARENA:  ARENA_1,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   "Three unarmored ranged attackers. Who the heck taught these guys to throw spears!?! Who thought that was a good idea?!",
		COST:   2,
		HP:     []int{52, 57, 62, 69, 75, 83, 91, X, 110, 121, 133, 146},
		DPS:    []int{18, 20, 22, 23, 26, 29, 32, X, 38, 42, 46, 51},
		DAM:    []int{24, 26, 29, 31, 35, 38, 42, X, 50, 55, 61, 67},
		HSPD:   1.3,
		TGTS:   AIR_AND_GROUND,
		SPD:    VERY_FAST,
		RNG:    5.5,
		DTIME:  1,
		COUNT:  3,
	}
	SKELETONS = Card{
		NAME:   "Skeletons",
		ARENA:  ARENA_2,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   "Four fast, very weak melee fighters. Swarm your enemies with this pile of bones!",
		COST:   1,
		HP:     []int{30, 33, 36, 39, 43, 48, 52, X, X, X, X, X},
		DPS:    []int{30, 33, 36, 39, 43, 48, 52, X, X, X, X, X},
		DAM:    []int{30, 33, 36, 39, 43, 48, 52, X, X, X, X, X},
		HSPD:   1,
		TGTS:   GROUND,
		SPD:    FAST,
		RNG:    MELEE,
		DTIME:  1,
		COUNT:  4,
	}
	MINIONS = Card{
		NAME:   "Minions",
		ARENA:  ARENA_2,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   "Three fast, unarmored flying attackers. Roses are red, minions are blue, they can fly, and will crush you!",
		COST:   3,
		HP:     []int{90, 99, 108, 119, 131, 144, X, X, 190, 209, 230, 252},
		DPS:    []int{40, 44, 48, 53, 58, 64, X, X, 84, 93, 102, 112},
		DAM:    []int{40, 44, 48, 53, 58, 64, X, X, 84, 93, 102, 112},
		HSPD:   1,
		TGTS:   AIR_AND_GROUND,
		SPD:    FAST,
		RNG:    2.5,
		DTIME:  1,
		COUNT:  3,
	}
	BARBARIANS = Card{
		NAME:   "Barbarians",
		ARENA:  ARENA_3,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   "A horde of melee attackers with mean mustaches and even meaner tempers.",
		COST:   5,
		HP:     []int{300, X, X, X, X, 480, 528, X, X, 699, 768, 843},
		DPS:    []int{50, X, X, X, X, 80, 88, X, X, 116, 128, 140},
		DAM:    []int{75, X, X, X, X, 120, 132, X, X, 174, 192, 210},
		HSPD:   1.5,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    MELEE,
		DTIME:  1,
		COUNT:  4,
	}
	MINION_HORDE = Card{
		NAME:   "Minion Horde",
		ARENA:  ARENA_4,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   "Six fast, unarmored flying attackers. Three's a crowd, six is a horde!",
		COST:   5,
		HP:     []int{90, 99, 108, 119, 131, 144, X, X, 190, 209, 230, 252},
		DPS:    []int{40, 44, 48, 53, 58, 64, X, X, 84, 93, 102, 112},
		DAM:    []int{40, 44, 48, 53, 58, 64, X, X, 84, 93, 102, 112},
		HSPD:   1,
		TGTS:   AIR_AND_GROUND,
		SPD:    FAST,
		RNG:    2.5,
		DTIME:  1,
		COUNT:  6,
	}
	GIANT = Card{
		NAME:   "Giant",
		ARENA:  ARENA_0,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   "Slow but durable, only attacks buildings. A real one-man wrecking crew!",
		COST:   5,
		HP:     []int{2000, 2200, 2420, 2660, X, X, X, 3860, X, 4660},
		DPS:    []int{80, 88, 96, 106, X, X, X, 154, X, 186},
		DAM:    []int{120, 132, 145, 159, X, X, X, 231, X, 279},
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
		HP:     []int{340, 374, 411, 452, 496, X, X, 656, 720, 792},
		DPS:    []int{81, 90, 98, 108, 119, X, X, 175, 192, 211},
		DAM:    []int{90, 99, 108, 119, 131, X, X, 193, 212, 233},
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
	GOBLINS,
	SPEAR_GOBLINS,
	SKELETONS,
	MINIONS,
	BARBARIANS,
	MINION_HORDE,
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
		// Header 1
		fmt.Printf("%*s", -mutableAttrNameLen, attrTitle)
		// Any field will do, not just "cards".
		for level := range card[RARITY].(*Rarity).cards {
			fmt.Printf(" | %*s", -attrValueLen, fmt.Sprintf("LV%d", level+1))
		}
		fmt.Println()

		// Header 2
		fmt.Printf("%*s", mutableAttrNameLen, strings.Repeat("-", mutableAttrNameLen))
		// Any field will do, not just "cards".
		for range card[RARITY].(*Rarity).cards {
			fmt.Printf(" | %*s", attrValueLen, strings.Repeat("-", attrValueLen))
		}
		fmt.Println()

		// Content
		for _, attr := range CARD_ATTRIBUTES {
			if !attr.isMutable {
				continue
			}
			if values, ok := card[attr]; ok {
				fmt.Printf("%*s", -mutableAttrNameLen, attr)
				switch values.(type) {
				case []int:
					for _, value := range values.([]int) {
						fmt.Printf(" | %*s", attrValueLen, attr.format(value))
					}
				default:
					panic("Unknown values type")
				}
				fmt.Println()
			}
		}

		// Footer 1
		fmt.Printf("%*s", -mutableAttrNameLen, CARDS_REQ)
		for _, cardsReq := range card[RARITY].(*Rarity).cards {
			fmt.Printf(" | %*s", attrValueLen, formatInt(cardsReq))
		}
		fmt.Println()

		// Footer 2
		fmt.Printf("%*s", -mutableAttrNameLen, GOLD_REQ)
		for _, goldReq := range card[RARITY].(*Rarity).gold {
			fmt.Printf(" | %*s", attrValueLen, formatInt(goldReq))
		}
		fmt.Println()

		// Footer 3
		fmt.Printf("%*s", -mutableAttrNameLen, EXP_GAIN)
		for _, expGain := range card[RARITY].(*Rarity).exp {
			fmt.Printf(" | %*s", attrValueLen, formatInt(expGain))
		}
		fmt.Println()

		fmt.Println()
	}
}
