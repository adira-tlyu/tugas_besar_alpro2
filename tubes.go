package main

import "fmt"

const NMAX int = 1024
const KEY string = "wZuNpR8h3vKD"

type daftarAsetKripto struct {
	nama         string
	harga        float64
	kapitalisasi float64
}

type daftarAsetPengguna struct {
	nama   string
	jumlah float64
	harga  float64
}

type riwayatTransaksi struct {
	nama   string
	jumlah float64
}

type arrAsetK [NMAX-1]daftarAsetKripto
type arrAsetP [NMAX-1]daftarAsetPengguna
type arrTransaksi [NMAX-1]riwayatTransaksi

func main() {
	var daftarK arrAsetK
	var daftarP arrAsetP
	var riwayat arrTransaksi
	var first, exit bool
	var s float64 = 100000
	var n, countDaftarK, countDaftarP, countRiwayat int
	var asciiDone string
	
	asciiDone = `
  /$$$$$$  /$$$$$$$$ /$$       /$$$$$$$$  /$$$$$$   /$$$$$$  /$$$$$$
 /$$__  $$| $$_____/| $$      | $$_____/ /$$__  $$ /$$__  $$|_  $$_/
| $$  \__/| $$      | $$      | $$      | $$  \__/| $$  \ $$  | $$  
|  $$$$$$ | $$$$$   | $$      | $$$$$   |  $$$$$$ | $$$$$$$$  | $$  
 \____  $$| $$__/   | $$      | $$__/    \____  $$| $$__  $$  | $$  
 /$$  \ $$| $$      | $$      | $$       /$$  \ $$| $$  | $$  | $$  
|  $$$$$$/| $$$$$$$$| $$$$$$$$| $$$$$$$$|  $$$$$$/| $$  | $$ /$$$$$$
 \______/ |________/|________/|________/ \______/ |__/  |__/|______/
                                                                    
                                                                    
                                                                    
`
	
	first = true
	testData(&countDaftarK, &daftarK)
	
	for !exit {
		showMenu(&first, s)
		fmt.Scan(&n)
		
		if n == 9 {
			fmt.Print(asciiDone)
			exit = true
		} else {
			switch n {
				case 1:
					beliAset(countDaftarK, &countDaftarP, &countRiwayat, daftarK, &daftarP, &riwayat, &s)
				case 2:
					jualAset(&countDaftarP, &countRiwayat, &daftarP, &riwayat, &s)
				case 3:
					tampilRiwayat(countRiwayat, riwayat)
				case 4:
					pilihSorting(countDaftarK, &daftarK)
				case 5:
					tampilAsetPengguna(countDaftarP, daftarP)
				case 6:
					tabelKoinK(countDaftarK, daftarK)
				case 7:
					admin(&countDaftarK, &countDaftarP, &daftarK, &daftarP, KEY)
				case 8:
					cariAset(countDaftarK, daftarK)
				default:
					fmt.Println("\n[INVALID]")
			}
		}
	}
}

func showMenu(first *bool, s float64) {
	var asciiMenu string
	
	asciiMenu = `
  /$$$$$$  /$$$$$$ /$$      /$$ /$$   /$$ /$$        /$$$$$$   /$$$$$$  /$$$$$$       /$$   /$$ /$$$$$$$  /$$$$$$ /$$$$$$$  /$$$$$$$$ /$$$$$$ 
 /$$__  $$|_  $$_/| $$$    /$$$| $$  | $$| $$       /$$__  $$ /$$__  $$|_  $$_/      | $$  /$$/| $$__  $$|_  $$_/| $$__  $$|__  $$__//$$__  $$
| $$  \__/  | $$  | $$$$  /$$$$| $$  | $$| $$      | $$  \ $$| $$  \__/  | $$        | $$ /$$/ | $$  \ $$  | $$  | $$  \ $$   | $$  | $$  \ $$
|  $$$$$$   | $$  | $$ $$/$$ $$| $$  | $$| $$      | $$$$$$$$|  $$$$$$   | $$        | $$$$$/  | $$$$$$$/  | $$  | $$$$$$$/   | $$  | $$  | $$
 \____  $$  | $$  | $$  $$$| $$| $$  | $$| $$      | $$__  $$ \____  $$  | $$        | $$  $$  | $$__  $$  | $$  | $$____/    | $$  | $$  | $$
 /$$  \ $$  | $$  | $$\  $ | $$| $$  | $$| $$      | $$  | $$ /$$  \ $$  | $$        | $$\  $$ | $$  \ $$  | $$  | $$         | $$  | $$  | $$
|  $$$$$$/ /$$$$$$| $$ \/  | $$|  $$$$$$/| $$$$$$$$| $$  | $$|  $$$$$$/ /$$$$$$      | $$ \  $$| $$  | $$ /$$$$$$| $$         | $$  |  $$$$$$/
 \______/ |______/|__/     |__/ \______/ |________/|__/  |__/ \______/ |______/      |__/  \__/|__/  |__/|______/|__/         |__/   \______/ 
                                                                                                                                              
                                                                                                                                              
                                                                                                                                              
`
	
	if *first {
		fmt.Print("\n\n\n", asciiMenu)
		*first = false
	}
	
	fmt.Printf("\nSALDO SAAT INI: $%.2f\n\n", s)
	fmt.Println("1. BELI ASET")
	fmt.Println("2. JUAL ASET")
	fmt.Println("3. RIWAYAT TRANSAKSI")
	fmt.Println("4. SORTING")
	fmt.Println("5. ASET PENGGUNA")
	fmt.Println("6. LIST KOIN")
	fmt.Println("7. ADMIN")
	fmt.Println("8. CARI ASET")
	fmt.Println("9. EXIT\n")
	fmt.Print("PILIH OPSI DARI MENU: ")
}

