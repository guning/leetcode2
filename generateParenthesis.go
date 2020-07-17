package main

import "fmt"

func generateParenthesis(n int) []string {
	ch := make(chan string, n)
	var res []string

	go func () {
		gp(ch, "", 0, 0, n)
		close(ch)
	}()
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		res = append(res, v)
	}

	return res
}

/*
                     (
        (                         )
 (              )             (
   )        (        )     (      )
     )        )    (          )     (
       )        )     )         )     )
相当于在限制条件下遍历打印这颗二叉树，对于每个节点
放“(”必须满足历史数量小于n
放“）”必须满足历史“（”数量大于“）”
 */

func gp(ch chan string, path string, i, j, n int) {
	if len(path) == n * 2 {
		ch <- path
		return
	}

	if i < n {
		gp(ch, path + "(", i+1, j, n)
	}

	if j < i {
		gp(ch, path + ")", i, j+1, n)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}