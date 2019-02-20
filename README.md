# taobao
[淘宝开放平台](http://open.taobao.com/api/api_list.htm)的go版的SDK. 内置sdk生成工具.

## autosdk工具包
可由淘宝提供的ApiMetadata.xml,自动生成sdk.
主要有两种生成方式:
* 分结构生成
* 一键生成所有

### 分结构生成
此方式结构清晰, 适用于按需生成, 且与淘宝官方提供的JAVA版本SDK结构一致. 
生成结构如下:
> filepackage/domain  
  filepackage/request  
  filepackage/response  

#### 使用示例
比如`得到单个商品信息`,[参见文档](http://open.taobao.com/api/api_cat_detail.htm?spm=a219a.7386789.0.0.8MDkjq&cat_id=4&category_id=102),其调用方法`taobao.item.get`.代码如下：
~~~
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
	filename = "./ApiMetadata.xml"
	//api文档的调用方法
	apiMethod = "taobao.item.get"
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
	///////////////
	////方式1//////
	///////////////
	if err = metadata.GenerateByApiMethod(apiMethod, dest); err != nil {
		panic(err)
	}
	///////////////
	////其他方式///
	///////////////

	//代码格式化：
	if err = autosdk.Format(dest); err != nil {
		panic(err)
	}
	//代码install:
	if err = autosdk.Install(path.Join(root, "domain"),
		path.Join(root, "request"),
		path.Join(root, "response")); err != nil {
		panic(err)
	}
}
~~~
```
///////////////
////方式2//////
///////////////
if err = metadata.Generate(dest); err != nil {
	panic(err)
}
```
> 注意：
	* 示例中`metadata.Generate(dest)`可一键生成所有. 由于是分文件存储, 其数量过多会致使`go install`等命令不通过.

### 一键生成所有
该方式会将所有结构和方法生成在一个文件, API调用更简单, 当然缺点是文件更大, 查找不方便.
结构为filepackage
#### 示例代码: 
```
///////////////
////方式3//////
///////////////
if err = metadata.GenerateAllInOne(dest); err != nil {
	panic(err)
}
```

## api工具包
* appKey,appSecret可以创建应用获取, [参考文档](http://open.taobao.com/doc/detail.htm?id=101618).
* sessionKey可以由工具生成,[工具地址](http://api.taobao.com/apitools/sessionPage.htm?spm=a219a.7395905.0.10).
* [淘宝测试工具](http://api.taobao.com/apitools/apiTools.htm?spm=a219a.7386781.0.0.0WxMxQ).

#### 使用示例
比如`得到单个商品信息`,可自行选择调用方式`http`或`https`:
~~~
package main

import (
	"github.com/one-han/taobao/api"
	"github.com/one-han/taobao/api/request"
	"github.com/one-han/taobao/api/response"
	"log"
)

var (
	serverUrl   = "" //为空则默认使用正式环境
	appKey      = "your appKey"
	appSecret   = "your appSecret"
	sessionKey  = "your sessionKey"
	accessToken = "your accessToken"
)

func main() {
	client := api.GetDefaultTaobaoClient(appKey, appSecret, serverUrl)
	//免签名方式(https)
	// client := api.GetHttpsTaobaoClient(serverUrl)

	var resp response.ItemGetResponse
	var req = new(request.ItemGetRequest)
	req.SetFields("desc")
	req.SetNumIid("21301327956")

	//免签名方式(https)时使用accessToken
	data, err := client.Excute(req, &resp, sessionKey) //"sessionKey or accessToken"
	log.Println("data:", string(data))                 //data为淘宝平台返回的字符串
	if err != nil {
		panic(err)
	}
	if resp.ErrorResponse != nil {
		panic(resp.Msg)
	}
	log.Println("desc:", resp.Item.Desc)
}
~~~
如果是使用上述`方式3`生成的SDK, API调用时仅引入`github.com/one-han/taobao/api`即可.

## 计划
* 目前是基于api调用方法的单个生成, 对于整个ApiMetadata.xml, 由于go install对文件数量的限制, 一键生成需要将文件合并. 已完成
* https免签调用. 已完成
* web授权支持.
* 完善多种加密方式md5/hmac.

## 说明
思路源自[changkong](https://github.com/changkong/open_taobao), 特此鸣谢.  
较之添加了对appKey,appSecret传入的支持. 
适用于多个店铺共用同样的appKey,appSecret不同的sessionKey.  

## 关于
本人未接受过专业训练, 业余爱好, 纯菜鸟一枚. 诚意接受各位的指正批评, 不接受人身攻击. 乐意各位的指点, 可请之喝酒玩乐!
