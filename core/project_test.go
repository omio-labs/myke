package core

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tidwall/gjson"

	"os"
	"path/filepath"
	"strings"
)

var _ = Describe("Project", func() {

	Describe("loadProjectJson", func() {
		It("Empty", func() {
			p := loadProjectJson(gjson.Parse("{}"))
			Expect(p.Name).To(BeEmpty())
			Expect(p.Desc).To(BeEmpty())
			Expect(p.Discover).To(BeEmpty())
			Expect(p.Discover).To(BeEmpty())
			Expect(p.Mixin).To(BeEmpty())
			Expect(p.Env).To(BeEmpty())
			Expect(p.EnvFiles).To(BeEmpty())
			Expect(p.Tags).To(BeEmpty())
			Expect(p.Tasks).To(BeEmpty())
		})

		It("Name", func() {
			p := loadProjectJson(gjson.Parse(`{ "project": "example" }`))
			Expect(p.Name).To(Equal("example"))
		})

		It("Desc", func() {
			p := loadProjectJson(gjson.Parse(`{ "desc": "example" }`))
			Expect(p.Desc).To(Equal("example"))
		})

		It("Discover", func() {
			p := loadProjectJson(gjson.Parse(`{ "discover": ["1", "2"] }`))
			Expect(p.Discover).To(Equal([]string{"1", "2"}))
		})

		It("Mixin", func() {
			p := loadProjectJson(gjson.Parse(`{ "mixin": ["1", "2"] }`))
			Expect(p.Mixin).To(Equal([]string{"1", "2"}))
		})

		It("Env", func() {
			p := loadProjectJson(gjson.Parse(`{ "env": { "a": "1", "b": "2", "c": "3" } }`))
			Expect(p.Env).To(HaveLen(3))
			Expect(p.Env["a"]).To(Equal("1"))
			Expect(p.Env["b"]).To(Equal("2"))
			Expect(p.Env["c"]).To(Equal("3"))
		})

		It("EnvFiles", func() {
			p := loadProjectJson(gjson.Parse(`{ "env_files": ["1", "2"] }`))
			Expect(p.EnvFiles).To(Equal([]string{"1", "2"}))
		})

		It("Tags", func() {
			p := loadProjectJson(gjson.Parse(`{ "tags": ["1", "2"] }`))
			Expect(p.Tags).To(Equal([]string{"1", "2"}))
		})

		Describe("Tasks", func() {
			It("None", func() {
				p := loadProjectJson(gjson.Parse(`{ "tasks": {} }`))
				Expect(p.Tasks).To(BeEmpty())
			})
			It("One", func() {
				p := loadProjectJson(gjson.Parse(`{ "tasks": { "test": {} } }`))
				Expect(p.Tasks).To(HaveLen(1))
				Expect(p.Tasks["test"].Name).To(Equal("test"))
			})
			It("Two", func() {
				p := loadProjectJson(gjson.Parse(`{ "tasks": { "test1": {}, "test2": {} } }`))
				var taskNames []string
				for _, t := range p.Tasks {
					taskNames = append(taskNames, t.Name)
				}
				Expect(taskNames).To(ConsistOf("test1", "test2"))
			})
		})
	})

	Describe("ParseProject", func() {
		It("examples", func() {
			path, err := filepath.Abs("../examples/myke.yml")
			p, err := ParseProject("../examples")
			Expect(err).ToNot(HaveOccurred())
			Expect(p.Src).To(Equal(path))
			Expect(p.Cwd).To(Equal(filepath.Dir(path)))
			Expect(p.Name).To(Equal("example"))
			Expect(p.Desc).To(Equal("example project suite"))
			Expect(p.Discover).To(Equal([]string{
				"env", "tag/tags1.yml", "tag/tags2.yml",
				"hooks", "template", "mixin",
			}))
			Expect(p.Env["PATH"]).To(HavePrefix(filepath.Join(p.Cwd, "bin")))
		})

		It("examples/env", func() {
			path, err := filepath.Abs("../examples/env/myke.yml")
			p, err := ParseProject("../examples/env")
			Expect(err).ToNot(HaveOccurred())
			Expect(p.Src).To(Equal(path))
			Expect(p.Cwd).To(Equal(filepath.Dir(path)))
			Expect(p.Name).To(Equal("env"))

			Expect(p.Env["KEY_YML"]).To(Equal("value_from_yml"))
			Expect(p.Env["KEY_ENVFILE"]).To(Equal("value_from_test.env"))
			Expect(p.Env["KEY_ENVFILE_LOCAL"]).To(Equal("value_from_test.env.local"))
			Expect(p.Env["KEY_MYKE"]).To(Equal("value_from_myke.env"))
			Expect(p.Env["KEY_MYKE_LOCAL"]).To(Equal("value_from_myke.env.local"))

			expectedPaths := strings.Join([]string{
				filepath.Join(p.Cwd, "path_from_myke.env.local"),
				filepath.Join(p.Cwd, "path_from_myke.env"),
				filepath.Join(p.Cwd, "path_from_test.env.local"),
				filepath.Join(p.Cwd, "path_from_test.env"),
				filepath.Join(p.Cwd, "path_from_yml"),
				filepath.Join(p.Cwd, "bin"),
			}, string(os.PathListSeparator))
			Expect(p.Env["PATH"]).To(HavePrefix(expectedPaths))
		})

		It("examples/mixin", func() {
			path, err := filepath.Abs("../examples/mixin/myke.yml")
			p, err := ParseProject("../examples/mixin")
			Expect(err).ToNot(HaveOccurred())
			Expect(p.Src).To(Equal(path))
			Expect(p.Cwd).To(Equal(filepath.Dir(path)))
			Expect(p.Name).To(Equal("mixin"))

			Expect(p.Env["KEY_1"]).To(Equal("value_parent_1"))
			Expect(p.Env["KEY_2"]).To(Equal("value_child_2"))
			Expect(p.Env["KEY_3"]).To(Equal("value_child_3"))

			Expect(p.Tasks).To(HaveLen(3))
			Expect(p.Tasks["task1"].Cmd).To(Equal("echo parent says $KEY_1"))
			Expect(p.Tasks["task2"].Cmd).To(Equal("echo child says $KEY_2"))
			Expect(p.Tasks["task3"].Cmd).To(Equal("echo child says $KEY_3"))

			expectedPaths := strings.Join([]string{
				filepath.Join(p.Cwd, "path_child"),
				filepath.Join(p.Cwd, "bin"),
				filepath.Join(p.Cwd, "parent", "path_parent"),
				filepath.Join(p.Cwd, "parent", "bin"),
			}, string(os.PathListSeparator))
			Expect(p.Env["PATH"]).To(HavePrefix(expectedPaths))
		})
	})

})
