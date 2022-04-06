package main
import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
	"time"
	"bufio"
)

const J = 5
const N = 10000
type ttl struct {
	tgl, bulan, tahun int
}
type profil struct {
	name *bufio.Scanner
	asal *bufio.Scanner
	id, del, cek, lolos, ujian, cekl int
	ultah ttl
	mtk, ipa, ips, ing, tpa, ting float64
	jurusan [J]int
}
var pendaftar [11]int
var bersyukur [11]int
var akun [N]profil
var waktu = time.Now()
var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}
func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

func main () {
	var dafma, benar, vlist string
	var login, dapat int
	i:=0
	menu()
	fmt.Print("Tulis disini: ")
	fmt.Scanln(&dafma)
	benar = kesalahan(dafma)
	for benar == "masuk" || benar == "list" {
		fmt.Println("Maaf belum ada akun yang terdaftar")
		fmt.Println("Mohon untuk daftar terlebih dahulu")
		fmt.Print("Tulis daftar disini: ")
		fmt.Scanln(&dafma)
		benar = kesalahan(dafma)
	}
	for benar == "daftar" || benar != "exit" {
		if benar == "daftar" {
			CallClear() 					
			create(i)						 
			menu()							
			fmt.Print("Tulis disini: ")
			fmt.Scanln(&dafma)
			benar = kesalahan(dafma)
			i++
		}
		for benar == "masuk" {
			fmt.Print("Masukkan ID anda: ")
			fmt.Scanln(&login)
			for cari(login,i) == -1 {
				fmt.Println("Maaf ID anda tidak ditemukan")
				fmt.Print("Silahkan coba disini: ")
				fmt.Scanln(&login)
			}
			fmt.Println("ID telah ditemukan")
			fmt.Println("Mohon tunggu sebentar...")
			time.Sleep(3 * time.Second)
			CallClear()
			dapat = cari(login,i)
			if akun[dapat].del == 1 {
				fmt.Println("Akun ini telah dihapus")
				fmt.Println("")
			} else {	
				tampil(dapat)
				fmt.Println("")
				setting(dapat)
			}
			menu()
			fmt.Print("Tulis disini: ")
			fmt.Scanln(&dafma)
			benar = kesalahan(dafma)
		}
		for benar  == "list" {
			CallClear()
			menu_list()
			fmt.Print("Tulis disini: ")
			fmt.Scanln(&vlist)
			listcorrect(&vlist)
			pilih_list(vlist, i)
			fmt.Println("")
			menu()
			fmt.Print("Tulis disini: ")
			fmt.Scanln(&dafma)
			benar = kesalahan(dafma)
		}
	}
	fmt.Println("Menunggu keluar...")
	time.Sleep(2 * time.Second)
	CallClear()
}
func menu_list () {
	fmt.Println("Tulis 'TPA' untuk menampilkan data berdasarkan nilai TPA")
	fmt.Println("Tulis 'TIng' untuk menampilkan data berdasarkan nilai Tes TOEFL")
	fmt.Println("Tulis 'jurusan' untuk menampilkan data berdasarkan jurusan")
	fmt.Println("")
}
func kesalahan(tulis string)string {
	for tulis != "daftar" && tulis != "masuk" && tulis != "list" && tulis != "exit" {
		fmt.Println("Tolong tulis dengan petunjuk yang ada")
		fmt.Print("Tulis ulang disini: ")
		fmt.Scanln(&tulis)
	}
	return tulis
}
func listcorrect(tulis *string) {
	for *tulis != "TPA" && *tulis != "TIng" && *tulis != "jurusan" {
		fmt.Println("Tolong tulis dengan petunjuk yang ada")
		fmt.Print("Tulis ulang disini: ")
		fmt.Scanln(*&tulis)
	}
}
func create(i int) {
	var cocok int
	fmt.Print("Masukkan nama anda: ")
	akun[i].name = bufio.NewScanner(os.Stdin)
	akun[i].name.Scan()
	fmt.Print("Asal Daerah: ")
	akun[i].asal = bufio.NewScanner(os.Stdin)
	akun[i].asal.Scan()
	fmt.Print("Tanggal Lahir(hr bln thn)(contoh: 27 12 2000): ")
	fmt.Scanln(&akun[i].ultah.tgl, &akun[i].ultah.bulan, &akun[i].ultah.tahun)
	cocok = jumlah_hari(akun[i].ultah.bulan, akun[i].ultah.tahun)
	for akun[i].ultah.tgl > cocok || akun[i].ultah.bulan > 12 || akun[i].ultah.tahun < 1000 || akun[i].ultah.tahun >= waktu.Year() {
		fmt.Println("Kesalahan pada tanggal lahir. Silahkan coba lagi")
		fmt.Print("Coba disini: ")
		fmt.Scanln(&akun[i].ultah.tgl, &akun[i].ultah.bulan, &akun[i].ultah.tahun)
		cocok = jumlah_hari(akun[i].ultah.bulan, akun[i].ultah.tahun)
	}
	pilihjurusan(i)
	input_nilai(i)
	akun[i].id = N + i
	fmt.Println("")
	fmt.Println("ID anda adalah: ", akun[i].id)
	fmt.Println("Gunakan ID anda untuk masuk")
	fmt.Println(" ")
}
func cari(data,i int)int {
	tengah:=0
	for akun[tengah].id != data && tengah < i {
		tengah++
	}
	if akun[tengah].id == data {
		return tengah
	} else {
		return -1
	}
}
func daftarjurusan () {
	fmt.Println("1. S1 Informatika             6. S1 Manajemen Bisnis Telekomunikasi dan Informatika")
	fmt.Println("2. S1 Teknik Telekomunikasi   7. S1 Akuntansi")
	fmt.Println("3. S1 Teknik Elektro          8. S1 Ilmu Komunikasi")
	fmt.Println("4. S1 Teknik Industri         9. S1 Administaris Bisnis")
	fmt.Println("5. S1 Sistem Informasi        10. S1 Desain Komunikasi Visual")
}
func tampil(tengah int) {
	var huruf string
	var nilai float64

	fmt.Println("Selamat anda berhasil masuk")
	fmt.Println("Nama: ", akun[tengah].name.Text())
	fmt.Println("Asal: ", akun[tengah].asal.Text())
	fmt.Println("Tanggal Lahir: ", akun[tengah].ultah.tgl, akun[tengah].ultah.bulan, akun[tengah].ultah.tahun)
	for k:=2; k<J; k++ {
		jurusan_ke_huruf(tengah, k, &huruf, &nilai)
		fmt.Println("Jurusan ", k-1, " : ", huruf)
	}
	fmt.Println("Nilai Rapor")
	fmt.Println("Matematika: ", akun[tengah].mtk)
	fmt.Println("inggris: ", akun[tengah].ing)
	if akun[tengah].cek == 1 {
		fmt.Println("IPA: ", akun[tengah].ipa)
	} else {
		fmt.Println("IPS: ", akun[tengah].ips)
	}
	fmt.Println("TPA: ", akun[tengah].tpa)
	fmt.Println("TOEFL: ", akun[tengah].ting)
}
func setting(tengah int) {
	var edit, huruf, ujn string
	var nilai, rata float64
	var temp, k int
	fmt.Println("Jika anda ingin mengedit data")
	menu_setting()
	fmt.Println("Jika anda tidak ingin mengedit data, tulis 'pass'")
	fmt.Println("")
	fmt.Println("Tulis 'ujian' untuk melakukan ujian")
	fmt.Println("Tulis 'kelulusan' untuk mengecek kelulusan")
	fmt.Println("KET: - Anda dapat mengecek kelulusan jika sudah ujian")
	fmt.Println("     - Jika anda sudah mengecek kelulusan anda sudah tidak dapat mengedit data")
	fmt.Println("Jika anda ingin menghapus akun, tulis 'hapus'")
	fmt.Println("")
	fmt.Print("Tulis disini: ")
	fmt.Scanln(&edit)
	cek_edit(&edit)
	for edit != "pass" && edit != "hapus" && edit != "kelulusan" && edit != "ujian" {
		if akun[tengah].cekl == 0 {
			pilih_edit(edit, tengah)
			fmt.Println("")
			fmt.Println("Jika anda ingin mengedit data lagi, tulis sesuai petunjuk")
			menu_setting()
			fmt.Println("Jika tidak tulis 'pass'")
			fmt.Println("")
			fmt.Print("Tulis disini: ")
			fmt.Scanln(&edit)
			cek_edit(&edit)
		} else {
			fmt.Println("Anda sudah tidak dapat mengedit data lagi")
			fmt.Println("Karena anda sudah mengecek kelulusan anda")
			edit = "pass"
		}
	}
	if edit == "kelulusan" && akun[tengah].ujian == 0 {
		fmt.Println("Mohon ujian terlebih dahulu")
		fmt.Print("Tulis daftar disini: ")
		fmt.Scanln(&edit)
		cek_edit(&edit)
		for edit != "ujian" {
			fmt.Println("Mohon ujian terlebih dahulu")
			fmt.Print("Tulis daftar disini: ")
			fmt.Scanln(&edit)
			cek_edit(&edit)
		}
	}
	if edit == "ujian" {
		if akun[tengah].cekl == 0 {
			fmt.Println("")
			tampil_ujian(&ujn)
			if ujn == "tes" {
				fmt.Print("Masukkan nilai TPA: ")
				fmt.Scanln(&akun[tengah].tpa)
				for akun[tengah].tpa > 1000 || akun[tengah].tpa < 0 {
					fmt.Println("Tolong masukkan nilai yang benar")
					fmt.Print("Tulis disini: ")
					fmt.Scanln(&akun[tengah].tpa)
				}
				fmt.Print("Masukkan nilai TOEFL: ")
				fmt.Scanln(&akun[tengah].ting)
				for akun[tengah].ting > 677 || akun[tengah].ting < 0 {
					fmt.Println("Tolong masukkan nilai yang benar")
					fmt.Print("Tulis disini: ")
					fmt.Scanln(&akun[tengah].ting)
				}
				fmt.Println("Nilai ujian telah masuk")
				fmt.Println("Anda sudah dapat mengecek kelulusan")
			}
			if ujn == "delete" {
				akun[tengah].tpa = 0
				akun[tengah].ting = 0
				fmt.Println("Nilai telah dihapus")
			}
			akun[tengah].ujian = 1
		} else {
			fmt.Println("Anda sudah tidak dapat ujian lagi")
			fmt.Println("Karena anda sudah mengecek kelulusan anda")
		}
	}
	if edit == "hapus" {
		akun[tengah].del = 1
		akun[tengah].ujian = 0
		akun[tengah].lolos = 0
		fmt.Println("Selamat akun anda telah dihapus")
		if akun[tengah].cek > 0 {
			rata = (((akun[tengah].ipa*0.5)+(akun[tengah].mtk*0.3)+(akun[tengah].ing*0.2))*0.4) + (((akun[tengah].tpa*0.07)+(akun[tengah].ting*0.045))*0.6)
		} else {
			rata = (((akun[tengah].ips*0.5)+(akun[tengah].mtk*0.3)+(akun[tengah].ing*0.2))*0.4) + (((akun[tengah].tpa*0.07)+(akun[tengah].ting*0.045))*0.6)
		}
		k = 2
		for k < J {
			temp = akun[tengah].jurusan[k]
			pendaftar[temp] = pendaftar[temp] - 1
			k++
		}
		k = 2
		for k < J && akun[tengah].lolos != 1 {
			temp = akun[tengah].jurusan[k]
			jurusan_ke_huruf(tengah, k, &huruf, &nilai)
			if rata >= nilai {
				akun[tengah].lolos = 1
				bersyukur[temp] = bersyukur[temp] - 1
			}
			k++
		}
	}
	if edit == "kelulusan" {
		if akun[tengah].cek > 0 {
			rata = (((akun[tengah].ipa*0.5)+(akun[tengah].mtk*0.3)+(akun[tengah].ing*0.2))*0.4) + (((akun[tengah].tpa*0.07)+(akun[tengah].ting*0.045))*0.6)
		} else {
			rata = (((akun[tengah].ips*0.5)+(akun[tengah].mtk*0.3)+(akun[tengah].ing*0.2))*0.4) + (((akun[tengah].tpa*0.07)+(akun[tengah].ting*0.045))*0.6)
		}
		CallClear()
		k = 2
		for k < J && akun[tengah].lolos != 1 {
			temp = akun[tengah].jurusan[k]
			jurusan_ke_huruf(tengah, k, &huruf, &nilai)
			if rata >= nilai {
				fmt.Println("Selamat anda diterima di prodi", huruf)
				akun[tengah].lolos = 1
				bersyukur[temp] = bersyukur[temp] + 1
			}
			k++
		}
		if akun[tengah].lolos == 0 {
			fmt.Println("Maaf anda kurang beruntung")
			fmt.Println("Jangan berkecil hati dan coba lagi")
		}
		akun[tengah].cekl = 1
	}
	fmt.Println("")
}
func cek_edit (edit *string) {
	for *edit != "nama" && *edit != "TL" && *edit != "asal" && *edit != "jurusan" && *edit != "pass" && *edit != "hapus" && *edit != "nilai" && *edit != "kelulusan" && *edit != "ujian" {
		fmt.Println("Maaf tulis dengan petunjuk yang ada")
		fmt.Print("Tulis ulang disini: ")
		fmt.Scanln(*&edit)
	}
}
func tampil_ujian (tulis *string) {
	var teldel string
	fmt.Println("Tulis 'tes' untuk memasukkan atau mengedit nilai ujian")
	fmt.Println("Tulis 'delete' untuk mengapus nilai ujian")
	fmt.Println("")
	fmt.Print("Tulis disini: ")
	fmt.Scanln(&teldel)
	for teldel != "tes" && teldel != "delete" {
		fmt.Println("Tolong tulis sesuai petunjuk")
		fmt.Print("Tulis disini: ")
		fmt.Scanln(&teldel)
	}
	*tulis = teldel
}
func menu_setting () {
	fmt.Println("Tulis 'nama' untuk mengedit nama")
	fmt.Println("Tulis 'TL' untuk mengedit tanggal lahir")
	fmt.Println("Tulis 'asal' untuk mengedit asal daerah")
	fmt.Println("Tulis 'nilai' untuk mengedit nilai rapor")
	fmt.Println("Tulis 'jurusan' untuk mengedit jurusan")
}
func menu () {
	fmt.Println("Selamat Datang di Pendaftaran Telkom University")
	fmt.Println(" ")
	fmt.Println("Tulis 'daftar' untuk mendaftar")
	fmt.Println("Tulis 'masuk' untuk Log in")
	fmt.Println("Tulis 'list' untuk melihat list data")
	fmt.Println("Tulis 'exit' untuk keluar")
	fmt.Println(" ")
}
func input_nilai(i int) {
	fmt.Println("Masukkan Nilai Rapor")
	fmt.Print("Matematika: ")
	fmt.Scanln(&akun[i].mtk)
	perbaikan(&akun[i].mtk)
	fmt.Print("inggris: ")
	fmt.Scanln(&akun[i].ing)
	perbaikan(&akun[i].ing)
	if akun[i].jurusan[2] <= 5 || akun[i].jurusan[3] <= 5 || akun[i].jurusan[4] <= 5 {
		fmt.Print("IPA: ")
		fmt.Scanln(&akun[i].ipa)
		perbaikan(&akun[i].ipa)
		akun[i].cek = 1
	} else {
		fmt.Print("IPS: ")
		fmt.Scanln(&akun[i].ips)
		perbaikan(&akun[i].ips)
		akun[i].cek = 0
	}
}
func perbaikan(nilai *float64) {
	for *nilai > 100 || *nilai < 0 {
		fmt.Println("Tolong masukkan nilai yang benar")
		fmt.Print("Tulis disini: ")
		fmt.Scanln(*&nilai)
	}
}
func kabisat(tahun int)bool {
	return tahun % 400 == 0 || tahun % 4 == 0 && tahun % 100 > 0
}
func jumlah_hari(bln,thn int)int {
	var hari int
	switch bln {
	case 1:
		hari = 31
	case 2:
		if kabisat(thn) {
			hari = 29
		} else {
			hari = 28
		}
	case 3:
		hari = 31
	case 4:
		hari = 30
	case 5:
		hari = 31
	case 6:
		hari = 30
	case 7:
		hari = 31
	case 8:
		hari = 31
	case 9:
		hari = 30
	case 10:
		hari = 31
	case 11:
		hari = 30
	case 12:
		hari = 31
	}
	return hari
}
func pilih_list(edit string, n int) {
	var sudah, lihatjurusan int
	switch edit {
	case "TPA":
		uruttpa(n)
		sudah = 0
		for j:=0; j<n; j++ {
			if akun[j].ujian == 1 {
				fmt.Println("nama:", akun[j].name.Text(), "   TPA:", akun[j].tpa)
				sudah = 1
			}
		}
		if sudah == 0 {
			fmt.Println("Belum ada akun yang sudah melakukan ujian TPA")
		}
	case "TIng":
		urutting(n)
		sudah = 0
		for k:=0; k<n; k++ {
			if akun[k].ujian == 1 {
				fmt.Println("nama:", akun[k].name.Text(), "   TOEFL:", akun[k].ting)
				sudah = 1
			}
		}
		if sudah == 0 {
			fmt.Println("Belum ada akun yang sudah melakukan tes bahasa inggris")
		}
	case "jurusan":
		fmt.Println("Tulis angka jurusan untuk melihat data dari jurusan tersebut")
		daftarjurusan()
		fmt.Print("Tulis disini(angka): ")
		fmt.Scanln(&lihatjurusan)
		for lihatjurusan > 10 || lihatjurusan <= 0 {
			fmt.Println("Maaf Daftar jurusan tidak ada")
			fmt.Print("Tolong masukkan ulang jurusan disini: ")
			fmt.Scanln(&lihatjurusan)
		}
		showjurusan(lihatjurusan)
	}
}
func showjurusan(major int) {
	switch major{
	case 1:
		fmt.Println("S1 Informatika")
	case 2:
		fmt.Println("S1 Teknik Telekomunikasi")
	case 3:
		fmt.Println("S1 Teknik Elektro")
	case 4:
		fmt.Println("S1 Teknik Industri")
	case 5:
		fmt.Println("S1 Sistem Informasi")
	case 6:
		fmt.Println("S1 Manajemen Bisnis Telekomunikasi dan Informasi")
	case 7:
		fmt.Println("S1 Akuntansi")
	case 8:
		fmt.Println("S1 Ilmu Komunikasi")
	case 9:
		fmt.Println("S1 Administrasi Bisnis")
	case 10:
		fmt.Println("S1 Desain Komunikasi Visual")
	}
	fmt.Println("Pendaftar: ", pendaftar[major])
	fmt.Println("Lulus: ", bersyukur[major])
	fmt.Println("")
	fmt.Println("ket: Data Lulus hanya mengecek data yang sudah mengecek kelulusan dan diterima")
}
func uruttpa(n int) {
	var temp profil
	var j int
	i:=1
	for i <= n {
		j = i - 1
		for j >= 0 && akun[j].tpa < akun[j+1].tpa  {
			temp = akun[j+1]
			akun[j+1] = akun[j]
			akun[j] = temp
			j = j -1
		}
		i++
	}
}
func urutting(n int) {
	var temp profil
	var j int
	i:=1
	for i <= n {
		j = i - 1
		for j >= 0 && akun[j].ting < akun[j+1].ting  {
			temp = akun[j+1]
			akun[j+1] = akun[j]
			akun[j] = temp
			j = j -1
		}
		i++
	}
}
func pilih_edit(edit string, tengah int) {
	var cocok, temp int
	switch edit {
	case "nama":
		fmt.Print("Masukkan nama anda disini: ")
		akun[tengah].name = bufio.NewScanner(os.Stdin)
		akun[tengah].name.Scan()
		fmt.Println("Selamat nama telah diperbaharui")
	case "TL":
		fmt.Print("Masukkan Tanggal Lahir(hr bln thn)(contoh: 27 12 2000): ")
		fmt.Scanln(&akun[tengah].ultah.tgl, &akun[tengah].ultah.bulan, &akun[tengah].ultah.tahun)
		cocok = jumlah_hari(akun[tengah].ultah.bulan, akun[tengah].ultah.tahun)
		for akun[tengah].ultah.tgl > cocok || akun[tengah].ultah.bulan > 12 || akun[tengah].ultah.tahun < 1000 || akun[tengah].ultah.tahun >= waktu.Year() {
			fmt.Println("Kesalahan pada tanggal atau bulan lahir. Silahkan coba lagi")
			fmt.Print("Coba disini: ")
			fmt.Scanln(&akun[tengah].ultah.tgl, &akun[tengah].ultah.bulan, &akun[tengah].ultah.tahun)
			cocok = jumlah_hari(akun[tengah].ultah.bulan, akun[tengah].ultah.tahun)
		}
		fmt.Println("Selamat tanggal lahir telah diperbaharui")
	case "asal":
		fmt.Print("Masukkan asal anda disini: ")
		akun[tengah].asal = bufio.NewScanner(os.Stdin)
		akun[tengah].asal.Scan()
		fmt.Println("Selamat asal telah diperbaharui")
	case "nilai":
		input_nilai(tengah)
		fmt.Println("Selamat nilai telah diperbaharui")
	case "jurusan":
		for k:=0; k < J; k++ {
			temp = akun[tengah].jurusan[k]
			pendaftar[temp] = pendaftar[temp] - 1
		}
		pilihjurusan(tengah)
		fmt.Println("Selamat jurusan telah diperbaharui")
	}
}
func pilihjurusan(i int) {
	var temp int
	fmt.Println("Tulis Jurusan anda sesuai dengan angka jurusan")
	daftarjurusan()
	fmt.Println("Pilih Jurusan (sesuai angka): ")
	for k:=2; k<J; k++ {
		fmt.Print("Pilihan ", k-1, " : ")
		fmt.Scanln(&akun[i].jurusan[k])
		for akun[i].jurusan[k] > 10 || akun[i].jurusan[k] <= 0 {
			fmt.Println("Maaf Daftar jurusan tidak ada")
			fmt.Print("Tolong masukkan ulang jurusan disini: ")
			fmt.Scanln(&akun[i].jurusan[k])
		}
		for akun[i].jurusan[k] == akun[i].jurusan[k-1] || akun[i].jurusan[k] == akun[i].jurusan[k-2] {
			fmt.Println("Maaf Jurusan tidak boleh sama")
			fmt.Print("Tolong masukkan ulang jurusan disini: ")
			fmt.Scanln(&akun[i].jurusan[k])
			for akun[i].jurusan[k] > 10 || akun[i].jurusan[k] < 0 {
				fmt.Println("Maaf Daftar jurusan tidak ada")
				fmt.Print("Tolong masukkan ulang jurusan disini: ")
				fmt.Scanln(&akun[i].jurusan[k])
			}
		}
		temp = akun[i].jurusan[k]
		pendaftar[temp] = pendaftar[temp] + 1
	}
}
func jurusan_ke_huruf(i, k int, huruf *string, nilai *float64) {
	switch akun[i].jurusan[k] {
	case 1:
		*huruf = "S1 Informatika"
		*nilai = 85
	case 2:
		*huruf = "S1 Teknik Telekomunikasi"
		*nilai = 83
	case 3:
		*huruf = "S1 Teknik Elektro"
		*nilai = 85
	case 4:
		*huruf = "S1 Teknik Industri"
		*nilai = 80
	case 5:
		*huruf = "S1 Sistem Informasi"
		*nilai = 78
	case 6:
		*huruf = "S1 Manajemen Bisnis Telekomunikasi dan Informasi"
		*nilai = 83
	case 7:
		*huruf = "S1 Akuntansi"
		*nilai = 80
	case 8:
		*huruf = "S1 Ilmu Komunikasi"
		*nilai = 79
	case 9:
		*huruf = "S1 Administrasi Bisnis"
		*nilai = 80
	case 10:
		*huruf = "S1 Desain Komunikasi Visual"
		*nilai = 78
	}
}