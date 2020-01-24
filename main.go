package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Value struct {
	Label string
	Num   int64
	Text  string
}

func main() {
	logger, err := NewLogger("sample", 0)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := logger.Close(); err != nil {
			log.New(os.Stderr, "", 0).Printf("failed logger.Close() err=%v\n", err)
		}
	}()

	nums := []int64{
		36028797018963970,
		36028797018963969,
		36028797018963968,
		36028797018963967,
		36028797018963966,
		36028797018963965,
		36028797018963964,
		36028797018963963,
		36028797018963962,
		36028797018963961,
		36028797018963960,
		36028797018963959,
		36028797018963958,
		36028797018963957,
		36028797018963956,
		36028797018963955,
		36028797018963954,
		36028797018963953,
		36028797018963952,
		36028797018963951,
		36028797018963950,
		36028797018963859,
		26028797018963969}

	for {
		for _, num := range nums {
			output(num)
		}
		logger.WriteString("world")
		time.Sleep(6 * time.Second)
	}
}

func output(num int64) {
	v := Value{
		Label: "36028797018963960",
		Num:   num,
		Text:  fmt.Sprint(num),
	}
	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", uint64(num))
	fmt.Println(string(j))
}
