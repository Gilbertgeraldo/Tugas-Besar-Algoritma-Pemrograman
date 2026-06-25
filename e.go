// package main

// import (
// 	"crypto/rand"
// 	"encoding/json"
// 	"fmt"
// 	"math"
// 	"os"
// 	"strings"
// )

// const MAXVAL int = 1000
// const JSON = "data_peminjam.json"

// type Pinjaman struct {
// 	Nama            string
// 	JumlahPinjaman  float64
// 	Tenor           int
// 	Status          string
// 	Bunga           float64
// 	BayaranPerbulan float64
// 	sisaPinjam      int
// 	sudahBayar      int
// 	SchemaBunga     string
// 	TotalBunga      float64
// 	TotalBayar      float64
// 	SudahBayar      int
// }

// type dataPeminjam [MAXVAL]Pinjaman

// var arrPeminjam dataPeminjam
// var countPeminjam int = 0

// func LoadData() {
// 	file, err := os.ReadFile(JSON)
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			return
// 		}
// 		fmt.Println("Error membaca file:", err)
// 		return
// 	}

// 	//json.Unmarshal fungsinya untuk membuka file dari Go dan mengisi data-data ke struct
// 	var tempSlice []Pinjaman
// 	err = json.Unmarshal(file, &tempSlice)
// 	if err != nil {
// 		fmt.Println("Error parsing data JSON:", err)
// 		return
// 	}

// 	// Pindahkan data dari slice ke array global arrPeminjam
// 	countPeminjam = len(tempSlice)
// 	for i := 0; i < countPeminjam; i++ {
// 		arrPeminjam[i] = tempSlice[i]
// 	}
// }

// func SaveData() {
// 	dataAktif := arrPeminjam[:countPeminjam]

// 	// Menggunakan MarshalIndent agar format JSON di file terlihat rapih
// 	fileJSON, err := json.MarshalIndent(dataAktif, "", "  ")
// 	if err != nil {
// 		fmt.Println("Error saat mengubah ke JSON:", err)
// 		return
// 	}

// 	//simpan ke json
// 	err = os.WriteFile(JSON, fileJSON, 0644)
// 	if err != nil {
// 		fmt.Println("Error saat menyimpan file:", err)
// 	}
// }

// func TampilkanDataDenganPagination(limit int) {
// 	if countPeminjam == 0 {
// 		fmt.Println("  Belum ada data peminjam!")
// 		return
// 	}

// 	// Menghitung total halaman
// 	totalPages := (countPeminjam + limit - 1) / limit
// 	currentPage := 1

// 	for {
// 		ClearScreen()
// 		garis1()
// 		fmt.Printf("  DATA PEMINJAM (Halaman %d dari %d)\n", currentPage, totalPages)
// 		garis1()

// 		TabelHeader()

// 		// Menentukan batas awal dan akhir indeks data yang akan di-print
// 		start := (currentPage - 1) * limit
// 		end := start + limit
// 		if end > countPeminjam {
// 			end = countPeminjam
// 		}

// 		for i := start; i < end; i++ {
// 			TabelBaris(arrPeminjam[i], i)
// 		}
// 		garis1()
// 		fmt.Printf("  Menampilkan data ke-%d sampai %d (Total: %d peminjam)\n", start+1, end, countPeminjam)
// 		garis2()

// 		fmt.Println("  Navigasi: [N]ext Page | [P]revious Page | [Q]uit untuk kembali")
// 		fmt.Print("  Pilihan Anda (N/P/Q) : ")

// 		var nav string
// 		fmt.Scan(&nav)

// 		if nav == "N" || nav == "n" {
// 			if currentPage < totalPages {
// 				currentPage++
// 			}
// 		} else if nav == "P" || nav == "p" {
// 			if currentPage > 1 {
// 				currentPage--
// 			}
// 		} else if nav == "Q" || nav == "q" {
// 			break
// 		}
// 	}
// }

// func ClearScreen() {
// 	//Untuk membersihkan terminal
// 	//didapat dari https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go/22892171#22892171
// 	fmt.Print("\033[H\033[2J")
// }

// func garis1() {
// 	fmt.Println("=================================================================================")
// }

// func garis2() {
// 	fmt.Println("-------------------------------------------------------------------------------")
// }

