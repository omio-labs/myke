package cli

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Workspace", func() {

	Describe("LoadWorkspace", func() {
		It("examples", func() {
			w := LoadWorkspace(".", "../examples")
			names := []string{}

			for _, p := range w.Projects {
				names = append(names, p.Name)
			}
			Expect(names).To(ConsistOf([]string{
				"example", "child", "depends", "env",
				"extends", "params", "tags1", "tags2",
			}))
		})
	})

})
