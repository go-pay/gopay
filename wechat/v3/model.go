package wechat

type OrderNoType uint8

type PlatformCertRsp struct {
	Code  int                 `json:"-"`
	Certs []*PlatformCertItem `json:"certs"`
	Error string              `json:"-"`
}

type EmptyRsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Error    string    `json:"-"`
}

// Prepay 支付Rsp
type PrepayRsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Response *Prepay   `json:"response,omitempty"`
	Error    string    `json:"-"`
}

// H5 支付Rsp
type H5Rsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Response *H5Url    `json:"response,omitempty"`
	Error    string    `json:"-"`
}

// Native 支付Rsp
type NativeRsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Response *Native   `json:"response,omitempty"`
	Error    string    `json:"-"`
}

// 查询订单 Rsp
type QueryOrderRsp struct {
	Code     int         `json:"-"`
	SignInfo *SignInfo   `json:"-"`
	Response *QueryOrder `json:"response,omitempty"`
	Error    string      `json:"-"`
}

// 关闭订单 Rsp
type CloseOrderRsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Error    string    `json:"-"`
}

// 退款 Rsp
type RefundRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *RefundOrderResponse `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 退款查询 Rsp
type RefundQueryRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *RefundQueryResponse `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 交易、资金账单 Rsp
type BillRsp struct {
	Code     int        `json:"-"`
	SignInfo *SignInfo  `json:"-"`
	Response *TradeBill `json:"response,omitempty"`
	Error    string     `json:"-"`
}

// 二级商户资金账单 Rsp
type EcommerceFundFlowBillRsp struct {
	Code     int           `json:"-"`
	SignInfo *SignInfo     `json:"-"`
	Response *DownloadBill `json:"response,omitempty"`
	Error    string        `json:"-"`
}

// 合单查询订单 Rsp
type CombineQueryOrderRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *CombineQueryOrder `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 服务商查询订单 Rsp
type PartnerQueryOrderRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *PartnerQueryOrder `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 创建支付分订单 Rsp
type ScoreOrderCreateRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ScoreOrderCreate `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 查询支付分订单 Rsp
type ScoreOrderQueryRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *ScoreOrderQuery `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 取消支付分订单 Rsp
type ScoreOrderCancelRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ScoreOrderCancel `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 修改订单金额 Rsp
type ScoreOrderModifyRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ScoreOrderModify `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 完结支付分订单 Rsp
type ScoreOrderCompleteRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Response *ScoreOrderComplete `json:"response,omitempty"`
	Error    string              `json:"-"`
}

// 商户发起催收扣款 Rsp
type ScoreOrderPayRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *ScoreOrderPay `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 同步服务订单信息 Rsp
type ScoreOrderSyncRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *ScoreOrderSync `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 创单结单合并 Rsp
type ScoreDirectCompleteRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *ScoreDirectComplete `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 创单结单合并 Rsp
type ScorePermissionRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *ScorePermission `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 查询用户授权记录（授权协议号） Rsp
type ScorePermissionQueryRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *ScorePermissionQuery `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 查询用户授权记录（openid） Rsp
type ScorePermissionOpenidQueryRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *ScorePermissionOpenidQuery `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 查询投诉单列表 Rsp
type ComplaintListRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *ComplaintList `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 查询投诉单详情 Rsp
type ComplaintDetailRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *ComplaintDetail `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 查询投诉单详情 Rsp
type ComplaintNegotiationHistoryRsp struct {
	Code     int                          `json:"-"`
	SignInfo *SignInfo                    `json:"-"`
	Response *ComplaintNegotiationHistory `json:"response,omitempty"`
	Error    string                       `json:"-"`
}

// 创建、查询、更新投诉通知回调地址 Rsp
type ComplaintNotifyUrlRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Response *ComplaintNotifyUrl `json:"response,omitempty"`
	Error    string              `json:"-"`
}

// 商户上传反馈图片 Rsp
type MediaUploadRsp struct {
	Code     int          `json:"-"`
	SignInfo *SignInfo    `json:"-"`
	Response *MediaUpload `json:"response,omitempty"`
	Error    string       `json:"-"`
}

// 请求分账 Rsp
type ProfitShareOrderRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ProfitShareOrder `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 查询分账结果 Rsp
type ProfitShareOrderQueryRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *ProfitShareOrderQuery `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 请求分账回退 Rsp
type ProfitShareReturnRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *ProfitShareReturn `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 请求分账回退 Rsp
type ProfitShareReturnResultRsp struct {
	Code     int                      `json:"-"`
	SignInfo *SignInfo                `json:"-"`
	Response *ProfitShareReturnResult `json:"response,omitempty"`
	Error    string                   `json:"-"`
}

// 解冻剩余资金 Rsp
type ProfitShareOrderUnfreezeRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *ProfitShareOrderUnfreeze `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

// 查询剩余待分金额 Rsp
type ProfitShareUnsplitAmountRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *ProfitShareUnsplitAmount `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

// 添加分账接收方 Rsp
type ProfitShareAddReceiverRsp struct {
	Code     int                     `json:"-"`
	SignInfo *SignInfo               `json:"-"`
	Response *ProfitShareAddReceiver `json:"response,omitempty"`
	Error    string                  `json:"-"`
}

// 删除分账接收方 Rsp
type ProfitShareDeleteReceiverRsp struct {
	Code     int                        `json:"-"`
	SignInfo *SignInfo                  `json:"-"`
	Response *ProfitShareDeleteReceiver `json:"response,omitempty"`
	Error    string                     `json:"-"`
}

type ProfitShareMerchantConfigsRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *ProfitShareMerchantConfigs `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 申请分账账单 Rsp
type ProfitShareBillsRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *ProfitShareBills `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 预受理领卡请求 Rsp
type DiscountCardApplyRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *DiscountCardApply `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 查询先享卡订单 Rsp
type DiscountCardQueryRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *DiscountCardQuery `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 服务人员注册 Rsp
type SmartGuideRegRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *SmartGuideReg `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 服务人员查询 Rsp
type SmartGuideQueryRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *SmartGuideQuery `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 商圈积分授权查询 Rsp
type BusinessAuthPointsQueryRsp struct {
	Code     int                      `json:"-"`
	SignInfo *SignInfo                `json:"-"`
	Response *BusinessAuthPointsQuery `json:"response,omitempty"`
	Error    string                   `json:"-"`
}

// 发起批量转账 Rsp
type TransferRsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Response *Transfer `json:"response,omitempty"`
	Error    string    `json:"-"`
}

// 微信批次单号查询批次单 Rsp
type TransferQueryRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *TransferQuery `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 微信批次单号查询批次单（服务商） Rsp
type PartnerTransferQueryRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *PartnerTransferQuery `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 微信明细单号查询明细单 Rsp
type TransferDetailRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *TransferDetailQuery `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 微信明细单号查询明细单（服务商） Rsp
type PartnerTransferDetailRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *PartnerTransferDetail `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 商家批次单号查询批次单 Rsp
type TransferMerchantQueryRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *TransferMerchantQuery `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 商家批次单号查询批次单（服务商） Rsp
type PartnerTransferMerchantQueryRsp struct {
	Code     int                           `json:"-"`
	SignInfo *SignInfo                     `json:"-"`
	Response *PartnerTransferMerchantQuery `json:"response,omitempty"`
	Error    string                        `json:"-"`
}

// 商家明细单号查询明细单 Rsp
type TransferMerchantDetailRsp struct {
	Code     int                     `json:"-"`
	SignInfo *SignInfo               `json:"-"`
	Response *TransferMerchantDetail `json:"response,omitempty"`
	Error    string                  `json:"-"`
}

// 商家明细单号查询明细单（服务商） Rsp
type PartnerTransferMerchantDetailRsp struct {
	Code     int                            `json:"-"`
	SignInfo *SignInfo                      `json:"-"`
	Response *PartnerTransferMerchantDetail `json:"response,omitempty"`
	Error    string                         `json:"-"`
}

// 转账电子回单申请受理 Rsp
type TransferReceiptRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *TransferReceipt `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 查询转账电子回单 Rsp
type TransferReceiptQueryRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *TransferReceiptQuery `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 转账明细电子回单受理 Rsp
type TransferDetailReceiptRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *TransferDetailReceipt `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 查询转账明细电子回单受理结果 Rsp
type TransferDetailReceiptQueryRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *TransferDetailReceiptQuery `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 查询特约商户账户实时余额 Rsp
type EcommerceBalanceRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *EcommerceBalance `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 查询账户实时余额 Rsp
type MerchantBalanceRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *MerchantBalance `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 商户银行来账查询 Rsp
type MerchantIncomeRecordRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *MerchantIncomeRecord `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 特约商户银行来账查询 Rsp
type PartnerIncomeRecordRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *PartnerIncomeRecord `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 特约商户进件提交申请单 Rsp
type Apply4SubSubmitRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *Apply4SubSubmit `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 特约商户进件申请单查询 Rsp
type Apply4SubQueryRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *Apply4SubQuery `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 特约商户查询结算账号 Rsp
type Apply4SubQuerySettlementRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *Apply4SubQuerySettlement `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

// 创建代金券批次 Rsp
type FavorBatchCreateRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *FavorBatchCreate `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 发放代金券批次 Rsp
type FavorBatchGrantRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *FavorBatchGrant `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 激活代金券批次 Rsp
type FavorBatchStartRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *FavorBatchStart `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 条件查询批次列表 Rsp
type FavorBatchListRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *FavorBatchList `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 查询批次详情 Rsp
type FavorBatchDetailRsp struct {
	Code     int         `json:"-"`
	SignInfo *SignInfo   `json:"-"`
	Response *FavorBatch `json:"response,omitempty"`
	Error    string      `json:"-"`
}

// 查询批次详情 Rsp
type FavorDetailRsp struct {
	Code     int          `json:"-"`
	SignInfo *SignInfo    `json:"-"`
	Response *FavorDetail `json:"response,omitempty"`
	Error    string       `json:"-"`
}

// 查询代金券可用商户 Rsp
type FavorMerchantRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *FavorMerchant `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 查询代金券可用单品 Rsp
type FavorItemsRsp struct {
	Code     int         `json:"-"`
	SignInfo *SignInfo   `json:"-"`
	Response *FavorItems `json:"response,omitempty"`
	Error    string      `json:"-"`
}

// 根据商户号查用户的券 Rsp
type FavorUserCouponsRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *FavorUserCoupons `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 下载批次核销明细 Rsp
type FavorUseFlowDownloadRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *FavorFlowDownload `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 下载批次退款明细 Rsp
type FavorRefundFlowDownloadRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *FavorFlowDownload `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 设置消息通知地址 Rsp
type FavorCallbackUrlSetRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *FavorCallbackUrl `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 暂停代金券批次 Rsp
type FavorBatchPauseRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *FavorBatchPause `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 重启代金券批次 Rsp
type FavorBatchRestartRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *FavorBatchRestart `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 创建商家券 Rsp
type BusiFavorCreateRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *FavorBatchCreate `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 查询商家券详情 Rsp
type BusiFavorBatchDetailRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *BusiFavorBatchDetail `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 核销用户券 Rsp
type BusiFavorUseRsp struct {
	Code     int           `json:"-"`
	SignInfo *SignInfo     `json:"-"`
	Response *BusiFavorUse `json:"response,omitempty"`
	Error    string        `json:"-"`
}

// 核销用户券 Rsp
type BusiFavorUserCouponsRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *BusiFavorUserCoupons `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 查询用户单张券详情 Rsp
type BusiFavorUserCouponDetailRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *BusiUserCoupon `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 上传预存code Rsp
type BusiFavorCodeUploadRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *BusiFavorCodeUpload `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 设置商家券事件通知地址 Rsp
type BusiFavorCallbackUrlSetRsp struct {
	Code     int                      `json:"-"`
	SignInfo *SignInfo                `json:"-"`
	Response *BusiFavorCallbackUrlSet `json:"response,omitempty"`
	Error    string                   `json:"-"`
}

// 查询商家券事件通知地址 Rsp
type BusiFavorCallbackUrlRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *BusiFavorCallbackUrl `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 关联订单信息 Rsp
type BusiFavorAssociateRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Response *BusiFavorAssociate `json:"response,omitempty"`
	Error    string              `json:"-"`
}

// 取消关联订单信息 Rsp
type BusiFavorDisassociateRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *BusiFavorDisassociate `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 修改批次预算 Rsp
type BusiFavorBatchUpdateRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *BusiFavorBatchUpdate `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 发放消费卡 Rsp
type BusiFavorSendRsp struct {
	Code     int            `json:"-"`
	SignInfo *SignInfo      `json:"-"`
	Response *BusiFavorSend `json:"response,omitempty"`
	Error    string         `json:"-"`
}

// 申请退券 Rsp
type BusiFavorReturnRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *BusiFavorReturn `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 使券失效 Rsp
type BusiFavorDeactivateRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *BusiFavorDeactivate `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 营销补差付款 Rsp
type BusiFavorSubsidyPayRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *BusiFavorSubsidyPay `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 查询营销补差付款单详情 Rsp
type BusiFavorSubsidyPayDetailRsp struct {
	Code     int                  `json:"-"`
	SignInfo *SignInfo            `json:"-"`
	Response *BusiFavorSubsidyPay `json:"response,omitempty"`
	Error    string               `json:"-"`
}

// 商户上传(营销专用)反馈图片 Rsp
type MarketMediaUploadRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *MarketMediaUpload `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 建立合作关系 Rsp
type PartnershipsBuildRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *PartnershipsBuild `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 终止合作关系 Rsp
type PartnershipsTerminateRsp struct {
	Code     int                    `json:"-"`
	SignInfo *SignInfo              `json:"-"`
	Response *PartnershipsTerminate `json:"response,omitempty"`
	Error    string                 `json:"-"`
}

// 查询合作关系列表 Rsp
type PartnershipsListRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *PartnershipsList `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// 二级商户进件 Rsp
type EcommerceApplyRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *EcommerceApply `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 查询申请状态 Rsp
type EcommerceApplyStatusRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *EcommerceApplyStatus `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 点金计划管理 Rsp
type GoldPlanManageRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *GoldPlanManage `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 特约商户余额提现 Rsp
type WithdrawRsp struct {
	Code     int       `json:"-"`
	SignInfo *SignInfo `json:"-"`
	Response *Withdraw `json:"response,omitempty"`
	Error    string    `json:"-"`
}

// 查询特约商户提现状态 Rsp
type WithdrawStatusRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *WithdrawStatus `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 电商平台预约提现 Rsp
type EcommerceWithdrawRsp struct {
	Code     int                `json:"-"`
	SignInfo *SignInfo          `json:"-"`
	Response *EcommerceWithdraw `json:"response,omitempty"`
	Error    string             `json:"-"`
}

// 电商平台查询预约提现状态 Rsp
type EcommerceWithdrawStatusRsp struct {
	Code     int                      `json:"-"`
	SignInfo *SignInfo                `json:"-"`
	Response *EcommerceWithdrawStatus `json:"response,omitempty"`
	Error    string                   `json:"-"`
}

