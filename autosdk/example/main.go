package main

import (
	"github.com/one-han/taobao/autosdk"
	"log"
	"os"
	"path"
)

var (
	//生成文件所在路径.
	root = "github.com/one-han/taobao/api"
	//ApiMetadata.xml文件
	filename = "/Users/han/go/src/github.com/one-han/taobao/autosdk/example/ApiMetadata.xml"
	//api文档的调用方法
	apiMethod = "taobao.tbk.tpwd.create"
	//生成文件所在路径(绝对路径).
	dest = path.Join(os.Getenv("GOPATH"), "src/github.com/one-han/taobao/api")
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	//文件读取
	var metadata *autosdk.Metadata
	metadata, err := metadata.UnmarshalFile(filename)
	if err != nil {
		panic(err)
	}
	//文件生成
	if err = metadata.GenerateByApiMethod(apiMethod, dest); err != nil {
		panic(err)
	}
	//if err = metadata.(dest); err != nil {
	//	panic(err)
	//}
	//代码格式化：
	if err = autosdk.Format(dest); err != nil {
		panic(err)
	}
	//代码install:
	if err = autosdk.Install(
		dest,
	// 	// path.Join(root, "domain"),
	// 	// path.Join(root, "request"),
	// 	// path.Join(root, "response"),
	); err != nil {
		panic(err)
	}
}
