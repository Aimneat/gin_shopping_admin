package response

type Categories struct {
	ID                uint         `json:"cat_id"`
	CategoriesName    string       `json:"cat_name"`
	CategoriesPid     uint         `json:"cat_pid"`
	CatgoriesLevel    uint         `json:"cat_level"`
	CategoriesDeleted bool         `json:"cat_deleted"`
	Children          []Categories `json:"children"`
}
