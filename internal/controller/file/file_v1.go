package file

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "life_notepad_api/api/file/v1"
	"life_notepad_api/internal/common/cos"
	"strings"
	"time"
)

func (c *ControllerV1) FileUpload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error) {
	if req.Files == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}

	names := make([]string, 0)
	for _, file := range req.Files {
		open, err := file.Open()
		if err != nil {
			return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
		}
		index := strings.LastIndex(file.Filename, ".")
		fileName := fmt.Sprintf("%s-%d", file.Filename[:index], gtime.Timestamp())
		if index >= 0 {
			fileName += file.Filename[index:]
		}
		err = cos.Cli.Upload(ctx, open, fileName)
		if err != nil {
			return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
		}
		names = append(names, fileName)
	}

	res = &v1.FileUploadRes{
		Names: names,
	}
	return
}

func (c *ControllerV1) FileUrls(ctx context.Context, req *v1.FileUrlsReq) (res *v1.FileUrlsRes, err error) {
	signUrls := make([]string, 0, len(req.Names))
	for _, fileName := range req.Names {
		signUrl, err := cos.Cli.GetPresignedURL(ctx, fileName, time.Hour*24)
		if err != nil {
			return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
		}
		signUrls = append(signUrls, signUrl)
	}
	res = &v1.FileUrlsRes{
		Urls: signUrls,
	}
	return

}
