package constant

const (
	WhereByID        = "id = ?"
	WhereByName      = "name = ?"
	WhereByUserName  = "username = ?"
	WhereByModelName = "model_name = ?"
	WhereByType      = "type = ?"
	WhereByContent   = "content = ?"
)

const (
	OrderDescById = "id desc"
)

const (
	FuzzySearch = "%"
)

const (
	GreaterThanCreatedAt = "created_at > ?"
	LessThanCreatedAt    = "created_at < ?"
)
