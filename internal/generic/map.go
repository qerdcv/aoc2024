package generic

func Map[T, V any](ts []T, f func(t T) V) []V {
	vs := make([]V, len(ts))
	for i, t := range ts {
		vs[i] = f(t)
	}

	return vs
}
