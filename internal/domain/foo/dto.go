package foo

type FooRegister struct {
	Name string `json:"name" binding:"required"`
}

type FooUpdate struct {
	Name string `json:"key" binding:"required"`
}

type FooView struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
