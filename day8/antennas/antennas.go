package antennas

import vec "adventofcode/geometry/vec2d"

type AntennaGroup struct {
	identity  string
	positions []vec.Vec2d
}

func InitGroup(id string) AntennaGroup {
	return AntennaGroup{identity: id, positions: []vec.Vec2d{}}
}

func (a AntennaGroup) GetId() string {
	return a.identity
}

func (a *AntennaGroup) AddAntenna(v vec.Vec2d) {
	a.positions = append(a.positions, v)
}

func (a AntennaGroup) AntiNodes() []vec.Vec2d {
	antiNodes := []vec.Vec2d{}
	for i := 0; i < len(a.positions); i++ {
		for j := i + 1; j < len(a.positions); j++ {
			antiNodeFromPair := antinodePair(a.positions[i], a.positions[j])
			antiNodes = append(antiNodes, antiNodeFromPair...)
		}
	}
	return antiNodes
}

func (a AntennaGroup) AntiNodesHarmonics(limitCorner vec.Vec2d) []vec.Vec2d {
	antiNodes := []vec.Vec2d{}
	for i := 0; i < len(a.positions); i++ {
		for j := i + 1; j < len(a.positions); j++ {
			antiNodeFromPair := antinodePairsHarmonics(a.positions[i], a.positions[j], limitCorner)
			antiNodes = append(antiNodes, antiNodeFromPair...)
		}
	}
	return antiNodes
}

func antinodePair(antenna1 vec.Vec2d, antenna2 vec.Vec2d) []vec.Vec2d {
	pairOfAntiNodes := []vec.Vec2d{}
	diff := vec.Subtract(antenna2, antenna1)
	antinode2 := vec.Add(antenna2, diff)
	diffFlipped := vec.ScalarMult(-1, diff)
	antinode1 := vec.Add(antenna1, diffFlipped)
	pairOfAntiNodes = append(pairOfAntiNodes, antinode1, antinode2)
	return pairOfAntiNodes
}

func antinodePairsHarmonics(antenna1 vec.Vec2d, antenna2 vec.Vec2d, limitCorner vec.Vec2d) []vec.Vec2d {
	antiNodes := []vec.Vec2d{}
	directionVec := vec.Subtract(antenna2, antenna1)
	forewardCandidate := antenna2
	for inbounds(forewardCandidate, limitCorner) {
		antiNodes = append(antiNodes, forewardCandidate)
		forewardCandidate = vec.Add(forewardCandidate, directionVec)
	}
	backwards := vec.ScalarMult(-1, directionVec)
	backwardsCandidate := antenna1
	for inbounds(backwardsCandidate, limitCorner) {
		antiNodes = append(antiNodes, backwardsCandidate)
		backwardsCandidate = vec.Add(backwardsCandidate, backwards)
	}

	return antiNodes
}

func inbounds(v vec.Vec2d, limitCorner vec.Vec2d) bool {
	xIn := 0 <= v.GetX() && v.GetX() < limitCorner.GetX()
	yIn := 0 <= v.GetY() && v.GetY() < limitCorner.GetY()
	return xIn && yIn
}
