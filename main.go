package main

import "yaml2go/task"

func main() {
	task.Convert(
		`
test:
  test1:
    test2: aaa
    test3: bbb
  test4:
    test5: ccc
  test6: ddd
test0: eee
`, "    ")
}
