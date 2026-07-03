package xomachine

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestXOmachine(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "XOmachine Suite")
}