func tambahAset(cDK *int, dK *arrAsetK) {
	var idxNew int

	idxNew = -1
	fmt.Println("\n[TAMBAH ASET]")
	fmt.Print("INPUT NAMA ASET: ")
	fmt.Scan(&dK[*cDK].nama)
	idxNew = cariIndexK(*cDK, *dK, dK[*cDK].nama)
	
	if idxNew != -1 {
		fmt.Println("[INVALID] [ASET SUDAH ADA]")
	} else {
		fmt.Print("INPUT HARGA ASET: ")
		fmt.Scan(&dK[*cDK].harga)
		fmt.Print("INPUT KAPITALISASI ASET: ")
		fmt.Scan(&dK[*cDK].kapitalisasi)
		fmt.Println("[TAMBAH ASET BERHASIL]")
		*cDK++
	}
}

func hapusAset(cDK, cDP *int, dK *arrAsetK, dP *arrAsetP) {
	var i, idxHapusK, idxHapusP int
	var hapus string
	
	fmt.Println("\n[HAPUS ASET]")
	fmt.Print("INPUT NAMA ASET YANG INGIN DIHAPUS: ")
	fmt.Scan(&hapus)
	idxHapusK = cariIndexK(*cDK, *dK, hapus)
	
	if idxHapusK == -1 {
		fmt.Println("[INVALID] [ASET TIDAK DITEMUKAN]")
	} else {
		for i = idxHapusK; i < *cDK-1; i++ {
			dK[i].nama = dK[i+1].nama
			dK[i].harga = dK[i+1].harga
			dK[i].kapitalisasi = dK[i+1].kapitalisasi
		}
		
		idxHapusP = cariIndexP(*cDP, *dP, hapus)
		*cDK--
		
		if idxHapusP != -1 {
			for i = idxHapusP; i < *cDP-1; i++ {
				dP[i].nama = dP[i+1].nama
				dP[i].harga = dP[i+1].harga
				dP[i].jumlah = dP[i+1].jumlah
			}
			
			*cDP--
		}
		
		fmt.Println("[HAPUS ASET BERHASIL]")
	}
}

func ubahAset(cDK, cDP int, dK *arrAsetK, dP *arrAsetP) {
	var idxUbahK, idxUbahP int
	var ubah string

	fmt.Println("\n[UBAH ASET]")
	fmt.Print("INPUT NAMA ASET YANG INGIN DIUBAH: ")
	fmt.Scan(&ubah)
	idxUbahK = cariIndexK(cDK, *dK, ubah)
	idxUbahP = cariIndexP(cDP, *dP, ubah)
	
	if idxUbahK == -1 {
		fmt.Println("[INVALID] [ASET TIDAK DITEMUKAN]")
	} else {
		fmt.Print("UBAH NAMA ASET: ")
		fmt.Scan(&dK[idxUbahK].nama)
		fmt.Print("UBAH HARGA ASET: ")
		fmt.Scan(&dK[idxUbahK].harga)
		fmt.Print("UBAH KAPITALISASI ASET: ")
		fmt.Scan(&dK[idxUbahK].kapitalisasi)
		fmt.Println("[UBAH ASET BERHASIL]")
	}
	
	if idxUbahP != -1 {
		dP[idxUbahP].nama = dK[idxUbahK].nama
		dP[idxUbahP].harga = dK[idxUbahK].harga
	}
}

