package calc_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var utlog = GinkgoWriter

func TestCalc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calc Suite")
}

// 仅当测试失败时，使用该函数打印的日志才会显示再控制台
// 如果需要再测试成功时，也显示打印日志，请使用ginkgo -v命令
func UTLogf(format string, v ...interface{}) {
	utlog.Write([]byte(fmt.Sprintf(format, v...)))
}

func UTLog(v ...interface{}) {
	utlog.Write([]byte(fmt.Sprintln(v...)))
}
