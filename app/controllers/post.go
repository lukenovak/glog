package controllers

import (
	"fmt"
	"github.com/lukenovak/goblog/app/models"
	"github.com/revel/revel"
	"github.com/lukenovak/goblog/app/services"
)

func (c App) ShowPost(id int) revel.Result {
	post, err := services.GetPostFromDB(id)
	if err != nil {
		c.Log.Error(err.Error())
		return c.RenderTemplate(ERROR_404)
	}
	return c.Render(post)
}

func (c App) CreatePost() revel.Result {
	var newPost models.IncomingPostJson
	c.Params.BindJSON(&newPost)
	c.Log.Info(fmt.Sprintf("%+v", newPost))
	if len(newPost.Title) + len(newPost.Body) < 2 {
		return c.RenderText("Post Title or Body is blank")
	}
	newId, err := services.CreatePostInDB(newPost)
	if err != nil {
		c.Log.Error(err.Error())
		c.RenderError(err)
	}
	c.Log.Info(fmt.Sprintf("Successfully created new post with id = %d", newId))
	return c.RenderText("nice!")
}