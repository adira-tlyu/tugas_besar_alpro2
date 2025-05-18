package main

import "fmt"

const NMAX int = 1024

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

type arrAsetK [NMAX - 1]daftarAsetKripto
type arrAsetP [NMAX - 1]daftarAsetPengguna
type arrTransaksi [NMAX - 1]riwayatTransaksi

func main() {
	var daftarK arrAsetK
	var daftarP arrAsetP
	var riwayat arrTransaksi
	var s float64 = 100000
	var n, countDaftarK, countDaftarP, countRiwayat int

	testData(&daftarK, &countDaftarK)

	for n < 7 {
		showMenu(s)
		fmt.Scan(&n)
		
		switch n {
		case 1:
			beliAset(countDaftarK, &countDaftarP, &countRiwayat, &s, daftarK, &daftarP, &riwayat)
		case 2:
			jualAset(&countDaftarP, &countRiwayat, &daftarP, &s, &riwayat)
		case 3:
			tampilRiwayat(riwayat, countRiwayat)
		case 4:
			pilihSorting(&daftarK, countDaftarK)
		case 5:
			tampilAsetPengguna(countDaftarP, daftarP)
		case 6:
			listKoin(daftarK, countDaftarK, true)
		case 7:
			fmt.Println("\n\t\t 	[PROGRAM SELESAI]")
		default:
			fmt.Println("\n\t\t     [INVALID]")
		}
	}
}

func showMenu(s float64) {
	fmt.Println("\n\t\t      ░▒▓██████▓▒░░▒▓███████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓███████▓▒░▒▓████████▓▒░▒▓██████▓▒░  ")
	fmt.Println("\t\t     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░ ")
	fmt.Println("\t\t     ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░ ")
	fmt.Println("\t\t     ░▒▓█▓▒░      ░▒▓███████▓▒░ ░▒▓██████▓▒░░▒▓███████▓▒░  ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░ ")
	fmt.Println("\t\t     ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒░        ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░ ")
	fmt.Println("\t\t     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒░        ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░ ")
	fmt.Println("\t\t      ░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒░        ░▒▓█▓▒░   ░▒▓██████▓▒░  ")
	fmt.Printf("\n\t\t     SALDO SAAT INI: $%.2f\n", s)
	fmt.Println("\t\t     1. BELI ASET")
	fmt.Println("\t\t     2. JUAL ASET")
	fmt.Println("\t\t     3. RIWAYAT TRANSAKSI")
	fmt.Println("\t\t     4. SORTING")
	fmt.Println("\t\t     5. ASET PENGGUNA")
	fmt.Println("\t\t     6. LIST KOIN")
	fmt.Println("\t\t     7. EXIT\n")
	fmt.Print("\t\t     PILIH OPSI DARI MENU: ")
}

func beliAset(cDK int, cDP, cR *int, s *float64, dK arrAsetK, dP *arrAsetP, r *arrTransaksi) {
	var idxBeliK, idxBeliP int
	var pembelian string
	var jumlahPembelian float64

	idxBeliK = -1
	idxBeliP = -1
	fmt.Printf("\n\t\t     [BELI ASET] [SALDO: %f]\n", *s)
	
	if cDK == 0 {
		fmt.Println("\t\t     [INVALID] [BELUM TERSEDIA ASET UNTUK DIBELI]")
	} else {
		fmt.Printf("\t\t     PILIH ASET YANG INGIN DIBELI: ")
		fmt.Scan(&pembelian)
		idxBeliK = cariIndexK(pembelian, dK, cDK)
		idxBeliP = cariIndexP(pembelian, *dP, *cDP)
		
		if idxBeliK == -1 {
			fmt.Println("\t\t     [INVALID] [ASET TIDAK DITEMUKAN]")
		} else {
			fmt.Printf("\t\t     INPUT JUMLAH %s YANG INGIN DIBELI: ", pembelian)
			fmt.Scan(&jumlahPembelian)
			
			if *s < (jumlahPembelian * dK[idxBeliK].harga) {
				fmt.Println("\t\t     [INVALID] [SALDO TIDAK CUKUP]")
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
				
				tambahRiwayatBeli(r, cR, pembelian, jumlahPembelian)
				fmt.Printf("\t\t     [BELI ASET BERHASIL] [SALDO: %f]\n", *s)
			}
		}
	}
}

