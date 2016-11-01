package cli

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Query", func() {

	Describe("ParseQuery", func() {
		It("basic", func() {
			q, err := ParseQuery("/task/[,]")
			Expect(err).ToNot(HaveOccurred())
			Expect(q.Task).To(Equal("task"))
			Expect(q.Tags).To(BeEmpty())
			Expect(q.Params).To(BeEmpty())
		})

		It("tag", func() {
			q, err := ParseQuery("/tag1/task/[,]")
			Expect(err).ToNot(HaveOccurred())
			Expect(q.Task).To(Equal("task"))
			Expect(q.Tags).To(ConsistOf("tag1"))
			Expect(q.Params).To(BeEmpty())
		})

		It("tags", func() {
			q, err := ParseQuery("/tag1/tag2/tag3/task/[,]")
			Expect(err).ToNot(HaveOccurred())
			Expect(q.Task).To(Equal("task"))
			Expect(q.Tags).To(ConsistOf("tag1", "tag2", "tag3"))
			Expect(q.Params).To(BeEmpty())
		})

		It("param", func() {
			q, err := ParseQuery("/tag1/task/[,a=1,]")
			Expect(err).ToNot(HaveOccurred())
			Expect(q.Task).To(Equal("task"))
			Expect(q.Tags).To(ConsistOf("tag1"))
			Expect(q.Params).To(HaveLen(1))
			Expect(q.Params["a"]).To(Equal("1"))
		})

		It("params", func() {
			q, err := ParseQuery("/tag1/task/[,a=1,,b=2,,c=3,]")
			Expect(err).ToNot(HaveOccurred())
			Expect(q.Task).To(Equal("task"))
			Expect(q.Tags).To(ConsistOf("tag1"))
			Expect(q.Params).To(HaveLen(3))
			Expect(q.Params["a"]).To(Equal("1"))
			Expect(q.Params["b"]).To(Equal("2"))
			Expect(q.Params["c"]).To(Equal("3"))
		})
	})

})
