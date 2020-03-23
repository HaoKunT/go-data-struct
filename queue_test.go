package datastruct

import (
	"fmt"
	"testing"
)

func TestQueueEn(t *testing.T) {
	q := NewQueue()
	q.En(1)
	q.En(2)
	q.En(3)

	if q.GetLength() != 3 {
		t.Errorf("Queue: En error: length should be 3, but get %d", q.GetLength())
	}
}

func TestQueueDe(t *testing.T) {
	q := NewQueue()
	q.En(1)
	q.En(2)
	q.En(3)

	if o := q.De(); o != 1 {
		t.Errorf("Queue: De error: it should get 1, but get %d", o)
	}
}

func TestQueueDeIsEmpty(t *testing.T) {
	q := NewQueue()

	if !q.IsEmpty() {
		t.Errorf("Queue: IsEmpty error: it should be true, but get %v", q.IsEmpty())
	}

	fmt.Println(q)

	q.En(1)
	q.En(2)
	q.En(3)

	if q.IsEmpty() {
		t.Errorf("Queue: IsEmpty error: it should be false, but get %v", q.IsEmpty())
	}
}
