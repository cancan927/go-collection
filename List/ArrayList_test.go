package List

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayList_Append(t *testing.T) {

	tests := []struct {
		name   string
		origin ArrayList[int]
		addVal int
		want   ArrayList[int]
	}{
		{
			name:   "empty case",
			origin: ArrayList[int]{},
			addVal: 10,
			want:   ArrayList[int]{[]int{10}, 1},
		}, {
			name:   "one element",
			origin: ArrayList[int]{[]int{10}, 1},
			addVal: 20,
			want:   ArrayList[int]{[]int{10, 20}, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.origin.Append(tt.addVal)
			assert.Equal(t, tt.want, tt.origin)

		})
	}
}

func TestArrayList_Clear(t *testing.T) {

	tests := []struct {
		name   string
		origin ArrayList[int]
		want   ArrayList[int]
	}{
		{
			name:   "empty case",
			origin: ArrayList[int]{},
			want:   ArrayList[int]{[]int{}, 0},
		}, {
			name:   "one element",
			origin: ArrayList[int]{[]int{10}, 1},
			want:   ArrayList[int]{[]int{}, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.origin.Clear()
			assert.Equal(t, tt.want, tt.origin)
		})
	}
}

func TestArrayList_Contains(t *testing.T) {
	tests := []struct {
		name   string
		origin ArrayList[int]
		target int
		want   bool
	}{
		{
			name:   "not exist case",
			origin: ArrayList[int]{},
			target: 10,
			want:   false,
		}, {
			name:   "exist case",
			origin: ArrayList[int]{[]int{10}, 1},
			target: 10,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.origin.Contains(tt.target)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestArrayList_Delete(t *testing.T) {
	idxErr := errors.New("index out of bounds")
	tests := []struct {
		name        string
		origin      ArrayList[int]
		targetIndex int
		want        ArrayList[int]
		wantErr     error
	}{
		{
			name:        "empty case",
			origin:      ArrayList[int]{},
			targetIndex: 10,
			want:        ArrayList[int]{},
			wantErr:     idxErr,
		}, {
			name:        "exist case",
			origin:      ArrayList[int]{[]int{10}, 1},
			targetIndex: 0,
			want:        ArrayList[int]{[]int{}, 0},
			wantErr:     nil,
		}, {
			name:        "normal case",
			origin:      ArrayList[int]{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9},
			targetIndex: 4,
			want:        ArrayList[int]{[]int{1, 2, 3, 4, 6, 7, 8, 9}, 8},
			wantErr:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.origin.Delete(tt.targetIndex)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.origin, tt.want)
		})
	}
}

func TestArrayList_Get(t *testing.T) {
	idxErr := errors.New("index out of bounds")
	tests := []struct {
		name    string
		origin  ArrayList[int]
		index   int
		want    int
		wantErr error
	}{
		{
			name:    "empty case",
			origin:  ArrayList[int]{},
			index:   10,
			want:    0, //默认值是0
			wantErr: idxErr,
		}, {
			name:    "exist case",
			origin:  ArrayList[int]{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9},
			index:   4,
			want:    5,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.origin.Get(tt.index)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestArrayList_Insert(t *testing.T) {
	idxErr := errors.New("index out of bounds")
	tests := []struct {
		name    string
		origin  ArrayList[int]
		index   int
		newVal  int
		want    ArrayList[int]
		wantErr error
	}{
		{
			name:    "empty case",
			origin:  ArrayList[int]{},
			index:   10,
			newVal:  99,
			want:    ArrayList[int]{},
			wantErr: idxErr,
		}, {
			name:    "exist case",
			origin:  ArrayList[int]{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9},
			index:   4,
			newVal:  99,
			want:    ArrayList[int]{[]int{1, 2, 3, 4, 99, 5, 6, 7, 8, 9}, 10},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantErr := tt.origin.Insert(tt.index, tt.newVal)
			assert.Equal(t, tt.want, tt.origin)
			assert.Equal(t, tt.wantErr, wantErr)
		})
	}
}

func TestArrayList_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		origin ArrayList[int]
		want   bool
	}{
		{
			name:   "empty case",
			origin: ArrayList[int]{},
			want:   true,
		}, {
			name:   "exist case",
			origin: ArrayList[int]{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.origin.IsEmpty()
			assert.Equal(t, tt.want, want)
		})
	}
}

func TestArrayList_Set(t *testing.T) {
	idxErr := errors.New("index out of bounds")
	tests := []struct {
		name    string
		origin  ArrayList[int]
		index   int
		newVal  int
		want    ArrayList[int]
		wantErr error
	}{
		{
			name:    "empty case",
			origin:  ArrayList[int]{},
			index:   10,
			newVal:  99,
			want:    ArrayList[int]{},
			wantErr: idxErr,
		}, {
			name:    "exist case",
			origin:  ArrayList[int]{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9},
			index:   4,
			newVal:  99,
			want:    ArrayList[int]{[]int{1, 2, 3, 4, 99, 6, 7, 8, 9}, 9},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wnatErr := tt.origin.Set(tt.index, tt.newVal)
			assert.Equal(t, tt.want, tt.origin)
			assert.Equal(t, tt.wantErr, wnatErr)
		})
	}
}

func TestArrayList_Size(t *testing.T) {
	tests := []struct {
		name   string
		origin ArrayList[int]
		want   int
	}{
		{
			name:   "empty case",
			origin: ArrayList[int]{},
			want:   0,
		}, {
			name:   "exist case",
			origin: ArrayList[int]{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9},
			want:   9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.origin.Size()
			assert.Equal(t, tt.want, want)
		})
	}
}
