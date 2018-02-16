package latte_test

import (
	"bytes"
	"html/template"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//. "github.com/golatte/latte"
)

var _ = Describe("View", func() {
	Describe("Test layouts", func() {
		It("Should find the layout", func() {
			var b bytes.Buffer
			t, err := template.New("layout.html").Parse("<layout>yield</layout>")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(t).ShouldNot(BeNil())
			err = t.Execute(&b, nil)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b.String()).Should(Equal("<layout>yield</layout>"))
		})
	})

})
