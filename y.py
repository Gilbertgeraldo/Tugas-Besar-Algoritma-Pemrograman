import json
import random
import itertools

# Variasi nama diperbanyak agar kombinasinya makin kaya
first_names = ["Danu", "Budi", "Ayu", "Siti", "Joko", "Rina", "Agus", "Putri", "Gilang", "Nisa", "Bima", "Cinta", "Wahyu", "Dewi", "Reza", "Tari", "Fajar", "Ahmad", "Muhammad", "Kevin"]
middle_names = ["Dwi", "Tri", "Nur", "Eka", "Setia", "Bakti", "Kurnia", "Maulana", "Fitri", "Indra", "Purnama", "Ramadhan", "Wahid", "Bagus"]
last_names = ["Pratama", "Santoso", "Wijaya", "Sari", "Kusuma", "Lestari", "Nugroho", "Hidayat", "Saputra", "Wibowo", "Siregar", "Pangestu", "Setiawan", "Hakim", "Nasution"]

# 1. Bangun variasi 2 kata (Nama Depan + Belakang)
kombinasi_2_kata = [f"{f} {l}" for f in first_names for l in last_names]

# 2. Bangun variasi 3 kata (Nama Depan + Tengah + Belakang)
kombinasi_3_kata = [f"{f} {m} {l}" for f in first_names for m in middle_names for l in last_names]

# Gabungkan semua kemungkinan nama (Totalnya bisa mencapai ribuan)
semua_kombinasi_nama = kombinasi_2_kata + kombinasi_3_kata

# 3. Ambil TEPAT 100 nama secara acak dan pastikan UNIK
nama_terpilih = random.sample(semua_kombinasi_nama, 100)

data_peminjam = []

# Lakukan perulangan persis sebanyak 100 kali sesuai nama yang ditarik
for nama in nama_terpilih:
    pinjaman = float(random.randint(10, 500) * 10000) # Random dari 100.000 sampai 5.000.000
    tenor = random.choice([3, 6, 12, 24, 36])
    bunga = float(random.randint(5, 15))
    schema = random.choice(["FLAT", "VARIABEL"])
    status = random.choice(["MENUNGGU", "AKTIF", "LUNAS", "MACET"])

    r = (bunga / 100.0) / 12.0
    
    # Kalkulasi sesuai rumus Golang-mu
    if schema == "FLAT":
        totalBunga = round(pinjaman * r * tenor, 2)
        totalBayar = round(pinjaman + totalBunga, 2)
        bayaran = round(totalBayar / tenor, 2)
    else: # VARIABEL / Anuitas
        if r == 0:
            bayaran = round(pinjaman / tenor, 2)
        else:
            rn = (1 + r) ** tenor
            bayaran = round(pinjaman * (r * rn / (rn - 1)), 2)
        totalBayar = round(bayaran * tenor, 2)
        totalBunga = round(totalBayar - pinjaman, 2)
        
    # Logika status pembayaran
    if status == "MENUNGGU":
        sudah_bayar = 0
    elif status == "LUNAS":
        sudah_bayar = tenor
    else:
        sudah_bayar = random.randint(1, tenor - 1) if tenor > 1 else 0

    peminjam = {
        "Nama": nama,
        "JumlahPinjaman": pinjaman,
        "Tenor": tenor,
        "Status": status,
        "Bunga": bunga,
        "BayaranPerbulan": bayaran,
        "SchemaBunga": schema,
        "TotalBunga": totalBunga,
        "TotalBayar": totalBayar,
        "SudahBayar": sudah_bayar
    }
    
    data_peminjam.append(peminjam)

# Simpan langsung ke file json
with open("data_peminjam.json", "w", encoding="utf-8") as file:
    json.dump(data_peminjam, file, indent=2)

print("Berhasil! File data_peminjam.json dengan TEPAT 100 data peminjam UNIK telah dibuat.")