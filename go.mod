module github.com/longpi1/user-interaction-system

go 1.22.2

replace (
	github.com/longpi1/user-interaction-system/comment-job => ./comment/comment-job
	github.com/longpi1/user-interaction-system/comment-service => ./comment/comment-service
)

require github.com/go-kratos/aegis v0.2.0 // indirect
