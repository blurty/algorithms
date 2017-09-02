## ZigZag Conversion

The string `"PAYPALISHIRING"` is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)

    P   A   H   N
    A P L S I I G
    Y   I   R

And then read line by line: `"PAHNAPLSIIGYIR"`

Write the code that will take a string and make this conversion given a number of rows:

    string convert(string text, int nRows);

`convert("PAYPALISHIRING", 3) should return `"PAHNAPLSIIGYIR"`.


## Approche

首先要搞清楚ZigZag Pattern是什么东西，用图形说话：

    四阶ZigZag
	*			*			*
	*		*	*		*	*
	*	*		*	*		*
	*			*			*

    三阶ZigZag
	p		a		h		n
	a	p	l	s	i	i	g
	y		i		r

可以看出，ZigZag其实就是个倒之字型的图形，可以研究一下这个图形的规律，从中找到解此题的算法，可以得出：

1 之字形行数为numRows,那么每次重复样出现前的间隔字符为space = numRows*2-2

2 第一行和最尾一行均存放一个字符，所以存储的字符是间隔为space的字符

3 中间行需要额外存储一个字符的，存储的字符位置是： space + j - 2*i（其中i为行数，j为该行第几个字符了）