##go 版本间的改变
####go1.5
参考: [1.5改变](https://tonybai.com/2015/07/10/some-changes-in-go-1-5/)

1. go在1.5之后不再用c语言,编译器用go重写了，如果想要编译的go的程序的话，不能再直接运行 ./make_bash  
2. GOMAXPROCS初始值由1改为了运行环境的CPU核数。
3. 简化跨平台编译 1.5之前的版本要想实现跨平台编译，需要到$GOROOT/src下重新执行一遍make.bash，执行前设置好目标环境的环境变量(GOOS和 GOARCH)，Go 1.5大大简化这个过程，使得跨平台编译几乎与普通编译一样简单
