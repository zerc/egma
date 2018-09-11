package renderers

import (
	"fmt"
	"game/models"
	"strings"
	"time"
)

// Renders everything to the terminal's stdout.
type Console struct {
}

// Draws a text provided.
func (c Console) DrawText(s string) error {
	fmt.Println("\033[H\033[2J")
	fmt.Println(s)
	return nil
}

// Starts rendering the scene in the infinite loop.
func (c Console) Start(s *models.Scene) {
	for {
		c.Draw(s)
		time.Sleep(10 * time.Millisecond)
	}
}

// Draws the whole 2D scene.
func (c Console) Draw(s *models.Scene) error {
	var tmp []string
	tmp = append(tmp, fmt.Sprintf("┏%s┓", strings.Repeat("━", s.Width)))

	for _, row := range s.Matrix {
		rowTmp := make([]string, s.Width+2)
		rowTmp = append(rowTmp, "┃")

		for _, obj := range row {
			rowTmp = append(rowTmp, Render(obj))
		}

		rowTmp = append(rowTmp, "┃")
		tmp = append(tmp, strings.Join(rowTmp, ""))
	}

	tmp = append(tmp, fmt.Sprintf("┗%s┛", strings.Repeat("━", s.Width)))
	rendered := strings.Join(tmp, "\n")

	fmt.Println("\033[H\033[2J") // clear the console
	fmt.Println(rendered)        // draw the scene

	return nil
}

// Renders the game object given.
func Render(o *models.GameObject) string {
	if o == nil {
		return " "
	} else if player, ok := (*o).(*models.Player); ok {
		if player.Colour == "red" {
			return "\033[1m\033[31m░\033[0m\033[21m"
		} else if player.Colour == "blue" {
			return "\033[1m\033[34m░\033[0m\033[21m"
		} else if player.Colour == "green" {
			return "\033[1m\033[32m░\033[0m\033[21m"
		} else if player.Colour == "orange" {
			return "\033[1m\033[36m░\033[0m\033[21m"
		} else {
			return "░"
		}
	} else {
		return " "
	}
}
