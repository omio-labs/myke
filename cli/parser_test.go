package cli

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"path/filepath"
)

var _ = Describe("Parser", func() {

	Describe("ParseFile", func() {
		It("Dir", func() {
			path, err := filepath.Abs("../examples/myke.yml")
			p, err := ParseFile("../examples")
			Expect(err).ToNot(HaveOccurred())
			Expect(p.Name).To(Equal("example"))
			Expect(p.Src).To(Equal(path))
		})

		It("Dummy", func() {
			p, err := ParseFile("../examples/extends")
			Expect(err).ToNot(HaveOccurred())
			Expect(p).ToNot(BeNil())
		})
	})

})
