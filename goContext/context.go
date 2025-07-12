package main

import (
	"context"
	"fmt"
)

func main() {
	contextParent := context.Background()
	ctx1 := context.WithValue(contextParent, "key1", "halo dunia")
	ctx2 := context.WithValue(ctx1, "key2", "halo maniest")
	ctx3 := context.WithValue(ctx2, "key3", "halo my")
	ctx4 := context.WithValue(contextParent, "key4", "halo cuy")
	ctx5 := context.WithValue(ctx1, "key5", "halo beg")

	fmt.Println(ctx5.Value("key5"))
	fmt.Println(ctx5.Value("key4"))
	fmt.Println(ctx5.Value("key3"))
	fmt.Println(ctx5.Value("key2"))
	fmt.Println(ctx5.Value("key1"))
	fmt.Println(ctx3.Value("key1"))
	fmt.Println(ctx4.Value("key1"))
}