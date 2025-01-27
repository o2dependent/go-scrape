package utils

func Filter[T any](ss []T, test func(T) bool) (r []T) {
	for _, s := range ss {
		if test(s) {
			r = append(r, s)
		}
	}
	return
}
