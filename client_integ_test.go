package coinbasegoclientv3_test

import (
	. "coinbase-go-client-v3"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {

	var (
		ctx        context.Context
		httpClient *http.Client

		client Client
	)

	BeforeEach(func() {
		ctx = context.Background()
		httpClient = &http.Client{
			Timeout: time.Second * 30,
		}
		Expect(httpClient).NotTo(BeNil())

		var err error
		client, err = NewClient(httpClient, "", os.Getenv("COINBASE_TEST_API_KEY"), os.Getenv("COINBASE_TEST_API_SECRET"))
		Expect(err).To(BeNil())
		Expect(client).NotTo(BeNil())
	})

	Context("IsActive", func() {
		It("IsActive", func() {
			isActive, err := client.IsActive(ctx)
			Expect(err).To(BeNil())
			Expect(isActive).To(BeTrue())
		})
	})

	Context("ListAccounts", func() {
		It("should successfully get the accounts", func() {
			accounts, err := client.ListAccounts(ctx)
			Expect(err).To(BeNil())
			Expect(accounts).NotTo(HaveLen(0))
		})
	})

	Context("GetAccount", func() {
		It("should successfully get the account", func() {
			accounts, err := client.ListAccounts(ctx)
			Expect(err).To(BeNil())
			Expect(accounts).NotTo(HaveLen(0))

			account, err := client.GetAccount(ctx, accounts[0].UUID.String())
			Expect(err).To(BeNil())
			Expect(account).NotTo(BeNil())

			Expect(account.Name).To(Equal(accounts[0].Name))
		})
	})

	Context("ListProducts", func() {
		It("should successfully get the products", func() {
			products, err := client.ListProducts(ctx)
			Expect(err).To(BeNil())
			Expect(products).NotTo(HaveLen(0))
		})
	})

	Context("GetProduct", func() {
		It("should successfully get the account", func() {
			products, err := client.ListProducts(ctx)
			Expect(err).To(BeNil())
			Expect(products).NotTo(HaveLen(0))

			product, err := client.GetProduct(ctx, products[0].ProductID)
			Expect(err).To(BeNil())
			Expect(product).NotTo(BeNil())

			Expect(product.BaseName).To(Equal(products[0].BaseName))
		})
	})

	Context("GetProductCandles", func() {
		It("should successfully get the product candles", func() {
			products, err := client.ListProducts(ctx)
			Expect(err).To(BeNil())
			Expect(products).NotTo(HaveLen(0))

			productCandles, err := client.GetProductCandles(
				ctx, products[0].ProductID,
				fmt.Sprintf("%d", time.Now().Add(time.Hour*-12).Unix()),
				fmt.Sprintf("%d", time.Now().Unix()), FiveMinuteGranularity)
			Expect(err).To(BeNil())
			Expect(productCandles).NotTo(HaveLen(0))
		})
	})

	Context("GetMarketTrades", func() {
		It("should successfully get the trades for a product", func() {
			products, err := client.ListProducts(ctx)
			Expect(err).To(BeNil())
			Expect(products).NotTo(HaveLen(0))

			trades, err := client.GetMarketTrades(ctx, products[0].ProductID, 10)
			Expect(err).To(BeNil())
			Expect(trades).NotTo(HaveLen(0))
			Expect(trades[0].ProductID).To(Equal(products[0].ProductID))
			/*
				TradeID: (string) (len=9) "496420852",
				ProductID: (string) (len=7) "BTC-USD",
				Price: (string) (len=8) "22878.64",
				Size: (string) (len=10) "0.00077741",
				Time: (string) (len=27) "2023-02-07T03:01:11.744790Z",
				Side: (coinbasegoclientv3.side) (len=3) "BUY",
				Bid: (string) "",
				Ask: (string) ""
			*/
		})
	})

})
