package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/mattn/go-tty"
	"github.com/sor4chi/go-terminal-game-sample/entity"
	"github.com/sor4chi/go-terminal-game-sample/field"
	"github.com/sor4chi/go-terminal-game-sample/object"
)

type Config struct {
	fieldWidth  int
	fieldHeight int
	objectCount int
}

func main() {
	config := Config{
		fieldWidth:  30,
		fieldHeight: 30,
		objectCount: 10,
	}
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer tty.Close()

	fmt.Println("Ready. Press any key ...")

	// init player
	player := entity.Entity{X: 0, Y: 0, Symbol: '@'}

	// set object
	objects := make([]object.Object, config.objectCount)
	for i := 0; i < config.objectCount; i++ {
		objects[i] = object.Object{X: rand.Intn(config.fieldWidth), Y: rand.Intn(config.fieldHeight), Symbol: '#'}
	}

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		player.MoveEntity(r)
		player.DisplayEntityPosition()
		field.DisplayField(player, objects, config.fieldWidth, config.fieldHeight)
	}
}
