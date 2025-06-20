// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	translator "github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/router/translator"
	user_gorm "github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/router/user_gorm"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	translator.Register(r)

	user_gorm.Register(r)
}