// 请求分账 Rsp
type EcommerceProfitShareRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *EcommerceProfitShare `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 查询分账结果 Rsp
type EcommerceProfitShareQueryRsp struct {
	Code     int                        `json:"-"`
	SignInfo *SignInfo                  `json:"-"`
	Response *EcommerceProfitShareQuery `json:"response,omitempty"`
	Error    string                     `json:"-"`
}

// 请求分账回退 Rsp
type EcommerceProfitShareReturnRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *EcommerceProfitShareReturn `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 查询分账回退结果 Rsp
type EcommerceProfitShareReturnResultRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *EcommerceProfitShareReturn `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 完结分账 Rsp
type EcommerceProfitShareFinishRsp struct {
	Code     int                         `json:"-"`
	SignInfo *SignInfo                   `json:"-"`
	Response *EcommerceProfitShareFinish `json:"response,omitempty"`
	Error    string                      `json:"-"`
}

// 查询订单剩余待分金额 Rsp
type EcommerceProfitShareUnsplitAmountRsp struct {
	Code     int                                `json:"-"`
	SignInfo *SignInfo                          `json:"-"`
	Response *EcommerceProfitShareUnsplitAmount `json:"response,omitempty"`
	Error    string                             `json:"-"`
}

// 添加分账接收方 Rsp
type EcommerceProfitShareAddReceiverRsp struct {
	Code     int                           `json:"-"`
	SignInfo *SignInfo                     `json:"-"`
	Response *EcommerceProfitShareReceiver `json:"response,omitempty"`
	Error    string                        `json:"-"`
}

// 删除分账接收方 Rsp
type EcommerceProfitShareDeleteReceiverRsp struct {
	Code     int                           `json:"-"`
	SignInfo *SignInfo                     `json:"-"`
	Response *EcommerceProfitShareReceiver `json:"response,omitempty"`
	Error    string                        `json:"-"`
}

// 请求补差 Rsp
type EcommerceSubsidiesRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Response *EcommerceSubsidies `json:"response,omitempty"`
	Error    string              `json:"-"`
}

// 请求补差回退 Rsp
type EcommerceSubsidiesReturnRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *EcommerceSubsidiesReturn `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

// 取消补差 Rsp
type EcommerceSubsidiesCancelRsp struct {
	Code     int                       `json:"-"`
	SignInfo *SignInfo                 `json:"-"`
	Response *EcommerceSubsidiesCancel `json:"response,omitempty"`
	Error    string                    `json:"-"`
}

// 申请退款 Rsp
type EcommerceRefundRsp struct {
	Code     int              `json:"-"`
	SignInfo *SignInfo        `json:"-"`
	Response *EcommerceRefund `json:"response,omitempty"`
	Error    string           `json:"-"`
}

// 申请退款 Rsp
type EcommerceRefundQueryRsp struct {
	Code     int                   `json:"-"`
	SignInfo *SignInfo             `json:"-"`
	Response *EcommerceRefundQuery `json:"response,omitempty"`
	Error    string                `json:"-"`
}

// 垫付退款回补 Rsp
type EcommerceRefundAdvanceRsp struct {
	Code     int                     `json:"-"`
	SignInfo *SignInfo               `json:"-"`
	Response *EcommerceRefundAdvance `json:"response,omitempty"`
	Error    string                  `json:"-"`
}

// 获取对私银行卡号开户银行 Rsp
type BankSearchBankRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *BankSearchBank `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 查询支持个人业务的银行列表 Rsp
type BankSearchPersonalListRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *BankSearchList `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 查询支持对公业务的银行列表 Rsp
type BankSearchCorporateListRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *BankSearchList `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 查询省份列表 Rsp
type BankSearchProvinceListRsp struct {
	Code     int                 `json:"-"`
	SignInfo *SignInfo           `json:"-"`
	Response *BankSearchProvince `json:"response,omitempty"`
	Error    string              `json:"-"`
}

// 查询城市列表 Rsp
type BankSearchCityListRsp struct {
	Code     int             `json:"-"`
	SignInfo *SignInfo       `json:"-"`
	Response *BankSearchCity `json:"response,omitempty"`
	Error    string          `json:"-"`
}

// 查询支行列表 Rsp
type BankSearchBranchListRsp struct {
	Code     int               `json:"-"`
	SignInfo *SignInfo         `json:"-"`
	Response *BankSearchBranch `json:"response,omitempty"`
	Error    string            `json:"-"`
}

// ==================================分割==================================

type JSAPIPayParams struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

type AppPayParams struct {
	Appid     string `json:"appid"`
	Partnerid string `json:"partnerid"`
	Prepayid  string `json:"prepayid"`
	Package   string `json:"package"`
	Noncestr  string `json:"noncestr"`
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
}

type AppletParams struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

type AppletScoreExtraData struct {
	MchId     string `json:"mch_id"`
	Package   string `json:"package"`
	TimeStamp string `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	SignType  string `json:"sign_type"`
	Sign      string `json:"sign"`
}

// ==================================分割==================================

type SignInfo struct {
	HeaderTimestamp string `json:"Wechatpay-Timestamp"`
	HeaderNonce     string `json:"Wechatpay-Nonce"`
	HeaderSignature string `json:"Wechatpay-Signature"`
	HeaderSerial    string `json:"Wechatpay-Serial"`
	SignBody        string `json:"sign_body"`
}

type PlatformCertItem struct {
	EffectiveTime string `json:"effective_time"`
	ExpireTime    string `json:"expire_time"`
	PublicKey     string `json:"public_key"`
	SerialNo      string `json:"serial_no"`
}

type PlatformCert struct {
	Data []*certData `json:"data"`
}

type certData struct {
	EffectiveTime      string       `json:"effective_time"`
	EncryptCertificate *encryptCert `json:"encrypt_certificate"`
	ExpireTime         string       `json:"expire_time"`
	SerialNo           string       `json:"serial_no"`
}

type encryptCert struct {
	Algorithm      string `json:"algorithm"`
	AssociatedData string `json:"associated_data"`
	Ciphertext     string `json:"ciphertext"`
	Nonce          string `json:"nonce"`
}

type Prepay struct {
	PrepayId string `json:"prepay_id"` // 预支付交易会话标识。用于后续接口调用中使用，该值有效期为2小时
}

type Native struct {
	CodeUrl string `json:"code_url"` // 此URL用于生成支付二维码，然后提供给用户扫码支付
}

type H5Url struct {
	H5Url string `json:"h5_url"` // h5_url为拉起微信支付收银台的中间页面，可通过访问该url来拉起微信客户端，完成支付，h5_url的有效期为5分钟
}

type Payer struct {
	Openid string `json:"openid"` // 用户在直连商户appid下的唯一标识
}

type Amount struct {
	Total         int    `json:"total,omitempty"`          // 订单总金额，单位为分
	PayerTotal    int    `json:"payer_total,omitempty"`    // 用户支付金额，单位为分
	Currency      string `json:"currency,omitempty"`       // CNY：人民币，境内商户号仅支持人民币
	PayerCurrency string `json:"payer_currency,omitempty"` // 用户支付币种
}

type RefundAmount struct {
	Total       int `json:"total,omitempty"`       // 订单总金额，单位为分，只能为整数
	Refund      int `json:"refund,omitempty"`      // 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额，如果有使用券，后台会按比例退
	PayerTotal  int `json:"payer_total,omitempty"` // 用户实际支付金额，单位为分，只能为整数
	PayerRefund int `json:"payer_refund"`          // 退款给用户的金额，不包含所有优惠券金额
}

type CombineAmount struct {
	TotalAmount   int    `json:"total_amount,omitempty"`   // 订单总金额，单位为分
	Currency      string `json:"currency,omitempty"`       // 标价币种：符合ISO 4217标准的三位字母代码，人民币：CNY
	PayerAmount   int    `json:"payer_amount"`             // 订单现金支付金额
	PayerCurrency string `json:"payer_currency,omitempty"` // 现金支付币种：符合ISO 4217标准的三位字母代码，默认人民币：CNY
}

type SceneInfo struct {
	DeviceId string `json:"device_id,omitempty"` // 商户端设备号（发起扣款请求的商户服务器设备号）
}

type PromotionDetail struct {
	Amount              int            `json:"amount"`                         // 优惠券面额
	CouponId            string         `json:"coupon_id"`                      // 券Id
	Name                string         `json:"name,omitempty"`                 // 优惠名称
	Scope               string         `json:"scope,omitempty"`                // 优惠范围：GLOBAL：全场代金券, SINGLE：单品优惠
	Type                string         `json:"type,omitempty"`                 // 优惠类型：CASH：充值, NOCASH：预充值
	StockId             string         `json:"stock_id,omitempty"`             // 活动Id
	WechatpayContribute int            `json:"wechatpay_contribute,omitempty"` // 微信出资，单位为分
	MerchantContribute  int            `json:"merchant_contribute,omitempty"`  // 商户出资，单位为分
	OtherContribute     int            `json:"other_contribute,omitempty"`     // 其他出资，单位为分
	Currency            string         `json:"currency,omitempty"`             // CNY：人民币，境内商户号仅支持人民币
	GoodsDetail         []*GoodsDetail `json:"goods_detail,omitempty"`         // 单品列表信息
}

type GoodsDetail struct {
	GoodsId         string `json:"goods_id"`                    // 商品编码
	Quantity        int    `json:"quantity"`                    // 用户购买的数量
	UnitPrice       int    `json:"unit_price"`                  // 商品单价，单位为分
	DiscountAmount  int    `json:"discount_amount"`             // 商品优惠金额
	GoodsRemark     string `json:"goods_remark,omitempty"`      // 商品备注信息
	MerchantGoodsId string `json:"merchant_goods_id,omitempty"` // 商户侧商品编码，服务商模式下无此字段
}

type QueryOrder struct {
	Appid           string             `json:"appid"`                      // 直连商户申请的公众号或移动应用appid
	Mchid           string             `json:"mchid"`                      // 直连商户的商户号，由微信支付生成并下发
	OutTradeNo      string             `json:"out_trade_no"`               // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	TransactionId   string             `json:"transaction_id"`             // 微信支付系统生成的订单号
	TradeType       string             `json:"trade_type"`                 // 交易类型，枚举值：JSAPI：公众号支付, NATIVE：扫码支付, APP：APP支付, MICROPAY：付款码支付, MWEB：H5支付, FACEPAY：刷脸支付
	TradeState      string             `json:"trade_state"`                // 交易状态，枚举值：SUCCESS：支付成功, REFUND：转入退款, NOTPAY：未支付, CLOSED：已关闭, REVOKED：已撤销（付款码支付）, USERPAYING：用户支付中（付款码支付）, PAYERROR：支付失败(其他原因，如银行返回失败)
	TradeStateDesc  string             `json:"trade_state_desc"`           // 交易状态描述
	BankType        string             `json:"bank_type,omitempty"`        // 银行类型，采用字符串类型的银行标识
	Attach          string             `json:"attach"`                     // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用
	SuccessTime     string             `json:"success_time,omitempty"`     // 支付完成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒
	Payer           *Payer             `json:"payer"`                      // 支付者信息
	Amount          *Amount            `json:"amount,omitempty"`           // 订单金额信息，当支付成功时返回该字段
	SceneInfo       *SceneInfo         `json:"scene_info,omitempty"`       // 支付场景描述
	PromotionDetail []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能，享受优惠时返回该字段
}

type SubOrders struct {
	Mchid         string         `json:"mchid"`               // 子单发起方商户号，必须与发起方Appid有绑定关系
	TradeType     string         `json:"trade_type"`          // 交易类型，枚举值：NATIVE：扫码支付，JSAPI：公众号支付，APP：APP支付，MWEB：H5支付
	TradeState    string         `json:"trade_state"`         // 交易状态，枚举值：SUCCESS：支付成功, REFUND：转入退款, NOTPAY：未支付, CLOSED：已关闭, REVOKED：已撤销（付款码支付）, USERPAYING：用户支付中（付款码支付）, PAYERROR：支付失败(其他原因，如银行返回失败)
	BankType      string         `json:"bank_type,omitempty"` // 银行类型，采用字符串类型的银行标识
	Attach        string         `json:"attach"`              // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用
	SuccessTime   string         `json:"success_time"`        // 支付完成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒
	TransactionId string         `json:"transaction_id"`      // 微信支付系统生成的订单号
	OutTradeNo    string         `json:"out_trade_no"`        // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	SubMchid      string         `json:"sub_mchid"`           // 二级商户商户号，由微信支付生成并下发。服务商子商户的商户号，被合单方。直连商户不用传二级商户号
	Amount        *CombineAmount `json:"amount"`              // 订单金额信息，当支付成功时返回该字段

}

type CombineQueryOrder struct {
	CombineAppid      string       `json:"combine_appid"`        // 合单发起方的appid
	CombineMchid      string       `json:"combine_mchid"`        // 合单发起方商户号
	CombineOutTradeNo string       `json:"combine_out_trade_no"` // 合单支付总订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一
	SceneInfo         *SceneInfo   `json:"scene_info,omitempty"` // 支付场景描述
	SubOrders         []*SubOrders `json:"sub_orders"`           // 最多支持子单条数：50
	CombinePayerInfo  *Payer       `json:"combine_payer_info"`   // 支付者信息
}

type TradeBill struct {
	HashType    string `json:"hash_type"`
	HashValue   string `json:"hash_value"`
	DownloadUrl string `json:"download_url"`
}

type BillDetail struct {
	BillSequence int    `json:"bill_sequence"` // 商户将多个文件按账单文件序号的顺序合并为完整的资金账单文件，起始值为1
	HashType     string `json:"hash_type"`
	HashValue    string `json:"hash_value"`
	DownloadUrl  string `json:"download_url"`
	EncryptKey   string `json:"encrypt_key"` // 加密账单文件使用的加密密钥。密钥用商户证书的公钥进行加密，然后进行Base64编码
	Nonce        string `json:"nonce"`       // 加密账单文件使用的随机字符串
}

type DownloadBill struct {
	DownloadBillCount int           `json:"download_bill_count"`
	DownloadBillList  []*BillDetail `json:"download_bill_list"`
}

type RefundOrderResponse struct {
	RefundId            string                        `json:"refund_id"`             // 微信支付退款号
	OutRefundNo         string                        `json:"out_refund_no"`         // 商户退款单号
	TransactionId       string                        `json:"transaction_id"`        // 微信支付系统生成的订单号
	OutTradeNo          string                        `json:"out_trade_no"`          // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	Channel             string                        `json:"channel"`               // 退款渠道
	UserReceivedAccount string                        `json:"user_received_account"` // 退款入账账户
	SuccessTime         string                        `json:"success_time"`          // 退款成功时间
	CreateTime          string                        `json:"create_time"`           // 退款创建时间
	Status              string                        `json:"status"`                // 退款状态
	FundsAccount        string                        `json:"funds_account"`         // 资金账户
	Amount              *RefundQueryAmount            `json:"amount"`                // 金额信息
	PromotionDetail     []*RefundQueryPromotionDetail `json:"promotion_detail"`      // 优惠退款信息
}

type RefundQueryResponse struct {
	RefundId            string                        `json:"refund_id"`             // 微信支付退款号
	OutRefundNo         string                        `json:"out_refund_no"`         // 商户退款单号
	TransactionId       string                        `json:"transaction_id"`        // 微信支付系统生成的订单号
	OutTradeNo          string                        `json:"out_trade_no"`          // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	Channel             string                        `json:"channel"`               // 退款渠道
	UserReceivedAccount string                        `json:"user_received_account"` // 退款入账账户
	SuccessTime         string                        `json:"success_time"`          // 退款成功时间
	CreateTime          string                        `json:"create_time"`           // 退款创建时间
	Status              string                        `json:"status"`                // 退款状态
	FundsAccount        string                        `json:"funds_account"`         // 资金账户
	Amount              *RefundQueryAmount            `json:"amount"`                // 金额信息
	PromotionDetail     []*RefundQueryPromotionDetail `json:"promotion_detail"`      // 优惠退款信息
}

type RefundQueryAmount struct {
	Total            int    `json:"total"`             // 订单总金额，单位为分
	Refund           int    `json:"refund"`            // 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额。
	PayerTotal       int    `json:"payer_total"`       // 用户支付金额，单位为分
	PayerRefund      int    `json:"payer_refund"`      // 用户退款金额，不包含所有优惠券金额
	SettlementRefund int    `json:"settlement_refund"` // 应结退款金额，去掉非充值代金券退款金额后的退款金额，单位为分
	DiscountRefund   int    `json:"discount_refund"`   // 优惠退款金额
	Currency         string `json:"currency"`          // CNY：人民币，境内商户号仅支持人民币
}

