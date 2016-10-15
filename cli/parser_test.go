package cli

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	It("should parse yml file and return object", func() {
		res, err := Parse("/go/src/myke/examples/myke.yml")
		Expect(err).ToNot(HaveOccurred())
		Expect(res).To(Equal(map[interface{}]interface{}{
			"project": "example",
			"desc": "example that includes all child projects",
			"tasks": map[interface{}]interface{}{
				"build": map[interface{}]interface{}{
			   	"cmd": "echo example build",
    			"desc": "run as `s2do build` or `s2do example/build`",
    		},
    	},
    	"includes": []interface{}{
		  	"child",
		  	"env",
		  	"tagging/tags1.yml",
		  	"tagging/tags2.yml",
		  	"depends",
		  	"params",
		  	"extends",
		  },
		}))
	})
})
