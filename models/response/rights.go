package response

type Menus struct {
	ID       uint    `json:"id"`
	AuthName string  `json:"authName"` //权限说明
	Path     string  `json:"path"`     //对应访问路径
	Children []Menus `json:"children"`
}
