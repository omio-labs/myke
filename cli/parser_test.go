package cli

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	It("should pass", func() {
		res, err := Parse("/some/path")
		Expect(res).To(BeNil())
		Expect(err).ToNot(BeNil())
	})
})
