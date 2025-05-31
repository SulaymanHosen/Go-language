package main
import "fmt"
var a=20
var b=90
func add(x int, y int)  {
	z:= x + y
	fmt.Println("The sum of", z)
}
func main() {
	p:= 10
q:=30
add(p,q) 
add(a,b) 
add(a,p)
add(b,q) 
add(a,q) 
add(p,b) 


}