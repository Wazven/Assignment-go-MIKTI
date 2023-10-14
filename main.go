package main

import (
	"fmt"
	"os"
)

//Menyimpan data tagihan dan tip
type Bill struct{
	Tagihan float64
	Tip float64
}

func main() {
	//Deklarasi data Uji
	tagihan := []float64{275, 40, 430}
	
	//fungsi input data nilai tagihan
	fmt.Print("Masukan Nilai Tagihan: ")
	var inputTagihan float64
	_, err := fmt.Scanln(&inputTagihan)
	if err != nil {
		fmt.Println("Masukan Tidak Valid: ", err)
		return
	}

	tagihan = append(tagihan, inputTagihan)

	//hitung total tip dan total tagihan
	var bills []*Bill
    for _, tagihan := range tagihan {
        bills = append(bills, calculateTip(tagihan))
    }
    // Menampilkan hasil pada console
    for _, bill := range bills {
        fmt.Printf("Tagihan %.2f, tipnya %.2f, dan total nilainya %.2f\n", bill.Tagihan, bill.Tip, bill.Tagihan+bill.Tip)
    }

	// membuat file invoice.txt
	if err := Invoice(bills); err != nil {
		fmt.Println("Gagal: ", err)
	}
}

func calculateTip(tagihan float64) *Bill {
	bill := &Bill{Tagihan: tagihan}
	if tagihan >= 50 && tagihan <= 300 {
		bill.Tip = tagihan * 0.15 // untuk sebagai tip 15% apabila lebih dari sama dengan 50 dan kurang dari sama dengan 300
	} else {
		bill.Tip = tagihan * 0.2 // untuk sebagai tip 20% apabila kurang dari sama dengan 50 dan lebih dari sama dengan 300
	}
	return bill
}

func Invoice(bills []*Bill) error {
	invoice, err := os.Create("invoice.txt")
    if err != nil {
        return err
    }
    defer invoice.Close()

    for _, bill := range bills {
        _, err := invoice.WriteString(fmt.Sprintf("Tagihan: %.2f\nTip: %.2f\nTotal: %.2f\n\n", bill.Tagihan, bill.Tip, bill.Tagihan+bill.Tip))
        if err != nil {
            return err
        }
    }

    return nil
}