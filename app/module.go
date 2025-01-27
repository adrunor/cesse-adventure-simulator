package app

type Repository interface {
	Insert(interface{}) (interface{}, error)
	Update(interface{}) (interface{}, error)
	Delete(interface{}) (interface{}, error)
	Find(interface{}, interface{}) (interface{}, error)
}

type Service interface{}

type Module struct {
	Repository Repository
	Service    Service
}
