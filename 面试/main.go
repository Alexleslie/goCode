package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		result := 0
		borad := [][]byte{}
		for j := 0; j < 3; j++ {
			var s string
			fmt.Scan(&s)
			temp := []byte{s[0], s[1], s[2]}
			borad = append(borad, temp)
		}
		dict1 := make(map[int]byte)
		dict2 := make(map[int]byte)
		dict3 := make(map[int]byte)
		dictlist := []map[int]byte{dict1, dict2, dict3}
		dictmap := make(map[int]map[int]byte)
		for j := 0; j < 3; j++ {
			var x, y int
			for k := 0; k < 3; k++ {
				fmt.Scan(&x)
				fmt.Scan(&y)
				dictlist[j][random(x, y)] = borad[x][y]
				dictmap[random(x, y)] = dictlist[j]
			}
		}

		var curr func(c int)

		curr = func(c int) {
			if c == 9 {
				result += 1
				return
			}
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if borad[i][j] != '*' {
						continue
					} else {
						for k := 0; k < 3; k++ {
							insert := strconv.Itoa(k + 1)[0]
							grida := gridCheck(dict1, dict2, dict3, i, j, k+1)
							for row := 0; row < 3; row++ {
								if borad[row][j] == insert {
									continue
								}
							}
							for col := 0; col < 3; col++ {
								if borad[i][col] == insert {
									continue
								}
							}
							if grida {
								dictmap[random(i, i)][random(i, j)] = insert
								borad[i][j] = insert
								curr(c + 1)
								borad[i][j] = '*'
								dictmap[random(i, i)][random(i, j)] = '*'
							}
						}
					}

				}

			}

		}
		all := 0
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if borad[x][y] != '*' {
					all += 1
				}
			}
		}
		curr(all)
		fmt.Println(result)
	}

}

func gridCheck(d1, d2, d3 map[int]byte, x, y, i int) bool {
	var d map[int]byte
	mapv := random(x, y)
	if d1[mapv] != '0' {
		d = d1
	} else if d2[mapv] != '0' {
		d = d2
	} else {
		d = d3
	}

	si := strconv.Itoa(i)
	for _, v := range d {
		if string(v) == si {
			return false
		}
	}
	return true
}

func random(x, y int) int {
	return (x+1)*11 + (y+1)*231
}
