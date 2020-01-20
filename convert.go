package main

/*
LEETCODEISHIRING

0         k*(2*numRows-2)
numRows-1 k*(2*numRows-2)+numRows-1
i         k*(2*numRows-2)+i , (k+1)*(2*numRows-2)-i

row k 4(k+1)-1 4k+1
1   0 1
1   0 3
1   1          5

numRows=3 1
	0 1 2 3 4 5 6 7
0 L   C   I   R
1 E T O E S I I G
2 E   D   H   N

当列数是：0 2 4 6 8，需要填入完整列
1 1
1 3
1 5
1 7
当列数是：1 3 5 7，需要填入在行1上填入

numRows=4 2
	0 1 2 3 4 5 6
0 L     D     R
1 E   O E   I I
2 E C   I H   N
3 T     S     G

当列数是：0 3 6 9，需要填入完整列
1 2
1 5
2 1
2 4
2 7
当列数是：1 4，需要在行2上填入
当列数是：2 5，需要在行1上填入

numRows=5 3
	0 1 2 3 4 5 6 7
0 L       I
1 E     E S     G
2 E   D   H   N
3 T O     I I
4 C       R

当列数是：0 4 8 12，需要填入完整列
1 3
1 7
2 2
2 6
3 1
3 5
当列数是：1 5，需要在行3上填入
当列数是：2 6，需要在行2上填入
当列数是：3 7，需要在行1上填入

列数：c
even: c%(numRows-1) == 0，填入完整列
odd:  c%(numRows-1) == ![0, numRows-1] (降序)
*/

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	var (
		rs     []byte
		delter = 2*numRows - 2
		b      = []byte(s)
	)
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j += delter {
			rs = append(rs, b[j+i])
			if i != 0 && i != numRows-1 && j+delter-i < len(s) {
				rs = append(rs, b[j+delter-i])
			}
		}
	}

	return string(rs)
}

func convertByRow(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	r := make([]string, numRows)

	var (
		row    int
		delter = -1
	)
	for _, c := range s {
		r[row] += string(c)
		if row == 0 || row == numRows-1 {
			delter = -delter
		}

		row += delter
	}

	var rs string
	for _, ss := range r {
		rs += ss
	}

	return rs
}

func convertWorkingHard(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	p := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		p[i] = make([]byte, len(s))
	}

	b := []byte(s)
	for col, i := 0, 0; i < len(b); col++ {
		// println(col)
		c := col % (numRows - 1)
		if c == 0 {
			for row := 0; row < numRows && i < len(b); row++ {
				p[row][col] = b[i]
				i++
			}
		} else {
			p[numRows-1-c][col] = b[i]
			i++
		}
	}

	var fs []byte
	for _, cols := range p {
		for _, col := range cols {
			if col > 0 {
				fs = append(fs, col)
				print(string(col), " ")
			} else {
				print("  ")
			}
		}
		println()
	}

	return string(fs)
}

// func main() {
//	// fs := map[int]string{
//	//		3: `L   C   I   R
//	// E T O E S I I G
//	// E   D   H   N`,
//	//		4: `L     D     R
//	// E   O E   I I
//	// E C   I H   N
//	// T     S     G`,
//	//	}
//	ss := map[int]string{
//		// 3: "LEETCODEISHIRING",
//		// 4: "LEETCODEISHIRING",
//		3: "PAYPALISHIRING",
//		1: "A",
//	}

//	for numRows, s := range ss {
//		// println(convert1(s, numRows))
//		println(convert(s, numRows))
//		// println(convertByRow(s, numRows))
//		// println(fs[numRows])
//		// println()
//	}
// }
