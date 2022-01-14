package yuque

import (
	"github.com/DesistDaydream/yuque-export/pkg/handler"
)

func NewDocsList() *DocsList {
	return &DocsList{}
}

func (d *DocsList) Get(h *handler.HandlerObject, name string) error {
	endpoint := "/repos/" + name + "/docs"

	yc := handler.NewYuqueClient(h.Opts)
	err := yc.Request("GET", endpoint, d)
	if err != nil {
		return err
	}

	return nil
}
func (d *DocsList) Handle(h *handler.HandlerObject) error {
	return nil
}
