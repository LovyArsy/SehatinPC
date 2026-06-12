package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Komponen struct {
	NoSeri     string
	Nama       string
	Suhu       int
	BebanKerja int
	Status     string
}

var dataKomponen []Komponen
var reader = bufio.NewReader(os.Stdin)

func main() {

	dataKomponen = append(dataKomponen,
		Komponen{"CPU001", "Intel i5-12400", 75, 60, hitungStatus(75, 60)},
		Komponen{"CPU002", "Ryzen 5 5600G", 70, 50, hitungStatus(70, 50)},

		Komponen{"RAM001", "DDR4 16GB", 45, 95, hitungStatus(45, 95)},
		Komponen{"RAM002", "DDR5 32GB", 50, 92, hitungStatus(50, 92)},

		Komponen{"GPU001", "RTX 3060", 85, 70, hitungStatus(85, 70)},
		Komponen{"GPU002", "GTX 1660 Super", 90, 80, hitungStatus(90, 80)},

		Komponen{"SSD001", "Kingston SSD 500GB", 90, 95, hitungStatus(90, 95)},
		Komponen{"SSD002", "Samsung SSD 1TB", 85, 98, hitungStatus(85, 98)},
	)

	for {

		tampilkanMenu()

		pilihan := inputInt("Pilih Menu : ")

		switch pilihan {

		case 1:
			tampilkanKomponen()

		case 2:
			tambahKomponen()

		case 3:
			editKomponen()

		case 4:
			hapusKomponen()
		
		case 5:
			cariStatus()
		
		case 6:
			cariNoSeri()
		
		case 7:
			selectionSort()
			fmt.Println("\nData berhasil diurutkan berdasarkan No Seri (Asc).")
		
		case 8:
			insertionSort()
			fmt.Println("\nData berhasil diurutkan berdasarkan No Seri (Asc).")
		
		case 9:
			tampilkanStatistik()

		case 10:
			fmt.Println("\nProgram selesai.")
			return

		default:
			fmt.Println("\nMenu tidak tersedia.")
		}
	}
}

func tampilkanMenu() {

	fmt.Println("\n===== SEHATIN PC =====")
	fmt.Println("1. Lihat Komponen")
	fmt.Println("2. Tambah Komponen")
	fmt.Println("3. Edit Komponen")
	fmt.Println("4. Hapus Komponen")
	fmt.Println("5. Cari Status (Sequential Search)")
	fmt.Println("6. Cari No Seri (Binary Search)")
	fmt.Println("7. Urutkan No Seri Asc (Selection Sort)")
	fmt.Println("8. Urutkan No Seri Asc (Insertion Sort)")
	fmt.Println("9. Statistik Komponen")
	fmt.Println("10. Keluar")
	fmt.Println("======================")
}

