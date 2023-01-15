package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"shop-v2/api/backend"
	"shop-v2/internal/consts"
	"shop-v2/utility/upload"
)

type cUpload struct {
}

var Upload = cUpload{}

func (c *cUpload) UploadImgTocloud(ctx context.Context, req *backend.UploadImgToCloudReq) (res *backend.UploadImgTocloudRes, err error) {
	//g.Dump(req)
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, consts.CodeMissingParameterMsg)
	}
	url, err := upload.UploadImgToCloud(ctx, req.File)
	if err != nil {
		return nil, err
	}
	return &backend.UploadImgTocloudRes{
		Url: url,
	}, nil
}