func jualAset(cDP *int, cR *int, dP *arrAsetP, s *float64, r *arrTransaksi) {
	var i, idxJual, idxHapusP int
	var penjualan string
	var jumlahPenjualan float64
	var isEmpty bool

	idxJual = -1
	fmt.Printf("\n\t\t     [JUAL ASET] [SALDO: %f]", *s)
	
	if *cDP == 0 {
		fmt.Println("\n\t\t     [INVALID] [BELUM TERSEDIA ASET UNTUK DIJUAL]")
		isEmpty = true
	}
	
	if !isEmpty {
		fmt.Print("\n\t\t     PILIH ASET UNTUK DIJUAL: ")
		fmt.Scan(&penjualan)
		idxJual = cariIndexP(penjualan, *dP, *cDP)
		
		if idxJual == -1 {
			fmt.Println("\t\t     [INVALID] [ASET TIDAK DITEMUKAN]")
		} else {
			fmt.Printf("\t\t     INPUT JUMLAH %s YANG INGIN DIJUAL: ", penjualan)
			fmt.Scan(&jumlahPenjualan)
			
			if jumlahPenjualan > dP[idxJual].jumlah {
				fmt.Println("\t\t     [INVALID] [JUMLAH TIDAK DIMILIKI]")
			} else if jumlahPenjualan == dP[idxJual].jumlah {
				*s += dP[idxJual].harga * jumlahPenjualan
				idxHapusP = cariIndexP(penjualan, *dP, *cDP)
				
				for i = idxHapusP; i < *cDP-1; i++ {
					dP[i].nama = dP[i+1].nama
					dP[i].harga = dP[i+1].harga
					dP[i].jumlah = dP[i+1].jumlah
				}
				
				*cDP--
				tambahRiwayatJual(r, cR, penjualan, jumlahPenjualan)
				fmt.Printf("\t\t     [JUAL ASET BERHASIL] [SALDO: %f]\n", *s)
			} else {
				*s += dP[idxJual].harga * jumlahPenjualan
				dP[idxJual].jumlah -= jumlahPenjualan
				tambahRiwayatJual(r, cR, penjualan, jumlahPenjualan)
				fmt.Printf("\t\t     [JUAL ASET BERHASIL] [SALDO: %f]\n", *s)
			}
		}
	}
}

func tampilRiwayat(r arrTransaksi, cR int) {
	var i int

	fmt.Println("\n\t\t     [RIWAYAT TRANSAKSI]")
	
	if cR <= 0 {
		fmt.Println("\t\t     [INVALID] [TIDAK ADA TRANSAKSI]")
	} else {
		fmt.Println("\t\t     TRANSAKSI\t\tJUMLAH")
		
		for i = 0; i < cR; i++ {
			fmt.Printf("\t\t     %s", r[i].nama)
			fmt.Printf("\t\t%f\n", r[i].jumlah)
		}
		
		fmt.Println("\t\t     [RIWAYAT TRANSAKSI]")
	}
}

func tampilAsetPengguna(cDP int, dP arrAsetP) {
	var i int

	fmt.Println("\n\t\t     [ASET PENGGUNA]")
	
	if cDP <= 0 {
		fmt.Println("\t\t     [INVALID] [TIDAK ADA ASET]")
	} else {
		fmt.Println("\t\t     NAMA ASET\t\tJUMLAH")
		
		for i = 0; i < cDP; i++ {
			fmt.Printf("\t\t     %s", dP[i].nama)
			fmt.Printf("\t\t%f\n", dP[i].jumlah)
		}
		
		fmt.Println("\t\t     [ASET PENGGUNA]")
	}
}

