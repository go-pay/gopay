// Command notifyserver 演示一个可运行的异步通知接收服务（接口文档 4.3 / 2.5）。
//
// 它完成：验签 → 幂等判重 → 更新本地订单 → 返回标准应答。
//
// 运行：
//
//	go run ./example/notifyserver
//
// 然后招行（或你的联调工具）向 http://<本机>:8080/cmb/notify 以 form-data POST 通知。
package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/go-pay/gopay/examples/cmbpay/config"
)

// orderStore 是一个演示用的内存幂等存储。生产环境应替换为数据库，
// 并在状态检查与更新之间使用行锁/唯一约束等手段做并发控制（接口文档 2.5）。
type orderStore struct {
	mu   sync.Mutex
	done map[string]bool // key: 商户订单号
}

func (s *orderStore) markIfNew(orderID string) (isNew bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.done[orderID] {
		return false
	}
	s.done[orderID] = true
	return true
}

func main() {
	client := config.MustClient()
	store := &orderStore{done: make(map[string]bool)}

	http.HandleFunc("/cmb/notify", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// 1) 验签并解析（内部完成 form-data 的 url_decode 与 SM2 验签）。
		data, err := client.ParseNotify(r)
		if err != nil {
			log.Printf("通知验签失败: %v", err)
			// 验签失败返回 FAIL，招行会按重试策略再次通知。
			_, _ = w.Write(client.NotifyFailBody("verify failed"))
			return
		}

		// 2) 幂等：同一通知可能重复投递，已处理过直接返回成功。
		if !store.markIfNew(data.OrderID) {
			log.Printf("通知重复，幂等跳过: 订单=%s", data.OrderID)
			_, _ = w.Write(client.NotifySuccessBody())
			return
		}

		// 3) 业务处理：招行仅在支付成功时发送支付结果通知。
		if data.IsPaySuccess() {
			log.Printf("支付成功: 订单=%s 平台单号=%s 金额(分)=%s 方式=%s 第三方单号=%s",
				data.OrderID, data.CmbOrderID, data.TxnAmt, data.PayType, data.ThirdOrderID)
			// TODO: 更新本地订单为已支付、触发发货等。
		} else {
			log.Printf("收到非成功状态通知: 订单=%s 状态=%s", data.OrderID, data.TradeState)
		}

		// 4) 返回成功应答，招行据此停止重复通知。
		_, _ = w.Write(client.NotifySuccessBody())
	})

	addr := ":8080"
	log.Printf("通知服务已启动: http://localhost%s/cmb/notify", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