type RefundQueryPromotionDetail struct {
	PromotionId  string                    `json:"promotion_id"`           // 券Id，券或立减金额
	Scope        string                    `json:"scope"`                  // 优惠范围，GLOBAL：全场代金券，SINGLE：单品优惠
	Type         string                    `json:"type"`                   // 优惠类型，COUPON：代金券，DISCOUNT：优惠券
	Amount       int                       `json:"amount"`                 // 优惠券面额，用户享受优惠的金额（优惠券面额=微信出资金额+商家出资金额+其他出资方金额），单位为分
	RefundAmount int                       `json:"refund_amount"`          // 优惠退款金额，单位为分
	GoodsDetail  []*RefundQueryGoodsDetail `json:"goods_detail,omitempty"` // 商品列表，优惠商品发送退款时返回商品信息
}

type RefundQueryGoodsDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`            // 商户侧商品编码
	WechatpayGoodsId string `json:"wechatpay_goods_id,omitempty"` // 微信侧商品编码
	GoodsName        string `json:"goods_name,omitempty"`         // 商品名称
	UnitPrice        int    `json:"unit_price"`                   // 商品单价金额
	RefundAmount     int    `json:"refund_amount"`                // 商品退款金额
	RefundQuantity   int    `json:"refund_quantity"`              // 商品退货数量
}

type PartnerQueryOrder struct {
	SpAppid         string             `json:"sp_appid"`                   // 服务商申请的公众号或移动应用appid。
	SpMchid         string             `json:"sp_mchid"`                   // 服务商户号，由微信支付生成并下发
	SubAppid        string             `json:"sub_appid"`                  // 子商户申请的公众号或移动应用appid。如果返回sub_appid，那么sub_openid一定会返回。
	SubMchid        string             `json:"sub_mchid"`                  // 子商户的商户号，有微信支付生成并下发。
	OutTradeNo      string             `json:"out_trade_no"`               // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	TransactionId   string             `json:"transaction_id"`             // 微信支付系统生成的订单号
	TradeType       string             `json:"trade_type"`                 // 交易类型，枚举值：JSAPI：公众号支付, NATIVE：扫码支付, APP：APP支付, MICROPAY：付款码支付, MWEB：H5支付, FACEPAY：刷脸支付
	TradeState      string             `json:"trade_state"`                // 交易状态，枚举值：SUCCESS：支付成功, REFUND：转入退款, NOTPAY：未支付, CLOSED：已关闭, REVOKED：已撤销（付款码支付）, USERPAYING：用户支付中（付款码支付）, PAYERROR：支付失败(其他原因，如银行返回失败)
	TradeStateDesc  string             `json:"trade_state_desc"`           // 交易状态描述
	BankType        string             `json:"bank_type,omitempty"`        // 银行类型，采用字符串类型的银行标识
	Attach          string             `json:"attach"`                     // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用
	SuccessTime     string             `json:"success_time,omitempty"`     // 支付完成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒
	Payer           *PartnerPayer      `json:"payer"`                      // 支付者信息
	Amount          *Amount            `json:"amount,omitempty"`           // 订单金额信息，当支付成功时返回该字段
	SceneInfo       *SceneInfo         `json:"scene_info,omitempty"`       // 支付场景描述
	PromotionDetail []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能，享受优惠时返回该字段
}

type PartnerPayer struct {
	SpOpenid  string `json:"sp_openid"`  // 用户在服务商appid下的唯一标识。
	SubOpenid string `json:"sub_openid"` // 用户在子商户appid下的唯一标识。 如果返回sub_appid，那么sub_openid一定会返回。
}

type ScoreOrderCreate struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	PostPayments        []*PostPayments  `json:"post_payments,omitempty"`     // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund"`                   // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range"`                  // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url,omitempty"`        // 商户接收用户确认订单或扣款成功回调通知的地址。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	Package             string           `json:"package"`                     // 用户跳转到微信侧小程序订单数据，需确认模式特有API中调起支付分-确认订单传入。该数据一小时内有效。
}

type PostPayments struct {
	Name        string `json:"name"`        // 付费项目名称
	Amount      int    `json:"amount"`      // 此付费项目总金额，大于等于0，单位为分，等于0时代表不需要扣费，只能为整数
	Description string `json:"description"` // 描述计费规则，不超过30个字符，超出报错处理。
	Count       int    `json:"count"`       // 付费项目的数量。
}

type PostDiscounts struct {
	Name        string `json:"name"`        // 优惠名称说明。
	Description string `json:"description"` // 优惠使用条件说明。
	Amount      int    `json:"amount"`      // 优惠金额，只能为整数
	Count       int    `json:"count"`       // 优惠的数量。
}

type RiskFund struct {
	Name        string `json:"name"`                  // 风险金名称。DEPOSIT：押金，ADVANCE：预付款，CASH_DEPOSIT：保证金，ESTIMATE_ORDER_COST：预估订单费用
	Description string `json:"description,omitempty"` // 风险说明
	Amount      int    `json:"amount"`                // 风险金额
}

type TimeRange struct {
	StartTime       string `json:"start_time"`                  // 服务开始时间，20091225091010
	StartTimeRemark string `json:"start_time_remark,omitempty"` // 服务开始时间备注
	EndTime         string `json:"end_time,omitempty"`          // 预计服务结束时间，20091225121010
	EndTimeRemark   string `json:"end_time_remark,omitempty"`   // 预计服务结束时间备注
}

type Location struct {
	StartLocation string `json:"start_location,omitempty"` // 服务开始地点
	EndLocation   string `json:"end_location,omitempty"`   // 服务结束位置
}

type ScoreOrderQuery struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	TotalAmount         int              `json:"total_amount,omitempty"`      // 总金额，大于等于0的数字，单位为分，只能为整数
	PostPayments        []*PostPayments  `json:"post_payments,omitempty"`     // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund"`                   // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range"`                  // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url"`                  // 商户接收用户确认订单或扣款成功回调通知的地址。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	NeedCollection      bool             `json:"need_collection,omitempty"`   // 是否需要收款
	Collection          *Collection      `json:"collection,omitempty"`        // 收款信息
	Openid              string           `json:"openid,omitempty"`            // 微信用户在商户对应appid下的唯一标识
}

type Collection struct {
	State        string     `json:"state"`                   // 收款状态，USER_PAYING：待支付，USER_PAID：已支付
	TotalAmount  int        `json:"total_amount,omitempty"`  // 总金额，大于等于0的数字，单位为分，只能为整数
	PayingAmount int        `json:"paying_amount,omitempty"` // 等待用户付款金额，只能为整数
	PaidAmount   int        `json:"paid_amount,omitempty"`   // 用户已付款的金额，只能为整数
	Details      []*Details `json:"details,omitempty"`       // 收款明细列表
}

type Details struct {
	Seq             int                `json:"seq,omitempty"`              // 收款序号
	Amount          int                `json:"amount,omitempty"`           // 单笔收款动作的金额，只能为整数
	PaidType        string             `json:"paid_type,omitempty"`        // 收款成功渠道，NEWTON：微信支付分，MCH：商户渠道
	PaidTime        string             `json:"paid_time,omitempty"`        // 支付成功时间，支持两种格式:yyyyMMddHHmmss和yyyyMMdd
	TransactionId   string             `json:"transaction_id,omitempty"`   // 结单交易单号，等于普通支付接口中的transaction_id
	PromotionDetail []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能，享受优惠时返回该字段
}

type ScoreOrderCancel struct {
	Appid      string `json:"appid"`        // 调用接口提交的公众账号Id。
	Mchid      string `json:"mchid"`        // 调用接口提交的商户号。
	ServiceId  string `json:"service_id"`   // 调用该接口提交的服务Id。
	OutOrderNo string `json:"out_order_no"` // 调用接口提交的商户服务订单号。
	OrderId    string `json:"order_id"`     // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
}

type ScoreOrderModify struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	TotalAmount         int              `json:"total_amount,omitempty"`      // 总金额，大于等于0的数字，单位为分，只能为整数
	PostPayments        []*PostPayments  `json:"post_payments,omitempty"`     // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund"`                   // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range"`                  // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url,omitempty"`        // 商户接收用户确认订单或扣款成功回调通知的地址。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	NeedCollection      bool             `json:"need_collection,omitempty"`   // 是否需要收款
	Collection          *Collection      `json:"collection,omitempty"`        // 收款信息
}

type ScoreOrderComplete struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	TotalAmount         int              `json:"total_amount"`                // 总金额，大于等于0的数字，单位为分，只能为整数
	PostPayments        []*PostPayments  `json:"post_payments"`               // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund"`                   // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range,omitempty"`        // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	NeedCollection      bool             `json:"need_collection,omitempty"`   // 是否需要收款
}

type ScoreOrderPay struct {
	Appid      string `json:"appid"`        // 调用接口提交的公众账号Id。
	Mchid      string `json:"mchid"`        // 调用接口提交的商户号。
	ServiceId  string `json:"service_id"`   // 调用该接口提交的服务Id。
	OutOrderNo string `json:"out_order_no"` // 调用接口提交的商户服务订单号。
	OrderId    string `json:"order_id"`     // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
}

type ScoreOrderSync struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	TotalAmount         int              `json:"total_amount"`                // 总金额，大于等于0的数字，单位为分，只能为整数
	PostPayments        []*PostPayments  `json:"post_payments,omitempty"`     // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	RiskFund            *RiskFund        `json:"risk_fund,omitempty"`         // 订单风险金信息
	TimeRange           *TimeRange       `json:"time_range,omitempty"`        // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url,omitempty"`        // 商户接收用户确认订单或扣款成功回调通知的地址。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	NeedCollection      bool             `json:"need_collection,omitempty"`   // 是否需要收款
	Collection          *Collection      `json:"collection,omitempty"`        // 收款信息
	Openid              string           `json:"openid"`                      // 微信用户在商户对应appid下的唯一标识
}

type ScoreDirectComplete struct {
	Appid               string           `json:"appid"`                       // 调用接口提交的公众账号Id。
	Mchid               string           `json:"mchid"`                       // 调用接口提交的商户号。
	OutOrderNo          string           `json:"out_order_no"`                // 调用接口提交的商户服务订单号。
	ServiceId           string           `json:"service_id"`                  // 调用该接口提交的服务Id。
	OrderId             string           `json:"order_id"`                    // 微信支付服务订单号，每个微信支付服务订单号与商户号下对应的商户服务订单号一一对应。
	ServiceIntroduction string           `json:"service_introduction"`        // 服务信息，用于介绍本订单所提供的服务。
	State               string           `json:"state"`                       // 表示当前单据状态。枚举值：CREATED：商户已创建服务订单，DOING：服务订单进行中，DONE：服务订单完成，REVOKED：商户取消服务订单，EXPIRED：服务订单已失效
	StateDescription    string           `json:"state_description,omitempty"` // 对服务订单"进行中"状态的附加说明。USER_CONFIRM：用户确认，MCH_COMPLETE：商户完结
	PostPayments        []*PostPayments  `json:"post_payments"`               // 后付费项目列表，最多包含100条付费项目。 如果传入，用户侧则显示此参数。
	PostDiscounts       []*PostDiscounts `json:"post_discounts,omitempty"`    // 后付费商户优惠，最多包含30条付费项目。 如果传入，用户侧则显示此参数。
	TimeRange           *TimeRange       `json:"time_range"`                  // 服务时间范围
	Location            *Location        `json:"location,omitempty"`          // 服务位置信息
	Attach              string           `json:"attach,omitempty"`            // 商户数据包,可存放本订单所需信息，需要先urlencode后传入，总长度不大于256字符,超出报错处理。
	NotifyUrl           string           `json:"notify_url,omitempty"`        // 商户接收用户确认订单或扣款成功回调通知的地址。
	TotalAmount         int              `json:"total_amount"`                // 金额：数字，必须≥0（单位：分）
}

type ScorePermission struct {
	ApplyPermissionsToken string `json:"apply_permissions_token"` // 用于跳转到微信侧小程序授权数据,跳转到微信侧小程序传入，时效性为1小时
}

type ScorePermissionQuery struct {
	Appid                    string `json:"appid"`                               // 调用接口提交的公众账号Id。
	Mchid                    string `json:"mchid"`                               // 调用接口提交的商户号。
	ServiceId                string `json:"service_id"`                          // 调用该接口提交的服务Id。
	Openid                   string `json:"openid,omitempty"`                    // 微信用户在商户对应appid下的唯一标识
	AuthorizationCode        string `json:"authorization_code"`                  // 预授权成功时的授权协议号。
	AuthorizationState       string `json:"authorization_state"`                 // 标识用户授权服务情况：UNAVAILABLE：用户未授权服务，AVAILABLE：用户已授权服务，UNBINDUSER：未绑定用户（已经预授权但未完成正式授权）
	NotifyUrl                string `json:"notify_url,omitempty"`                // 商户接收用户确认订单或扣款成功回调通知的地址。
	CancelAuthorizationTime  string `json:"cancel_authorization_time,omitempty"` // 最近一次解除授权时间, 示例值：2015-05-20T13:29:35.120+08:00
	AuthorizationSuccessTime string `json:"authorization_success_time"`          // 最近一次授权成功时间, 示例值：2015-05-20T13:29:35.120+08:00
}

type ScorePermissionOpenidQuery struct {
	Appid                    string `json:"appid"`                               // 调用接口提交的公众账号Id。
	Mchid                    string `json:"mchid"`                               // 调用接口提交的商户号。
	ServiceId                string `json:"service_id"`                          // 调用该接口提交的服务Id。
	Openid                   string `json:"openid,omitempty"`                    // 微信用户在商户对应appid下的唯一标识
	AuthorizationCode        string `json:"authorization_code"`                  // 预授权成功时的授权协议号。
	AuthorizationState       string `json:"authorization_state"`                 // 标识用户授权服务情况：UNAVAILABLE：用户未授权服务，AVAILABLE：用户已授权服务，UNBINDUSER：未绑定用户（已经预授权但未完成正式授权）
	CancelAuthorizationTime  string `json:"cancel_authorization_time,omitempty"` // 最近一次解除授权时间, 示例值：2015-05-20T13:29:35.120+08:00
	AuthorizationSuccessTime string `json:"authorization_success_time"`          // 最近一次授权成功时间, 示例值：2015-05-20T13:29:35.120+08:00
}

type ComplaintList struct {
	Data       []*ComplaintListItem `json:"data,omitempty"`        // 用户投诉信息详情
	Limit      int                  `json:"limit"`                 // 设置该次请求返回的最大投诉条数，范围【1,50】
	Offset     int                  `json:"offset"`                // 该次请求的分页开始位置，从0开始计数，例如offset=10，表示从第11条记录开始返回。
	TotalCount int                  `json:"total_count,omitempty"` // 投诉总条数，当offset=0时返回
}

type ComplaintListItem struct {
	ComplaintId           string                `json:"complaint_id"`                   // 投诉单对应的投诉单号
	ComplaintTime         string                `json:"complaint_time"`                 // 投诉时间, 例如：2015-05-20T13:29:35.120+08:00表示北京时间2015年05月20日13点29分35秒
	ComplaintDetail       string                `json:"complaint_detail"`               // 投诉的具体描述
	ComplaintState        string                `json:"complaint_state"`                // 投诉单状态, PENDING：待处理, PROCESSING：处理中, PROCESSED：已处理完成
	PayerPhone            string                `json:"payer_phone,omitempty"`          // 投诉人联系方式。该字段已做加密处理
	ComplaintOrderInfo    []*ComplaintOrderInfo `json:"complaint_order_info,omitempty"` // 投诉单关联订单信息
	ComplaintFullRefunded bool                  `json:"complaint_full_refunded"`        // 投诉单下所有订单是否已全部全额退款
	IncomingUserResponse  bool                  `json:"incoming_user_response"`         // 投诉单是否有待回复的用户留言
	UserComplaintTimes    int                   `json:"user_complaint_times"`           // 用户投诉次数
}

