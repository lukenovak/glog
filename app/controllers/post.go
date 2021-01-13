package controllers

import (
	"database/sql"
	"fmt"
	"github.com/revel/revel"
	"glog/app/constants"
	"glog/app/services"
	"strings"
)

func (c App) ShowPost(id int) revel.Result {
	DB_CONN, _ := sql.Open(constants.POSTGRES, constants.PsqlInfo)
	c.Log.Info(fmt.Sprintf("DB_CONN is %p", DB_CONN))
	post, err := services.GetPostFromDB(id, DB_CONN)
	if err != nil {
		c.Log.Error(err.Error())
		return c.RenderTemplate(ERROR_404)
	}
	c.Log.Info(fmt.Sprintf("%d, %s, %s", post.Id, post.Title, post.Body))
	if len(post.Body) < 2 {
		post.Body = strings.Split(post.Body[0], "\n")
	}
	return c.Render(post)
}

func (c App) CreatePost() revel.Result {
	dbConn, _ := sql.Open(constants.POSTGRES, constants.PsqlInfo)
	defer dbConn.Close()
	var newPost constants.IncomingPostJson
	c.Params.BindJSON(&newPost)
	c.Log.Info(fmt.Sprintf("%+v", newPost))
	if len(newPost.Title) + len(newPost.Body) < 2 {
		return c.RenderText("you done fucked up")
	}
	newId, err := services.CreatePostInDB(newPost, dbConn)
	if err != nil {
		c.Log.Error(err.Error())
		c.RenderError(err)
	}
	c.Log.Info(fmt.Sprintf("Successfully created new post with id = %d", newId))
	return c.RenderText("nice!")
}