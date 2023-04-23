package server

import (
	"container/list"
	"fmt"
)

type KQueue struct {
	Queue *list.List
}

func (c *KQueue) Enqueue(value Channel) {
	c.Queue.PushBack(value)
}

func (c *KQueue) Dequeue() error {
	if c.Queue.Len() > 0 {
		ele := c.Queue.Front()
		c.Queue.Remove(ele)
	}
	return fmt.Errorf("POP ERR: EMPTY")
}

func (c *KQueue) Front() (Channel, error) {
	if c.Queue.Len() > 0 {
		if val, ok := c.Queue.Front().Value.(Channel); ok {
			return val, nil
		}
		return Channel{}, fmt.Errorf("PEEP ERR: Queue DATATYPE IS INCORRECT")
	}
	return Channel{}, fmt.Errorf("PEEP ERROR: Queue IS EMPTY")
}

func (c *KQueue) Size() int {
	return c.Queue.Len()
}

func (c *KQueue) Empty() bool {
	return c.Queue.Len() == 0
}
