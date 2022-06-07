package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type NegativeFeedbackType int64
const (
	NegativeFeedbackType_UNINTERESTED_AUTHOR NegativeFeedbackType = 1
	NegativeFeedbackType_BLOCK_KEYWORDS NegativeFeedbackType = 2
	NegativeFeedbackType_LESS_SIMILAR_CONTENT NegativeFeedbackType = 3
	NegativeFeedbackType_UNINTERESTED_CONTENT_TYPE NegativeFeedbackType = 4
	NegativeFeedbackType_NOT_BELONG_SECTION NegativeFeedbackType = 5
	NegativeFeedbackType_REPEAT_CONTENT NegativeFeedbackType = 6
	NegativeFeedbackType_LOW_QUALITY_CONTENT NegativeFeedbackType = 7
)

type Msg struct {
	MemberID int64 `thrift:"member_id,1,required" db:"member_id" json:"member_id"`
	ContentType string `thrift:"content_type,2,required" db:"content_type" json:"content_type"`
	ContentID int64 `thrift:"content_id,3,required" db:"content_id" json:"content_id"`
	NegativeFeedbackType NegativeFeedbackType `thrift:"negative_feedback_type,4,required" db:"negative_feedback_type" json:"negative_feedback_type"`
	Timestamp int64 `thrift:"timestamp,5,required" db:"timestamp" json:"timestamp"`
	Extra *string `thrift:"extra,6" db:"extra" json:"extra,omitempty"`
}

