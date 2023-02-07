package coinbasegoclientv3_test

import (
	. "coinbase-go-client-v3"
	. "coinbase-go-client-v3/mocks"
	"os"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {

	var (
		ctrl       *gomock.Controller
		httpClient HTTPClient

		client Client
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		httpClient = NewMockHTTPClient(ctrl)
		Expect(httpClient).NotTo(BeNil())

		var err error
		client, err = NewClient(httpClient, os.Getenv("COINBASE_TEST_API_KEY"), os.Getenv("COINBASE_TEST_API_SECRET"))
		Expect(err).To(BeNil())
		Expect(client).NotTo(BeNil())
	})

	Context("TODO", func() {
		It("TODO", func() {

		})
	})

})
