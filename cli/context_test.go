package cli

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Context", func() {

	Describe("Discover", func() {
		It("examples", func() {
			c := make(chan Project)
			go Discover(".", "../examples", c)

			names := []string{}
			for p := range c {
				names = append(names, p.Name)
			}
			Expect(names).To(ConsistOf([]string{
				"example", "child", "depends", "env",
				"extends", "params", "tags1", "tags2",
			}))
		})
	})

})
