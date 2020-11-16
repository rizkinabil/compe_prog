package main

import (
	"fmt"
)

const N = 200
var JumlahTim, JumlahSkor, jmlJenis1,jmlJenis2,jmlJenis3 int 
var arrLomba Jenislomba

type ArrayGlobal [N]Tim

type Tim struct {
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
type Jenislomba struct{
	Datamining [N]Tim
	Pemrograman [N]Tim
	Animasi [N]Tim
}
type Database struct {
	arraydata   [N]Tim
	
	JumlahLomba int
}

func Create(Array *ArrayGlobal, input *Tim) {
	var (
		namatim, tanya, namauniv string
		skor,jmltim, jmluniv, jmlskor, jenis int
		
	)
	status := true
	
	for status == true {

		fmt.Printf("pilih jenis lomba\n"+
				   "1. DATA MINING\n"+
				   "2. PEMROGRAMAN\n"+
				   "3. ANIMASI\n")
		fmt.Println("pilih jenis lomba : *(pilihan jenis lomba yang diisi harus urut dari 1-3)")
		fmt.Scan(&jenis)
		if jenis == 1 {
			fmt.Println("LOMBA DATA MINING")
			fmt.Println("======================")
			fmt.Println("masukkan nama tim : ")
			fmt.Scan(&namatim)
			Array[JumlahTim].Nama = namatim
			arrLomba.Datamining[JumlahTim].Nama = namatim
			jmltim += 1
			fmt.Println("masukkan asal universitas : ")
			fmt.Scan(&namauniv)
			Array[JumlahTim].Univ = namauniv
			arrLomba.Datamining[JumlahTim].Univ = namauniv
			jmluniv += 1
			fmt.Println("masukkan nilai skor : ")
			fmt.Scan(&skor)
			Array[JumlahTim].Skor = skor
			arrLomba.Datamining[JumlahTim].Skor = skor
			jmlskor += 1
			JumlahTim += 1
			jmlJenis1++
		}
		if jenis == 2 {
			fmt.Println("LOMBA PEMROGRMAN")
			fmt.Println("======================")
			fmt.Println("masukkan nama tim : ")
			fmt.Scan(&namatim)
			Array[JumlahTim].Nama = namatim
			arrLomba.Pemrograman[JumlahTim].Nama = namatim
			jmltim += 1
			fmt.Println("masukkan asal universitas : ")
			fmt.Scan(&namauniv)
			Array[JumlahTim].Univ = namauniv
			arrLomba.Pemrograman[JumlahTim].Univ = namauniv
			jmluniv += 1
			fmt.Println("masukkan nilai skor : ")
			fmt.Scan(&skor)
			Array[JumlahTim].Skor = skor
			arrLomba.Pemrograman[JumlahTim].Skor = skor
			jmlskor += 1
			jmlJenis2++
			JumlahTim += 1
		}
		if jenis == 3 {
			fmt.Println("LOMBA ANIMASI")
			fmt.Println("======================")
			fmt.Println("masukkan nama tim : ")
			fmt.Scan(&namatim)
			Array[JumlahTim].Nama = namatim
			arrLomba.Animasi[JumlahTim].Nama = namatim
			jmltim += 1
			fmt.Println("masukkan asal universitas : ")
			fmt.Scan(&namauniv)
			Array[JumlahTim].Univ = namauniv
			arrLomba.Animasi[JumlahTim].Univ = namauniv
			jmluniv += 1
			fmt.Println("masukkan nilai skor : ")
			fmt.Scan(&skor)
			Array[JumlahTim].Skor = skor
			arrLomba.Animasi[JumlahTim].Skor = skor
			jmlskor += 1
			jmlJenis3++
			JumlahTim += 1
		}
		
		JumlahSkor += 1
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

func SortingDM(jmlJenis1 int ) {

	var (min, j int
		kosong Tim
	)

	for i:=0;i<jmlJenis1-1;i++ {
		min=i
		j=i+1
		for j<jmlJenis1 {
			if arrLomba.Datamining[j].Skor < arrLomba.Datamining[min].Skor {
				min=j

			}
			j++
		}
		kosong=arrLomba.Datamining[min]
		arrLomba.Datamining[min]=arrLomba.Datamining[i]
		arrLomba.Datamining[i]=kosong
		
	}
}

func SortingPR(jmlJenis2 int ) {
	var (min, j int
		kosong Tim
	)

	for i:=0;i<jmlJenis2-1;i++ {
		min=i
		j=i+1
		for j<jmlJenis2 {
			if arrLomba.Pemrograman[j].Skor < arrLomba.Pemrograman[min].Skor {
				min=j

			}
			j++
		}
		kosong=arrLomba.Pemrograman[min]
		arrLomba.Pemrograman[min]=arrLomba.Pemrograman[i]
		arrLomba.Pemrograman[i]=kosong
		
	}
}

func SortingAM(jmlJenis3 int ) {
	var (min, j int
		kosong Tim
	)

	for i:=0;i<jmlJenis3-1;i++ {
		min=i
		j=i+1
		for j<jmlJenis3 {
			if arrLomba.Animasi[j].Skor < arrLomba.Animasi[min].Skor {
				min=j

			}
			j++
		}
		kosong=arrLomba.Animasi[min]
		arrLomba.Animasi[min]=arrLomba.Animasi[i]
		arrLomba.Animasi[i]=kosong
		
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
		array ArrayGlobal
		arrayP Penyisihan
		arraySemi Semifinal
		arrayF Final
		input Tim
		hasil int
		cari string
		menu int	
		i int	
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
"	8. Exit menu\n") 
fmt.Println("pilih menu nomer berapa? *isi data jika belum")
fmt.Scan(&menu)
for menu != 8 {
	if menu == 1 {
		Create(&array, &input)
		Sorting(&array, JumlahTim)
		fmt.Println("data tim yang ada")
		fmt.Println("Lomba data mining :")
		SortingDM(jmlJenis1)
		i=0
		for i < JumlahTim {
			fmt.Println(arrLomba.Datamining[i])
			i += 1
		}
		fmt.Println("Lomba pemrograman :")
		SortingPR(jmlJenis2)
		i=1
		for i < JumlahTim {
			fmt.Println(arrLomba.Pemrograman[i])
			i += 1
		}
		fmt.Println("Lomba Animasi :")
		SortingAM(jmlJenis3)
		i=2
		for i < JumlahTim {
			fmt.Println(arrLomba.Animasi[i])
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
			k += 1
		}
	}
	if menu == 6 {
		semifinal(&arrayP, &arraySemi, JumlahTim)
		l:=0
		for l<JumlahTim {
			fmt.Println(arraySemi.ListTim2[l])
			l += 1
		}
	} 
	if menu == 7 {
		final(&arraySemi, &arrayF, JumlahTim) 
		m:=0
		for m<JumlahTim {
			fmt.Println(arrayF.ListTim3[m])
			m += 1
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