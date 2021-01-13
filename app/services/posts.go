package services

import (
	"database/sql"
	"fmt"
	"glog/app/constants"
	"strings"
)

const CREATE_POST_QUERY = "INSERT INTO post_data (title, post_date, body_text, post_author) VALUES ($1, NOW(), $2, -1)"
const GET_ONE_POST_BY_ID_QUERY = "SELECT * FROM post_data WHERE post_id = $1"


// gets up to the given number of posts and returns them
func GetNumPosts(n int) ([]constants.Post, error) {
	dbConn, err := sql.Open(constants.POSTGRES, constants.PsqlInfo)
	defer dbConn.Close()
	if err != nil {
		return nil, err
	}
	return GetNumMostRecentPostsFromDB(n, dbConn)
}

func CreatePostInDB(post constants.IncomingPostJson, db *sql.DB) (int64, error) {
	bodyString := strings.Join(post.Body, "\n")
	result, err := db.Exec(CREATE_POST_QUERY, post.Title, bodyString)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// takes an open connection and gets the post from the db connection
func GetPostFromDB(postId int, db *sql.DB) (constants.Post, error) {
	var emptyPost constants.Post
	rows, err := db.Query(GET_ONE_POST_BY_ID_QUERY, postId)
	if rows.Next() {
		return scanToPost(rows)
	}
	err = fmt.Errorf("no row found")
	return emptyPost, err
}

// gets the number given of the most recent posts
func GetNumMostRecentPostsFromDB(n int, db *sql.DB) ([]constants.Post, error) {
	unpreppedQuery := fmt.Sprintf("SELECT * FROM post_data ORDER BY post_date LIMIT %d", n)
	rows, err := db.Query(unpreppedQuery)
	if err != nil {
		return nil, err
	}
	var postList []constants.Post
	for rows.Next() {
		var post constants.Post
		post, err = scanToPost(rows)
		postList = append(postList, post)
	}
	return postList, err
}

// scans a single row to a post
func scanToPost(rows *sql.Rows) (constants.Post, error) {
	var post constants.Post
	var date, author, unseparatedBody string
	err := rows.Scan(&post.Id, &date, &author, &post.Title, &unseparatedBody)
	post.Body = strings.Split(unseparatedBody, "\n")
	return post, err
}
