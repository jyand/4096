package main
import (
        "fmt"
        //"os"
        //"bufio"
        "math"
)

const N = 4

func PlaceRandom(squares [N][N]uint16) [N][N]uint16 {
        rand.Seed(time.Now().UnixNano())
        ri := rand.Intn(N)
        rand.Seed(time.Now().UnixNano())
        rj := rand.Intn(N)
        if squares[ri][rj] != 0 {
                PlaceRandom(squares)
        } else {
                squares[ri][rj] = 2
        }
        return squares
}

func InitGame() [N][N]uint16{
        squares := [N][N]uint16{}
        for j := 0 ; j < N ; j++ {
                for i := 0 ; i < N ; i++ {
                        squares[i][j] = 0
                }
        }
        return squares
}

// transforms the game's state each turn
func Move(prev [N][N]uint16, m int, n int) [N][N]uint16 {
}

/*polymorphic wrapper for Move()
func Directional(s string) [N][N]uint16 {
        switch s {
        default: return Move()
        case "j": return Move()
        case "k": return Move()
        case "l": return Move()
        }
}*/

func main() {
}
