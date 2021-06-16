package generic

func Choose(b bool, t, f string) string {
	if b {
		return t
	} else {
		return f
	}
}

/*
func xChoose[T any](b bool, t T, f T) T {
	if b {
		return t
	} else {
		return f
	}
}
*/
