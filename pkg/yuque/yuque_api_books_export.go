package yuque

import (
	"encoding/json"
	"fmt"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
)

// ReqBodyForGetExportURL is
type ReqBodyForExportData struct {
	Type         string `json:"type"`
	Force        int    `json:"force"`
	Title        string `json:"title"`
	TocNodeUUID  string `json:"toc_node_uuid"`
	TocNodeURL   string `json:"toc_node_url"`
	WithChildren bool   `json:"with_children"`
}

func NewExportsData() *ExportsData {
	return &ExportsData{}
}

func (e *ExportsData) Get(h *handler.HandlerObject, name string) error {
	endpoint := fmt.Sprintf("/books/%s/export", h.Namespace)

	// 根据节点信息，配置当前待导出节点的请求体信息
	// 解析请求体
	reqBodyByte, err := json.Marshal(ReqBodyForExportData{
		Type:  "lakebook",
		Force: 0,
		// Title:        toc.Title,
		// TocNodeUUID:  toc.UUID,
		// TocNodeURL:   toc.URL,
		TocNodeUUID:  name,
		WithChildren: true,
	})
	if err != nil {
		return err
	}

	yc := handler.NewYuqueClient(h.Flags)
	err = yc.Request("POST", endpoint, reqBodyByte, e)
	if err != nil {
		return err
	}

	return nil
}

func (e *ExportsData) Handle(h *handler.HandlerObject) error {
	panic("not implemented") // TODO: Implement
}