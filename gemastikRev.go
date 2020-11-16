package main

import (
	"fmt"
)

const N = 200
var JumlahTim int 

var array ArrayGlobal
type ArrayGlobal [N]Tim

type Tim struct {
	Jenis string
	Nama string
	Univ string
	Skor int
}
type Penyisihan struct {
	ListTim1 		[N]Tim
	JumlahTim2      int
	JumlahTimLolos  int
}
type Semifinal struct {
	ListTim2       [N]Tim
	JumlahTim3     int
	JumlahTimLolos int
}

type Final struct {
	ListTim3   [N]Tim
	JumlahTim4 int
}
type Lolos struct {
	Final      Final
	Semifinal  Semifinal
	Penyisihan Penyisihan
}

type Database struct {
	arraydata   [N]Tim
	
	JumlahLomba int
}

func Create(array *ArrayGlobal, Input *Tim) {
	var (
		namatim, tanya, namauniv, jenis string
		skor,jmltim, jmluniv, jmlskor, countjenis int
		
	)
	status := true
	
	for status == true {
		fmt.Println("masukkan nama tim : ")
		fmt.Scan(&namatim)
		array[jmltim].Nama = namatim
		
		jmltim += 1
		JumlahTim += 1
		fmt.Println("masukkan asal universitas : ")
		fmt.Scan(&namauniv)
		array[jmluniv].Univ = namauniv
		
		jmluniv += 1
		fmt.Println("masukkan nilai skor : ")
		fmt.Scan(&skor)
		array[jmlskor].Skor = skor
		
		jmlskor += 1
		fmt.Printf("pilih jenis lomba\n"+
				   "1. DATA MINING\n"+
				   "2. PEMROGRAMAN\n"+
				   "3. ANIMASI\n")
		fmt.Println("ketik jenis lomba : ")
		fmt.Scan(&jenis)
		array[countjenis].Jenis=jenis

		countjenis += 1
		fmt.Println("masukkan tim lagi ? *(iya/tidak)")
		fmt.Scan(&tanya)
		if tanya == "iya" {
			status = true
		} else if tanya == "tidak" {
			status = false
		}
	}

}


func Sorting(tab *ArrayGlobal, JumlahTim int ) {

	fmt.Println("proses dimulai")
	fmt.Println("=====================")
	var (min, j int
		kosong Tim
	)

	for i:=0;i<JumlahTim-1;i++ {
		min=i
		j=i+1
		for j<JumlahTim {
			if tab[j].Skor < tab[min].Skor {
				min=j

			}
			j++
		}
		kosong=tab[min]
		tab[min]=tab[i]
		tab[i]=kosong
		
	}
}

func Search(array *ArrayGlobal, cari string) int {
	fmt.Printf("mulai search\n"+
			   "===================\n")
	
	var hasil, awal, tgh, akhir int
	
	awal = 0
	akhir = JumlahTim-1
	tgh = (awal + akhir) / 2
	for akhir >= tgh && cari != array[tgh].Nama {
		if cari > array[tgh].Nama {
			awal = tgh+1
		}else{
			akhir = tgh-1
		}
		tgh = (awal+akhir) / 2
	}
	if cari == array[tgh].Nama {
		hasil=tgh
	}
	return hasil
}
func edit(array *ArrayGlobal, hasil int) {
	var (
		tanya, nama, univ string
			skor int
		)
		
		fmt.Printf("Apakah ingin merubah nama Tim ? *(iya/tidak) %v\n " , array[hasil].Nama)
		fmt.Scan(&tanya)
		if tanya == "iya" {
			fmt.Println("Nama tim yang diubah: ", array[hasil].Nama)
			fmt.Print("Nama baru: ")
			fmt.Scan(&nama)
			array[hasil].Nama = nama
		}
		fmt.Println("Apakah ingin merubah nama Universitas ? *(iya/tidak) ")
		fmt.Scan(&tanya)
		if tanya == "iya" {
			fmt.Println("Nama universitas yang diubah: ", array[hasil].Univ)
			fmt.Print("Nama baru: ")
			fmt.Scan(&univ)
			array[hasil].Univ = univ
		}
		fmt.Println("Apakah ingin merubah nilai Skor ? *(iya/tidak) ")
		fmt.Scan(&tanya)
		if tanya == "iya" {
			fmt.Println("Nilai skor yang diubah: ", array[hasil].Skor)
			fmt.Print("Skor baru: ")
			fmt.Scan(&skor)
			array[hasil].Skor = skor
		}
}
func delete(array *ArrayGlobal, hasil int) {
	var (
		tanya string
		kosong string
	)
	kosong = " "
	fmt.Println("Apakah ingin menghapus nama Tim ? *(iya/tidak) ")
	fmt.Scan(&tanya)
	if tanya == "iya" {
		fmt.Println("Nama Tim yang dihapus :", array[hasil].Nama)
		array[hasil].Nama = kosong
	}
	fmt.Println("Apakah ingin menghapus nama Universitas ? *(iya/tidak) ")
	fmt.Scan(&tanya)
	if tanya == "iya" {
		fmt.Println("Nama Universitas yang dihapus: ", array[hasil].Univ)
		array[hasil].Univ = kosong
	}
	fmt.Println("Apakah ingin menghapus nilai Skor ? *(iya/tidak) ")
	fmt.Scan(&tanya)
	if tanya == "iya" {
		fmt.Println("Nilai skor yang dihapus: ", array[hasil].Skor)
		array[hasil].Skor = 0
	}		
}
func penyisihan(array *ArrayGlobal, arrayP *Penyisihan, JumlahTim int) {
	var (
		 temp1 string
		 temp2 string
		 temp3 int
	)

	i := 0 
	j := 0
	for i <JumlahTim-1 {
		if  array[i].Skor >= 125 {
			temp1 = array[i].Nama
			arrayP.ListTim1[j].Nama = temp1
			temp2 = array[i].Univ
			arrayP.ListTim1[j].Univ = temp2
			temp3 = array[i].Skor
			arrayP.ListTim1[j].Skor = temp3
			j += 1	
		}
		i += 1
	}	
}
func semifinal(arraylist *Penyisihan, arraySemi *Semifinal, JumlahTim int) {
	var (
		temp1 string
		temp2 string
		temp3 int
   )

   i := 0 
   j := 0
   for i <JumlahTim-1 {
	   if  arraylist.ListTim1[i].Skor >= 350 {
		   temp1 = arraylist.ListTim1[i].Nama
		   arraySemi.ListTim2[j].Nama = temp1
		   temp2 = arraylist.ListTim1[i].Univ
		   arraySemi.ListTim2[j].Univ = temp2
		   temp3 = arraylist.ListTim1[i].Skor
		   arraySemi.ListTim2[j].Skor = temp3
		   j += 1	
	   }
	   i += 1
   }	
}
func final(arraySemi *Semifinal, arrayF *Final, JumlahTim int) {
	var (
		temp1 string
		temp2 string
		temp3 int
   )

   i := 0 
   j := 0
   for i <JumlahTim-1 {
	   if  arraySemi.ListTim2[i].Skor >= 500 {
		   temp1 = arraySemi.ListTim2[i].Nama
		   arrayF.ListTim3[j].Nama = temp1
		   temp2 = arraySemi.ListTim2[i].Univ
		   arrayF.ListTim3[j].Univ = temp2
		   temp3 = arraySemi.ListTim2[i].Skor
		   arrayF.ListTim3[j].Skor = temp3
		   j += 1	
	   }
	   i += 1
   }	
}


