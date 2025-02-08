package utils

func Map[I, O any, S ~[]I](s S, mf func(I) O) []O {
	out := make([]O, 0, len(s))
	for i := range len(s) {
		out = append(out, mf(s[i]))
	}
	return out
}

func Filter[E any, S ~[]E](s S, ff func(E) bool) S {
	out := make(S, 0)
	for i := range s {
		if ff(s[i]) {
			out = append(out, s[i])
		}
	}
	return out
}

func Contains[E any, S ~[]E](s S, ff func(E) bool) bool {
	return len(Filter(s, ff)) > 0
}

func Group[E any, S ~[]E](s S, kf func(E) string) map[string][]E {
	grouped := make(map[string][]E)

	for _, i := range s {
		key := kf(i)
		grouped[key] = append(grouped[key], i)
	}
	return grouped
}

func ToMap[T any, V comparable](src []T, key func(T) V) map[V]T {
	result := make(map[V]T)
	for _, v := range src {
		result[key(v)] = v
	}
	return result
}
