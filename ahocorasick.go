package ahocorasick

import (
	"container/list"
	"sync"
)

type node struct {
	child      map[rune]*node
	output     []string
	fail       *node
	isTerminal bool
}

func newNode() *node {
	return &node{
		child: make(map[rune]*node),
	}
}

type Matcher struct {
	root   *node
	locker sync.RWMutex
}

func NewMatcher() *Matcher {
	root := newNode()
	root.fail = root
	return &Matcher{
		root:   root,
		locker: sync.RWMutex{},
	}
}

/*
The failure function is construccted from the goto function. We shall compute
the failure for all states of depth 1,then for all states of depth 2, and so on,
until the failure function has been computed for all states. The failure
function for states of depth d is computed from the failure function for the
states of depth less than d. The states of depth d can be determined from the
nonfail values of the goto function of the states of depth d-1. We consider each
state r of depth d - 1 and perform below actions.

1.If goto(r, a) == fail for all a, do nothing.
2.Otherwise, for each symbol a such that goto(r, a) = s, do the following:
    (a) Set state = f(r)
    (b) Execute the statement state = fail(state) zero or more times, until a
    value for state is obtained such that goto(state, a) != fail. (Note that
    since goto(O, a) != fail for all a, such a state will always be found.)
    (c) Set fail(s) = goto(state, a).
*/
func (m *Matcher) buildFail() {
	l := list.New()
	for _, c := range m.root.child {
		l.PushBack(c)
		c.fail = m.root
	}
	for l.Len() > 0 {
		r := l.Remove(l.Front()).(*node)
		for b, c := range r.child {
			l.PushBack(c)
			fnode := r.fail
			for fnode.child[b] == nil && fnode != m.root {
				fnode = fnode.fail
			}
			var exist bool
			if c.fail, exist = fnode.child[b]; !exist {
				c.fail = m.root
			}
			c.output = append(c.output, c.fail.output...)
		}
	}
}

/*
Insert a entry to trie tree, and build fail functions.
*/
func (m *Matcher) Insert(entry string) {
	m.locker.Lock()
	defer m.locker.Unlock()

	m.add(entry)
	m.buildFail()
}

func (m *Matcher) add(entry string) {
	cur := m.root
	for _, runeValue := range entry {
		if _, exist := cur.child[runeValue]; !exist {
			cur.child[runeValue] = newNode()
		}
		cur = cur.child[runeValue]
	}
	cur.output = []string{entry}
	cur.isTerminal = true
}
func (m *Matcher) delete(entry string) {
	l := list.New()
	cur := m.root
	for _, runeValue := range entry {
		if _, exist := cur.child[runeValue]; !exist {
			return
		}
		if cur.isTerminal {
			l = nil
		} else {
			l.PushBack(cur)
		}
		cur = cur.child[runeValue]
	}
}
func (m *Matcher) Delete(entry string) {
	m.locker.Lock()
	defer m.locker.Unlock()
}

/*
Build a trie tree first, and then build fail functions.
*/
func (m *Matcher) Build(dictionary []string) {
	m.locker.Lock()
	defer m.locker.Unlock()
	for _, entry := range dictionary {
		m.add(entry)
	}
	// Now let's build fail function.
	m.buildFail()
}

func (m *Matcher) Search(str string) (wordIndex map[string][]int) {
	m.locker.RLock()
	defer m.locker.RUnlock()
	wordIndex = make(map[string][]int)
	state := m.root
	runeSlice := []rune(str)
	for k, v := range runeSlice {
		for state.child[v] == nil {
			//state 0(root) accept any rune(i.e not fail.)
			if state == m.root {
				break
			}
			state = state.fail
		}
		state = state.child[v]
		if state == nil {
			state = m.root
		}

		if state.output != nil {
			for _, v := range state.output {
				wordIndex[v] = append(wordIndex[v], k)
			}
		}
	}
	return
}
