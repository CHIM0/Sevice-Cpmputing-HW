package main

import (
	flag "github.com/spf13/pflag"
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"io"
)
//setting pflag
var s = flag.IntP("sNum","s",1,"Input start page.")
var e = flag.IntP("eNum", "e", 1,"Input end page")
var l = flag.IntP("lNum", "l", 72,"Input num of lines in each page")
var f = flag.BoolP("f", "f", false,"Use /f as flag for page")
var d=flag.StringP("d","d","", "input destination")
var argsnum int
var filename string

func args_check( ) bool{
	argsnum = argsnum-flag.NArg()
	if(*s < 1 || *e < 1 || *e<*s){
		fmt.Println(os.Stderr, "input wrong s/e page ")
		return false
	}
	if(argsnum < 3){ 	//check num of args
		fmt.Println(os.Stderr, "arguments not enough ")
		return false
	}
	if(argsnum >= 5 && *f == true && *l != 72 ){	// check if -f and -lNumber is both exit
		fmt.Println(os.Stderr, "shouldn't use -f and -l at the same time")
		return false
	}
	if(flag.NArg()>1){
		fmt.Println(os.Stderr, "Input too much arguments")
		return false
	}
	filename = flag.Arg(0)
	return true
}

func selpg() {
	cmd := exec.Command( *d)
	cmd_in, _ := cmd.StdinPipe()
	if *d != "" {
		cmd.Stdout = os.Stdout
		cmd.Start() //Wait()
	}
	//judge filename
	if filename != "" {
		input_file, err := os.OpenFile(filename,os.O_RDWR,0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		line_nums, page_nums:= 1, 1
		file_in := bufio.NewReader(input_file)
		for { //read line from file
			line, _, err := file_in.ReadLine()
			if err != io.EOF && err != nil {
				fmt.Println(err)
				return
			}else if err == io.EOF{
				break
			}
			if page_nums >= *s && page_nums <= *e{
				if *d == "" {
					fmt.Println(string(line))
				} else {
					fmt.Fprintln(cmd_in, string(line))//write to pipe
				}
			}
			line_nums++
			if (*f == false){
				if line_nums > *l {
					line_nums = 1
					page_nums++
				}
			} else {
				if string(line) == "\f" {
					page_nums++
				}
			}
		}
		if *d != "" {
			cmd_in.Close()
			cmd.Wait()
		}
	} else { ////stdin scanner with no filename
		std_input := bufio.NewScanner(os.Stdin)
		line_nums := 1
		page_nums := 1
		out := ""
		for std_input.Scan() {
			line := std_input.Text()
			line += "\n"
			if page_nums >= *s && page_nums <= *e{
				out += line
			}
			line_nums++
			if (*f == false){
				if line_nums> *l {
					line_nums = 1
					page_nums++
				}
			} else {
				if string(line) == "\f" {
					page_nums++
				}
			}
		}
		if *d == "" {
			fmt.Print(out)
		} else {
			fmt.Fprint(cmd_in, out)//write to pipe
			cmd_in.Close()
			cmd.Wait()
		}
	}
}
//main function
func main() {
	argsnum = len(os.Args)
	flag.Parse()
	if(args_check()){
		selpg()
	}
}