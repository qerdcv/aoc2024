package generic

func Map[T, V any](ts []T, f func(t T) V) []V {
	vs := make([]V, len(ts))
	for i, t := range ts {
		vs[i] = f(t)
	}

	return vs
}

func Filter[T any](ts []T, f func(t T) bool) []T {
	newTS := make([]T, 0, len(ts))
	for _, t := range ts {
		if f(t) {
			newTS = append(newTS, t)
		}
	}

	return newTS
}

func ForEach[T any](ts []T, f func(t T)) {
	for _, t := range ts {
		f(t)
	}
}
