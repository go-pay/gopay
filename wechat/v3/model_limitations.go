package wechat

type LimitationsRsp struct {
	Code        int                 `json:"-"`
	SignInfo    *SignInfo           `json:"-"`
	Response    *LimitationsRspBody `json:"response,omitempty"`
	ErrResponse ErrResponse         `json:"err_response,omitempty"`
	Error       string              `json:"-"`
}

type LimitationsRspBody struct {
	Mchid                  string                   `json:"mchid"`                   // 商户号
	LimitedFunctions       []*string                `json:"limited_functions"`       // 受限功能列表
	OtherLimitedFunctions  string                   `json:"other_limited_functions"` // 其他受限功能描述
	RecoverySpecifications []*RecoverySpecification `json:"recovery_specifications"` // 恢复方案列表
}

type RecoverySpecification struct {
	LimitationCaseId         string   `json:"limitation_case_id"`         // 限制案例ID
	LimitationReasonType     string   `json:"limitation_reason_type"`     // 限制原因类型
	LimitationReason         string   `json:"limitation_reason"`          // 限制原因
	LimitationReasonDescribe string   `json:"limitation_reason_describe"` // 限制原因描述
	RelateLimitations        []string `json:"relate_limitations"`         // 相关受限功能列表
	OtherRelateLimitations   string   `json:"other_relate_limitations"`   // 其他相关受限功能描述
	RecoverWay               string   `json:"recover_way"`                // 恢复方式
	RecoverWayParam          string   `json:"recover_way_param"`          // 恢复方式参数
	RecoverHelpUrl           string   `json:"recover_help_url"`           // 恢复帮助URL
	LimitationActionType     string   `json:"limitation_action_type"`     // 限制措施类型
	LimitationStartDate      string   `json:"limitation_start_date"`      // 限制开始时间，遵循RFC3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2025-06-08T10:34:56+08:00
	LimitationDate           string   `json:"limitation_date"`            // 限制持续时间，遵循RFC3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2025-06-08T10:34:56+08:00
}
