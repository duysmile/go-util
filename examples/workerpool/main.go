package main

import (
	"fmt"
	"github.com/duysmile/go-util/workerpool"
)

func main() {
	pool := workerpool.NewPool(workerpool.Config{
		NumOfWorker: 2,
	})

	pool.RegisterHandler(func(input interface{}) workerpool.Response {
		fmt.Println("handling: ", input)
		return workerpool.Response{
			Data:  input,
			Error: nil,
		}
	})

	iChan, oChan := pool.Run()

	data := []int{1, 2, 3, 4, 5, 6}
	go func(pool *workerpool.Pool, data []int) {
		for _, i := range data {
			iChan <- i
		}
		pool.Finish()
	}(pool, data)

	results := make([]interface{}, 0, len(data))
	for result := range oChan {
		results = append(results, result.Data)
	}

	fmt.Println("results: ", results)
}
