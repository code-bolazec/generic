package generic

func ChooseString(b bool, t, f string) string {
	if b {
		return t
	} else {
		return f
	}
}

func ChooseAny(b bool, t, f interface{}) interface{} {
	if b {
		return t
	} else {
		return f
	}
}

/*
func ChooseG[T any](b bool, t, f T) T {
	if b {
		return t
	} else {
		return f
	}
}
*/
