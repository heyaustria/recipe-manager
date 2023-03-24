package recipe

type Writer interface {
	Create(*Model) (*Model, error)
	Update(*Model) (*Model, error)
}

type Reader interface {
	Read(uint64) (*Model, error)
	List() (*List, error)
}

type Service interface {
	Reader
	Writer
}

type Repository interface {
	Reader
	Writer
}
