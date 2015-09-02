# go-shebang

Go-shebang is a runner for Go source files like a regular shell sript.

  - runs Go source files like a regular shell script files (bash, etc..)
  - works well with script args
  - simple and secure


> For a long time I was looking for a simple way to create a comfortable scripts, and finally I solved this problem in one evening.

### Version
1.0

### Installation

Go-shebang has a simple installation:

```sh
$ go get github.com/amaxcz/go-shebang
$ sudo ln -sfn $GOPATH/bin/go-shebang /usr/local/bin/go-shebang
$ /usr/local/bin/go-shebang
```

### How to use

You need to create new test file with content:
```
    #!/usr/local/bin/go-shebang
    package main
    
    import "fmt"
    
    func main() {
        fmt.Println("Hello World.")
    }
```
Next, fix permissions run:
```sh
$ chmod +x testfile
```

And run it, with:
```sh
$ ./testfile
```
You should see:
```sh
$ ./go-shebang testfile
Hello World.
```

Enjoy!

