package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/brandonc/advent2020/internal/mapping"
	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "12 [input file]",
		Short: "Runs the day 12 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day12)
		},
	})
}

func day12(file *os.File) error {
	scanner, err := tools.Readlines(file)
	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	var ship = &mapping.CartesianLoc{};

	for instr := range scanner {
		if len(instr) < 2 {
			return fmt.Errorf("invalid input: '%s'", instr)
		}

		action, value := instr[0:1], instr[1:]
		valueI, err := strconv.Atoi(value)

		if err != nil {
			return fmt.Errorf("invalid value: '%s'", value)
		}

		switch action {
		case "N":
			ship.MoveNorth(valueI)
		case "S":
			ship.MoveSouth(valueI)
		case "E":
			ship.MoveEast(valueI)
		case "W":
			ship.MoveWest(valueI)
		case "L":
			ship.Turn(mapping.L, valueI)
		case "R":
			ship.Turn(mapping.R, valueI)
		case "F":
			ship.Forward(valueI)
		}
	}

	fmt.Printf("manhattan distance from start = %d (part one)\n", ship.DistFromStart())

	shipWay := &mapping.CartesianLoc{}
	shipWay.MoveWaypointEast(10)
	shipWay.MoveWaypointNorth(1)

	file.Seek(0, 0)
	scanner, err = tools.Readlines(file)

	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	for instr := range scanner {
		action, value := instr[0:1], instr[1:]
		valueI, _ := strconv.Atoi(value)
	
		switch action {
		case "N":
			shipWay.MoveWaypointNorth(valueI)
		case "S":
			shipWay.MoveWaypointSouth(valueI)
		case "E":
			shipWay.MoveWaypointEast(valueI)
		case "W":
			shipWay.MoveWaypointWest(valueI)
		case "L":
			shipWay.RotateWaypoint(mapping.L, valueI)
		case "R":
			shipWay.RotateWaypoint(mapping.R, valueI)
		case "F":
			shipWay.ToWaypoint(valueI)
		}
	}

	fmt.Printf("manhattan distance from start = %d (part two)\n", shipWay.DistFromStart())

	return nil
}