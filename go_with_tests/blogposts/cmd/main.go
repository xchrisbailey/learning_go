package main

import (
	"log"
	"os"

	"githhub.com/xchrisbailey/learning_go/go_with_tests/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(posts)
}
