package person

// Person 结构体定义
type Person struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	// 其他属性
}
