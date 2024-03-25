package main

import (
	"contacts1/components"
	archiver "contacts1/internal"
	"context"
	"fmt"
	"strings"
)

func main() {

	c := components.Contacts{}
	c.New("test1", "testemail1")
	c.New("test2", "testemail2")

	fmt.Println(string(*c.Bytes()))

	a := archiver.Archiver{}
	rd := strings.NewReader(string(*c.Bytes()))
	// ctx, _ := context.WithCancel(context.Background())
	go a.Run(context.Background(), rd)

}
