This simple example works perfectly using a CLI, but if you try to execute from Goland you get:

```bash
GOROOT=/usr/local/go #gosetup
GOPATH=/home/xxx/go #gosetup
/usr/local/go/bin/go build -o /tmp/___go_build_main_go /home/xxx/Documents/src/hello-module/main.go #gosetup
# command-line-arguments
./main.go:11:14: undefined: hello
./main.go:12:14: undefined: Hello

Compilation finished with exit code 2
```
