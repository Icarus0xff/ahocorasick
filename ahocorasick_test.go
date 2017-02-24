package ahocorasick

import (
	"reflect"
	"strings"
	"testing"
)

func assert(t *testing.T, b bool) {
	if !b {
		t.Fail()
	}
}

func TestEmpty(t *testing.T) {
	m := NewMatcher()
	m.Build([]string{})
	hits := m.Search("The various modes of worship, which prevailed in the Roman world, were all considered by the people, as equally true; by the Philosopher, as equally false; and by the magistrate, as equally useful.")
	assert(t, len(hits) == 0)
}

func TestSearchNil(t *testing.T) {
	m := NewMatcher()
	m.Build([]string{"foo", "bar", "baz"})
	hits := m.Search("")
	assert(t, len(hits) == 0)
}

func TestLongEntry(t *testing.T) {
	m := NewMatcher()
	m.Insert("The various modes of worship, which prevailed in the Roman world, were all considered by the people, as equally true; by the Philosopher, as equally false; and by the magistrate, as equally useful.")
	hits := m.Search("The various modes of worship, which prevailed in the Roman world, were all considered by the people, as equally true; by the Philosopher, as equally false; and by the magistrate, as equally useful.")
	assert(t, len(hits) == 1)
	assert(t, reflect.DeepEqual(hits["The various modes of worship, which prevailed in the Roman world, were all considered by the people, as equally true; by the Philosopher, as equally false; and by the magistrate, as equally useful."], []int{196}))

}

func TestBuild(t *testing.T) {
	m := NewMatcher()
	dictionary := []string{"Philosopher", "Philosophe", "Philosoph", "Philosop",
		"Philoso", "Philos", "Philo", "Phil", "Phi", "Ph", "P"}
	m.Build(dictionary)
	hits := m.Search("The various modes of worship, which prevailed in the Roman world, were all considered by the people, as equally true; by the Philosopher, as equally false; and by the magistrate, as equally useful.")
	assert(t, len(hits) == len(dictionary))
	assert(t, reflect.DeepEqual(hits["Philosopher"], []int{135}))
	assert(t, reflect.DeepEqual(hits["Philosophe"], []int{134}))
	assert(t, reflect.DeepEqual(hits["Philosoph"], []int{133}))
	assert(t, reflect.DeepEqual(hits["Philosop"], []int{132}))
	assert(t, reflect.DeepEqual(hits["Philoso"], []int{131}))
	assert(t, reflect.DeepEqual(hits["Philos"], []int{130}))
	assert(t, reflect.DeepEqual(hits["Philo"], []int{129}))
	assert(t, reflect.DeepEqual(hits["Phil"], []int{128}))
	assert(t, reflect.DeepEqual(hits["Phi"], []int{127}))
	assert(t, reflect.DeepEqual(hits["Ph"], []int{126}))
	assert(t, reflect.DeepEqual(hits["P"], []int{125}))
}

func TestBuildChinese(t *testing.T) {
	m := NewMatcher()
	dictionary := []string{"罗马帝国寰宇之内的各式各样的宗教信仰和膜拜", "罗马帝国寰宇之内的各式各样的宗教信仰和膜", "罗马帝国寰宇之内的各式各样的宗教信仰和", "罗马帝国寰宇之内的各式各样的宗教信仰", "罗马帝国寰宇之内的各式各样的宗教信", "罗马帝国寰宇之内的各式各样的宗教", "罗马帝国寰宇之内的各式各样的宗", "罗马帝国寰宇之内的各式各样的", "罗马帝国寰宇之内的各式各样", "罗马帝国寰宇之内的各式各", "罗马帝国寰宇之内的各式", "罗马帝国寰宇之内的各", "罗马帝国寰宇之内的", "罗马帝国寰宇之内", "罗马帝国寰宇之", "罗马帝国寰宇", "罗马帝国寰", "罗马帝国", "罗马帝", "罗马", "罗"}
	m.Build(dictionary)
	hits := m.Search("流行于罗马帝国寰宇之内的各式各样的宗教信仰和膜拜，一般人民看来都是同样灵验；明哲之士看来，同样荒诞；统治阶级看来，同样有用。")
	assert(t, len(hits) == len(dictionary))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之内的各式各样的宗教信仰和膜拜"], []int{23}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之内的各式各样的宗教信仰和膜"], []int{22}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之内的各式各样的宗教信仰和"], []int{21}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之内的各式各样的宗教信仰"], []int{20}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之内的各式各样的宗教信"], []int{19}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之内的各式各样的宗教"], []int{18}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之内的各式各样的宗"], []int{17}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之内的各式各样的"], []int{16}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之内的各式各样"], []int{15}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之内的各式各"], []int{14}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之内的各式"], []int{13}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之内的各"], []int{12}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之内的"], []int{11}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之内"], []int{10}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇之"], []int{9}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰宇"], []int{8}))
	assert(t, reflect.DeepEqual(hits["罗马帝国寰"], []int{7}))
	assert(t, reflect.DeepEqual(hits["罗马帝国"], []int{6}))
	assert(t, reflect.DeepEqual(hits["罗马帝"], []int{5}))
	assert(t, reflect.DeepEqual(hits["罗马"], []int{4}))
	assert(t, reflect.DeepEqual(hits["罗"], []int{3}))
}

