package main

import (
	"context"
	"fmt"
	"time"
)

type paramKey struct {
}

func main() {
	c := context.WithValue(context.Background(), paramKey{}, "abc")
	c, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()
	mainTask(c)
}

func mainTask(c context.Context) {
	fmt.Printf("main task started with param %q\n", c.Value(paramKey{}))
	smallTask(c, "task1", 4*time.Second)
	smallTask(c, "task2", 2*time.Second)

}

func smallTask(c context.Context, name string, d time.Duration) {
	fmt.Printf("%s started with param %q \n", name, c.Value(paramKey{}))
	select {
	case <-time.After(d):
		fmt.Printf("%s cancelled\n", name)

	case <-c.Done():
		fmt.Printf("%s done\n", name)
	}
}
