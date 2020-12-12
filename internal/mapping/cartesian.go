package mapping

import "fmt"

var (
	// E is East
	E = 0
	// N is North
	N = 1
	// W is West
	W = 2
	// S is South
	S = 4
	// F is Forward
	F = 5
	// L is Left
	L = 6
	// R is Right
	R = 7
)

// CartesianLoc tracks an object with a direction through cartesian space
type CartesianLoc struct {
	dir int
	y int
	x int
	wy int
	wx int
}

// MoveNorth moves the object north while maintaining its direction
func (l *CartesianLoc) MoveNorth(amount int) {
	l.y -= amount
}

// MoveSouth moves the object South while maintaining its direction
func (l *CartesianLoc) MoveSouth(amount int) {
	l.y += amount
}

// MoveEast moves the object East while maintaining its direction
func (l *CartesianLoc) MoveEast(amount int) {
	l.x += amount
}

// MoveWest moves the object West while maintaining its direction
func (l *CartesianLoc) MoveWest(amount int) {
	l.wx -= amount
}

// MoveWaypointNorth moves the waypoint north while maintaining its direction
func (l *CartesianLoc) MoveWaypointNorth(amount int) {
	l.wy += amount
}

// MoveWaypointSouth moves the waypoint South while maintaining its direction
func (l *CartesianLoc) MoveWaypointSouth(amount int) {
	l.wy -= amount
}

// MoveWaypointEast moves the waypoint East while maintaining its direction
func (l *CartesianLoc) MoveWaypointEast(amount int) {
	l.wx += amount
}

// MoveWaypointWest moves the waypoint West while maintaining its direction
func (l *CartesianLoc) MoveWaypointWest(amount int) {
	l.wx -= amount
}

// Turn turns the object left or right by the specified number of degrees
func (l *CartesianLoc) Turn(lOrR int, amount int) error {
	if amount < 0 && amount % 90 != 0 {
		return fmt.Errorf("%d is not a valid amount", amount)
	}

	for v := amount; v != 0; v -= 90 {
		switch l.dir {
		case N:
			if lOrR == L {
				l.dir = W
			} else {
				l.dir = E
			}
		case S:
			if lOrR == L {
				l.dir = E
			} else {
				l.dir = W
			}
		case E:
			if lOrR == L {
				l.dir = N
			} else {
				l.dir = S
			}
		case W:
			if lOrR == L {
				l.dir = S
			} else {
				l.dir = N
			}
		}
	}

	return nil
}

// RotateWaypoint rotates the waypoint around the ship
func (l *CartesianLoc) RotateWaypoint(lOrR int, amount int) error {
	if amount < 0 && amount % 90 != 0 {
		return fmt.Errorf("%d is not a valid amount", amount)
	}

	for v := amount; v != 0; v -= 90 {
		switch lOrR {
		case R:
			l.wx, l.wy = l.wy, -l.wx
		case L:
			l.wx, l.wy = -l.wy, l.wx
		}
	}

	return nil
}

// Forward moves the object forward in the amount specified
func (l *CartesianLoc) Forward(amount int) {
	switch l.dir {
	case N:
		l.y -= amount
	case S:
		l.y += amount
	case E:
		l.x += amount
	case W:
		l.x -= amount
	}
}

// ToWaypoint moves the ship to the waypoint the specified number of times
func (l *CartesianLoc) ToWaypoint(times int) {
	l.y += l.wy * times
	l.x += l.wx * times
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// DistFromStart is the manhattan distance from the origin
func (l *CartesianLoc) DistFromStart() int {
	return abs(l.y) + abs(l.x)
}