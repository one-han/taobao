package domain

/* 该数据结构保存宝贝描述对应的规范化信息 */
type DescModuleInfo struct {
	/* 代表宝贝描述中规范化打标使用到的模块id列表，以逗号分隔。 */
	AnchorModuleIds string `json:"anchor_module_ids"`

	/* 类型代表规范化打标的类型：人工或自动化 */
	Type int64 `json:"type"`
}