func inputString(pesan string) string {

	fmt.Print(pesan)

	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func inputInt(pesan string) int {

	for {

		fmt.Print(pesan)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		angka, err := strconv.Atoi(input)

		if err == nil {
			return angka
		}

		fmt.Println("Input harus berupa angka!")
	}
}

func hitungStatus(suhu int, beban int) string {

	if suhu > 80 && beban > 90 {
		return "Lag & Overheat"
	}

	if suhu > 80 {
		return "Overheat"
	}

	if beban > 90 {
		return "Lag"
	}

	return "Normal"
}

func tampilkanKomponen() {

	fmt.Println("\n===== DATA KOMPONEN =====")

	if len(dataKomponen) == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	for i := 0; i < len(dataKomponen); i++ {

		fmt.Printf(
			"\n%d\nNo Seri      : %s\nNama         : %s\nSuhu         : %d°C\nBeban Kerja  : %d%%\nStatus       : %s\n",
			i+1,
			dataKomponen[i].NoSeri,
			dataKomponen[i].Nama,
			dataKomponen[i].Suhu,
			dataKomponen[i].BebanKerja,
			dataKomponen[i].Status,
		)
	}
}

func tambahKomponen() {

	var dataBaru Komponen

	fmt.Println("\n===== TAMBAH KOMPONEN =====")

	for {
		dataBaru.NoSeri = inputString("No Seri      : ")

		if !noSeriSudahAda(dataBaru.NoSeri) {
			break
		}

		fmt.Println("No Seri sudah terdaftar, silakan masukkan No Seri lain.")
	}

	dataBaru.Nama = inputString("Nama         : ")
	dataBaru.Suhu = inputInt("Suhu         : ")
	dataBaru.BebanKerja = inputInt("Beban Kerja  : ")

	dataBaru.Status = hitungStatus(
		dataBaru.Suhu,
		dataBaru.BebanKerja,
	)

	dataKomponen = append(dataKomponen, dataBaru)

	fmt.Println("\nData berhasil ditambahkan.")
}

func noSeriSudahAda(noSeri string) bool {

	for i := 0; i < len(dataKomponen); i++ {
		if dataKomponen[i].NoSeri == noSeri {
			return true
		}
	}

	return false
}

func editKomponen() {

	fmt.Println("\n===== EDIT KOMPONEN =====")

	noSeri := inputString("Masukkan No Seri : ")

	for i := 0; i < len(dataKomponen); i++ {

		if dataKomponen[i].NoSeri == noSeri {

			fmt.Println("\nData ditemukan.")

			dataKomponen[i].Nama = inputString("Nama Baru         : ")
			dataKomponen[i].Suhu = inputInt("Suhu Baru         : ")
			dataKomponen[i].BebanKerja = inputInt("Beban Kerja Baru  : ")

			dataKomponen[i].Status = hitungStatus(
				dataKomponen[i].Suhu,
				dataKomponen[i].BebanKerja,
			)

			fmt.Println("\nData berhasil diperbarui.")
			return
		}
	}

	fmt.Println("\nData tidak ditemukan.")
}

func hapusKomponen() {

	fmt.Println("\n===== HAPUS KOMPONEN =====")

	noSeri := inputString("Masukkan No Seri : ")

	for i := 0; i < len(dataKomponen); i++ {

		if dataKomponen[i].NoSeri == noSeri {

			dataKomponen = append(
				dataKomponen[:i],
				dataKomponen[i+1:]...,
			)

			fmt.Println("\nData berhasil dihapus.")
			return
		}
	}

	fmt.Println("\nData tidak ditemukan.")
}

//Sequential
func cariStatus() {

	fmt.Println("\n===== CARI STATUS =====")
	fmt.Println("1. Normal")
	fmt.Println("2. Lag")
	fmt.Println("3. Overheat")
	fmt.Println("4. Lag & Overheat")

	pilihan := inputInt("Pilih Status : ")

	var status string

	switch pilihan {
	case 1:
		status = "Normal"
	case 2:
		status = "Lag"
	case 3:
		status = "Overheat"
	case 4:
		status = "Lag & Overheat"
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	sequentialSearchStatus(status)
}

func sequentialSearchStatus(status string) {

	ditemukan := false

	for i := 0; i < len(dataKomponen); i++ {

		if dataKomponen[i].Status == status {

			fmt.Printf(
				"\nNo Seri      : %s\nNama         : %s\nSuhu         : %d°C\nBeban Kerja  : %d%%\nStatus       : %s\n",
				dataKomponen[i].NoSeri,
				dataKomponen[i].Nama,
				dataKomponen[i].Suhu,
				dataKomponen[i].BebanKerja,
				dataKomponen[i].Status,
			)

			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("\nData tidak ditemukan.")
	}
}

//Binary
func cariNoSeri() {

	fmt.Println("\n===== BINARY SEARCH =====")

	noSeri := inputString("Masukkan No Seri : ")

	selectionSort()

	index := binarySearchNoSeri(noSeri)

	if index == -1 {

		fmt.Println("\nData tidak ditemukan.")
		return
	}

	fmt.Println("\nData ditemukan")

	fmt.Printf(
		"No Seri      : %s\nNama         : %s\nSuhu         : %d°C\nBeban Kerja  : %d%%\nStatus       : %s\n",
		dataKomponen[index].NoSeri,
		dataKomponen[index].Nama,
		dataKomponen[index].Suhu,
		dataKomponen[index].BebanKerja,
		dataKomponen[index].Status,
	)
}

func binarySearchNoSeri(noSeri string) int {

	kiri := 0
	kanan := len(dataKomponen) - 1

	for kiri <= kanan {

		tengah := (kiri + kanan) / 2

		if dataKomponen[tengah].NoSeri == noSeri {
			return tengah
		}

		if noSeri < dataKomponen[tengah].NoSeri {

			kanan = tengah - 1

		} else {

			kiri = tengah + 1
		}
	}

	return -1
}

//Selection 
func selectionSort() {

	for i := 0; i < len(dataKomponen)-1; i++ {

		indexTerkecil := i

		for j := i + 1; j < len(dataKomponen); j++ {

			if dataKomponen[j].NoSeri < dataKomponen[indexTerkecil].NoSeri {

				indexTerkecil = j
			}
		}

		dataKomponen[i], dataKomponen[indexTerkecil] = dataKomponen[indexTerkecil], dataKomponen[i]
	}
}

//Insertion
func insertionSort() {

	for i := 1; i < len(dataKomponen); i++ {

		key := dataKomponen[i]

		j := i - 1

		for j >= 0 && dataKomponen[j].NoSeri > key.NoSeri {

			dataKomponen[j+1] = dataKomponen[j]

			j--
		}

		dataKomponen[j+1] = key
	}
}

//Statistik
func tampilkanStatistik() {

	fmt.Println("\n+++ SEHATIN PC +++")

	if len(dataKomponen) == 0 {
		fmt.Println("Belum ada data komponen.")
		return
	}

	totalSuhu := 0
	jumlahNormal := 0
	jumlahBermasalah := 0

	for i := 0; i < len(dataKomponen); i++ {

		totalSuhu += dataKomponen[i].Suhu

		if dataKomponen[i].Status == "Normal" {

			jumlahNormal++

		} else {

			jumlahBermasalah++
		}
	}

	rataRataSuhu := float64(totalSuhu) / float64(len(dataKomponen))

	fmt.Printf("\nJumlah Komponen      : %d\n", len(dataKomponen))
	fmt.Printf("Komponen Normal      : %d\n", jumlahNormal)
	fmt.Printf("Komponen Bermasalah  : %d\n", jumlahBermasalah)
	fmt.Printf("Rata-rata Suhu       : %.2f °C\n", rataRataSuhu)

	fmt.Println("\n++++++++++++++++++++")
}