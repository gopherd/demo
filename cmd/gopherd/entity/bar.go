package entity

type Bar struct {
	id int64
}

func NewBar(id int64) *Bar {
	return &Bar{
		id: id,
	}
}

func (bar *Bar) ID() int64 {
	return bar.id
}
