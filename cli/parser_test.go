package cli

import (
  "os"
  "io/ioutil"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/MakeNowJust/heredoc/dot"
)

var _ = Describe("Parser", func() {

  Describe("ParseYaml", func() {
    It("parses yaml as object", func() {
      yml := []byte(D(`
        str: testString
        arr:
          - one
          - two
        map:
          key1:
            nest1: value1
          key2:
            nest2: value2
        longtext: |-
          foo
            bar
          baz
      `))

      file, err := ioutil.TempFile(os.TempDir(), "myke")
      defer os.Remove(file.Name())
      Expect(err).NotTo(HaveOccurred())
      
      err = ioutil.WriteFile(file.Name(), yml, 0644)
      Expect(err).NotTo(HaveOccurred())

      v, err := ParseYaml(file.Name())
      Expect(err).NotTo(HaveOccurred())
      Expect(v.Get("str").String()).To(Equal("testString"))
      Expect(v.Get("arr").Array()[0].String()).To(Equal("one"))
      Expect(v.Get("arr").Array()[1].String()).To(Equal("two"))
      Expect(v.Get("map").Get("key1").Get("nest1").String()).To(Equal("value1"))
      Expect(v.Get("map").Get("key2").Get("nest2").String()).To(Equal("value2"))
      Expect(v.Get("longtext").String()).To(Equal("foo\n  bar\nbaz"))
    })
  })

})
