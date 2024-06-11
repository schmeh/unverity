//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
	"unverity/cmd/wasm/swaps"
)

func main() {
	fmt.Println("Dissections wasm loaded")
	js.Global().Set("dissect", dissectWrapper())
	<-make(chan struct{})
}

func dissectWrapper() js.Func {
	dissectFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 4 {
			return "Invalid number of arguments passed"
		}
		soloCallout := args[0].String()
		currentShapes := []string{args[1].String(), args[2].String(), args[3].String()}
		s, err := swaps.GetSwaps(soloCallout, currentShapes)
		if err != nil {
			fmt.Printf("unable to get dissections %s\n", err)
			return err.Error()
		}
		var rtn string
		for i, i2 := range s {
			rtn += fmt.Sprintf("%d. %s", i+1, i2.Export())
		}
		return rtn
	})
	return dissectFunc
}
