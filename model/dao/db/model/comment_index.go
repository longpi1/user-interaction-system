package model

import (
	"gorm.io/gorm"
	"user-interaction-system/libary/constant"
	"user-interaction-system/model/dao/db"
)

/*
comment_index：索引表
记录评论的索引
同样记录对应的主题，方便后续查询
通过 root 和 parent 记录是否是根评论以及子评论的上级
floor 记录评论层级，也需要更新主题表中的楼层数，
*/
type CommentIndex struct {
	gorm.Model
	CommentId       uint `gorm:"index"` // 评论id
	ResourceId      uint `gorm:"index"` // 评论所关联的资源id
	UserID          uint `gorm:"index"` //  发表者id
	UserName        string
	IP              string
	IPArea          string
	Pid             uint // 父评论 ID
	Type            uint // 评论类型，文字评论、图评等
	IsCollapsed     bool `gorm:"default:false"` // 折叠
	IsPending       bool `gorm:"default:false"` // 待审
	IsPinned        bool `gorm:"default:false"` // 置顶
	State           uint // 状态
	LikeCount       uint // 点赞数
	HateCount       uint // 点踩数
	ReplyCount      uint // 回复数
	RootReplayCount uint // 根评论回复数
	FloorCount      uint // 评论层数
}

func InsertCommentIndex(commentIndex *CommentIndex) error {
	err := db.GetClient().Create(&commentIndex).Error
	return err
}

func InsertBatchCommentIndex(commentIndexs []*CommentIndex) error {
	err := db.GetClient().Create(&commentIndexs).Error
	return err
}

func DeletCommentIndex(commentIndex *CommentIndex) error {
	err := db.GetClient().Unscoped().Delete(&commentIndex).Error
	return err
}

func FindCommentIndexById(id string) (CommentIndex, error) {
	var commentIndex CommentIndex
	err := db.GetClient().Where(constant.WhereByID, id).First(&commentIndex).Error
	return commentIndex, err
}

func UpdatCommentIndex(commentIndex *CommentIndex) error {
	err := db.GetClient().Updates(&commentIndex).Error
	return err
}

func GetCommentIndexList(param CommentParamsList) (commentIndexs []CommentIndex, err error) {
	tx := db.GetClient()
	if param.Type != 0 {
		tx = tx.Where(constant.WhereByType, param.Type)
	}
	if param.Username != "" {
		tx = tx.Where(constant.WhereByUserName, param.Username)
	}
	if param.Content != "" {
		tx = tx.Where(constant.WhereByContent, constant.FuzzySearch+param.Content+constant.FuzzySearch)
	}
	if param.Limit == 0 {
		param.Limit = constant.DefaultLimit
	}
	err = tx.Order(constant.OrderDescById).Limit(param.Limit).Offset(param.Offset).Find(&commentIndexs).Error
	return commentIndexs, err
}

func DeleteCommentIndexByTime(deleteTime int) error {
	err := db.GetClient().Where(constant.LessThanCreatedAt, deleteTime).Delete(&CommentIndex{}).Error
	return err
}
