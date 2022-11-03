package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	constructor := MyLinkedList{
		Head: &Node{},
		Size: 0,
	}

	obj := Constructor()
	if !reflect.DeepEqual(obj, constructor) {
		t.Errorf("Constructor() = %v, expected %v", obj, constructor)
	}
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	if got := obj.Get(1); !reflect.DeepEqual(got, 2) {
		t.Errorf("Get() = %v, expected %v", got, 2)
	}
	obj.DeleteAtIndex(1)
	if got := obj.Get(1); !reflect.DeepEqual(got, 3) {
		t.Errorf("Get() = %v, expected %v", got, 3)
	}
}
