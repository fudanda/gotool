package router

import (
	"write_code_in_/app/api/hello"
	"write_code_in_/app/api/video"
	"write_code_in_/app/api/word"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", hello.Hello)
		group.ALL("/word", word.Create)
		group.ALL("/files", word.GetFile)
		group.ALL("/download/:name", word.Download)
		group.ALL("/delete/:name", word.Delete)
		group.ALL("/video", video.List)

	})
}
