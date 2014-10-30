package main

import (
	"github.com/ohohco/taobao/autosdk"
	"log"
	"os"
	"path"
)

var (
	root     = "github.com/ohohco/taobao/api"
	filename = "./ApiMetadata.xml"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	var metadata *autosdk.Metadata
	metadata, err := metadata.UnmarshalFile(filename)
	if err != nil {
		log.Fatalln("error: %s", err)
	}

	//1. 根据api方法名生成sdk
	//参考地址：http://open.taobao.com/api/api_cat_detail.htm?spm=a219a.7386789.0.0.IG5dsm&cat_id=4&category_id=102
	rootPath := path.Join(os.Getenv("GOPATH"), "src/"+root)
	log.Println("rootPath:", rootPath)
	if err = metadata.GeneratorByApiMethod("taobao.item.get", rootPath); err != nil {
		log.Fatalf("error: %s", err)
	}

	//2. 格式化go文件
	if err = autosdk.Format(rootPath); err != nil {
		log.Fatalf("error: %s", err)
	}

	//3. install go文件
	if err = autosdk.Install(path.Join(root, "domain"), path.Join(root, "request"), path.Join(root, "response")); err != nil {
		log.Fatalf("error: %s", err)
	}
}