type ComplaintOrderInfo struct {
	TransactionId string `json:"transaction_id"` // 投诉单关联的微信订单号
	OutTradeNo    string `json:"out_trade_no"`   // 投诉单关联的商户订单号
	Amount        int    `json:"amount"`         // 订单金额，单位（分）
}

type ComplaintDetail struct {
	ComplaintId           string                `json:"complaint_id"`                   // 投诉单对应的投诉单号
	ComplaintTime         string                `json:"complaint_time"`                 // 投诉时间, 例如：2015-05-20T13:29:35.120+08:00表示北京时间2015年05月20日13点29分35秒
	ComplaintDetail       string                `json:"complaint_detail"`               // 投诉的具体描述
	ComplaintedMchid      string                `json:"complainted_mchid,omitempty"`    // 投诉单对应的被诉商户号。
	ComplaintState        string                `json:"complaint_state"`                // 投诉单状态, PENDING：待处理, PROCESSING：处理中, PROCESSED：已处理完成
	PayerPhone            string                `json:"payer_phone,omitempty"`          // 投诉人联系方式。该字段已做加密处理
	PayerOpenid           string                `json:"payer_openid"`                   // 投诉人在商户appid下的唯一标识
	ComplaintOrderInfo    []*ComplaintOrderInfo `json:"complaint_order_info,omitempty"` // 投诉单关联订单信息
	ComplaintFullRefunded bool                  `json:"complaint_full_refunded"`        // 投诉单下所有订单是否已全部全额退款
	IncomingUserResponse  bool                  `json:"incoming_user_response"`         // 投诉单是否有待回复的用户留言
	UserComplaintTimes    int                   `json:"user_complaint_times"`           // 用户投诉次数
}

type ComplaintNegotiationHistory struct {
	Data       []*ComplaintNegotiationHistoryItem `json:"data,omitempty"`        // 投诉协商历史
	Limit      int                                `json:"limit"`                 // 设置该次请求返回的最大投诉条数，范围【1,50】
	Offset     int                                `json:"offset"`                // 该次请求的分页开始位置，从0开始计数，例如offset=10，表示从第11条记录开始返回。
	TotalCount int                                `json:"total_count,omitempty"` // 投诉总条数，当offset=0时返回
}

type ComplaintNegotiationHistoryItem struct {
	LogId          string   `json:"log_id"`          // 操作流水号
	Operator       string   `json:"operator"`        // 当前投诉协商记录的操作人
	OperateTime    string   `json:"operate_time"`    // 当前投诉协商记录的操作时间
	OperateType    string   `json:"operate_type"`    // 当前投诉协商记录的操作类型
	OperateDetails string   `json:"operate_details"` // 当前投诉协商记录的具体内容
	ImageList      []string `json:"image_list"`      // 当前投诉协商记录提交的图片凭证（url格式），最多返回4张图片，url有效时间为1小时。如未查询到协商历史图片凭证，则返回空数组。
}

type ComplaintNotifyUrl struct {
	Mchid string `json:"mchid"` // 返回创建回调地址的商户号，由微信支付生成并下发。
	Url   string `json:"url"`   // 通知地址，仅支持https。
}

type MediaUpload struct {
	MediaId string `json:"media_id"` // 微信返回的媒体文件标识Id。
}

type ProfitShareOrder struct {
	SubMchid      string                   `json:"sub_mchid"`           // 二级商户号
	TransactionId string                   `json:"transaction_id"`      // 微信订单号
	OutOrderNo    string                   `json:"out_order_no"`        // 商户分账单号
	OrderId       string                   `json:"order_id"`            // 微信分账单号
	State         string                   `json:"state"`               // 分账单状态（每个接收方的分账结果请查看receivers中的result字段）:PROCESSING：处理中,FINISHED：分账完成
	Receivers     []*ProfitSharingReceiver `json:"receivers,omitempty"` // 分账接收方列表
}

type ProfitShareOrderQuery struct {
	SubMchid      string                   `json:"sub_mchid,omitempty"` // 子商户号，即分账的出资商户号【服务商模式】
	TransactionId string                   `json:"transaction_id"`      // 微信订单号
	OutOrderNo    string                   `json:"out_order_no"`        // 商户分账单号
	OrderId       string                   `json:"order_id"`            // 微信分账单号
	State         string                   `json:"state"`               // 分账单状态（每个接收方的分账结果请查看receivers中的result字段）:PROCESSING：处理中,FINISHED：分账完成
	Receivers     []*ProfitSharingReceiver `json:"receivers,omitempty"` // 分账接收方列表
}

type ProfitShareReturn struct {
	SubMchid    string `json:"sub_mchid"`             // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	OrderId     string `json:"order_id"`              // 微信分账单号，微信系统返回的唯一标识
	OutOrderNo  string `json:"out_order_no"`          // 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OutReturnNo string `json:"out_return_no"`         // 此回退单号是商户在自己后台生成的一个新的回退单号，在商户后台唯一
	ReturnId    string `json:"return_id"`             // 微信分账回退单号，微信系统返回的唯一标识
	ReturnMchid string `json:"return_mchid"`          // 只能对原分账请求中成功分给商户接收方进行回退
	Amount      int    `json:"amount"`                // 需要从分账接收方回退的金额，单位为分，只能为整数
	Description string `json:"description"`           // 分账回退的原因描述
	Result      string `json:"result"`                // 回退结果: PROCESSING：处理中，SUCCESS：已成功，FAILED：已失败
	FailReason  string `json:"fail_reason,omitempty"` // 失败原因: ACCOUNT_ABNORMAL : 分账接收方账户异常，TIME_OUT_CLOSED : 超时关单
	CreateTime  string `json:"create_time"`           // 分账回退创建时间
	FinishTime  string `json:"finish_time"`           // 分账回退完成时间
}

type ProfitShareReturnResult struct {
	SubMchid    string `json:"sub_mchid,omitempty"`   // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	OrderId     string `json:"order_id"`              // 微信分账单号，微信系统返回的唯一标识
	OutOrderNo  string `json:"out_order_no"`          // 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OutReturnNo string `json:"out_return_no"`         // 此回退单号是商户在自己后台生成的一个新的回退单号，在商户后台唯一
	ReturnId    string `json:"return_id"`             // 微信分账回退单号，微信系统返回的唯一标识
	ReturnMchid string `json:"return_mchid"`          // 只能对原分账请求中成功分给商户接收方进行回退
	Amount      int    `json:"amount"`                // 需要从分账接收方回退的金额，单位为分，只能为整数
	Description string `json:"description"`           // 分账回退的原因描述
	Result      string `json:"result"`                // 回退结果: PROCESSING：处理中，SUCCESS：已成功，FAILED：已失败
	FailReason  string `json:"fail_reason,omitempty"` // 失败原因: ACCOUNT_ABNORMAL : 分账接收方账户异常，TIME_OUT_CLOSED : 超时关单
	CreateTime  string `json:"create_time"`           // 分账回退创建时间
	FinishTime  string `json:"finish_time"`           // 分账回退完成时间
}

type ProfitShareOrderUnfreeze struct {
	SubMchid      string                   `json:"sub_mchid,omitempty"` // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	TransactionId string                   `json:"transaction_id"`      // 微信支付订单号
	OutOrderNo    string                   `json:"out_order_no"`        // 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OrderId       string                   `json:"order_id"`            // 微信分账单号，微信系统返回的唯一标识
	State         string                   `json:"state"`               // 分账单状态（每个接收方的分账结果请查看receivers中的result字段）:PROCESSING：处理中,FINISHED：分账完成
	Receivers     []*ProfitSharingReceiver `json:"receivers,omitempty"` // 分账接收方列表
}

type ProfitShareUnsplitAmount struct {
	TransactionId string `json:"transaction_id"` // 微信支付订单号
	UnsplitAmount int    `json:"unsplit_amount"` // 订单剩余待分金额，整数，单位为分
}

// 分账接收方
type ProfitSharingReceiver struct {
	Amount      int    `json:"amount"`      // 分账金额
	Description string `json:"description"` // 分账描述
	Type        string `json:"type"`        // 分账接收方类型
	Account     string `json:"account"`     // 分账接收方帐号
	Result      string `json:"result"`      // 分账结果,PENDING：待分账,SUCCESS：分账成功,CLOSED：已关闭
	FailReason  string `json:"fail_reason"` // 分账失败原因ACCOUNT_ABNORMAL : 分账接收账户异常、NO_RELATION : 分账关系已解除、RECEIVER_HIGH_RISK : 高风险接收方、RECEIVER_REAL_NAME_NOT_VERIFIED : 接收方未实名、NO_AUTH : 分账权限已解除
	CreateTime  string `json:"create_time"` // 分账创建时间,遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE
	FinishTime  string `json:"finish_time"` // 分账完成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE
	DetailId    string `json:"detail_id"`   // 分账明细单号, 微信分账明细单号，每笔分账业务执行的明细单号，可与资金账单对账使用
}

type ProfitShareAddReceiver struct {
	SubMchid       string `json:"sub_mchid,omitempty"`       // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	Type           string `json:"type"`                      // 分账接收方类型MERCHANT_ID：商户ID,PERSONAL_OPENID：个人openid（由父商户APPID转换得到）
	Account        string `json:"account"`                   // 分账接收方帐号
	Name           string `json:"name,omitempty"`            // 分账接收方类型是MERCHANT_ID时，是商户全称（必传），当商户是小微商户或个体户时，是开户人姓名 分账接收方类型是PERSONAL_OPENID时，是个人姓名（选传，传则校验）
	RelationType   string `json:"relation_type"`             // 商户与接收方的关系。STORE：门店STAFF：员工	STORE_OWNER：店主	PARTNER：合作伙伴	HEADQUARTER：总部	BRAND：品牌方	DISTRIBUTOR：分销商	USER：用户	SUPPLIER： 供应商	CUSTOM：自定义
	CustomRelation string `json:"custom_relation,omitempty"` // 子商户与接收方具体的关系，本字段最多10个字。当字段relation_type的值为CUSTOM时，本字段必填;当字段relation_type的值不为CUSTOM时，本字段无需填写。
}

type ProfitShareDeleteReceiver struct {
	SubMchid string `json:"sub_mchid,omitempty"` // 子商户号，分账回退的接收商户，对应原分账出资的分账方商户【服务商模式】
	Type     string `json:"type"`                // 分账接收方类型MERCHANT_ID：商户ID,PERSONAL_OPENID：个人openid（由父商户APPID转换得到）
	Account  string `json:"account"`             // 分账接收方帐号
}

type ProfitShareMerchantConfigs struct {
	SubMchId string `json:"sub_mchid"` // 子商户号
	MaxRatio int    `json:"max_ratio"` // 最大分账比例 (单位万分比，比如2000表示20%)
}

type ProfitShareBills struct {
	DownloadUrl string `json:"download_url"` // 下载地址	原始账单（gzip需要解压缩）的摘要值，用于校验文件的完整性。
	HashType    string `json:"hash_type"`    // 哈希类型 原始账单（gzip需要解压缩）的摘要算法，用于校验文件的完整性
	HashValue   string `json:"hash_value"`   // 哈希值	供下一步请求账单文件的下载地址，该地址30s内有效。
}

type DiscountCardApply struct {
	PrepareCardToken string `json:"prepare_card_token"` // 预领卡请求token，在引导用户进入先享卡领卡时，需要传入prepare_card_token
}

type DiscountCardQuery struct {
	CardId         string `json:"card_id"`          // 先享卡Id，唯一标识一个先享卡。
	CardTemplateId string `json:"card_template_id"` // 先享卡卡模板Id，唯一定义此资源的标识。
	Openid         string `json:"openid"`           // 微信用户在商户对应appid下的唯一标识
	OutCardCode    string `json:"out_card_code"`    // 商户领卡号
	Appid          string `json:"appid"`            // 公众账号Id
	Mchid          string `json:"mchid"`            // 商户号
	TimeRange      *struct {
		BeginTime string `json:"begin_time"` // 约定开始时间
		EndTime   string `json:"end_time"`   // 约定结束时间
	} `json:"time_range"` // 用户领取先享卡之后，约定的生效时间和到期时间。
	State            string          `json:"state"`                       // 先享卡的守约状态：ONGOING：约定进行中，SETTLING：约定到期核对中，FINISHED：已完成约定，UNFINISHED：未完成约定
	UnfinishedReason string          `json:"unfinished_reason,omitempty"` // 用户未完成约定的原因
	TotalAmount      int             `json:"total_amount,omitempty"`      // 享受优惠总金额
	PayInformation   *PayInformation `json:"pay_information,omitempty"`   // 用户退回优惠的付款信息
	CreateTime       string          `json:"create_time"`                 // 创卡时间
	Objectives       []*Objective    `json:"objectives"`                  // 用户先享卡目标列表
	Rewards          []*Reward       `json:"rewards"`                     // 用户先享卡优惠列表
	SharerOpenid     string          `json:"sharer_openid,omitempty"`     // 邀请者用户标识
}

type PayInformation struct {
	PayAmount     int    `json:"pay_amount"`               // 用户需要退回优惠而付款的金额，单位为：分
	PayState      string `json:"pay_state"`                // 用户付款状态：PAYING：付款中，PAID：已付款
	TransactionId string `json:"transaction_id,omitempty"` // 微信支付订单号，仅在订单成功收款时才返回
	PayTime       string `json:"pay_time,omitempty"`       // 用户成功支付的时间，仅在订单成功收款时才返回
}

type Objective struct {
	ObjectiveId                string                       `json:"objective_id"`                           // 由先享卡平台生成，唯一标识一个先享卡目标。商户需要记录该目标Id，进行同步用户记录
	Name                       string                       `json:"name"`                                   // 目标的名称
	Count                      int                          `json:"count"`                                  // 履约目标需要完成的数量，必须大于0
	Unit                       string                       `json:"unit"`                                   // 目标的单位
	Description                string                       `json:"description"`                            // 对先享卡目标的补充信息
	ObjectiveCompletionRecords []*ObjectiveCompletionRecord `json:"objective_completion_records,omitempty"` // 用户完成的目标明细列表
}

type ObjectiveCompletionRecord struct {
	ObjectiveCompletionSerialNo string `json:"objective_completion_serial_no"` // 目标完成流水号
	ObjectiveId                 string `json:"objective_id"`                   // 微信先享卡为每个先享卡目标分配的唯一Id
	CompletionTime              string `json:"completion_time"`                // 用户履约行为发生的时间
	CompletionType              string `json:"completion_type"`                // 目标完成类型： INCREASE：增加数量，DECREASE：减少数量
	Description                 string `json:"description"`                    // 用户本次履约的描述
	CompletionCount             int    `json:"completion_count"`               // 用户本次履约的数量，必须大于0
	Remark                      string `json:"remark,omitempty"`               // 对于用户履约情况的一些补充信息
}

type Reward struct {
	RewardId           string               `json:"reward_id"`                      // 由先享卡平台生成，唯一标识一个先享卡目标。商户需要记录该优惠Id，进行同步用户记录
	Name               string               `json:"name"`                           // 优惠名称
	CountType          string               `json:"count_type"`                     // 优惠数量的类型标识：COUNT_UNLIMITED：不限数量，COUNT_LIMIT：有限数量
	Count              int                  `json:"count"`                          // 本项优惠可使用的数量，必须大于0
	Unit               string               `json:"unit"`                           // 优惠的单位
	Amount             int                  `json:"amount"`                         // 优惠金额，此项优惠对应的优惠总金额，单位：分，必须大于0
	Description        string               `json:"description,omitempty"`          // 对先享卡优惠的补充信息
	RewardUsageRecords []*RewardUsageRecord `json:"reward_usage_records,omitempty"` // 优惠使用记录列表
}

