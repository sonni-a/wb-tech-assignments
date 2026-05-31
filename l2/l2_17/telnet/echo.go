package main
import ("fmt"; "io"; "net")
func main() {
    l, _ := net.Listen("tcp", "127.0.0.1:0")
    fmt.Println("echo on", l.Addr())
    c, _ := l.Accept()
    io.Copy(c, c)
}