func beliAset(cDK int, cDP, cR *int, dK arrAsetK, dP *arrAsetP, r *arrTransaksi, s *float64) {
	var idxBeliK, idxBeliP int
	var pembelian string
	var jumlahPembelian float64
	
	idxBeliK = -1
	idxBeliP = -1
	fmt.Printf("\n[BELI ASET] [SALDO: %.2f]\n", *s)
	
	if cDK == 0 {
		fmt.Println("[INVALID] [BELUM TERSEDIA ASET UNTUK DIBELI]")
	} else {
		fmt.Printf("PILIH ASET YANG INGIN DIBELI: ")
		fmt.Scan(&pembelian)
		idxBeliK = cariIndexK(cDK, dK, pembelian)
		idxBeliP = cariIndexP(*cDP, *dP, pembelian)
		
		if idxBeliK == -1 {
			fmt.Println("[INVALID] [ASET TIDAK DITEMUKAN]")
		} else {
			fmt.Printf("INPUT JUMLAH %s YANG INGIN DIBELI: ", pembelian)
			fmt.Scan(&jumlahPembelian)
			
			if *s < (jumlahPembelian * dK[idxBeliK].harga) {
				fmt.Println("[INVALID] [SALDO TIDAK CUKUP]")
			} else {
				*s -= jumlahPembelian * dK[idxBeliK].harga
				
				if idxBeliP == -1 {
					dP[*cDP].nama = pembelian
					dP[*cDP].jumlah = jumlahPembelian
					dP[*cDP].harga = dK[idxBeliK].harga
					*cDP++
				} else {
					dP[idxBeliP].jumlah += jumlahPembelian
				}
				
				tambahRiwayatBeli(cR, jumlahPembelian, pembelian, r)
				fmt.Printf("[BELI ASET BERHASIL] [SALDO: %.2f]\n", *s)
			}
		}
	}
}

func jualAset(cDP *int, cR *int, dP *arrAsetP, r *arrTransaksi, s *float64) {
	var i, idxJual, idxHapusP int
	var penjualan string
	var jumlahPenjualan float64
	var isEmpty bool
	
	idxJual = -1
	fmt.Printf("\n[JUAL ASET] [SALDO: %.2f]", *s)
	
	if *cDP == 0 {
		fmt.Println("\n[INVALID] [BELUM TERSEDIA ASET UNTUK DIJUAL]")
		isEmpty = true
	}
	
	if !isEmpty {
		fmt.Print("\nPILIH ASET UNTUK DIJUAL: ")
		fmt.Scan(&penjualan)
		idxJual = cariIndexP(*cDP, *dP, penjualan)
		
		if idxJual == -1 {
			fmt.Println("[INVALID] [ASET TIDAK DITEMUKAN]")
		} else {
			fmt.Printf("INPUT JUMLAH %s YANG INGIN DIJUAL: ", penjualan)
			fmt.Scan(&jumlahPenjualan)
			
			if jumlahPenjualan > dP[idxJual].jumlah {
				fmt.Println("[INVALID] [JUMLAH TIDAK DIMILIKI]")
			} else if jumlahPenjualan == dP[idxJual].jumlah {
				*s += dP[idxJual].harga * jumlahPenjualan
				idxHapusP = cariIndexP(*cDP, *dP, penjualan)
				
				for i = idxHapusP; i < *cDP-1; i++ {
					dP[i].nama = dP[i+1].nama
					dP[i].harga = dP[i+1].harga
					dP[i].jumlah = dP[i+1].jumlah
				}
				
				*cDP--
				tambahRiwayatJual(cR, jumlahPenjualan, penjualan, r)
				fmt.Printf("[JUAL ASET BERHASIL] [SALDO: %.2f]\n", *s)
			} else {
				*s += dP[idxJual].harga * jumlahPenjualan
				dP[idxJual].jumlah -= jumlahPenjualan
				tambahRiwayatJual(cR, jumlahPenjualan, penjualan, r)
				fmt.Printf("[JUAL ASET BERHASIL] [SALDO: %.2f]\n", *s)
			}
		}
	}
}