// func DetailPeminjaman(p Pinjaman, idx int) {
// 	garis2()
// 	fmt.Printf("  No.             : %d\n", idx+1)
// 	fmt.Printf("  Nama            : %s\n", p.Nama)
// 	fmt.Printf("  Jumlah Pinjaman : Rp %.2f\n", p.JumlahPinjaman)
// 	fmt.Printf("  Tenor           : %d bulan\n", p.Tenor)
// 	fmt.Printf("  Skema Bunga     : %s\n", p.SchemaBunga)
// 	fmt.Printf("  Bunga Per Tahun : %.2f%%\n", p.Bunga)
// 	fmt.Printf("  Bayaran/Bulan   : Rp %.2f\n", p.BayaranPerbulan)
// 	fmt.Printf("  Total Bunga     : Rp %.2f\n", p.TotalBunga)
// 	fmt.Printf("  Total Bayar     : Rp %.2f\n", p.TotalBayar)
// 	fmt.Printf("  Sudah Bayar     : %d dari %d bulan\n", p.SudahBayar, p.Tenor)
// 	fmt.Printf("  Status          : %s\n", p.Status)
// 	garis2()
// }

// func TabelHeader() {
// 	garis1()
// 	fmt.Printf("| %-3s | %-15s | %-14s | %-5s | %-8s | %-9s | %-12s |\n",
// 		"NO", "Nama", "Pinjaman (Rp)", "Tenor", "Skema", "Status", "Bayaran/Bulan")
// 	garis1()
// }

// func TabelBaris(p Pinjaman, idx int) {
// 	fmt.Printf("| %-3d | %-15s | %14.0f | %5d | %-8s | %-9s | %12.0f |\n",
// 		idx+1, p.Nama, p.JumlahPinjaman, p.Tenor, p.SchemaBunga, p.Status, p.BayaranPerbulan)
// }

// func allTable() {
// 	if countPeminjam == 0 {
// 		fmt.Println("  Belum ada data peminjam!")
// 		return
// 	}
// 	TabelHeader()
// 	for i := 0; i < countPeminjam; i++ {
// 		TabelBaris(arrPeminjam[i], i)
// 	}
// 	garis1()
// 	fmt.Printf("  Total : %d peminjam\n", countPeminjam)
// }

// func enter() {
// 	fmt.Println("\nTekan ENTER untuk kembali ke Menu Utama...")
// 	fmt.Scanln()
// 	fmt.Scanln()
// }

// // INI RUMUS UNTUK KALKULASI BUNGA
// // Bunga variabel anuitas
// // C = p * r(1 + r)^n / ((1 + r)^n - 1)
// // didapat dari https://www.bfi.co.id/id/blog/bunga-anuitas

// func HitungAnuitas(pokok, bunga float64, tenor int) (bayaran, totalBunga, totalBayar float64) {
// 	r := (bunga / 100.0) / 12.0
// 	if r == 0 {
// 		bayaran = math.Round((pokok/float64(tenor))*100) / 100
// 		return
// 	}
// 	rn := math.Pow(1+r, float64(tenor))
// 	bayaran = math.Round(pokok*(r*rn/(rn-1))*100) / 100
// 	totalBayar = math.Round(bayaran*float64(tenor)*100) / 100
// 	totalBunga = math.Round((totalBayar-pokok)*100) / 100
// 	return
// }

// // Hitung bunga yang bersifat Flat, artinya bunga tetap dari pokok awal setiap bulan
// // r           = Bunga / (100 * 12)
// // TotalBunga  = pokok * r * Tenor
// // Bayaran/bln = (pokok + totalBunga) / tenor
// // didapat dari : https://www.bfi.co.id/id/blog/bunga-flat-adalah-pengertian-kelebihan-dan-cara-menghitungnya

// func bungaFlat(pokok, bunga float64, tenor int) (bayaran, totalBunga, totalBayar float64) {
// 	r := (bunga / 100.0) / 12.0
// 	totalBunga = math.Round(pokok*r*float64(tenor)*100) / 100
// 	totalBayar = math.Round((pokok+totalBunga)*100) / 100
// 	bayaran = math.Round((totalBayar/float64(tenor))*100) / 100
// 	return
// }

