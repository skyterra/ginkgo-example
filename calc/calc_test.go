package calc_test

import (
	"ginkgo-example/calc"
	"math"
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/*
 * Context内的It block是顺序执行的
 * Context间是随机的
 *
 * NOTE：如果想要It也是随机执行的话,使用命令 ginkgo --randomizeAllSpecs
 */

var _ = Describe("Calc", func() {
	Context("Add Positive", func() {
		It("should be succeed", func() {
			UTLog("test 10 + 20")
			Expect(calc.Add(10, 20) == 30).Should(BeTrue())
		})

		It("should be succeed", func() {
			UTLog("test MaxInt32 + 1")
			Expect(calc.Add(math.MaxInt32, 1) == math.MinInt32).Should(BeTrue())
		})
	})

	Context("Add Negative", func() {
		It("should be succeed", func() {
			UTLog("test -20 + (-10)")
			Expect(calc.Add(-20, -10) == -30).Should(BeTrue())
		})

		It("should be succeed", func() {
			UTLog("test MinInt32 - 1")
			Expect(calc.Add(math.MinInt32, -1) == math.MaxInt32).Should(BeTrue())
		})
	})

	Context("Devide 0", func() {
		It("should be failed", func() {
			UTLog("test 10 / 0")
			_, err := calc.Divide(10, 0)
			Expect(err).ShouldNot(Succeed())
		})
	})

	// ginkgo版的benchmark测试函数
	Measure("Add Benchmark", func(b Benchmarker) {
		runtime := b.Time("runtime", func() {
			a := rand.Int31()
			b := rand.Int31()
			calc.Add(a, b)
		})

		Ω(runtime.Seconds()).Should(BeNumerically("<", 0.2), "Add() shouldn't take too long.")
	}, 10)

})

// golang自带的benchmark测试函数，benchmark测试函数必须以"Benchmark"开头，这个golang要求的
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc.Add(rand.Int31(), rand.Int31())
	}
}
