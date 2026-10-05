package main
import "fmt"

type date struct {
	tanggal int
	bulan   string
	tahun   int
}

func main() {
	var tgl date
	tgl.tanggal = 10
	fmt.Scan(&tgl.bulan, &tgl.tahun)
	tgl.tanggal += 10
	fmt.Println(tgl.tanggal)
	fmt.Println(tgl.tanggal, tgl.bulan, tgl.tahun)

}