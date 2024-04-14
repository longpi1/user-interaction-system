package constant

const (
	WhereByID         = "id = ?"
	WhereByPlatform   = "Platform = ?"
	WhereByName       = "name = ?"
	WhereByUserID     = "uid = ?"
	WhereByType       = "type = ?"
	WhereByResourceID = "resource_id = ?"
	WhereByPID        = "pid = ?"
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

const (
	RelationTableName      = "relation"
	RelationCountTableName = "relation_count"
)
