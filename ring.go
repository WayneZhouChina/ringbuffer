package ring

type Ring struct {
	Size  int
	Start int
	End   int
	Elem  []interface{}
}

func New(size int) *Ring {
	if size < 0 {
		panic("Size must greater than 0")
	}
	return &Ring{Size: size, Elem: make([]interface{}, size)}
}

func (r *Ring) Get() interface{} {
	e := r.Elem[r.Start%r.Size]
	r.Start++
	return e
}

func (r *Ring) Set(i interface{}) {
	r.Elem[r.End%r.Size] = i
	r.End++
}

func (r *Ring) Len() int {
	return len(r.Elem)
}