// func hitungDanSet(p *Pinjaman) {
// 	var bayaran, totalBunga, totalBayar float64
// 	if p.SchemaBunga == "FLAT" {
// 		bayaran, totalBunga, totalBayar = bungaFlat(p.JumlahPinjaman, p.Bunga, p.Tenor)
// 	} else {
// 		bayaran, totalBunga, totalBayar = HitungAnuitas(p.JumlahPinjaman, p.Bunga, p.Tenor)
// 	}
// 	p.BayaranPerbulan = bayaran
// 	p.TotalBunga = totalBunga
// 	p.TotalBayar = totalBayar
// }

// // Rumus Tabel Amortisasi
// // - Angsuran Bunga = Sisa Pokok × Suku Bunga Bulanan
// // - Angsuran Pokok = Total Angsuran - Angsuran Bunga
// // - Sisa Pokok     = Sisa Pokok Bulan Lalu - Angsuran Pokok
// // didapat dari: https://kledo.com/blog/rumus-kalkulator-amortisasi/
// func cetakAmortisasi(p Pinjaman) {
// 	garis1()
// 	fmt.Println("  Tabel Amortisasi (Jadwal Cicilan Anda)")
// 	garis1()
// 	fmt.Printf("  %-5s | %-14s | %-14s | %-14s | %-14s\n",
// 		"Bulan", "Bayaran (Rp)", "Pokok (Rp)", "Bunga (Rp)", "Sisa Pokok (Rp)")
// 	garis1()

// 	sisaPokok := p.JumlahPinjaman
// 	r := (p.Bunga / 100.0) / 12.0

// 	for bulan := 1; bulan <= p.Tenor; bulan++ {
// 		var bungaBulan, pokokBulan float64

// 		if p.SchemaBunga == "FLAT" {
// 			bungaBulan = math.Round(p.JumlahPinjaman*r*100) / 100
// 			pokokBulan = math.Round((p.BayaranPerbulan-bungaBulan)*100) / 100
// 		} else {
// 			bungaBulan = math.Round(sisaPokok*r*100) / 100
// 			pokokBulan = math.Round((p.BayaranPerbulan-bungaBulan)*100) / 100
// 		}

// 		sisaPokok = math.Round((sisaPokok-pokokBulan)*100) / 100
// 		if bulan == p.Tenor {
// 			sisaPokok = 0
// 		}

// 		fmt.Printf("  %-5d | %14.2f | %14.2f | %14.2f | %14.2f\n",
// 			bulan, p.BayaranPerbulan, pokokBulan, bungaBulan, sisaPokok)
// 	}
// 	garis1()
// }

// func inputSkema() string {
// 	fmt.Println("  Skema Bunga :")
// 	fmt.Println("    1. FLAT     - Bunga kamu tetap dari pokok awal.")
// 	fmt.Println("    2. VARIABEL - Bunga anuitas, menurun tiap bulan.")
// 	fmt.Print("  Pilih (1/2) : ")
// 	var s string
// 	fmt.Scan(&s)
// 	if s == "1" {
// 		return "FLAT"
// 	}
// 	return "VARIABEL"
// }

// // =====================================================================
// // CRUD DATA PEMINJAM
// // =====================================================================

// func tambahPeminjam() {
// 	ClearScreen()
// 	garis1()
// 	fmt.Println(" Tambah Data Peminjam")
// 	garis1()

// 	if countPeminjam > 1000 {
// 		fmt.Println("  Data penuh! Maksimal 1000 peminjam.")
// 		return
// 	}

// 	var p Pinjaman

// 	fmt.Print("  Nama peminjam        : ")
// 	fmt.Scan(&p.Nama)
// 	fmt.Print("  Jumlah pinjaman (Rp) : ")
// 	fmt.Scan(&p.JumlahPinjaman)
// 	fmt.Print("  Tenor (bulan)        : ")
// 	fmt.Scan(&p.Tenor)
// 	p.SchemaBunga = inputSkema()
// 	fmt.Print("  Bunga per tahun (%)  : ")
// 	fmt.Scan(&p.Bunga)

// 	if p.JumlahPinjaman <= 0 || p.Tenor <= 0 || p.Bunga < 0 {
// 		fmt.Println("  Mohon maaf, jumlah pinjaman dan tenor harus lebih dari 0.")
// 		return
// 	}

