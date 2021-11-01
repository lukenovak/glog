package services

import (
	"database/sql"
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/lukenovak/goblog/app"
	"github.com/lukenovak/goblog/app/models"
	"html/template"
	"strings"
)

const CREATE_POST_QUERY = "INSERT INTO post (title, post_date, body_text, post_author) VALUES ($1, NOW(), $2, -1)"
const GET_ONE_POST_BY_ID_QUERY = "SELECT * FROM post WHERE post_id = $1"

// TODO: Move these to the controller file- this is a mess, and there's no need for this number of functions

func CreatePostInDB(post models.IncomingPostJson) (int64, error) {
	bodyString := strings.Join(post.Body, "\n")
	result, err := app.DB.Exec(CREATE_POST_QUERY, post.Title, bodyString)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// takes an open connection and gets the post from the db connection
func GetPostFromDB(postId int) (models.Post, error) {
	var emptyPost models.Post
	rows, err := app.DB.Query(GET_ONE_POST_BY_ID_QUERY, postId)
	if rows.Next() {
		return scanToPost(rows)
	}
	err = fmt.Errorf("no row found")
	return emptyPost, err
}

// gets the number given of the most recent posts
func GetNumMostRecentPostsFromDB(n int) ([]models.Post, error) {
	unpreppedQuery := fmt.Sprintf("SELECT * FROM post ORDER BY post_date LIMIT %d", n)
	rows, err := app.DB.Query(unpreppedQuery)
	if err != nil {
		return nil, err
	}
	var postList []models.Post
	for rows.Next() {
		var post models.Post
		post, err = scanToPost(rows)
		postList = append(postList, post)
	}
	return postList, err
}

// scans a single row to a post
func scanToPost(rows *sql.Rows) (models.Post, error) {
	var post models.Post
	var date, author, unseparatedBody string
	err := rows.Scan(&post.Id, &date, &author, &post.Title, &unseparatedBody)
	post.Body = template.HTML(markdown.ToHTML([]byte(unseparatedBody), nil, nil))
	return post, err
}
