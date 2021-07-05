# ginkgo-example
ginkgo 使用例子

## 常规UT测试
./calc目录下位常规测试，包括benchmark类测试  

## 一个简单个mock server类测试
./mock目录下使用Gin作为web-frame进行开发，ginkgo使用mock server对api进行UT测试

## 执行测试
make test

## 显示测试详情
make detail

## 执行基准测试
make benchmark

## 覆盖率统计
make cover-html
