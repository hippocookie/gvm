package main

import (
	"fmt"
	"gvm/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%s Xjre:%s class:%s args:%v\n",
		cmd.cpOption, cmd.XjreOption, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("cloud not find or load main class %s, err:%s\n", cmd.class, err)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