func TestInsert(t *testing.T) {
	m := NewMatcher()
	m.Insert("Philosopher")
	m.Insert("Philosophe")
	m.Insert("Philosoph")
	m.Insert("Philosop")
	m.Insert("Philoso")
	m.Insert("Philos")
	m.Insert("Philo")
	m.Insert("Phil")
	m.Insert("Phi")
	m.Insert("Ph")
	m.Insert("P")
	hits := m.Search("The various modes of worship, which prevailed in the Roman world, were all considered by the people, as equally true; by the Philosopher, as equally false; and by the magistrate, as equally useful.")
	assert(t, reflect.DeepEqual(hits["Philosopher"], []int{135}))
	assert(t, reflect.DeepEqual(hits["Philosophe"], []int{134}))
	assert(t, reflect.DeepEqual(hits["Philosoph"], []int{133}))
	assert(t, reflect.DeepEqual(hits["Philosop"], []int{132}))
	assert(t, reflect.DeepEqual(hits["Philoso"], []int{131}))
	assert(t, reflect.DeepEqual(hits["Philos"], []int{130}))
	assert(t, reflect.DeepEqual(hits["Philo"], []int{129}))
	assert(t, reflect.DeepEqual(hits["Phil"], []int{128}))
	assert(t, reflect.DeepEqual(hits["Phi"], []int{127}))
	assert(t, reflect.DeepEqual(hits["Ph"], []int{126}))
	assert(t, reflect.DeepEqual(hits["P"], []int{125}))

}

// TestInsertAfterBuild ...
func TestInsertAfterBuild(t *testing.T) {
	m := NewMatcher()
	dictionary := []string{"Philosopher", "Philosophe", "Philosoph", "Philosop"}
	m.Build(dictionary)
	m.Insert("Philoso")
	m.Insert("Philos")
	m.Insert("Philo")
	m.Insert("Phil")
	m.Insert("Phi")
	m.Insert("Ph")
	m.Insert("P")
	hits := m.Search("The various modes of worship, which prevailed in the Roman world, were all considered by the people, as equally true; by the Philosopher, as equally false; and by the magistrate, as equally useful.")
	assert(t, len(hits) == 11)
	assert(t, reflect.DeepEqual(hits["Philosopher"], []int{135}))
	assert(t, reflect.DeepEqual(hits["Philosophe"], []int{134}))
	assert(t, reflect.DeepEqual(hits["Philosoph"], []int{133}))
	assert(t, reflect.DeepEqual(hits["Philosop"], []int{132}))
	assert(t, reflect.DeepEqual(hits["Philoso"], []int{131}))
	assert(t, reflect.DeepEqual(hits["Philos"], []int{130}))
	assert(t, reflect.DeepEqual(hits["Philo"], []int{129}))
	assert(t, reflect.DeepEqual(hits["Phil"], []int{128}))
	assert(t, reflect.DeepEqual(hits["Phi"], []int{127}))
	assert(t, reflect.DeepEqual(hits["Ph"], []int{126}))
	assert(t, reflect.DeepEqual(hits["P"], []int{125}))
}

func benchmarkSingleString(b *testing.B, pattern []string, text string) {
	m := NewMatcher()
	m.Build(pattern)
	b.SetBytes(int64(len(text)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Search(text)
	}
}

func BenchmarkChineseMatch(b *testing.B) {
	dictionary := []string{"罗马帝国寰宇之内的各式各样的宗教信仰和膜拜", "罗马帝国寰宇之内的各式各样的宗教信仰和膜", "罗马帝国寰宇之内的各式各样的宗教信仰和", "罗马帝国寰宇之内的各式各样的宗教信仰", "罗马帝国寰宇之内的各式各样的宗教信", "罗马帝国寰宇之内的各式各样的宗教", "罗马帝国寰宇之内的各式各样的宗", "罗马帝国寰宇之内的各式各样的", "罗马帝国寰宇之内的各式各样", "罗马帝国寰宇之内的各式各", "罗马帝国寰宇之内的各式", "罗马帝国寰宇之内的各", "罗马帝国寰宇之内的", "罗马帝国寰宇之内", "罗马帝国寰宇之", "罗马帝国寰宇", "罗马帝国寰", "罗马帝国", "罗马帝", "罗马", "罗"}
	benchmarkSingleString(b, dictionary, strings.Repeat("流行于罗马帝国寰宇之内的各式各样的宗教信仰和膜拜，一般人民看来都是同样灵验；明哲之士看来，同样荒诞；统治阶级看来，同样有用。", 100))
}
