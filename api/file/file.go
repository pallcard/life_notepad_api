package file

import (
	"context"
	v1 "life_notepad_api/api/file/v1"
)

type IFileV1 interface {
	FileUpload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error)
	FileUrls(ctx context.Context, req *v1.FileUrlsReq) (res *v1.FileUrlsRes, err error)
}
