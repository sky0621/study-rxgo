package main

import (
	"fmt"

	"net/http"

	"github.com/reactivex/rxgo/errors"
	"github.com/reactivex/rxgo/iterable"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func main() {
	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			fmt.Printf("Processing: %v\n", item)
		},

		ErrHandler: func(err error) {
			fmt.Printf("Encounter error: %v\n", err)
		},

		DoneHandler: func() {
			fmt.Printf("Done!")
		},
	}

	it, err := iterable.New([]interface{}{1, 2, 3, 4, errors.New(http.StatusOK), 5})
	if err != nil {
		panic(err)
	}

	source := observable.From(it)
	sub := source.Subscribe(watcher)

	<-sub
}
