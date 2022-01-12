package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

// 是否有必要抽象出来这层接口？
type YuqueData interface {
	Get()
}

// 处理语雀 API 是要使用的 HTTP 处理器
func HttpHandler(method, url, token string, data interface{}) (interface{}, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
