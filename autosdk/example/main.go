package main

import (
	// "github.com/ohohco/taobao/api/domain"
	"github.com/ohohco/taobao/autosdk"
	"log"
	"os"
	"path"
)

var (
	root     = path.Join(os.Getenv("GOPATH"), "src/github.com/ohohco/taobao/api")
	filename = "./ApiMetadata.xml"
)

func main() {
	var metadata *autosdk.Metadata
	metadata, err := metadata.UnmarshalFile(filename)
	if err != nil {
		log.Fatalln("error: %s", err)
	}
	if err = metadata.GeneratorByApiMethod("taobao.item.get", root); err != nil {
		log.Fatalln("error: %s", err)
	}
	if err = autosdk.Format(root); err != nil {
		log.Fatalln("error: %s", err)
	}
	root = "github.com/ohohco/taobao/api"
	if err = autosdk.Install(path.Join(root, "domain"), path.Join(root, "request"), path.Join(root, "response")); err != nil {
		log.Fatalln("error: %s", err)
	}
}
