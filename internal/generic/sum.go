package generic

type number interface {
	int | uint
}

func Sum[T number](ts []T) T {
	s := T(0)
	for _, t := range ts {
		s += t
	}

	return s
}
