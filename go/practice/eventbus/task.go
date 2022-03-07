/*
 * @Author: 馬濤
 * @Date: 2021-01-15 11:55:10
 * @LastEditTime: 2021-01-15 13:45:17
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: /cf/eventbus/task.go
 * @MT is your father.
 */

package eventbus

// 任务
type task struct {
	// 执行函数
	f func(handler *EventHandler, args ...interface{})
	// 处理
	eventHandler *EventHandler
	// 参数
	args []interface{}
}