// 	hitungDanSet(&p)
// 	p.Status = "MENUNGGU"
// 	p.sudahBayar = 0

// 	arrPeminjam[countPeminjam] = p
// 	countPeminjam++

// 	SaveData()

// 	fmt.Println("\n  === HASIL KALKULASI ===")
// 	DetailPeminjaman(p, countPeminjam-1)

// 	var pil string
// 	fmt.Print("  Apakah anda ingin menampilkan tabel amortisasi? (Y/N) : ")
// 	fmt.Scan(&pil)
// 	if pil == "Y" || pil == "y" {
// 		cetakAmortisasi(p)
// 	}

// 	fmt.Printf("\n  Peminjaman \"%s\" berhasil ditambahkan!\n", p.Nama)
// }

// func ubahPeminjam() {
// 	ClearScreen()
// 	garis1()
// 	fmt.Println("  UBAH DATA PEMINJAM")
// 	garis1()

// 	if countPeminjam == 0 {
// 		fmt.Println("  Data peminjam belum ada!")
// 		return
// 	}
// 	allTable()

// 	fmt.Print("\n  Masukkan No. peminjam yang akan diubah : ")
// 	var no int
// 	fmt.Scan(&no)

// 	if no < 1 || no > countPeminjam {
// 		fmt.Println("  Nomor tidak valid!")
// 		return
// 	}

// 	idx := no - 1
// 	p := &arrPeminjam[idx]

// 	fmt.Println("\n  Data saat ini:")
// 	DetailPeminjaman(*p, idx)

// 	fmt.Println("\n  Yang ingin diubah :")
// 	fmt.Println("    1. Nama")
// 	fmt.Println("    2. Jumlah Pinjaman") // sudah hitung ulang secara otomatis bila diubah
// 	fmt.Println("    3. Tenor")
// 	fmt.Println("    4. Skema dan Bunga") // sudah hitung ulang secara otomatis
// 	fmt.Println("    5. Status Pembayaran")
// 	fmt.Println("    6. Jumlah bulan yang sudah dibayar")
// 	fmt.Print("  Pilihan : ")

// 	var pil string
// 	fmt.Scan(&pil)

// 	switch pil {
// 	case "1":
// 		fmt.Print("  Nama baru : ")
// 		fmt.Scan(&p.Nama)
// 	case "2":
// 		fmt.Print("  Jumlah pinjaman baru (Rp) : ")
// 		fmt.Scan(&p.JumlahPinjaman)
// 		hitungDanSet(p)
// 	case "3":
// 		fmt.Print("  Tenor baru (bulan) : ")
// 		fmt.Scan(&p.Tenor)
// 		hitungDanSet(p)
// 	case "4":
// 		p.SchemaBunga = inputSkema()
// 		fmt.Print("  Bunga per tahun (%) baru : ")
// 		fmt.Scan(&p.Bunga)
// 		if p.Bunga < 0 {
// 			fmt.Println("  Bunga tidak boleh negatif.")
// 		}
// 		hitungDanSet(p)
// 	case "5":
// 		fmt.Println("  Pilih status :")
// 		fmt.Println("    1. MENUNGGU")
// 		fmt.Println("    2. AKTIF")
// 		fmt.Println("    3. LUNAS")
// 		fmt.Println("    4. MACET")
// 		fmt.Print("  Pilihan : ")
// 		var sp string
// 		fmt.Scan(&sp)
// 		switch sp {
// 		case "1":
// 			p.Status = "MENUNGGU"
// 		case "2":
// 			p.Status = "AKTIF"
// 		case "3":
// 			p.Status = "LUNAS"
// 		case "4":
// 			p.Status = "MACET"
// 		default:
// 			fmt.Println("  Pilihan tidak valid.") 
// 		SaveData()
// 		fmt.Printf("\n  Data peminjam \"%s\" berhasil diubah!\n", p.Nama)
// 		DetailPeminjaman(*p, idx)
// 		enter()
// 	}
// }
// }

// func hapusPeminjam() {
// 	ClearScreen()
// 	garis1()
// 	fmt.Println("  HAPUS DATA PEMINJAM")
// 	garis1()

// 	if countPeminjam == 0 {
// 		fmt.Println("  Belum ada data peminjam!")
// 		return
// 	}
// 	allTable()

