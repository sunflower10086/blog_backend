package constants

// MQProducer Sender 的 key
const (
	MQProducerPageViewName   = "page_view"
	MQProducerLikeActionName = "like_event"
)

// MQProducer Receiver 的 key (配置文件中的 name 和 topic 进行拼接，name_topic )
// 如果后续添加了 consumer 则需要在此处添加对应的 name
const (
	MQConsumerPageViewName   = "page_view_consumer_blog.page_views"
	MQConsumerLikeActionName = "like_event_consumer_blog.like_events"
)
