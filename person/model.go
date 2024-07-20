package person

// Person 结构体定义
type Person struct {
	ID        int64  `json:"id"`
	Name      string `json:"name" binding:"required"`
	Birthdate string `json:"birthdate"`
	// 其他属性 关系等 目前只支持一个关系一个关系的绑定，后续可以批量关系绑定
	Relation *Relation `json:"relation"`
}

type Relation struct {
	RelationId   int64        `json:"relationId"`
	RelationType RelationType `json:"relationType"`
}

type RelationType string

const (
	// 父母及子女
	RelationTypeFather   RelationType = "Father"
	RelationTypeMother   RelationType = "Mother"
	RelationTypeSon      RelationType = "Son"
	RelationTypeDaughter RelationType = "Daughter"

	// 祖父母及孙子女
	RelationTypeGrandfather              RelationType = "Grandfather"
	RelationTypeGrandmother              RelationType = "Grandmother"
	RelationTypeMaternalGrandfather      RelationType = "MaternalGrandfather"
	RelationTypeMaternalGrandGrandmother RelationType = "MaternalGrandGrandmother"
	RelationTypeGrandson                 RelationType = "Grandson"
	RelationTypeGranddaughter            RelationType = "Granddaughter"
	RelationTypeMaternalGrandson         RelationType = "MaternalGrandson"
	RelationTypeMaternalGranddaughter    RelationType = "MaternalGranddaughter"

	// 兄弟姐妹
	RelationTypeElderBrother   RelationType = "ElderBrother"
	RelationTypeYoungerBrother RelationType = "YoungerBrother"
	RelationTypeElderSister    RelationType = "ElderSister"
	RelationTypeYoungerSister  RelationType = "YoungerSister"

	// 配偶
	RelationTypeHusband RelationType = "Husband"
	RelationTypeWife    RelationType = "Wife"

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
