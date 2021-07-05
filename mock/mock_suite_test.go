package mock_test

import (
	"ginkgo-example/mock"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var router *gin.Engine
var servUrl string

func TestMock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mock Suite")
}

var _ = BeforeSuite(func() {
	router = mock.NewRouter()

	serv := httptest.NewServer(router)
	servUrl = serv.URL
})
