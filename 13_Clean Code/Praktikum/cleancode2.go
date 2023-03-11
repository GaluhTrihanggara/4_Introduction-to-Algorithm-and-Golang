package main

type kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

type mobil struct {
	kendaraan
}

func (m *mobil) berjalan() {
	m.tambahKecepatan(10)
}

func (m *mobil) tambahKecepatan(kecepatanBaru int) {
	m.kecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	mobilLamban := mobil{}
	mobilLamban.berjalan()
}

// Penjelasan perbaikan kode:
// 1. Penulisan kode diubah agar lebih konsisten, misalnya pada penulisan spasi dan penulisan huruf kapital.
// 2. Nama variabel diubah agar lebih deskriptif, seperti totalRoda dan kecepatanPerJam pada struct kendaraan.
// 3. Nama fungsi diubah menjadi lebih deskriptif, seperti tambahKecepatan.
// 4. Penulisan operator += digunakan untuk memperpendek penulisan.
// 5. Nama variabel pada fungsi main diubah agar lebih deskriptif dan sesuai dengan aturan camel case.
// 6. Penulisan fungsi berjalan diubah agar lebih rapi dengan tidak menuliskan penambahan kecepatan langsung di dalamnya.