func tampilRiwayat(cR int, r arrTransaksi) {
	if cR <= 0 {
		fmt.Println("[INVALID] [TIDAK ADA TRANSAKSI]")
	} else {
		tabelRiwayat(cR, r)
	}
}

func tampilAsetPengguna(cDP int, dP arrAsetP) {
	if cDP <= 0 {
		fmt.Println("\n[INVALID] [TIDAK ADA ASET]")
	} else {
		tabelKoinP(cDP, dP)
	}
}

func tambahRiwayatBeli(cR *int, jA float64, nA string, r *arrTransaksi) {
	r[*cR].nama = "BELI " + nA
	r[*cR].jumlah = jA
	*cR++
}

func tambahRiwayatJual(cR *int, jA float64, nA string, r *arrTransaksi) {
	r[*cR].nama = "JUAL " + nA
	r[*cR].jumlah = jA
	*cR++
}

func cariIndexP(cDP int, dP arrAsetP, txt string) int {
	var i, idx int
	
	idx = -1
	
	for i = 0; i < cDP; i++ {
		if txt == dP[i].nama {
			idx = i
		}
	}
	
	return idx
}

func cariIndexK(cDK int, dK arrAsetK, txt string) int {
	var i, idx int
	
	idx = -1
	
	for i = 0; i < cDK; i++ {
		if txt == dK[i].nama {
			idx = i
		}
	}
	
	return idx
}

func testData(cDK *int, dK *arrAsetK) {
	dK[0].nama = "BTC"
	dK[0].harga = 101430.75
	dK[0].kapitalisasi = 213031267532645
	dK[1].nama = "ETH"
	dK[1].harga = 2065.89
	dK[1].kapitalisasi = 248543789753
	dK[2].nama = "BNB"
	dK[2].harga = 617.99
	dK[2].kapitalisasi = 87019472447
	dK[3].nama = "SOL"
	dK[3].harga = 160.18
	dK[3].kapitalisasi = 83068492358
	dK[4].nama = "XRP"
	dK[4].harga = 2.2442
	dK[4].kapitalisasi = 131135864789
	dK[5].nama = "ADA"
	dK[5].harga = 0.7333
	dK[5].kapitalisasi = 25883453452
	dK[6].nama = "DOGE"
	dK[6].harga = 0.19076
	dK[6].kapitalisasi = 28412455663
	dK[7].nama = "PEPE"
	dK[7].harga = 0.001038
	dK[7].kapitalisasi = 4385323432
	dK[8].nama = "AVAX"
	dK[8].harga = 21.31
	dK[8].kapitalisasi = 8895321553
	dK[9].nama = "TON"
	dK[9].harga = 3.184
	dK[9].kapitalisasi = 7913255312
	dK[10].nama = "SHIB"
	dK[10].harga = 0.1403
	dK[10].kapitalisasi = 8257248784
	dK[11].nama = "KAITO"
	dK[11].harga = 1.3558
	dK[11].kapitalisasi = 329032842
	dK[12].nama = "AIXBT"
	dK[12].harga = 0.1946
	dK[12].kapitalisasi = 179432312
	*cDK = 13
}

func selectionSort(ascending bool, berdasarkan string, cDK int, dK *arrAsetK) {
	var i, j int
	var minIdx int
	
	for i = 0; i < cDK-1; i++ {
		minIdx = i
		
		for j = i + 1; j < cDK; j++ {
			if berdasarkan == "harga" {
				if ascending && dK[j].harga < dK[minIdx].harga || !ascending && dK[j].harga > dK[minIdx].harga {
					minIdx = j
				}
			} else if berdasarkan == "kapitalisasi" {
				if ascending && dK[j].kapitalisasi < dK[minIdx].kapitalisasi || !ascending && dK[j].kapitalisasi > dK[minIdx].kapitalisasi {
					minIdx = j
				}
			}
		}
		
		dK[i], dK[minIdx] = dK[minIdx], dK[i]
	}
	
	fmt.Print("\n[SELESAI: DATA DIURUTKAN BERDASARKAN ")
	
	if berdasarkan == "harga" {
		fmt.Print("HARGA ")
	} else {
		fmt.Print("KAPITALISASI ")
	}
	
	if ascending {
		fmt.Println("ASCENDING]")
	} else {
		fmt.Println("DESCENDING]")
	}
}

