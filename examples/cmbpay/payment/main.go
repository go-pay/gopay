// Command payment 演示一次完整的支付闭环：
//
//	收款码申请 → 轮询支付结果查询 → 退款 → 退款结果查询
//
// 并额外演示含敏感字段（terminalInfo）加密的付款码收款。
//
// 直接运行（使用接口文档附录 1 的联调示例参数）：
//
//	go run ./example/payment
//
// 或通过环境变量指定真实商户参数后运行（见 example/config）。
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-pay/gopay/cmbpay"
	"github.com/go-pay/gopay/examples/cmbpay/config"
)

func main() {
	client := config.MustClient()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	orderID := newOrderID()
	const cashier = "V003109945"

	// 1) 收款码申请（4.1）。金额单位为分。
	qr, err := client.QrCodeApply(ctx, &cmbpay.QrCodeApplyReq{
		MerID:        client.MerID(),
		OrderID:      orderID,
		UserID:       cashier,
		NotifyURL:    config.NotifyURL(),
		TxnAmt:       "1",
		PayValidTime: "900",
		Body:         "示例商品",
	})

	fmt.Println(client.MerID(), orderID, cashier)

	if err != nil {
		exit("收款码申请", err)
	}
	fmt.Printf("✓ 收款码申请成功\n  平台订单号: %s\n  二维码: %s\n", qr.BizContent.CmbOrderID, qr.BizContent.QrCode)

	// 2) 轮询支付结果查询（4.2）。
	// 实践建议：申请二维码约 15 秒后首查，之后每 5-10 秒一次，最多约 10 次；
	// 到达终态或超时后应调用关单接口。此处为演示仅快速轮询数次。
	final, err := pollOrder(ctx, client, orderID, cashier, 3, 2*time.Second)
	if err != nil {
		exit("支付结果查询", err)
	}
	fmt.Printf("✓ 查询到交易状态: %s（%s）\n", final.BizContent.TradeState, tradeStateText(final.BizContent.TradeState))

	// 若未支付，演示关单（4.7）。
	if final.BizContent.TradeState != cmbpay.TradeStateSuccess {
		if _, err := client.Close(ctx, &cmbpay.CloseReq{
			MerID:   client.MerID(),
			OrderID: orderID,
			UserID:  cashier,
		}); err != nil {
			handleErr("关闭订单", err)
		} else {
			fmt.Println("✓ 订单已关闭，可用新订单号重新发起支付")
		}
	}

	// 3) 若已支付，演示退款（4.4）与退款查询（4.5）。
	if final.BizContent.TradeState == cmbpay.TradeStateSuccess {
		refundID := newOrderID()
		if _, err := client.Refund(ctx, &cmbpay.RefundReq{
			MerID:      client.MerID(),
			OrderID:    refundID,                    // 商户退款单号，需唯一
			CmbOrderID: final.BizContent.CmbOrderID, // 原支付平台订单号
			UserID:     cashier,
			RefundAmt:  "1",
		}); err != nil {
			handleErr("退款申请", err)
		} else {
			fmt.Println("✓ 退款申请已受理")
			rq, err := client.RefundQuery(ctx, &cmbpay.RefundQueryReq{
				MerID:   client.MerID(),
				OrderID: refundID,
				UserID:  cashier,
			})
			if err != nil {
				handleErr("退款结果查询", err)
			} else {
				fmt.Printf("✓ 退款状态: %s\n", rq.BizContent.TradeState)
			}
		}
	}

	// 4) 含敏感字段加密的付款码收款（4.8 + 2.4.4）。
	demoBarcodePay(ctx, client, cashier)
}

