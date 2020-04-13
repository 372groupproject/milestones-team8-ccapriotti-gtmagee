package main

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
)

func main(){
	var board [3][3]byte
	board = clearBoard(board)
	printBoard(board)
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	var playerOneTurn bool = true
	gameOver := false
	moveCount := 0;
	for !gameOver {
		var marker byte
		if playerOneTurn {
			fmt.Println("Player One, Make Your Move!")
			marker = 'X'
		}else{
			fmt.Println("Player Two, Make Your Move!")
			marker = 'O'
		}
		row,col := getCoords(reader, board)
		board[row][col] = marker
		printBoard(board)
		moveCount++
		if (vertWin(board) || horWin(board) || diagWin(board)){
			fmt.Printf("%c Has Won!\n", marker)
			gameOver = true
		} else if (moveCount == 9){
			fmt.Printf("ITS A DRAW...")
			gameOver = true
		}else{
			playerOneTurn = !playerOneTurn
		}
	}
}

func horWin(board [3][3]byte) bool{
	for row := 0; row < len(board); row++ {
		diff := false
		for col := 0; !diff && col < len(board[row]) - 1; col++ {
			if (board[row][col] != board[row][col + 1] || board[row][col] == '-'){
				diff = true
			}
		}
		if (!diff){
			return true
		}
	}
	return false
}

func vertWin(board [3][3]byte) bool{
	for col := 0; col < len(board[0]); col++ {
		diff := false
		for row := 0; !diff && row < len(board) - 1; row++ {
			if (board[row][col] != board[row + 1][col] || board[row][col] == '-'){
				diff = true
			}
		}
		if (!diff){
			return true
		}
	}
	return false
}

func diagWin(board [3][3]byte) bool{
	return ((board[0][0] == board[1][1] && board[1][1] == board[2][2]) || 
		   (board[0][2] == board[1][1] && board[1][1] == board[2][0])) &&
		   (board[1][1] != '-')
}

func printBoard(board [3][3]byte){
	fmt.Println()
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]) - 1; col++ {
			fmt.Printf("%c,", board[row][col])
		}
		fmt.Printf("%c\n", board[row][len(board[row]) - 1])
	}
	fmt.Println()
}

func clearBoard(board [3][3]byte) [3][3]byte{
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++{
			board[row][col] = '-'
		}
	}
	return board
}

func getCoords(reader *bufio.Reader, board [3][3]byte) (int, int){
    fmt.Println("Specify Row and Column '(row,col)' of Your Move From 0 to 2 : ")
    coordsRead := false
    var row int
    var col int
    var err1 error
    var err2 error
    for !coordsRead {
        inp, _ := reader.ReadString('\n')
        coords := strings.Split(strings.Trim(strings.TrimSpace(inp), "() "), ",")
        if len(coords) != 2{
        	coords = append(coords, "-1")
        	coords[1] = "-1"
        }
        row, err1 = strconv.Atoi(strings.TrimSpace(coords[0]))
        col, err2 = strconv.Atoi(strings.TrimSpace(coords[1]))
        if (err1 != nil || err2 != nil || row > 2 || row < 0 || col > 2 || col < 0) {
            fmt.Println("Bad Input.  Need Input of Form: '(row,col)'")
        }else if board[row][col] != '-'{
        	fmt.Println("Already A Piece There, Choose Somewhere Else...")
        }else{
            coordsRead = true
        }
    }
    return row, col
}