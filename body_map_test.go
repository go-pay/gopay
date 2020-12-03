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
	sceneInfo := make(map[string]map[string]string)
	h5Info := make(map[string]string)
	h5Info["type"] = "Wap"
	h5Info["wap_url"] = "http://www.gopay.ink"
	h5Info["wap_name"] = "H5测试支付"
	sceneInfo["h5_info"] = h5Info
	bm.Set("scene_info", sceneInfo)
	xlog.Debug(bm) // map[scene_info:map[h5_info:map[type:Wap wap_name:H5测试支付 wap_url:http://www.gopay.ink]]]

	bm.Reset()
	xlog.Debug(bm) // []

	bm.SetBodyMap("scene_info", func(bm BodyMap) {
		bm.SetBodyMap("h5_info", func(bm BodyMap) {
			bm.Set("type", "Wap").
				Set("wap_url", "http://www.gopay.ink").
				Set("wap_name", "H5测试支付")
		})
	}).Set("额外", "哇").
		Set("sada", "qedfs")
	xlog.Debug(bm) // map[scene_info:map[h5_info:map[type:Wap wap_name:H5测试支付 wap_url:http://www.gopay.ink]]]
}
