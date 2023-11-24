package main

type laptopSize float64

func (l laptopSize) getSize() laptopSize {
	return l
}

func main() {
	l := laptopSize(133.5)
	println(l.getSize())
}
