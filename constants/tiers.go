package constants

const (
	Free   int = 0
	Bronze int = 1
	Silver int = 2
	Gold   int = 3
)

var LevelToTier = map[int]string{
	0: "free",
	1: "bronze",
	2: "silver",
	3: "gold",
	4: "silver",
	5: "gold",
	6: "gold",
}
