package mytools

const PAI = 3.14

func GetC(r float64) float64 {
	return 2 * PAI * r
}

func GetS(r float64) float64 {
	return PAI * r * r
}

func GetL(r float64) float64 {
	return 2 * r
}