// 	fmt.Print("\n  Masukkan No. peminjam yang ingin dihapus : ")
// 	var no int
// 	fmt.Scan(&no)

// 	if no < 1 || no > countPeminjam {
// 		fmt.Println("  Nomor tidak valid!")
// 		return
// 	}

// 	idx := no - 1
// 	nama := arrPeminjam[idx].Nama

// 	fmt.Printf("  Apakah anda yakin untuk menghapus \"%s\"? (Y/N) : ", nama)
// 	var konfirmasi string
// 	fmt.Scan(&konfirmasi)

// 	if konfirmasi != "Y" && konfirmasi != "y" {
// 		fmt.Println("  Penghapusan dibatalkan.")
// 		return
// 	}

// 	for i := idx; i < countPeminjam-1; i++ {
// 		arrPeminjam[i] = arrPeminjam[i+1]
// 	}
// 	countPeminjam--
// 	SaveData()
// 	fmt.Printf("  Peminjam \"%s\" berhasil dihapus!\n", nama)
// }

// func binaSearch(arr dataPeminjam, n int, target string) int {
// 	left := 0
// 	right := n - 1
// 	targetLowerCase := strings.ToLower(target)
	
// 	for left <= right {
// 		mid := (left + right) / 2
// 		midLowerCase := strings.ToLower(arr[mid].Nama)
		
// 		// Menggunakan HasPrefix agar ketik "Agus" bisa mendeteksi "Agus Kusuma"
// 		if midLowerCase == targetLowerCase || strings.HasPrefix(midLowerCase, targetLowerCase) {
// 			return mid
// 		} else if midLowerCase < targetLowerCase { // Typo pembanding target sudah diperbaiki
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}
// 	return -1
// }

// func GenerateRandomID()string {
// 	b := make([]byte,2)
// 	rand.Read(b)
// 	return fmt.Sprintf("ID-%X",b)
// }

// func InsertionSort(arr *dataPeminjam, n int, fieldName string, isAscending bool) {
// 	sub := 0
// 	for sub < n-1 {
// 		ins := sub + 1
// 		t := sub
// 		for t >= 0 && Swap(arr, ins, t, fieldName, isAscending) {
// 			arr[ins], arr[t] = arr[t], arr[ins]
// 			ins--
// 			t--
// 		}
// 		sub++
// 	}
// }

// func selectionSort(arr *dataPeminjam, n int, fieldName string, isAscending bool) {
// 	for i := 0; i < n-1; i++ {
// 		minIdx := i
// 		for j := i + 1; j < n; j++ {
// 			if Swap(arr, j, minIdx, fieldName, isAscending) {
// 				minIdx = j
// 			}
// 		}
// 		arr[i], arr[minIdx] = arr[minIdx], arr[i]
// 	}
// }

// func Swap(arr *dataPeminjam, ins, t int, fieldName string, isAscending bool) bool {
// 	var status bool
// 	switch fieldName {
// 	case "Pinjaman":
// 		status = arr[ins].JumlahPinjaman < arr[t].JumlahPinjaman
// 	case "Tenor":
// 		status = arr[ins].Tenor < arr[t].Tenor
// 	case "Nama":
// 		status = strings.ToLower(arr[ins].Nama) < strings.ToLower(arr[t].Nama)
// 	case "ID":
// 		status = arr[ins].Nama < arr[t].Nama
// 	default:
// 		return false
// 	}
// 	if isAscending {
// 		return status
// 	} else {
// 		return !status
// 	}
// }

// func MenuSorting() {
// 	ClearScreen()
// 	garis1()
// 	fmt.Println("  URUTKAN DATA PEMINJAM")
// 	garis1()

// 	if countPeminjam == 0 {
// 		fmt.Println("  Belum ada data peminjam!")
// 		enter()
// 		return
// 	}

// 	fmt.Println("  Urutkan berdasarkan :")
// 	fmt.Println("    1. Jumlah Pinjaman (Selection Sort)")
// 	fmt.Println("    2. Tenor           (Selection Sort)")
// 	fmt.Println("    3. Jumlah Pinjaman (Insertion Sort)")
// 	fmt.Println("    4. Tenor           (Insertion Sort)")
// 	fmt.Println("	 5. Nama			(Insertion Sort)")
// 	fmt.Println("	 6. Nama			(Selection Sort)")