type RewardUsageRecord struct {
	RewardUsageSerialNo string `json:"reward_usage_serial_no"` // 优惠使用记录流水号
	RewardId            string `json:"reward_id"`              // 微信先享卡为每个先享卡优惠分配的唯一Id
	UsageTime           string `json:"usage_time"`             // 用户使用优惠的时间
	UsageType           string `json:"usage_type"`             // 目标完成类型：INCREASE：增加数量，DECREASE：减少数量
	Description         string `json:"description"`            // 用户获得奖励的描述
	UsageCount          int    `json:"usage_count"`            // 用户本次获得的奖励数量，必须大于0
	Amount              int    `json:"amount"`                 // 优惠金额，用户此项本次享受的优惠对应的优惠总金额，单位：分，必须大于0
	Remark              string `json:"remark,omitempty"`       // 对于用户奖励情况的一些补充信息
}

type SmartGuideReg struct {
	GuideId string `json:"guide_id"` // 服务人员在服务人员系统中的唯一标识
}

type SmartGuideQuery struct {
	Data       []*SmartGuide `json:"data"`        // 服务人员列表
	TotalCount int           `json:"total_count"` // 服务人员数量
	Limit      int           `json:"limit"`       // 该次请求可返回的最大资源条数，不大于10
	Offset     int           `json:"offset"`      // 该次请求资源的起始位置，默认值为0
}

type SmartGuide struct {
	GuideId string `json:"guide_id"`          // 服务人员在服务人员系统中的唯一标识
	StoreId int    `json:"store_id"`          // 门店在微信支付商户平台的唯一标识
	Name    string `json:"name"`              // 服务人员姓名
	Mobile  string `json:"mobile"`            // 员工在商户个人/企业微信通讯录上设置的手机号码（加密信息，需解密）
	Userid  string `json:"userid,omitempty"`  // 员工在商户企业微信通讯录使用的唯一标识，使用企业微信商家时返回
	WorkId  string `json:"work_id,omitempty"` // 服务人员通过小程序注册时填写的工号，使用个人微信商家时返回
}

type BusinessAuthPointsQuery struct {
	Openid          string `json:"openid"`                     // 顾客授权时使用的小程序上的openid
	AuthorizeState  string `json:"authorize_state"`            // 顾客授权商圈积分结果：UNAUTHORIZED：未授权，AUTHORIZED：已授权，DEAUTHORIZED：已取消授权
	AuthorizeTime   string `json:"authorize_time,omitempty"`   // 顾客成功授权商圈积分的时间
	DeauthorizeTime string `json:"deauthorize_time,omitempty"` // 顾客关闭授权商圈积分的时间
}

type Transfer struct {
	OutBatchNo string `json:"out_batch_no"` // 商户系统内部的商家批次单号
	BatchId    string `json:"batch_id"`     // 微信批次单号，微信商家转账系统返回的唯一标识
	CreateTime string `json:"create_time"`  // 批次受理成功时返回
}

type TransferQuery struct {
	TransferBatch      *TransferBatch    `json:"transfer_batch"`                 // 转账批次单基本信息
	TransferDetailList []*TransferDetail `json:"transfer_detail_list,omitempty"` // 当批次状态为“FINISHED”（已完成），且成功查询到转账明细单时返回
}

type TransferBatch struct {
	Mchid         string `json:"mchid"`                    // 微信支付分配的商户号
	OutBatchNo    string `json:"out_batch_no"`             // 商户系统内部的商家批次单号
	BatchId       string `json:"batch_id"`                 // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid         string `json:"appid"`                    // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	BatchStatus   string `json:"batch_status"`             // 批次状态
	BatchType     string `json:"batch_type"`               // 批次类型
	BatchName     string `json:"batch_name"`               // 该笔批量转账的名称
	BatchRemark   string `json:"batch_remark"`             // 转账说明，UTF8编码，最多允许32个字符
	CloseReason   string `json:"close_reason,omitempty"`   // 如果批次单状态为“CLOSED”（已关闭），则有关闭原因
	TotalAmount   int    `json:"total_amount"`             // 转账金额单位为分
	TotalNum      int    `json:"total_num"`                // 一个转账批次单最多发起三千笔转账
	CreateTime    string `json:"create_time,omitempty"`    // 批次受理成功时返回
	UpdateTime    string `json:"update_time,omitempty"`    // 批次最近一次状态变更的时间
	SuccessAmount int    `json:"success_amount,omitempty"` // 转账成功的金额，单位为分
	SuccessNum    int    `json:"success_num,omitempty"`    // 转账成功的笔数
	FailAmount    int    `json:"fail_amount,omitempty"`    // 转账失败的金额，单位为分
	FailNum       int    `json:"fail_num,omitempty"`       // 转账失败的笔数
}

type TransferDetail struct {
	DetailId     string `json:"detail_id"`     // 微信明细单号
	OutDetailNo  string `json:"out_detail_no"` // 商家明细单号
	DetailStatus string `json:"detail_status"` // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
}

type PartnerTransferQuery struct {
	SpMchid            string            `json:"sp_mchid"`                       // 微信支付分配的服务商商户号
	SubMchid           string            `json:"sub_mchid"`                      // 微信支付分配的特约商户号
	OutBatchNo         string            `json:"out_batch_no"`                   // 商户系统内部的商家批次单号
	BatchId            string            `json:"batch_id"`                       // 微信批次单号，微信商家转账系统返回的唯一标识
	SpAppid            string            `json:"sp_appid,omitempty"`             // 微信分配的服务商商户公众账号Id，特约商户授权类型为FUND_AUTHORIZATION_TYPE时才有该字段
	SubAppid           string            `json:"sub_appid"`                      // 微信分配的特约商户公众账号Id。特约商户appid
	BatchStatus        string            `json:"batch_status"`                   // 批次状态
	BatchType          string            `json:"batch_type"`                     // 批次类型
	AuthorizationType  string            `json:"authorization_type"`             // 特约商户授权类型
	BatchName          string            `json:"batch_name"`                     // 该笔批量转账的名称
	BatchRemark        string            `json:"batch_remark"`                   // 转账说明，UTF8编码，最多允许32个字符
	CloseReason        string            `json:"close_reason,omitempty"`         // 如果批次单状态为“CLOSED”（已关闭），则有关闭原因
	TotalAmount        int               `json:"total_amount"`                   // 转账金额单位为分
	TotalNum           int               `json:"total_num"`                      // 一个转账批次单最多发起三千笔转账
	CreateTime         string            `json:"create_time,omitempty"`          // 批次受理成功时返回
	UpdateTime         string            `json:"update_time,omitempty"`          // 批次最近一次状态变更的时间
	SuccessAmount      int               `json:"success_amount,omitempty"`       // 转账成功的金额，单位为分
	SuccessNum         int               `json:"success_num,omitempty"`          // 转账成功的笔数
	FailAmount         int               `json:"fail_amount,omitempty"`          // 转账失败的金额，单位为分
	FailNum            int               `json:"fail_num,omitempty"`             // 转账失败的笔数
	TransferPurpose    string            `json:"transfer_purpose"`               // 批量转账用途
	TransferDetailList []*TransferDetail `json:"transfer_detail_list,omitempty"` // 当批次状态为“FINISHED”（已完成），且成功查询到转账明细单时返回
}

type TransferDetailQuery struct {
	Mchid          string `json:"mchid"`                 // 微信支付分配的商户号
	OutBatchNo     string `json:"out_batch_no"`          // 商户系统内部的商家批次单号
	BatchId        string `json:"batch_id"`              // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid          string `json:"appid"`                 // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	OutDetailNo    string `json:"out_detail_no"`         // 商家明细单号
	DetailId       string `json:"detail_id"`             // 微信明细单号
	DetailStatus   string `json:"detail_status"`         // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
	TransferAmount int    `json:"transfer_amount"`       // 转账金额单位为分
	TransferRemark string `json:"transfer_remark"`       // 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string `json:"fail_reason,omitempty"` // 如果转账失败则有失败原因
	Openid         string `json:"openid"`                // 用户在直连商户appid下的唯一标识
	UserName       string `json:"user_name"`             // 收款方姓名（加密）
	InitiateTime   string `json:"initiate_time"`         // 转账发起的时间
	UpdateTime     string `json:"update_time"`           // 明细最后一次状态变更的时间
}

type PartnerTransferDetail struct {
	SpMchid        string `json:"sp_mchid"`              // 微信支付分配的服务商商户号
	OutBatchNo     string `json:"out_batch_no"`          // 商户系统内部的商家批次单号
	BatchId        string `json:"batch_id"`              // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid          string `json:"appid"`                 // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	OutDetailNo    string `json:"out_detail_no"`         // 商家明细单号
	DetailId       string `json:"detail_id"`             // 微信明细单号
	DetailStatus   string `json:"detail_status"`         // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
	TransferAmount int    `json:"transfer_amount"`       // 转账金额单位为分
	TransferRemark string `json:"transfer_remark"`       // 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string `json:"fail_reason,omitempty"` // 如果转账失败则有失败原因
	Openid         string `json:"openid"`                // 用户在直连商户appid下的唯一标识
	Username       string `json:"username"`              // 收款方姓名（加密）
	InitiateTime   string `json:"initiate_time"`         // 转账发起的时间
	UpdateTime     string `json:"update_time"`           // 明细最后一次状态变更的时间
}

type TransferMerchantQuery struct {
	TransferBatch      *TransferBatch    `json:"transfer_batch"`                 // 转账批次单基本信息
	TransferDetailList []*TransferDetail `json:"transfer_detail_list,omitempty"` // 当批次状态为“FINISHED”（已完成），且成功查询到转账明细单时返回
	Offset             int               `json:"offset,omitempty"`               // 该次请求资源（转账明细单）的起始位置
	Limit              int               `json:"limit,omitempty"`                // 该次请求可返回的最大资源（转账明细单）条数
}

type PartnerTransferMerchantQuery struct {
	SpMchid            string            `json:"sp_mchid"`                       // 微信支付分配的服务商商户号
	SubMchid           string            `json:"sub_mchid"`                      // 微信支付分配的特约商户号
	OutBatchNo         string            `json:"out_batch_no"`                   // 商户系统内部的商家批次单号
	BatchId            string            `json:"batch_id"`                       // 微信批次单号，微信商家转账系统返回的唯一标识
	SpAppid            string            `json:"sp_appid,omitempty"`             // 微信分配的服务商商户公众账号Id，特约商户授权类型为FUND_AUTHORIZATION_TYPE时才有该字段
	SubAppid           string            `json:"sub_appid"`                      // 微信分配的特约商户公众账号Id。特约商户appid
	BatchStatus        string            `json:"batch_status"`                   // 批次状态
	BatchType          string            `json:"batch_type"`                     // 批次类型
	AuthorizationType  string            `json:"authorization_type"`             // 特约商户授权类型
	BatchName          string            `json:"batch_name"`                     // 该笔批量转账的名称
	BatchRemark        string            `json:"batch_remark"`                   // 转账说明，UTF8编码，最多允许32个字符
	CloseReason        string            `json:"close_reason,omitempty"`         // 如果批次单状态为“CLOSED”（已关闭），则有关闭原因
	TotalAmount        int               `json:"total_amount"`                   // 转账金额单位为分
	TotalNum           int               `json:"total_num"`                      // 一个转账批次单最多发起三千笔转账
	CreateTime         string            `json:"create_time,omitempty"`          // 批次受理成功时返回
	UpdateTime         string            `json:"update_time,omitempty"`          // 批次最近一次状态变更的时间
	SuccessAmount      int               `json:"success_amount,omitempty"`       // 转账成功的金额，单位为分
	SuccessNum         int               `json:"success_num,omitempty"`          // 转账成功的笔数
	FailAmount         int               `json:"fail_amount,omitempty"`          // 转账失败的金额，单位为分
	FailNum            int               `json:"fail_num,omitempty"`             // 转账失败的笔数
	TransferPurpose    string            `json:"transfer_purpose"`               // 批量转账用途
	TransferDetailList []*TransferDetail `json:"transfer_detail_list,omitempty"` // 当批次状态为“FINISHED”（已完成），且成功查询到转账明细单时返回
}

type TransferMerchantDetail struct {
	OutBatchNo     string `json:"out_batch_no"`          // 商户系统内部的商家批次单号
	BatchId        string `json:"batch_id"`              // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid          string `json:"appid"`                 // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	OutDetailNo    string `json:"out_detail_no"`         // 商家明细单号
	DetailId       string `json:"detail_id"`             // 微信明细单号
	DetailStatus   string `json:"detail_status"`         // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
	TransferAmount int    `json:"transfer_amount"`       // 转账金额单位为分
	TransferRemark string `json:"transfer_remark"`       // 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string `json:"fail_reason,omitempty"` // 如果转账失败则有失败原因
	Openid         string `json:"openid"`                // 用户在直连商户appid下的唯一标识
	UserName       string `json:"user_name"`             // 收款方姓名（加密）
	InitiateTime   string `json:"initiate_time"`         // 转账发起的时间
	UpdateTime     string `json:"update_time"`           // 明细最后一次状态变更的时间
}

type PartnerTransferMerchantDetail struct {
	SpMchid        string `json:"sp_mchid"`              // 微信支付分配的服务商商户号
	OutBatchNo     string `json:"out_batch_no"`          // 商户系统内部的商家批次单号
	BatchId        string `json:"batch_id"`              // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid          string `json:"appid"`                 // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	OutDetailNo    string `json:"out_detail_no"`         // 商家明细单号
	DetailId       string `json:"detail_id"`             // 微信明细单号
	DetailStatus   string `json:"detail_status"`         // 明细状态：PROCESSING：转账中，SUCCESS：转账成功，FAIL：转账失败
	TransferAmount int    `json:"transfer_amount"`       // 转账金额单位为分
	TransferRemark string `json:"transfer_remark"`       // 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string `json:"fail_reason,omitempty"` // 如果转账失败则有失败原因
	Openid         string `json:"openid"`                // 用户在直连商户appid下的唯一标识
	Username       string `json:"username"`              // 收款方姓名（加密）
	InitiateTime   string `json:"initiate_time"`         // 转账发起的时间
	UpdateTime     string `json:"update_time"`           // 明细最后一次状态变更的时间
}

type TransferReceipt struct {
	OutBatchNo      string `json:"out_batch_no"`               // 商户系统内部的商家批次单号
	SignatureNo     string `json:"signature_no"`               // 电子回单申请单号，申请单据的唯一标识
	SignatureStatus string `json:"signature_status,omitempty"` // 电子回单状态：ACCEPTED:已受理，电子签章已受理成功，FINISHED:已完成。电子签章已处理完成
	HashType        string `json:"hash_type,omitempty"`        // 电子回单文件的hash方法，回单状态为：FINISHED时返回。
	HashValue       string `json:"hash_value,omitempty"`       // 电子回单文件的hash值，用于下载之后验证文件的完整、正确性，回单状态为：FINISHED时返回。
	DownloadUrl     string `json:"download_url,omitempty"`     // 电子回单文件的下载地址，回单状态为：FINISHED时返回
	CreateTime      string `json:"create_time,omitempty"`      // 电子签章单创建时间
	UpdateTime      string `json:"update_time,omitempty"`      // 电子签章单最近一次状态变更的时间
}

