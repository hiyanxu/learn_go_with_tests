package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)
func main() {
	// 线上配置.
	data := `{
"storage": {
  "name": "content-deliver-cost-data",
  "conn": "redis://:8902ec5cb4fa8dbc58a0d739036fcac8@wredis-content-deliver-content-deliver-cost-data.global.rack.zhihu.com:26379",
  "timeout": 1000
}
}`

	// 测试环境配置.
//	data := `{
//  "storage": {
//    "name": "content-deliver-storage",
//    "conn": "redis://redis-content-deliver-content-deliver-storage.aspasia.svc.cluster.local:6379",
//    "timeout": 1000
//  },
//  "pubsub_cache": {
//    "name": "deliver-pubsub",
//    "config": {
//      "master": "redis://redis-content-deliver-content-deliver-rpc.aspasia.svc.cluster.local:6379",
//      "slave": [
//        "redis://redis-content-deliver-content-deliver-rpc.aspasia.svc.cluster.local:6379"
//      ]
//    }
//  }
//}`
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sEnc2, _ := base64.StdEncoding.DecodeString("ewogICJzdG9yYWdlIjogewogICAgIm5hbWUiOiAiY29udGVudC1kZWxpdmVyLXN0b3JhZ2UiLAogICAgImNvbm4iOiAicmVkaXM6Ly9yZWRpcy1jb250ZW50LWRlbGl2ZXItY29udGVudC1kZWxpdmVyLXN0b3JhZ2UuYXNwYXNpYS5zdmMuY2x1c3Rlci5sb2NhbDo2Mzc5IiwKICAgICJ0aW1lb3V0IjogMTAwMAogIH0sCiAgImNhY2hlIjogewogICAgIm5hbWUiOiAiY29udGVudC1kZWxpdmVyLWNhY2hlIiwKICAgICJjb25uIjogInJlZGlzOi8vcmVkaXMtY29udGVudC1kZWxpdmVyLWNvbnRlbnQtZGVsaXZlci1jYWNoZS5hc3Bhc2lhLnN2Yy5jbHVzdGVyLmxvY2FsOjYzNzkiLAogICAgInRpbWVvdXQiOiAxMDAwCiAgfSwKICAicHVic3ViIjogewogICAgIm5hbWUiOiAiZGVsaXZlci1wdWJzdWIiLAogICAgImNvbm4iOiAicmVkaXM6Ly9yZWRpcy1jb250ZW50LWRlbGl2ZXItY29udGVudC1kZWxpdmVyLXJwYy5hc3Bhc2lhLnN2Yy5jbHVzdGVyLmxvY2FsOjYzNzkiLAogICAgInRpbWVvdXQiOiAxMDAwCiAgfSwKICAibmxwIjogewogICAgIm5hbWUiOiAiYWQtbmxwLWNyb3dkLXRyaWdnZXIiLAogICAgImNvbm4iOiAicmVkaXM6Ly9yZWRpcy1hZC1ubHAtY3Jvd2QtdHJpZ2dlci10ZXN0LmFzcGFzaWEuc3ZjLmNsdXN0ZXIubG9jYWw6NjM3OSIsCiAgICAidGltZW91dCI6IDEwMDAKICB9Cn0=")
	fmt.Println(string(sEnc2))

	d, err := strconv.ParseInt("1", 10, 64)
	fmt.Println(d, err)

	t := time.Now().AddDate(0, 0, -31).Unix()
	fmt.Println(t)
}