// 	fmt.Print("  Pilihan : ")

// 	var pil string
// 	fmt.Scan(&pil)

// 	switch pil {
// 	case "1":
// 		fmt.Println("Mau Mengurutkan jumlah Pinjaman dari apa dulu nih : ")
// 		fmt.Println("1.	Ascending(Terkecil->Terbesar)")
// 		fmt.Println("2.	Descending(Terbesar->Terkecil)")
// 		var pi string
// 		fmt.Scan(&pi)
// 		switch pi {
// 		case "1":
// 			selectionSort(&arrPeminjam, countPeminjam, "Pinjaman", true)
// 			fmt.Println("Data diurutkan berdasarkan Pinjaman Terkecil ke Terbesar (Ascending).")
// 		case "2":
// 			selectionSort(&arrPeminjam, countPeminjam, "Pinjaman", false)
// 			fmt.Println("Data diurutkan berdasarkan jumlah pinjaman Terbesar ke Terkecil (Descending).")
// 		default:
// 			fmt.Println("	Error! Pilihan anda tidak valid.")
// 		}
// 	case "2":
// 		fmt.Println("Mau mengurutkan jumlah Tenor dari apa dulu nih : ")
// 		fmt.Println("1.	Ascending(Terkecil->Terbesar)")
// 		fmt.Println("2.	Descending(Terbesar->Terkecil)")
// 		var pi string
// 		fmt.Scan(&pi)
// 		switch pi {
// 		case "1":
// 			selectionSort(&arrPeminjam, countPeminjam, "Tenor", true)
// 			fmt.Println("Data diurutkan berdasarkan jumlah Tenor Terkecil ke Terbesar(Ascending).")
// 		case "2":
// 			selectionSort(&arrPeminjam, countPeminjam, "Tenor", false)
// 			fmt.Println("Data diurutkan berdasarkan jumlah Tenor Terbesar ke Terkecil (Descending).")
// 		default:
// 			fmt.Println("	Error! Pilihan anda tidak valid.")
// 		}
// 	case "3":
// 		fmt.Println("Mau Mengurutkan jumlah Pinjaman dari apa dulu nih : ")
// 		fmt.Println("1.	Ascending(Terkecil->Terbesar)")
// 		fmt.Println("2.	Descending(Terbesar->Terkecil)")
// 		var pi string
// 		fmt.Scan(&pi)
// 		switch pi {
// 		case "1":
// 			InsertionSort(&arrPeminjam, countPeminjam, "Pinjaman", true)
// 			fmt.Println("Data diurutkan berdasarkan jumlah pinjaman Terkecil ke Terbesar(Ascending).")
// 		case "2":
// 			InsertionSort(&arrPeminjam, countPeminjam, "Pinjaman", false)
// 			fmt.Println("Data diurutkan berdasarkan jumlah pinjaman Terbesar ke Terkecil(Descending).")
// 		default:
// 			fmt.Println("	Error! Pilihan anda tidak valid.")
// 		}
// 	case "4":
// 		fmt.Println("Mau mengurutkan jumlah Tenor dari apa dulu nih : ")
// 		fmt.Println("1.	Ascending(Terkecil->Terbesar)")
// 		fmt.Println("2.	Descending(Terbesar->Terkecil)")
// 		var pi string
// 		fmt.Scan(&pi)
// 		switch pi {
// 		case "1":
// 			InsertionSort(&arrPeminjam, countPeminjam, "Tenor", true)
// 			fmt.Println("Data diurutkan berdasarkan jumlah Tenor Terkecil ke Terbesar(Ascending).")
// 		case "2":
// 			InsertionSort(&arrPeminjam, countPeminjam, "Tenor", false)
// 			fmt.Println("Data diurutkan berdasarkan jumlah Tenor Terbesar ke Terkecil (Descending).")
// 		default:
// 			fmt.Println("	Error! Pilihan anda tidak valid!")
// 		}
// 	case "5":
// 		fmt.Println("Mau mengurutkan Nama dari apa dulu nih : ")
// 		fmt.Println("1.	Ascending(A->Z)")
// 		fmt.Println("2.	Descending(Z->A)")
// 		var pi string
// 		fmt.Scan(&pi)
// 		switch pi {
// 		case "1":
// 			InsertionSort(&arrPeminjam, countPeminjam, "Nama", true)

