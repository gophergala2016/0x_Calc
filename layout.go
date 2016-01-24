// main_layout.go
package main

import (
	"fmt"
	"os"
	"strconv"
	_ "unsafe"

	gdk "github.com/mattn/go-gtk/gdk"
	_ "github.com/mattn/go-gtk/glib"
	gtk "github.com/mattn/go-gtk/gtk"
)

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
// Constants
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
const (
	// For window size
	UI_Width  = 640
	UI_Height = 360

	// for Positioning
	No_Resize = false // No resize in Pack#()
	No_Shrink = false // No shrink in Pack#()

	Homogeneous   = false
	Heterogeneous = true

	Default_Spacing = 1
)

// Main function.
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

func main() {
	defer catch() // catch Panic

	Start()
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
// Type : UI
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

const (
	CVT_HEX = 16
	CVT_DEC = 10
	CVT_OCT = 8
)

type UI struct {
	// GUI components
	Win        *gtk.Window // Main Window
	Calc_Frame *gtk.Frame  // Frame for Calculation
	Nums_Frame *gtk.Frame  // Frame for Number buttons
	Oper_Frame *gtk.Frame  // Frame for Operator buttons
	Lbl_prev   *gtk.Label  // Label for previous result
	Lbl_lhs    *gtk.Label  // Label for Left hand side value
	Lbl_rhs    *gtk.Label  // Label for Right hand side value

	Btn_map map[string]*gtk.Button // Map of Buttons

	// Member for Operation
	mode     int
	csr      bool
	prev     int              // previous result
	lhs      int              // operand type is int
	rhs      int              // operand type is int
	Ch_Event chan interface{} // Channel for event streaming
}

// Constructor function
//  	1. Create Components
//  	2. Setup Layout
//  	3. Event - Callback
func (this *UI) Construct() {
	defer catch()
	// 1. Creating Components
	// ---- ---- ---- ---- ---- ---- ---- ----

	// Create the Main Window
	// Set title & size
	this.Win = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	if this.Win == nil {
		panic("UI::Construct() : Window allocation Failed")
	}
	window := this.Win               // Window aliasin
	window.SetTitle("0x_Calculator") // SetTitle

	this.Calc_Frame = gtk.NewFrame("Calculation")
	this.Nums_Frame = gtk.NewFrame("Numbers")
	this.Oper_Frame = gtk.NewFrame("Operation")

	this.Lbl_prev = gtk.NewLabel("(Previous)")
	this.Lbl_lhs = gtk.NewLabel("(LHS)")
	this.Lbl_rhs = gtk.NewLabel("(RHS)")

	this.Btn_map = make(map[string]*gtk.Button)
	this.Ch_Event = make(chan interface{})

	this.csr = false
	this.mode = CVT_DEC

	// 2. Setup Layout
	// ---- ---- ---- ---- ---- ---- ---- ----
	this.init_Calc()
	this.init_Nums()
	this.init_Oper()
	this.put_frames()

	// 3. Event - Callback connection
	// ---- ---- ---- ---- ---- ---- ---- ----
	this.init_Events()

	// 4. Left overs
	// ---- ---- ---- ---- ---- ---- ---- ----
	window.SetSizeRequest(UI_Width, UI_Height)
	window.ShowAll()
}

// Destructor function
func (this *UI) Destruct() {
	// Panic handler for Safe clean up
	defer catch()

	fmt.Println("UI::Destruct() : this UI is destructing...")
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
// UI : Initializers
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

// Frame - Calculation
// This frame contains radix(16,10,8) and result labels
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
func (this *UI) init_Calc() {
	// In this function, the designated frame is Calc_Frame
	_Frame := this.Calc_Frame
	if _Frame == nil {
		panic("UI::init_Calc() : nil Frame received")
	}

	// (inner) Box of Calculation
	fm_calc_box := gtk.NewHBox(false, 1)
	if fm_calc_box == nil {
		panic("UI::init_Calc() : HBox allocation Failed")
	}

	_Frame.Add(fm_calc_box)

	// Box for Radix Buttons.
	box_rdx := gtk.NewVBox(false, 1)
	if box_rdx == nil {
		panic("UI::init_Calc() : VBox allocation Failed")
	}

	btn_hex := button("Hex") // [Hex] : Hexadecimal
	btn_dec := button("Dec") // [Dec] : Decimal
	btn_oct := button("Oct") // [Oct] : Octal
	box_rdx.Add(btn_hex)
	box_rdx.Add(btn_dec)
	box_rdx.Add(btn_oct)

	// Insert radix buttons into the map
	this.Btn_map["Hex"] = btn_hex
	this.Btn_map["Dec"] = btn_dec
	this.Btn_map["Oct"] = btn_oct

	// Box for Result Labels
	box_labels := gtk.NewVBox(false, 1)
	if box_labels == nil {
		panic("UI::init_Calc() : VBox allocation Failed")
	}

	// Place previous result
	box_labels.Add(this.Lbl_prev)

	// Place left and right operand
	box_LnR := gtk.NewHBox(false, 3)
	if box_LnR == nil {
		panic("UI::init_Calc() : HBox allocation Failed")
	}
	box_LnR.Add(this.Lbl_lhs)
	box_LnR.Add(this.Lbl_rhs)
	box_labels.Add(box_LnR)

	// Add both Boxes (Radix & Result) to frame box
	fm_calc_box.Add(box_rdx)
	fm_calc_box.Add(box_labels)

	fmt.Println("UI::init_Calc() done.")
}

// Frame - Numbers
// This frame contains number buttons for calculation
//  	Hexadecimal	: 0 ~ 9, A ~ F
//  	Decimal   	: 0 ~ 9
//  	Octal     	: 0 ~ 7
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
func (this *UI) init_Nums() {
	// In this function, the designated frame is Nums_Frame
	_Frame := this.Nums_Frame

	if _Frame == nil {
		panic("UI::init_Nums() : nil Frame received")
	}

	// (inner) Box of Numbers
	fm_nums_box := gtk.NewVBox(false, 1)
	if fm_nums_box == nil {
		panic("UI::init_Nums() : VBox allocation Failed")
	}
	// Add to given frame
	_Frame.Add(fm_nums_box)

	// Table Initialization
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
	tbl_nums := gtk.NewTable(5, 4, false)
	if tbl_nums == nil {
		panic("UI::init_Nums() : Table allocation Failed")
	}
	// Jagged slice of buttons?
	// nums := [][]*gtk.Button{}

	// Button for Number
	num := [...]*gtk.Button{
		// 0~7 : Oct
		button("0"), button("1"), button("2"), button("3"),
		button("4"), button("5"), button("6"), button("7"),
		// 0~9 : Dec
		button("8"), button("9"),
		// A~F : Hex
		button("A"), button("B"), button("C"),
		button("D"), button("E"), button("F"),
	}

	// Insert all Buttons to map
	for idx, btn := range num {
		//fmt.Println("UI::init_Nums() : Index :\t", idx, "Button :\t", btn.GetLabel())
		s_idx := strconv.Itoa(idx)
		this.Btn_map[s_idx] = btn
	}

	// Place buttons into the table
	tbl_nums.Attach(num[0], 0, 1, 3, 4, gtk.FILL, gtk.FILL, 1, 1) // 0
	tbl_nums.Attach(num[1], 0, 1, 2, 3, gtk.FILL, gtk.FILL, 1, 1) // 1
	tbl_nums.Attach(num[2], 1, 2, 2, 3, gtk.FILL, gtk.FILL, 1, 1)
	tbl_nums.Attach(num[3], 2, 3, 2, 3, gtk.FILL, gtk.FILL, 1, 1)
	tbl_nums.Attach(num[4], 0, 1, 1, 2, gtk.FILL, gtk.FILL, 1, 1) // 4
	tbl_nums.Attach(num[5], 1, 2, 1, 2, gtk.FILL, gtk.FILL, 1, 1)
	tbl_nums.Attach(num[6], 2, 3, 1, 2, gtk.FILL, gtk.FILL, 1, 1)
	tbl_nums.Attach(num[7], 0, 1, 0, 1, gtk.FILL, gtk.FILL, 1, 1) // 7
	tbl_nums.Attach(num[8], 1, 2, 0, 1, gtk.FILL, gtk.FILL, 1, 1)
	tbl_nums.Attach(num[9], 2, 3, 0, 1, gtk.FILL, gtk.FILL, 1, 1)
	tbl_nums.Attach(num[10], 3, 4, 2, 3, gtk.FILL, gtk.FILL, 1, 1) // A
	tbl_nums.Attach(num[11], 4, 5, 2, 3, gtk.FILL, gtk.FILL, 1, 1)
	tbl_nums.Attach(num[12], 3, 4, 1, 2, gtk.FILL, gtk.FILL, 1, 1)
	tbl_nums.Attach(num[13], 4, 5, 1, 2, gtk.FILL, gtk.FILL, 1, 1)
	tbl_nums.Attach(num[14], 3, 4, 0, 1, gtk.FILL, gtk.FILL, 1, 1)
	tbl_nums.Attach(num[15], 4, 5, 0, 1, gtk.FILL, gtk.FILL, 1, 1) // F

	// Add the table to box
	fm_nums_box.Add(tbl_nums)
	fmt.Println("UI::init_Nums() done.")
}

// Frame - Operations
// This frame contains operations.
//  	ADD, SUB, MUL, DIV, MOD
//  	AND, OR, XOR, NOT
//  	LSHFT, RSHFT
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
func (this *UI) init_Oper() {
	// In this function, the designated frame is Oper_Frame
	_Frame := this.Oper_Frame
	if _Frame == nil {
		panic("UI::init_Oper() : nil Frame received")
	}

	// (inner) Box of Operations
	fm_oper_box := gtk.NewVBox(false, 1)
	if fm_oper_box == nil {
		panic("UI::init_Nums() : VBox allocation Failed")
	}

	_Frame.Add(fm_oper_box)

	tbl_opers := gtk.NewTable(5, 3, false)
	if tbl_opers == nil {
		panic("UI::init_Nums() : Table allocation Failed")
	}

	// Operation Buttons
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

	// 'oper' is Jagged slice of buttons
	oper := [][]*gtk.Button{}
	// slice of Arithmetic
	oper_arit := []*gtk.Button{
		button("ADD"), button("SUB"), button("MUL"), button("DIV"),
		button("MOD")}
	// slice of Bitwise
	oper_bit := []*gtk.Button{
		button("AND"), button("OR"), button("XOR"), button("NOT")}
	// slice of Bit Shift
	oper_shft := []*gtk.Button{
		button("LSHIFT"), button("RSHIFT")}

	// Compose the jagged slice
	oper = append(oper, oper_arit)
	oper = append(oper, oper_bit)
	oper = append(oper, oper_shft)

	// Iterate jagged slice and place them into the table
	for r, btn_slice := range oper {
		// r : row
		// btn_slice : slice of buttons
		for c, btn := range btn_slice {
			// c : column
			// btn == btn_slice[c] == oper[row][col]
			// Place the button to table
			tbl_opers.Attach(btn, uint(c), uint(c)+1, uint(r), uint(r)+1,
				gtk.FILL, gtk.FILL, 1, 1)
		}
	}

	btn_done := button("=")
	tbl_opers.Attach(btn_done, uint(4), uint(4)+1, uint(2), uint(2)+1,
		gtk.FILL, gtk.FILL, 1, 1)

	// Insert all buttons to button map
	this.Btn_map["="] = btn_done

	this.Btn_map["ADD"] = oper_arit[0]
	this.Btn_map["SUB"] = oper_arit[1]
	this.Btn_map["MUL"] = oper_arit[2]
	this.Btn_map["DIV"] = oper_arit[3]
	this.Btn_map["MOD"] = oper_arit[4]
	this.Btn_map["AND"] = oper_bit[0]
	this.Btn_map["OR"] = oper_bit[1]
	this.Btn_map["XOR"] = oper_bit[2]
	this.Btn_map["NOT"] = oper_bit[3]
	this.Btn_map["LSHFT"] = oper_shft[0]
	this.Btn_map["RSHFT"] = oper_shft[1]
	fm_oper_box.Add(tbl_opers)

	fmt.Println("UI::init_Oper() done.")
}

// Locate frames on the window
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
func (this *UI) put_frames() {
	box_win := gtk.NewVBox(Heterogeneous, Default_Spacing)
	if box_win == nil {
		panic("UI::put_frames() : VBox allocation Failed")
	}

	vpan1 := gtk.NewVPaned()
	if vpan1 == nil {
		panic("UI::put_frames() : VPaned allocation failed")
	}

	vpan1.Pack1(this.Calc_Frame, No_Resize, No_Shrink) // Calc : Top half
	hpan1 := gtk.NewHPaned()
	if hpan1 == nil {
		panic("UI::put_frames() : HPaned allocation failed")
	}

	hpan1.Pack1(this.Nums_Frame, No_Resize, No_Shrink) // Nums : Bottom-Left
	hpan1.Pack2(this.Oper_Frame, No_Resize, No_Shrink) // Oper : Bottom-Right

	vpan1.Pack2(hpan1, No_Resize, No_Shrink)
	box_win.Add(vpan1)

	if this.Win == nil {
		panic("UI::put_frames() : nil Window received")
	}
	// Place all Layout
	this.Win.Add(box_win)
	fmt.Println("UI::put_frames() done.")
}

// UI Event - Callback binding
//  	Window : "destroy"
//  	Button : "clicked"
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
func (this *UI) init_Events() {
	// on Exit -> Quit the program
	this.Win.Connect("destroy", gtk.MainQuit)

	// Set Listening : Keyboard
	this.Win.SetEvents(int(gdk.BUTTON_PRESS_MASK))

	// Map name aliasing
	btn := this.Btn_map

	// Format Buttons
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
	btn["Hex"].Connect("clicked", func() { this.switch_format(16) })
	btn["Dec"].Connect("clicked", func() { this.switch_format(10) })
	btn["Oct"].Connect("clicked", func() { this.switch_format(8) })

	// Number Buttons
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
	btn["0"].Connect("clicked", func() { this.handle_num(0) })
	btn["1"].Connect("clicked", func() { this.handle_num(1) })
	btn["2"].Connect("clicked", func() { this.handle_num(2) })
	btn["3"].Connect("clicked", func() { this.handle_num(3) })
	btn["4"].Connect("clicked", func() { this.handle_num(4) })
	btn["5"].Connect("clicked", func() { this.handle_num(5) })
	btn["6"].Connect("clicked", func() { this.handle_num(6) })
	btn["7"].Connect("clicked", func() { this.handle_num(7) })
	btn["8"].Connect("clicked", func() { this.handle_num(8) })
	btn["9"].Connect("clicked", func() { this.handle_num(9) })
	btn["10"].Connect("clicked", func() { this.handle_num(10) })
	btn["11"].Connect("clicked", func() { this.handle_num(11) })
	btn["12"].Connect("clicked", func() { this.handle_num(12) })
	btn["13"].Connect("clicked", func() { this.handle_num(13) })
	btn["14"].Connect("clicked", func() { this.handle_num(14) })
	btn["15"].Connect("clicked", func() { this.handle_num(15) })

	// Operator Buttons
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
	btn["ADD"].Connect("clicked", func() { this.handle_arith(OP_ADD) })
	btn["SUB"].Connect("clicked", func() { this.handle_arith(OP_SUB) })
	btn["MUL"].Connect("clicked", func() { this.handle_arith(OP_MUL) })
	btn["DIV"].Connect("clicked", func() { this.handle_arith(OP_DIV) })
	btn["MOD"].Connect("clicked", func() { this.handle_arith(OP_MOD) })
	btn["AND"].Connect("clicked", func() { this.handle_bit(OP_AND) })
	btn["OR"].Connect("clicked", func() { this.handle_bit(OP_OR) })
	btn["XOR"].Connect("clicked", func() { this.handle_bit(OP_XOR) })
	btn["NOT"].Connect("clicked", func() { this.handle_bit(OP_NOT) })
	btn["LSHFT"].Connect("clicked", func() { this.handle_shft(OP_LSHFT) })
	btn["RSHFT"].Connect("clicked", func() { this.handle_shft(OP_RSHFT) })

}

// Switch the radix format.
//  	Supports only 3.
//  	16(hex), 10(dec), 8(oct)
func (this *UI) switch_format(_rdx int) {
	//prev := this.mode
	switch _rdx {
	case 8:
		this.mode = CVT_OCT
		fmt.Println("Radix : ", this.mode)
	case 10:
		this.mode = CVT_DEC
		fmt.Println("Radix : ", this.mode)
	case 16:
		this.mode = CVT_HEX
		fmt.Println("Radix : ", this.mode)
	default:
		break
	}

	// convert all strings based on the mode
	// this.mode
	// this.set_prev()	this.pre .toOct()
	// this.set_left()	this.lhs .toHex()
	// this.set_right()	this.rhs .toDec()
}

// Handler for Number Buttons
//  	0(0x0) ~ 15(0xF)
func (this *UI) handle_num(_val int) {

	if this.csr == false {
		// Writing Left hand side
		res := this.lhs * this.mode
		res += _val
		this.lhs = res
		this.set_left(strconv.Itoa(this.lhs))
	} else {
		// Writing Right hand side
		res := this.rhs * this.mode
		res += _val
		this.rhs = res
		this.set_right(strconv.Itoa(this.rhs))
	}
	fmt.Println("Handle : Num : ", _val)
}

// Handler for Arithmetic operation
//  	ADD, SUB, MUL, DIV, MOD
func (this *UI) handle_arith(_code int) {
	if this.csr == false {
		// move focus from left to right
		this.csr = true
	} else {
		// handle the code :
		//  	Calculate the inner value
		switch _code {
		case OP_ADD:
			this.prev = this.lhs + this.rhs
			fmt.Println("Handle : ADD")
		case OP_SUB:
			this.prev = this.lhs - this.rhs
			fmt.Println("Handle : SUB")
		case OP_MUL:
			this.prev = this.lhs * this.rhs
			fmt.Println("Handle : MUL")
		case OP_DIV:
			this.prev = this.lhs / this.rhs
			fmt.Println("Handle : DIV")
		case OP_MOD:
			this.prev = this.lhs % this.rhs
			fmt.Println("Handle : MOD")
		default:
			return
		}

		// Display the result
		this.Lbl_prev.SetLabel(strconv.Itoa(this.prev))
		// clear the operands
		this.Lbl_lhs.SetLabel("")
		this.lhs = 0
		this.Lbl_rhs.SetLabel("")
		this.rhs = 0
		// Reset the cursor
		this.csr = false
	}
}

// Handler for Bitwise operation
//  	AND, OR, XOR, NOT
func (this *UI) handle_bit(_code int) {
	// move focus from left to right
	if this.csr != true {
		this.csr = true
	}
	else{
		// handle the code
		switch _code {
		case OP_AND:
			fmt.Println("Handle : AND")
		case OP_OR:
			fmt.Println("Handle : OR")
		case OP_XOR:
			fmt.Println("Handle : XOR")
		case OP_NOT:
			fmt.Println("Handle : NOT")
		default:
			return
	}
	}
}

// Handler for Shift operation
//  	Left Shift, Right Shift
func (this *UI) handle_shft(_code int) {
	switch _code {
	case OP_LSHFT:
		fmt.Println("Handle : L-shift")
	case OP_RSHFT:
		fmt.Println("Handle : R-shift")
	default:
		return
	}
}

// Change the Label of previous result
func (this *UI) set_prev(_str string) {
	this.Lbl_prev.SetLabel(_str)
}

// Change the Label of left operand
func (this *UI) set_left(_str string) {
	this.Lbl_lhs.SetLabel(_str)
}

// Change the Label of right operand
func (this *UI) set_right(_str string) {
	this.Lbl_rhs.SetLabel(_str)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
// Type : Operand and Operator
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

// Operand Redefinition
// type Operand int

// Operation Codes
const (
	OP_ADD = iota
	OP_SUB
	OP_MUL
	OP_DIV
	OP_MOD
	OP_AND
	OP_OR
	OP_XOR
	OP_NOT
	OP_LSHFT
	OP_RSHFT
)

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
// Utilities
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

// Default Panic Handelr
func catch() {
	s := recover()
	if s != nil {
		// print the content
		fmt.Println(s)
	}
}

// wrapper of gtk.NewButton() with label
func button(_lbl string) *gtk.Button {
	return gtk.NewButtonWithLabel(_lbl)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
// GUI exports
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

func Start() {
	defer catch() // Panic Handler

	// Initiate GTK
	gtk.Init(&os.Args)

	// Window Setup
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
	main_ui := new(UI)

	main_ui.Construct()
	defer main_ui.Destruct()

	fmt.Println("UI Objects construction finished.")
	fmt.Println("Starting the UI...")

	// Start the UI
	gtk.Main()
}

// Ready for event streaming
/*
func Setup_Events(_Window *gtk.Window) chan interface{} {
	defer catch()

	if _Window == nil {
		panic("init_Events() : nil window received")
	}

	// on Exit -> Quit the program
	_Window.Connect("destroy", gtk.MainQuit)

	ev_chan := make(chan interface{})

	_Window.Connect("key-press-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		ev_chan <- *(**gdk.EventKey)(unsafe.Pointer(&arg))
	})

	// Set the keyboard events
	_Window.SetEvents(int(gdk.BUTTON_PRESS_MASK))

	return ev_chan
}
*/

/*
	// Initialte evnets
	event := Setup_Events(win)

	go func() {
		for {
			e := <-event
			switch ev := e.(type) {
			case *gdk.EventKey:
				fmt.Println("key-press-event:", ev.Keyval)
				break
			}
		}
	}()
*/
