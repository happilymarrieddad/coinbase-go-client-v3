package coinbasegoclientv3

type coinbaseStatusCodes string

const (
	/* https://docs.cloud.coinbase.com/sign-in-with-coinbase/docs/status-codes */
	OKCoinbaseStatusCodes                  coinbaseStatusCodes = "OK Successful request"
	CreatedCoinbaseStatusCodes             coinbaseStatusCodes = "Created New object saved"
	NoContentCoinbaseStatusCodes           coinbaseStatusCodes = "No content Object deleted"
	BadRequestCoinbaseStatusCodes          coinbaseStatusCodes = "Bad Request Returns JSON with the error message"
	UnauthorizedCoinbaseStatusCodes        coinbaseStatusCodes = "Unauthorized Couldn't authenticate your request"
	TokenRequiredCoinbaseStatusCodes       coinbaseStatusCodes = "2FA Token required Re-try request with user’s 2FA token as CB-2FA-Token header"
	InvalidScopeCoinbaseStatusCodes        coinbaseStatusCodes = "Invalid scope User hasn't authorized necessary scope"
	NotFoundCoinbaseStatusCodes            coinbaseStatusCodes = "Not Found No such object"
	TooManyRequestsCoinbaseStatusCodes     coinbaseStatusCodes = "Too Many Requests Your connection is being rate limited"
	InternalServerErrorCoinbaseStatusCodes coinbaseStatusCodes = "Internal Server Error Something went wrong"
	ServiceUnavailableCoinbaseStatusCodes  coinbaseStatusCodes = "Service Unavailable Your connection is being throttled or the service is down for maintenance"
)

var CoinbaseStatusCodesMap map[int]coinbaseStatusCodes

func init() {
	CoinbaseStatusCodesMap = make(map[int]coinbaseStatusCodes)

	CoinbaseStatusCodesMap[200] = OKCoinbaseStatusCodes
	CoinbaseStatusCodesMap[201] = CreatedCoinbaseStatusCodes
	CoinbaseStatusCodesMap[204] = NoContentCoinbaseStatusCodes
	CoinbaseStatusCodesMap[400] = BadRequestCoinbaseStatusCodes
	CoinbaseStatusCodesMap[401] = UnauthorizedCoinbaseStatusCodes
	CoinbaseStatusCodesMap[402] = TokenRequiredCoinbaseStatusCodes
	CoinbaseStatusCodesMap[403] = InvalidScopeCoinbaseStatusCodes
	CoinbaseStatusCodesMap[404] = NotFoundCoinbaseStatusCodes
	CoinbaseStatusCodesMap[429] = TooManyRequestsCoinbaseStatusCodes
	CoinbaseStatusCodesMap[500] = InternalServerErrorCoinbaseStatusCodes
	CoinbaseStatusCodesMap[503] = ServiceUnavailableCoinbaseStatusCodes
}
