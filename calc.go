//////////////////////////
//			//
//	  0x_Calc	//
//			//
//			//
//			//
//////////////////////////

package main

import (
	_ "encoding/hex"
	_ "fmt"
	"strconv"
	_ "strings"
)

// Store val, operator etc....
var Data struct {
	l        string //
	val      string // stores val result
	operator string // the next operator to use
}

/*
// Split
func Calc() {
	fmt.Print("Split by SPACE\n Numerical Formula : ")
	fmt.Scan("%s",Data.l)

	words := strings.Split(Data.l, " ")
	for i := range words {		// i for space

		if (i/2 == 1) && (words[i] == "+" || words[i] == "-" || words[i] == "*" || words[i] == "/" || words[i] =="%" || words[i] == "&" || words[i] == "|" || words[i] == "^|" || words[i] == "^&" || words[i] == "<<" || words[i] == ">>" || word s[i] == "^") {		// Unsigned Number

			for j := range words {		// j for word
				if j/2 == 1 {

					alpa := strings.Split(words[j], []byte(words[j]))
					for a :=range alpa {
						if (alpa[0:2] == "0x") && (alpa[a] >= "0" && alpa[a] <= "9") || (alpa[a] >= "A" && alpa[a] <= "F") { // 0x -> Hex, 0 -> Oct, else -> float, int
							// This val is Hex
							fmt.Printf("Hex : %s\n",words[j])
								Data.operator = words[i]
								function(Data.operator)
						}else if (alpa[0] == "0") && (alpa[a] >= "0" && alpa[a] <= "7") {
							// This val is Oct
							fmt.Printf("Oct : %s\n",words[j])
						}else {
							// This val is float or int
							fmr.Printf("float :%s\n",words[j])
						}
					}
				}
			}
		} else if (i/2 == 0) && (words[i] == "+" || words[i] == "-" || words[i] == "*" || words[i] == "/" || words[i] =="%" || words[i] == "&" || words[i] == "|" || words[i] == "^|" || words[i] == "^&" || words[i] == "<<" || words[i] == ">>" || words[i] == "^") {		// Signed Number
			 for j := range words {          // j for word
                                if (j/2 == 1) {

                                        alpa := strings.Split(words[j], []byte(words[j]))
                                        for a := range alpa {
                                                if (alpa[0:2] == "0x") && (alpa[a] >= "0" && alpa[a] <= "9") || (alpa[a] >= "A" && alpa[a] <= "F") { // 0x -> Hex, 0 -> Oct, else -> float, int
                                                        // This val is Hex
                                                        fmt.Printf("Hex : %s\n",words[j])
                                                                Data.operator = words[i]
                                                                function(Data.operator)
                                                }else if (alpa[0] == "0") && (alpa[a] >= "0" && alpa[a] <= "7") {
                                                        // This val is Oct
                                                        fmt.Printf("Oct : %s\n",words[j])
                                                }else {
                                                        // This val is float or int
                                                        fmr.Printf("float :%s\n",words[j])
                                                }
                                        }
				}
			}
		}else {			// Wrong Syntex Or Operator
			fmt.Println("err")
		}
	}
}
*/

/*
// Calculation
func function(_oper string) {
	switch _oper {
	case "+":
		sum(Data.val, tmp)
		//Data.val += tmp
	case "-":
		sub(Data.val, tmp)
		//Data.val -= tmp
	case "*":
		mul(Data.val, tmp)
		//Data.val *= tmp
	case "/":
		div(Data.val, tmp)
		//Data.val /= tmp
	case "%":
		etc(Data.val, tmp)
		//Data.val %= tmp
	case "&":
		and(Data.val, tmp)
		//Data.val &= tmp
	case "|":
		or(Data.val, tmp)
		//Data.val |= tmp
	case "^|":
		xor(Data.val, tmp)
		//Data.val ^= tmp
	case "^&":
		notand(Data.val, tmp)
		//Data.val &^= tmp
	case "<":
		lshift(Data.val, tmp)
		//Data.val = Data.val << uint64(tmp)
	case ">":
		rshift(Data.val, tmp)
		//Data.val = Data.val >> uint64(tmp)
	case "^":
		rever(Data.val)
		//Data.val = ^Data.val
	default:
		fmt.Println("err")
	}
}
*/
//---------------------------------------------------------------

//Operator
func sum(val string, val2 string) {
	//val += val2
	//return val
}

func sub(val string, val2 string) {
	//val -= val2
	//return val
}

func mul(val string, val2 string) {
	//val *= val2
	//return val
}

func div(val string, val2 string) {
	//val /= val2
	//return val
}

func etc(val string, val2 string) {
	//val %= val2
	//return val
}

func and(val string, val2 string) {
	//val &= val2
	//return val
}

func or(val string, val2 string) {
	//val |= val2
	//return val
}

func xor(val string, val2 string) {
	//val ^|= val2
	//return val
}

func notand(val string, val2 string) {
	//val ^%= val2
	//return val
}

func lshift(val string, val2 int) {
	//val =<< val2
	//return val
}

func rshift(val string, val2 int) {
	//val =>> val2
	//return val
}

func rever(val string) {
	//val = ^val
	//return val
}

/*
//String to Float
func atoiconvert(val string, val2 string) {
	i, _ := strconv.ParseFloat(val, 64)
	j, _ := strconv.ParseFloat(val2, 64)

	ftoiconvert(i, j)
}
*/

/*
//Float to Int
func ftoiconvert(val float32, val2 float32) {
	a := int(val)
	b := int(val2)

	tmp1 := int(val) - a
	tmp2 := int(val2) - b

	if !(tmp1 < 0.0) || !(tmp2 < 0.0) {
		return val, val2
	} else {
		return a, b
	}
}
*/

// Dec int -> Hex string
func Dec_to_HexStr(_lhs int, _rhs int) (hex_l string, hex_r string) {
	hex1 := strconv.FormatInt(int64(_lhs), 16)
	hex2 := strconv.FormatInt(int64(_rhs), 16)
	return hex1, hex2
}

func Dec_to_OctStr(_lhs int, _rhs int) (oct_l string, oct_r string) {
	oct1 := strconv.FormatInt(int64(_lhs), 8)
	oct2 := strconv.FormatInt(int64(_rhs), 8)
	return oct1, oct2
}

/*
//Int to String || Float to String
func tosconvert(float32 val, float32 val2) {
	s1 := strconv.FormatFloat(val, 'f', 6, 32)
	s2 := strconv.FormatFloat(val2, 'f', 6, 32)

	//------
}
*/
