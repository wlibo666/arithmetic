package queue

import (
	"testing"
)

func TestInit(t *testing.T) {
	queue := Init()
	t.Logf("init queue size:%d", queue.Size)

	queue.Put(123)
	queue.Put(456)
	queue.Put(789)
	t.Logf("queue size:%d", queue.Size)
	node := queue.Pop()
	t.Logf("pop:%v,size:%d", node, queue.Size)
	node = queue.Pop()
	t.Logf("pop:%v,size:%d", node, queue.Size)

	queue.Put("aaa")
	queue.Put("bbb")
	node = queue.Pop()
	t.Logf("pop:%v,size:%d", node, queue.Size)

	queue.ForEach(func(node interface{}) error {
		t.Logf("node value:%v\n", node)
		return nil
	})
}
