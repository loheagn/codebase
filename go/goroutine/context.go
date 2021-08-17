package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	childCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	go func() {
		//for {
			select {
			case <- childCtx.Done():
				fmt.Println(childCtx.Err())
				fmt.Println("Child go routine end.")
			default:
				//fmt.Println(childCtx.Err())
			}
		//}
	}()
	select {
	case <- ctx.Done():
		fmt.Println("Parent go routine end.")
		return
	}
}
