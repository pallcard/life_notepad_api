package cos

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	cli *cos.Client
}

var Cli = new(Client)

func init() {
	secretID, _ := g.Cfg().Get(gctx.New(), "config.secretID")
	secretKey, _ := g.Cfg().Get(gctx.New(), "config.secretKey")
	cosUrl, _ := g.Cfg().Get(gctx.New(), "config.cosUrl")
	u, _ := url.Parse(cosUrl.String())
	su, _ := url.Parse("https://cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	Cli.cli = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID.String(),
			SecretKey: secretKey.String(),
		},
	})
}

func (s *Client) Upload(ctx context.Context, file io.Reader, fileName string) error {
	_, err := s.cli.Object.Put(ctx, fileName, file, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Client) GetPresignedURL(ctx context.Context, fileName string, expired time.Duration) (string, error) {
	// 获取预签名 URL
	secretID, _ := g.Cfg().Get(gctx.New(), "config.secretID")
	secretKey, _ := g.Cfg().Get(gctx.New(), "config.secretKey")
	signedURL, err := s.cli.Object.GetPresignedURL(ctx,
		http.MethodGet,
		fileName,
		secretID.String(), secretKey.String(),
		expired, nil)
	if err != nil {
		return "", err
	}

	return signedURL.String(), nil
}
