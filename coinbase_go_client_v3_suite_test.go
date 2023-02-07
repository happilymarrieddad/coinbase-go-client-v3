package coinbasegoclientv3_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCoinbaseGoClientV3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CoinbaseGoClientV3 Suite")
}
