package main

import (
	"github.com/OpenAudioMc/documentation/md_loader"
	"github.com/OpenAudioMc/documentation/renderer"
)

func main() {
	// load source pages
	var pages = md_loader.LoadPages()
	renderer.RenderPages(pages)
}
