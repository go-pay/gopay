package paygateway

import (
	"errors"
	"fmt"
)

func formatCNYFenToYuanString(fen int64) (string, error) {
	if fen < 0 {
		return "", errors.New("amount must be >= 0")
	}
	return fmt.Sprintf("%d.%02d", fen/100, fen%100), nil
}
