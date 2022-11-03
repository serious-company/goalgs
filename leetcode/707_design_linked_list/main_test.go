package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type addAtIndex struct {
		index int
		val   int
	}

	type args struct {
		Get           int
		AddAtHead     int
		AddAtTail     int
		AddAtIndex    addAtIndex
		DeleteAtIndex int
	}

	type expected struct {
		Constructor MyLinkedList
		Get1        int
		Get2        int
	}

	tests := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "test-1",
			args: args{
				AddAtHead: 1,
				AddAtTail: 3,
				AddAtIndex: addAtIndex{
					index: 1,
					val:   2,
				},
				Get:           1,
				DeleteAtIndex: 1,
			},
			expected: expected{
				Constructor: MyLinkedList{
					Head: &Node{},
					Size: 0,
				},
				Get1: 2,
				Get2: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := Constructor()

			if !reflect.DeepEqual(obj, tt.expected.Constructor) {
				t.Errorf("Constructor() = %v, expected %v", obj, tt.expected.Constructor)
			}
			obj.AddAtHead(tt.args.AddAtHead)
			obj.AddAtTail(tt.args.AddAtTail)
			obj.AddAtIndex(tt.args.AddAtIndex.index, tt.args.AddAtIndex.val)
			if got := obj.Get(tt.args.Get); !reflect.DeepEqual(got, tt.expected.Get1) {
				t.Errorf("Get() = %v, expected %v", got, tt.expected.Get2)
			}
			obj.DeleteAtIndex(tt.args.DeleteAtIndex)
			if got := obj.Get(tt.args.Get); !reflect.DeepEqual(got, tt.expected.Get2) {
				t.Errorf("Get() = %v, expected %v", got, tt.expected.Get2)
			}
		})
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
