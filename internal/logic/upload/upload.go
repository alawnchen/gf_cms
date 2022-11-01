package upload

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/os/gtime"
	"os"
)

type sUpload struct{}

var (
	insUpload = sUpload{}
)

func init() {
	service.RegisterUpload(New())
}

func New() *sUpload {
	return &sUpload{}
}

func Upload() *sUpload {
	return &insUpload
}

// SingleUploadFile 上传文件
func (*sUpload) SingleUploadFile(ctx context.Context, in model.FileUploadInput, dir string) (out *backendApi.UploadFileRes, err error) {
	serverRoot := service.Util().ServerRoot()
	os.MkdirAll(serverRoot, 0777)
	os.Chmod(serverRoot, 0777)
	fullUploadDir := "/upload/" + dir + "/" + gtime.Date()
	fullDir := serverRoot + fullUploadDir
	filename, err := in.File.Save(fullDir, in.RandomName)
	if err != nil {
		return nil, err
	}
	url := fullUploadDir + "/" + filename
	out = &backendApi.UploadFileRes{
		Name: filename,
		Url:  url,
	}
	return
}
