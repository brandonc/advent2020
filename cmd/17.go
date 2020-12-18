package cmd

import (
	"fmt"
	"os"

	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "17 [input file]",
		Short: "Runs the day 17 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day17)
		},
	})
}

type point3d struct {
	x int
	y int
	z int
}

type cube3d struct {
	location point3d
	active bool
}

func (p *point3d) neighbors3d() <-chan point3d {
	result := make(chan point3d)

	go func() {
		for x := p.x - 1; x <= p.x + 1; x++ {
			for y := p.y - 1; y <= p.y + 1; y++ {
				for z := p.z - 1; z <= p.z + 1; z++ {
					if p.x == x && p.y == y && p.z == z {
						// p (self)
						continue
					}
					result <- point3d{
						x: x, y: y, z: z,
					}
				}
			}
		}
		close(result)
	}()
	return result
}

func countActive3d(space map[point3d]*cube3d) int {
	var count = 0
	for _, cube := range space {
		if cube.active {
			count++
		}
	}

	return count
}

func (p *point3d) countActiveNeighbors3d(space map[point3d]*cube3d) int {
	var count = 0
	for neighbor := range p.neighbors3d() {
		if n, exists := space[neighbor]; exists {
			if n.active {
				count++
			}
		}
	}
	return count
}

func simulate3d(space map[point3d]*cube3d, cycles int) {
	for i := 0; i < cycles; i++ {
		var remap = make(map[point3d]bool)
		for p, cube := range space {
			for neighborLocation := range p.neighbors3d() {
				if _, exists := space[neighborLocation]; !exists {
					if neighborLocation.countActiveNeighbors3d(space) == 3 {
						remap[neighborLocation] = true
					}
				}
			}

			activeNeighbors := p.countActiveNeighbors3d(space)
			if cube.active && activeNeighbors != 2 && activeNeighbors != 3 {
				remap[cube.location] = false
				continue
			}
			if !cube.active && activeNeighbors == 3 {
				remap[cube.location] = true
			}
		}

		for p, active := range remap {
			if c, exists := space[p]; exists {
				c.active = active
			} else {
				space[p] = &cube3d{
					location: p,
					active: active,
				}
			}
		}
	}
}

type point4d struct {
	x int
	y int
	z int
	w int
}

type cube4d struct {
	location point4d
	active bool
}

func (p *point4d) neighbors4d() <-chan point4d {
	result := make(chan point4d)

	go func() {
		for x := p.x - 1; x <= p.x + 1; x++ {
			for y := p.y - 1; y <= p.y + 1; y++ {
				for z := p.z - 1; z <= p.z + 1; z++ {
					for w := p.w - 1; w <= p.w + 1; w++ {
						if p.x == x && p.y == y && p.z == z && p.w == w {
							// p (self)
							continue
						}
						result <- point4d{
							x: x, y: y, z: z, w: w,
						}
					}
				}
			}
		}
		close(result)
	}()
	return result
}

func countActive4d(space map[point4d]*cube4d) int {
	var count = 0
	for _, cube := range space {
		if cube.active {
			count++
		}
	}

	return count
}

func (p *point4d) countActiveNeighbors4d(space map[point4d]*cube4d) int {
	var count = 0
	for neighbor := range p.neighbors4d() {
		if n, exists := space[neighbor]; exists {
			if n.active {
				count++
			}
		}
	}
	return count
}

func simulate4d(space map[point4d]*cube4d, cycles int) {
	for i := 0; i < cycles; i++ {
		var remap = make(map[point4d]bool)
		for p, cube := range space {
			for neighborLocation := range p.neighbors4d() {
				if _, exists := space[neighborLocation]; !exists {
					if neighborLocation.countActiveNeighbors4d(space) == 3 {
						remap[neighborLocation] = true
					}
				}
			}

			activeNeighbors := p.countActiveNeighbors4d(space)
			if cube.active && activeNeighbors != 2 && activeNeighbors != 3 {
				remap[cube.location] = false
				continue
			}
			if !cube.active && activeNeighbors == 3 {
				remap[cube.location] = true
			}
		}

		for p, active := range remap {
			if c, exists := space[p]; exists {
				c.active = active
			} else {
				space[p] = &cube4d{
					location: p,
					active: active,
				}
			}
		}
	}
}

func day17(file *os.File) error {
	scanner, err := tools.Readlines(file)

	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	var space3d = make(map[point3d]*cube3d)
	var space4d = make(map[point4d]*cube4d)

	var y = 0
	for line := range scanner {
		var x = 0
		for _, r := range line {
			var active = false
			if r == '#' {
				active = true
			}
			key3d := point3d{ x, y, 1 }
			key4d := point4d{ x, y, 1, 1 }
			space3d[key3d] = &cube3d{location: key3d, active: active }
			space4d[key4d] = &cube4d{location: key4d, active: active }
			
			x++
		}
		y++
	}

	simulate3d(space3d, 6)
	simulate4d(space4d, 6)

	fmt.Printf("after 6 cycles the number of active cubes in 3d space is %d (part one)\n", countActive3d(space3d))
	fmt.Printf("after 6 cycles the number of active cubes in 4d space is %d (part two)\n", countActive4d(space4d))
	
	return nil
}