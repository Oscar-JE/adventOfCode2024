package particle

import vec "adventofcode/geometry/vec2d"

type Particle struct {
	initialPosition vec.Vec2d
	velocity        vec.Vec2d
}

func Init(position vec.Vec2d, velocity vec.Vec2d) Particle {
	return Particle{initialPosition: position, velocity: velocity}
}

func (p Particle) ProjectedPosition(nrOfTimeUnits int) vec.Vec2d {
	return vec.Add(p.initialPosition, vec.ScalarMult(nrOfTimeUnits, p.velocity))
}
