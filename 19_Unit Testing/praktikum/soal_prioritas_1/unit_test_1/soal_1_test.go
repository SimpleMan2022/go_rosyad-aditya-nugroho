package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRectangleArea(t *testing.T) {
	t.Run("even", func(t *testing.T) {
		result := RectangleArea(4, 2)
		assert.Equal(t, "even rectangle", result, "Result must be even rectangle")
		fmt.Println("Test RectangleArea Done")
	})
	t.Run("odd", func(t *testing.T) {
		result := RectangleArea(7, 3)
		assert.Equal(t, "odd rectangle", result, "Result must be odd rectangle")
		fmt.Println("Test RectangleArea Done")
	})
}

func TestRectanglePerimeter(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		result := RectanglePerimeter(4, 6)
		assert.Equal(t, true, result, "Result must be True")
	})
	t.Run("false", func(t *testing.T) {
		result := RectanglePerimeter(2, 7)
		assert.Equal(t, false, result, "Result must be False")
	})
	fmt.Println("Test Rectangle Perimeter Done")
}

func TestSquareArea(t *testing.T) {
	t.Run("even", func(t *testing.T) {
		result := SquareArea(6)
		assert.Equal(t, "even square", result, "Result must be even square")
	})
	t.Run("odd", func(t *testing.T) {
		result := SquareArea(3)
		assert.Equal(t, "odd square", result, "Result must be odd square")
	})
	fmt.Println("Test Square Area Done")
}

func TestSquarePerimeter(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		result := SquarePerimeter(10)
		assert.Equal(t, true, result, "Result must be true")
	})
	t.Run("odd", func(t *testing.T) {
		result := SquarePerimeter(8)
		assert.Equal(t, false, result, "Result must be false")
	})
	fmt.Println("Test Square Perimeter Done")
}
