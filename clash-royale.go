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
		exp:   []int{0, 4, 5, 6, 10, 25, 50, 100, X, X, X, X},
	}
	RARE = &Rarity{
		name:  "Rare",
		cards: []int{0, 2, 4, 10, 20, 50, 100, 200, X, X},
		gold:  []int{0, 50, 150, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 6, 10, 25, 50, 100, X, X, X, X},
	}
	EPIC = &Rarity{
		name:  "Epic",
		cards: []int{0, 2, 4, 10, 20, 50, 100, 200},
		gold:  []int{0, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 25, 50, 100, X, X, X, X},
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

type CardAttribute interface {
	CardAttribute() // Tag
	String() string
}

type FixedCardAttribute struct {
	name        string
	formatValue func(value interface{}) string
}

func (attr *FixedCardAttribute) CardAttribute() {
}

func (attr *FixedCardAttribute) String() string {
	return attr.name
}

type UpgradableCardAttribute struct {
	name         string
	formatValues func(values interface{}) []string
}

func (attr *UpgradableCardAttribute) CardAttribute() {
}

func (attr *UpgradableCardAttribute) String() string {
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

func formatInts(values interface{}) []string {
	ints := values.([]int)
	strings := make([]string, len(ints))
	for i, v := range ints {
		strings[i] = formatInt(v)
	}
	return strings
}

func formatTimes(values interface{}) []string {
	ints := values.([]int)
	strings := make([]string, len(ints))
	for i, v := range ints {
		strings[i] = formatTime(v)
	}
	return strings
}

var (
	NAME = &FixedCardAttribute{
		"Name",
		formatString,
	}
	ARENA = &FixedCardAttribute{
		"Arena",
		func(value interface{}) string {
			return value.(*Arena).String()
		},
	}
	RARITY = &FixedCardAttribute{
		"Rarity",
		func(value interface{}) string {
			return value.(*Rarity).String()
		},
	}
	TYPE = &FixedCardAttribute{
		"Type",
		func(value interface{}) string {
			return string(value.(Type))
		},
	}
	DESC = &FixedCardAttribute{
		"Description",
		formatString,
	}
	COST = &FixedCardAttribute{
		"Elixir Cost",
		formatInt,
	}
	HP = &UpgradableCardAttribute{
		"Hitpoints",
		formatInts,
	}
	DPS = &UpgradableCardAttribute{
		"Damage per Second",
		formatInts,
	}
	DAM = &UpgradableCardAttribute{
		"Damage",
		formatInts,
	}
	ADAM = &UpgradableCardAttribute{
		"Area Damage",
		formatInts,
	}
	DDAM = &UpgradableCardAttribute{
		"Death Damage",
		formatInts,
	}
	SKE_LV = &UpgradableCardAttribute{
		"Skeleton Level",
		formatInts,
	}
	SGO_LV = &UpgradableCardAttribute{
		"Spear Goblin Level",
		formatInts,
	}
	SSPD = &FixedCardAttribute{
		"Spawn Speed",
		formatTime,
	}
	HSPD = &FixedCardAttribute{
		"Hit Speed",
		formatTime,
	}
	TGTS = &FixedCardAttribute{
		"Targets",
		func(value interface{}) string {
			return string(value.(Targets))
		},
	}
	SPD = &FixedCardAttribute{
		"Speed",
		func(value interface{}) string {
			return string(value.(Speed))
		},
	}
	RNG = &FixedCardAttribute{
		"Range",
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
	DTIME = &FixedCardAttribute{
		"Deploy Time",
		formatTime,
	}
	COUNT = &FixedCardAttribute{
		"Count",
		func(value interface{}) string {
			return fmt.Sprintf("x %d", value.(int))
		},
	}
)

var CARD_ATTRIBUTES = [...]CardAttribute{
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
	SSPD,
	HSPD,
	TGTS,
	SPD,
	RNG,
	DTIME,
	COUNT,
}

type Card map[CardAttribute]interface{}

var (
	KNIGHT = Card{
		NAME:   "Knight",
		ARENA:  ARENA_0,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   "A tough melee fighter. The Barbarian's handsome, cultured cousin. Rumor has it that he was knighted based on the sheer awesomeness of his mustache alone.",
		COST:   3,
		HP:     []int{600, 660, 726, 798, 876, 960, 1056, 1158, 1272, 1398, 1536, 1686},
		DPS:    []int{68, 74, 81, 90, 99, 109, 120, 130, 144, 158, 174, 190},
		DAM:    []int{75, 82, 90, 99, 109, 120, 132, 144, 159, 174, 192, 210},
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
		HP:     []int{125, 137, 151, 166, 182, 200, 220, 241, 265, 291, 320, 351},
		DPS:    []int{33, 36, 40, 44, 48, 53, 58, 64, 70, 77, 85, 93},
		DAM:    []int{40, 44, 48, 53, 58, 64, 70, 77, 84, 93, 102, 112},
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
		DESC:   "Small, lightly protected skeleton that throws bombs. Deals area damage that can wipe out a swarm of enemies.",
		COST:   3,
		HP:     []int{150, 165, 181, 199, 219, 240, 264, 289, 318, 349, 384, 421},
		DPS:    []int{52, 57, 63, 70, 76, 84, 92, 101, 111, 122, 134, 147},
		ADAM:   []int{100, 110, 121, 133, 146, 160, 176, 193, 212, 233, 256, 281},
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
		HP:     []int{80, 88, 96, 106, 116, 128, 140, X, 169, 186, 204, 224},
		DPS:    []int{45, 50, 54, 60, 66, 72, 80, X, 96, 105, 116, 127},
		DAM:    []int{50, 55, 60, 66, 73, 80, 88, X, 106, 116, 128, 140},
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
		HP:     []int{52, 57, 62, 69, 75, 83, 91, 100, 110, 121, 133, 146},
		DPS:    []int{18, 20, 22, 23, 26, 29, 32, 35, 38, 42, 46, 51},
		DAM:    []int{24, 26, 29, 31, 35, 38, 42, 46, 50, 55, 61, 67},
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
		HP:     []int{30, 33, 36, 39, 43, 48, 52, 57, 63, 69, 76, 84},
		DPS:    []int{30, 33, 36, 39, 43, 48, 52, 57, 63, 69, 76, 84},
		DAM:    []int{30, 33, 36, 39, 43, 48, 52, 57, 63, 69, 76, 84},
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
		HP:     []int{90, 99, 108, 119, 131, 144, 158, 173, 190, 209, 230, 252},
		DPS:    []int{40, 44, 48, 53, 58, 64, 70, 77, 84, 93, 102, 112},
		DAM:    []int{40, 44, 48, 53, 58, 64, 70, 77, 84, 93, 102, 112},
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
		HP:     []int{300, X, X, 399, 438, 480, 528, 579, X, 699, 768, 843},
		DPS:    []int{50, X, X, 66, 72, 80, 88, 96, X, 116, 128, 140},
		DAM:    []int{75, X, X, 99, 109, 120, 132, 144, X, 174, 192, 210},
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
		HP:     []int{90, 99, 108, 119, 131, 144, 158, 173, 190, 209, 230, 252},
		DPS:    []int{40, 44, 48, 53, 58, 64, 70, 77, 84, 93, 102, 112},
		DAM:    []int{40, 44, 48, 53, 58, 64, 70, 77, 84, 93, 102, 112},
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
		HP:     []int{2000, 2200, 2420, 2660, 2920, X, 3520, 3860, X, 4660},
		DPS:    []int{80, 88, 96, 106, 116, X, 140, 154, X, 186},
		DAM:    []int{120, 132, 145, 159, 175, X, 221, 231, X, 279},
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
		DESC:   "Don't be fooled by her delicately coiffed hair, the Musketeer is a mean shot with her trusty boomstick.",
		COST:   4,
		HP:     []int{340, 374, 411, 452, 496, X, X, 656, 720, 792},
		DPS:    []int{90, X, 110, 120, 132, X, X, 175, 192, 211},
		DAM:    []int{100, 110, 121, 133, 146, 160, 176, 193, 212, 233},
		HSPD:   1.1,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RNG:    6.5,
		DTIME:  1,
	}
	MINI_PEKKA = Card{
		NAME:   "Mini P.E.K.K.A",
		ARENA:  ARENA_0,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   "The arena is a certified butterfly-free zone. No distractions for P.E.K.K.A, only destruction.",
		COST:   4,
		HP:     []int{600, 660, 726, 798, 876, X, X, X, X, X},
		DPS:    []int{180, 198, 218, 240, 263, X, 317, 348, X, X},
		DAM:    []int{325, 357, 393, 432, 474, X, 572, 627, X, X},
		HSPD:   1.8,
		TGTS:   GROUND,
		SPD:    FAST,
		RNG:    MELEE,
		DTIME:  1,
	}
	VALKYRIE = Card{
		NAME:   "Valkyrie",
		ARENA:  ARENA_1,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   "Tough melee fighter, deals area damage around her. Swarm or horde, no problem! She can take them all out with a few spins.",
		COST:   4,
		HP:     []int{800, 880, 968, 1064, 1168, 1280, 1408, 1544, 1696, 1864},
		DPS:    []int{73, X, X, X, X, X, X, 141, X, X},
		DAM:    []int{110, 121, 133, 146, 160, 176, 193, 212, 233, 256},
		HSPD:   1.5,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    MELEE,
		DTIME:  1,
	}
	HOG_RIDER = Card{
		NAME:   "Hog Rider",
		ARENA:  ARENA_4,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   `Fast melee troop that targets buildings and can jump over the river. He followed the echoing call of "Hog Riderrrrr" all the way through the arena doors.`,
		COST:   4,
		HP:     []int{800, 880, 968, 1064, 1168, 1280, 1408, 1544, 1696, 1864},
		DPS:    []int{106, 116, 128, 141, 155, X, X, 205, 226, 248},
		DAM:    []int{160, 176, 193, 212, 233, X, X, 308, 339, 372},
		HSPD:   1.5,
		TGTS:   BUILDINGS,
		SPD:    VERY_FAST,
		RNG:    MELEE,
		DTIME:  1,
	}
	WIZARD = Card{
		NAME:   "Wizard",
		ARENA:  ARENA_5,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   "The most awesome man to ever set foot in the arena, the Wizard will blow you away with his handsomeness... and/or fireballs.",
		COST:   5,
		HP:     []int{340, X, 411, 452, 496, X, X, X, X, X},
		DPS:    []int{76, X, 92, 101, 111, X, X, X, X, X},
		ADAM:   []int{130, X, 157, 172, 189, X, X, X, X, X},
		HSPD:   1.7,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RNG:    5.5,
		DTIME:  1,
	}
	WITCH = Card{
		NAME:   "Witch",
		ARENA:  ARENA_0,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   "Summons skeletons, shoots destructo beams, has glowing pink eyes that unfortunately don't shoot lasers.",
		COST:   5,
		HP:     []int{500, 550, 605, X, 730, 800, 880, 968},
		DPS:    []int{51, 55, 61, X, 74, 81, 90, X},
		ADAM:   []int{36, 39, 43, 48, 52, 57, 63, 69},
		SKE_LV: []int{6, 7, 8, 9, 10, 11, 12, X},
		SSPD:   7.5,
		HSPD:   0.7,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RNG:    5.5,
		DTIME:  1,
	}
	SKELETON_ARMY = Card{
		NAME:   "Skeleton Army",
		ARENA:  ARENA_0,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   "Summons an army of skeletons. Meet Larry and his friends Harry, Gerry, Terry, Mary, etc.",
		COST:   4,
		HP:     []int{30, 33, 36, 39, 43, 48, 52, 57},
		DPS:    []int{30, 33, 36, 39, 43, 48, 52, 57},
		DAM:    []int{30, 33, 36, 39, 43, 48, 52, 57},
		HSPD:   1,
		TGTS:   GROUND,
		SPD:    FAST,
		RNG:    MELEE,
		DTIME:  1,
		COUNT:  20,
	}
	BABY_DRAGON = Card{
		NAME:   "Baby Dragon",
		ARENA:  ARENA_0,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   "Flying troop that deals area damage. Baby dragons hatch cute, hungry, and ready for a barbeque.",
		COST:   4,
		HP:     []int{800, 880, 968, 1064, 1168, 1280, 1408, 1544},
		DPS:    []int{55, 61, 67, 73, 81, 88, 97, 107},
		ADAM:   []int{100, 110, 121, 133, 146, 160, 176, 193},
		HSPD:   1.8,
		TGTS:   AIR_AND_GROUND,
		SPD:    FAST,
		RNG:    3.5,
		DTIME:  1,
	}
	PRINCE = Card{
		NAME:   "Prince",
		ARENA:  ARENA_0,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   "Don't let the little pony fool you. Once the Prince gets a running start, you WILL be trampled. Does 2x damage once he gets charging.",
		COST:   5,
		HP:     []int{1100, 1210, 1331, 1463, 1606, 1760, 1936, X},
		DPS:    []int{146, 161, 177, 194, 214, 234, 258, X},
		DAM:    []int{220, 242, 266, 292, 321, 352, 387, X},
		HSPD:   1.5,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    2.5,
		DTIME:  1,
	}
	GIANT_SKELETON = Card{}
	BALLOON        = Card{}
	PEKKA          = Card{}
	GOLEM          = Card{}
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
	MINI_PEKKA,
	VALKYRIE,
	HOG_RIDER,
	WIZARD,
	WITCH,
	SKELETON_ARMY,
	BABY_DRAGON,
	PRINCE,
}

// --- Table ---

type Table struct {
	headers                        map[int]string
	contents                       map[int]map[int]string
	maxLens                        map[int]int
	minRow, maxRow, minCol, maxCol int
}

func NewTable() *Table {
	return &Table{
		make(map[int]string),
		make(map[int]map[int]string),
		make(map[int]int),
		0, 0, 0, 0,
	}
}

func (t *Table) SetHeader(col int, header string) {
	t.headers[col] = header
	if length := len(header); length > t.maxLens[col] {
		t.maxLens[col] = length
	}
}

func (t *Table) GetHeader(col int) string {
	header, ok := t.headers[col]
	if !ok {
		return ""
	}
	return header
}

func (t *Table) SetContent(row, col int, content string) {
	t2, ok := t.contents[row]
	if !ok {
		t2 = make(map[int]string)
		t.contents[row] = t2
	}
	t2[col] = content
	if length := len(content); length > t.maxLens[col] {
		t.maxLens[col] = length
	}
	if row < t.minRow {
		t.minRow = row
	}
	if row > t.maxRow {
		t.maxRow = row
	}
	if col < t.minCol {
		t.minCol = col
	}
	if col > t.maxCol {
		t.maxCol = col
	}
}

func (t *Table) GetContent(row, col int) string {
	t2, ok := t.contents[row]
	if !ok {
		return ""
	}
	content, ok := t2[col]
	if !ok {
		return ""
	}
	return content
}

func (t *Table) GetMaxLen(col int) int {
	maxLen, ok := t.maxLens[col]
	if !ok {
		return 0
	}
	return maxLen
}

func (t *Table) Print() {
	for col, sep := t.minCol, ""; col <= t.maxCol; col++ {
		fmt.Printf("%s%*s", sep, -t.GetMaxLen(col), t.GetHeader(col))
		sep = " | "
	}
	fmt.Println()
	for col, sep := t.minCol, ""; col <= t.maxCol; col++ {
		fmt.Printf("%s%s", sep, strings.Repeat("-", t.GetMaxLen(col)))
		sep = " | "
	}
	fmt.Println()
	for row := t.minRow; row <= t.maxRow; row++ {
		for col, sep := t.minCol, ""; col <= t.maxCol; col++ {
			fmt.Printf("%s%*s", sep, -t.GetMaxLen(col), t.GetContent(row, col))
			sep = " | "
		}
		fmt.Println()
	}
	fmt.Println()
}

/*
func main() {
	t := NewTable()
	t.SetHeader(0, "Attribute")
	t.SetHeader(1, "Value")
	t.SetContent(0, 0, "Name")
	t.SetContent(0, 1, "Knight")
	t.Print()
}
*/

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
	upgradableAttrNameLen := attrTitleLen
	for _, attr := range CARD_ATTRIBUTES {
		attrNameLen := len(attr.String())
		switch attr.(type) {
		case *FixedCardAttribute:
			if attrNameLen > fixedAttrNameLen {
				fixedAttrNameLen = attrNameLen
			}
		case *UpgradableCardAttribute:
			if attrNameLen > upgradableAttrNameLen {
				upgradableAttrNameLen = attrNameLen
			}
		}
	}
	for _, attr := range RARITY_ATTRIBUTES {
		attrNameLen := len(attr)
		if attrNameLen > upgradableAttrNameLen {
			upgradableAttrNameLen = attrNameLen
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
			fattr, ok := attr.(*FixedCardAttribute)
			if !ok {
				continue
			}
			if value, ok := card[fattr]; ok {
				fmt.Printf("%*s | %s\n", -fixedAttrNameLen, fattr, fattr.formatValue(value))
			}
		}
		fmt.Println()

		// Upgradable Attribute Table
		// Header 1
		fmt.Printf("%*s", -upgradableAttrNameLen, attrTitle)
		// Any field will do, not just "cards".
		for level := range card[RARITY].(*Rarity).cards {
			fmt.Printf(" | %*s", -attrValueLen, fmt.Sprintf("LV%d", level+1))
		}
		fmt.Println()

		// Header 2
		fmt.Printf("%*s", upgradableAttrNameLen, strings.Repeat("-", upgradableAttrNameLen))
		// Any field will do, not just "cards".
		for range card[RARITY].(*Rarity).cards {
			fmt.Printf(" | %*s", attrValueLen, strings.Repeat("-", attrValueLen))
		}
		fmt.Println()

		// Content
		for _, attr := range CARD_ATTRIBUTES {
			uattr, ok := attr.(*UpgradableCardAttribute)
			if !ok {
				continue
			}
			if values, ok := card[attr]; ok {
				fmt.Printf("%*s", -upgradableAttrNameLen, uattr)
				for _, fvalue := range uattr.formatValues(values) {
					fmt.Printf(" | %*s", attrValueLen, fvalue)
				}
				fmt.Println()
			}
		}

		// Footer 1
		fmt.Printf("%*s", -upgradableAttrNameLen, CARDS_REQ)
		for _, cardsReq := range card[RARITY].(*Rarity).cards {
			fmt.Printf(" | %*s", attrValueLen, formatInt(cardsReq))
		}
		fmt.Println()

		// Footer 2
		fmt.Printf("%*s", -upgradableAttrNameLen, GOLD_REQ)
		for _, goldReq := range card[RARITY].(*Rarity).gold {
			fmt.Printf(" | %*s", attrValueLen, formatInt(goldReq))
		}
		fmt.Println()

		// Footer 3
		fmt.Printf("%*s", -upgradableAttrNameLen, EXP_GAIN)
		for _, expGain := range card[RARITY].(*Rarity).exp {
			fmt.Printf(" | %*s", attrValueLen, formatInt(expGain))
		}
		fmt.Println()

		fmt.Println()
	}
}
