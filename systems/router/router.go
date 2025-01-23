package router

import "github.com/gin-gonic/gin"

type Impl struct {
	router *gin.Engine
}

func (r *Impl) GetRouter() *gin.Engine {
	return r.router
}

func (r *Impl) Run() {
	if err := r.router.Run(":8080"); err != nil {
		panic(err)
	}
}

func NewRouter() *Impl {
	return &Impl{
		router: gin.Default(),
	}
}
