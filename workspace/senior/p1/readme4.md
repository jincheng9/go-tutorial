# Go testing缓存导致测试没执行的问题

## Methods

- `go clean -testcache`: expires all test results

- use non-cacheable flags on your test run. The idiomatic way is to use `-count=1`

  

## References

* https://stackoverflow.com/questions/48882691/force-retesting-or-disable-test-caching