// 			fmt.Println("Data diurutkan berdasarkan Nama dari A ke Z(Ascending).")
// 		case "2":
// 			InsertionSort(&arrPeminjam, countPeminjam, "Nama", false)

// 			fmt.Println("Data diurutkan berdasarkan Nama dari Z ke A(Descending).")
// 		default:
// 			fmt.Println("	Error! Pilihan anda tidak valid!")
// 		}
// 	case "6":
// 		fmt.Println("Mau mengurutkan Nama dari apa dulu nih : ")
// 		fmt.Println("1.	Ascending(A->Z)")
// 		fmt.Println("2.	Descending(Z->A)")
// 		var pi string
// 		fmt.Scan(&pi)
// 		switch pi {
// 		case "1":
// 			selectionSort(&arrPeminjam, countPeminjam, "Nama", true)

// 			fmt.Println("Data diurutkan berdasarkan Nama dari A ke Z(Ascending).")
// 		case "2":
// 			selectionSort(&arrPeminjam, countPeminjam, "Nama", false)

// 			fmt.Println("Data diurutkan berdasarkan Nama dari Z ke A(Descending).")
// 		default:
// 			fmt.Println("	Error! Pilihan anda tidak valid!")
// 		}
// 	default:
// 		fmt.Println("  Error! Pilihan tidak valid.")
// 		enter()
// 		return
// 	}
// 	TampilkanDataDenganPagination(20)
// }

// func MenuSearching() {
// 	ClearScreen()
// 	garis1()
// 	fmt.Println("  CARI DATA PEMINJAM")
// 	garis1()

// 	if countPeminjam == 0 {
// 		fmt.Println("  Belum ada data peminjam!")
// 		enter()
// 		return
// 	}

// 	fmt.Println("  Pilih metode pencarian :")
// 	fmt.Println("    1. Sequential Search")
// 	fmt.Println("    2. Binary Search")
// 	fmt.Print("  Pilihan : ")

// 	var pil string
// 	fmt.Scan(&pil)

// 	switch pil {
// 	case "1":
// 		SequentialSearch()
// 	case "2":
// 		var keyword string
// 		fmt.Print("  Masukkan nama peminjam yang dicari : ")
// 		fmt.Scan(&keyword)
// 		tampilkanBinarySearch(arrPeminjam, countPeminjam, keyword)
// 	default:
// 		fmt.Println("  Error! Pilihan tidak valid.")
// 	}
// 	enter()
// }

// func Laporan() {
// 	ClearScreen()
// 	garis1()
// 	fmt.Println("  LAPORAN SISTEM PEMINJAMAN")
// 	garis1()

// 	if countPeminjam == 0 {
// 		fmt.Println("  Belum ada data peminjam.")
// 		enter()
// 		return
// 	}

// 	var totalPokok, totalBayar, totalBunga float64
// 	var cMen, cAk, cLun, cMcet int

// 	for i := 0; i < countPeminjam; i++ {
// 		p := arrPeminjam[i]
// 		totalPokok += p.JumlahPinjaman
// 		totalBayar += p.TotalBayar
// 		totalBunga += p.TotalBunga

// 		switch p.Status {
// 		case "MENUNGGU":
// 			cMen++
// 		case "AKTIF":
// 			cAk++
// 		case "LUNAS":
// 			cLun++
// 		case "MACET":
// 			cMcet++
// 		}
// 	}

// 	fmt.Printf("  Total Peminjam       : %d orang\n", countPeminjam)
// 	fmt.Printf("  Total Pokok Pinjaman : Rp %.2f\n", totalPokok)
// 	fmt.Printf("  Total Bunga          : Rp %.2f\n", totalBunga)
// 	fmt.Printf("  Total Nilai Bayar    : Rp %.2f\n", totalBayar)
// 	garis2()
// 	fmt.Println("  Status Pembayaran :")
// 	fmt.Printf("    MENUNGGU : %d peminjam\n", cMen)
// 	fmt.Printf("    AKTIF    : %d peminjam\n", cAk)
// 	fmt.Printf("    LUNAS    : %d peminjam\n", cLun)
// 	fmt.Printf("    MACET    : %d peminjam\n", cMcet)
// 	garis2()
// 	fmt.Println("\n  Daftar Semua Peminjam :")
// 	TampilkanDataDenganPagination(20)
// }

