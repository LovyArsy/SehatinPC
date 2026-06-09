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
		Komponen{"GPU001", "RTX 3060", 85, 70, hitungStatus(85, 70)},
		Komponen{"RAM001", "DDR4 16GB", 45, 95, hitungStatus(45, 95)},
		Komponen{"SSD001", "Kingston SSD 500GB", 90, 95, hitungStatus(90, 95)},
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
	fmt.Println("5. Keluar")
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

	dataBaru.NoSeri = inputString("No Seri      : ")
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