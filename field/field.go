package field

import (
	"fmt"

	"github.com/sor4chi/go-terminal-game-sample/entity"
	"github.com/sor4chi/go-terminal-game-sample/object"
)

func DisplayField(player entity.Entity, objects []object.Object, fieldWidth int, fieldHeight int) {

	if player.X < 0 || player.X >= fieldWidth || player.Y < 0 || player.Y >= fieldHeight {
		fmt.Println("You are out of field")
	}
	var field [256][256]rune
	var result string
	// if player exists, display player.Symbol at player.X, player.Y
	// if for each objects exist, display object.Symbol at object.X, object.Y
	// if player.X or player.Y is out of field, display "You are out of field"
	field[player.Y][player.X] = player.Symbol
	for _, obj := range objects {
		field[obj.Y][obj.X] = obj.Symbol
	}
	for i := 0; i < fieldHeight; i++ {
		for j := 0; j < fieldWidth; j++ {
			if i == 0 || i == fieldHeight-1 || j == 0 || j == fieldWidth-1 {
				result += "*"
			} else if field[i][j] == 0 {
				result += " "
			} else {
				result += string(field[i][j])
			}
		}
		result += "\n"
	}
	fmt.Println(result)
}
