package cli

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    . "github.com/MakeNowJust/heredoc/dot"
)

var _ = Describe("Parser", func() {

	yml := []byte(D(`
		name: testProject
	`))

	Describe("Project", func() {
		It("", func() {
			p := Project{}
			err := Parse(yml, &p)
			Expect(err).ToNot(HaveOccurred())
		})
	})

})
