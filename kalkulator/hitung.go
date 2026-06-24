package kalkulator

func Tambah(a int, b int) int {
	return a + b
}

func Selisih(a int, b int ) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}