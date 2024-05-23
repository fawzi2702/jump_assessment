package utils

func Map[T, V any](Ts []T, fn func(T) V) []V {
	Vs := make([]V, len(Ts))

	for idx, t := range Ts {
		Vs[idx] = fn(t)
	}

	return Vs
}
