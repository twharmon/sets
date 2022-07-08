package sets_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/twharmon/sets"
)

func assertEqual(t *testing.T, want, got interface{}) {
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestLen(t *testing.T) {
	s := sets.New(1, 2)
	want := 2
	got := s.Len()
	assertEqual(t, want, got)
}

func TestClear(t *testing.T) {
	s := sets.New(1, 2)
	want := 0
	s.Clear()
	got := s.Len()
	assertEqual(t, want, got)
}

func TestAdd(t *testing.T) {
	s := sets.New(1, 2)
	want := []int{1, 2, 3}
	s.Add(3)
	got := s.Slice()
	sort.Ints(got)
	assertEqual(t, want, got)
}

func TestContains(t *testing.T) {
	s := sets.New(1, 2)
	t.Run("true", func(t *testing.T) {
		want := true
		got := s.Contains(2)
		assertEqual(t, want, got)
	})
	t.Run("false", func(t *testing.T) {
		want := false
		got := s.Contains(3)
		assertEqual(t, want, got)
	})
}

func TestRemove(t *testing.T) {
	s := sets.New(1, 2)
	want := []int{2}
	s.Remove(1)
	got := s.Slice()
	sort.Ints(got)
	assertEqual(t, want, got)
}

func TestMerge(t *testing.T) {
	a := sets.New(1, 2)
	b := sets.New(3, 2)
	want := []int{1, 2, 3}
	a.Merge(b)
	got := a.Slice()
	sort.Ints(got)
	assertEqual(t, want, got)
}

func TestUnion(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		a := sets.New(1, 2)
		b := sets.New(3, 2)
		want := []int{1, 2, 3}
		c := sets.Union(a, b)
		got := c.Slice()
		sort.Ints(got)
		assertEqual(t, want, got)
	})
	t.Run("empty", func(t *testing.T) {
		want := []int{}
		c := sets.Union[int]()
		got := c.Slice()
		assertEqual(t, want, got)
	})
}

func TestIntersection(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		a := sets.New(1, 2)
		b := sets.New(3, 2)
		want := []int{2}
		c := sets.Intersection(a, b)
		got := c.Slice()
		assertEqual(t, want, got)
	})
	t.Run("empty", func(t *testing.T) {
		want := []int{}
		c := sets.Intersection[int]()
		got := c.Slice()
		assertEqual(t, want, got)
	})
}