type TransferReceiptQuery struct {
	OutBatchNo      string `json:"out_batch_no"`               // 商户系统内部的商家批次单号
	SignatureNo     string `json:"signature_no"`               // 电子回单申请单号，申请单据的唯一标识
	SignatureStatus string `json:"signature_status,omitempty"` // 电子回单状态：ACCEPTED:已受理，电子签章已受理成功，FINISHED:已完成。电子签章已处理完成
	HashType        string `json:"hash_type,omitempty"`        // 电子回单文件的hash方法，回单状态为：FINISHED时返回。
	HashValue       string `json:"hash_value,omitempty"`       // 电子回单文件的hash值，用于下载之后验证文件的完整、正确性，回单状态为：FINISHED时返回。
	DownloadUrl     string `json:"download_url,omitempty"`     // 电子回单文件的下载地址，回单状态为：FINISHED时返回
	CreateTime      string `json:"create_time,omitempty"`      // 电子签章单创建时间
	UpdateTime      string `json:"update_time,omitempty"`      // 电子签章单最近一次状态变更的时间
}

type TransferDetailReceipt struct {
	AcceptType      string `json:"accept_type"`                // 电子回单受理类型
	OutBatchNo      string `json:"out_batch_no,omitempty"`     // 商户系统内部的商家批次单号
	OutDetailNo     string `json:"out_detail_no"`              // 商家明细单号
	SignatureNo     string `json:"signature_no"`               // 电子回单申请单号，申请单据的唯一标识
	SignatureStatus string `json:"signature_status,omitempty"` // 电子回单状态：ACCEPTED:已受理，电子签章已受理成功，FINISHED:已完成。电子签章已处理完成
	HashType        string `json:"hash_type,omitempty"`        // 电子回单文件的hash方法，回单状态为：FINISHED时返回。
	HashValue       string `json:"hash_value,omitempty"`       // 电子回单文件的hash值，用于下载之后验证文件的完整、正确性，回单状态为：FINISHED时返回。
	DownloadUrl     string `json:"download_url,omitempty"`     // 电子回单文件的下载地址，回单状态为：FINISHED时返回
}

type TransferDetailReceiptQuery struct {
	AcceptType      string `json:"accept_type"`                // 电子回单受理类型
	OutBatchNo      string `json:"out_batch_no,omitempty"`     // 商户系统内部的商家批次单号
	OutDetailNo     string `json:"out_detail_no"`              // 商家明细单号
	SignatureNo     string `json:"signature_no"`               // 电子回单申请单号，申请单据的唯一标识
	SignatureStatus string `json:"signature_status,omitempty"` // 电子回单状态：ACCEPTED:已受理，电子签章已受理成功，FINISHED:已完成。电子签章已处理完成
	HashType        string `json:"hash_type,omitempty"`        // 电子回单文件的hash方法，回单状态为：FINISHED时返回。
	HashValue       string `json:"hash_value,omitempty"`       // 电子回单文件的hash值，用于下载之后验证文件的完整、正确性，回单状态为：FINISHED时返回。
	DownloadUrl     string `json:"download_url,omitempty"`     // 电子回单文件的下载地址，回单状态为：FINISHED时返回
}

type EcommerceBalance struct {
	SubMchid        string `json:"sub_mchid"`                // 特约商户号
	AccountType     string `json:"account_type,omitempty"`   // 账户类型
	AvailableAmount int    `json:"available_amount"`         // 可用余额（单位：分），此余额可做提现操作
	PendingAmount   int    `json:"pending_amount,omitempty"` // 不可用余额（单位：分）
}

type MerchantBalance struct {
	AvailableAmount int `json:"available_amount"`         // 可用余额（单位：分），此余额可做提现操作
	PendingAmount   int `json:"pending_amount,omitempty"` // 不可用余额（单位：分）
}

type MerchantIncomeRecord struct {
	Data       []*IncomeData `json:"data,omitempty"` // 单次查询返回的银行来账记录列表结果数组，如果查询结果为空时，则为空数组
	Links      *Link         `json:"links"`          // 返回前后页和当前页面的访问链接
	Offset     int           `json:"offset"`         // 该次请求资源的起始位置，请求中包含偏移量时应答消息返回相同偏移量，否则返回默认值0
	Limit      int           `json:"limit"`          // 经过条件筛选，本次查询到的银行来账记录条数
	TotalCount int           `json:"total_count"`    // 经过条件筛选，查询到的银行来账记录总数
}

type IncomeData struct {
	Mchid             string `json:"mchid"`               // 微信支付分配的商户号
	AccountType       string `json:"account_type"`        // 需查询银行来账记录商户的账户类型：BASIC：基本账户，OPERATION：运营账户，FEES：手续费账户
	IncomeRecordType  string `json:"income_record_type"`  // 银行来账类型
	IncomeRecordId    string `json:"income_record_id"`    // 银行来账的微信单号
	Amount            int    `json:"amount"`              // 银行来账金额，单位为分，只能为整数
	SuccessTime       string `json:"success_time"`        // 银行来账完成时间
	BankName          string `json:"bank_name"`           // 银行来账的付款方银行名称，由于部分银行的数据获取限制，该字段有可能为空
	BankAccountName   string `json:"bank_account_name"`   // 银行来账的付款方银行账户信息，户名为全称、明文，由于部分银行的数据获取限制，该字段有可能为空
	BankAccountNumber string `json:"bank_account_number"` // 四位掩码+付款方银行卡尾号后四位
	RechargeRemark    string `json:"recharge_remark"`     // 银行备注
}

type Link struct {
	Next string `json:"next"` // 下一页链接
	Prev string `json:"prev"` // 上一页链接
	Self string `json:"self"` // 当前链接
}

type PartnerIncomeRecord struct {
	Data       []*PartnerIncomeData `json:"data,omitempty"` // 单次查询返回的银行来账记录列表结果数组，如果查询结果为空时，则为空数组
	Links      *Link                `json:"links"`          // 返回前后页和当前页面的访问链接
	Offset     int                  `json:"offset"`         // 该次请求资源的起始位置，请求中包含偏移量时应答消息返回相同偏移量，否则返回默认值0
	Limit      int                  `json:"limit"`          // 经过条件筛选，本次查询到的银行来账记录条数
	TotalCount int                  `json:"total_count"`    // 经过条件筛选，查询到的银行来账记录总数
}

type PartnerIncomeData struct {
	SubMchid          string `json:"sub_mchid"`           // 特约商户号
	AccountType       string `json:"account_type"`        // 需查询银行来账记录商户的账户类型：BASIC：基本账户，OPERATION：运营账户，FEES：手续费账户
	IncomeRecordType  string `json:"income_record_type"`  // 银行来账类型
	IncomeRecordId    string `json:"income_record_id"`    // 银行来账的微信单号
	Amount            int    `json:"amount"`              // 银行来账金额，单位为分，只能为整数
	SuccessTime       string `json:"success_time"`        // 银行来账完成时间
	BankName          string `json:"bank_name"`           // 银行来账的付款方银行名称，由于部分银行的数据获取限制，该字段有可能为空
	BankAccountName   string `json:"bank_account_name"`   // 银行来账的付款方银行账户信息，户名为全称、明文，由于部分银行的数据获取限制，该字段有可能为空
	BankAccountNumber string `json:"bank_account_number"` // 四位掩码+付款方银行卡尾号后四位
	RechargeRemark    string `json:"recharge_remark"`     // 银行备注
}

type Apply4SubSubmit struct {
	ApplymentId int64 `json:"applyment_id"` // 微信支付申请单号
}

type Apply4SubQuery struct {
	BusinessCode      string                      `json:"business_code"`       // 业务申请编号
	ApplymentId       int64                       `json:"applyment_id"`        // 微信支付申请单号
	SubMchid          string                      `json:"sub_mchid"`           // 特约商户号
	SignUrl           string                      `json:"sign_url"`            // 超级管理员签约链接
	ApplymentState    string                      `json:"applyment_state"`     // 申请单状态
	ApplymentStateMsg string                      `json:"applyment_state_msg"` // 申请状态描述
	AuditDetail       []*Applyment4SubAuditDetail `json:"audit_detail"`        // 驳回原因详情
}

type Applyment4SubAuditDetail struct {
	Field        string `json:"field"`         // 字段名
	FieldName    string `json:"field_name"`    // 字段名称
	RejectReason string `json:"reject_reason"` // 驳回原因
}

type Apply4SubQuerySettlement struct {
	AccountType      string `json:"account_type"`       // 账户类型
	AccountBank      string `json:"account_bank"`       // 开户银行
	BankName         string `json:"bank_name"`          // 开户银行全称（含支行）
	BankBranchId     string `json:"bank_branch_id"`     // 开户银行联行号
	AccountNumber    string `json:"account_number"`     // 银行账号
	VerifyResult     string `json:"verify_result"`      // 汇款验证结果
	VerifyFailReason string `json:"verify_fail_reason"` // 汇款验证失败原因
}

type FavorBatchCreate struct {
	StockId    string `json:"stock_id"`    // 批次号
	CreateTime string `json:"create_time"` // 创建时间，遵循rfc3339标准格式
}

type FavorBatchGrant struct {
	CouponId string `json:"coupon_id"` // 微信为代金券唯一分配的id
}

type FavorBatchStart struct {
	StockId   string `json:"stock_id"`   // 微信为每个代金券批次分配的唯一Id
	StartTime string `json:"start_time"` // 生效时间，遵循rfc3339标准格式
}

type FavorBatchList struct {
	Data       []*FavorBatch `json:"data,omitempty"` // 批次详情列表
	TotalCount int           `json:"total_count"`    // 批次总数
	Offset     int           `json:"offset"`         // 分页页码
	Limit      int           `json:"limit"`          // 分页大小
}

type FavorBatch struct {
	StockId            string        `json:"stock_id"`             // 微信为每个代金券批次分配的唯一Id
	StockCreatorMchid  string        `json:"stock_creator_mchid"`  // 创建批次的商户号
	StockName          string        `json:"stock_name"`           // 批次名称
	Status             string        `json:"status"`               // 批次状态
	CreateTime         string        `json:"create_time"`          // 创建时间，遵循rfc3339标准格式
	Description        string        `json:"description"`          // 使用说明
	StockUseRule       *StockUseRule `json:"stock_use_rule"`       // 普通发券批次特定信息
	AvailableBeginTime string        `json:"available_begin_time"` // 可用开始时间，遵循rfc3339标准格式
	AvailableEndTime   string        `json:"available_end_time"`   // 可用结束时间，遵循rfc3339标准格式
	DistributedCoupons int           `json:"distributed_coupons"`  // 已发券数量
	NoCash             bool          `json:"no_cash"`              // 是否无资金流
	StartTime          string        `json:"start_time"`           // 激活批次的时间
	StopTime           string        `json:"stop_time"`            // 终止批次的时间
	CutToMessage       *CutToMessage `json:"cut_to_message"`       // 单品优惠特定信息
	Singleitem         bool          `json:"singleitem"`           // 是否单品优惠
	StockType          string        `json:"stock_type"`           // 批次类型
}

type StockUseRule struct {
	MaxCoupons        int                `json:"max_coupons"`          // 发放总上限
	MaxAmount         int                `json:"max_amount"`           // 总预算
	MaxAmountByDay    int                `json:"max_amount_by_day"`    // 当天发放上限金额
	FixedNormalCoupon *FixedNormalCoupon `json:"fixed_normal_coupon"`  // 固定面额发券批次特定信息
	MaxCouponsPerUser int                `json:"max_coupons_per_user"` // 单个用户可领个数
	CouponType        string             `json:"coupon_type"`          // 券类型
	GoodsTag          []string           `json:"goods_tag,omitempty"`  // 订单优惠标记
	TradeType         []string           `json:"trade_type"`           // 支付方式
	CombineUse        bool               `json:"combine_use"`          // 是否可叠加其他优惠
}

type FixedNormalCoupon struct {
	CouponAmount       int `json:"coupon_amount"`       // 面额，单位：分
	TransactionMinimum int `json:"transaction_minimum"` // 使用券金额门槛，单位：分
}

type CutToMessage struct {
	SinglePriceMax int `json:"single_price_max"` // 可用优惠的商品最高单价，单位：分
	CutToPrice     int `json:"cut_to_price"`     // 减至后的优惠单价，单位：分
}

type FavorDetail struct {
	StockId                 string                   `json:"stock_id"`                  // 微信为每个代金券批次分配的唯一Id
	StockCreatorMchid       string                   `json:"stock_creator_mchid"`       // 创建批次的商户号
	CouponId                string                   `json:"coupon_id"`                 // 微信为代金券唯一分配的id
	CutToMessage            *CutToMessage            `json:"cut_to_message"`            // 单品优惠特定信息
	CouponName              string                   `json:"coupon_name"`               // 代金券名称
	Status                  string                   `json:"status"`                    // 代金券状态
	Description             string                   `json:"description"`               // 使用说明
	CreateTime              string                   `json:"create_time"`               // 领券时间，遵循rfc3339标准格式
	CouponType              string                   `json:"coupon_type"`               // 券类型
	NoCash                  bool                     `json:"no_cash"`                   // 是否无资金流
	AvailableBeginTime      string                   `json:"available_begin_time"`      // 可用开始时间，遵循rfc3339标准格式
	AvailableEndTime        string                   `json:"available_end_time"`        // 可用结束时间，遵循rfc3339标准格式
	Singleitem              bool                     `json:"singleitem"`                // 是否单品优惠
	NormalCouponInformation *NormalCouponInformation `json:"normal_coupon_information"` // 普通满减券面额、门槛信息
}

type NormalCouponInformation struct {
	CouponAmount       int `json:"coupon_amount"`       // 面额，单位：分
	TransactionMinimum int `json:"transaction_minimum"` // 使用券金额门槛，单位：分
}

type FavorMerchant struct {
	StockId    string   `json:"stock_id"`       // 批次号
	Data       []string `json:"data,omitempty"` // 可用商户列表
	TotalCount int      `json:"total_count"`    // 批次总数
	Offset     int      `json:"offset"`         // 分页页码
	Limit      int      `json:"limit"`          // 分页大小
}

type FavorItems struct {
	StockId    string   `json:"stock_id"`       // 批次号
	Data       []string `json:"data,omitempty"` // 可用商户列表
	TotalCount int      `json:"total_count"`    // 批次总数
	Offset     int      `json:"offset"`         // 分页页码
	Limit      int      `json:"limit"`          // 分页大小
}

type FavorUserCoupons struct {
	Data       []*UserCoupon `json:"data,omitempty"` // 批次详情列表
	TotalCount int           `json:"total_count"`    // 批次总数
	Offset     int           `json:"offset"`         // 分页页码
	Limit      int           `json:"limit"`          // 分页大小
}

type UserCoupon struct {
	StockCreatorMchid       string                   `json:"stock_creator_mchid"` // 创建批次的商户号
	StockId                 string                   `json:"stock_id"`            // 批次号
	CouponId                string                   `json:"coupon_id"`
	CouponName              string                   `json:"coupon_name"`
	CouponType              string                   `json:"coupon_type"`
	CutToMessage            *CutToMessage            `json:"cut_to_message"`            // 单品优惠特定信息
	Status                  string                   `json:"status"`                    // 代金券状态
	Description             string                   `json:"description"`               // 使用说明
	CreateTime              string                   `json:"create_time"`               // 领券时间，遵循rfc3339标准格式
	NoCash                  bool                     `json:"no_cash"`                   // 是否无资金流
	AvailableBeginTime      string                   `json:"available_begin_time"`      // 可用开始时间，遵循rfc3339标准格式
	AvailableEndTime        string                   `json:"available_end_time"`        // 可用结束时间，遵循rfc3339标准格式
	Singleitem              bool                     `json:"singleitem"`                // 是否单品优惠
	NormalCouponInformation *NormalCouponInformation `json:"normal_coupon_information"` // 普通满减券面额、门槛信息
	ConsumeInformation      *ConsumeInformation      `json:"consume_information"`       // 已实扣代金券信息
}

