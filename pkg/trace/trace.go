package trace

import (
	"runtime"
	"strings"
	"regexp"
	"fmt"
)

//
// Return the caller function name
// 0 -> returns the current caller function
// 1 -> returns the current caller parent
// etc.
//
var rx, _ = regexp.Compile("\\.func\\d+$")
func GetCallerFunctionName(backLevel int) string {

	var pc, tryAgain = make([]uintptr, backLevel  + 5), true
	runtime.Callers(backLevel, pc)
	var fn *runtime.Func

	for i:=0; i < len(pc); i++ {
		fmt.Printf("lvl=%d, fn=%s, size=%d\n", backLevel, runtime.FuncForPC(pc[i]).Name(), len(pc))
	}
	for i:=0; i < len(pc) && tryAgain; i++ {
		fn = runtime.FuncForPC(pc[i])
		tryAgain = rx.MatchString(fn.Name())
	}
	if index := strings.LastIndex(fn.Name(), "."); index != -1 {
		return fn.Name()[index + 1:]
	}
	return ""
}