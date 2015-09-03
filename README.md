# go-shebang

Go-shebang is a runner for Go source files like a regular shell sript.

  - works with **shebang** like #!/bin/bash
  - runs Go lang source files like a regular shell script files (bash, etc..)
  - works well with script args $1 $2 etc ...
  - simple and secure


> For a long time I was looking for a simple way to create a comfortable scripts, and finally I solved this problem in one night.

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

### Debian/Ubuntu workaround
You can try one more solution for files with .go extension, only for Linux.

```sh
apt-get install binfmt-support

cat > /usr/local/bin/gorun << EOF
#!/bin/sh
/usr/bin/go run "\$@"
EOF

chmod 755 /usr/local/bin/gorun

update-binfmts --install go /usr/local/bin/gorun --extension go
```

### See also

  - [About wrappers] http://tldp.org/LDP/abs/html/wrapper.html
  - [gorun] https://wiki.ubuntu.com/gorun and http://bazaar.launchpad.net/~niemeyer/gorun/trunk/view/head:/gorun.go
  - [go lang cookbook] http://golangcookbook.com/chapters/running/shebang/


###Enjoy!

