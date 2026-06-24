/*========================================SATU=======================================================*/
// package main

// /*
// go.mod itu manajemen direktori atau path dalam proyek
// go mod init Golangs : modul modul (mod singkatan modules) pada program Golang ini, root path nya adalah Golangs
// barulah dia bisa nyari direktori kalkulator lewat "Golangs/kalkulator"
// kalo tanpa inisialisasi perintah itu, Golangs akan dianggap standard library bawaan. makanya karena Golangs bukan-
// -standard library, compiler akan kasih tau kalo "package Golangs is not in std"
// */

// import (
// 	"fmt"
// 	"Golangs/kalkulator"
// )

// func main(){
// 	angkuy := kalkulator.Selisih(11, 14)
// 	fmt.Println("Hasil penjumlahan adalah : ", angkuy)
// 	angka := 10
// 	angki := angka + 10
// 	for i := range 10 { // Looping dari 0 hingga 9
// 		if (angki+i)%2 == 0 {
// 			fmt.Println(angki+i, "adalah bilangan genap")
// 		} else {
// 			fmt.Println(angki+i, "adalah bilangan ganjil")
// 		}
// 	}
// 	fmt.Println("Hello, World!")
// 	fmt.Println(angki)
// }
/*========================================SATU=======================================================*/
/*========================================DUA=======================================================*/
// package main

// import (
// 	"net/http" //ya internet gitu ya
// 	"github.com/gin-gonic/gin" //ini frameworknya
// )

// func main() {
// 	// 1. Inisialisasi framework Gin
// 	r := gin.Default() //ini router. variabel objek mesin utama dari Gin; 
// 					   // tugasnya mengatur semua traffic data yg masuk ke server db

// 	// 2. Membuat rute (endpoint) HTTP GET /halo
// 	//  
// 	r.GET("/halo", func(c *gin.Context){ // GET berarti user mau mengambil/membaca data aja.
// 		// "/halo" cuma nama untuk situs localhost nya. bisa diganti jadi /aloha dll
// 		// c itu context. sebuah kantong ajaib yang bisa membawa seluruh konteks yang user perlu. 
// 			// *gin.Context anggap aja ya c memberi/mendapat context dari gin
// 		//c.JSON() ---> View
// 		c.JSON(http.StatusOK, gin.H{ //c.JSON() untuk membalas pengguna yang mengakses link
// 		//c.JSON() : isinya adalah seperti Map. ada key, value nya gitu.
// 			"status" : "sukses",
// 			"message" : "Halo Pengguna! Server Gin kamu sudah berhasil berjalan.",
// 		})
// 	})

// 	// 3. Menjalankan server di port 8080
// 	r.Run(":8080")
// }

// /*
// 1. Arti Huruf r
// 	r adalah singkatan dari Router. 
// 	Ini adalah variabel objek mesin utama dari Gin yang kita buat di baris sebelumnya (r := gin.Default()). 
// 	Variabel ini bertugas mengatur semua lalu lintas data yang masuk ke server kamu.
// 2. Arti Kata .GET (HTTP Method)
// 	Ini disebut dengan HTTP Method / HTTP Verb. 
// 	Ini adalah cara client (browser) memberi tahu server tentang aksi atau niat apa yang ingin mereka lakukan:
// 	- (R) .GET: Digunakan saat client hanya ingin mengambil/membaca data dari server (seperti membuka halaman web biasa).
// 	- (C).POST: Digunakan saat client ingin mengirim/membuat data baru ke server (seperti submit form registrasi atau login).
// 	- (U).PUT / .PATCH: Digunakan saat client ingin mengubah/mengupdate data yang sudah ada di server.
// 	- (D).DELETE: Digunakan saat client ingin menghapus data di server.
// 3. Arti Teks "/halo" (Path / Endpoint)
// 	Ini adalah alamat jalan (path) spesifik dari URL.
// 	Jika pengguna membuka http://localhost:8080/halo, maka kode di dalam fungsi ini akan berjalan.
// 	Jika pengguna membuka http://localhost:8080/produk, maka server akan membalas eror 404 Not Found 
// 	karena alamat /produk belum kita daftarkan di sistem routing.
// 4. Arti func(c *gin.Context) (objek yang dibawa masuk ke Controller{func(...) itu baru disebut Controller/Pemberi Perintah})
// Ini adalah otak pengeksekusi atau fungsi yang akan otomatis berjalan ketika ada orang yang mengetik alamat /halo 
// dengan metode GET.c *gin.Context: 
// c adalah singkatan dari Context. 
// Ini adalah "kantong ajaib" bawaan Gin yang berisi semua informasi tentang permintaan si pengguna 
// (seperti: apa IP-nya, apa data yang dia kirim, dan lewat browser apa dia membukanya).
// Kita juga menggunakan variabel c ini untuk mengirim balik balasan ke pengguna, 
// contohnya lewat perintah c.JSON() di dalam kode tersebut.
// */
/*========================================DUA=======================================================*/
/*========================================TIGA=======================================================*/
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// ==========================================
// 1. MODEL(Struktur Data)
// ==========================================
type Pengguna struct{
	ID 		int `json:"id"`
	Nama 	string `json:"nama"`
	Peran 	string `json:"peran"`
}

// Simulasi fungsi database untuk mengambil data pengguna berdasarkan ID

func AmbilDataPenggunaDariDB(id int) (Pengguna, bool) {
	// Anggap saja ini data yang diambil dari MySQL / PostgreSQL
	databaseSimulasi := map[int]Pengguna{
		1: {ID: 1, Nama: "Budi", Peran: "Developer Backend"},
		2: {ID: 2, Nama: "Chloud", Peran: "DevOps Engineer"},
	}

	data, ditemukan := databaseSimulasi[id] // ini comma-ok idiom
	return data, ditemukan
	// Baris kode yang kamu tanyakan memanfaatkan fitur unik di bahasa Go bernama 
	// "Map Map-Index Expression" (atau sering disebut Comma-ok Idiom).
}

func main() {
	r := gin.Default()

	// ==========================================
	// 2. COMPONENT: CONTROLLER (Pengeksekusi Logika)
	// ==========================================
	r.GET("/pengguna", func(c *gin.Context) {
		// Controller meminta bantuan Model untuk mengambil data (Misal kita cari ID nomor 2)
		idTarget := 2
		dataModel, ada := AmbilDataPenggunaDariDB(idTarget) //ini comma-ok idiom

		if !ada {
			// ==========================================
			// 3. COMPONENT: VIEW (Jika Data Eror / Tidak Ada)
			// ==========================================
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "gagal",
				"message": "Data pengguna tidak ditemukan di database",
			})
			return
		}

		// ==========================================
		// 3. COMPONENT: VIEW (Jika Data Sukses Ditemukan)
		// ==========================================
		// Controller menyerahkan data dari Model ke View (Format JSON) untuk dikirim ke browser
		c.JSON(http.StatusOK, gin.H{
			"status": "sukses",
			"data":   dataModel, // Menyisipkan data struct Model ke dalam JSON
		})
	})

	r.Run(":8080")
}
/*========================================TIGA=======================================================*/
