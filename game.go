package main
import (
        "fmt"
        //"os"
        //"bufio"
        "math"
)

var uint8 N = 4

func PlaceRandom(squares [N][N]uint16) {
        rand.Seed(time.Now().UnixNano())
        ri := rand.Intn(N)
        rand.Seed(time.Now().UnixNano())
        rj := rand.Intn(N)
        if squares[ri][rj] != 0 {
                PlaceRandom()
        }
        squares[ri][rj] = 2
}

func InitGame(uint8 N) [N][N]uint16{
        var squares [N][N] uint16
        for j := 0 ; j < N ; j++ {
                for i := 0 ; i < N ; i++ {
                        squares[i][j] = 0
                }
        }
        PlaceRandom()
        PlaceRandom()
        return squares
}

// transforms the game's state each turn
func Move(prev [N][N]uint16, k uint8, n uint8) [N][N]uint16 {
}

//polymorphic wrapper for Move()
func Directional(s string) [N][N]uint16 {
        switch s {
        default: return Move()
        case "j": return Move()
        case "k": return Move()
        case "l": return Move()
        }
}

func main() {

        var n float64 = 4
        if math.Exp2(n) == math.Pow(2, n) {
                fmt.Println("ya")
        }
}
