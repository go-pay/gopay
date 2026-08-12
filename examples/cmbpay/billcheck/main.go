package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-pay/gopay/cmbpay"
	"github.com/go-pay/gopay/examples/cmbpay/config"
)

func main() {

	client := config.MustClient()

	// 查询对账文件（使用对账接口专用主机地址）
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &cmbpay.BillRecordReq{
		MerchantNo: client.MerID(),
		BillDate:   "2026-07-10",
		BillType:   "JH_JZ",
	}

	// 使用 BillRecord 方法，传入对账接口专用主机地址
	resp, err := client.BillRecord(ctx, cmbpay.BillCheckHostUAT, req)

	fmt.Println(resp, err)

	if err != nil {
		log.Fatalf("查询对账文件失败: %v", err)
	}

	// 处理响应
	if resp.IsBillReady() {
		fmt.Printf("对账文件已生成\n")
		fmt.Printf("商户号: %s\n", resp.Data.MerchantNo)
		fmt.Printf("账单类型: %s\n", resp.Data.BillType)
		fmt.Printf("账单日期: %s\n", resp.Data.BillDate)
		fmt.Printf("下载链接: %s\n", resp.Data.DownloadUrl)
		fmt.Printf("清分状态: %s (%s)\n", resp.Data.SettleStatus, resp.Data.GetSettleStatusDesc())
	} else if resp.IsBillNotGenerated() {
		fmt.Printf("账单还未生成，请稍后再试\n")
	} else if resp.IsNoBill() {
		fmt.Printf("当日无账单\n")
	}
}
