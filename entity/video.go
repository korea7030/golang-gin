package entity

import "time"

/*
field에 binding:"조건" 또는 validate:"함수이름" 로 구현
테스트 해봤을 때 binding을 먼저 수행 후 validate를 수행하는 것으로 보임
binding min,max 시 붙여서 써야 함. 띄워 쓰면 다른 binding 변수로 착각
gorm: ORM 사용할 때 쓰임(in SQLite).
*/
type Person struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"firstname" binding: "required" gorm:"type:varchar(32)"`
	LastName  string `json:"lastname" binding: "required" gorm:"type:varchar(32)"`
	Age       int8   `json:"age" binding: "gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email" gorm:"type:varchar(256)"`
}
type Video struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `json:"title" validate:"is-cool" binding:"min=2,max=100" gorm:"type:varchar(100)"`
	Description string    `json:"description" binding:"max=200" gorm:"type:varchar(200)"`
	URL         string    `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author      Person    `json:"author" binding: "required" gorm:"foreignkey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
