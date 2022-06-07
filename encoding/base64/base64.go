package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)
func main() {
	data := `{
  "storage": {
    "name": "content-deliver-storage",
    "conn": "redis://:4111b9434afab6db81ac27b7a403d722@wredis-content-deliver-content-deliver-storage.global.rack.zhihu.com:26379",
    "timeout": 1000
  },
  "cache": {
    "name": "content-deliver-cache",
    "conn": "redis://:fed5077226a1e209792382849b78c84d@credis-content-deliver-cluster3036666.pek01.rack.zhihu.com:26379",
    "timeout": 1000
  },
  "pubsub": {
    "name": "deliver-pubsub",
    "conn": "redis://:1de84cceb87dbc17abb2dd9726593929@hredis-content-deliver-storage3090938.pek01.rack.zhihu.com:35197",
    "timeout": 1000
  },
  "nlp": {
    "name": "ad-nlp-crowd-trigger",
    "conn": "redis://:a9b854b5329dde5f5af2a6ebb0e5d309@rredis-ad-nlp-crowd-trigger-test.pek01.rack.zhihu.com:26379",
    "timeout": 1000
  }
}`
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sEnc2, _ := base64.StdEncoding.DecodeString("ewogICJzdG9yYWdlIjogewogICAgIm5hbWUiOiAiY29udGVudC1kZWxpdmVyLXN0b3JhZ2UiLAogICAgImNvbm4iOiAicmVkaXM6Ly9yZWRpcy1jb250ZW50LWRlbGl2ZXItY29udGVudC1kZWxpdmVyLXN0b3JhZ2UuYXNwYXNpYS5zdmMuY2x1c3Rlci5sb2NhbDo2Mzc5IiwKICAgICJ0aW1lb3V0IjogMTAwMAogIH0sCiAgImNhY2hlIjogewogICAgIm5hbWUiOiAiY29udGVudC1kZWxpdmVyLWNhY2hlIiwKICAgICJjb25uIjogInJlZGlzOi8vcmVkaXMtY29udGVudC1kZWxpdmVyLWNvbnRlbnQtZGVsaXZlci1jYWNoZS5hc3Bhc2lhLnN2Yy5jbHVzdGVyLmxvY2FsOjYzNzkiLAogICAgInRpbWVvdXQiOiAxMDAwCiAgfSwKICAicHVic3ViIjogewogICAgIm5hbWUiOiAiZGVsaXZlci1wdWJzdWIiLAogICAgImNvbm4iOiAicmVkaXM6Ly9yZWRpcy1jb250ZW50LWRlbGl2ZXItY29udGVudC1kZWxpdmVyLXJwYy5hc3Bhc2lhLnN2Yy5jbHVzdGVyLmxvY2FsOjYzNzkiLAogICAgInRpbWVvdXQiOiAxMDAwCiAgfSwKICAibmxwIjogewogICAgIm5hbWUiOiAiYWQtbmxwLWNyb3dkLXRyaWdnZXIiLAogICAgImNvbm4iOiAicmVkaXM6Ly9yZWRpcy1hZC1ubHAtY3Jvd2QtdHJpZ2dlci10ZXN0LmFzcGFzaWEuc3ZjLmNsdXN0ZXIubG9jYWw6NjM3OSIsCiAgICAidGltZW91dCI6IDEwMDAKICB9Cn0=")
	fmt.Println(string(sEnc2))

	d, err := strconv.ParseInt("1", 10, 64)
	fmt.Println(d, err)

	t := time.Now().AddDate(0, 0, -31).Unix()
	fmt.Println(t)
}
