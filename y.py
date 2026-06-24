import json
import random

# Variasi nama untuk menguji fitur Search dan Sorting
first_names = ["Danu", "Budi", "Ayu", "Siti", "Joko", "Rina", "Agus", "Putri", "Gilang", "Nisa", "Bima", "Cinta"]
last_names = ["Pratama", "Santoso", "Wijaya", "Sari", "Kusuma", "Lestari", "Nugroho", "Hidayat", "Saputra"]

data_peminjam = []

for _ in range(1000):
    nama = f"{random.choice(first_names)} {random.choice(last_names)}"
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
with open("data_peminjam.json", "w") as file:
    json.dump(data_peminjam, file, indent=2)

print("Berhasil! File data_peminjam.json dengan 1000 data telah dibuat.")