func pilihSorting(cDK int, dK *arrAsetK) {
	var opsi int
	
	fmt.Println("\n[PILIH JENIS SORTING]")
	fmt.Println("1. HARGA (ASCENDING)")
	fmt.Println("2. HARGA (DESCENDING)")
	fmt.Println("3. KAPITALISASI (ASCENDING)")
	fmt.Println("4. KAPITALISASI (DESCENDING)\n")
	fmt.Print("PILIH OPSI: ")
	fmt.Scan(&opsi)
	
	switch opsi {
		case 1:
			selectionSort(true, "harga", cDK, dK)
		case 2:
			selectionSort(false, "harga", cDK, dK)
		case 3:
			selectionSort(true, "kapitalisasi", cDK, dK)
		case 4:
			selectionSort(false, "kapitalisasi", cDK, dK)
		default:
			fmt.Println("\n[INVALID OPTION]")
	}
}

func admin(cDK, cDP *int, dK *arrAsetK, dP *arrAsetP, key string) {
	var inputKey string
	var nAdmin int
	
	fmt.Print("\nINPUT ADMIN KEY: ")
	fmt.Scan(&inputKey)
	
	if inputKey != key {
		fmt.Println("[INVALID] [SALAH KEY]")
	} else {
		fmt.Println("\n1. TAMBAH ASET")
		fmt.Println("2. HAPUS ASET")
		fmt.Println("3. UBAH ASET")
		fmt.Print("\nPILIH OPSI DARI MENU ADMIN: ")
		fmt.Scan(&nAdmin)
		
		switch nAdmin {
			case 1:
				tambahAset(cDK, dK)
			case 2:
				hapusAset(cDK, cDP, dK, dP)
			case 3:
				ubahAset(*cDK, *cDP, dK, dP)
			default:
				fmt.Println("\n[INVALID]")
		}
	}
}

func tabelKoinK(cDK int, dK arrAsetK) {
	fmt.Println("\n[DAFTAR KOIN TERSEDIA]")
	fmt.Println("====================================================")
	fmt.Printf("%-12s | %-15s | %-20s\n", "NAMA", "HARGA", "KAPITALISASI")
	fmt.Println("-------------+-----------------+--------------------")
	
	for i := 0; i < cDK; i++ {
		fmt.Printf("%-12s | $%-14.2f | $%-19.2f\n", dK[i].nama, dK[i].harga, dK[i].kapitalisasi)
	}
	
	fmt.Println("====================================================")
}

func tabelKoinP(cDP int, dP arrAsetP) {
	fmt.Println("\n[DAFTAR KOIN PENGGUNA]")
	fmt.Println("===========================================")
	fmt.Printf("%-12s | %-15s | %-15s\n", "NAMA", "HARGA", "JUMLAH")
	fmt.Println("-------------+-----------------+-----------")
	
	for i := 0; i < cDP; i++ {
		fmt.Printf("%-12s | $%-14.2f | %-14.8f\n", dP[i].nama, dP[i].harga, dP[i].jumlah)
	}
	
	fmt.Println("===========================================")
}

func tabelRiwayat(cR int, r arrTransaksi) {
	fmt.Println("\n[DAFTAR TRANSAKSI]")
	fmt.Println("=========================")
	fmt.Printf("%-12s | %-15s\n", "NAMA", "JUMLAH")
	fmt.Println("-------------+-----------")
	
	for i := 0; i < cR; i++ {
		fmt.Printf("%-12s | %-14.8f\n", r[i].nama, r[i].jumlah)
	}
	
	fmt.Println("=========================")
}

func cariAset(cDK int, dK arrAsetK) {
	var i, j int
	var txt string
	var sama bool
	
	fmt.Print("\nINPUT PENCARIAN: ")
	fmt.Scan(&txt)
	fmt.Println("\n[HASIL PENCARIAN]")
	fmt.Println("====================================================")
	fmt.Printf("%-12s | %-15s | %-20s\n", "NAMA", "HARGA", "KAPITALISASI")
	fmt.Println("-------------+-----------------+--------------------")
	
	for i = 0; i < cDK; i++ {
		sama = true
		
		for j = 0; j < len(txt); j++ {
			if txt[j] != dK[i].nama[j] {
				sama = false
			}
		}
		
		if sama {
			fmt.Printf("%-12s | $%-14.2f | $%-19.2f\n", dK[i].nama, dK[i].harga, dK[i].kapitalisasi)
		}
	}
	
	fmt.Println("====================================================")
}
