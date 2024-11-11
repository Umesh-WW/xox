package types

type Iboard [3][3]string

type GameState string

// Define constants for the possible game states
const (
	GameOver  GameState = "over"
	XWin      GameState = "x-win"
	YWin      GameState = "y-win"
	GameError GameState = "error"
	Play      GameState = "play"
)

type Game struct {
	Id    int
	Board Iboard
	State GameState
}
