package Tree

import (
	"reflect"
	"testing"
)

func Test_inOrder(t *testing.T) {

	tests := []struct {
		name  string
		input TreeNode
		want  []int
	}{
		{
			name:  "nil case",
			input: TreeNode{},
			want:  []int{0}, //这里因为初始化时候会默认赋0值，所以会有一个元素
		}, {
			name: "single element",
			input: TreeNode{
				value: 19,
				left:  nil,
				right: nil,
			},
			want: []int{19},
		}, {
			name: "three elements",
			input: TreeNode{
				value: 1,
				left: &TreeNode{
					value: 2,
					left:  nil,
					right: nil,
				},
				right: &TreeNode{
					value: 3,
					left:  nil,
					right: nil,
				},
			},
			want: []int{2, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inOrder(&tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postOrder(t *testing.T) {
	tests := []struct {
		name  string
		input TreeNode
		want  []int
	}{
		{
			name:  "nil case",
			input: TreeNode{},
			want:  []int{0}, //这里因为初始化时候会默认赋0值，所以会有一个元素
		}, {
			name: "single element",
			input: TreeNode{
				value: 19,
				left:  nil,
				right: nil,
			},
			want: []int{19},
		}, {
			name: "three elements",
			input: TreeNode{
				value: 1,
				left: &TreeNode{
					value: 2,
					left:  nil,
					right: nil,
				},
				right: &TreeNode{
					value: 3,
					left:  nil,
					right: nil,
				},
			},
			want: []int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postOrder(&tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preOrder(t *testing.T) {
	tests := []struct {
		name  string
		input TreeNode
		want  []int
	}{
		{
			name:  "nil case",
			input: TreeNode{},
			want:  []int{0}, //这里因为初始化时候会默认赋0值，所以会有一个元素
		}, {
			name: "single element",
			input: TreeNode{
				value: 19,
				left:  nil,
				right: nil,
			},
			want: []int{19},
		}, {
			name: "three elements",
			input: TreeNode{
				value: 1,
				left: &TreeNode{
					value: 2,
					left:  nil,
					right: nil,
				},
				right: &TreeNode{
					value: 3,
					left:  nil,
					right: nil,
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preOrder(&tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
