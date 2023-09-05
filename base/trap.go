package trap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//示例 2：
//
//输入：height = [4,2,0,3,2,5]
//输出：9

func TrapV1(height []int) int {
	sum := 0

	for i := 1; i < len(height)-1; i++ {
		maxLeft := 0
		for j := i - 1; j >= 0; j-- {
			if maxLeft < height[j] {
				maxLeft = height[j]
			}
		}
		maxRight := 0
		for j := i + 1; j < len(height); j++ {
			if maxRight < height[j] {
				maxRight = height[j]
			}
		}

		cur := 0
		if maxLeft > maxRight {
			cur = (maxRight - height[i])
		} else {
			cur = (maxLeft - height[i])
		}
		if cur > 0 {
			sum += cur
		}
	}

	return sum
}

func TrapV2(height []int) int {
	sum := 0
	maxLeft := make([]int, len(height))
	maxLeft[0] = height[0]
	for i := 1; i < len(height); i++ {
		if maxLeft[i-1] < height[i] {
			maxLeft[i] = height[i]
		} else {
			maxLeft[i] = maxLeft[i-1]
		}
	}
	maxRight := make([]int, len(height))
	maxRight[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		if maxRight[i+1] < height[i] {
			maxRight[i] = height[i]
		} else {
			maxRight[i] = maxRight[i+1]
		}
	}

	for i := 1; i < len(height)-1; i++ {
		cur := 0
		if maxLeft[i] > maxRight[i] {
			cur = (maxRight[i] - height[i])
		} else {
			cur = (maxLeft[i] - height[i])
		}
		if cur > 0 {
			sum += cur
		}
	}

	return sum
}

func TestTrap(t *testing.T) {
	assert.Equal(t, 6, TrapV1([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), "should be equal")
	assert.Equal(t, 9, TrapV1([]int{4, 2, 0, 3, 2, 5}), "should be equal")
	assert.Equal(t, 6, TrapV2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), "should be equal")
	assert.Equal(t, 9, TrapV2([]int{4, 2, 0, 3, 2, 5}), "should be equal")
}
