package main

import (
	flag "github.com/spf13/pflag"
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"io"
	"io/ioutil"
)

var s = flag.IntP("sNum","s",1,"Input start page.")
var e = flag.IntP("eNum", "e", 1,"Input end page")
var l = flag.IntP("lNum", "l", 72,"Input num of lines in each page")
var f = flag.BoolP("f", "f", false,"Use /f as flag for page")
var d=flag.StringP("d","d","", "input destination")
var argsnum int
var filename string

func args_check( ){
	argsnum = argsnum-flag.NArg()
	// check s/e page
	if(*s < 1 || *e < 1 || *e<*s){
		fmt.Println(os.Stderr, "input wrong s/e page ")
		os.Exit(1)
	}
	//check num of args
	if(argsnum < 3){
		fmt.Println(os.Stderr, "arguments not enough ")
		os.Exit(1)
	}
	// check if -f and -lNumber is both exit
	if(argsnum >= 5){
		fmt.Println(os.Stderr, "shouldn't use -f and -l at the same time")
		os.Exit(1)
	}
	if(flag.NArg()>1){
		fmt.Println(os.Stderr, "Input to much arguments")
	}
	filename = flag.Arg(0)
}
func selpg() {
	var cmd *exec.Cmd
	var cmd_in io.WriteCloser
	var cmd_out io.ReadCloser
	if *d != "" {
		cmd = exec.Command("bash", "-c", *d)
		cmd_in, _ = cmd.StdinPipe()
		cmd_out, _ = cmd.StdoutPipe()
		//执行设定的命令
		cmd.Start()
	}
	if filename != "" {
		inf, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		line_count := 1
		page_count := 1
		fin := bufio.NewReader(inf)
		for {
			//读取输入文件中的一行数据
			line, _, err := fin.ReadLine()
			if err != io.EOF && err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if err == io.EOF {
				break
			}
			if page_count >= *s && page_count <= *e{
				if *d == "" {
					//打印到屏幕
					fmt.Println(string(line))
				} else {
					//写入文件中
					fmt.Fprintln(cmd_in, string(line))
				}
			}
			line_count++
			if (*f == false){
				if line_count > *l {
					line_count = 1
					page_count++
				}
			} else {
				if string(line) == "\f" {
					page_count++
				}
			}
		}
		if *d != "" {
			cmd_in.Close()
			cmdBytes, err := ioutil.ReadAll(cmd_out)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print(string(cmdBytes))
			//等待command退出
			cmd.Wait()
		}
	} else {
		//从标准输入读取内容
		ns := bufio.NewScanner(os.Stdin)
		line_count := 1
		page_count := 1
		out := ""

		for ns.Scan() {
			line := ns.Text()
			line += "\n"
			if page_count >= *s && page_count <= *e{
				out += line
			}
			line_count++
			if (*f == false){
				if line_count > *l {
					line_count = 1
					page_count++
				}
			} else {
				if string(line) == "\f" {
					page_count++
				}
			}
		}
		if *d == "" {
			fmt.Print(out)
		} else {
			fmt.Fprint(cmd_in, out)
			cmd_in.Close()
			cmdBytes, err := ioutil.ReadAll(cmd_out)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print(string(cmdBytes))
			//等待command退出
			cmd.Wait()
		}
	}
}

func main() {
	argsnum = len(os.Args)
	flag.Parse()
	args_check()
	selpg()
}