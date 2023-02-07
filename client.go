package coinbasegoclientv3

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

//go:generate mockgen -destination=./mocks/Client.go -package=mock_client coinbase-go-client-v3 Client
type Client interface {
	// IsActive tells you if the api is up
	IsActive(ctx context.Context) (bool, error)
	ListAccounts(ctx context.Context) ([]*Account, error)
	GetAccount(ctx context.Context, uuidStr string) (*Account, error)
	ListProducts(ctx context.Context) ([]*Product, error)
	GetProduct(ctx context.Context, uuidStr string) (*Product, error)
	GetProductCandles(
		ctx context.Context,
		uuidStr string,
		startTimeInUnixTime string,
		endTimeInUnixTime string,
		gran Granularity,
	) ([]*ProductCandle, error)
	GetMarketTrades(ctx context.Context, uuidStr string, numOfTradesToReturn int) ([]*MarketTrade, error)
	GetTransactionHistory(ctx context.Context, startTimeInRFC3339NanoTime string,
		endTimeInRFC3339NanoTime string, userNativeCurrency string, typ ProductType) (*TransactionSummary, error)
	// CreateOrder not tested yet...
	CreateOrder(ctx context.Context, createBody *CreateOrderBody) (newOrd *CreateOrderData, err error)
	// CancelOrders not tested yet...
	CancelOrders(ctx context.Context, orderIDs []string) error
	// GetOrder not tested yet...
	GetOrder(ctx context.Context, orderID string) (*Order, error)
	// ListOrders doesn't work
	// 		https://forums.coinbasecloud.dev/t/listorders-error-orderexecutionstatus/2699
	ListOrders(
		ctx context.Context,
		productID string,
		startDateInUnixTime string,
		userNativeCurrency string,
		orderType OrderType,
		productType ProductType,
		opts *ListOrdersOpts,
	) (data *ListOrderData, err error)
	// ListFills not tested yet...
	ListFills(
		ctx context.Context, orderID, productID,
		startSequenceTimestampInUnixTime, endSequenceTimestampInUnixTime string,
		opts *ListFillsOpts) (*ListFillsData, error)
}

func NewClient(
	httpClient HTTPClient, hostURI, coinbaseKey, coinbaseSecret string,
) (Client, error) {
	if len(hostURI) == 0 {
		hostURI = "https://api.coinbase.com"
	}

	return &client{
		httpClient:     httpClient,
		hostURI:        hostURI,
		coinbaseKey:    coinbaseKey,
		coinbaseSecret: coinbaseSecret,
	}, nil
}

type client struct {
	httpClient     HTTPClient
	hostURI        string
	coinbaseKey    string
	coinbaseSecret string
}