type ConsumeInformation struct {
	ConsumeTime   string               `json:"consume_time"`           // 核销时间，遵循rfc3339标准格式
	ConsumeMchid  string               `json:"consume_mchid"`          // 核销商户号
	TransactionId string               `json:"transaction_id"`         // 核销订单号
	GoodsDetail   []*CouponGoodsDetail `json:"goods_detail,omitempty"` // 商户下单单品信息
}

type CouponGoodsDetail struct {
	GoodsId        string `json:"goods_id"`        // 商品编码
	Quantity       int    `json:"quantity"`        // 商品数量
	Price          int    `json:"price"`           // 单品价格，单位为分
	DiscountAmount int    `json:"discount_amount"` // 商品优惠金额
}

type FavorFlowDownload struct {
	Url       string `json:"url"`        // 流水文件下载链接，30s内有效
	HashValue string `json:"hash_value"` // 文件内容的哈希值，防止篡改
	HashType  string `json:"hash_type"`  // 哈希算法类型，目前只支持sha1
}

type FavorCallbackUrl struct {
	UpdateTime string `json:"update_time"` // 修改时间，遵循rfc3339标准格式
	NotifyUrl  string `json:"notify_url"`  // 通知地址
}

type FavorBatchPause struct {
	PauseTime string `json:"pause_time"` // 暂停时间，遵循rfc3339标准格式
	StockId   string `json:"stock_id"`   // 批次号
}

type FavorBatchRestart struct {
	RestartTime string `json:"restart_time"` // 生效时间，遵循rfc3339标准格式
	StockId     string `json:"stock_id"`     // 批次号
}

type BusiFavorBatchDetail struct {
	StockId              string                `json:"stock_id"`                         // 批次号
	StockName            string                `json:"stock_name"`                       // 商家券批次名称
	BelongMerchant       string                `json:"belong_merchant"`                  // 批次归属商户号
	Comment              string                `json:"comment"`                          // 批次备注
	GoodsName            string                `json:"goods_name"`                       // 适用商品范围
	StockType            string                `json:"stock_type"`                       // 批次类型
	StockState           string                `json:"stock_state"`                      // 批次状态
	CouponUseRule        *CouponUseRule        `json:"coupon_use_rule"`                  // 券核销相关规则
	StockSendRule        *StockSendRule        `json:"stock_send_rule"`                  // 券发放相关规则
	CustomEntrance       *CustomEntrance       `json:"custom_entrance,omitempty"`        // 自定义入口
	DisplayPatternInfo   *DisplayPatternInfo   `json:"display_pattern_info,omitempty"`   // 创建批次时的样式信息
	CouponCodeMode       string                `json:"coupon_code_mode"`                 // 券code模式
	CouponCodeCount      *CouponCodeCount      `json:"coupon_code_count,omitempty"`      // 券code数量
	NotifyConfig         *NotifyConfig         `json:"notify_config,omitempty"`          // 事件通知配置
	SendCountInformation *SendCountInformation `json:"send_count_information,omitempty"` // 批次发放情况
}

type CouponUseRule struct {
	CouponAvailableTime *CouponAvailableTime `json:"coupon_available_time"`
	FixedNormalCoupon   *struct {
		DiscountAmount     int `json:"discount_amount"`     // 优惠金额，单位：分
		TransactionMinimum int `json:"transaction_minimum"` // 消费门槛，单位：分
	} `json:"fixed_normal_coupon,omitempty"` // 固定面额满减券使用规则
	DiscountCoupon *struct {
		DiscountPercent    int `json:"discount_percent"`    // 折扣百分比，例如：86为八六折
		TransactionMinimum int `json:"transaction_minimum"` // 消费门槛，单位：分
	} `json:"discount_coupon,omitempty"` // 折扣券使用规则
	ExchangeCoupon *struct {
		ExchangePrice      int `json:"exchange_price"`      // 单品换购价，单位：分
		TransactionMinimum int `json:"transaction_minimum"` // 消费门槛，单位：分
	} `json:"exchange_coupon,omitempty"` // 换购券使用规则
	UseMethod         string `json:"use_method"`          // 核销方式
	MiniProgramsAppid string `json:"mini_programs_appid"` // 小程序appid
	MiniProgramsPath  string `json:"mini_programs_path"`  // 小程序path
}

type CouponAvailableTime struct {
	AvailableBeginTime       string           `json:"available_begin_time"`        // 批次开始时间，遵循rfc3339标准格式
	AvailableEndTime         string           `json:"available_end_time"`          // 批次结束时间，遵循rfc3339标准格式
	AvailableDayAfterReceive int              `json:"available_day_after_receive"` // 生效后N天内有效
	AvailableWeek            *AvailableWeek   `json:"available_week"`              // 固定周期有效时间段
	IrregularyAvaliableTime  []*AvailableTime `json:"irregulary_avaliable_time"`   // 无规律的有效时间段
	WaitDaysAfterReceive     int              `json:"wait_days_after_receive"`     // 领取后N天开始生效
}

type AvailableWeek struct {
	WeekDay          []int            `json:"week_day"` // 可用星期数，0代表周日，1代表周一，以此类推
	AvailableDayTime []*AvailableTime `json:"available_day_time"`
}

type AvailableTime struct {
	BeginTime int `json:"begin_time"` // 开始时间
	EndTime   int `json:"end_time"`   // 结束时间
}

type StockSendRule struct {
	MaxAmount          int  `json:"max_amount"`           // 批次总预算
	MaxCoupons         int  `json:"max_coupons"`          // 批次最大发放个数
	MaxCouponsPerUser  int  `json:"max_coupons_per_user"` // 用户最大可领个数
	MaxAmountByDay     int  `json:"max_amount_by_day"`    // 单天发放上限金额
	MaxCouponsByDay    int  `json:"max_coupons_by_day"`   // 单天发放上限个数
	NaturalPersonLimit bool `json:"natural_person_limit"` // 是否开启自然人限制
	PreventApiAbuse    bool `json:"prevent_api_abuse"`    // 可疑账号拦截
	Transferable       bool `json:"transferable"`         // 是否允许转赠
	Shareable          bool `json:"shareable"`            // 是否允许分享链接
}

type CustomEntrance struct {
	MiniProgramsInfo *MiniProgramsInfo `json:"mini_programs_info"` // 小程序入口
	Appid            string            `json:"appid"`              // 商户公众号appid
	HallId           string            `json:"hall_id"`            // 营销馆id
	StoreId          string            `json:"store_id"`           // 可用门店id
	CodeDisplayMode  string            `json:"code_display_mode"`  // code展示模式
}

type MiniProgramsInfo struct {
	MiniProgramsAppid string `json:"mini_programs_appid"` // 商家小程序appid
	MiniProgramsPath  string `json:"mini_programs_path"`  // 商家小程序path
	EntranceWords     string `json:"entrance_words"`      // 入口文案
	GuidingWords      string `json:"guiding_words"`       // 引导文案
}

type DisplayPatternInfo struct {
	Description     string `json:"description,omitempty"` // 使用说明
	MerchantLogoUrl string `json:"merchant_logo_url"`     // 商户logo
	MerchantName    string `json:"merchant_name"`         // 商户名称
	BackgroundColor string `json:"background_color"`      // 背景颜色
	CouponImageUrl  string `json:"coupon_image_url"`      // 券详情图片
}

type CouponCodeCount struct {
	TotalCount     int `json:"total_count"`     // 该批次总共已上传的code总数
	AvailableCount int `json:"available_count"` // 该批次当前可用的code数
}

type NotifyConfig struct {
	NotifyAppid string `json:"notify_appid"` // 事件通知appid
}

type SendCountInformation struct {
	TotalSendNum    int `json:"total_send_num"`    // 已发放券张数
	TotalSendAmount int `json:"total_send_amount"` // 已发放券金额
	TodaySendNum    int `json:"today_send_num"`    // 单天已发放券张数
	TodaySendAmount int `json:"today_send_amount"` // 单天已发放券金额
}

type BusiFavorUse struct {
	StockId          string `json:"stock_id"`           // 批次号
	Openid           string `json:"openid"`             // 用户在公众号内的唯一身份标识
	WechatpayUseTime string `json:"wechatpay_use_time"` // 系统成功核销券的时间，遵循rfc3339标准
}

type BusiFavorUserCoupons struct {
	Data       []*BusiUserCoupon `json:"data,omitempty"` // 结果集
	TotalCount int               `json:"total_count"`    // 批次总数
	Offset     int               `json:"offset"`         // 分页页码
	Limit      int               `json:"limit"`          // 分页大小
}

type BusiUserCoupon struct {
	BelongMerchant     string              `json:"belong_merchant"`                // 批次归属商户号
	StockName          string              `json:"stock_name"`                     // 批次名称
	Comment            string              `json:"comment"`                        // 批次备注
	GoodsName          string              `json:"goods_name"`                     // 适用商品范围
	StockType          string              `json:"stock_type"`                     // 批次类型
	Transferable       bool                `json:"transferable"`                   // 是否允许转赠
	Shareable          bool                `json:"shareable"`                      // 是否允许分享链接
	CouponState        string              `json:"coupon_state"`                   // 商家券状态
	DisplayPatternInfo *DisplayPatternInfo `json:"display_pattern_info,omitempty"` // 创建批次时的样式信息
	CouponUseRule      *CouponUseRule      `json:"coupon_use_rule"`                // 券核销相关规则
	CustomEntrance     *CustomEntrance     `json:"custom_entrance,omitempty"`      // 自定义入口
	CouponCode         string              `json:"coupon_code"`                    // 券的唯一标识
	StockId            string              `json:"stock_id"`                       // 批次号
	AvailableStartTime string              `json:"available_start_time"`           // 券可使用开始时间
	ExpireTime         string              `json:"expire_time"`                    // 券过期时间
	ReceiveTime        string              `json:"receive_time"`                   // 券领券时间
	SendRequestNo      string              `json:"send_request_no"`                // 发券请求单号
	UseRequestNo       string              `json:"use_request_no"`                 // 核销请求单号
	UseTime            string              `json:"use_time"`                       // 券核销时间
}

type BusiFavorCodeUpload struct {
	StockId        string      `json:"stock_id"`        // 批次号
	TotalCount     int         `json:"total_count"`     // 去重后上传code总数
	SuccessCount   int         `json:"success_count"`   // 上传成功code个数
	SuccessCodes   []string    `json:"success_codes"`   // 上传成功的code列表
	SuccessTime    string      `json:"success_time"`    // 上传成功时间
	FailCount      int         `json:"fail_count"`      // 上传失败code个数
	FailCodes      []*FailCode `json:"fail_codes"`      // 上传失败的code及原因
	ExistCodes     []string    `json:"exist_codes"`     // 已存在的code列表
	DuplicateCodes []string    `json:"duplicate_codes"` // 本次请求中重复的code列表
}

type FailCode struct {
	CouponCode string `json:"coupon_code"` // 上传失败的券code
	Code       string `json:"code"`        // 上传失败错误码
	Message    string `json:"message"`     // 上传失败错误信息
}

type BusiFavorCallbackUrlSet struct {
	UpdateTime string `json:"update_time"` // 修改时间
	NotifyUrl  string `json:"notify_url"`  // 通知URL地址
	Mchid      string `json:"mchid"`       // 商户号
}

type BusiFavorCallbackUrl struct {
	NotifyUrl string `json:"notify_url"` // 通知URL地址
	Mchid     string `json:"mchid"`      // 商户号
}

type BusiFavorAssociate struct {
	WechatpayAssociateTime string `json:"wechatpay_associate_time"` // 关联成功时间
}

type BusiFavorDisassociate struct {
	WechatpayDisassociateTime string `json:"wechatpay_disassociate_time"` // 取消关联成功时间
}

type BusiFavorBatchUpdate struct {
	MaxCoupons      int `json:"max_coupons"`        // 批次当前最大发放个数
	MaxCouponsByDay int `json:"max_coupons_by_day"` // 当前单天发放上限个数
}

type BusiFavorSend struct {
	CardCode string `json:"card_code"` // 消费卡code
}

type BusiFavorReturn struct {
	WechatpayReturnTime string `json:"wechatpay_return_time"` // 微信退券成功的时间
}

type BusiFavorDeactivate struct {
	WechatpayDeactivateTime string `json:"wechatpay_deactivate_time"` // 券成功失效的时间
}

type BusiFavorSubsidyPay struct {
	SubsidyReceiptId string `json:"subsidy_receipt_id"` // 补差付款单号
	StockId          string `json:"stock_id"`           // 商家券批次号
	CouponCode       string `json:"coupon_code"`        // 券的唯一标识
	TransactionId    string `json:"transaction_id"`     // 微信支付系统生成的订单号
	PayerMerchant    string `json:"payer_merchant"`     // 营销补差扣款商户号
	PayeeMerchant    string `json:"payee_merchant"`     // 营销补差入账商户号
	Amount           int    `json:"amount"`             // 补差付款金额
	Description      string `json:"description"`        // 补差付款描述
	Status           string `json:"status"`             // 补差付款单据状态
	FailReason       string `json:"fail_reason"`        // 补差付款失败原因
	SuccessTime      string `json:"success_time"`       // 补差付款完成时间
	OutSubsidyNo     string `json:"out_subsidy_no"`     // 业务请求唯一单号
	CreateTime       string `json:"create_time"`        // 补差付款发起时间
}

type MarketMediaUpload struct {
	MediaUrl string `json:"media_url"` // 微信返回的媒体文件URL地址
}

type PartnershipsBuild struct {
	Partner        *Partner        `json:"partner"`         // 合作方相关的信息
	AuthorizedData *AuthorizedData `json:"authorized_data"` // 被授权的数据
	State          string          `json:"state"`           // 合作状态，ESTABLISHED：已建立，TERMINATED：已终止
	BuildTime      string          `json:"build_time"`      // 建立合作关系时间
	CreateTime     string          `json:"create_time"`     // 创建时间
	UpdateTime     string          `json:"update_time"`     // 更新时间
}

type Partner struct {
	Appid      string `json:"appid"`       // 合作方APPID
	Type       string `json:"type"`        // 合作方类别
	MerchantId string `json:"merchant_id"` // 合作方商户Id
}

type AuthorizedData struct {
	BusinessType string   `json:"business_type"` // 授权业务类别
	Scenarios    []string `json:"scenarios"`     // 授权场景
	StockId      string   `json:"stock_id"`      // 授权批次Id
}

type PartnershipsTerminate struct {
	TerminateTime string `json:"terminate_time"` // 终止合作关系时间
}

type PartnershipsList struct {
	Data       []*Partnerships `json:"data,omitempty"` // 合作关系结果集
	TotalCount int             `json:"total_count"`    // 批次总数
	Offset     int             `json:"offset"`         // 分页页码
	Limit      int             `json:"limit"`          // 分页大小
}

type Partnerships struct {
	Partner        *Partner        `json:"partner"`         // 合作方相关的信息
	AuthorizedData *AuthorizedData `json:"authorized_data"` // 被授权的数据
	BuildTime      string          `json:"build_time"`      // 建立合作关系时间
	TerminateTime  string          `json:"terminate_time"`  // 终止合作关系时间
	CreateTime     string          `json:"create_time"`     // 创建时间
	UpdateTime     string          `json:"update_time"`     // 更新时间
}

type EcommerceApply struct {
	ApplymentId  int64  `json:"applyment_id"`
	OutRequestNo string `json:"out_request_no"`
}

