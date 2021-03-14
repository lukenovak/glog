package constants

import "html/template"

type Post struct {
	Id int
	Title string
	Body template.HTML // array of paragraphs
}

type IncomingPostJson struct {
	Title string
	Body []string
}