package main

import "fmt"

func main() {
	//初始化
	game := new(Game)
	manager := NewGameManager()
	game.Process = "玩了30%的进度"
	//保存进度信息
	save1 := game.CreateGameSaver()
	manager.AddGameSaver("进度1", save1)
	game.Process = "玩了40%的进度"
	fmt.Println(game.Process)
	//恢复之前的存档
	game.RestoreGame(manager.GetGameSaverByKey("进度1"))
	fmt.Println(game.Process)
}

//游戏，发起人
type Game struct {
	Process string
}

func (game *Game) RestoreGame(gameSaver *GameSaver) {
	game.Process = gameSaver.Process
}

func (game *Game) CreateGameSaver() *GameSaver {
	gameSaver := new(GameSaver)
	gameSaver.Process = game.Process
	return gameSaver
}

//游戏存档-备忘录
type GameSaver struct {
	Process string
}

//游戏存档管理者
type GameManager struct {
	gameSavers map[string]*GameSaver
}

func (m *GameManager) AddGameSaver(key string, saver *GameSaver) {
	m.gameSavers[key] = saver
}

func (m *GameManager) GetGameSaverByKey(key string) *GameSaver {
	return m.gameSavers[key]
}

func NewGameManager() *GameManager {
	manager := new(GameManager)
	manager.gameSavers = make(map[string]*GameSaver)
	return manager
}
