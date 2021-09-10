package main

import (
	"fmt"
	"sort"
	"strings"
)

var ID_CS = make(OrderList, len(NAME_CS))
var NAME_CS = map[int]string{}
var HOURS_CS = map[int]int{}
var PRICE_CS = map[int]int{}

func main(){
	// Risky Kurniawan - ARS University
	MenuUtama:
	var menu, submenu strings
	fmt.Println("==========================================")
	fmt.Println("          APLIKASI WARNET")
	fmt.Println("==========================================")
	fmt.Println("Silahkan pilih menu dibawah ini : ")
	fmt.Println("1. Tampil Data")
	fmt.Println("2. Input Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Rata - Rata Jumlah Jam Customers")
	fmt.Println("5. Jam penggunaa paling sedikit (3 Down)")
	fmt.Println("6. Pengguna dibawah jam rata - rata")
	fmt.Println("0. Keluar Aplikasi")
	fmt.Println("==========================================")
	fmt.Print("Pilihan anda adalah : ")
	fmt.Scan(&menu)
	fmt.Println()

	Pil:
	if menu == "0" {
		fmt.Println("==========================================")
		fmt.Println("           ** Terimakasih **")
		fmt.Println("    Risky Kurniawan - ARS University")
		fmt.Println("==========================================")
	}else if menu == "1" {
		SelectData()
	}else if menu == "2" {
		InputData()
	}else if menu == "3" {
		DeleteData()
	}else if menu == "4" {
		AverageCS()
	}else if menu == "5" {
		Top3Down()
	}else if menu == "6" {
		DownAverageCS()
	}else{
		fmt.Println("==========================================")
		fmt.Println("          Menu Tidak Tersedia")
		fmt.Println("      Silahkan pilih menu kembali")
		fmt.Println("==========================================")
	}

	PilSub:

	if menu !="0" {
		fmt.Println("9. Kembali ke menu")
		fmt.Println("0. Keluar Aplikasi")
		fmt.Print("Pilihan anda adalah : ")
		fmt.Scan(&submenu)

		if submenu == "9" {
			goto MenuUtama
		}else if submenu == "0" {
			menu = "0"
			goto Pil
		}else{
			fmt.Println("==========================================")
			fmt.Println("          Menu Tidak Tersedia")
			fmt.Println("      Silahkan pilih menu kembali")
			fmt.Println("==========================================")
			goto PilSub
		}
	}
}

