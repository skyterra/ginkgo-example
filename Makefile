.PHONY: build check pre-install detail test cover cover-html cover-func clean

build:
	go build .

# 静态检查
check:
	go vet ./...

# 第三方依赖
pre-install:
	go get -u github.com/onsi/ginkgo/ginkgo

# 使用ginkgo命令
detail:
	ginkgo ./...

# 仅执行benchmark测试
# -v 显示命令详情
# -run="none" 仅执行以"none"开头的单元测试函数，因为没有以"none"开头的单侧函数，所以这里是不想执行单侧，仅执行benchmark
benchmark:
	go test -v -run="none" -bench="Benchmark" ./...

# 执行单元测试
test:
	go test ./...

# 统计覆盖率
cover:
	go test ./... -coverprofile cover.profile

# 打开浏览器显示覆盖统计信息
cover-html:cover
	go tool cover -html=cover.profile

# 显示函数覆盖统计
cover-func:cover
	go tool cover -func=cover.profile

# 清理
clean:
	rm cover.profile
