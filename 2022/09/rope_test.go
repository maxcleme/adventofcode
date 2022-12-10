package _09

import (
	"reflect"
	"testing"
)

func TestPosition_Follow(t *testing.T) {
	tests := []struct {
		before Position
		head   Position
		after  Position
	}{
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: 0, y: 0},
			after:  Position{x: 0, y: 0},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: 1, y: 0},
			after:  Position{x: 0, y: 0},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: 0, y: 1},
			after:  Position{x: 0, y: 0},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: -1, y: 0},
			after:  Position{x: 0, y: 0},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: 0, y: -1},
			after:  Position{x: 0, y: 0},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: 1, y: 1},
			after:  Position{x: 0, y: 0},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: -1, y: 1},
			after:  Position{x: 0, y: 0},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: 1, y: -1},
			after:  Position{x: 0, y: 0},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: -1, y: -1},
			after:  Position{x: 0, y: 0},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: 2, y: 0},
			after:  Position{x: 1, y: 0},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: 0, y: 2},
			after:  Position{x: 0, y: 1},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: -2, y: 0},
			after:  Position{x: -1, y: 0},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: 0, y: -2},
			after:  Position{x: 0, y: -1},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: 2, y: 1},
			after:  Position{x: 1, y: 1},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: 1, y: 2},
			after:  Position{x: 1, y: 1},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: -2, y: -1},
			after:  Position{x: -1, y: -1},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: -1, y: -2},
			after:  Position{x: -1, y: -1},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: 2, y: 2},
			after:  Position{x: 1, y: 1},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: -2, y: -2},
			after:  Position{x: -1, y: -1},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: -2, y: 2},
			after:  Position{x: -1, y: 1},
		},
		{
			before: Position{x: 0, y: 0},
			head:   Position{x: -2, y: 2},
			after:  Position{x: -1, y: 1},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.before.Follow(tt.head); !reflect.DeepEqual(got, tt.after) {
				t.Errorf("Follow() = %v, want %v", got, tt.after)
			}
		})
	}
}
