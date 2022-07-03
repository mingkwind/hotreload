# Golang 配置文件热加载

使用 `github.com/fsnotify/fsnotify`，监听 `Create`、`Remove`、`Write`、`Chmod`、`Rename` 等事件，实现对文件状态的实时监听，当文件有变化时执行已注册的回调函数（如下），实现对配置的重新加载。

```go
// CallbackFunc 配置回调函数
// filename：文件相对路径
type CallbackFunc func(filename string) error
```

以读取配置为例，来说明如何实现热加载：

## 定义配置文件对应结构体

配置文件 `config.json` 内容：  
```html
{
    "model":"aaaaa"
}
```

为配置文件定义结构：

```go
type Config struct {
	Model string `json:"model"`
}
```

## 定义读取配置文件的方法

```go
var cnf Config

func loadConfig(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	err = jsoniter.Unmarshal(file, &cnf)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("config:%+v\n", cnf)

	return nil
}
```

## 注册回调

```go
// 1. 注册配置文件监听及回调
hotreload.Register("conf/config.json", loadConfig)
```


## 启动监听

```go
// 2. 启动监听
hotreload.Watcher()
```

## 完整 demo

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/LeungGeorge/grit/lib/hotreload"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
)

type Config struct {
	Model string `json:"model"`
}

var cnf Config

func loadConfig(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	err = jsoniter.Unmarshal(file, &cnf)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("config:%+v\n", cnf)

	return nil
}

func main() {
	r := gin.Default()

    // 1. 注册配置文件监听及回调
	hotreload.Register("conf/config.json", loadConfig)
    // 2. 启动监听
	hotreload.Watcher()

	r.Run()
}
```

输出（可以看到 `demo` 读取到 `conf/config.json` 的内容）：  
```html
config:{Model:aaaaa}
```
