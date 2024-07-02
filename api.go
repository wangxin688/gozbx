package gozbx

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/go-resty/resty/v2"
)

var requestId uint64

type ZbxAPI struct {
	token     string
	client    *resty.Client
	id        uint64
	HostGroup *HosGroupImpl
}

type Request struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
	Id      uint64 `json:"id"`
}

type ZbxApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type Response struct {
	JsonRpc string      `json:"jsonrpc"`
	Result  any         `json:"result"`
	Error   ZbxApiError `json:"error"`
	Id      uint64      `json:"id"`
}

func (rsp *Response) IsError() error {
	if rsp.Error.Code != 0 {
		return fmt.Errorf("Request failed: %d %s %s", rsp.Error.Code, rsp.Error.Data, rsp.Error.Message)
	}
	return nil
}

func (rsp *Response) GetResult(v any) error {
	json.Unmarshal(rsp.Result.([]byte), v)
	return nil
}

func (z *ZbxAPI) NewZbxAPI(url string) *ZbxAPI {
	url = strings.TrimSuffix(url, "/")
	client := resty.New().SetBaseURL(url).SetHeader("Content-Type", "application/json-rpc")
	return &ZbxAPI{
		token:     "",
		client:    client,
		id:        requestId,
		HostGroup: &HosGroupImpl{z: z},
	}
}

func (z *ZbxAPI) SetToken(token string) {
	z.client.SetHeader("Authorization", "Bearer "+token)
}

func (z *ZbxAPI) SetTimeout(timeout int) {
	z.client.SetTimeout(time.Duration(timeout) * time.Second)
}

func (z *ZbxAPI) SetRetryCount(retry int) {
	z.client.SetRetryCount(retry)
}

func (z *ZbxAPI) Login(username, password string) error {
	params := map[string]string{"username": username, "password": password}
	request := &Request{
		JsonRpc: "2.0",
		Method:  "user.login",
		Params:  params,
		Id:      atomic.AddUint64(&requestId, 1),
	}
	resp, err := z.client.R().SetBody(request).Post("/api_jsonrpc.php")
	if err != nil {
		return err
	}
	rsp := &Response{}
	if err := json.Unmarshal(resp.Body(), &rsp); err != nil {
		return err
	}

	if err := rsp.IsError(); err != nil {
		return err
	}
	z.token = rsp.Result.(string)
	return nil
}

//
func (z *ZbxAPI) Rpc(req *Request) (*Response, error) {
	if z == nil {
		return nil, fmt.Errorf("ZbxAPI instance is nil")
	}
	if z.token == "" {
		return nil, fmt.Errorf("authentication required, please login first or set token first")
	}
	req.Id = atomic.AddUint64(&requestId, 1)
	req.JsonRpc = "2.0"
	resp, err := z.client.R().SetBody(req).Post("/api_jsonrpc.php")
	if err != nil {
		return nil, err
	}
	rsp := &Response{}
	err = json.Unmarshal(resp.Body(), rsp)
	if err != nil {
		return nil, err
	}
	if err = rsp.IsError(); err != nil {
		return nil, err
	}
	return rsp, nil
}
