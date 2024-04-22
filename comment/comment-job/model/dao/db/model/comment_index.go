package model

import (
	"comment-job/libary/constant"
	"comment-job/model/dao/db"

	"gorm.io/gorm"
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
	ResourceID      int64  `gorm:"index;comment:'评论所关联的资源id'"` // 评论所关联的资源id
	ResourceType    string `gorm:"comment:'评论所关联的资源类型'"`       // 评论所关联的资源类型
	UserID          int64  `gorm:"index;comment:'发表者id'"`      // 发表者id
	UserName        string `gorm:"comment:'发表者名称'"`            // 发表者名称
	IP              string `gorm:"comment:'发表者ip'"`            // 发表者ip
	IPArea          string `gorm:"comment:'ip属地'"`             // ip属地
	PID             int64  `gorm:"comment:'父评论ID'"`            // 父评论 ID
	Type            uint   `gorm:"comment:'评论类型'"`             // 评论类型，文字评论、图评等
	IsCollapsed     bool   `gorm:"default:false;comment:'折叠'"` // 折叠
	IsPending       bool   `gorm:"default:false;comment:'待审'"` // 待审
	IsPinned        bool   `gorm:"default:false;comment:'置顶'"` // 置顶
	IsHighLight     bool   `gorm:"default:false;comment:'高亮'"` // 高亮
	State           uint   `gorm:"comment:'状态'"`               // 状态
	LikeCount       uint   `gorm:"comment:'点赞数'"`              // 点赞数
	HateCount       uint   `gorm:"comment:'点踩数'"`              // 点踩数
	ReplyCount      uint   `gorm:"comment:'回复数'"`              // 回复数
	RootReplayCount uint   `gorm:"comment:'根评论回复数'"`           // 根评论回复数
	FloorCount      uint   `gorm:"comment:'评论层数'"`             // 评论层数
}

// TableName 自定义表名
func (CommentIndex) TableName() string {
	return constant.CommentIndexTableName
}

func InsertCommentIndex(commentIndex *CommentIndex) (uint, error) {
	err := db.GetClient().Create(&commentIndex).Error
	return commentIndex.ID, err
}

func InsertCommentIndexWithTx(tx *gorm.DB, commentIndex *CommentIndex) (uint, error) {
	if err := tx.Where("resource_id = ? AND pid = ?", commentIndex.ResourceID, commentIndex.PID).Last(&commentIndex).Error; err != nil {
		commentIndex.FloorCount = 1
	} else {
		commentIndex.FloorCount++
	}
	err := db.GetClient().Create(&commentIndex).Error
	return commentIndex.ID, err
}

func InsertBatchCommentIndex(commentIndexs []*CommentIndex) error {
	err := db.GetClient().Create(&commentIndexs).Error
	return err
}

func DeleteCommentIndex(commentIndex *CommentIndex) error {
	err := db.GetClient().Unscoped().Delete(&commentIndex).Error
	return err
}

func DeleteCommentIndexWithTx(tx *gorm.DB, commentID uint) error {
	err := tx.Where(constant.WhereByID, commentID).Delete(&CommentIndex{}).Error
	return err
}

func FindCommentIndexById(id int64) (CommentIndex, error) {
	var commentIndex CommentIndex
	err := db.GetClient().Where(constant.WhereByCommentID, id).First(&commentIndex).Error
	return commentIndex, err
}

func UpdateCommentIndex(commentIndex *CommentIndex) error {
	err := db.GetClient().Updates(&commentIndex).Error
	return err
}

func GetCommentIndexList(param CommentParamsList) (commentIndexs []CommentIndex, err error) {
	tx := db.GetClient()
	if param.ResourceId != 0 {
		tx = tx.Where(constant.WhereByResourceID, param.ResourceId)
	}
	if param.Pid != 0 {
		tx = tx.Where(constant.WhereByPID, param.Pid)
	}
	if param.Limit == 0 {
		param.Limit = constant.DefaultLimit
	}
	err = tx.Order(constant.OrderDescById).Limit(param.Limit).Offset(param.Offset).Find(&commentIndexs).Error
	return commentIndexs, err
}

func GetCommentListCount(param CommentParamsList) (count int64, err error) {
	tx := db.GetClient()
	if param.ResourceId != 0 {
		tx = tx.Where(constant.WhereByResourceID, param.ResourceId)
	}
	if param.Pid != 0 {
		tx = tx.Where(constant.WhereByPID, param.Pid)
	}
	err = tx.Count(&count).Error
	return count, err
}

func DeleteCommentIndexByTime(deleteTime int) error {
	err := db.GetClient().Where(constant.LessThanCreatedAt, deleteTime).Delete(&CommentIndex{}).Error
	return err
}

// DeleteChildCommentsWithTx 递归删除子评论
func DeleteChildCommentsWithTx(tx *gorm.DB, commentID uint) error {
	// 查询当前评论的所有子评论
	var childComments []CommentIndex
	if err := tx.Where("pid = ?", commentID).Find(&childComments).Error; err != nil {
		return err
	}

	// 遍历子评论
	for _, childComment := range childComments {
		// 删除子评论的内容
		if err := DeleteCommentContentWithTx(tx, childComment.ID); err != nil {
			return err
		}

		// 删除子评论的索引
		if err := DeleteCommentIndexWithTx(tx, childComment.ID); err != nil {
			return err
		}

		// 更新用户评论数量
		if err := DecreaseCommentCount(tx, uint(childComment.UserID)); err != nil {
			return err
		}

		// 递归删除子评论的子评论
		if err := DeleteChildCommentsWithTx(tx, childComment.ID); err != nil {
			return err
		}
	}

	return nil
}
