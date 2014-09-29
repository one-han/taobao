package autosdk

type Metadata struct {
	VersionNo string    `xml:"versionNo,attr"`
	Structs   []*Struct `xml:"structs>struct"`
	Apis      []*Api    `xml:"apis>api"`
}

/**
* 抽象出公有的属性
 */
type BaseObject struct {
	Name string `xml:"name"`
	Desc string `xml:"desc"`
	// CustomerName string
	// LangType     string
}

/**
* 属性
 */
type Prop struct {
	BaseObject
	Type  string `xml:"type"`
	Level string `xml:"level"`
}

/*
 * struct
 */
type Struct struct {
	BaseObject
	Props []*Prop `xml:"props>prop"`
}

/* struct end*/
type Api struct {
	BaseObject
	Request
	Response
}

type Request struct {
	BaseObject
	Params []*RequestParam `xml:"request>param"`
}
type Response struct {
	BaseObject
	// HasObject bool
	Params []*ResponseParam `xml:"response>param"`
}

type RequestParam struct {
	BaseObject
	Type        string `xml:"type"`     // Field List/String/Number/Boolean/Price/byte[]/Date
	Required    string `xml:"required"` // required/optional/special 其中 special 多半参照 desc 的说明
	MaxValue    string `xml:"max_value"`
	MaxLength   string `xml:"max_length"`
	MaxListSize string `xml:"max_list_size"`
	MinValue    string `xml:"min_value"`
}

type ResponseParam struct {
	BaseObject
	Type  string `xml:"type"`
	Level string `xml:"level"`
}
