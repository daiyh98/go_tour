package main

import (
	"github.com/go-programming-tour-book/tour/cmd"
	"log"
)

//type Name string //定制化命令行参数类型，需要实现flag.Parse()源码中Value相关的两个接口
//
//func (i *Name) String() string {
//	return fmt.Sprintf(string(*i))
//}
//
//func (i *Name) Set(value string) error {
//	if len(*i) > 0 {
//		return errors.New("name flag already set")
//	}
//
//	*i = Name("daiyh:" + value)
//	return nil
//}

//func main() {
//	//var name string
//	var name Name
//	flag.Var(&name, "name", "帮助信息")
//	//flag.StringVar(&name, "name", "Go语言编程之旅", "帮助信息")
//	//flag.StringVar(&name, "n", "Go语言编程之旅", "帮助信息")
//	flag.Parse() //解析并绑定命令行参数
//	//args := flag.Args() //将传入的参数赋值给args变量，如果没有上一步解析，那返回的是nil
//	//if len(args) <= 0 {
//	//	return
//	//}
//
//	//switch args[0] {
//	//case "go":
//	//	//flag.NewFlagSet方法会返回带有指定名称和错误处理属性的空命令集，相当于创建一个新的命令集去支持子命令
//	//	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
//	//	goCmd.StringVar(&name, "name", "Go语言", "帮助信息")//如果用-n设置name参数就会报错
//	//	_ = goCmd.Parse(args[1:])
//	//case "php":
//	//	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
//	//	phpCmd.StringVar(&name, "n", "PHP语言", "帮助信息")//如果用-name设置name参数就会报错
//	//	_ = phpCmd.Parse(args[1:])
//	//}
//
//	log.Printf("name: %s", name)
//}

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
