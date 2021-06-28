// Package itertools provides a translation of the python standard library module itertools.
// Many of the functions have been brought over, althought not all.
// In this implementation, chan interface{} has been used as all iterators; if more specific types are necessary,
// feel free to copy the code to your project to be implemented with more specific types.
package itertools

import (
	"sort"
	"strings"
	"sync"
)

type Iter chan interface{}
type Predicate func (interface{}) bool
type Mapper func (interface{}) interface{}
type MultiMapper func (...interface{}) interface{}
type Reducer func (memo interface{}, element interface{}) interface{}
func Permutations(r int, els ... interface{}) (c[]string) {

	//# permutations('ABCD', 2) --> AB AC AD BA BC BD CA CB CD DA DB DC
	//# permutations(range(3)) --> 012 021 102 120 201 210
	pool := make([]string, 0)
	n := len(els)
	if r >= n {
		return nil
	}
	// 返回结果的总个数
	m := SectionMutiplication(n, n-r+1)

	indices := make([]int, n)
	indicesReverse := make([]int, r)
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	for i := 0; i < r; i++ {
		indicesReverse[i] = i
	}

	for _, v := range els {
		pool = append(pool, v.(string))
	}

	// 获取第一个
	first := pool[indices[0]: r]
	firstItem := strings.Join(first, ",")
	c = append(c, firstItem)
	sort.Sort(sort.Reverse(sort.IntSlice(indicesReverse)))

	var cycles []int

	for i := n; i > n - r; i-- {
		cycles = append(cycles, i)
	}
	lc := int64(len(c))
	for m > lc {
		for _, ii := range indicesReverse {
			cycles[ii] -= 1
			if cycles[ii] == 0 {
				//indices[i:] = indices[i+1:] + indices[i:i+1]
				tmp := append(indices[ii+1:], indices[ii: ii+1]...)
				indices = append(indices[0: len(indices)-len(tmp)], tmp...)
				//if ii == 0 {
				//	indices = tmp
				//} else {
				//	indices = append(indices[0: ii], tmp...)
				//}

				cycles[ii] = n - ii
			} else {
				j := cycles[ii]
				j = len(indices)-j

				iiv := indices[ii]
				jv := indices[j]
				indices[ii] = jv
				indices[j] = iiv
				var it []string
				for _,v := range indices[0:r] {
					it = append(it, pool[v])
				}
				c = append(c, strings.Join(it, ","))
				lc = int64(len(c))
				break
			}
		}
	}
	return c
}

func factorial(n3 int) int64 {
	n := int64(n3)
	if n < 0 {
		return 0
	}
	facVal := int64(1)
	i := int64(1)
	for i <= n {
		facVal *= i
		i++
	}
	return facVal
}

//SectionMutiplication 求区间内的乘积
func SectionMutiplication(max int, min int) (res int64) {
	res = int64(max)
	for max > min {
		max = max - 1
		res *= int64(max)
	}
	return res
}
func Combinations(r int, els[]string) (c []string){
 	println(len(els))

	pool := make([]string, 0)
	//# combinations('ABCD', 2) --> AB AC AD BC BD CD
	//# combinations(range(4), 3) --> 012 013 023 123
	n := len(els)
	if r > n {
		return nil
	}
	// 返回结果的总个数
	m := SectionMutiplication(n, r)/factorial(r)
	indices := make([]int, r)
	indicesReverse := make([]int, r)
	for i := 0; i < r; i++ {
		indices[i] = i
		indicesReverse[i] = i
	}
	for _, v := range els {
		pool = append(pool, v)
	}
	// 获取第一个
	first := pool[indices[0]: r]

	firstItem := strings.Join(first, ",")
	c = append(c, firstItem)
	if r == n {
		return c
	}
	sort.Sort(sort.Reverse(sort.IntSlice(indicesReverse)))
	println(m, len(c), "ken")
	lc := int64(len(c))
	for m > lc {
		var i = 0
		for ii := range indicesReverse {
			i = indicesReverse[ii]
			if indices[i] != i + n - r {
				break
			}
		}
		indices[i] += 1
		var tmp []int
		if i + 1 < r {
			tmp = make([]int, r-(i+1))
			var ti = 0
			for tv := i+1; tv < r; tv++ {
				tmp[ti] = tv
				ti++
			}
		}
		for _,j := range tmp {
			indices[j] = indices[j - 1] + 1
		}
		var it []string
		for _,v := range indices {
			it = append(it, pool[v])
		}
		c = append(c, strings.Join(it, ","))
		lc = int64(len(c))
	}

	return c
}
func New(els ... interface{}) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

