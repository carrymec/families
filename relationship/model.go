package relationship

type Relationship struct {
	Id       int64             `json:"id"`
	TypeName string            `json:"typeName" binding:"required"`
	Desc     string            `json:"desc"`
	Tags     map[string]string `json:"tags"` // 关系支持打标签
}

type Query struct {
	TypeName string `json:"typeName" binding:"required"`
	Tags     string `json:"tags"`
}
