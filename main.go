package main
import (
	"net/http"
	"strconv"
	"sync"
	"github.com/gin-gonic/gin"
)
// ==========================================
// 1. MODEL (Struktur Data)
// ==========================================
type Pengguna struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Peran string `json:"peran"`
}
// Simulasi Database Global (Disimpan di Memory RAM)
// Kita tambahkan sync.RWMutex karena Gin berjalan secara konkuren (multithreading).
// Mutex ini bertugas menjaga agar data tidak bentrok/error saat dibaca & ditulis bersamaan.
var (
	databaseSimulasi = map[int]Pengguna{
		1: {ID: 1, Nama: "Budi", Peran: "Developer Backend"},
		2: {ID: 2, Nama: "Chloud", Peran: "DevOps Engineer"},
	}
	dbMutex sync.RWMutex
	lastID  = 2 // Digunakan untuk auto-increment ID baru saat Create
)
func main() {
	r := gin.Default()
	// ==========================================
	// [CREATE] - POST /pengguna
	// ==========================================
	r.POST("/pengguna", func(c *gin.Context) {
		var input struct {
			Nama  string `json:"nama" binding:"required"`
			Peran string `json:"peran" binding:"required"`
		}
		// Validasi input JSON dari client
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "gagal",
				"error":  "Nama dan Peran harus diisi",
			})
			return
		}
		// Kunci database untuk menulis data baru
		dbMutex.Lock()
		lastID++
		baru := Pengguna{
			ID:    lastID,
			Nama:  input.Nama,
			Peran: input.Peran,
		}
		databaseSimulasi[lastID] = baru
		dbMutex.Unlock()
		c.JSON(http.StatusCreated, gin.H{
			"status":  "sukses",
			"message": "Pengguna berhasil ditambahkan",
			"data":    baru,
		})
	})
	// ==========================================
	// [READ ALL] - GET /pengguna
	// ==========================================
	r.GET("/pengguna", func(c *gin.Context) {
		dbMutex.RLock()
		defer dbMutex.RUnlock() //pake defer agar lebih rapi

		// Mengubah Map menjadi Slice/Array agar rapi saat ditampilkan di JSON
		var list []Pengguna
		for _, v := range databaseSimulasi {
			list = append(list, v)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "sukses",
			"data":   list,
		})
	})
	// ==========================================
	// [READ ONE] - GET /pengguna/:id
	// ==========================================
	r.GET("/pengguna/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam) // Mengubah string ID menjadi int
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "gagal", "error": "ID harus angka"})
			return
		}
		dbMutex.RLock()
		data, ada := databaseSimulasi[id]
		dbMutex.RUnlock()
		if !ada {
			c.JSON(http.StatusNotFound, gin.H{"status": "gagal", "message": "User tidak ditemukan"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "sukses", "data": data})
	})
	// ==========================================
	// [UPDATE] - PUT /pengguna/:id
	// ==========================================
	r.PUT("/pengguna/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "gagal", "error": "ID harus angka"})
			return
		}
		var input struct {
			Nama  string `json:"nama" binding:"required"`
			Peran string `json:"peran" binding:"required"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "gagal", "error": err.Error()})
			return
		}
		dbMutex.Lock()
		defer dbMutex.Unlock()
		// Cek apakah data yang mau diupdate ada di database
		if _, ada := databaseSimulasi[id]; !ada {
			c.JSON(http.StatusNotFound, gin.H{"status": "gagal", "message": "User tidak ditemukan"})
			return
		}
		// Update data
		databaseSimulasi[id] = Pengguna{
			ID:    id,
			Nama:  input.Nama,
			Peran: input.Peran,
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "sukses",
			"message": "Data berhasil diperbarui",
			"data":    databaseSimulasi[id],
		})
	})
	// ==========================================
	// [DELETE] - DELETE /pengguna/:id
	// ==========================================
	r.DELETE("/pengguna/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "gagal", "error": "ID harus angka"})
			return
		}
		dbMutex.Lock()
		defer dbMutex.Unlock()
		// Cek apakah data ada
		if _, ada := databaseSimulasi[id]; !ada {
			c.JSON(http.StatusNotFound, gin.H{"status": "gagal", "message": "User tidak ditemukan"})
			return
		}
		// Hapus dari map database
		delete(databaseSimulasi, id)
		c.JSON(http.StatusOK, gin.H{
			"status":  "sukses",
			"message": "User berhasil dihapus",
		})
	})
	r.Run(":8080")
}