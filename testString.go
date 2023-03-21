package main
//汇编命令： GOOS=linux GOARCH=amd64 go tool compile -S testString.go
func test() string{
	str:="hello"
	return str
}
