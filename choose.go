package generic

func Choose(b bool, t, f string) string {
	if b {
		return t
	} else {
		return f
	}
}


func ChooseG[T any](b bool, t, f T) T {
	if b {
		return t
	} else {
		return f
	}
}

