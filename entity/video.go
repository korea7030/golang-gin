package entity

/*
field에 binding:"조건" 또는 validate:"함수이름" 로 구현
테스트 해봤을 때 binding을 먼저 수행 후 validate를 수행하는 것으로 보임
binding min,max 시 붙여서 써야 함. 띄워 쓰면 다른 binding 변수로 착각
*/
type Person struct {
	FirstName string `json:"firstname" binding: "required"`
	LastName  string `json:"lastname" binding: "required"`
	Age       int8   `json:"age" binding: "gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}
type Video struct {
	Title       string `json:"title" validate:"is-cool" binding:"min=2,max=10"`
	Description string `json:"description" binding:"max=20"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding: "required"`
}
