package msg

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	//ERROR 失败
	ERROR = 1
	//SUCCESS 成功
	SUCCESS = 0
)

// JSONResponse 返回json
type JSONResponse struct {
	Code    int         `json:"code"` // 错误码((0:成功, 1:失败, >1:错误码))
	Data    interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
	Message string      `json:"msg"`  // 提示信息
}

// Result 返回json
func Result(r *ghttp.Request, code int, data interface{}, message string) {
	if err := r.Response.WriteJson(JSONResponse{
		code,
		data,
		message,
	}); err != nil {
		g.Log().Error(err)
	}
}

//Success 返回成功
func Success(r *ghttp.Request, data interface{}, message string) {
	Result(r, SUCCESS, data, message)
}

// Fail 返回失败
func Fail(r *ghttp.Request, data interface{}, message string) {
	Result(r, ERROR, data, message)

}

// JSON 标准返回结果数据结构封装。
func JSON(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	_ = r.Response.WriteJson(JSONResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// JSONExitgo 返回JSON数据并退出当前HTTP执行函数。
func JSONExitgo(r *ghttp.Request, err int, msg string, data ...interface{}) {
	JSON(r, err, msg, data...)
	r.Exit()
}
