package queue

type Node struct {
	Value interface{}
	Next  *Node
}

type Queue struct {
	Head *Node
	Tail *Node
	Size int64
}

func Init() *Queue {
	queue := &Queue{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
	return queue
}

func (queue *Queue) Put(value interface{}) {
	node := &Node{
		Value: value,
		Next:  nil,
	}
	tmpTail := queue.Tail
	// 空队列
	if tmpTail == nil {
		queue.Head = node
		queue.Tail = node
	} else {
		tmpTail.Next = node
		queue.Tail = node
	}
	queue.Size++
}

func (queue *Queue) Pop() interface{} {
	if queue.Head == nil {
		return nil
	}
	node := queue.Head
	queue.Head = node.Next
	queue.Size--
	if queue.Size == 0 {
		queue.Tail = nil
	}
	return node.Value
}

func (queue *Queue) Empty() bool {
	if queue.Size <= 0 {
		return true
	}
	return false
}

func (queue *Queue) ForEach(process func(value interface{}) error) error {
	tmp := queue.Head
	for {
		err := process(tmp.Value)
		if err != nil {
			return err
		}
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	return nil
}
