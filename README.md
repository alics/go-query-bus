# go-query-bus


## Install

Use go get.
```sh
$ go get github.com/alics/go-query-bus
```

Then import the package into your own code:
```
import "github.com/alics/go-query-bus/core"
```

## Usage
```go
package main

import (
	"context"
	"fmt"
	"github.com/alics/go-query-bus/core"
	"log"
	"os"
)
import . "github.com/ahmetb/go-linq/v3"

type personQueryHandler struct {
}

func (p personQueryHandler) Handle(ctx context.Context, filter interface{}) (result interface{}, err error) {
	var owners []Person

	f, _ := filter.(*PersonQueryFilter)

	From(people).Where(func(c interface{}) bool {
		return c.(Person).Name == f.Name
	}).Select(func(c interface{}) interface{} {
		return c.(Person)
	}).ToSlice(&owners)

	return owners, nil
}

type Person struct {
	Name string
	Age  int
}

var p1 = Person{"Ali", 32}
var p2 = Person{"Reza", 27}
var p3 = Person{"Kaveh", 64}
var p4 = Person{"Farhad", 56}

var people = []Person{p1, p2, p3, p4}

type PersonQueryFilter struct {
	Name string
}

func NewAPersonQueryHandler() core.IQueryHandler {
	return &personQueryHandler{}
}
func main() {
	h := NewAPersonQueryHandler()
	bus := core.New()
	err := bus.Register(&PersonQueryFilter{}, h)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	filter := PersonQueryFilter{
		Name: "Ali",
	}
	result, err := bus.Execute(context.Background(), &filter)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Print(result)
	}
}
```
