package app

type Repository interface{}

type Service interface{}

type Module struct {
	Repository Repository
	Service    Service
}
