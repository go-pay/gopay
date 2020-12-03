/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2020/12/02 20:20
*/

package gopay

import (
	"testing"

	"github.com/iGoogle-ink/gotil/xlog"
)

func TestBodyMapSetBodyMap(t *testing.T) {
	bm := make(BodyMap)
	// 1、配合map使用
	sceneInfo := make(map[string]map[string]string)
	h5Info := make(map[string]string)
	h5Info["type"] = "Wap"
	h5Info["wap_url"] = "http://www.gopay.ink"
	h5Info["wap_name"] = "H5测试支付"
	sceneInfo["h5_info"] = h5Info
	bm.Set("scene_info", sceneInfo)
	xlog.Debug("配合map使用：", bm) // map[scene_info:map[h5_info:map[type:Wap wap_name:H5测试支付 wap_url:http://www.gopay.ink]]]

	bm.Reset()
	xlog.Debug(bm) // []

	// 2、基础用法
	bm.Set("1key", "1value")
	bm.Set("2key", "2value")
	bm.Set("3key", "3value")
	xlog.Debug("基础用法：", bm)

	bm.Reset()
	xlog.Debug(bm) // []

	// 3、链式用法
	bm.Set("4key", "4value").
		Set("5key", "5value").
		Set("6key", "6value")
	xlog.Debug("链式用法：", bm)

	bm.Reset()
	xlog.Debug(bm) // []

	// 4、高级用法
	bm.SetBodyMap("scene_info", func(bm BodyMap) {
		bm.SetBodyMap("h5_info", func(bm BodyMap) {
			bm.Set("type", "Wap").
				Set("wap_url", "http://www.gopay.ink").
				Set("wap_name", "H5测试支付")
		})
	}).Set("7key", "7value").
		Set("8key", "8value")
	xlog.Debug("高级用法：", bm) // map[scene_info:map[h5_info:map[type:Wap wap_name:H5测试支付 wap_url:http://www.gopay.ink]]]

}
