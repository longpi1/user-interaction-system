package event

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"

	"github.com/mlogclub/simple/common/jsons"
)

func TestEvent(t *testing.T) {
	RegisterHandler(reflect.TypeOf(model.CommentIndex{}), func(i interface{}) {
		fmt.Println("处理用户1")
		fmt.Println(jsons.ToStr(i))
	})
	RegisterHandler(reflect.TypeOf(model.CommentIndex{}), func(i interface{}) {
		fmt.Println("处理用户2")
		fmt.Println(jsons.ToStr(i))
	})
	Send(model.CommentIndex{})
	//w.Wait()
}
