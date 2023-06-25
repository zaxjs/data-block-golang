# data-block-golang
Data block golang package

### Use
``` go

import (
  dataBlock "github.com/zaxjs/data-block-golang"
)
... 
opts := dataBlock.Options{Key: "http://localhost:8089/data-block-service-api/v1/open, Api: "Y2wwemk4aWtnMDAwMjA4bDQ4c3VrZzB5bA==", Ttl: "5m", ShowSysField: true, ShowGroupInfo: true}
myBlock, _ := dataBlock.New(opts) // 建议配置为全局单例对象

res1, _ := myBlock.GetBlock([]string{ "TEST_BLOCK","TEST_MISC" }, nil)
fmt.Println("[GetBlock]:", res1)

res2, _ := myBlock.GetBlock([]string{ "TEST_BLOCK","TEST_MISC" }, &dataBlock.Options{ShowSysField: false, ShowGroupInfo: false})
fmt.Println("[GetBlock]:", res2)

res3, _ := myBlock.GetKv([]string{ "TEST_BLOCK","TEST_MISC" })
fmt.Println("[GetKv]:", res2)
  ...
  
```

### Reference
[![Go Reference](https://pkg.go.dev/badge/github.com/zaxjs/data-block-golang/tree/main.svg)](https://pkg.go.dev/github.com/zaxjs/data-block-golang/tree/main)