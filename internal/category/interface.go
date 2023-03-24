package category

type Writer interface {
	Create(*Model) (*Model, error)
	Update(*Model) (*Model, error)
}

type Reader interface {
	Read(int64) (*Model, error)
	List() (*List, error)
	Exists([]*Model) error
}

type Service interface {
	Reader
	Writer
}

type Repository interface {
	Reader
	Writer
}
