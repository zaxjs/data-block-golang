package data_block

import (
	"fmt"
	"log"

	"github.com/bytedance/sonic"
)

func init() {
	// Do some init
	log.SetPrefix("[data-block]: ")
	log.SetFlags(0)
}

func ExposeFn1() {
	data := [2]string{"a", "b"}
	// Marshal
	output, err := sonic.Marshal(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
	// Unmarshal
	err = sonic.Unmarshal(output, &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}

func ExposeFn2() {
	fmt.Println("test public member")
}
