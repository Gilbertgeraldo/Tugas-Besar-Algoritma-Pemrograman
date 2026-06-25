package main

import (
	"crypto/rand"
	"fmt"
)

// GenerateRandomID menghasilkan ID random berupa 4 karakter hexadesimal
func GenerateRandomID12() string {
	b := make([]byte, 2)
	rand.Read(b)
	return fmt.Sprintf("ID-%X", b)
}

// SearchByID mencari peminjam berdasarkan ID (Sequential Search)
func SearchByID(id string) {
	found := false
	fmt.Println()
	garis1()
	fmt.Println("  HASIL SEQUENTIAL SEARCH BERDASARKAN ID")
	garis1()

	for i := 0; i < countPeminjam; i++ {
		if arrPeminjam[i].ID == id {
			DetailPeminjaman(arrPeminjam[i], i)
			found = true
			break // ID bersifat unik, jadi bisa langsung break setelah ketemu
		}
	}

	if !found {
		fmt.Printf("  Peminjam dengan ID \"%s\" tidak ditemukan.\n", id)
	}
}

// SortByID mengurutkan array berdasarkan ID secara Ascending
// Diperlukan agar Binary Search bisa berjalan dengan benar
func SortByID() {
	for i := 1; i < countPeminjam; i++ {
		key := arrPeminjam[i]
		j := i - 1
		for j >= 0 && arrPeminjam[j].ID > key.ID {
			arrPeminjam[j+1] = arrPeminjam[j]
			j = j - 1
		}
		arrPeminjam[j+1] = key
	}
}

// BinarySearchByID mencari peminjam berdasarkan ID menggunakan metode Binary Search
func BinarySearchByID(id string) {
	// Urutkan terlebih dahulu menggunakan metode pengurutan (misal: Insertion Sort)
	SortByID()

	left := 0
	right := countPeminjam - 1
	foundIdx := -1

	for left <= right {
		mid := (left + right) / 2
		if arrPeminjam[mid].ID == id {
			foundIdx = mid
			break
		} else if arrPeminjam[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println()
	garis1()
	fmt.Println("  HASIL BINARY SEARCH BERDASARKAN ID")
	garis1()

	if foundIdx != -1 {
		DetailPeminjaman(arrPeminjam[foundIdx], foundIdx)
	} else {
		fmt.Printf("  Peminjam dengan ID \"%s\" tidak ditemukan.\n", id)
	}
}
