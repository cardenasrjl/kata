package main

import "fmt"

func main() {
	
	/*board := [9][9]int{
		{0,7,2,0,0,4,9,0,0},
		{3,0,4,0,8,9,1,0,0},
		{8,1,9,0,0,6,2,5,4},
		{7,0,1,0,0,0,0,9,5},
		{7,0,0,0,0,2,0,7,0},
		{0,0,0,8,0,7,0,1,2},
		{4,0,5,0,0,1,6,2,0},
		{2,3,7,0,0,0,5,0,1},
		{0,0,0,0,2,5,7,0,0},
		
	}*/
	board := [9][9]int{}
	
	newBoard, result := sodoku(board,0,0)
	fmt.Printf("%+v %t", newBoard, result)
}

func sodoku(board [9][9]int, row, col int) ([9][9]int,bool) {
	if len(board) <= col {
		return sodoku(board, row+1,0)
	}

	if len(board) <= row {
		return board,true
	}

	if board[row][col] > 0 {
		return sodoku(board, row,col+1)
	}
	
	//the value has ben set already
	for i:=0;i<len(board);i++ {
		if boardValid(board, i+1, row, col) {
			board[row][col] = i+1
			temp, result := sodoku(board,row,col+1)
			if result {
				return temp, result
			} else {
				board[row][col] = 0
			}
		}
	}
	
	return board,false
}

func boardValid(board [9][9]int, value, row, col int) bool {
	//check the row
	for iRow := 0;iRow<len(board);iRow++ {
		if board[iRow][col] == value {
			return false
		}
	} 
	
	//cehck the colum
	for iCol := 0;iCol<len(board);iCol++ {
		if board[row][iCol] == value {
			return false
		}
	}

	//check the square
	sqRow := getSquare(row)
	sqCol := getSquare(col)
	for iRow := sqRow;iRow < sqRow+3;iRow++ {
		for iCol := sqCol;iCol < sqCol+3;iCol++ {
			if board[iRow][iCol] == value {
				return false
			}
		}
	}
	
	return true
}

func getSquare(val int) int {

	result := 0
	if val >=3 && val <=5 {
		result =3
	}
	
	if val >=6 && val <=8 {
		result =6
	}
	return result
}
