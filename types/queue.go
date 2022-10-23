package types

// QueueItem Элемент очереди, а очередь это массив таких элементов
type QueueItem struct {
	Id          int    `json:"Id"`
	Time        string `json:"Time"`
	Service     string `json:"Service"`
	Workstation int    `json:"Workstation"`
	Status      int    `json:"Status"`
}
