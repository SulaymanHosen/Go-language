package main
import "fmt"

func main() {
    var a = 100
    fmt.Println("Outside block, a =", a)

    {
        var a = 200 // নতুন a শুধুমাত্র এই block-এ
        fmt.Println("Inside block, a =", a)
    }

    fmt.Println("After block, a =", a)
}