func tambahRiwayatBeli(r *arrTransaksi, cR *int, nA string, jA float64) {
	r[*cR].nama = "BELI " + nA
	r[*cR].jumlah = jA
	*cR++
}

func tambahRiwayatJual(r *arrTransaksi, cR *int, nA string, jA float64) {
	r[*cR].nama = "JUAL " + nA
	r[*cR].jumlah = jA
	*cR++
}

func cariIndexP(txt string, dP arrAsetP, cDP int) int {
	var i, idx int

	idx = -1
	
	for i = 0; i < cDP; i++ {
		if txt == dP[i].nama {
			idx = i
		}
	}
	
	return idx
}

func cariIndexK(txt string, dK arrAsetK, cDK int) int {
	var i, idx int

	idx = -1
	
	for i = 0; i < cDK; i++ {
		if txt == dK[i].nama {
			idx = i
		}
	}
	
	return idx
}

func testData(dK *arrAsetK, cDK *int) {
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

func selectionSort(dK *arrAsetK, cDK int, berdasarkan string, ascending bool) {
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
	
	fmt.Print("\n\t\t     [SELESAI: DATA DIURUTKAN BERDASARKAN ")
	
	if berdasarkan == "harga" {
		fmt.Print("HARGA ")
	} else {
		fmt.Print("KAPITALISASI ")
	}
	
	if ascending {
		fmt.Println("TERKECIL -> TERBESAR]")
	} else {
		fmt.Println("TERBESAR -> TERKECIL]")
	}
}

func listKoin(dK arrAsetK, cDK int, ascending bool) {
	var i, j int
	var temp daftarAsetKripto

	for i = 1; i < cDK; i++ {
		temp = dK[i]
		j = i - 1
		
		if ascending {
			for j >= 0 && dK[j].kapitalisasi > temp.kapitalisasi {
				dK[j+1] = dK[j]
				j--
			}
		} else {
			for j >= 0 && dK[j].kapitalisasi < temp.kapitalisasi {
				dK[j+1] = dK[j]
				j--
			}
		}
		
		dK[j+1] = temp
	}
	
	fmt.Println("\n\t\t     [DAFTAR KOIN TERSEDIA]")
	fmt.Println("\t\t     ================================================")
	fmt.Printf("\t\t     %-12s | %-15s | %-20s\n", "Nama", "Harga", "Kapitalisasi")
	fmt.Println("\t\t     -----------------------------------------------")
	
	for i := 0; i < cDK; i++ {
		fmt.Printf("\t\t     %-12s | $%-14.2f | $%-19.2f\n", dK[i].nama, dK[i].harga, dK[i].kapitalisasi)
	}
	
	fmt.Println("\t\t     ================================================")
}

func pilihSorting(dK *arrAsetK, cDK int) {
	var opsi int
	
	fmt.Println("\n\t\t     [PILIH JENIS SORTING]")
	fmt.Println("\t\t     1. Harga (Terkecil -> Terbesar)")
	fmt.Println("\t\t     2. Harga (Terbesar -> Terkecil)")
	fmt.Println("\t\t     3. Kapitalisasi (Terkecil -> Terbesar)")
	fmt.Println("\t\t     4. Kapitalisasi (Terbesar -> Terkecil)")
	fmt.Print("\t\t     Pilih opsi: ")
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		selectionSort(dK, cDK, "harga", true)
	case 2:
		selectionSort(dK, cDK, "harga", false)
	case 3:
		selectionSort(dK, cDK, "kapitalisasi", true)
	case 4:
		selectionSort(dK, cDK, "kapitalisasi", false)
	default:
		fmt.Println("\t\t     [INVALID OPTION]")
	}
}
