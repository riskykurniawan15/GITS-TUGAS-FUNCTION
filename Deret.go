package main

import (
 "fmt"
 "strings"
 "strconv"
)

func main() {
	// Risky Kurniawan - ARS University
	var jml int
	fmt.Println("===========================================")
	fmt.Println("             DERET REKURSIF")
	fmt.Println("===========================================")
	fmt.Print("Masukan jumlah deret : ")
	fmt.Scan(&jml)
	fmt.Println("===========================================")
	fmt.Println("1. ", getDeret(jml, 1, 6))
	fmt.Println("2. ", getDeret(jml, 5, 10))
	fmt.Println("3. ", getDeret(jml, 1, 7))
	fmt.Println("4. ", getDeret_z(jml, 1, 1))
	fmt.Println("5. ", getDeret_z(jml, 5, 0))
	fmt.Println("===========================================")
}

func runDeret(n int, jml int, awal int, kelipatan int, sum int, temp string) string {
	temp = fmt.Sprintf("%d", awal) 

	if n != jml {
		temp += fmt.Sprintf(" + ")
	}

	if n == jml {
		return fmt.Sprintf("%s = %d", temp, sum)
	}
	sum += awal*kelipatan
	temp += runDeret(n+1, jml, awal*kelipatan, kelipatan, sum, temp)

	return temp
}

func getDeret(jml int, awal int, kelipatan int) string {
	var deret string = runDeret(1, jml, awal, kelipatan, awal, "")
	result := strings.Split(deret, " = ")
	return result[1] + " = " + result[0]
}

func runDeret_z(n int, jml int, awal int, kelipatan int, sum int, zero string, temp string) string {
	var nilai int
	nilai, _ = strconv.Atoi(fmt.Sprintf("%d%s", awal, zero))
	sum += nilai
	zero += "0" 
	temp = fmt.Sprintf("%d", nilai) 

	if n != jml {
		temp += fmt.Sprintf(" + ")
	}

	if n == jml {
		return fmt.Sprintf("%s = %d", temp, sum)
	}
	
	temp += runDeret_z(n+1, jml, awal+kelipatan, kelipatan, sum, zero, temp)

	return temp
}

func getDeret_z(jml int, awal int, kelipatan int) string {	
	var deret string = runDeret_z(1, jml, awal, kelipatan, 0, "", "")
	result := strings.Split(deret, " = ")
	return result[1] + " = " + result[0]
}