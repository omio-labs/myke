package cli

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tidwall/gjson"
)

var _ = Describe("Parser", func() {

	Describe("ParseProject", func() {
		It("Empty", func() {
			p := ParseProject(gjson.Parse("{}"))
			Expect(p.Name).To(BeEmpty())
			Expect(p.Desc).To(BeEmpty())
			Expect(p.Includes).To(BeEmpty())
			Expect(p.Includes).To(BeEmpty())
			Expect(p.Extends).To(BeEmpty())
			Expect(p.Env).To(BeEmpty())
			Expect(p.EnvFiles).To(BeEmpty())
			Expect(p.Tags).To(BeEmpty())
			Expect(p.Tasks).To(BeEmpty())
		})

		It("Name", func() {
			p := ParseProject(gjson.Parse(`{ "project": "example" }`))
			Expect(p.Name).To(Equal("example"))
		})

		It("Desc", func() {
			p := ParseProject(gjson.Parse(`{ "desc": "example" }`))
			Expect(p.Desc).To(Equal("example"))
		})

		It("Includes", func() {
			p := ParseProject(gjson.Parse(`{ "includes": ["1", "2"] }`))
			Expect(p.Includes).To(Equal([]string{"1", "2"}))
		})

		It("Extends", func() {
			p := ParseProject(gjson.Parse(`{ "extends": ["1", "2"] }`))
			Expect(p.Extends).To(Equal([]string{"1", "2"}))
		})

		It("Env", func() {
			p := ParseProject(gjson.Parse(`{ "env": { "a": "1", "b": "2", "c": "3" } }`))
			Expect(p.Env).To(HaveLen(3))
			Expect(p.Env["a"]).To(Equal("1"))
			Expect(p.Env["b"]).To(Equal("2"))
			Expect(p.Env["c"]).To(Equal("3"))
		})

		It("EnvFiles", func() {
			p := ParseProject(gjson.Parse(`{ "env_files": ["1", "2"] }`))
			Expect(p.EnvFiles).To(Equal([]string{"1", "2"}))
		})

		It("Tags", func() {
			p := ParseProject(gjson.Parse(`{ "tags": ["1", "2"] }`))
			Expect(p.Tags).To(Equal([]string{"1", "2"}))
		})

		Describe("Tasks", func() {
			It("None", func() {
				p := ParseProject(gjson.Parse(`{ "tasks": {} }`))
				Expect(p.Tasks).To(BeEmpty())
			})
			It("One", func() {
				p := ParseProject(gjson.Parse(`{ "tasks": { "test": {} } }`))
				Expect(p.Tasks).To(HaveLen(1))
				Expect(p.Tasks[0].Name).To(Equal("test"))
			})
			It("Two", func() {
				p := ParseProject(gjson.Parse(`{ "tasks": { "test1": {}, "test2": {} } }`))
				var taskNames []string
				for _, t := range p.Tasks {
					taskNames = append(taskNames, t.Name)
				}
				Expect(taskNames).To(ConsistOf("test1", "test2"))
			})
		})
	})

	Describe("ParseTask", func() {
		It("Empty", func() {
			t := ParseTask("", gjson.Parse("{}"))
			Expect(t.Name).To(BeEmpty())
			Expect(t.Desc).To(BeEmpty())
			Expect(t.Cmd).To(BeEmpty())
			Expect(t.Before).To(BeEmpty())
			Expect(t.After).To(BeEmpty())
		})

		It("Name", func() {
			t := ParseTask("task", gjson.Parse(`{}`))
			Expect(t.Name).To(Equal("task"))
		})

		It("Desc", func() {
			t := ParseTask("", gjson.Parse(`{ "desc": "example" }`))
			Expect(t.Desc).To(Equal("example"))
		})

		It("Cmd", func() {
			t := ParseTask("", gjson.Parse(`{ "cmd": "echo" }`))
			Expect(t.Cmd).To(Equal("echo"))
		})

		It("Before", func() {
			t := ParseTask("", gjson.Parse(`{ "before": ["1", "2"] }`))
			Expect(t.Before).To(Equal([]string{"1", "2"}))
		})

		It("After", func() {
			t := ParseTask("", gjson.Parse(`{ "after": ["1", "2"] }`))
			Expect(t.After).To(Equal([]string{"1", "2"}))
		})
	})

})
