package nicepkg

// +build gofuzz

func Fuzz(data []byte) int {
	if err := CoolFunc(string(data)); err != nil {
		return 0
	}
	return 1
}
