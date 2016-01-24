// comm
package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {

	Server()
	fmt.Scanln()
}

// Shared type
// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

type Operands struct {
	Lhs, Rhs int64
}

type Result struct {
	Value int64
}

func ADD(_l int64, _r int64) Result {
	return Result{Value: _l + _r}
}
func SUB(_l int64, _r int64) Result {
	return Result{Value: _l - _r}
}
func MUL(_l int64, _r int64) Result {
	return Result{Value: _l * _r}
}
func DIV(_l int64, _r int64) Result {
	return Result{Value: _l / _r}
}

func catch() { // Panic handler
	s := recover()
	if s != nil {
		fmt.Println(s)
	}
}

// Server side
// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
func Server() {
	defer catch()

	lsner, err := net.Listen("tcp", ":65097")

	if err != nil {
		// Listener allocation Failed
		fmt.Println(err)
		panic("Server Listener failed")
	} else {
		// Listener allocation Success
		// Reserve closing
		defer lsner.Close()
	}
	fmt.Println("Listening...")

	for {
		conn, err := lsner.Accept()
		if err != nil {
			// Client connection Failed
			// Do Nothing
			fmt.Println("Accept failed")
			continue
		} else {
			fmt.Println("Accepted")
		}

		content := make([]byte, 4096)

		_, err = conn.Read(content)
		Alert(err, "Server Read Failed")

		var opnd Operands
		err = json.Unmarshal(content, &opnd)
		Alert(err, "Unmarshal Failed")

		var res Result
		res = ADD(opnd.Lhs, opnd.Rhs)
		fmt.Println(res)

		bt, err := json.Marshal(res)
		Alert(err, "JSON encoding failed")

		_, err = conn.Write(bt)
		Alert(err, "Server Read Failed")

	}
}

func Alert(_e error, _cmt string) {
	if _e != nil {
		fmt.Println(_e)
		panic(_cmt)
	}
}
