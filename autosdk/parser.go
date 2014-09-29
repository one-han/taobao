package autosdk

import (
	"encoding/xml"
	"io/ioutil"
)

func (m *Metadata) Unmarshal(data []byte) (*Metadata, error) {
	return m, xml.Unmarshal(data, &m)
}

/**
* 将配置文件解析为Metadata结构体
 */
func (m *Metadata) UnmarshalFile(filename string) (*Metadata, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return m.Unmarshal(data)
}

func (m *Metadata) Parse() (map[string]*Struct, map[string]*Api) {
	if m == nil {
		return nil, nil
	}
	structMap := make(map[string]*Struct)
	for _, s := range m.Structs {
		structMap[s.Name] = s
	}
	apiMap := make(map[string]*Api)
	for _, api := range m.Apis {
		apiMap[api.Name] = api
	}
	return structMap, apiMap
}
