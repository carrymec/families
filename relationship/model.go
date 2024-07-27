package relationship

type Relationship struct {
	TypeName string            `json:"typeName" binding:"required"`
	Desc     string            `json:"desc"`
	Tags     map[string]string `json:"tags"` // 关系支持打标签
}
