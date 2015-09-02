# go-shebang

Go-shebang is runner for Go source files like a regular shell sript.

Installation: go get github.com/amaxcz/go-shebang

It's will be installed to $GOPATH/bin/

sudo ln -sfn $GOPATH/bin/go-shebang /usr/local/bin/go-shebang





Example code. Put in to any file, chmod +x file, run it.

<pre>
	#!/usr/local/bin/go-shebang
	package main
	
	import "fmt"
	
	func main() {
	    fmt.Println("Hello World.")
	}
</pre>