func main() {

	
	var (
		
		arrayP Penyisihan
		arraySemi Semifinal
		arrayF Final
		Input Tim
		hasil int
		cari string
		menu int		
	)
	
	

fmt.Printf("\n"+
"Pilih Menu : \n"+
"	1. Input data Tim\n"+
"	2. Cari data Tim\n"+
"	3. Edit data Tim\n"+ 
"	4. Delete data Tim\n"+
"	5. Tim lolos penyisihan\n"+
"	6. Tim lolos semifinal\n"+
"	7. Tim pemenang\n"+
"	8. Tampilkan data lomba"
"	9. Exit menu\n") 
fmt.Println("pilih menu nomer berapa? *isi data jika belum")
fmt.Scan(&menu)
for menu != 9 {
	if menu == 1 {
		Create(&array, &Input)
		Sorting(&array, JumlahTim)
		i := 0
		fmt.Println("data tim yang ada")
	
		
		for i < JumlahTim {
			fmt.Println(array[i])
			i += 1
		}
		
		
	} 
	if menu == 2 {
		Sorting(&array, JumlahTim)
		fmt.Print("nama tim yang ingin dicari :")
		fmt.Scan(&cari)
		hasil=Search(&array, cari)
		fmt.Println(array[hasil])
	}
	if menu == 3 {
		Sorting(&array, JumlahTim)
		fmt.Print("nama tim yang ingin diedit :")
		fmt.Scan(&cari)
		hasil=Search(&array, cari)
		edit(&array, hasil)
		Sorting(&array, JumlahTim)
		fmt.Println(array[hasil])
	}
	if menu == 4 {
		Sorting(&array, JumlahTim)
		fmt.Print("nama tim yang ingin didelete :")
		fmt.Scan(&cari)
		hasil=Search(&array, cari)
		delete(&array, hasil)
		Sorting(&array, JumlahTim)
		Sorting(&array, JumlahTim)
		j:=0
		for j<JumlahTim {
			fmt.Println(array[j])
			j++
		}
	}
	if menu == 5 {
		penyisihan(&array, &arrayP, JumlahTim)
		k:=0
		for k<JumlahTim {
			fmt.Println(arrayP.ListTim1[k])
			k++
		}
	}
	if menu == 6 {
		semifinal(&arrayP, &arraySemi, JumlahTim)
		l:=0
		for l<JumlahTim {
			fmt.Println(arraySemi.ListTim2[l])
			l++
		}
	} 
	if menu == 7 {
		final(&arraySemi, &arrayF, JumlahTim) 
		m:=0
		for m<JumlahTim {
			fmt.Println(arrayF.ListTim3[m])
			m++
		}
	if menu == 9 {
		
	}
	}
	fmt.Printf("\n"+
	"Pilih Menu : \n"+
	"	1. Input data Tim\n"+
	"	2. Cari data Tim\n"+
	"	3. Edit data Tim\n"+ 
	"	4. Delete data Tim\n"+
	"	5. Tim lolos penyisihan\n"+
	"	6. Tim lolos semifinal\n"+
	"	7. Tim pemenang\n"+
	"	8. Exit menu\n") 
	fmt.Println("pilih menu nomer berapa? *isi data jika belum")
	fmt.Scan(&menu)
}
}