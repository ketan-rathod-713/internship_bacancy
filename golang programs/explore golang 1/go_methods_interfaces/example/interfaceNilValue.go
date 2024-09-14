package example

// TODO what to do
type I interface {
	M()
}

func InterfaceNilValue() {
	var i I
	describe(i)
	i.M() // There is no type hence it will produce a run time error // not compile time ha ha
}
