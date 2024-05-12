package persist

type Persist interface {
	Close()
	Insert(name string) error
	FindById(id string) ([]string, error)
}
