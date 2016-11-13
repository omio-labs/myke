package core

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/tidwall/gjson"
)

var _ = Describe("Task", func() {

	Describe("loadTaskJson", func() {
		It("Empty", func() {
			t := loadTaskJson("", gjson.Parse("{}"))
			Expect(t.Name).To(BeEmpty())
			Expect(t.Desc).To(BeEmpty())
			Expect(t.Cmd).To(BeEmpty())
			Expect(t.Before).To(BeEmpty())
			Expect(t.After).To(BeEmpty())
		})

		It("Name", func() {
			t := loadTaskJson("task", gjson.Parse(`{}`))
			Expect(t.Name).To(Equal("task"))
		})

		It("Desc", func() {
			t := loadTaskJson("", gjson.Parse(`{ "desc": "example" }`))
			Expect(t.Desc).To(Equal("example"))
		})

		It("Cmd", func() {
			t := loadTaskJson("", gjson.Parse(`{ "cmd": "echo" }`))
			Expect(t.Cmd).To(Equal("echo"))
		})

		It("Before", func() {
			t := loadTaskJson("", gjson.Parse(`{ "before": ["1", "2"] }`))
			Expect(t.Before).To(Equal([]string{"1", "2"}))
		})

		It("After", func() {
			t := loadTaskJson("", gjson.Parse(`{ "after": ["1", "2"] }`))
			Expect(t.After).To(Equal([]string{"1", "2"}))
		})
	})

})
