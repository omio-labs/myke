package core

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
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
			q2 := Query{Task:"task"}
			Expect(q1.Match(&p, &t)).To(BeTrue())
			Expect(q2.Match(&p, &t)).To(BeFalse())
		})

		It("task name glob match", func() {
			p := Project{}
			t := Task{Name:"task1"}
			q1 := Query{Task:"*task*"}
			q2 := Query{Task:"*2*"}
			Expect(q1.Match(&p, &t)).To(BeTrue())
			Expect(q2.Match(&p, &t)).To(BeFalse())
		})

		It("project match", func() {
			p := Project{Name:"project1"}
			t := Task{Name:"task"}
			q1 := Query{Task:"task", Tags:[]string{"project1"}}
			q2 := Query{Task:"task", Tags:[]string{"project2"}}
			Expect(q1.Match(&p, &t)).To(BeTrue())
			Expect(q2.Match(&p, &t)).To(BeFalse())
		})

		It("project glob match", func() {
			p := Project{Name:"project1"}
			t := Task{Name:"task"}
			q1 := Query{Task:"task", Tags:[]string{"*project*"}}
			q2 := Query{Task:"task", Tags:[]string{"*2*"}}
			Expect(q1.Match(&p, &t)).To(BeTrue())
			Expect(q2.Match(&p, &t)).To(BeFalse())
		})

		It("tags match", func() {
			p := Project{Name:"project", Tags:[]string{"tag1", "tag2", "tag3"}}
			t := Task{Name:"task"}
			q1 := Query{Task:"task", Tags:[]string{"tag1", "tag2"}}
			q2 := Query{Task:"task", Tags:[]string{"tag3", "tag4"}}
			Expect(q1.Match(&p, &t)).To(BeTrue())
			Expect(q2.Match(&p, &t)).To(BeFalse())
		})

		It("tags glob match", func() {
			p := Project{Name:"project", Tags:[]string{"tag1", "tag2", "tag3"}}
			t := Task{Name:"task"}
			q1 := Query{Task:"task", Tags:[]string{"*tag*"}}
			q2 := Query{Task:"task", Tags:[]string{"*tag*", "tag4"}}
			Expect(q1.Match(&p, &t)).To(BeTrue())
			Expect(q2.Match(&p, &t)).To(BeFalse())
		})

		It("tags and project match", func() {
			p := Project{Name:"project", Tags:[]string{"tag1", "tag2", "tag3"}}
			t := Task{Name:"task"}
			q1 := Query{Task:"task", Tags:[]string{"tag1", "tag2", "project"}}
			q2 := Query{Task:"task", Tags:[]string{"tag3", "tag4", "project"}}
			Expect(q1.Match(&p, &t)).To(BeTrue())
			Expect(q2.Match(&p, &t)).To(BeFalse())
		})

		It("tags and project glob match", func() {
			p := Project{Name:"project", Tags:[]string{"tag1", "tag2", "tag3"}}
			t := Task{Name:"task"}
			q1 := Query{Task:"*task*", Tags:[]string{"tag?", "project*"}}
			q2 := Query{Task:"*task*", Tags:[]string{"tag4", "project*"}}
			Expect(q1.Match(&p, &t)).To(BeTrue())
			Expect(q2.Match(&p, &t)).To(BeFalse())
		})
	})

	Describe("Search", func() {
		w := ParseWorkspace("../examples")

		It("Match All", func() {
			q, _ := ParseQuery("*")
			m := q.Search(&w)
			Expect(m).To(HaveLen(14))
		})

		It("example/build", func() {
			q, _ := ParseQuery("example/build")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("example/build"))
		})

		It("child/build", func() {
			q, _ := ParseQuery("child/build")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("child/build"))
		})

		It("env/env", func() {
			q, _ := ParseQuery("env/env")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("env/env"))
		})

		It("tags1/tag", func() {
			q, _ := ParseQuery("tags1/tag")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("tags1/tag"))
		})

		It("tags2/tag", func() {
			q, _ := ParseQuery("tags2/tag")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("tags2/tag"))
		})

		It("build", func() {
			q, _ := ParseQuery("build")
			m := q.Search(&w)
			Expect(m).To(HaveLen(2))
			Expect(fullNames(m)).To(ConsistOf("example/build", "child/build"))
		})

		It("test", func() {
			q, _ := ParseQuery("test")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("child/test"))
		})

		It("tagA/tag", func() {
			q, _ := ParseQuery("tagA/tag")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("tags1/tag"))
		})

		It("tagB/tag", func() {
			q, _ := ParseQuery("tagB/tag")
			m := q.Search(&w)
			Expect(m).To(HaveLen(2))
			Expect(fullNames(m)).To(ConsistOf("tags1/tag", "tags2/tag"))
		})

		It("tagC/tag", func() {
			q, _ := ParseQuery("tagC/tag")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("tags2/tag"))
		})

		It("depends/itself", func() {
			q, _ := ParseQuery("depends/itself")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("depends/itself"))
		})

		It("depends/before", func() {
			q, _ := ParseQuery("depends/before")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("depends/before"))
		})

		It("depends/after", func() {
			q, _ := ParseQuery("depends/after")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("depends/after"))
		})

		It("depends/before_after", func() {
			q, _ := ParseQuery("depends/before_after")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("depends/before_after"))
		})

		It("extends/task1", func() {
			q, _ := ParseQuery("extends/task1")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("extends/task1"))
		})

		It("extends/task2", func() {
			q, _ := ParseQuery("extends/task2")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("extends/task2"))
		})

		It("extends/task3", func() {
			q, _ := ParseQuery("extends/task3")
			m := q.Search(&w)
			Expect(m).To(HaveLen(1))
			Expect(fullNames(m)).To(ConsistOf("extends/task3"))
		})
	})

})

func fullNames(ms []QueryMatch) []string {
	s := make([]string, len(ms))
	for i, m := range ms {
		s[i] = fmt.Sprintf("%v/%v", m.Project.Name, m.Task.Name)
	}
	return s
}