func main() {
	//topic := "msg.feed-root.negative-feedback"
	topic := "data.bidding-sync-service.plutus-engine-decoupling-content"
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	//msg := &Msg{
	//	MemberID:             1,
	//	ContentType:          "ANSWER",
	//	ContentID:            1,
	//	NegativeFeedbackType: NegativeFeedbackType_BLOCK_KEYWORDS,
	//	Timestamp:            time.Now().Unix(),
	//	Extra:                nil,
	//}
	//jMsg, _ := json.Marshal(msg)  // 105475
	//jMsg := `{"member_id":101561295,"content_type":"ANSWER","content_id":462786461,"negative_feedback_type":"LESS_SIMILAR_CONTENT","timestamp":1646815687238,"extra":null,"setNegativeFeedbackType":true,"negativeFeedbackType":"LESS_SIMILAR_CONTENT","memberId":101561295,"contentId":462786461,"setContentId":true,"setExtra":false,"setMemberId":true,"setContentType":true,"setTimestamp":true,"contentType":"ANSWER"}`
	jMsg := `{
  "msg_id": 717152151134285827,
  "msg_version": 1000014,
  "msg_type": "AD_UPDATE",
  "ad_id": 975911,
  "campaign_id": 510316,
  "user_id": 11,
  "agent_id": 7,
  "ad_name": "contentMarking_ad_1651145694017",
  "template_name": "card_app_feed",
  "date_begin": "2022-04-28",
  "date_end": "",
  "status": 1,
  "zone_ids": "[9999,8899,10001,10002,10003,10005,105,8,1297,1298,1549,1550,10007,10008]",
  "position_ids": "[1]",
  "targeting": "{\"os\":[],\"network\":[],\"gender\":[],\"area\":{},\"interest\":[],\"keyword\":[],\"time\":{},\"mobileOperator\":{},\"equipmentPrice\":[],\"appCategory\":[],\"orientationCrowd\":[],\"excludeCrowd\":[\"6670\",\"121\",\"79\"],\"industry\":{},\"industry1\":\"\",\"industry2\":\"\",\"industry3\":[],\"crowdPaid\":1,\"appInclusion\":[],\"appExclusion\":[],\"locationArea\":{},\"allArea\":{},\"excludeAppCategory\":[],\"age\":[]}",
  "bid_type": 0,
  "price": 200,
  "search_bid": 200,
  "campaign_daily_budget": 222200,
  "product_id": 913,
  "product_uid": 1533338112728582,
  "industry_id": 100010059,
  "user_industry_id": 100010015,
  "strategy": 2,
  "campaign_serving_level": "{\"2022-04-28\":0,\"2022-04-29\":0}",
  "user_serving_level": "{\"2022-04-28\":3,\"2022-04-29\":3}",
  "explore": 0,
  "probability": 100,
  "creative_targeting": "[{\"asset\":\"{\\\"logo\\\":{\\\"w\\\":250,\\\"h\\\":250,\\\"token\\\":\\\"v2-725659c077fa2d7244993ecf02ff94e6\\\",\\\"format\\\":\\\"png\\\",\\\"path\\\":\\\"https://pic3.zhimg.com/v2-725659c077fa2d7244993ecf02ff94e6_250x250.png\\\",\\\"url\\\":\\\"https://pic3.zhimg.com/v2-725659c077fa2d7244993ecf02ff94e6_250x250.png\\\"},\\\"sponsored\\\":{\\\"value\\\":\\\"效果引擎测试产品_zj\\\"},\\\"title\\\":{\\\"value\\\":\\\"测试回答？\\\"}}\",\"landing_url\":\"https://www.zhihu.dev/question/6512/answer/4915\",\"imptrackers\":\"[]\",\"clicktrackers\":\"[]\",\"validtrackers\":\"[]\",\"uid\":49,\"content_attach_uid\":0,\"entity_id\":880,\"creative_id\":21071910,\"new_asset\":\"{\\\"id\\\":21071910,\\\"brand_name\\\":\\\"效果引擎测试产品_zj\\\",\\\"brand_logo\\\":\\\"https://pic3.zhimg.com/v2-725659c077fa2d7244993ecf02ff94e6_250x250.png\\\",\\\"title\\\":\\\"测试回答？\\\",\\\"desc\\\":\\\"\\\",\\\"img_size\\\":0,\\\"landing_url\\\":\\\"https://www.zhihu.dev/question/6512/answer/4915\\\",\\\"cta\\\":\\\"\\\",\\\"img_full_screen\\\":false}\",\"author_full_name\":\"robot1\",\"author_avatar_path\":\"https://pic4.zhimg.com/v2-abed1a8c04700ba7d72b45195223e0ff.jpeg\",\"content_group_id\":0,\"content_plugin_relation\":\"[\\\"2-99\\\"]\",\"activity_info\":\"\",\"zm_status\":2}]",
  "target_type": "content",
  "template_id": 6,
  "parent_id": 0,
  "msg_time": 1651148890830,
  "silence_download": 0,
  "office_only": 0,
  "apk_name": "",
  "bundle_id": "",
  "ocpc_status": 0,
  "conversion_price": 0,
  "new_template_id": 11,
  "new_template_name": "big_image_plutus",
  "ocpx_goal": 0,
  "ocpx_second_start": 0,
  "deep_conversion_goal": 0,
  "deep_conversion_type": 0,
  "deep_conversion_status": 0,
  "deep_conversion_price": 0,
  "ad_update_time": 1651145694,
  "put_speed": 0,
  "uniform_serving_level": 0,
  "user_daily_budget": 34000,
  "user_support_stage": 1,
  "user_support_start_date": 1583749451000,
  "search_word": "[]",
  "bid_coefficient": "{\"fuzzy\":-1,\"exact\":-1,\"content_page\":-1.0}",
  "is_smart_targeting": 0,
  "experiment_count": 0,
  "zone_bid_factor": "[{\"zoneId\":9999,\"factor\":100,\"valid\":0},{\"zoneId\":8899,\"factor\":100,\"valid\":0},{\"zoneId\":10001,\"factor\":100,\"valid\":0},{\"zoneId\":10005,\"factor\":100,\"valid\":1},{\"zoneId\":105,\"factor\":100,\"valid\":0},{\"zoneId\":8,\"factor\":100,\"valid\":0},{\"zoneId\":1297,\"factor\":100,\"valid\":0},{\"zoneId\":1298,\"factor\":100,\"valid\":0},{\"zoneId\":1549,\"factor\":100,\"valid\":0},{\"zoneId\":1550,\"factor\":100,\"valid\":0},{\"zoneId\":10007,\"factor\":100,\"valid\":0},{\"zoneId\":10008,\"factor\":100,\"valid\":0}]",
  "intelligent_conv_level": -1,
  "biz_type": 1,
  "content_target": 0,
  "zm_serving_level": "{\"2022-04-28\":4,\"2022-04-29\":4}",
  "dpa_info": ""
}`

	m := &sarama.ProducerMessage{}
	m.Topic = topic
	m.Value = sarama.StringEncoder(jMsg)

	producer, err := sarama.NewSyncProducer([]string{"kafka-dev01.dev.rack.zhihu.com:3390", "kafka-dev02.dev.rack.zhihu.com:3390", "kafka-dev03.dev.rack.zhihu.com:3390"}, config)
	if err != nil {
		fmt.Printf("producer err: %+v\n", err)
		return
	}
	defer producer.Close()

	message, offset, err := producer.SendMessage(m)
	if err != nil {
		fmt.Printf("send msg err: %+v\n", err)
		return
	}
	fmt.Printf("message: %+v, offset: %d", message, offset)
}