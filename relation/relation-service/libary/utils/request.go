package utils

import "relation-service/libary/conf"

// ConvertType 将string类型的类型转换为int
func ConvertType(relationType string) int {
	typeMap := conf.GetMapConfig().TypeMap
	return typeMap[relationType]
}

// ConvertPlatform 将string类型的平台转换为int
func ConvertPlatform(Platform string) int {
	platform := conf.GetMapConfig().PlatformMap
	return platform[Platform]
}