type Order struct {
	Key   int
	Value int
}
type OrderList []Order
func (p OrderList) Len() int           { return len(p) }
func (p OrderList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p OrderList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func Ordering(){
	ID_CS = make(OrderList, len(NAME_CS))

	i := 0
	for k, _ := range NAME_CS {
		ID_CS[i] = Order{k, k}
		i++
	}
	
	sort.Sort(ID_CS)
}

func SelectData(){
	fmt.Println("==========================================")
	fmt.Println("              TAMPIL DATA")
	fmt.Println("==========================================")	
	for _, ID := range ID_CS {
		fmt.Printf("ID:%d -> [%s - %d Jam] Total(%s) \n", ID.Key, NAME_CS[ID.Key], HOURS_CS[ID.Key], currency(float64(PRICE_CS[ID.Key])))
	}	
	if len(NAME_CS) == 0 {
		fmt.Println("          NO RECORD DATA")
	}
	fmt.Println("==========================================")	
	fmt.Println("Jumlah Record =", len(NAME_CS))	
	fmt.Println("==========================================")	
}

func InputData(){
	var ID, hours int
	var name string
	var notif, err string = "",""
	fmt.Println("==========================================")
	fmt.Println("               INPUT DATA")
	ID_input:
	fmt.Println("==========================================")	
	fmt.Print("Masukan Kode ID (number) : ")
	fmt.Scan(&ID)
	if NAME_CS[ID] != "" {
		fmt.Println("Kode ID telah digunakan silahkan input kembali")
		goto ID_input
	}
	fmt.Print("Masukan Nama Pelanggan   : ")
	fmt.Scan(&name)
	fmt.Print("Masukan Jumlah Jam       : ")
	fmt.Scan(&hours)
	fmt.Println("==========================================")	

	_ , err = validation(name, "name")
	notif += err

	if notif == "" {
		NAME_CS[ID] = name
		HOURS_CS[ID] = hours
		PRICE_CS[ID] = (hours*60)*1000
		Ordering()
		fmt.Println("            INPUT BERHASIL")	
	}else{
		fmt.Println("NOTIF: ", notif)	
		fmt.Println("==========================================")	
		fmt.Println("            INPUT GAGAL")	
	}

	fmt.Println("==========================================")	
}

func validation(value string, name string) (string, string){
	if value == "" {
		return value, fmt.Sprintf("required %s;", name)
	}else{
		return "",""
	}
}

func DeleteData() {
	var id int
	fmt.Println("==========================================")
	fmt.Println("               HAPUS DATA")
	fmt.Println("==========================================")
	fmt.Print("Masukan ID Data : ")
	fmt.Scan(&id)
	fmt.Println("==========================================")
	if NAME_CS[id] != "" {
		delete(NAME_CS, id)
		delete(HOURS_CS, id)
		delete(PRICE_CS, id)
		Ordering()
		fmt.Println("        DATA BERHASIL DIHAPUS")
	}else{
		fmt.Println("        ID Tidak Ditemukan")
	}
	fmt.Println("==========================================")
}

func AVG() (float64, float64){
	var sum_hours int = 0
	for _, ID := range ID_CS {
	  sum_hours += HOURS_CS[ID.Key]
	}	

	return float64(sum_hours), (float64(sum_hours) / float64(len(NAME_CS)))
}

func AverageCS(){
	var sum_hours, avgvalue float64 = AVG()
	fmt.Println("==========================================")
	fmt.Println("        JUMLAH JAM RATA - RATA")
	fmt.Println("==========================================")	
	fmt.Println("Jumlah Record =", len(NAME_CS))	
	fmt.Printf("Jumlah Jam    = %.0f Jam \n", sum_hours)	
	fmt.Printf("Rata - Rata   = %.2f Jam \n", avgvalue)	
	fmt.Println("==========================================")	
}

func DownAverageCS(){
	var _, avgvalue float64 = AVG()
	fmt.Println("==========================================")
	fmt.Println("   TAMPIL DATA CS DIBAWAH RATA - RATA")
	fmt.Println("==========================================")	
	fmt.Printf("Rata - Rata   = %.2f Jam \n", avgvalue)	
	fmt.Println("==========================================")	
	var exist bool = false
	for _, ID := range ID_CS {
		if float64(HOURS_CS[ID.Key]) < avgvalue {
			exist = true
			fmt.Printf("ID:%d -> [%s - %d Jam] Total(%s) \n", ID.Key, NAME_CS[ID.Key], HOURS_CS[ID.Key], currency(float64(PRICE_CS[ID.Key])))
		}
	}	
	if exist == false {
		fmt.Println("          NO RECORD DATA")
	}
	fmt.Println("==========================================")	
}

func Top3Down(){
	H_CS := make(OrderList, len(HOURS_CS))

	i := 0
	for k, v := range HOURS_CS {
		H_CS[i] = Order{k, v}
		i++
	}
	
	sort.Sort(H_CS)

	fmt.Println("==========================================")
	fmt.Println("   TAMPIL 3 JAM PENGGUNAAN TER RENDAH")
	fmt.Println("==========================================")	
	var counter = 0;
	for _, ID := range H_CS {
		fmt.Printf("ID:%d -> [%s - %d Jam] Total(%s) \n", ID.Key, NAME_CS[ID.Key], HOURS_CS[ID.Key], currency(float64(PRICE_CS[ID.Key])))
		counter++
		if counter==3 {
			break
		}
	}	
	if counter == 0 {
		fmt.Println("          NO RECORD DATA")
	}
	fmt.Println("==========================================")	
}

func currency(number float64) string{
	// Modifikasi Dari tugas sebelumnya
	var money, sim string
	money = "Rp. "
	
	sim = "."
	
	valnum := strings.Split(fmt.Sprintf("%.2f", number), ".")[0]
	var reverse string = ""

	var index int = 1
	
	for i:=(strings.Count(valnum, "")-2); i>=0; i--{
		if (index%3) == 1 && index != 1{
			reverse += sim
		}
		reverse += string(valnum[i])
		
		index++
	}

	valnum = reverse

	for i:=(strings.Count(valnum, "")-2); i>=0; i--{
		money += string(valnum[i])
	}

	money += "," + strings.Split(fmt.Sprintf("%.2f", number), ".")[1]
	
	return money
}