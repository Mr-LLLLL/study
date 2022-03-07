/*
 * @Author: 馬濤
 * @Date: 2021-01-23 11:29:41
 * @LastEditTime: 2021-01-23 17:50:58
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: /cf/eventbus/dirver/options/server.go
 * @MT is your father.
 */

package options

// Server Server
type Server struct {
	Addrs []string
	Topic string
}
