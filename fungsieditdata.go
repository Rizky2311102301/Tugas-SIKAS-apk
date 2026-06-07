func editData() {
	if len(datamahasiswa) == 0 {
		fmt.Println("\n✗ Belum ada Data Mahasiswa!")
		return
	}
	var inputNIM string
	fmt.Print("\nMasukkan NIM Mahasiswa yang akan diedit: ")
	fmt.Scanln(&inputNIM)
	idx := -1
	for i := 0; i < len(datamahasiswa); i++ {
		if datamahasiswa[i].NIM == inputNIM {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Println("\n✗ Mahasiswa dengan NIM tidak ditemukan!")
		return
	}
	fmt.Println("\nDATA MAHASISWA")
	fmt.Printf("NIM         : %s\n", datamahasiswa[idx].NIM)
	fmt.Printf("Nama        : %s\n", datamahasiswa[idx].Nama)
	fmt.Printf("Total Bayar : Rp %d\n", datamahasiswa[idx].Totalbayar)
	fmt.Printf("Status      : %s\n", statusLunas(datamahasiswa[idx].StatusLunas))
	if !datamahasiswa[idx].StatusLunas {
		fmt.Printf("Sisa Tunggakan: Rp %d\n", TARGET_KAS-datamahasiswa[idx].Totalbayar)
	}
	fmt.Println("\nPILIH DATA YANG AKAN DIEDIT")
	fmt.Println("1. Edit NIM")
	fmt.Println("2. Edit Nama")
	fmt.Println("3. Edit Total Bayar")
	fmt.Print("Pilih: ")
	var pilihan int
	fmt.Scanln(&pilihan)
	if pilihan == 1 {
		var NIMbaru string
		fmt.Print("\nMasukkan NIM baru: ")
		fmt.Scanln(&NIMbaru)
		NIMlama := datamahasiswa[idx].NIM
		datamahasiswa[idx].NIM = NIMbaru
		for i := 0; i < len(riwayatPembayaran); i++ {
			if riwayatPembayaran[i].NIM == NIMlama {
				riwayatPembayaran[i].NIM = NIMbaru
			}
		}
		fmt.Println("\nDATA SETELAH DIEDIT")
		fmt.Printf("NIM         : %s\n", datamahasiswa[idx].NIM)
		fmt.Printf("Nama        : %s\n", datamahasiswa[idx].Nama)
		fmt.Printf("Total Bayar : Rp %d\n", datamahasiswa[idx].Totalbayar)
		fmt.Printf("Status      : %s\n", statusLunas(datamahasiswa[idx].StatusLunas))
		if !datamahasiswa[idx].StatusLunas {
			fmt.Printf("Sisa Tunggakan: Rp %d\n", TARGET_KAS-datamahasiswa[idx].Totalbayar)
		}
		fmt.Println("\n✓ NIM berhasil diubah!")
	} else if pilihan == 2 {
		var namabaru string
		fmt.Print("\nMasukkan Nama Baru: ")
		fmt.Scanln(&namabaru)
		datamahasiswa[idx].Nama = namabaru
		for i := 0; i < len(riwayatPembayaran); i++ {
			if riwayatPembayaran[i].NIM == datamahasiswa[idx].NIM {
				riwayatPembayaran[i].Nama = namabaru
			}
		}
		fmt.Println("\nDATA SETELAH DIEDIT")
		fmt.Printf("NIM         : %s\n", datamahasiswa[idx].NIM)
		fmt.Printf("Nama        : %s\n", datamahasiswa[idx].Nama)
		fmt.Printf("Total Bayar : Rp %d\n", datamahasiswa[idx].Totalbayar)
		fmt.Printf("Status      : %s\n", statusLunas(datamahasiswa[idx].StatusLunas))
		if !datamahasiswa[idx].StatusLunas {
			fmt.Printf("Sisa Tunggakan: Rp %d\n", TARGET_KAS-datamahasiswa[idx].Totalbayar)
		}
		fmt.Println("\n✓ Nama berhasil diubah!")
	} else if pilihan == 3 {
		fmt.Printf("\nTotal Bayar Saat Ini: Rp %d\n", datamahasiswa[idx].Totalbayar)
		fmt.Printf("Sisa Tunggakan Saat Ini: Rp %d\n", TARGET_KAS-datamahasiswa[idx].Totalbayar)
		var totalbaru int
		fmt.Print("Masukkan Total Bayar baru: Rp ")
		fmt.Scanln(&totalbaru)
		if totalbaru < 0 {
			fmt.Println("\n✗ Total bayar tidak boleh negatif!")
			return
		}
		datamahasiswa[idx].Totalbayar = totalbaru
		if datamahasiswa[idx].Totalbayar >= TARGET_KAS {
			datamahasiswa[idx].StatusLunas = true
		} else {
			datamahasiswa[idx].StatusLunas = false
		}
		fmt.Println("\nDATA SETELAH DIEDIT")
		fmt.Printf("NIM         : %s\n", datamahasiswa[idx].NIM)
		fmt.Printf("Nama        : %s\n", datamahasiswa[idx].Nama)
		fmt.Printf("Total Bayar : Rp %d\n", datamahasiswa[idx].Totalbayar)
		fmt.Printf("Status      : %s\n", statusLunas(datamahasiswa[idx].StatusLunas))
		if !datamahasiswa[idx].StatusLunas {
			fmt.Printf("Sisa Tunggakan: Rp %d\n", TARGET_KAS-datamahasiswa[idx].Totalbayar)
		}
		fmt.Println("\n✓ Total bayar berhasil diubah!")
	} else {
		fmt.Println("\n✗ Pilihan tidak valid! Silakan pilih 1, 2, atau 3.")
		return
	}
}