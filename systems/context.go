package systems

import "github.com/gin-gonic/gin"

type Env interface {
	Get(string) (string, error)
}

type Database interface{}

type Router interface {
	GetRouter() *gin.Engine
	Run()
}

type Context struct {
	env      Env
	database Database
	router   Router
}

func (c *Context) Env() Env {
	return c.env
}

func (c *Context) Database() Database {
	return c.database
}

func (c *Context) Router() Router {
	return c.router
}

func NewContext(env Env, database Database, router Router) *Context {
	return &Context{
		env:      env,
		database: database,
		router:   router,
	}
}
