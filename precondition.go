package hsm

func Precondition(srv Service, expression bool, message string) {
	if !expression {
		srv.Logger().Panic(message)
		panic(message)
	}
}
