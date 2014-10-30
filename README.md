# taobao
[淘宝开放平台](http://open.taobao.com/api/api_list.htm)的go版的SDK. 内置sdk生成工具.

## autosdk
可由淘宝提供的ApiMetadata.xml,自动生成sdk.

### 使用示例
比如`得到单个商品信息`,[参见文档](http://open.taobao.com/api/api_cat_detail.htm?spm=a219a.7386789.0.0.8MDkjq&cat_id=4&category_id=102),其调用方法`taobao.item.get`.代码如下：
~~~
//filename : ApiMetadata.xml文件
var metadata *autosdk.Metadata
metadata, err := metadata.UnmarshalFile(filename)
if err != nil {
	panic(err)
}
//taobao.item.get : api文档的调用方法
//yourcreatedpath : 生成文件所在路径.(绝对路径：path.Join(os.Getenv("GOPATH"), "src/github.com/ohohco/taobao/api"))
if err = metadata.GeneratorByApiMethod("taobao.item.get", yourcreatedpath); err != nil {
	panic(err)
}
~~~
代码格式化：
~~~
if err = autosdk.Format(yourcreatedpath); err != nil {
	panic(err)
}
~~~
代码install:
~~~
//root : 生成文件所在路径.(github.com/ohohco/taobao/api)
if err = autosdk.Install(path.Join(root, "domain"), path.Join(root, "request"), path.Join(root, "response")); err != nil {
	panic(err)
}
~~~

###api###
* appKey,appSecret可以创建应用获取, [参考文档](http://open.taobao.com/doc/detail.htm?id=101618).
* sessionKey可以由工具生成,[工具地址](http://api.taobao.com/apitools/sessionPage.htm?spm=a219a.7395905.0.10).
* [淘宝测试工具](http://api.taobao.com/apitools/apiTools.htm?spm=a219a.7386781.0.0.0WxMxQ).

使用示例
比如`得到单个商品信息`:
~~~
client := api.DefaultTaobaoClient{
	"http://gw.api.taobao.com/router/rest",
	"your appKey",
	"your appSecret",
}
var resp response.ItemGetResponse
var req = new(request.ItemGetRequest)
req.SetFields("detail_url,num_iid,title,nick,typel")
req.SetNumIid("41195705534")
data, err := client.Excute(req, &resp, "your sessionKey")
if err != nil {
	panic(err)
}
log.Println("data:", string(data))//data为淘宝平台返回的字符串

if resp.ErrorResponse != nil {
	panic(resp.Msg)
}
log.Println("item:", resp.Item)
~~~

## 计划
* 目前是基于api调用方法的单个生成, 对于整个ApiMetadata.xml, 由于go install对文件数量的限制, 一键生成需要将文件合并.未完成
* https免签调用.未完成
* web支持.未完成

## 说明
本人未接受过专业训练, 纯菜鸟一枚.诚意接受各位的指正批评,不接受人身攻击.乐意各位的指点,可请之喝酒玩乐!

## 关于
思路源自[changkong](https://github.com/changkong/open_taobao),特别鸣谢.