// func cetakJudul() {
// 	judul := []string{
// 		`________          __                                   ______   __                               `,
// 		`/        |        /  |                                 /      \ /  |                              `,
// 		`$$$$$$$$/__    __ $$ |____    ______    _______       /$$$$$$  |$$ |  ______    ______    ______  `,
// 		`   $$ | /  |  /  |$$      \  /      \  /       |      $$ |__$$ |$$ | /      \  /      \  /      \`,
// 		`   $$ | $$ |  $$ |$$$$$$$  |/$$$$$$  |/$$$$$$$/       $$    $$ |$$ |/$$$$$$  |/$$$$$$  |/$$$$$$  |`,
// 		`   $$ | $$ |  $$ |$$ |  $$ |$$    $$ |$$      \       $$$$$$$$ |$$ |$$ |  $$ |$$ |  $$/ $$ |  $$ |`,
// 		`   $$ | $$ \__$$ |$$ |__$$ |$$$$$$$$/  $$$$$$  |      $$ |  $$ |$$ |$$ |__$$ |$$ |      $$ \__$$ |`,
// 		`   $$ | $$    $$/ $$    $$/ $$       |/     $$/       $$ |  $$ |$$ |$$    $$/ $$ |      $$    $$/  `,
// 		`   $$/   $$$$$$/  $$$$$$$/   $$$$$$$/ $$$$$$$/        $$/   $$/ $$/ $$$$$$$/  $$/        $$$$$$/   `,
// 		`                                                                    $$ |                            `,
// 		`                                                                    $$ |                            `,
// 		`                                                                    $$/                             `,
// 	}

// 	warna := []string{
// 		"\033[31m", // Merah
// 		"\033[33m", // Kuning
// 		"\033[32m", // Hijau
// 		"\033[36m", // Cyan
// 		"\033[34m", // Biru
// 		"\033[35m", // Ungu
// 		"\033[91m", // Merah terang
// 		"\033[92m", // Hijau terang
// 		"\033[93m", // Kuning terang
// 		"\033[94m", // Biru terang
// 		"\033[95m", // Ungu terang
// 		"\033[96m", // Cyan terang
// 	}

// 	for i, baris := range judul {
// 		fmt.Println(warna[i%len(warna)] + baris + "\033[0m")
// 	}
// }

// func main() {
// 	LoadData()
// 	for {
// 		ClearScreen()
// 		cetakJudul()
// 		garis2()
// 		fmt.Println("  Anggota :")
// 		fmt.Println("  1. Gilbert Geraldo (103052500054)")
// 		fmt.Println("  2. Jafar Shiddiq   (103052500002)")
// 		garis2()
// 		fmt.Println("  1. Tambah Peminjam")
// 		fmt.Println("  2. Ubah Data Peminjam")
// 		fmt.Println("  3. Hapus Peminjam")
// 		fmt.Println("  4. Lihat Semua Pinjaman")
// 		fmt.Println("  5. Cari Peminjam")
// 		fmt.Println("  6. Urutkan Peminjaman")
// 		fmt.Println("  7. Laporan")
// 		fmt.Println("  0. Keluar")
// 		garis1()
// 		fmt.Print("  Silahkan masukkan pilihan anda : ")

// 		var p string
// 		fmt.Scan(&p)

// 		switch p {
// 		case "1":
// 			tambahPeminjam()
// 		case "2":
// 			ubahPeminjam()
// 		case "3":
// 			hapusPeminjam()
// 		case "4":
// 			ClearScreen()
// 			garis1()
// 			fmt.Println("  DAFTAR SEMUA PEMINJAM")
// 			garis1()
// 			allTable()
// 			enter()
// 		case "5":
// 			MenuSearching()
// 		case "6":
// 			MenuSorting()
// 		case "7":
// 			Laporan()
// 		case "0":
// 			fmt.Println("  Terima kasih... Sampai Jumpa!")
// 			return
// 		default:
// 			fmt.Println("  Pilihan tidak valid.")
// 			enter()
// 		}
// 	}
// }