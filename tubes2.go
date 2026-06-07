package main

import "fmt"

const NMAX = 1000

type warga struct {
	id, nama, tanggal, jenis string
	berat float64
}
type arrWarga [NMAX]warga

func main(){
	var data arrWarga
	var banyakW, menu int
	var cek string
	var hasil int
	
	banyakW = 0
	fmt.Println("selamat datang di aplikasi Sampahku")
	fmt.Println("apakah ingin menggunakan aplikasi(YES/NO)")
	fmt.Scan(&cek)
		if cek != "YES" && cek != "NO" {
			fmt.Println("\nmines literasi")
		}else if cek == "NO" {
			fmt.Println("\nterima kasih")
		}else{	
		for cek != "NO"   {
			fmt.Println(" ")
			fmt.Println("DAFTAR MENU")
			fmt.Println("1. Tambah Data Warga & Setoran")
			fmt.Println("2. Tampilkan Semua Data Warga")
			fmt.Println("3. Ubah Data Warga")
			fmt.Println("4. Hapus Data Warga")
			fmt.Println("5. Urutkan ID Ascending (Selection Sort)")
			fmt.Println("6. Urutkan Berat Descending (Insertion Sort)")
			fmt.Println("7. Cari Warga Berdasarkan Nama (Sequential)")
			fmt.Println("8. Cari Warga Berdasarkan ID (Binary)")
			fmt.Println("9. Tampilkan Statistik Mingguan")
			fmt.Println("0. Keluar Aplikasi")
			fmt.Print("Pilih menu (0-9): ")
			fmt.Scan(&menu)

		switch menu {
		case 1:
			inputData(&data, &banyakW)
		case 2:
			outputData(&data, banyakW)
		case 3:
			ubahData(&data, banyakW)
		case 4: 
			removeData(&data, &banyakW)
		case 5:
			selecSortSampah(&data, banyakW)
			fmt.Println("\ndata setelah diurutkan")
			outputData(&data, banyakW)
		case 6:
			insertSortSampah(&data, banyakW)
			fmt.Println("\ndata setelah diurutkan")
			outputData(&data, banyakW)
		case 7:
			hasil = seqSearchNama(data, banyakW)
			if hasil != -1 {
				fmt.Println("Tanggal :", data[hasil].tanggal)
        		fmt.Println("ID      :", data[hasil].id)
        		fmt.Println("Nama    :", data[hasil].nama)
        		fmt.Println("Berat   :", data[hasil].berat, "kg")
        		fmt.Println("Jenis   :", data[hasil].jenis)
			} else {
				fmt.Println("Tidak ditemukan")
			}	
		case 8:
			selecSortSampah(&data, banyakW)
			hasil = binSearchId(data, banyakW)
			if hasil != -1 {
				fmt.Println("Tanggal :", data[hasil].tanggal)
        		fmt.Println("ID      :", data[hasil].id)
        		fmt.Println("Nama    :", data[hasil].nama)
        		fmt.Println("Berat   :", data[hasil].berat, "kg")
        		fmt.Println("Jenis   :", data[hasil].jenis)
			} else {
				fmt.Println("Tidak ditemukan")
			}
		
		case 9:
			insertSortSampah(&data, banyakW)
			statistikSampah(&data, banyakW)
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi Sampahku!")
			cek = "NO"
		}
		}
		}	
}

func inputData(data *arrWarga, banyakW *int){
	var i int
	var lanjut string

	i = *banyakW
	lanjut = "yes"
	for i < NMAX && lanjut != "no" && lanjut != "NO" {
		fmt.Print("Masukkan Tanggal Transaksi (DD-MM-YYYY): ")
		fmt.Scan(&data[i].tanggal)
		fmt.Print("Masukkan ID Warga (5 Digit): ")
		fmt.Scan(&data[i].id)
		for len(data[i].id) != 5 {
			fmt.Println("Eorr")
			fmt.Print("Masukkan ID Warga: ")
			fmt.Scan(&data[i].id)
		}
		fmt.Print("Masukkan Nama Warga: ")
		fmt.Scan(&data[i].nama)
		fmt.Print("Masukkan Berat Sampah (kg): ")
		fmt.Scan(&data[i].berat)
		fmt.Print("Masukkan jenis sampah(organik/anorganik) : ")
		fmt.Scan(&data[i].jenis)
		i = i + 1
		fmt.Println("Data berhasil disimpan")
		fmt.Println("\nApakah ingin menambahkan data lagi? (yes/no)")
		fmt.Scan(&lanjut)
	}
	*banyakW = i
}

