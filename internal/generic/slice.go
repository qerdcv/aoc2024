package generic

func PopStart[T any](ts []T) ([]T, T) {
	if len(ts) == 0 {
		return ts, *new(T)
	}

	t := ts[0]
	newTs := ts[1:]
	return newTs, t
}

func PopEnd[T any](ts []T) ([]T, T) {
	if len(ts) == 0 {
		return ts, *new(T)
	}

	t := ts[len(ts)-1]
	newTs := ts[:len(ts)-1]
	return newTs, t
}
