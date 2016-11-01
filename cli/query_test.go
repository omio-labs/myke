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

		It("project", func() {
			q, err := ParseQuery("/project/task/[,]")
			Expect(err).ToNot(HaveOccurred())
			Expect(q.Task).To(Equal("task"))
			Expect(q.Tags).To(ConsistOf("project"))
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

	Describe("Match", func() {
		It("task name match", func() {
			p := Project{}
			t := Task{Name:"task1"}
			q1 := Query{Task:"task1"}
			q2 := Query{Task:"task2"}
			Expect(q1.Matches(p, t)).To(BeTrue())
			Expect(q2.Matches(p, t)).To(BeFalse())
		})

		It("project match", func() {
			p := Project{Name:"project1"}
			t := Task{Name:"task"}
			q1 := Query{Task:"task", Tags:[]string{"project1"}}
			q2 := Query{Task:"task", Tags:[]string{"project2"}}
			Expect(q1.Matches(p, t)).To(BeTrue())
			Expect(q2.Matches(p, t)).To(BeFalse())
		})

		It("tags match", func() {
			p := Project{Name:"project", Tags:[]string{"tag1", "tag2", "tag3"}}
			t := Task{Name:"task"}
			q1 := Query{Task:"task", Tags:[]string{"tag1", "tag2"}}
			q2 := Query{Task:"task", Tags:[]string{"tag3", "tag4"}}
			Expect(q1.Matches(p, t)).To(BeTrue())
			Expect(q2.Matches(p, t)).To(BeFalse())
		})

		It("tags and project match", func() {
			p := Project{Name:"project", Tags:[]string{"tag1", "tag2", "tag3"}}
			t := Task{Name:"task"}
			q1 := Query{Task:"task", Tags:[]string{"tag1", "tag2", "project"}}
			q2 := Query{Task:"task", Tags:[]string{"tag3", "tag4", "project"}}
			Expect(q1.Matches(p, t)).To(BeTrue())
			Expect(q2.Matches(p, t)).To(BeFalse())
		})
	})

})
