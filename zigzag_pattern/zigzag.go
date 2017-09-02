package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	/*	0	1	2	3	4	5	6
		0	*			*			*
		1	*		*	*		*	*
		2	*	*		*	*		*
		3	*			*			*

			p		a		h		n
			a	p	l	s	i	i	g
			y		i		r
	*/
	if numRows < 2 || len(s) < 3 || len(s) <= numRows {
		return s
	}
	ret := make([]byte, len(s))
	space := 2*numRows - 2
	index := 0
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += space {
			ret[index] = s[j]
			index++
			if i != 0 && i != numRows-1 && space+j-2*i < len(s) {
				ret[index] = s[space+j-2*i]
				index++
			}
		}
	}
	return string(ret)
}

func main() {
	s := "ABC"
	fmt.Println(convert(s, 2))
}
