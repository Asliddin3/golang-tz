package models

type Person struct {
	ID         int    `gorm:"type:bigint;primaryKey" json:"id"`
	Name       string `gorm:"type:varchar(250);default:null" json:"name"`
	Surname    string `gorm:"type:varchar(250);default:null" json:"surname"`
	Patronymic string `gorm:"type:varchar(250);default:null" json:"patronymic"`
	Age        int    `gorm:"type:smallint;default:0" json:"age"`
	Gender     string `gorm:"type:varchar(10);default:null" json:"gender"`
	Nation     string `gorm:"type:varchar(3);default:null" json:"nation"`
}

type PersonRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}
type PersonUpdateRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Nation     string `json:"nation"`
}
type PersonFilter struct {
	MultiSearch string `json:"multiSearch" form:"multiSearch"`
	Name        string `json:"name" form:"name"`
	Age         int    `json:"age" form:"age"`
	Gender      string `json:"gender" form:"gender"`
	Nation      string `json:"nation" form:"nation"`
	Page        int    `json:"page" form:"page"`
	PageSize    int    `json:"pageSize" form:"pageSize"`
}
type PersonFilterResponse struct {
	Person   []Person `json:"person" form:"person"`
	Page     int      `json:"page" form:"page"`
	PageSize int      `json:"pageSize" form:"pageSize"`
	Count    int      `json:"count" form:"count"`
}
