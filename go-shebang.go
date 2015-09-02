//
//	go-shebang - Runner for Go source files.
//

/*

Copyright (c) 2015, Aleksey Maximov <amaxcz@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/


package main


import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"path"
	"io/ioutil"
	"os"
	"os/exec"
	"bufio"
)

func main() {
	var err error
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: " + path.Base(os.Args[0]) + " <sourcefile> [args]")
		os.Exit(1)
	}

	_, err = os.Stat(filepath.Join(runtime.GOROOT(), "bin", "go"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: can't access GO lang executable")
		os.Exit(1)
	}

	_, err = exec.LookPath("go")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: can't find GO lang tools in PATH")
		os.Exit(1)
	}

	err = Run(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: " + err.Error())
		os.Exit(1)
	}
}




func Run(args []string) (err error) {

	srcfile := args[0]

	srcfile, err = filepath.Abs(srcfile)
	if err != nil {
		return err
	}

	srcfile, err = filepath.EvalSymlinks(srcfile)
	if err != nil {
		return err
	}
	
	
	tmpdir := os.TempDir()
	if err != nil {
		return err
	}
	
	tmpdir, err = ioutil.TempDir(tmpdir, strconv.Itoa(os.Getpid()) + "-")
	if err != nil {
		return err
	}
	
	
	runfile := filepath.Join(tmpdir, path.Base(srcfile))	

	defer os.Remove(tmpdir)
	defer os.Remove(runfile)

	err = CompileFile(tmpdir, srcfile, runfile)
	if err != nil {
		return err
	}

	err = Exec(append([]string{runfile}, args[1:]...))
	
	return nil
}

func CompileFile(tmpdir, srcfile, runfile string) (err error) {
	tmpfile := path.Base(srcfile)
	dstfile := filepath.Join(tmpdir, tmpfile + ".go")	

	f, err := os.Open(srcfile)
	if err != nil {
	    return err
	}
	defer f.Close()
	
	r := bufio.NewReaderSize(f, 16)
	content, _, err := r.ReadLine()
	if err != nil {
	    return err
	}
	
	fmt.Println(os.Stderr,  content)
	
	if len(content) > 2 && content[0] == '#' && content[1] == '!' {
		content, err := ioutil.ReadFile(srcfile)
	    if err != nil {
	        return err
	    }
		content[0] = '/'
		content[1] = '/'
		ioutil.WriteFile(dstfile, content, 0600)
		defer os.Remove(dstfile)
	}

	n := GetArchMagic()
	cc := filepath.Join(tmpdir, tmpfile + "." +  n)
	ld := filepath.Join(tmpdir, tmpfile + "." + "out")

	gobin, err := exec.LookPath("go")
	if err != nil {
		return err
	}

	err = Exec([]string{gobin, "tool", n + "g", "-o", cc, dstfile})
	if err != nil {
		return err
	}
	defer os.Remove(cc)

	err = Exec([]string{gobin, "tool", n + "l", "-o", ld, cc})
	if err != nil {
		return err
	}
	
	err = os.Rename(ld, runfile)
	if err != nil {
		return err
	}
	
	return nil
}


func Exec(args []string) error {
	var cmd *exec.Cmd
	if len(args) > 1 {
		cmd = exec.Command(args[0], args[1:]...)
	} else {
		cmd = exec.Command(args[0], []string{""}...)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	if err != nil {
		fmt.Fprintln(os.Stderr,  err)
		return err
	}
	return nil
}



func GetArchMagic() string {
	switch runtime.GOARCH {
	case "arm":
		return "5"
	case "386":
		return "8"
	case "amd64":
		return "6"
	}
	panic("Fatal error: unsupported GOARCH: " + runtime.GOARCH)
}

