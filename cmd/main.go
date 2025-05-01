package main

import (
	"bufio"
	"fmt"
	"os"

	task "github.com/jachant/YadroTest/task"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int
	fmt.Fscan(in, &n)
	containers := make([][]uint64, n)
	for i := 0; i < n; i++ {
		containers[i] = make([]uint64, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &containers[i][j])
		}
	}
	result := task.Task1(containers)
	fmt.Fprintln(out, result)
}
