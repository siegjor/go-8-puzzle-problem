package utils

import "n-puzzle-problem/nodes"

type Comparator interface {
	Equals(v *nodes.Node) bool
}
