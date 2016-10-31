package cli

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"path/filepath"
	"strings"
)

var _ = Describe("Parser", func() {

	Describe("ParseProject", func() {
		It("examples", func() {
			path, err := filepath.Abs("../examples/myke.yml")
			p, err := ParseProject("../examples")
			Expect(err).ToNot(HaveOccurred())
			Expect(p.Src).To(Equal(path))
			Expect(p.Cwd).To(Equal(filepath.Dir(path)))
			Expect(p.Name).To(Equal("example"))
			Expect(p.Desc).To(Equal("example project suite"))
			Expect(p.Includes).To(Equal([]string{
				"child", "env", "tagging/tags1.yml", "tagging/tags2.yml",
				"depends", "params", "extends",
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
			Expect(p.Desc).To(Equal("usage of env vars"))

			expectedPaths := strings.Join([]string{
				filepath.Join(p.Cwd, "env_local_file_path"),
				filepath.Join(p.Cwd, "env_file_path"),
				filepath.Join(p.Cwd, "env_yml_path"),
				filepath.Join(p.Cwd, "bin"),
			}, string(os.PathListSeparator))
			Expect(p.Env["PATH"]).To(HavePrefix(expectedPaths))
		})

		It("examples/extends", func() {
			path, err := filepath.Abs("../examples/extends/myke.yml")
			p, err := ParseProject("../examples/extends")
			Expect(err).ToNot(HaveOccurred())
			Expect(p.Src).To(Equal(path))
			Expect(p.Cwd).To(Equal(filepath.Dir(path)))
			Expect(p.Name).To(Equal("extends"))
			Expect(p.Desc).To(Equal("demonstrates how one yml file can extend another"))
			Expect(p.Env["KEY_1"]).To(Equal("value_parent_1"))
			Expect(p.Env["KEY_2"]).To(Equal("value_child_2"))
			Expect(p.Env["KEY_3"]).To(Equal("value_child_3"))

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
