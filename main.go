package main

/*
go.mod itu manajemen direktori atau path dalam proyek
go mod init Golangs : modul modul (mod singkatan modules) pada program Golang ini, root path nya adalah Golangs
barulah dia bisa nyari direktori kalkulator lewat "Golangs/kalkulator"
kalo tanpa inisialisasi perintah itu, Golangs akan dianggap standard library bawaan. makanya karena Golangs bukan-
-standard library, compiler akan kasih tau kalo "package Golangs is not in std"
*/

import (
	"fmt"
	"Golangs/kalkulator"
)

func main(){
	angkuy := kalkulator.Selisih(11, 14)
	fmt.Println("Hasil penjumlahan adalah : ", angkuy)
	angka := 10
	angki := angka + 10
	for i := range 10 { // Looping dari 0 hingga 9
		if (angki+i)%2 == 0 {
			fmt.Println(angki+i, "adalah bilangan genap")
		} else {
			fmt.Println(angki+i, "adalah bilangan ganjil")
		}
	}
	fmt.Println("Hello, World!")
	fmt.Println(angki)
} 



