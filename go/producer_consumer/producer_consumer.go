package producer_consumer

func ProduceConsume(produce func() interface{}, consume func(interface{}), producerNum, consumerNum, capcity int) {
	buf := make(chan interface{}, capcity)
	produceWrapper := func(ch chan<- interface{}) {
		for {
			ch <- produce()
		}
	}
}
