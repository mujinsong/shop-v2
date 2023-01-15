package upload

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"os"
)

func UploadImgToCloud(ctx context.Context, file *ghttp.UploadFile) (url string, err error) {
	dirpath := "/tmp/"
	name, err := file.Save(dirpath, true)
	if err != nil {
		return "", err
	}
	localFile := dirpath + name

	bucket := g.Cfg().MustGet(ctx, "qiniu.bucket").String()
	accesskey := g.Cfg().MustGet(ctx, "qiniu.accesskey").String()
	secretkey := g.Cfg().MustGet(ctx, "qiniu.secretkey").String()

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	mac := qbox.NewMac(accesskey, secretkey)

	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneXinjiapo
	cfg.UseHTTPS = true
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	key := name
	err = formUploader.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	//g.Dump(err)
	if err != nil {
		return "", err
	}
	fmt.Println(ret.Key, ret.Hash, ret.PersistentID)
	//删除本地临时文件
	err = os.RemoveAll(localFile)
	if err != nil {
		return "", err
	}
	url = g.Cfg().MustGet(ctx, "qiniu.url").String() + ret.Key
	return
}
