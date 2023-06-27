# data-block-golang
Data block golang package

### Use
``` go
  import (
    "fmt"
    dataBlock "github.com/zaxjs/data-block-golang"
  )

  opts := dataBlock.Options{Api: "API", Key: "KEY", Ttl: "5m", ShowSysField: true, ShowGroupInfo: true, ShowRawData: false}
  myBlock, _ := dataBlock.New(opts) // 建议配置为全局单例对象

  res1, _ := myBlock.Block([]string{ "TEST_BLOCK","TEST_MISC" }, nil)
  fmt.Println("[Block]:", res1)

  res2, _ := myBlock.Block([]string{ "TEST_BLOCK","TEST_MISC" }, &dataBlock.Options{ShowSysField: false, ShowGroupInfo: false, ShowRawData: false})
  fmt.Println("[Block]:", res2)

  res3, _ := myBlock.Kv([]string{ "TEST_BLOCK","TEST_MISC" })
  fmt.Println("[Kv]:", res2)
```

### Test
``` go
go test .
```

### Docs
[![Go Reference](https://pkg.go.dev/badge/github.com/zaxjs/data-block-golang/tree/main.svg)](https://pkg.go.dev/github.com/zaxjs/data-block-golang/tree/main)