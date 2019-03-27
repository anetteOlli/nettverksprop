package main

/**
For å kjøre programmet må man først kjøre i terminal komandoen:
 go get "github.com/jinzhu/gorm"
go get "github.com/go-sql-driver/mysql"


 */
import (
	"bufio"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"strconv"
	"strings"
)

type konto struct {
	//gorm.Model
	Kontonummer int `gorm:"primary_key";"AUTO_INCREMENT"`
	Kunde string `gorm:"type:varchar(150)"`
	Penger int
}


func main(){
	var rikeKontoer []konto
	db, err := OpprettForbindelse()
	if err !=nil {
		return
	}
	defer db.Close()

	//db.Debug().DropTableIfExists() vs db.DropTableIfExists() --> debug() skriver ut sql spørringa
	//db.Debug().DropTableIfExists(&konto{}) //drop tables if they exists
	//db.Debug().AutoMigrate(&konto{}) //auto creates tables based on struct konto

	fortsette := true
	for fortsette{
		reader := bufio.NewReader(os.Stdin) //tilsvarer java scanner
		fmt.Print("for å legge inn person i db skriv ADD, \n for å slette skriv DELETE, \n for å overføre penger skriv TRANSFER,\n for å hente ut en bruker, skriv GETONE,\n for å endre navn CHANGE,\nfor å avslutte skriv STOP")
		valg,_:=reader.ReadString('\n')
		switch valg {
		case "ADD\n":
			fmt.Print("skriv inn navnet på eier av konto")
			navnRead, _ := reader.ReadString('\n')
			navn := strings.Trim(navnRead, "\n")
			fmt.Print("skriv inn beløp på kontoen:")
			moneyRead, _ := reader.ReadString('\n')
			moneyReadTrim :=strings.Trim(moneyRead, "\n")
			money, err := strconv.Atoi(moneyReadTrim) //konverterer string til int
			for err!=nil{
				fmt.Print("skriv inn beløp på kontoen, beløpet må være et heltall")
				moneyRead, _ = reader.ReadString('\n')
				moneyReadTrim = strings.Trim(moneyRead, "\n")
				money, err = strconv.Atoi(moneyReadTrim)
			}
			LagBruker(db, navn, money)
		case "DELETE\n":
			fmt.Print("skriv navnet på brukeren:")
			navnRead,_ := reader.ReadString('\n')
			navn :=strings.Trim(navnRead, "\n")
			SlettBruker(db, navn)
		case "CHANGE\n":
			fmt.Print("skriv inn det gamle brukernavnet")
			oldNavnRead, _ := reader.ReadString('\n')
			oldNavn := strings.Trim(oldNavnRead, "\n")
			fmt.Print("skriv inn det nye navnet")
			newNameRead,_:=reader.ReadString('\n')
			newNavn := strings.Trim(newNameRead, "\n")
			OppdaterNavn(db, oldNavn, newNavn)
		case "TRANSFER\n":
			fmt.Print("hvem skal overføre penger?")
			fmt.Print("hvem skal overføre penger?")
			donorRead,_:=reader.ReadString('\n')
			donor :=strings.Trim(donorRead,"\n")
			fmt.Print("hvem skal få penger?")
			motakerRead,_:=reader.ReadString('\n')
			motaker := strings.Trim(motakerRead,"\n")
			fmt.Print("hvor mye penger skal overføres?")
			pengerRead,_:=reader.ReadString('\n')
			pengerRead = strings.Trim(pengerRead, "\n")
			penger, err := strconv.Atoi(pengerRead)
			for err!=nil{
				fmt.Print("du må skrive inn et heltall")
				pengerRead,_ = reader.ReadString('\n')
				pengerRead = strings.Trim(pengerRead, "\n")
				penger, err = strconv.Atoi(pengerRead)
			}
			TransferMOney(db, donor, motaker, penger)
		case "GETONE\n":
			fmt.Print("skriv inn navn")
			navnRead,_ := reader.ReadString('\n')
			navn := strings.Trim(navnRead, "\n")
			bruker := GetKonto(db, navn)
			fmt.Print(bruker.Kontonummer, bruker.Kunde, strconv.Itoa(bruker.Penger) + "\n")
		case "RIKE\n":
			fmt.Print("hvor rik?, skriv inn minste beløp")
			pengerRead, _ := reader.ReadString('\n')
			pengerRead = strings.Trim(pengerRead, "\n")
			penger, err := strconv.Atoi(pengerRead)
			for err!=nil{
				fmt.Print("du må skrive inn et heltall")
				pengerRead, _ = reader.ReadString('\n')
				pengerRead = strings.Trim(pengerRead, "\n")
				penger, err = strconv.Atoi(pengerRead)
			}
			rikeKontoer = GetRikeKontoer(db, penger)
			fmt.Print(rikeKontoer)

		default:
			fmt.Print("du valgte å avslutte ", valg)
			fortsette = false
		}
	}



}

