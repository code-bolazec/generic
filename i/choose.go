package i

func Choose[T any](b bool, t T, f T) T {
	if b {
		return t
	} else {
		return f
	}
}
