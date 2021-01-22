package main
import (
        "fmt"
        //"os"
        //"bufio"
        "time"
        "math/rand"
)

const N = 4

// fix: doesn't always place
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
        board := [N][N]uint16{}
        for j := 0 ; j < N ; j++ {
                for i := 0 ; i < N ; i++ {
                        board[i][j] = 0
                }
        }
        return PlaceRandom(board)
}

func MoveLeft(board [N][N]uint16, j int) [N][N]uint16 {
        for i := 1 ; i < N ; i++ {
                if board[i][j] == 0 {
                        board[i-1][j] = board[i][j]
                        board[i][j] = 0
                }
                if board[i][j] == board[i-1][j] {
                        board[i-1][j] = 2*board[i][j]
                        board[i][j] = 0
                }
        }
        if j == 0 {
                return PlaceRandom(board)
        }
        return MoveLeft(board, j-1)
}

func PrintMatrix(a [N][N]uint16) {
        for j := 0 ; j < N ; j++ {
                for i := 0 ; i < N ; i ++ {
                        if a[i][j] == 0 {
                                fmt.Printf("[ ]")
                        } else {
                                fmt.Printf("[%d]", a[i][j])
                        }
                }
                fmt.Print("\n")
        }
}

func main() {
        game := PlaceRandom(PlaceRandom(InitGame()))
        PrintMatrix(game)
        fmt.Println()
        game = MoveLeft(game, N-1)
        PrintMatrix(game)
}
