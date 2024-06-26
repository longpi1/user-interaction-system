package constant

// 资源类型相关常量
const (
	// ResourceTopic 话题
	ResourceTopic = "topic"
	// ResourceNews 新闻
	ResourceNews = "news"
	// ResourceArticle 文章
	ResourceArticle = "Article"
	// ResourceVideo 视频
	ResourceVideo = "video"
	// ResourceComment 评论
	ResourceComment = "comment"
)

// 状态相关数据常量
const (
	StateWaitReview = iota
	StateReviewed
	StateReviewFail
	StateDelete
)

// 展示相关数据常量
const (
	DisplayOk = 1
)
