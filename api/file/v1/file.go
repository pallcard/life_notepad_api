package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/fileUpload" tags:"File" mime:"multipart/form-data" method:"post" summary:"You first hello api"`
	Files  []*ghttp.UploadFile `json:"files" type:"file" dc:"选择上传文件"`
}

type FileUploadRes struct {
	Names []string
}

type FileUrlsReq struct {
	g.Meta `path:"/fileUrls" tags:"File" method:"post" summary:"获取笔记列表"`
	Names  []string
}

type FileUrlsRes struct {
	g.Meta `mime:"application/json"`
	Urls   []string
}
