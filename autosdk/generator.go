package autosdk

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func findRelatedStructs(s *Struct, structs map[string]*Struct, list []*Struct) []*Struct {
	for _, prop := range s.Props {
		if prop.Level == "Object" || prop.Level == "Object Array" {
			domainName := prop.Type
			list = append(list, structs[domainName])
			findRelatedStructs(structs[domainName], structs, list)
		}
	}
	return list
}

/**
* **推荐使用**
* 根据api的获取方法生成对应即依赖的文件
* methodname api的请求名
* 生成文件指定的根路径
 */
func (m *Metadata) GenerateByApiMethod(methodName, root string) (err error) {
	if m == nil {
		return errors.New("the Metadata is nil")
	}
	if err = makeDir(root); err != nil {
		return
	}
	structMap, apiMap := m.Parse()

	list := []*Struct{}
	api := apiMap[methodName]

	api.Request.Name = api.Name
	api.Response.Name = api.Name
	if err = api.Generate(root); err != nil {
		return
	}
	for _, param := range api.Response.Params {
		if hasObject(param.Level) {
			s := structMap[param.Type]
			list = append(list, s)

			list = findRelatedStructs(s, structMap, list)
		}
	}

	for _, s := range list {
		if err = s.Generate(root); err != nil {
			break
		}
	}
	return
}
func makeDir(root string) error {
	domainPath := filepath.Join(root, "domain")
	if err := os.MkdirAll(domainPath, 0777); err != nil {
		return err
	}
	requestPath := filepath.Join(root, "request")
	if err := os.MkdirAll(requestPath, 0777); err != nil {
		return err
	}
	responsePath := filepath.Join(root, "response")
	if err := os.MkdirAll(responsePath, 0777); err != nil {
		return err
	}
	return nil
}

/**
* **不推荐使用（windows下会由于文件数过多而无法执行 go install等命令）**
* 将整个metadata.xml生成对应的api文件
*
 */
func (m *Metadata) Generate(root string) (err error) {
	if m == nil {
		return errors.New("the Metadata is nil")
	}
	if err = makeDir(root); err != nil {
		return
	}

	for _, s := range m.Structs {
		if err = s.Generate(root); err != nil {
			break
		}
	}
	for _, api := range m.Apis {
		if err = api.Generate(root); err != nil {
			break
		}
	}
	return
}
func (m *Metadata) GenerateAllInOne(root string) (err error) {
	if m == nil {
		return errors.New("the Metadata is nil")
	}
	if err := os.MkdirAll(root, 0777); err != nil {
		return err
	}

	//structs
	filename := filepath.Join(root, "api.go")
	return generate(filename, AllInOneTmpl, m)
}
func (s *Struct) Generate(root string) error {
	filename := filepath.Join(root, "domain", s.Name+".go")
	if err := generate(filename, StructTmpl, s); err != nil {
		return err
	}
	return nil
}
func (api *Api) Generate(root string) error {
	filename := filepath.Join(root, "request", GetRequestResponseStructName(api.Name, "Request")+".go")
	if err := generate(filename, RequestTmpl, api.Request); err != nil {
		return err
	}

	filename = filepath.Join(root, "response", GetRequestResponseStructName(api.Name, "Response")+".go")
	if err := generate(filename, ResponseTmpl, api.Response); err != nil {
		return err
	}
	return nil
}

func generate(filename, tmpl string, inter interface{}) error {
	funcs := make(map[string]interface{})
	funcs["GetHump"] = GetHump
	funcs["GetRequestResponseStructName"] = GetRequestResponseStructName
	funcs["GetRequestResponseStructJsonName"] = GetRequestResponseStructJsonName
	funcs["parseType"] = parseType
	funcs["hasObject"] = hasObject
	funcs["isFloat64"] = isFloat64
	funcs["xmlName"] = xmlName
	funcs["ResponeField"] = ResponeField

	t, err := template.New("t").Funcs(funcs).Parse(strings.Replace(tmpl, "'", "`", -1))
	if err != nil {
		return err
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, inter)
}
