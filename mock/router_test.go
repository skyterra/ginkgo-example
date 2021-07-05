package mock_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Router", func() {
	Context("Echo", func() {
		It("should be succeed", func() {
			// 静态调用api函数，不需要动server

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/echo", nil)
			router.ServeHTTP(w, req)
			Expect(w.Code == http.StatusOK).Should(BeTrue())

			data, _ := ioutil.ReadAll(w.Body)
			Expect(string(data) == "通了").Should(BeTrue())
		})
	})

	Context("Echo", func() {
		It("should be succeed", func() {
			// 动态调用api函数，需要启动server

			resp, err := http.Get(fmt.Sprintf("%s/echo", servUrl))
			Expect(err).Should(Succeed())
			Expect(resp.StatusCode == http.StatusOK).Should(BeTrue())

			data, _ := ioutil.ReadAll(resp.Body)
			Expect(string(data) == "通了").Should(BeTrue())
		})
	})
})
