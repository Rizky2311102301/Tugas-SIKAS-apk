func pembayarankas() {
	var input string
	fmt.Print("Masukkan NIM atau Nama Mahasiswa: ")
	fmt.Scan(&input)
	idx := -1
	i := 0
	for i < len(datamahasiswa) {
		if datamahasiswa[i].NIM == input || datamahasiswa[i].Nama == input {
			idx = i
		}
		i++
	}
	if idx != -1 {
		fmt.Println("\nData ditemukan!")
		fmt.Println("NIM  :", datamahasiswa[idx].NIM)
		fmt.Println("Nama :", datamahasiswa[idx].Nama)
		var nominal int
		fmt.Print("Masukkan nominal pembayaran: ")
		fmt.Scan(&nominal)
		datamahasiswa[idx].Totalbayar += nominal
		if datamahasiswa[idx].Totalbayar >= TARGET_KAS {
			datamahasiswa[idx].StatusLunas = true
		}
		fmt.Println("\nPEMBAYARAN BERHASIL")
		fmt.Println("Total Bayar :", datamahasiswa[idx].Totalbayar)
		if datamahasiswa[idx].StatusLunas {
			fmt.Println("Status : LUNAS")
		} else {
			sisa := TARGET_KAS - datamahasiswa[idx].Totalbayar
			fmt.Println("Status : BELUM LUNAS")
			fmt.Println("Sisa Tunggakan :", sisa)
		}
	} else {
		fmt.Println("\nData mahasiswa tidak ditemukan!")
	}
}