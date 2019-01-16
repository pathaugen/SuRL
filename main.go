
package main

import (
  "fmt"
	//"net/http"
	//"strconv"
	//"strings"
	//"bufio"
	"os"
	"os/exec"
	//"time"
	//"log"
)


var debug = false // Set to true to debug the question 2 output


// ========== START: Golang Console Colors ========== ========== ========== ==========
// Golang Console Colors
// Example: fmt.Print( cRed + "HelloWorld" + cClr )
var cClr				= "\u001b[0m"

var cBold				= "\u001b[1m"

var cBlack			= "\u001b[30m"
var cRed				= "\u001b[31m"
var cGreen			= "\u001b[32m"
var cYellow			= "\u001b[33m"
var cBlue				= "\u001b[34m"
var cMagenta		= "\u001b[35m"
var cCyan				= "\u001b[36m"
var cWhite			= "\u001b[37m"

var cBlackBG		= "\u001b[40m"
var cRedBG			= "\u001b[41m"
var cGreenBG		= "\u001b[42m"
var cYellowBG		= "\u001b[43m"
var cBlueBG			= "\u001b[44m"
var cMagentaBG	= "\u001b[45m"
var cCyanBG			= "\u001b[46m"
var cWhiteBG		= "\u001b[47m"
// ========== END: Golang Console Colors ========== ========== ========== ==========


// ========== START: Console Splash ========== ========== ========== ==========
var appinfo = `
  ` + cBlue + `====================================================` + cBold + cCyan + `
  .d88888b            888888ba  dP
  88.    "'           88    `+"`"+`8b 88
  `+"`"+`Y88888b. dP    dP a88aaaa8P' 88
        `+"`"+`8b 88    88  88   `+"`"+`8b. 88
  d8'   .8P 88.  .88  88     88 88
   Y88888P  `+"`"+`88888P'  dP     dP 88888888P
  ooooooooooooooooooooooooooooooooooooooo` + cClr + `
  ` + cCyan + `Suricata Rule Loader` + cClr + `
  ` + cCyan + `https://github.com/` + cYellow + `pathaugen` + cCyan + `/SuRL` + cClr + `
  ` + cBlue + `====================================================` + cClr + `
`
// ========== END: Console Splash ========== ========== ========== ==========


func main() {

  // Clear Screen
  cmd := exec.Command("cmd", "/c", "cls")
  cmd.Stdout = os.Stdout
  cmd.Run()

  // Output Simplification
  breakspace := ("\n")
	breakline := ( breakspace + cBlue + "  ====================================================" + cClr + breakspace )

  fmt.Print( appinfo )
  fmt.Print( breakspace )




  // ========== START: Main Logic ========== ========== ========== ==========
  fmt.Print( "SuRL Output" )
  // ========== END: Main Logic ========== ========== ========== ==========




  fmt.Print( breakspace )
  fmt.Print( breakline )
}
