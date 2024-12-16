package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type NegativeFeedbackType int64

const (
	NegativeFeedbackType_UNINTERESTED_AUTHOR       NegativeFeedbackType = 1
	NegativeFeedbackType_BLOCK_KEYWORDS            NegativeFeedbackType = 2
	NegativeFeedbackType_LESS_SIMILAR_CONTENT      NegativeFeedbackType = 3
	NegativeFeedbackType_UNINTERESTED_CONTENT_TYPE NegativeFeedbackType = 4
	NegativeFeedbackType_NOT_BELONG_SECTION        NegativeFeedbackType = 5
	NegativeFeedbackType_REPEAT_CONTENT            NegativeFeedbackType = 6
	NegativeFeedbackType_LOW_QUALITY_CONTENT       NegativeFeedbackType = 7
)

type Msg struct {
	MemberID             int64                `thrift:"member_id,1,required" db:"member_id" json:"member_id"`
	ContentType          string               `thrift:"content_type,2,required" db:"content_type" json:"content_type"`
	ContentID            int64                `thrift:"content_id,3,required" db:"content_id" json:"content_id"`
	NegativeFeedbackType NegativeFeedbackType `thrift:"negative_feedback_type,4,required" db:"negative_feedback_type" json:"negative_feedback_type"`
	Timestamp            int64                `thrift:"timestamp,5,required" db:"timestamp" json:"timestamp"`
	Extra                *string              `thrift:"extra,6" db:"extra" json:"extra,omitempty"`
}

func main() {
	//topic := "msg.feed-root.negative-feedback"
	//topic := "data.bidding-sync-service.plutus-engine-decoupling-content"
	topic := "data.bidding-sync-service.plutus-engine-decoupling-creative"
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
	   {"msg_id":975672,"msg_version":2985710074,"msg_type":"AD_ADD","msg_time":1635246945507,"ad_id":975672,"campaign_id":8860,"user_id":8860,"agent_id":8860,"media_type":"[1,2,3]","ad_name":"contentMarking_ad_1611302815235","template_name":"card_app_feed","new_template_name":"big_image_plutus","date_begin":"2021-10-26","date_end":"","status":1,"zone_ids":"[9999,10001,1297,1298,1549,1550]","position_ids":"[1, 2]","bid_type":5,"price":20,"ocpx_status":0,"campaign_daily_budget":111100,"user_daily_budget":1000000000,"product_id":8860,"campaign_serving_level":"{\"2021-10-26\": 0}","user_serving_level":"{\"2021-10-26\": 0}","zm_serving_level":"{\"2021-10-26\": 0}","creative_targeting":"[{\"creative_id\":21072155,\"searchWords\":[],\"version\":214,\"template_id\":7,\"template_name\":\"card_app_feed_word\",\"new_template_id\":48,\"new_template_name\":\"app_feed_word\", \"entity_id\": 423353, \"landing_url\": \"https://zhuanlan.zhihu.dev/p/43761393\", \"content-attach-uid\":22222}]","target_type":"content","template_id":8860,"new_template_id":8860,"ad_update_time":8860,"targeting":"{\"searchWords\":[\"文章\",\"标题\"], \"os\":[],\"network\":[],\"gender\":[],\"area\":{\"310000\":[150700, 150800],\"120000\":[]},\"locationArea\":{\"110000\":[],\"140000\":[]},\"allArea\":{\"810000\":[810002],\"820000\":[820001]},\"interest\":[],\"keyword\":[],\"time\":{},\"industry\":{}}","zone_bid_factor":"[{\"zoneId\":10001,\"factor\":100,\"valid\":1},{\"zoneId\":8899,\"factor\":100,\"valid\":1},{\"zoneId\":8,\"factor\":100,\"valid\":1},{\"zoneId\":105,\"factor\":50,\"valid\":1},{\"zoneId\":9999,\"factor\":100,\"valid\":1},{\"zoneId\":10005,\"factor\":100,\"valid\":1}]","content_target":0,"biz_type":1,"industry_id":8004,"user_industry_id":16000,"search_bid":0,"product_uid":8860,"strategy":2,"explore":0,"probability":100,"parent_id":0,"silence_download":0,"office_only":0,"apk_name":"apk_name","bundle_id":"bundle_id","ocpc_status":0,"conversion_price":1111,"ocpx_goal":14,"ocpx_second_start":0,"user_support_stage":2,"put_speed":0,"uniform_serving_level":0,"user_support_start_date":8860,"search_word":"[]","bid_coefficient":"{\"fuzzy\":-1,\"exact\":-1,\"content_page\":-1.0}","is_smart_targeting":0,"experiment_count":0,"intelligent_conv_level":-1,"traffic_source":2,"media_type":"[1,2,3]"}
	}`
	//	jMsg := `{
	//  "customer_id": 2,
	//  "customer_token": "1oF0A87NxGVmEzHc",
	//  "qps": 100,
	//  "advertiser_ids": "23137,23093,23192,23084",
	//  "bid_url": "https://promotion-partner.kuaishou.com/rest/n/rta/zhihu",
	//  "block_rule": "{\"time\":[],\"zoneIds\":[],\"frequence\":0,\"invalidRequestAvailable\":1}",
	//  "request_fields": "",
	//  "status": 1,
	//  "msg_version": 1,
	//  "msg_id": 1111,
	//  "msg_type": "RTA_INFO_UPDATE"
	//}`

	m := &sarama.ProducerMessage{}
	m.Topic = topic
	m.Value = sarama.StringEncoder(jMsg)

	producer, err := sarama.NewSyncProducer([]string{"", "", ""}, config)
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
