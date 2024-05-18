package constant

const (
	DefaultLimit  = 20
	DefaultOffset = 0
)

// 动作
const (
	ActionLike = iota
	ActionDisLike
	ActionReport
	ActionHighLight
	ActionTop
)

// type
const (
	AddComment = iota
	DeleteComment
	AuditCommentPass
	AuditCommentNoPass
	UpdateComment
)