func (c *client) ListAccounts(ctx context.Context) ([]*Account, error) {
	accounts := []*Account{}
	var cursor string

	for {
		data := struct {
			Accounts []*Account `json:"accounts"`
			HasNext  bool       `json:"has_next"`
			Cursor   string     `json:"cursor"`
			Size     int        `json:"size"`
		}{}

		uri := "/api/v3/brokerage/accounts"
		params := map[string]string{
			"limit": "250",
		}
		if len(cursor) > 0 {
			params["cursor"] = cursor
		}

		res, err := c.makeRequest(ctx, GETHttpMethod, uri, params, nil)
		if err = c.handleErrorStatusCode(res, err); err != nil {
			return nil, err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(body, &data); err != nil {
			return nil, err
		}

		accounts = append(accounts, data.Accounts...)
		time.Sleep(time.Millisecond * 50) // The coinbase rate limit is 50 per a second so this should be well below

		if data.HasNext {
			cursor = data.Cursor
			continue
		}

		cursor = ""
		break
	}

	return accounts, nil
}

func (c *client) GetAccount(ctx context.Context, uuidStr string) (*Account, error) {
	data := struct {
		Account *Account `json:"account"`
	}{}

	uri := fmt.Sprintf("/api/v3/brokerage/accounts/%s/", uuidStr)

	res, err := c.makeRequest(ctx, GETHttpMethod, uri, nil, nil)
	if err = c.handleErrorStatusCode(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data.Account, nil
}

func (c *client) ListProducts(ctx context.Context) ([]*Product, error) {
	data := struct {
		Products    []*Product `json:"products"`
		NumProducts int        `json:"num_products"`
	}{}

	uri := "/api/v3/brokerage/products"
	// params := map[string]string{
	// 	//"limit": "5",
	// }
	// if len(cursor) > 0 {
	// 	params["cursor"] = cursor
	// }

	res, err := c.makeRequest(ctx, GETHttpMethod, uri, nil, nil)
	if err = c.handleErrorStatusCode(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data.Products, nil
}

func (c *client) GetProduct(ctx context.Context, uuidStr string) (*Product, error) {
	product := new(Product)

	uri := fmt.Sprintf("/api/v3/brokerage/products/%s/", uuidStr)

	res, err := c.makeRequest(ctx, GETHttpMethod, uri, nil, nil)
	if err = c.handleErrorStatusCode(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &product); err != nil {
		return nil, err
	}

	return product, nil
}

func (c *client) GetProductCandles(
	ctx context.Context,
	uuidStr string,
	startTimeInUnixTime string,
	endTimeInUnixTime string,
	gran Granularity,
) ([]*ProductCandle, error) {
	data := struct {
		Candles []*ProductCandle `json:"candles"`
	}{}

	uri := fmt.Sprintf("/api/v3/brokerage/products/%s/candles", uuidStr)

	params := map[string]string{
		"start":       startTimeInUnixTime,
		"end":         endTimeInUnixTime,
		"granularity": string(gran),
	}

	res, err := c.makeRequest(ctx, GETHttpMethod, uri, params, nil)
	if err = c.handleErrorStatusCode(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data.Candles, nil
}

func (c *client) GetMarketTrades(ctx context.Context, uuidStr string, numOfTradesToReturn int) ([]*MarketTrade, error) {
	data := struct {
		Trades []*MarketTrade `json:"trades"`
	}{}

	uri := fmt.Sprintf("/api/v3/brokerage/products/%s/ticker", uuidStr)

	params := map[string]string{
		"limit": fmt.Sprintf("%d", numOfTradesToReturn),
	}

	res, err := c.makeRequest(ctx, GETHttpMethod, uri, params, nil)
	if err = c.handleErrorStatusCode(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data.Trades, nil
}

func (c *client) GetTransactionHistory(ctx context.Context, startTimeInRFC3339NanoTime string,
	endTimeInRFC3339NanoTime string, userNativeCurrency string, typ ProductType) (*TransactionSummary, error) {
	ts := new(TransactionSummary)

	uri := "/api/v3/brokerage/transaction_summary"

	params := map[string]string{
		"start_date":           startTimeInRFC3339NanoTime,
		"end_date":             endTimeInRFC3339NanoTime,
		"user_native_currency": userNativeCurrency,
		"product_type":         string(typ),
	}

	res, err := c.makeRequest(ctx, GETHttpMethod, uri, params, nil)
	if err = c.handleErrorStatusCode(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, ts); err != nil {
		return nil, err
	}

	return ts, nil
}

func (c *client) CreateOrder(ctx context.Context, createBody *CreateOrderBody) (newOrd *CreateOrderData, err error) {
	newOrd = new(CreateOrderData)

	if createBody == nil {
		return nil, errors.New("create order body required and cannot be nil")
	}

	bodyBts, err := json.Marshal(createBody)
	if err != nil {
		return nil, err
	}

	uri := "/api/v3/brokerage/orders"

	res, err := c.makeRequest(ctx, POSTHttpMethod, uri, nil, bytes.NewBuffer(bodyBts))
	if err = c.handleErrorStatusCode(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(resBody, newOrd); err != nil {
		return nil, err
	}

	return newOrd, nil
}

func (c *client) CancelOrders(ctx context.Context, orderIDs []string) error {
	data := struct {
		Results []struct {
			Success       bool   `json:"success"`
			FailureReason string `json:"failure_reason"`
			OrderID       string `json:"order_id"`
		} `json:"results"`
	}{}

	body := struct {
		OrderIDs []string `json:"order_ids"`
	}{
		OrderIDs: orderIDs,
	}

	bodyBts, err := json.Marshal(body)
	if err != nil {
		return err
	}

	uri := "/api/v3/brokerage/orders/batch_cancel"

	res, err := c.makeRequest(ctx, POSTHttpMethod, uri, nil, bytes.NewBuffer(bodyBts))
	if err = c.handleErrorStatusCode(res, err); err != nil {
		return err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(resBody, &data); err != nil {
		return err
	}

	if len(data.Results) > 0 {
		if data.Results[0].Success {
			return nil
		}

		return errors.New(data.Results[0].FailureReason)
	}

	bts, err := json.Marshal(data.Results)
	if err != nil {
		return err
	}

	// This should never happen
	return errors.New("unknown error: " + string(bts))
}

func (c *client) GetOrder(ctx context.Context, orderID string) (ord *Order, err error) {
	data := struct {
		Order *Order `json:"order"`
	}{}

	uri := fmt.Sprintf("/api/v3/brokerage/orders/historical/%s", orderID)

	res, err := c.makeRequest(ctx, GETHttpMethod, uri, nil, nil)
	if err = c.handleErrorStatusCode(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data.Order, nil
}

func (c *client) ListOrders(
	ctx context.Context,
	productID string,
	startDateInUnixTime string,
	userNativeCurrency string,
	orderType OrderType,
	productType ProductType,
	opts *ListOrdersOpts,
) (data *ListOrderData, err error) {
	data = new(ListOrderData)

	uri := "/api/v3/brokerage/orders/historical/batch"

	params := map[string]string{
		"product_id":           productID,
		"start_date":           startDateInUnixTime,
		"user_native_currency": userNativeCurrency,
		"order_type":           string(orderType),
		"product_type":         string(productType),
	}

	if opts != nil {
		AssignStrIfSetToMap("order_status", StrArrToStr(opts.OrderStatus), params)
		AssignStrIfSetToMap("limit", fmt.Sprintf("%d", opts.Limit), params)
		AssignStrIfSetToMap("cursor", opts.Cursor, params)
		AssignStrIfSetToMap("end_date", opts.EndDateInUnixTime, params)
		AssignStrIfSetToMap("order_side", string(opts.OrderSide), params)
	}

	res, err := c.makeRequest(ctx, GETHttpMethod, uri, params, nil)
	if err = c.handleErrorStatusCode(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, data); err != nil {
		return nil, err
	}

	return data, nil
}

func (c *client) ListFills(
	ctx context.Context, orderID, productID,
	startSequenceTimestampInUnixTime, endSequenceTimestampInUnixTime string,
	opts *ListFillsOpts) (data *ListFillsData, err error) {
	data = new(ListFillsData)

	uri := "/api/v3/brokerage/orders/historical/batch"

	params := map[string]string{
		"order_id":                 orderID,
		"product_id":               productID,
		"start_sequence_timestamp": startSequenceTimestampInUnixTime,
		"end_sequence_timestamp":   endSequenceTimestampInUnixTime,
	}

	if opts != nil {
		AssignStrIfSetToMap("limit", fmt.Sprintf("%d", opts.Limit), params)
		AssignStrIfSetToMap("cursor", opts.Cursor, params)
	}

	res, err := c.makeRequest(ctx, GETHttpMethod, uri, params, nil)
	if err = c.handleErrorStatusCode(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, data); err != nil {
		return nil, err
	}

	return data, nil
}

func (c *client) IsActive(ctx context.Context) (bool, error) {
	res, err := c.makeRequest(ctx, GETHttpMethod, "/api/v3/brokerage/accounts", nil, nil)
	if err = c.handleErrorStatusCode(res, err); err != nil {
		return false, err
	}

	return true, nil
}

func (c *client) handleErrorStatusCode(res *http.Response, err error) error {
	if err != nil {
		return err
	} else if res.StatusCode/100 != 2 {
		if res.StatusCode == 400 {
			body, _ := io.ReadAll(res.Body)
			return fmt.Errorf("%s :%s", NewCoinbaseGoClientErr(res.StatusCode), string(body))
		}
		return NewCoinbaseGoClientErr(res.StatusCode)
	}

	return nil
}

func (c *client) makeRequest(ctx context.Context, method httpMethod, endpoint string, queryParams map[string]string, body io.Reader) (res *http.Response, err error) {
	url := fmt.Sprintf("%s%s", c.hostURI, endpoint)

	if queryParams != nil {
		qryParamSymbl := "?"
		for key, val := range queryParams {
			url += fmt.Sprintf("%s%s=%s", qryParamSymbl, key, val)
			if qryParamSymbl == "?" {
				qryParamSymbl = "&"
			}
		}
	}

	req, err := http.NewRequestWithContext(ctx, string(method), url, body)
	if err != nil {
		return nil, err
	}

	var bodyStr string

	if body != nil {
		bts, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		bodyStr = string(bts)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	ts := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	message := ts + string(method) + endpoint + bodyStr

	h := hmac.New(sha256.New, []byte(c.coinbaseSecret))
	h.Write([]byte(message))
	signature := hex.EncodeToString(h.Sum(nil))

	req.Header.Set("CB-ACCESS-KEY", c.coinbaseKey)
	req.Header.Set("CB-ACCESS-TIMESTAMP", ts)
	req.Header.Set("CB-ACCESS-SIGN", signature)

	return c.httpClient.Do(req)
}
