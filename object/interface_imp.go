package object

type Integer int

func (i *Integer) Add(x Integer) {
	*i += x
}