func LagBruker(db *gorm.DB, navn string, penger int){

	fmt.Print("kommet så langt, lest verdiene: " + strconv.Itoa(penger) + "  " + navn)
	person :=&konto{Kunde:navn, Penger:penger} //lager et person objekt av konto type
	db.Debug().Create(person) //opretter person i databasen
	return
}

func SlettBruker(db *gorm.DB, navn string){

	person :=&konto{}
	//First() vs Find(): First() finner kjører med LIMIT 1
	db.Debug().First(&person, "Kunde=?", navn) //ikke spesielt funksjonelt men dette tar altså å legger inn sql resultatet inn i person
	fmt.Print(person.Kontonummer, person.Penger, person.Kunde)

	if person.Kontonummer !=0 {
		//NB! hvis raden vi skal slette mangler primary key, så slettes alt i tabellen
		db.Debug().Delete(&person)
	}else{
		fmt.Print("personen finnes ikke")
	}
}
func OppdaterNavn(db *gorm.DB, oldName string, newName string){
	var person konto
	db.Debug().First(&person, "Kunde=?", oldName)
	fmt.Print(person.Kunde, person.Kontonummer, person.Penger)
	person.Kunde = newName
	db.Debug().Save(&person)
}

func TransferMOney(db *gorm.DB, donor string, mottaker string, penger int){

	donorPerson :=&konto{}
	db.Debug().First(&donorPerson, "Kunde=?", donor)
	//her burde man sikkert ha hatt en sjekk på at kunden faktisk finnes, MEEEEN det driter jeg i

	mottakerPerson :=&konto{}
	db.Debug().First(&mottakerPerson, "Kunde=?", mottaker)
	//igjen her mangler det en del if-sjekker, men de bryr vi oss ikke om

	//her får man lov til å overføre mer penger enn man har, null stress med minus på konto
	donorPerson.Penger = donorPerson.Penger - penger
	mottakerPerson.Penger = mottakerPerson.Penger + penger

	//så bare lagre dette i databasen og så er vi good
	db.Debug().Save(&donorPerson)
	db.Debug().Save(&mottakerPerson)

}

func OpprettForbindelse() (*gorm.DB, error){
	mysqlAdress := "anettosi:LN8iIcr6@tcp(mysql.stud.iie.ntnu.no:3306)/anettosi?charset=utf8&parseTime=True&loc=Local"
	db, err :=gorm.Open("mysql", mysqlAdress)
	//defer db.Close() //lukker db for oss når vi er ferdig med den for oss, trenger ikke å ha db.close() på slutten av funksjonen som vi må i java
	if err != nil{
		fmt.Print("Connection failed to open   ", err)
		return db, err
	}
	fmt.Print("connection established")
	return db, nil
}

func GetKonto(db *gorm.DB, navn string)konto{
	var konto konto
	db.Debug().First(&konto, "Kunde=?", navn)
	return konto
}
func GetRikeKontoer(db *gorm.DB,  penger int)[]konto{
	var kontoer []konto
	db.Debug().Find(&kontoer, "Penger>=?", penger)
	return kontoer
}
func WipeDatabase(db *gorm.DB){
	db.Debug().DropTableIfExists(&konto{})
}
