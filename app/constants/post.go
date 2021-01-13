package constants

const PostPath = "ShowPost.html"

type Post struct {
	Id int
	Title string
	Body []string // array of paragraphs
}

type IncomingPostJson struct {
	Title string
	Body []string
}