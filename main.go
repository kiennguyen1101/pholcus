package main

import (
	"github.com/henrylee2cn/pholcus/exec"
	_ "github.com/henrylee2cn/pholcus_lib" // This is a publicly maintained spider rule base
	// _ "pholcus_lib_pte" // You can also freely add your own rule base
)

func main() {
	// Set the runtime default action interface and start running
	// Before running the software, you can set the -a_ui parameter to "web", "gui" or "cmd" to specify the operation interface for this run.
	// where "gui" only supports Windows systems
	exec.DefaultRun("web")
}
