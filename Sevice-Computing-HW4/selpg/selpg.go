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
	// check s/e page
	if(*s < 1 || *e < 1 || *e<*s){
		fmt.Println(os.Stderr, "input wrong s/e page ")
		return false
	}
	//check num of args
	if(argsnum < 3){
		fmt.Println(os.Stderr, "arguments not enough ")
		return false
	}
	// check if -f and -lNumber is both exit
	if(argsnum >= 5 && *f == true && *l != 72 ){
		fmt.Println(os.Stderr, "shouldn't use -f and -l at the same time")
		return false
	}
	if(flag.NArg()>1){
		fmt.Println(os.Stderr, "Input to much arguments")
		return false
	}
	filename = flag.Arg(0)
	return true
}

func selpg() {
	//set cmd
	var cmd *exec.Cmd
	var cmd_in io.WriteCloser
	var cmd_out io.ReadCloser
	//setting -d pipe
	if *d != "" {
		cmd = exec.Command("bash", "-c", *d)
		cmd_in, _ = cmd.StdinPipe()
		cmd_out, _ = cmd.StdoutPipe()
		cmd.Start()
	}
	//judge filename
	if filename != "" {
		inf, err := os.OpenFile(filename,os.O_RDWR,0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		line_count := 1
		page_count := 1
		fin := bufio.NewReader(inf)
		for {
			//read from file
			line, _, err := fin.ReadLine()
			if err != io.EOF && err != nil {
				fmt.Println(err)
				return
			}
			if err == io.EOF {
				break
			}
			if page_count >= *s && page_count <= *e{
				if *d == "" {
					//stdout
					fmt.Println(string(line))
				} else {
					//write to file
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
			cmd.Wait()
		}
	} else { ////stdin scanner with no filename
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