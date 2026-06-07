func HapusData() {
	if len(Datamahasiswa) == 0 {
		fmt.Println("\n✗ Belum ada data Mahasiswa!")
		return
	}
	var nim string
	fmt.Print("\nMasukkan NIM Mahasiswa yang akan dihapus: ")
	fmt.Scanln(&nim)
	idx := -1
	ditemukan := false
	for i := 0; i < len(Datamahasiswa) && !ditemukan; i++ {
		if Datamahasiswa[i].NIM == nim {
			idx = i
			ditemukan = true
		}
	}
	if idx == -1 {
		fmt.Println("✗ Mahasiswa dengan NIM tersebut tidak ditemukan!")
		return
	}
	fmt.Println("\n=== Detail Data yang Akan Dihapus ===")
	fmt.Printf("NIM        : %s\n", Datamahasiswa[idx].NIM)
	fmt.Printf("Nama       : %s\n", Datamahasiswa[idx].Nama)
	fmt.Printf("Total Bayar: Rp %d\n", Datamahasiswa[idx].Totalbayar)
	fmt.Print("\nYakin ingin menghapus data ini? (yes/no): ")
	var konfirmasi string
	fmt.Scanln(&konfirmasi)
	if konfirmasi == "yes" || konfirmasi == "YES" || konfirmasi == "y" || konfirmasi == "Y" {
		for i := idx; i < len(Datamahasiswa)-1; i++ {
			Datamahasiswa[i] = Datamahasiswa[i+1]
		}
		Datamahasiswa = Datamahasiswa[:len(Datamahasiswa)-1]
		fmt.Println("\n✓ Data berhasil dihapus!")
	} else {
		fmt.Println("\n✗ Penghapusan data dibatalkan.")
	}
}