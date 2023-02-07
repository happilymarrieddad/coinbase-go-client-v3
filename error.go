package coinbasegoclientv3

import (
	"fmt"
)

type coinbaseGoClientErr error

func NewCoinbaseGoClientErr(errorCode int) coinbaseGoClientErr {
	apiErr, ok := CoinbaseStatusCodesMap[errorCode]
	if !ok {
		return fmt.Errorf("%d: unreconized error", errorCode)
	}

	return fmt.Errorf("%d: %s", errorCode, apiErr)
}
