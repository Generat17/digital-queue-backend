package types

// QueueItem Элемент очереди, а очередь это массив таких элементов
type QueueItem struct {
	Id      int    `json:"Id"`
	Time    string `json:"Time"`
	Service string `json:"Service"`
}

// QueueItemNumber Структура, нужная для возвращения значения в API
type QueueItemNumber struct {
	Ticket int `json:"TicketID"`
}
