// main_layout.go
package main

import (
	"fmt"
	"os"
	"unsafe"

	gdk "github.com/mattn/go-gtk/gdk"
	glib "github.com/mattn/go-gtk/glib"
	gtk "github.com/mattn/go-gtk/gtk"
)

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
func main() {
	defer catch() // catch Panic

	Start()
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
// Utilities
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

// Basic Panic Handelr
func catch() {
	s := recover()
	if s != nil {
		// print the content
		fmt.Println(s)
	}
}

// wrapper of button with new label
func button(_lbl string) *gtk.Button {
	return gtk.NewButtonWithLabel(_lbl)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
// Initializers
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

// Program Menu
func init_Menu(_MenuBar *gtk.MenuBar) {
	if _MenuBar == nil {
		panic("init_Menu() : nil received")
	}
	// Vertical Box for menu
	box_menu := gtk.NewVBox(false, 1)
	// MenuBar - menu
	mb_menu := gtk.NewMenuBar()
	box_menu.PackStart(mb_menu, false, false, 0)

	// Menu Items

	// [File]
	mi_file := gtk.NewMenuItemWithMnemonic("_File2")
	mb_menu.Append(mi_file)
	// Submenu for [File]
	subm_file := gtk.NewMenu()
	mi_file.SetSubmenu(subm_file)

	mi_exit := gtk.NewMenuItemWithMnemonic("_Exit2")
	mb_menu.Append(mi_exit)

	mi_exit.Connect("activate", func() {
		gtk.MainQuit()
	})

	// Add the menubox
	//win.Add(box_menu)
}

// Calculation Frame
func init_Calc(_Frame *gtk.Frame) {
	if _Frame == nil {
		panic("init_Calc() : nil received")
	}

	// (inner) Box of Calculation
	fm_calc_box := gtk.NewHBox(false, 1)
	_Frame.Add(fm_calc_box)

	// Box for Radix Buttons.
	box_rdx := gtk.NewVBox(false, 1)
	btn_hex := button("Hex") // [Hex] : Hexadecimal
	btn_dec := button("Dec") // [Dec] : Decimal
	btn_oct := button("Oct") // [Oct] : Octal
	box_rdx.Add(btn_hex)
	box_rdx.Add(btn_dec)
	box_rdx.Add(btn_oct)

	// Box for Result Labels
	box_labels := gtk.NewVBox(false, 1)
	lbl_prev := gtk.NewLabel("Previous Result") // Previous Calculation
	lbl_late := gtk.NewLabel("Current Result")  // Latest Calculaltion
	box_labels.Add(lbl_prev)
	box_labels.Add(lbl_late)

	// Add both Boxes (Radix & Result) to frame box
	fm_calc_box.Add(box_rdx)
	fm_calc_box.Add(box_labels)
}

// Number Button Frame
func init_Nums(_Frame *gtk.Frame) {
	if _Frame == nil {
		panic("init_Nums() : nil received")
	}

	// (inner) Box of Numbers
	fm_nums_box := gtk.NewVBox(false, 1)
	// Add to given frame
	_Frame.Add(fm_nums_box)

	// Table Initialization
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
	tbl_nums := gtk.NewTable(5, 4, false)

	// Jagged slice of buttons
	// nums := [][]*gtk.Button{}

	// Button for Number
	num := [17]*gtk.Button{
		// 0~7 : Oct
		button("0"), button("1"), button("2"), button("3"),
		button("4"), button("5"), button("6"), button("7"),
		// 0~9 : Dec
		button("8"), button("9"),
		// 0~F : Hex
		button("A"), button("B"), button("C"),
		button("D"), button("E"), button("F"),
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
}

// Operator Frame
func init_Oper(_Frame *gtk.Frame) {
	if _Frame == nil {
		panic("init_Oper() : nil received")
	}

	// (inner) Box of Operations
	fm_oper_box := gtk.NewVBox(false, 1)
	_Frame.Add(fm_oper_box)

	tbl_opers := gtk.NewTable(5, 3, false)

	// Operation Buttons
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
	// Button for Number
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

	fm_oper_box.Add(tbl_opers)

}

// Ready for event streaming
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

func Setup_UI(_Window *gtk.Window) {
	if _Window == nil {
		panic("init_UI() : nil received ")
	}

	box_win := gtk.NewVBox(Homogeneous, Default_Spacing)

	// Menu Bar
	//  	File :
	//  	View :
	//  	Help :
	// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
	//mb_Menu := gtk.NewMenuBar()

	// Frame - Calculation
	// This frame contains radix(16,10,8) and result labels
	// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
	fm_calc := gtk.NewFrame("Calculation")
	init_Calc(fm_calc)

	// Frame - Numbers
	// This frame contains number buttons for calculation
	//  	Hexadecimal	: 0 ~ 9, A ~ F
	//  	Decimal   	: 0 ~ 9
	//  	Octal     	: 0 ~ 7
	// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
	fm_nums := gtk.NewFrame("Numbers")
	init_Nums(fm_nums)

	// Frame - Operations
	// This frame contains operations.
	//  	ADD, SUB, MUL, DIV, MOD
	//  	AND, OR, XOR, NOT
	//  	LSHFT, RSHFT
	// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
	fm_oper := gtk.NewFrame("Operations")
	init_Oper(fm_oper)

	// Frame Positionings
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
	vpan1 := gtk.NewVPaned()
	vpan1.Pack1(fm_calc, No_Resize, No_Shrink)

	hpan1 := gtk.NewHPaned()
	hpan1.Pack1(fm_nums, No_Resize, No_Shrink)
	hpan1.Pack2(fm_oper, No_Resize, No_Shrink)

	vpan1.Pack2(hpan1, No_Resize, No_Shrink)

	box_win.Add(vpan1)

	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
	_Window.Add(box_win)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
// GUI exports
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

func Start() {
	defer catch() // Panic Handler

	// Initiate GTK
	gtk.Init(&os.Args)

	// Window Setup
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

	// Create the Main Window
	// Set title & size
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("0x_Calc")

	// on Exit -> Quit the program
	win.Connect("destroy", gtk.MainQuit)

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

	// Initialize UI
	Setup_UI(win)

	win.SetSizeRequest(UI_Width, UI_Height)
	win.ShowAll()

	// Start the UI
	gtk.Main()
}
