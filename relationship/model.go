package relationship

import "github/carrymec/families/common"

type Relationship struct {
	TypeName string            `json:"typeName" binding:"required"`
	Desc     string            `json:"desc"`
	Tags     map[string]string `json:"tags"` // 关系支持打标签
}

const (
	// 父母及子女
	RelationTypeFather   common.RelationType = "Father"
	RelationTypeMother   common.RelationType = "Mother"
	RelationTypeSon      common.RelationType = "Son"
	RelationTypeDaughter common.RelationType = "Daughter"

	// 祖父母及孙子女
	RelationTypeGrandfather              common.RelationType = "Grandfather"
	RelationTypeGrandmother              common.RelationType = "Grandmother"
	RelationTypeMaternalGrandfather      common.RelationType = "MaternalGrandfather"
	RelationTypeMaternalGrandGrandmother common.RelationType = "MaternalGrandGrandmother"
	RelationTypeGrandson                 common.RelationType = "Grandson"
	RelationTypeGranddaughter            common.RelationType = "Granddaughter"
	RelationTypeMaternalGrandson         common.RelationType = "MaternalGrandson"
	RelationTypeMaternalGranddaughter    common.RelationType = "MaternalGranddaughter"

	// 兄弟姐妹
	RelationTypeElderBrother   common.RelationType = "ElderBrother"
	RelationTypeYoungerBrother common.RelationType = "YoungerBrother"
	RelationTypeElderSister    common.RelationType = "ElderSister"
	RelationTypeYoungerSister  common.RelationType = "YoungerSister"

	// 配偶
	RelationTypeHusband common.RelationType = "Husband"
	RelationTypeWife    common.RelationType = "Wife"

	/*
		TODO
		兄弟姐妹的配偶

		嫂子 (Sister-in-law, wife of elder brother)
		弟媳 (Sister-in-law, wife of younger brother)
		姐夫 (Brother-in-law, husband of elder sister)
		妹夫 (Brother-in-law, husband of younger sister)
		叔伯及姑姨
		父系亲属

		叔叔 (Uncle, father's younger brother)
		伯伯 (Uncle, father's elder brother)
		姑姑 (Aunt, father's sister)
		婶婶 (Aunt, wife of father's younger brother)
		伯母 (Aunt, wife of father's elder brother)
		母系亲属

		舅舅 (Uncle, mother's brother)
		姨妈 (Aunt, mother's sister)
		舅妈 (Aunt, wife of mother's brother)
		姨夫 (Uncle, husband of mother's sister)
		其他亲属
		堂表兄弟姐妹
		堂兄 (Cousin, elder male cousin from father's side)
		堂弟 (Cousin, younger male cousin from father's side)
		堂姐 (Cousin, elder female cousin from father's side)
		堂妹 (Cousin, younger female cousin from father's side)
		表兄 (Cousin, elder male cousin from mother's side)
		表弟 (Cousin, younger male cousin from mother's side)
		表姐 (Cousin, elder female cousin from mother's side)
		表妹 (Cousin, younger female cousin from mother's side)
		继亲关系
		继父母及继子女

		继父 (Stepfather)
		继母 (Stepmother)
		继子 (Stepson)
		继女 (Stepdaughter)
		半兄弟姐妹

		同父异母兄弟 (Half-brother, same father)
		同父异母姐妹 (Half-sister, same father)
		同母异父兄弟 (Half-brother, same mother)
		同母异父姐妹 (Half-sister, same mother)
	*/
)
