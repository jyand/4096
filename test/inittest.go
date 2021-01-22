package main
import (
        "fmt"
        //"os"
        //"bufio"
        "time"
        "math/rand"
)

const N = 4

func PlaceRandom(squares [N][N]uint16) [N][N]uint16 {
        rand.Seed(time.Now().UnixNano())
        ri := rand.Intn(N)
        rand.Seed(time.Now().UnixNano())
        rj := rand.Intn(N)
        if squares[ri][rj] == 0 {
                squares[ri][rj] = 2
                return squares
        }
        return PlaceRandom(squares)
}

func InitGame() [N][N]uint16{
        squares := [N][N]uint16{}
        for j := 0 ; j < N ; j++ {
                for i := 0 ; i < N ; i++ {
                        squares[i][j] = 0
                }
        }
        return PlaceRandom(squares)
}

func main() {
        a := InitGame()
        for j := 0 ; j < N ; j++ {
                for i := N-1 ; i > 0 ; i -= 1 {
                        if a[i][j] == 0 {
                                fmt.Printf("[ ]")
                        } else {
                                fmt.Printf("[%d]", a[i][j])
                        }
                }
                fmt.Print("\n")
        }
}
