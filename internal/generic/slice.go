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

func Insert[T any](ts []T, t T, idx int) []T {
	ts = append(ts[:idx+1], ts[idx:]...)
	ts[idx] = t

	return ts
}

func Pop[T any](ts []T, idx int) ([]T, T) {
	v := ts[idx]
	ts = append(ts[:idx], ts[idx+1:]...)
	return ts, v
}
