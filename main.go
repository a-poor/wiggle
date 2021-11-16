package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/go-vgo/robotgo"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "wiggle",
		Usage: "Move the mouse around to keep your screen from sleeping!",
		Action: func(c *cli.Context) error {
			// Start running the spinner
			s := spinner.New(
				spinner.CharSets[11],
				100*time.Millisecond,
			)
			s.Color("green")
			s.Suffix = " Wiggling..."
			s.FinalMSG = "Done Wiggling!"
			s.Start()
			s.Reverse() // Spin clockwise
			defer s.Stop()

			// Get the screen size
			width, height := robotgo.GetScreenSize()

			// Create a RNG
			rs := rand.NewSource(time.Now().UnixNano())
			rng := rand.New(rs)

			// Loop forever
			for {
				// Pick a random x and y
				x := rng.Intn(width)
				y := rng.Intn(height)

				// Move the mouse to a random location
				robotgo.MoveSmooth(x, y)

				// Sleep
				time.Sleep(time.Second * 5)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done.")
}
