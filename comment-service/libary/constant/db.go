package constant

const (
	WhereByID            = "id = ?"
	WhereByCommentID     = "comment_id = ?"
	WhereByName          = "name = ?"
	WhereByUserName      = "username = ?"
	WhereByUserID        = "user_id = ?"
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

const (
	ResourceTableName       = "resource"
	CommentIndexTableName   = "comment_index"
	CommentContentTableName = "comment_content"
	UserTableName           = "user"
	UserCommentTableName    = "user_comment"
)
