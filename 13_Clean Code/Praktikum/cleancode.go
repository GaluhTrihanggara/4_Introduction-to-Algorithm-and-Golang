// kekuranganya sebanyak 5 kesalahan dalam program ini yaitu

package main

// 1. Variabel "username" dan "password" seharusnya bertipe string bukan int.
// karena umumnya username dan password merupakan string yang terdiri dari karakter-karakter.
type user struct {
	id       int
	username int
	password int
}

// 2. konsisten  dalam penamnaan variabel "t" pada "userservice".
// sebaiknya diberikan nama yang lebih bisa deskriptif dan konsisten dalam penamaan sebuah variabel.
//3 Kepemilikan struktur antara "userservice" dan "user" kurang jelas.
// perlu ditambahkan komentar atau dokumentasi yang menjelaskan bahwa "userservice" adalah layanan yang menyediakan operasi-operasi terhadap "user".
type userservice struct {
	t []user
}

// 4 Penggunaan getter pada "getallusers" dan "getuserbyid" tidak efektif
// karena pada penggunaan getter hanya mengembalikan nilai dari variabel internal tanpa melakukan apa-apa.
// jadi perlu dipikikan kembali apakah penggunaan getter ini dibutuhkan atau tidak.
func (u userservice) getallusers() []user {
	return u.t
}

// 5 Tidak ada validasi input pada "getuserbyid".
// karena diperlukanya tambahan validasi untuk memastikan bahwa input id adalah sebuah bilangan bulat positif.
func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
