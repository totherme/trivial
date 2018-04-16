package hello_test

import (
	"io/ioutil"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/totherme/trivial/hello"
)

var _ = Describe("Handler", func() {
	It("always says hello", func() {
		response := httptest.NewRecorder()
		hello.Handler(response, nil)
		body, err := ioutil.ReadAll(response.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(body).To(ContainSubstring("Hello the cloud!"))
	})
})