func Int64(els ... int64) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

func Int32(els ... int32) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

func Float64(els ... float64) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

func Float32(els ... float32) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

func Uint(els ... uint) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}
func Uint64(els ... uint64) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

func Uint32(els ... uint32) Iter {
	c := make(Iter)
	go func () {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

func List(it Iter) []interface{} {
	arr := make([]interface{}, 0, 1)
	for el := range it {
		arr = append(arr, el)
	}
	return arr
}

// Count from i to infinity
func Count(i int) Iter {
	c := make(Iter)
	go func () {
		for ; true; i++ {
			c <- i
		}
	}()
	return c
}

// Cycle through an iterator infinitely (requires memory)
func Cycle(it Iter) Iter {
	c, a := make(Iter), make([]interface{}, 0, 1)
	go func () {
		for el := range it {
			a = append(a, el)
			c <- el
		}
		for {
			for _, el := range a {
				c <- el
			}
		}
	}()
	return c
}

// Repeat an element n times or infinitely
func Repeat(el interface{}, n ...int) Iter {
	c := make(Iter)
	go func () {
		for i := 0; len(n) == 0 || i < n[0]; i++ {
			c <- el
		}
		close(c)
	}()
	return c
}

// Chain together multiple iterators
func Chain(its ...Iter) Iter {
	c := make(Iter)
	go func() {
		for _, it := range its {
			for el := range it {
				c <- el
			}
		}
		close(c)
	}()
	return c
}

// Elements after pred(el) == true
func DropWhile(pred Predicate, it Iter) Iter {
	c := make(Iter)
	go func () {
		for el := range it {
			if drop := pred(el); !drop {
				c <- el
				break
			}
		}
		for el := range it {
			c <- el
		}
		close(c)
	}()
	return c
}


// Elements before pred(el) == false
func TakeWhile(pred Predicate, it Iter) Iter {
	c := make(Iter)
	go func () {
		for el := range it {
			if take := pred(el); take {
				c <- el
			} else {
				break
			}
		}
		close(c)
	}()
	return c
}

// Filter out any elements where pred(el) == false
func Filter(pred Predicate, it Iter) Iter {
	c := make(Iter)
	go func () {
		for el := range it {
			if keep := pred(el); keep {
				c <- el
			}
		}
		close(c)
	}()
	return c
}

// Filter out any elements where pred(el) == true
func FilterFalse(pred Predicate, it Iter) Iter {
	c := make(Iter)
	go func () {
		for el := range it {
			if drop := pred(el); !drop {
				c <- el
			}
		}
		close(c)
	}()
	return c
}

// Sub-iterator from start (inclusive) to [stop (exclusive) every [step (default 1)]]
func Slice(it Iter, startstopstep...int) Iter {
	start, stop, step := 0, 0, 1
	if len(startstopstep) == 1 {
		start = startstopstep[0]
	} else if len(startstopstep) == 2 {
		start, stop = startstopstep[0], startstopstep[1]
	} else	if len(startstopstep) >= 3 {
		start, stop, step = startstopstep[0], startstopstep[1], startstopstep[2]
	}

	c := make(Iter)
	go func () {
		i := 0
		// Start
		for el := range it {
			if i >= start {
				c <- el // inclusive
				break
			}
			i += 1
		}

		// Stop
		i, j := i + 1, 1
		for el := range it {
			if stop > 0 && i >= stop {
				break
			} else if j % step == 0 {
				c <- el
			}

			i, j = i + 1, j + 1
		}

		close(c)
	}()
	return c
}

// Map an iterator to fn(el) for el in it
func Map(fn Mapper, it Iter) Iter {
	c := make(Iter)
	go func () {
		for el := range it {
			c <- fn(el)
		}
		close(c)
	}()
	return c
}

// Map p, q, ... to fn(pEl, qEl, ...)
// Breaks on first closed channel
func MultiMap(fn MultiMapper, its ...Iter) Iter {
	c := make(Iter)
	go func() {
Outer:
		for {
			els := make([]interface{}, len(its))
			for i, it := range its {
				if el, ok := <- it; ok {
					els[i] = el
				} else {
					break Outer
				}
			}
			c <- fn(els...)
		}
		close(c)
	}()
	return c
}

// Map p, q, ... to fn(pEl, qEl, ...)
// Breaks on last closed channel
func MultiMapLongest(fn MultiMapper, its ...Iter) Iter {
	c := make(Iter)
	go func() {
		for {
			els := make([]interface{}, len(its))
			n := 0
			for i, it := range its {
				if el, ok := <- it; ok {
					els[i] = el
				} else {
					n += 1
				}
			}
			if n < len(its) {
				c <- fn(els...)
			} else {
				break
			}
		}
		close(c)
	}()
	return c
}

// Map an iterator if arrays to a fn(els...)
// Iter must be an iterator of []interface{} (possibly created by Zip)
// If not, Starmap will act like MultiMap with a single iterator
func Starmap(fn MultiMapper, it Iter) Iter {
	c := make(Iter)
	go func() {
		for els := range it {
			if elements, ok := els.([]interface{}); ok {
				c <- fn(elements...)
			} else {
				c <- fn(els)
			}
		}
		close(c)
	}()
	return c
}

// Zip up multiple interators into one
// Close on shortest iterator
func Zip(its ...Iter) Iter {
	c := make(Iter)
	go func() {
		defer close(c)
		for {
			els := make([]interface{}, len(its))
			for i, it := range its {
				if el, ok := <- it; ok {
					els[i] = el
				} else {
					return
				}
			}
			c <- els
		}
	}()
	return c
}

// Zip up multiple iterators into one
// Close on longest iterator
func ZipLongest(its ...Iter) Iter {
	c := make(Iter)
	go func() {
		for {
			els := make([]interface{}, len(its))
			n := 0
			for i, it := range its {
				if el, ok := <- it; ok {
					els[i] = el
				} else {
					n += 1
				}
			}
			if n < len(its) {
				c <- els
			} else {
				break
			}
		}
		close(c)
	}()
	return c
}

// Reduce the iterator (aka fold) from the left
func Reduce(it Iter, red Reducer, memo interface{}) interface{} {
	for el := range it {
		memo = red(memo, el)
	}
	return memo
}

// Split an iterator into n multiple iterators
// Requires memory to keep values for n iterators
func Tee(it Iter, n int) []Iter {
	deques := make([][]interface{}, n)
	iters := make([]Iter, n)
	for i := 0; i < n; i++ {
		iters[i] = make(Iter)
	}

	mutex := new(sync.Mutex)

	gen := func(myiter Iter, i int) {
		for {
			if len(deques[i]) == 0 {
				mutex.Lock()
				if len(deques[i]) == 0 {
					if newval, ok := <- it; ok {
						for i, d := range deques {
							deques[i] = append(d, newval)
						}
					} else {
						mutex.Unlock()
						close(myiter)
						break
					}
				}
				mutex.Unlock()
			}
			var popped interface{}
			popped, deques[i] = deques[i][0], deques[i][1:]
			myiter <- popped
		}
	}
	for i, iter := range iters {
		go gen(iter, i)
	}
	return iters
}

// Helper to tee just into two iterators
func Tee2(it Iter) (Iter, Iter) {
	iters := Tee(it, 2)
	return iters[0], iters[1]
}