func outputData( data *arrWarga, banyakW int){
	var i int
	
	fmt.Println("\nDaftar Warga")
	 for i = 0; i < banyakW; i++ {
        fmt.Println("\nTanggal :", data[i].tanggal)
        fmt.Println("ID      :", data[i].id)
        fmt.Println("Nama    :", data[i].nama)
        fmt.Println("Berat   :", data[i].berat, "kg")
        fmt.Println("Jenis   :", data[i].jenis)
        fmt.Println("---") 
    }
}

func selecSortSampah(data *arrWarga, banyakW int){
	var pass, idx, i int
	var temp warga
	
	pass = 1
	for pass <= banyakW-1 {
		idx = pass - 1
		i = pass
		for i < banyakW {
			if data[idx].id > data[i].id{
				idx = i
			}
			i = i + 1
		}
		temp = data[pass-1]
		data[pass-1] = data[idx]
		data[idx] = temp
		pass = pass + 1
	}
}

func insertSortSampah(data *arrWarga, banyakW int){
	var pass, i int
	var temp warga
	
	pass = 1
	for pass <= banyakW-1 {
		i = pass
		temp = data[pass]
		for i > 0 && temp.berat > data[i-1].berat {
			data[i] = data[i-1]
			i = i - 1
		}
		data[i] = temp
		pass = pass + 1
	}
}

func seqSearchNama(data arrWarga, banyakW int) int {
	var found, i int
	var target string
	
	fmt.Print("Target nama:")
	fmt.Scan(&target)
	found = -1
	i = 0
	for found == -1 && i < banyakW {
		if data[i].nama == target {
			found = i
		}
		i = i+1
	}
	return found
}

func binSearchId(data arrWarga, banyakW int) int {
	var left, right, mid, found int
	var x string
	
	fmt.Print("Target ID: ")
	fmt.Scan(&x)
	left = 0
	right = banyakW - 1
	found = -1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if x < data[mid].id {
			right = mid - 1
		} else if x > data[mid].id{
			left = mid + 1
		} else {
			found = mid
		}
		mid = (right+left)/2
	}
	return found
}

func removeData(data *arrWarga, banyakW *int) {
	var i, found int
	found = binSearchId(*data, *banyakW)
	if found == -1 {
		fmt.Println("\nID tidak ditemukan")
	} else {
		i = found
		for i <= *banyakW-2 {
			data[i] = data[i+1]
			i = i + 1
		}
		*banyakW = *banyakW - 1
		fmt.Println("Data berhasil di hapus")
	}
}

func ubahData(data *arrWarga, banyakW int) {
	var found int

	found = seqSearchNama(*data, banyakW)
	if found == -1 {
		fmt.Println("Id tidak ditemukan")
	} else {
		fmt.Print("Masukkan Tanggal Transaksi (DD-MM-YYYY): ")
		fmt.Scan(&data[found].tanggal)
		fmt.Print("Masukkan ID Warga: ")
		fmt.Scan(&data[found].id)
		fmt.Print("Masukkan Nama Warga: ")
		fmt.Scan(&data[found].nama)
		fmt.Print("Masukkan Berat Sampah (kg): ")
		fmt.Scan(&data[found].berat)
		fmt.Print("Masukkan jenis sampah(organik/anorganik) : ")
		fmt.Scan(&data[found].jenis)
	}
}

func minMaxSampah( data *arrWarga, banyakW int){
	fmt.Println("\nWarga dengan berat sampah paling banyak: ", data[0].nama, "| berat: ", data[0].berat, "Kg")
	fmt.Println("Warga dengan berat sampah paling sedikit: ", data[banyakW-1].nama, "| berat: ", data[banyakW-1].berat, "Kg")
}

func statistikSampah(data *arrWarga, banyakW int){
	var i int
	var totalOr, totalAn, total float64

	total = 0
	totalOr = 0
	totalAn = 0

	for i = 0; i<banyakW; i++{
		if data[i].jenis == "organik" {
			totalOr = totalOr + data[i].berat
		} else if data[i].jenis == "anorganik" {
			totalAn = totalAn + data[i].berat
		}
		total = total + data[i].berat
	}

	fmt.Println("\nStatistik Sampah Mingguan")
	fmt.Println("\nJumlah warga: ", banyakW)
	
	minMaxSampah(data, banyakW)

	fmt.Printf("\nTotal sampah keseluruhan   : %.2f kg\n", total)
	fmt.Printf("Rata-rata berat mingguan   : %.2f kg\n", total/float64(banyakW))
	fmt.Printf("Total sampah organik       : %.2f kg\n", totalOr)
	fmt.Printf("Total sampah anorganik     : %.2f kg\n", totalAn)
}
