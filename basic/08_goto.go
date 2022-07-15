package basic

func TestGoto() {
	err := true
	if err {
		goto handleErr
	}
handleErr:
	panic("there is an error")
}
