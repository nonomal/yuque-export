package yuque

import (
	"fmt"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/sirupsen/logrus"
)

func NewDocDetail() *DocDetailData {
	return &DocDetailData{}
}

func (dd *DocDetailData) Get(h *handler.HandlerObject, name string) error {
	// 获取文档详情 URL
	endpoint := "/repos/" + fmt.Sprint(h.Namespace) + "/docs/" + name

	yc := handler.NewYuqueClient(h.Opts)
	err := yc.Request("GET", endpoint, dd)
	if err != nil {
		return err
	}

	return nil
}

func (dd *DocDetailData) Handle(h *handler.HandlerObject) error {
	return nil
}

func (dd *DocDetailData) GetDocDetailBodyHTML(h *handler.HandlerObject, slug string) (string, string, error) {
	err := dd.Get(h, slug)
	if err != nil {
		logrus.Error("获取 HTML 格式的文档详情失败")
		return "", "", err
	}
	return dd.Data.BodyHTML, dd.Data.Title, nil
}
