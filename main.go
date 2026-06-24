package main

import (
	"net/http" //ya internet gitu ya
	"github.com/gin-gonic/gin" //ini frameworknya
)

func main() {
	// 1. Inisialisasi framework Gin
	r := gin.Default() //ini router. variabel objek mesin utama dari Gin; 
					   // tugasnya mengatur semua traffic data yg masuk ke server db

	// 2. Membuat rute (endpoint) HTTP GET /halo
	//  
	r.GET("/halo", func(c *gin.Context){ // GET berarti user mau mengambil/membaca data aja.
		// "/halo" cuma nama untuk situs localhost nya. bisa diganti jadi /aloha dll
		// c itu context. sebuah kantong ajaib yang bisa membawa seluruh konteks yang user perlu. 
			// *gin.Context anggap aja ya c memberi/mendapat context dari gin
		//c.JSON() ---> View
		c.JSON(http.StatusOK, gin.H{ //c.JSON() untuk membalas pengguna yang mengakses link
		//c.JSON() : isinya adalah seperti Map. ada key, value nya gitu.
			"status" : "sukses",
			"message" : "Halo Pengguna! Server Gin kamu sudah berhasil berjalan.",
		})
	})

	// 3. Menjalankan server di port 8080
	r.Run(":8080")
}

/*
==================================================PENJELASAN===============================================================
1. Arti Huruf r
	r adalah singkatan dari Router. 
	Ini adalah variabel objek mesin utama dari Gin yang kita buat di baris sebelumnya (r := gin.Default()). 
	Variabel ini bertugas mengatur semua lalu lintas data yang masuk ke server kamu.
2. Arti Kata .GET (HTTP Method)
	Ini disebut dengan HTTP Method / HTTP Verb. 
	Ini adalah cara client (browser) memberi tahu server tentang aksi atau niat apa yang ingin mereka lakukan:
	- (R) .GET: Digunakan saat client hanya ingin mengambil/membaca data dari server (seperti membuka halaman web biasa).
	- (C).POST: Digunakan saat client ingin mengirim/membuat data baru ke server (seperti submit form registrasi atau login).
	- (U).PUT / .PATCH: Digunakan saat client ingin mengubah/mengupdate data yang sudah ada di server.
	- (D).DELETE: Digunakan saat client ingin menghapus data di server.
3. Arti Teks "/halo" (Path / Endpoint)
	Ini adalah alamat jalan (path) spesifik dari URL.
	Jika pengguna membuka http://localhost:8080/halo, maka kode di dalam fungsi ini akan berjalan.
	Jika pengguna membuka http://localhost:8080/produk, maka server akan membalas eror 404 Not Found 
	karena alamat /produk belum kita daftarkan di sistem routing.
4. Arti func(c *gin.Context) (objek yang dibawa masuk ke Controller{func(...) itu baru disebut Controller/Pemberi Perintah})
Ini adalah otak pengeksekusi atau fungsi yang akan otomatis berjalan ketika ada orang yang mengetik alamat /halo 
dengan metode GET.c *gin.Context: 
c adalah singkatan dari Context. 
Ini adalah "kantong ajaib" bawaan Gin yang berisi semua informasi tentang permintaan si pengguna 
(seperti: apa IP-nya, apa data yang dia kirim, dan lewat browser apa dia membukanya).
Kita juga menggunakan variabel c ini untuk mengirim balik balasan ke pengguna, 
contohnya lewat perintah c.JSON() di dalam kode tersebut.
*/
