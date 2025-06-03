package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Peminjam struct {
	Nama        string
	Pinjaman    float64
	Tenor       int
	Bunga       float64
	Cicilan     float64
	StatusBayar string
}

var daftar []Peminjam
var reader = bufio.NewReader(os.Stdin)

func hitungCicilan(pinjaman float64, tenor int, bunga float64) float64 {
	total := pinjaman + (pinjaman * bunga * float64(tenor) / 12)
	return total / float64(tenor)
}

func containsNumber(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}

func tambahPeminjam() {
	var nama string
	var pinjaman float64
	var tenor int
	var reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Masukkan nama lengkap peminjam: ")
		namaInput, _ := reader.ReadString('\n')
		nama = strings.TrimSpace(namaInput)

		if nama == "" {
			fmt.Println("❌ Nama tidak boleh kosong. Silakan masukkan ulang.")
			continue
		} else if containsNumber(nama) {
			fmt.Println("❌ Nama tidak boleh mengandung angka. Silakan masukkan ulang.")
			continue
		}
		break
	}

	for {
		fmt.Print("Masukkan jumlah pinjaman: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		pinjaman, err = strconv.ParseFloat(input, 64)
		if err != nil || pinjaman <= 0 {
			fmt.Println("❌ Jumlah pinjaman harus berupa angka dan lebih dari 0. Silakan masukkan ulang.")
			continue
		}
		break
	}

	for {
		fmt.Print("Masukkan lama pinjaman (dalam bulan): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		tenor, err = strconv.Atoi(input)
		if err != nil || tenor <= 0 {
			fmt.Println("❌ Tenor harus berupa angka dan lebih dari 0. Silakan masukkan ulang.")
			continue
		}
		break
	}

	bunga := 0.05
	cicilan := hitungCicilan(pinjaman, tenor, bunga)
	peminjam := Peminjam{nama, pinjaman, tenor, bunga, cicilan, "Belum Lunas"}
	daftar = append(daftar, peminjam)

	fmt.Printf("✅ Data peminjam berhasil ditambahkan.\n")
	fmt.Printf("Nama: %s | Pinjaman: %.0f | Tenor: %d bulan | Bunga: %.2f | Cicilan: %.0f per bulan\n",
		nama, pinjaman, tenor, bunga, cicilan)
}

func ubahPeminjam() {
	var index int
	var namaCari string

	for {
		fmt.Print("Nama lengkap peminjam yang ingin diperbarui: ")
		namaInput, _ := reader.ReadString('\n')
		namaCari = strings.TrimSpace(namaInput)

		if namaCari == "" {
			fmt.Println("❌ Nama tidak boleh kosong. Silakan masukkan ulang.")
			continue
		}
		if containsNumber(namaCari) {
			fmt.Println("❌ Nama tidak boleh mengandung angka. Silakan masukkan ulang.")
			continue
		}

		namaCariLower := strings.ToLower(namaCari)

		index = -1
		for i, p := range daftar {
			if strings.ToLower(p.Nama) == namaCariLower {
				index = i
				break
			}
		}

		if index == -1 {
			fmt.Println("❌ Data tidak ditemukan. Silakan masukkan nama yang benar.")
		} else {
			break
		}
	}

	var pinjaman float64
	var tenor int

	for {
		fmt.Print("Masukkan jumlah pinjaman baru: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("❌ Input tidak boleh kosong. Silakan coba lagi.")
			continue
		}

		var err error
		pinjaman, err = strconv.ParseFloat(input, 64)
		if err != nil || pinjaman <= 0 {
			fmt.Println("❌ Jumlah pinjaman harus berupa angka dan lebih dari 0. Silakan coba lagi.")
			continue
		}
		break
	}

	for {
		fmt.Print("Masukkan tenor baru (bulan): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("❌ Input tidak boleh kosong. Silakan coba lagi.")
			continue
		}

		var err error
		tenor, err = strconv.Atoi(input)
		if err != nil || tenor <= 0 {
			fmt.Println("❌ Tenor harus berupa angka dan lebih dari 0. Silakan coba lagi.")
			continue
		}
		break
	}

	bunga := 0.05
	cicilan := hitungCicilan(pinjaman, tenor, bunga)

	daftar[index] = Peminjam{
		Nama:        daftar[index].Nama,
		Pinjaman:    pinjaman,
		Tenor:       tenor,
		Bunga:       bunga,
		Cicilan:     cicilan,
		StatusBayar: daftar[index].StatusBayar,
	}

	fmt.Println("\n🔄 Data peminjam berhasil diperbarui.\n")
}

func hapusPeminjam() {
	fmt.Print("Masukkan nama lengkap peminjam yang ingin dihapus: ")
	namaInput, _ := reader.ReadString('\n')
	namaCari := strings.TrimSpace(namaInput)
	namaCariLower := strings.ToLower(namaCari)

	index := -1
	for i, p := range daftar {
		if strings.ToLower(p.Nama) == namaCariLower {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("❌ Peminjam tidak ditemukan.")
		return
	}

	fmt.Print("Apakah Anda yakin ingin menghapus data ini? (ketik: hapus / batalkan): ")
	confirmInput, _ := reader.ReadString('\n')
	confirm := strings.TrimSpace(confirmInput)

	if strings.ToLower(confirm) == "hapus" {
		daftar = append(daftar[:index], daftar[index+1:]...)
		fmt.Println("🗑️  Data berhasil dihapus.")
	} else {
		fmt.Println("✅ Penghapusan dibatalkan.")
	}
}

func sequentialSearch(nama string) *Peminjam {
	nama = strings.ToLower(strings.TrimSpace(nama))

	for i := 0; i < len(daftar); i++ {
		if strings.ToLower(strings.TrimSpace(daftar[i].Nama)) == nama {
			return &daftar[i]
		}
	}
	return nil
}

func sortByNama(data []Peminjam) []Peminjam {
	hasil := make([]Peminjam, len(data))
	copy(hasil, data)

	for i := 0; i < len(hasil)-1; i++ {
		min := i
		for j := i + 1; j < len(hasil); j++ {
			if strings.ToLower(hasil[j].Nama) < strings.ToLower(hasil[min].Nama) {
				min = j
			}
		}
		hasil[i], hasil[min] = hasil[min], hasil[i]
	}
	return hasil
}

func binarySearch(nama string) *Peminjam {
	dataUrut := sortByNama(daftar)

	nama = strings.ToLower(strings.TrimSpace(nama))
	low, high := 0, len(dataUrut)-1

	for low <= high {
		mid := (low + high) / 2
		midNama := strings.ToLower(strings.TrimSpace(dataUrut[mid].Nama))

		if midNama == nama {
			return &dataUrut[mid]
		} else if midNama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}

func selectionSortByPinjaman(ascending bool) {
	for i := 0; i < len(daftar)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(daftar); j++ {
			if ascending && daftar[j].Pinjaman < daftar[minIdx].Pinjaman {
				minIdx = j
			}
			if !ascending && daftar[j].Pinjaman > daftar[minIdx].Pinjaman {
				minIdx = j
			}
		}
		daftar[i], daftar[minIdx] = daftar[minIdx], daftar[i]
	}

	if ascending {
		fmt.Println("📌Data setelah diurutkan berdasarkan jumlah pinjaman (terkecil ke terbesar):")
	} else {
		fmt.Println("📌Data setelah diurutkan berdasarkan jumlah pinjaman (terbesar ke terkecil):")
	}

	tampilkanSemua()
}

func insertionSortByTenor(ascending bool) {
	for i := 1; i < len(daftar); i++ {
		key := daftar[i]
		j := i - 1
		if ascending {
			for j >= 0 && daftar[j].Tenor > key.Tenor {
				daftar[j+1] = daftar[j]
				j--
			}
		} else {
			for j >= 0 && daftar[j].Tenor < key.Tenor {
				daftar[j+1] = daftar[j]
				j--
			}
		}
		daftar[j+1] = key
	}

	if ascending {
		fmt.Println("📌Data setelah diurutkan berdasarkan tenor (terkecil ke terbesar):")
	} else {
		fmt.Println("📌Data setelah diurutkan berdasarkan tenor (terbesar ke terkecil):")
	}

	tampilkanSemua()
}

func tampilkanSemua() {
	fmt.Printf("%-3s %-20s %-15s %-8s %-10s %-18s %-15s %-15s %-20s\n",
		"No", "Nama", "Pinjaman", "Tenor", "Bunga", "Cicilan Per Bulan", "Total Bunga", "Total Akhir", "Status")
	fmt.Println(strings.Repeat("-", 123))

	for i, p := range daftar {
		bungaTotal := p.Pinjaman * p.Bunga * float64(p.Tenor) / 12
		totalAkhir := p.Pinjaman + bungaTotal

		fmt.Printf("%-3d %-20s %-15.0f %-8d %-10.2f %-18.0f %-15.0f %-15.0f %-20s\n",
			i+1, p.Nama, p.Pinjaman, p.Tenor, p.Bunga, p.Cicilan, bungaTotal, totalAkhir, p.StatusBayar)
	}
}

func laporan() {
	fmt.Println("📊 Ringkasan Laporan Peminjaman:")
	fmt.Printf("%-3s %-20s %-15s %-8s %-10s %-20s %-20s %-20s %-20s\n",
		"No", "Nama", "Pinjaman", "Tenor", "Bunga", "Cicilan Per Bulan", "Total Bunga", "Total Akhir", "Status")
	fmt.Println(strings.Repeat("-", 123))

	totalPinjaman := 0.0
	totalBunga := 0.0

	for i, p := range daftar {
		bunga := p.Pinjaman * p.Bunga * float64(p.Tenor) / 12
		total := p.Pinjaman + bunga
		totalPinjaman += p.Pinjaman
		totalBunga += bunga

		fmt.Printf("%-3d %-20s %-15.0f %-8d %-10.2f %-20.0f %-20.0f %-20.0f %-20s\n",
			i+1, p.Nama, p.Pinjaman, p.Tenor, p.Bunga, p.Cicilan, bunga, total, p.StatusBayar)
	}

	totalAkhir := totalPinjaman + totalBunga
	fmt.Println(strings.Repeat("-", 123))
	fmt.Printf("%-24s: %.0f\n", "💰 Total Pinjaman     ", totalPinjaman)
	fmt.Printf("%-24s: %.0f\n", "📈 Total Bunga        ", totalBunga)
	fmt.Printf("%-24s: %.0f\n", "🧾 Total yang Dibayar ", totalAkhir)
}

func dataPeminjaman() {
	daftar = append(daftar,
		Peminjam{"Vito Naryama", 10000000, 12, 0.05, hitungCicilan(10000000, 12, 0.05), "Belum Lunas"},
		Peminjam{"Agung Ramadhan", 5000000, 6, 0.05, hitungCicilan(5000000, 6, 0.05), "Belum Lunas"},
		Peminjam{"Fathur Rizal", 7000000, 10, 0.05, hitungCicilan(7000000, 10, 0.05), "Belum Lunas"},
		Peminjam{"Galang Saputra", 8000000, 8, 0.05, hitungCicilan(8000000, 8, 0.05), "Belum Lunas"},
		Peminjam{"Omar Nadiv", 15000000, 24, 0.05, hitungCicilan(15000000, 24, 0.05), "Belum Lunas"},
		Peminjam{"Mario Sebastian", 6000000, 36, 0.05, hitungCicilan(6000000, 36, 0.05), "Belum Lunas"},
		Peminjam{"Farrel Gaska", 9000000, 18, 0.05, hitungCicilan(9000000, 18, 0.05), "Belum Lunas"},
		Peminjam{"Fatih Zaki", 4000000, 9, 0.05, hitungCicilan(4000000, 9, 0.05), "Belum Lunas"},
		Peminjam{"Naufal Ammar", 11000000, 15, 0.05, hitungCicilan(11000000, 15, 0.05), "Belum Lunas"},
		Peminjam{"Rafi Tribuana", 20000000, 3, 0.05, hitungCicilan(20000000, 3, 0.05), "Belum Lunas"},
	)
}

func menu() {
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("\n======= Menu Utama =======")
		fmt.Println("1. Tambah Data Peminjam")
		fmt.Println("2. Edit Data Peminjam")
		fmt.Println("3. Hapus Data Peminjam")
		fmt.Println("4. Lihat Semua Peminjam")
		fmt.Println("5. Cari Peminjam (Manual)")
		fmt.Println("6. Cari Peminjam (Cepat)")
		fmt.Println("7. Urutkan Berdasarkan Pinjaman")
		fmt.Println("8. Urutkan Berdasarkan Tenor")
		fmt.Println("9. Lihat Ringkasan Laporan")
		fmt.Println("0. Keluar dari Aplikasi")

		var pilihan int
		for {
			fmt.Print("Silakan pilih menu (0-9): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			var err error
			pilihan, err = strconv.Atoi(input)
			if err != nil {
				fmt.Println("❗ Pilihan tidak mengandung huruf. Silakan coba lagi")
				continue
			}
			break
		}

		fmt.Println()

		switch pilihan {
		case 1:
			tambahPeminjam()
		case 2:
			ubahPeminjam()
		case 3:
			tampilkanSemua()
			hapusPeminjam()
		case 4:
			tampilkanSemua()
		case 5:
			for {
				fmt.Print("Masukkan Nama lengkap yang dicari: ")
				input, _ := reader.ReadString('\n')
				nama := strings.ToLower(strings.TrimSpace(input))
				p := sequentialSearch(nama)
				if p != nil {
					bungaTotal := p.Pinjaman * p.Bunga * float64(p.Tenor) / 12
					totalAkhir := p.Pinjaman + bungaTotal
					fmt.Println("📌 Data ditemukan:")
					fmt.Printf("%-3s %-20s %-15s %-8s %-10s %-15s %-15s %-15s %-15s\n",
						"No", "Nama", "Pinjaman", "Tenor", "Bunga", "Cicilan ", "Total Bunga", "Total Akhir", "Status")
					fmt.Println(strings.Repeat("-", 123))
					fmt.Printf("%-3d %-20s %-15.0f %-8d %-10.2f %-15.0f %-15.0f %-15.0f %-15s\n",
						1, p.Nama, p.Pinjaman, p.Tenor, p.Bunga, p.Cicilan, bungaTotal, totalAkhir, p.StatusBayar)
				} else {
					fmt.Println("❌ Data peminjam tidak ditemukan.")
				}
				break
			}
		case 6:
			for {
				fmt.Print("Masukkan Nama lengkap yang dicari: ")
				input, _ := reader.ReadString('\n')
				nama := strings.TrimSpace(input)
				p := binarySearch(nama)
				if p != nil {
					bungaTotal := p.Pinjaman * p.Bunga * float64(p.Tenor) / 12
					totalAkhir := p.Pinjaman + bungaTotal
					fmt.Println("📌 Data ditemukan:")
					fmt.Printf("%-3s %-20s %-15s %-8s %-10s %-15s %-15s %-15s %-15s\n",
						"No", "Nama", "Pinjaman", "Tenor", "Bunga", "Cicilan", "Total Bunga", "Total Akhir", "Status")
					fmt.Println(strings.Repeat("-", 123))
					fmt.Printf("%-3d %-20s %-15.0f %-8d %-10.2f %-15.0f %-15.0f %-15.0f %-15s\n",
						1, p.Nama, p.Pinjaman, p.Tenor, p.Bunga, p.Cicilan, bungaTotal, totalAkhir, p.StatusBayar)
				} else {
					fmt.Println("❌ Data peminjam tidak ditemukan.")
				}
				break
			}
		case 7:
			var urut int
			for {
				fmt.Println("Urutkan pinjaman berdasarkan:")
				fmt.Println("1. Terkecil ke terbesar")
				fmt.Println("2. Terbesar ke terkecil")
				fmt.Print("Pilih (1/2): ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)
				var err error
				urut, err = strconv.Atoi(input)
				if err != nil || (urut != 1 && urut != 2) {
					fmt.Println("❗ Pilihan tidak valid. Silakan coba lagi.")
					continue
				}
				if urut == 1 {
					insertionSortByTenor(true)
				} else if urut == 2 {
					insertionSortByTenor(false)
				}
				break
			}

		case 8:
			for {
				var urut int
				fmt.Println("Urutkan tenor berdasarkan:")
				fmt.Println("1. Terkecil ke terbesar")
				fmt.Println("2. Terbesar ke terkecil")
				fmt.Print("Pilih (1/2): ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)
				var err error
				urut, err = strconv.Atoi(input)
				if err != nil || (urut != 1 && urut != 2) {
					fmt.Println("❗ Pilihan tidak valid. Silakan coba lagi.")
					continue
				}
				if urut == 1 {
					insertionSortByTenor(true)
				} else if urut == 2 {
					insertionSortByTenor(false)
				}
				break
			}
		case 9:
			laporan()
		case 0:
			fmt.Println("👋 Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("❗ Pilihan tidak valid. Silakan coba lagi.")
		}

	}
}

func main() {
	dataPeminjaman()
	menu()
}
