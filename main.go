package main

import (
	"fmt"
	"rbac_manager/core"
	"rbac_manager/global"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	global.Conf = core.GetConfig()
	core.InitLogger("logs")
	global.Log.Info(fmt.Sprintf("%#v\n", global.Conf))
}
