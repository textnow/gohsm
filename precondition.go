package hsm

func Precondition(expression bool, message string) {
	if !expression {
		panic(message)
	}
}
