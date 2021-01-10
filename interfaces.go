package main

import "fmt"

type feed interface {
	Render() string
}
type news struct {
}

func (n news) Render() string {
	return "news"
}

type images struct {
}

func (img images) Render() string {
	return "image"
}
func main() {
	f := []feed{news{}, images{}}
	for _, Feed := range f {
		fmt.Println(Feed.Render())
	}
}
