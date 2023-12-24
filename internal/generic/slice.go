package generic

func Map[T, E any](s []T, f func(int, T) E) []E {
	newS := make([]E, len(s))

	for idx, sv := range s {
		newS[idx] = f(idx, sv)
	}

	return newS
}
