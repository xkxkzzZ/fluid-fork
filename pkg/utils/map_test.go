/*
Copyright 2023 The Fluid Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestContainsAll(t *testing.T) {
	type args struct {
		m     map[string]string
		slice []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				m: map[string]string{
					"foo": "1",
				},
				slice: nil,
			},
			want: true,
		},
		{
			name: "nil map",
			args: args{
				m:     nil,
				slice: []string{"foo", "bar"},
			},
			want: false,
		},
		{
			name: "contains all",
			args: args{
				m: map[string]string{
					"foo": "1",
					"bar": "2",
					"xxx": "3",
				},
				slice: []string{"foo", "xxx"},
			},
			want: true,
		},
		{
			name: "contains some",
			args: args{
				m: map[string]string{
					"foo": "1",
					"bar": "2",
					"xxx": "3",
				},
				slice: []string{"foo", "yyy"},
			},
			want: false,
		},
		{
			name: "contains none",
			args: args{
				m: map[string]string{
					"foo": "1",
					"bar": "2",
					"xxx": "3",
				},
				slice: []string{"aaa", "bbb"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAll(tt.args.m, tt.args.slice); got != tt.want {
				t.Errorf("ContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionMapsWithOverride(t *testing.T) {
	type args struct {
		map1 map[string]string
		map2 map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "no_shared_elements",
			args: args{
				map1: map[string]string{"key1": "value1", "key2": "value2"},
				map2: map[string]string{"keyA": "valueA", "keyB": "valueB"},
			},
			want: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"keyA": "valueA",
				"keyB": "valueB",
			},
		},
		{
			name: "with_shared_elements",
			args: args{
				map1: map[string]string{"key1": "value1", "key2": "value2"},
				map2: map[string]string{"key2": "new_value", "key3": "value3"},
			},
			want: map[string]string{
				"key1": "value1",
				"key2": "new_value",
				"key3": "value3",
			},
		},
		{
			name: "nil_map",
			args: args{
				map1: map[string]string{"key1": "value1", "key2": "value2"},
				map2: nil,
			},
			want: map[string]string{"key1": "value1", "key2": "value2"},
		},
		{
			name: "nil_maps",
			args: args{
				map1: nil,
				map2: nil,
			},
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnionMapsWithOverride(tt.args.map1, tt.args.map2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionMapsWithOverride() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectIntegerSets(t *testing.T) {
	type args struct {
		map1 map[int]bool
		map2 map[int]bool
	}
	tests := []struct {
		name string
		args args
		want map[int]bool
	}{
		{
			name: "empty_map1",
			args: args{
				map1: map[int]bool{},
				map2: map[int]bool{100: true, 101: true},
			},
			want: map[int]bool{},
		},
		{
			name: "empty_map2",
			args: args{
				map1: map[int]bool{100: true, 101: true},
				map2: map[int]bool{},
			},
			want: map[int]bool{},
		},
		{
			name: "empty_intersection",
			args: args{
				map1: map[int]bool{100: true, 101: true},
				map2: map[int]bool{102: true},
			},
			want: map[int]bool{},
		},
		{
			name: "basic_intersection",
			args: args{
				map1: map[int]bool{100: true, 42: true, 101: true},
				map2: map[int]bool{100: true, 101: true, 102: true},
			},
			want: map[int]bool{100: true, 101: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectIntegerSets(tt.args.map1, tt.args.map2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectIntegerSets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedKeys(t *testing.T) {
	type args struct {
		m map[string]struct{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty_map",
			args: args{
				m: map[string]struct{}{},
			},
			want: []string{},
		},
		{
			name: "map_with_elements",
			args: args{
				m: map[string]struct{}{
					"bbb":  {},
					"aaa":  {},
					"aa":   {},
					"bbbb": {},
				},
			},
			want: []string{"aa", "aaa", "bbb", "bbbb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderedKeys(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

var _ = Describe("KeyMatched", func() {
	Context("when the map is empty", func() {
		It("should return false for any key", func() {
			m := map[string]int{}
			Expect(KeyMatched(m, "key1")).To(BeFalse())
		})
	})

	Context("when the map contains the key", func() {
		It("should return true", func() {
			m := map[string]int{"key1": 1, "key2": 2}
			Expect(KeyMatched(m, "key1")).To(BeTrue())
		})
	})

	Context("when the map does not contain the key", func() {
		It("should return false", func() {
			m := map[string]int{"key1": 1, "key2": 2}
			Expect(KeyMatched(m, "key3")).To(BeFalse())
		})
	})

	Context("when the map is nil", func() {
		It("should return false for any key", func() {
			var m map[string]int
			Expect(KeyMatched(m, "key1")).To(BeFalse())
		})
	})

	Context("when the query key is empty", func() {
		It("should return false", func() {
			m := map[string]string{"key1": "value1"}
			Expect(KeyMatched(m, "")).To(BeFalse())
		})
	})
})

var _ = Describe("KeyValueMatched", func() {
	Context("when the map is empty", func() {
		It("should return false for any key-value pair", func() {
			m := map[string]int{}
			Expect(KeyValueMatched(m, "key1", 1)).To(BeFalse())
		})
	})

	Context("when the map contains the key with the matching value", func() {
		It("should return true", func() {
			m := map[string]int{"key1": 1, "key2": 2}
			Expect(KeyValueMatched(m, "key1", 1)).To(BeTrue())
		})
	})

	Context("when the map contains the key with a different value", func() {
		It("should return false", func() {
			m := map[string]int{"key1": 1, "key2": 2}
			Expect(KeyValueMatched(m, "key1", 2)).To(BeFalse())
		})
	})

	Context("when the map does not contain the key", func() {
		It("should return false", func() {
			m := map[string]int{"key1": 1, "key2": 2}
			Expect(KeyValueMatched(m, "key3", 3)).To(BeFalse())
		})
	})

	Context("when the map is nil", func() {
		It("should return false for any key-value pair", func() {
			var m map[string]int
			Expect(KeyValueMatched(m, "key1", 1)).To(BeFalse())
		})
	})
})
