package tests

import (
	goaltenapi "github.com/qbhy/goal-ten-api"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTenAPI(t *testing.T) {
	api := goaltenapi.Client{BaseUrl: "https://tenapi.cn/v2"}

	var (
		err  error
		hots []goaltenapi.Hot
	)

	hots, err = api.BaiduHots()
	assert.NoError(t, err, err)
	assert.NotEmpty(t, hots, "热点数据获取失败")
	time.Sleep(time.Second)

	hots, err = api.DouyinHots()
	assert.NoError(t, err, err)
	assert.NotEmpty(t, hots, "热点数据获取失败")
	time.Sleep(time.Second)

	hots, err = api.WeiboHots()
	assert.NoError(t, err, err)
	assert.NotEmpty(t, hots, "热点数据获取失败")
	time.Sleep(time.Second)

	hots, err = api.ZhihuHots()
	assert.NoError(t, err, err)
	assert.NotEmpty(t, hots, "热点数据获取失败")
	time.Sleep(time.Second)

	hots, err = api.BiliHots()
	assert.NoError(t, err, err)
	assert.NotEmpty(t, hots, "热点数据获取失败")
	time.Sleep(time.Second)

	hots, err = api.ToutiaoHots()
	assert.NoError(t, err, err)
	assert.NotEmpty(t, hots, "热点数据获取失败")
	time.Sleep(time.Second)

	// { "code": 201, "msg": "error" }
	//hots, err = api.ToutiaoHotNews()
	//assert.NoError(t, err, err)
	//assert.NotEmpty(t, hots, "热点数据获取失败")
	//time.Sleep(time.Second)

}
