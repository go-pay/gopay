package alipay

import (
	"testing"
)

func TestBizErr_BizErrCheck(t *testing.T) {
	bizErrRsp := ErrorResponse{
		Code: "40004",
		Msg:  "NOT_FOUND",
	}
	if bizErrCheck(bizErrRsp) == nil {
		t.Fail()
	}

	noBizErrRsp := ErrorResponse{
		Code: "10000",
		Msg:  "SUCCEED",
	}

	if bizErrCheck(noBizErrRsp) != nil {
		t.Fail()
	}
}

func TestBizErr_AsBizError(t *testing.T) {
	bizErrRsp := ErrorResponse{
		Code: "40004",
		Msg:  "NOT_FOUND",
	}
	noBizErrRsp := ErrorResponse{
		Code: "10000",
		Msg:  "SUCCEED",
	}
	var err error
	err = bizErrCheck(bizErrRsp)
	if _, ok := IsBizError(err); !ok {
		t.Fail()
	}

	err = bizErrCheck(noBizErrRsp)
	if _, ok := IsBizError(err); !ok {
		t.Fail()
	}
}
