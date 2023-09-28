package main

import "C"
import (
	"fmt"
	"go_leran/src/routers"
)

func main() {
	r := routers.SetupRouter()
	err := r.Run()
	if err != nil {
		fmt.Printf("start failed, err: %v", err)
	}

}
