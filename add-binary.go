package main

import "fmt"

func addBinary(a string, b string) string {
	var res []rune
	ar := []rune(a)
	br := []rune(b)
	ai := len(ar)-1
	bi := len(br)-1
	var carry rune = '0'
	for ai>=0 || bi >=0 || carry == '1'{
		var ac rune = '0'
		var bc rune = '0'
		if ai>=0 {
			ac = ar[ai]
			ai--
		}
		if bi>=0 {
			bc = br[bi]
			bi--
		}
		r,c := add(ac,bc,carry)
		carry = c
		res = append([]rune{r},res...)
	}

	return string(res)
}

func add(a rune,b rune, carry rune) (rune, rune)  {
	r := a-'0'+b-'0'+carry-'0'
	r1 := r%2+'0'
	c := r/2+'0'
	return r1, c
}

func main() {
	ss := addBinary("11","1")
	fmt.Println(ss)

}