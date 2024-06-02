package goal_ten_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Hot struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Hot  string `json:"hot,omitempty"`
}

type HotResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []Hot  `json:"data"`
}

type Client struct {
	BaseUrl string
}

func (c *Client) fetchHots(endpoint string) ([]Hot, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response HotResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if response.Msg != "success" {
		return nil, errors.New(response.Msg)
	}

	return response.Data, nil
}

func (c *Client) BaiduHots() ([]Hot, error) {
	return c.fetchHots("baiduhot")
}

func (c *Client) DouyinHots() ([]Hot, error) {
	return c.fetchHots("douyinhot")
}

func (c *Client) WeiboHots() ([]Hot, error) {
	return c.fetchHots("weibohot")
}

func (c *Client) ZhihuHots() ([]Hot, error) {
	return c.fetchHots("zhihuhot")
}

func (c *Client) BiliHots() ([]Hot, error) {
	return c.fetchHots("bilihot")
}

func (c *Client) ToutiaoHots() ([]Hot, error) {
	return c.fetchHots("toutiaohot")
}

func (c *Client) ToutiaoHotNews() ([]Hot, error) {
	return c.fetchHots("toutiaohotnew")
}
