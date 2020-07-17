package main

import "fmt"

func convert(s string, numRows int) string {
    bs := []byte(s)
    res := make([]byte, len(bs))
    k := 0
    tmp := (numRows-1) * 2

    for i:=1; i<= numRows/2; i++ {
    	interval := (numRows - i) * 2
		ind := i-1
		if i == 1 {
			interval = tmp
			for next := ind; next < len(bs); next = next + interval{
				res[k] = bs[next]
				k++
			}
		} else {
			flag := true
			cnt := 1
			for next := ind; next < len(bs); {
				res[k] = bs[next]
				k++
				if flag {
					next = next + interval
					flag = false
				} else {
					flag = true
					next = i-1+tmp*cnt
					cnt++
				}
			}
		}
	}

    for i:=numRows/2+1; i<=numRows; i++ {
		interval := (numRows - i) * 2
		ind := i-1
		if i == numRows{
			interval = tmp
			for next := ind; next < len(bs); next = next + interval{
				res[k] = bs[next]
				k++
			}
		} else {
			flag := true
			cnt := 1
			for next := ind; next < len(bs); {
				res[k] = bs[next]
				k++
				if flag {
					next = next + interval
					flag = false
				} else {
					flag = true
					next = i-1+tmp*cnt
					cnt++
				}
			}
		}
	}
    return string(res)
}


func main() {
	s := "PAYPALISHIRING"
	numRows:= 4
	fmt.Println(convert(s, numRows))
}