// pollOrder 轮询支付结果查询，直到到达终态或达到最大次数。
func pollOrder(ctx context.Context, c *cmbpay.Client, orderID, cashier string, maxTimes int, interval time.Duration) (*cmbpay.OrderQueryResp, error) {
	var last *cmbpay.OrderQueryResp
	for i := 0; i < maxTimes; i++ {
		resp, err := c.OrderQuery(ctx, &cmbpay.OrderQueryReq{
			MerID:   c.MerID(),
			OrderID: orderID,
			UserID:  cashier,
		})
		if err != nil {
			var apiErr *cmbpay.APIError
			// SYSTERM_ERROR 表示结果不明确，应继续查询（接口文档 2.2 第 10 条）。
			if errors.As(err, &apiErr) && apiErr.IsSystemError() {
				time.Sleep(interval)
				continue
			}
			return nil, err
		}
		last = resp
		if isFinalState(resp.BizContent.TradeState) {
			return resp, nil
		}
		time.Sleep(interval)
	}
	if last == nil {
		return nil, errors.New("未查询到订单")
	}
	return last, nil
}

// demoBarcodePay 演示带 terminalInfo 加密的付款码收款。
func demoBarcodePay(ctx context.Context, c *cmbpay.Client, cashier string) {
	enc, err := cmbpay.NewEncryptor()
	if err != nil {
		handleErr("创建加密器", err)
		return
	}
	termCipher, err := enc.Encrypt(`{"location":"+37.12/-121.213","mobile_country_cd":"460"}`)
	if err != nil {
		handleErr("终端信息加密", err)
		return
	}
	pay, err := c.PayEncrypted(ctx, &cmbpay.PayReq{
		MerID:               c.MerID(),
		OrderID:             newOrderID(),
		UserID:              cashier,
		AuthCode:            "134650123456789012", // 用户付款码
		TxnAmt:              "1",
		EncryptTerminalInfo: termCipher,
	}, enc)
	if err != nil {
		handleErr("付款码收款", err)
		return
	}
	fmt.Printf("✓ 付款码收款状态: %s\n", pay.BizContent.TradeState)
}

func isFinalState(s string) bool {
	switch s {
	case cmbpay.TradeStateSuccess, cmbpay.TradeStateClosed,
		cmbpay.TradeStateCanceled, cmbpay.TradeStateFail, cmbpay.TradeStateRefunding:
		return true
	}
	return false
}

func tradeStateText(s string) string {
	switch s {
	case cmbpay.TradeStateSuccess:
		return "交易成功"
	case cmbpay.TradeStateClosed:
		return "订单已关闭"
	case cmbpay.TradeStateCanceled:
		return "交易已撤销"
	case cmbpay.TradeStateFail:
		return "交易失败"
	case cmbpay.TradeStateRefunding:
		return "转入退款"
	case cmbpay.TradeStateProcess:
		return "交易进行中"
	default:
		return "未知"
	}
}

// newOrderID 生成一个商户下唯一的订单号（时间戳 + 纳秒）。
// 生产环境建议结合业务前缀与分布式唯一序列生成。
func newOrderID() string {
	now := time.Now()
	return fmt.Sprintf("D%s%09d", now.Format("20060102150405"), now.Nanosecond())
}

// handleErr 打印错误但不中断流程（用于演示中非关键步骤）。
func handleErr(scene string, err error) {
	var apiErr *cmbpay.APIError
	if errors.As(err, &apiErr) {
		if apiErr.IsSystemError() {
			log.Printf("• %s 遇到 SYSTERM_ERROR，交易结果不明确，需发起查询确认", scene)
			return
		}
		if apiErr.ErrDescription != "" {
			log.Printf("• %s 失败: returnCode=%s respCode=%s errCode=%s msg=%s errDescription=%s",
				scene, apiErr.ReturnCode, apiErr.RespCode, apiErr.ErrCode, apiErr.RespMsg, apiErr.ErrDescription)
		} else {
			log.Printf("• %s 失败: returnCode=%s respCode=%s errCode=%s msg=%s",
				scene, apiErr.ReturnCode, apiErr.RespCode, apiErr.ErrCode, apiErr.RespMsg)
		}
		return
	}
	log.Printf("• %s 请求错误: %v", scene, err)
}

// exit 打印错误并退出（用于演示中的关键步骤）。
func exit(scene string, err error) {
	handleErr(scene, err)
	log.Fatalf("%s 失败，终止示例", scene)
}
