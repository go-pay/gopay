package paygateway

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func formatCNYFenToYuanString(fen int64) (string, error) {
	if fen < 0 {
		return "", errors.New("amount must be >= 0")
	}
	return fmt.Sprintf("%d.%02d", fen/100, fen%100), nil
}

func parseCNYYuanStringToFen(yuan string) (int64, error) {
	yuan = strings.TrimSpace(yuan)
	if yuan == "" {
		return 0, errors.New("amount is empty")
	}
	if strings.HasPrefix(yuan, "-") {
		return 0, errors.New("amount must be >= 0")
	}
	parts := strings.Split(yuan, ".")
	if len(parts) > 2 {
		return 0, errors.New("invalid amount format")
	}
	intPart := parts[0]
	if intPart == "" {
		intPart = "0"
	}
	i, err := strconv.ParseInt(intPart, 10, 64)
	if err != nil {
		return 0, errors.New("invalid amount format")
	}
	var frac int64
	if len(parts) == 2 {
		dec := parts[1]
		switch len(dec) {
		case 0:
			frac = 0
		case 1:
			d, err := strconv.ParseInt(dec, 10, 64)
			if err != nil {
				return 0, errors.New("invalid amount format")
			}
			frac = d * 10
		case 2:
			d, err := strconv.ParseInt(dec, 10, 64)
			if err != nil {
				return 0, errors.New("invalid amount format")
			}
			frac = d
		default:
			return 0, errors.New("invalid amount format")
		}
	}
	return i*100 + frac, nil
}
