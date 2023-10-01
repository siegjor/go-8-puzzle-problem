package heuristics

type Strategy string

const (
	UNIFORM_COST           Strategy = "UNIFORM_COST"
	A_STAR_MANHATTAN       Strategy = "A_STAR_MANHATTAN"
	A_STAR_LINEAR_CONFLICT Strategy = "A_STAR_LINEAR_CONFLICT"
)
