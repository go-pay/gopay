package wechat

// 商圈积分授权查询 Rsp
type BusinessAuthPointsQueryRsp struct {
	Code     int                      `json:"-"`
	SignInfo *SignInfo                `json:"-"`
	Response *BusinessAuthPointsQuery `json:"response,omitempty"`
	Error    string                   `json:"-"`
}

// 商圈会员待积分状态查询 Rsp
type BusinessPointsStatusQueryRsp struct {
	Code     int                        `json:"-"`
	SignInfo *SignInfo                  `json:"-"`
	Response *BusinessPointsStatusQuery `json:"response,omitempty"`
	Error    string                     `json:"-"`
}

// =========================================================分割=========================================================

type BusinessAuthPointsQuery struct {
	Openid          string `json:"openid"`                     // 顾客授权时使用的小程序上的openid
	AuthorizeState  string `json:"authorize_state"`            // 顾客授权商圈积分结果：UNAUTHORIZED：未授权，AUTHORIZED：已授权，DEAUTHORIZED：已取消授权
	AuthorizeTime   string `json:"authorize_time,omitempty"`   // 顾客成功授权商圈积分的时间
	DeauthorizeTime string `json:"deauthorize_time,omitempty"` // 顾客关闭授权商圈积分的时间
}

type BusinessPointsStatusQuery struct {
	PointsCommitStatus string `json:"points_commit_status"` // 顾客关闭授权商圈积分的时间
}
