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
	type args struct {
		errRsp ErrorResponse
	}
	tests := []struct {
		name           string
		args           args
		wantIsBizError bool
	}{
		{
			name: "",
			args: args{
				errRsp: ErrorResponse{
					Code: "40004",
					Msg:  "NOT_FOUND",
				},
			},
			wantIsBizError: true,
		},
		{
			name: "",
			args: args{
				errRsp: ErrorResponse{
					Code: "10000",
					Msg:  "SUCCEED",
				},
			},
			wantIsBizError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := bizErrCheck(tt.args.errRsp)
			_, ok := IsBizError(err)
			if ok != tt.wantIsBizError {
				t.Errorf("isBizError got = %v, want %v", ok, tt.wantIsBizError)
			}
		})
	}
}
