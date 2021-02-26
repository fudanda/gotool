package video

import (
	"write_code_in_/app/model/video"

	"write_code_in_/app/ship/msg"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// List is a demonstration route handler for output "Hello World!".
func List(r *ghttp.Request) {
	var req *video.SearchReq
	var data []interface{}
	//获取参数
	if err := r.Parse(&req); err != nil {
		msg.Fail(r, data, "失败")
	}
	total, list, err := video.GetList(req)
	if err != nil {
		msg.Fail(r, data, "失败")
	}
	res := g.Map{
		"total": total,
		"list":  list,
	}
	msg.Success(r, res, "成功")
}
