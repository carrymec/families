package person

import "github/carrymec/families/common"

// Person 结构体定义
type Person struct {
	ID        int64  `json:"id"`
	Name      string `json:"name" binding:"required"`
	Birthdate string `json:"birthdate"`
	Note      string `json:"note"`
	Avatar    Avatar `json:"avatar"`
	// 其他属性 关系等 目前只支持一个关系一个关系的绑定，后续可以批量关系绑定，还可以绑定出生地居住地等信息
	Relation *Relation `json:"relation"`
}

// 头像
type Avatar struct {
	Key      string `json:"key"`
	FileName string `json:"FileName"`
	Url      string `json:"url"`
}

// 关系
type Relation struct {
	RelationId   int64               `json:"relationId"`
	RelationType common.RelationType `json:"relationType"`
	// TODO 支持关系属性，如夫妻关系，可以给这个关系加上结婚时间等
	//Labels map[string]string `json:"labels"`
}

type Query struct {
	Name         string              `json:"name"`
	RelationType common.RelationType `json:"relationType"`
	Page         int64               `json:"page"`
	PageSize     int64               `json:"pageSize"`
	//TODO 其他条件
}
