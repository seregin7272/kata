package llist

import "testing"

func TestPutOk(t *testing.T) {

	lfu := Constructor(2)

	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Put(3, 3)

	if lfu.Get(1) != -1 {
		t.Errorf("expected -1 but got 1 %d", lfu.Get(1))
	}

	if lfu.Get(2) != 2 {
		t.Error("element should be 2")
	}

	if lfu.Get(3) != 3 {
		t.Error("element should be 3")
	}

}

func TestGetOk1(t *testing.T) {

	lfu := Constructor(2)

	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Get(1)
	lfu.Put(3, 3)

	if lfu.Get(2) != -1 {
		t.Errorf("expected -1 but got %d", lfu.Get(2))
	}

	if lfu.Get(1) != 1 {
		t.Errorf("expected 1 but got %d", lfu.Get(1))
	}

	if lfu.Get(3) != 3 {
		t.Error("element should be 3")
	}

}

func TestGetOk2(t *testing.T) {

	lfu := Constructor(2)

	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Get(1)
	lfu.Put(3, 3)

	if lfu.Get(2) != -1 {
		t.Errorf("expected -1 but got %d", lfu.Get(2))
	}

	if lfu.Get(1) != 1 {
		t.Errorf("expected 1 but got %d", lfu.Get(1))
	}

	if lfu.Get(3) != 3 {
		t.Error("element should be 3")
	}

}

func TestGetOk3(t *testing.T) {

	lfu := Constructor(2)

	lfu.Put(1, 1) // cache=[1,_], cnt(1)=1
	lfu.Put(2, 2) //  cache=[2,1], cnt(2)=1, cnt(1)=1

	val := lfu.Get(1) // 1 // cache=[1,2], cnt(2)=1, cnt(1)=2
	if val != 1 {
		t.Errorf("expected 1 but got %d", val)
	}

	lfu.Put(3, 3) // cache=[3,1], cnt(3)=1, cnt(1)=2

	val = lfu.Get(2) // -1
	if val != -1 {
		t.Errorf("expected -1 but got %d", val)
	}

	val = lfu.Get(3) // 3 // cache=[3,1], cnt(3)=2, cnt(1)=2
	if val != 3 {
		t.Errorf("expected 3 but got %d", val)
	}

	lfu.Put(4, 4) // cache=[4,3], cnt(4)=1, cnt(3)=2

	val = lfu.Get(1) // -1
	if val != -1 {
		t.Errorf("expected -1 but got %d", val)
	}
	val = lfu.Get(3) // 3 // cache=[3,4], cnt(4)=1, cnt(3)=3
	if val != 3 {
		t.Errorf("expected 3 but got %d", val)
	}

	val = lfu.Get(4) // 4 // cache=[4,3], cnt(4)=2, cnt(3)=3
	if val != 4 {
		t.Errorf("expected 4 but got %d", val)
	}

}

func TestGetOk4(t *testing.T) {

	lfu := Constructor(2)

	lfu.Put(2, 1) // cache=[2,_], cnt(2)=1
	lfu.Put(2, 2) //  cache=[2,_], cnt(2)=2

	val := lfu.Get(2) // 2 //  cache=[2,_], cnt(2)=3
	if val != 2 {
		t.Errorf("expected 2 but got %d", val)
	}

	lfu.Put(1, 1) //  cache=[2,1], cnt(2)=3 cnt(1)=1
	lfu.Put(3, 3) // cache=[3,2], cnt(3)=1, cnt(2)=3

	val = lfu.Get(2) // 2
	if val != 2 {
		t.Errorf("expected 2 but got %d", val)
	}
	val = lfu.Get(3) // 3 // cache=[3,4], cnt(4)=1, cnt(3)=3
	if val != 3 {
		t.Errorf("expected 3 but got %d", val)
	}

}
