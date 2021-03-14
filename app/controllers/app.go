package controllers

import (
	_ "github.com/lib/pq"
	"github.com/revel/revel"
	"github.com/lukenovak/glog/app/services"
)

type App struct {
	*revel.Controller
}

const ERROR_404 = "errors/404.html"
const ERROR_500 = "errors/500.html"

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) NewPost() revel.Result {
	return c.Render()
}

func (c App) Posts() revel.Result {
	posts, err := services.GetNumPosts(10)
	if err != nil {
		return c.RenderError(err)
	}
	return c.Render(posts)
}