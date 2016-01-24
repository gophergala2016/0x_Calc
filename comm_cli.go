// comm
package main

import (
	"encoding/json"
	"fmt"
	"net"
	_ "time"
)

func main() {

	go Client()

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

func catch() { // Panic handler
	s := recover()
	if s != nil {
		fmt.Println(s)
	}
}

// Client side
// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
func Client() {
	defer catch()

	cli, err := net.Dial("tcp", "127.0.0.1:65097")
	Alert(err, "Dial Failed")

	// Dial Success -> Reserve closing
	defer cli.Close()

	opnd := Operands{Lhs: 60, Rhs: 20}

	json_opnd, err := json.Marshal(opnd)
	Alert(err, "JSON encoding failed")

	cli.Write(json_opnd)

	var json_res []byte
	_, err = cli.Read(json_res)
	Alert(err, "JSON Result read failed")

	var res Result

	err = json.Unmarshal(json_res, res)
	Alert(err, "JSON Unmarshal Failed")

	fmt.Println(res)
}

func Alert(_e error, _cmt string) {
	if _e != nil {
		fmt.Println(_e)
		panic(_cmt)
	}
}
