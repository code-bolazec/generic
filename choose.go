package generic

func If[T any](b bool, t, f T) T {
	if b {
		return t
	} else {
		return f
	}
}
