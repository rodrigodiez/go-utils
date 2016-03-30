package slice

import "testing"

type IndexText struct {
  Haystack  []string
  Index     int
}

func TestContainsStringReturnsTrueIfStringInSlice(t *testing.T) {
  haystacks := [][]string{
    {"Foo", "Bar"},
    {"Bar", "Foo"},
    {"Foo", "Foo"},
    {"Foo"},
  }

  for _, haystack := range(haystacks) {
    if v := ContainsString("Foo", &haystack); v != true {
      t.Error("Expected", true, "got", v)
    }
  }
}

func TestContainsStringReturnsFalseIfStringNotInSlice(t *testing.T) {
  haystacks := [][]string{
    {"Bar", "Baz"},
    {"Bar"},
    {},
  }

  for _, haystack := range(haystacks) {
    if v := ContainsString("Foo", &haystack); v != false {
      t.Error("Expected", false, "got", v)
    }
  }
}

func TestIndexOfStringReturnsFirstIndexWhereStringIsFound(t *testing.T) {
  tests := []IndexText{
    {[]string{"Foo", "Bar"}, 0},
    {[]string{"Bar", "Foo"}, 1},
    {[]string{"Foo", "Foo"}, 0},
    {[]string{"Foo"}, 0},
  }

  for _, test := range(tests) {
    if v := IndexOfString("Foo", &test.Haystack); v != test.Index {
      t.Error("Expected", test.Index, "got", v)
    }
  }
}

func TestIndexOfStringReturnsMinusOneWhereStringIsNotFound(t *testing.T) {
  tests := []IndexText{
    {[]string{"Bar", "Baz"}, -1},
    {[]string{"Bar"}, -1},
    {[]string{}, -1},
  }

  for _, test := range(tests) {
    if v := IndexOfString("Foo", &test.Haystack); v != test.Index {
      t.Error("Expected", test.Index, "got", v)
    }
  }
}
