package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddItemsToSet(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3)
	assert.Equal(t, 3, s.Length())
	s.Add(1, 2, 3)
	assert.Equal(t, 3, s.Length()) // 添加重复元素，集合长度不变
}

func TestCheckItemInSet(t *testing.T) {
	s := NewSet[int]()
	assert.False(t, s.Contains(1)) // 集合为空，不包含元素
	s.Add(1, 2, 3)
	assert.True(t, s.Contains(1))
	assert.False(t, s.Contains(4))
}

func TestConvertSetToList(t *testing.T) {
	s := NewSet[int]()
	list := s.ToList()
	assert.Equal(t, 0, len(list)) // 空集合转换为列表，长度为0
	s.Add(1, 2, 3)
	list = s.ToList()
	assert.ElementsMatch(t, []int{1, 2, 3}, list)
}

func TestCheckSetLength(t *testing.T) {
	s := NewSet[int]()
	assert.Equal(t, 0, s.Length()) // 新建的集合长度为0
	s.Add(1, 2, 3)
	assert.Equal(t, 3, s.Length())
}

func TestRemoveItemFromSet(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3)
	s.Remove(2)
	assert.False(t, s.Contains(2))
	s.Remove(4)
	assert.Equal(t, 2, s.Length()) // 删除不存在的元素，集合长度不变
}

func TestClearSet(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3)
	s.Clear()
	assert.Equal(t, 0, s.Length())
	assert.False(t, s.Contains(1)) // 清空集合后，不包含任何元素
}

func TestUnionOfTwoSets(t *testing.T) {
	s1 := NewSet[int]()
	s1.Add(1, 2, 3)

	s2 := NewSet[int]()
	s2.Add(3, 4, 5)
	union := s1.Union(s2)
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, union.ToList())

	s3 := NewSet[int]()
	union = s1.Union(s3)
	assert.ElementsMatch(t, []int{1, 2, 3}, union.ToList()) // 与空集合求并集，结果不变
}

func TestIntersectionOfTwoSets(t *testing.T) {
	s1 := NewSet[int]()
	s1.Add(1, 2, 3)

	s2 := NewSet[int]()
	s2.Add(3, 4, 5)
	intersection := s1.Intersection(s2)
	assert.ElementsMatch(t, []int{3}, intersection.ToList())

	s3 := NewSet[int]()
	intersection = s1.Intersection(s3)
	assert.Equal(t, 0, intersection.Length()) // 与空集合求交集，结果为空集
}

func TestDifferenceOfTwoSets(t *testing.T) {
	s1 := NewSet[int]()
	s1.Add(1, 2, 3)

	s2 := NewSet[int]()
	s2.Add(3, 4, 5)
	difference := s1.Difference(s2)
	assert.ElementsMatch(t, []int{1, 2}, difference.ToList())

	s3 := NewSet[int]()
	difference = s1.Difference(s3)
	assert.ElementsMatch(t, []int{1, 2, 3}, difference.ToList()) // 与空集合求差集，结果不变
}
