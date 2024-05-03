package controller

import (
	"relation-service/libary/conf"

	"github.com/longpi1/gopkg/libary/log"
)

func VerifyTypeAndPlatform(relationType int, platform int) bool {
	mapConfig := conf.GetMapConfig()
	var hasType, hasPlatform bool
	for _, value := range mapConfig.TypeMap {
		if value == relationType {
			hasType = true
		}
	}

	for _, value := range mapConfig.PlatformMap {
		if value == platform {
			hasPlatform = true
		}
	}
	if !hasType {
		log.Error("类型不存在")
	}
	if !hasPlatform {
		log.Error("平台不存在")
	}
	return hasType && hasPlatform
}
