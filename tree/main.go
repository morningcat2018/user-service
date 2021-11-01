package main

import "tree/avl"

func main() {
	root := avl.BuildTreeByData(17)
	root.Put(20)
	root.Put(10)
	root.Put(19)
	root.Put(30)
	root.Put(23)
	//root.Put(36)
	//root.Put(2)
	root.MidOrderErgodic()
}
