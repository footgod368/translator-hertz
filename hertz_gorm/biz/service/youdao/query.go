package youdao

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/utils"
	"github.com/cloudwego/hertz/pkg/common/json"
	"io"
	"net/http"
)

const YouDaoAPIURL = "http://dict.youdao.com/jsonapi"

func Query(ctx context.Context, text string) (*YouDaoResponse, error) {
	response, err := http.Get(utils.AppendURLParams(YouDaoAPIURL, map[string]string{"q": text}))
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("youdao api error: %s", response.Status)
	}
	youDaoResponse := &YouDaoResponse{}
	if err = json.Unmarshal(body, youDaoResponse); err != nil {
		return nil, err
	}
	return youDaoResponse, nil
}
