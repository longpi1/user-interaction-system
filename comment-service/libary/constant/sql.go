package constant

const (
	WhereByCommentID     = "comment_id = ?"
	WhereByName          = "name = ?"
	WhereByUserName      = "username = ?"
	WhereByModelName     = "model_name = ?"
	WhereByType          = "type = ?"
	WhereByContent       = "content = ?"
	WhereByResourceID    = "resource_id = ?"
	WhereByResourceTitle = "resource_title = ?"
	WhereByPID           = "pid = ?"
)

const (
	OrderDescById         = "id desc"
	OrderDescByFloorCount = "floor_count desc"
)

const (
	FuzzySearch = "%"
)

const (
	GreaterThanCreatedAt = "created_at > ?"
	LessThanCreatedAt    = "created_at < ?"
)