type EcommerceApplyStatus struct {
	ApplymentState     string            `json:"applyment_state"`
	ApplymentStateDesc string            `json:"applyment_state_desc"`
	SignState          string            `json:"sign_state,omitempty"`
	SignUrl            string            `json:"sign_url,omitempty"`
	SubMchid           string            `json:"sub_mchid,omitempty"`
	AccountValidation  AccountValidation `json:"account_validation"`
	AuditDetail        []*AuditDetail    `json:"audit_detail,omitempty"`
	LegalValidationUrl string            `json:"legal_validation_url,omitempty"`
	OutRequestNo       string            `json:"out_request_no"`
	ApplymentId        int64             `json:"applyment_id"`
}

type AccountValidation struct {
	AccountName              string `json:"account_name"`
	AccountNo                string `json:"account_no,omitempty"`
	PayAmount                int    `json:"pay_amount"`
	DestinationAccountNumber string `json:"destination_account_number"`
	DestinationAccountName   string `json:"destination_account_name"`
	DestinationAccountBank   string `json:"destination_account_bank"`
	City                     string `json:"city"`
	Remark                   string `json:"remark"`
	Deadline                 string `json:"deadline"`
}

type AuditDetail struct {
	ParamName    string `json:"param_name"`
	RejectReason string `json:"reject_reason"`
}

type GoldPlanManage struct {
	SubMchid string `json:"sub_mchid"`
}

type Withdraw struct {
	SubMchid    string `json:"sub_mchid"`              // 服务商特约商户号，由微信支付生成并下发。
	WithdrawId  string `json:"withdraw_id"`            // 微信支付提现单号
	AccountType string `json:"account_type,omitempty"` // 出款账户类型
}

type WithdrawStatus struct {
	SubMchid      string `json:"sub_mchid"`      // 服务商特约商户号，由微信支付生成并下发。
	SpMchid       string `json:"sp_mchid"`       // 服务商户号
	Status        string `json:"status"`         // 提现单状态：CREATE_SUCCESS：受理成功，SUCCESS：提现成功，FAIL：提现失败，REFUND：提现退票，CLOSE：关单，INIT：业务单已创建
	WithdrawId    string `json:"withdraw_id"`    // 微信支付提现单号
	OutRequestNo  string `json:"out_request_no"` // 商户提现单号
	Amount        int    `json:"amount"`         // 提现金额
	CreateTime    string `json:"create_time"`    // 创建时间
	UpdateTime    string `json:"update_time"`    // 更新时间
	Reason        string `json:"reason"`         // 提现失败原因，仅在提现失败、退票、关单时有值
	Remark        string `json:"remark"`         // 商户对提现单的备注，若发起提现时未传入相应值或输入不合法，则该值为空
	BankMemo      string `json:"bank_memo"`      // 展示在收款银行系统中的附言，由数字、字母、汉字组成（能否成功展示依赖银行系统支持）。若发起提现时未传入相应值或输入不合法，则该值为空
	AccountType   string `json:"account_type"`   // 出款账户类型
	AccountNumber string `json:"account_number"` // 服务商提现入账的银行账号，仅显示后四位。
	AccountBank   string `json:"account_bank"`   // 服务商提现入账的开户银行
	BankName      string `json:"bank_name"`      // 服务商提现入账的开户银行全称（含支行）
}

type EcommerceWithdraw struct {
	WithdrawId   string `json:"withdraw_id"`    // 微信支付预约提现单号
	OutRequestNo string `json:"out_request_no"` // 商户预约提现单号
}

type EcommerceWithdrawStatus struct {
	Status        string `json:"status"`         // 提现单状态：CREATE_SUCCESS：受理成功，SUCCESS：提现成功，FAIL：提现失败，REFUND：提现退票，CLOSE：关单，INIT：业务单已创建
	WithdrawId    string `json:"withdraw_id"`    // 微信支付提现单号
	OutRequestNo  string `json:"out_request_no"` // 商户提现单号
	Amount        int    `json:"amount"`         // 提现金额
	CreateTime    string `json:"create_time"`    // 创建时间
	UpdateTime    string `json:"update_time"`    // 更新时间
	Reason        string `json:"reason"`         // 提现失败原因，仅在提现失败、退票、关单时有值
	Remark        string `json:"remark"`         // 商户对提现单的备注，若发起提现时未传入相应值或输入不合法，则该值为空
	BankMemo      string `json:"bank_memo"`      // 展示在收款银行系统中的附言，由数字、字母、汉字组成（能否成功展示依赖银行系统支持）。若发起提现时未传入相应值或输入不合法，则该值为空
	AccountType   string `json:"account_type"`   // 出款账户类型
	AccountNumber string `json:"account_number"` // 服务商提现入账的银行账号，仅显示后四位。
	AccountBank   string `json:"account_bank"`   // 服务商提现入账的开户银行
	BankName      string `json:"bank_name"`      // 服务商提现入账的开户银行全称（含支行）
	Solution      string `json:"solution"`       // 提现失败解决方案，仅在提现失败、退票、关单时有值
}

type EcommerceProfitShare struct {
	SubMchid      string               `json:"sub_mchid"`           // 二级商户号
	TransactionId string               `json:"transaction_id"`      // 微信订单号
	OutOrderNo    string               `json:"out_order_no"`        // 商户分账单号
	OrderId       string               `json:"order_id"`            // 微信分账单号
	Status        string               `json:"status"`              // 分账单状态:PROCESSING：处理中,FINISHED：分账完成
	Receivers     []*EcommerceReceiver `json:"receivers,omitempty"` // 分账接收方列表
}

type EcommerceReceiver struct {
	Amount          int    `json:"amount"`                // 分账金额
	Description     string `json:"description"`           // 分账描述
	Type            string `json:"type"`                  // 分账接收方类型
	ReceiverAccount string `json:"receiver_account"`      // 分账接收方帐号
	ReceiverMchid   string `json:"receiver_mchid"`        // 分账接收商户号
	Result          string `json:"result"`                // 分账结果,PENDING：待分账,SUCCESS：分账成功,CLOSED：已关闭
	DetailId        string `json:"detail_id"`             // 分账明细单号
	FailReason      string `json:"fail_reason,omitempty"` // 分账失败原因ACCOUNT_ABNORMAL : 分账接收账户异常、NO_RELATION : 分账关系已解除、RECEIVER_HIGH_RISK : 高风险接收方、RECEIVER_REAL_NAME_NOT_VERIFIED : 接收方未实名、NO_AUTH : 分账权限已解除
	FinishTime      string `json:"finish_time"`           // 分账完成时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE
}

type EcommerceProfitShareQuery struct {
	SubMchid          string               `json:"sub_mchid"`                    // 二级商户号
	TransactionId     string               `json:"transaction_id"`               // 微信订单号
	OutOrderNo        string               `json:"out_order_no"`                 // 商户分账单号
	OrderId           string               `json:"order_id"`                     // 微信分账单号
	Status            string               `json:"status"`                       // 分账单状态:PROCESSING：处理中,FINISHED：分账完成
	Receivers         []*EcommerceReceiver `json:"receivers,omitempty"`          // 分账接收方列表
	FinishAmount      int                  `json:"finish_amount,omitempty"`      // 分账完结的分账金额，单位为分，仅当查询分账完结的执行结果时，存在本字段。
	FinishDescription string               `json:"finish_description,omitempty"` // 分账完结的原因描述，仅当查询分账完结的执行结果时，存在本字段。
}

type EcommerceProfitShareReturn struct {
	SubMchid    string `json:"sub_mchid"`             // 二级商户号
	OrderId     string `json:"order_id"`              // 微信分账单号，微信系统返回的唯一标识
	OutOrderNo  string `json:"out_order_no"`          // 商户分账单号
	OutReturnNo string `json:"out_return_no"`         // 商户回退单号
	ReturnNo    string `json:"return_no"`             // 微信分账回退单号，微信支付系统返回的唯一标识。
	ReturnMchid string `json:"return_mchid"`          // 只能对原分账请求中成功分给商户接收方进行回退
	Amount      int    `json:"amount"`                // 需要从分账接收方回退的金额，单位为分，只能为整数
	Result      string `json:"result"`                // 回退结果: PROCESSING：处理中，SUCCESS：已成功，FAILED：已失败
	FailReason  string `json:"fail_reason,omitempty"` // 失败原因: ACCOUNT_ABNORMAL : 分账接收方账户异常，TIME_OUT_CLOSED : 超时关单
	FinishTime  string `json:"finish_time"`           // 分账回退完成时间
}

type EcommerceProfitShareFinish struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionId string `json:"transaction_id"` // 微信订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	OrderId       string `json:"order_id"`       // 微信分账单号
}

type EcommerceProfitShareUnsplitAmount struct {
	TransactionId string `json:"transaction_id"` // 微信订单号
	UnsplitAmount int    `json:"unsplit_amount"` // 订单剩余待分金额
}

type EcommerceProfitShareReceiver struct {
	Type    string `json:"type"`    // 接收方类型:MERCHANT_ID：商户,PERSONAL_OPENID：个人
	Account string `json:"account"` // 分账接收方帐号
}

type EcommerceSubsidies struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionId string `json:"transaction_id"` // 微信订单号
	SubsidyId     string `json:"subsidy_id"`     // 微信补差单号
	Description   string `json:"description"`    // 补差描述
	Amount        int    `json:"amount"`         // 补差金额
	Result        string `json:"result"`         // 补差单结果
	SuccessTime   string `json:"success_time"`   // 补差完成时间
}

type EcommerceSubsidiesReturn struct {
	SubMchid        string `json:"sub_mchid"`         // 二级商户号
	TransactionId   string `json:"transaction_id"`    // 微信订单号
	SubsidyRefundId string `json:"subsidy_refund_id"` // 微信补差回退单号
	RefundId        string `json:"refund_id"`         // 微信退款单号
	OutOrderNo      string `json:"out_order_no"`      // 商户补差回退单号
	Amount          int    `json:"amount"`            // 补差回退金额
	Description     string `json:"description"`       // 补差回退描述
	Result          string `json:"result"`            // 补差回退结果
	SuccessTime     string `json:"success_time"`      // 补差回退完成时间
}

type EcommerceSubsidiesCancel struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionId string `json:"transaction_id"` // 微信订单号
	Result        string `json:"result"`         // 取消补差结果
	Description   string `json:"description"`    // 取消补差描述
}

type EcommerceRefund struct {
	RefundId        string                 `json:"refund_id"`        // 微信支付退款号
	OutRefundNo     string                 `json:"out_refund_no"`    // 商户退款单号
	CreateTime      string                 `json:"create_time"`      // 退款创建时间
	RefundAccount   string                 `json:"refund_account"`   // 退款资金来源
	Amount          *EcommerceRefundAmount `json:"amount"`           // 金额信息
	PromotionDetail []*PromotionDetailItem `json:"promotion_detail"` // 优惠退款信息
}

type EcommerceRefundAmount struct {
	Refund         int    `json:"refund"`          // 退款金额
	PayerRefund    int    `json:"payer_refund"`    // 用户退款金额
	DiscountRefund int    `json:"discount_refund"` // 优惠退款金额
	Currency       string `json:"currency"`        // 退款币种
}

type PromotionDetailItem struct {
	PromotionId  string `json:"promotion_id"`  // 券Id，券或立减金额
	Scope        string `json:"scope"`         // 优惠范围，GLOBAL：全场代金券，SINGLE：单品优惠
	Type         string `json:"type"`          // 优惠类型，COUPON：代金券，DISCOUNT：优惠券
	Amount       int    `json:"amount"`        // 优惠券面额，用户享受优惠的金额（优惠券面额=微信出资金额+商家出资金额+其他出资方金额），单位为分
	RefundAmount int    `json:"refund_amount"` // 优惠退款金额，单位为分
}

type EcommerceRefundQuery struct {
	RefundId            string                 `json:"refund_id"`                // 微信支付退款号
	OutRefundNo         string                 `json:"out_refund_no"`            // 商户退款单号
	TransactionId       string                 `json:"transaction_id"`           // 微信支付系统生成的订单号
	OutTradeNo          string                 `json:"out_trade_no"`             // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	Channel             string                 `json:"channel"`                  // 退款渠道
	UserReceivedAccount string                 `json:"user_received_account"`    // 退款入账账户
	SuccessTime         string                 `json:"success_time"`             // 退款成功时间
	CreateTime          string                 `json:"create_time"`              // 退款创建时间
	Status              string                 `json:"status"`                   // 退款状态
	RefundAccount       string                 `json:"refund_account,omitempty"` // 退款出资商户
	FundsAccount        string                 `json:"funds_account"`            // 资金账户
	Amount              *EcommerceRefundAmount `json:"amount"`                   // 金额信息
	PromotionDetail     []*PromotionDetailItem `json:"promotion_detail"`         // 优惠退款信息
}

type EcommerceRefundAdvance struct {
	RefundId        string `json:"refund_id"`         // 微信支付退款号
	AdvanceReturnId string `json:"advance_return_id"` // 微信回补单号
	ReturnAmount    int    `json:"return_amount"`     // 垫付回补金额
	PayerMchid      string `json:"payer_mchid"`       // 出款方商户号
	PayerAccount    string `json:"payer_account"`     // 出款方账户
	PayeeMchid      string `json:"payee_mchid"`       // 入账方商户号
	PayeeAccount    string `json:"payee_account"`     // 入账方账户
	Result          string `json:"result"`            // 垫付回补结果
	SuccessTime     string `json:"success_time"`      // 垫付回补完成时间
}

type BankSearchBank struct {
	TotalCount int         `json:"total_count"`    // 查询数据总条数
	Data       []*BankInfo `json:"data,omitempty"` // 银行列表
}

type BankInfo struct {
	BankAlias       string `json:"bank_alias"`        // 银行别名
	BankAliasCode   string `json:"bank_alias_code"`   // 银行别名编码
	AccountBank     string `json:"account_bank"`      // 开户银行
	AccountBankCode int    `json:"account_bank_code"` // 开户银行编码
	NeedBankBranch  bool   `json:"need_bank_branch"`  // 是否需要填写支行
}

type BankSearchList struct {
	TotalCount int         `json:"total_count"`    // 查询数据总条数
	Count      int         `json:"count"`          // 本次查询数据条数
	Offset     int         `json:"offset"`         // 本次查询偏移量
	Data       []*BankInfo `json:"data,omitempty"` // 银行列表
	Links      *Link       `json:"links"`          // 分页链接
}

type BankSearchProvince struct {
	TotalCount int             `json:"total_count"`    // 查询数据总条数
	Data       []*ProvinceInfo `json:"data,omitempty"` // 省份列表
}

type ProvinceInfo struct {
	ProvinceName string `json:"province_name"` // 省份名称
	ProvinceCode int    `json:"province_code"` // 省份编码
}

type BankSearchCity struct {
	TotalCount int         `json:"total_count"`    // 查询数据总条数
	Data       []*CityInfo `json:"data,omitempty"` // 城市列表
}

type CityInfo struct {
	CityName string `json:"city_name"` // 城市名称
	CityCode int    `json:"city_code"` // 城市编码
}

type BankSearchBranch struct {
	TotalCount      int               `json:"total_count"`       // 查询数据总条数
	Count           int               `json:"count"`             // 本次查询数据条数
	Offset          int               `json:"offset"`            // 本次查询偏移量
	BankAlias       string            `json:"bank_alias"`        // 银行别名
	BankAliasCode   string            `json:"bank_alias_code"`   // 银行别名编码
	AccountBank     string            `json:"account_bank"`      // 开户银行
	AccountBankCode int               `json:"account_bank_code"` // 开户银行编码
	Data            []*BankBranchInfo `json:"data,omitempty"`    // 支行列表
	Links           *Link             `json:"links"`             // 分页链接
}

type BankBranchInfo struct {
	BankBranchName string `json:"bank_branch_name"` // 开户银行支行名称
	BankBranchId   string `json:"bank_branch_id"`   // 开户银行支行联行号
}
