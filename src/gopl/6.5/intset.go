package main

import (
	"fmt"
)

type Intset struct {
	words []uint64
}

func (list *Intset) Has(i int) bool {
	block, index := i/64, uint(i%64)
	has := (block < len(list.words) && list.words[block]&(1<<index) != 0)
	return has
}

func (list *Intset) Add(i int) {
	block, index := i/64, uint(i%64)
	for block >= len(list.words) {
		list.words = append(list.words, 0)

	}
	list.words[block] |= (1 << index)
}

func (dlist *Intset) Union(slist *Intset) {
	for block, v := range slist.words {
		if block < len(dlist.words) {
			dlist.words[block] |= v
		} else {
			dlist.words = append(dlist.words, v)
		}
	}
}

func (dlist *Intset) Intersect(slist *Intset) *Intset {
	var temp2 Intset
	for b, v := range slist.words {
		if b > len(dlist.words) {
			break
		}
		for i := 0; i < 64; i++ {
			if 0x1&(v>>uint(i)) != 0 && 0x1&(dlist.words[b]>>uint(i)) != 0 {
				temp2.Add(b*64 + i)
			}
		}

	}
	return &temp2
}

func (dlist *Intset) Difference(slist *Intset) *Intset {
	var temp2 Intset
	temp2.words = make([]uint64, len(dlist.words))
	copy(temp2.words, dlist.words)

	for b, v := range dlist.words {
		if b > len(slist.words) {
			break
		}
		for i := 0; i < 64; i++ {
			if 0x1&(v>>uint(i)) != 0 && 0x1&(slist.words[b]>>uint(i)) != 0 {
				temp2.Remove(b*64 + i)
			}
		}

	}

	return &temp2
}

func (dlist *Intset) SymmetricDifference(slist *Intset) *Intset {
	var temp2 Intset
	for b, v := range dlist.words {
		if b > len(slist.words) {
			break
		}
		for i := 0; i < 64; i++ {
			if 0x1&(v>>uint(i)) != 0 && 0x1&(slist.words[b]>>uint(i)) == 0 {
				temp2.Add(b*64 + i)
			}
			if 0x1&(v>>uint(i)) == 0 && 0x1&(slist.words[b]>>uint(i)) != 0 {
				temp2.Add(b*64 + i)
			}
		}
		copy(temp2.words[len(temp2.words):], slist.words[b:])
		copy(temp2.words[len(temp2.words):], dlist.words[b:])
	}
	return &temp2
}

func (list *Intset) String() string {
	var o = 0
	var tempString = ""
	var outString = ""
	for b, v := range list.words {
		if v == 0 {
			continue
		}
		for i := 0; i < 64; i++ {
			o = 0
			tempString = ""
			if 0x1&(v>>uint(i)) != 0 {
				o = b*64 + i
				tempString = fmt.Sprintf("%d", o)
				if len(outString) > 0 {
					outString += ","
				}
				outString += tempString
			}
		}
	}
	return outString
}

func (list *Intset) Len() int {
	var lens = 0
	for _, v := range list.words {
		if v == 0 {
			continue
		}
		for i := 0; i < 64; i++ {
			if 0x1&(v>>uint(i)) != 0 {
				lens++
			}
		}
	}
	return lens

}

func (list *Intset) Remove(num int) {
	for b, v := range list.words {
		if v == 0 {
			continue
		}
		for i := 0; i < 64; i++ {
			if 0x1&(v>>uint(i)) != 0 {
				if num == b*64+i {
					v &= (^(1 << uint(i)))
					list.words[b] = v
				}
			}
		}
	}
}

func (list *Intset) Clear() {
	list.words = nil
}

func (list *Intset) Copy() *Intset {
	var x Intset
	x = *list
	return &x
}

func main() {
	var x, y Intset
	var z *Intset = &x
	/*
		x.Add(1)
		x.Add(25)
		y.Add(5)
		y.Add(1)
		fmt.Println(x.Has(1))
		fmt.Println(x.String())
		x.Union(&y)
		fmt.Println(x.String())
		fmt.Println(x.Len())
		x.Remove(5)
		fmt.Println(x.String())
		x.Clear()
		x.Add(66)
		x.Add(67)
		fmt.Println(x.String())
		z = x.Copy()
		fmt.Println(z.String())
	*/
	x.Add(1)
	x.Add(5)
	x.Add(25)

	y.Add(5)
	y.Add(25)
	z = x.Intersect(&y)
	fmt.Println(z.String())
	z = x.Difference(&y)
	fmt.Println(x.String())
	fmt.Println(z.String())
	z = x.SymmetricDifference(&y)
	fmt.Println(z.String())
}
