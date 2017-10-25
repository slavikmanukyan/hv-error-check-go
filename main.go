package main

import "fmt"

func getVerticalCode(s string) int32 {
	r := []rune(s)
	v := rune(0)
	for _, char := range r {
		v ^= char
	}
	return v
}

func errorCheck(num int) byte {
	mask := 0x80000000
	count := byte(0)
	for mask > 0 {
		check := num & mask
		if check == mask {
			count ^= 1
		}
		mask >>= 1
	}
	return count
}

func getHorizontalCode(s string) []byte {
	r := []rune(s)
	h := make([]byte, len(s))
	for index, char := range r {
		h[index] = errorCheck(int(char))
	}
	return h
}

func main() {
	s := "ԱԲԳԴԵԶԷԸԹԺ ԻԼԽ"
	chars := []rune(s)
	v := getVerticalCode(s)
	h := getHorizontalCode(s)

	for index, char := range chars {
		binChars := fmt.Sprintf("%016b", char)
		for _, b := range binChars {
			fmt.Printf(" %c ", b)
		}
		fmt.Printf("| %d\n", h[index])
	}
	for i := 1; i <= 16; i++ {
		fmt.Print("---")
	}
	fmt.Println()
	binChars := fmt.Sprintf("%016b", v)
	for _, b := range binChars {
		fmt.Printf(" %c ", b)
	}
}

/* Result:

0  0  0  0  0  1  0  1  0  0  1  1  0  0  0  1 | 1
0  0  0  0  0  1  0  1  0  0  1  1  0  0  1  0 | 1
0  0  0  0  0  1  0  1  0  0  1  1  0  0  1  1 | 0
0  0  0  0  0  1  0  1  0  0  1  1  0  1  0  0 | 1
0  0  0  0  0  1  0  1  0  0  1  1  0  1  0  1 | 0
0  0  0  0  0  1  0  1  0  0  1  1  0  1  1  0 | 0
0  0  0  0  0  1  0  1  0  0  1  1  0  1  1  1 | 1
0  0  0  0  0  1  0  1  0  0  1  1  1  0  0  0 | 1
0  0  0  0  0  1  0  1  0  0  1  1  1  0  0  1 | 0
0  0  0  0  0  1  0  1  0  0  1  1  1  0  1  0 | 0
0  0  0  0  0  0  0  0  0  0  1  0  0  0  0  0 | 1
0  0  0  0  0  1  0  1  0  0  1  1  1  0  1  1 | 1
0  0  0  0  0  1  0  1  0  0  1  1  1  1  0  0 | 0
0  0  0  0  0  1  0  1  0  0  1  1  1  1  0  1 | 1
------------------------------------------------
0  0  0  0  0  1  0  1  0  0  0  1  0  0  0  1

*/
