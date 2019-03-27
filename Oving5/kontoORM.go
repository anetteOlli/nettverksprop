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
	mysqlAdress := "anettosi:LN8iIcr6@tcp(mysql.stud.iie.ntnu.no:3306)/anettosi?charset=utf8&parseTime=True&loc=Local"
	db, err :=gorm.Open("mysql", mysqlAdress)
	defer db.Close() //lukker db for oss når vi er ferdig med den for oss, trenger ikke å ha db.close() på slutten av funksjonen som vi må i java
	if err != nil{
		fmt.Print("Connection failed to open   ", err)
		return
	}
	fmt.Print("connection established")


	//db.Debug().DropTableIfExists(&konto{}) //drop tables if they exists
	//db.Debug().AutoMigrate(&konto{}) //auto creates tables based on struct konto

	fortsette := true
	for fortsette{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("for å legge inn person i db skriv ADD, \n for å slette skriv DELETE, \n for å overføre penger skriv TRANSFER, \n for å endre navn CHANGE,\nfor å avslutte skriv STOP")
		valg,_:=reader.ReadString('\n')
		switch valg {
		case "ADD\n":
			lagBruker(db)
		case "DELETE\n":
			slettBruker(db)
		case "CHANGE\n":
			oppdaterNavn(db)
		case "TRANSFER\n":
			transferMOney(db)
		default:
			fmt.Print("du valgte å avslutte ", valg)
			fortsette = false
		}
	}



}

func lagBruker(db *gorm.DB){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("skriv inn navnet på eier av konto")
	navnRead, _ := reader.ReadString('\n')
	navn := strings.Trim(navnRead, "\n")
	fmt.Print("skriv inn beløp på kontoen:")
	moneyRead, _ := reader.ReadString('\n')
	moneyReadTrim :=strings.Trim(moneyRead, "\n")
	money, err := strconv.Atoi(moneyReadTrim)
	for err!=nil{
		fmt.Print("skriv inn beløp på kontoen, beløpet må være et heltall")
		moneyRead, _ = reader.ReadString('\n')
		moneyReadTrim = strings.Trim(moneyRead, "\n")
		money, err = strconv.Atoi(moneyReadTrim)
	}
	fmt.Print("kommet så langt, lest verdiene: " + strconv.Itoa(money) + "  " + navn)
	person :=&konto{Kunde:navn, Penger:money}
	db.Debug().Create(person) //db.Create() vs db.Debug().Create() , debug skriver ut sql setninga ut til konsoll
	return

}
func slettBruker(db *gorm.DB){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("skriv navnet på brukeren:")
	navnRead,_ := reader.ReadString('\n')
	navn :=strings.Trim(navnRead, "\n")

	person :=&konto{}
	db.Debug().First(&person, "Kunde=?", navn) //ikke spesielt funksjonelt men dette tar altså å legger inn sql resultatet inn i person
	fmt.Print(person.Kontonummer, person.Penger, person.Kunde)
	db.Debug().Delete(&person)


}
func oppdaterNavn(db *gorm.DB){
	reader :=bufio.NewReader(os.Stdin)
	fmt.Print("skriv inn det gamle brukernavnet")
	navnRead, _ := reader.ReadString('\n')
	navn := strings.Trim(navnRead, "\n") //hive bort linjeskiftet fra navnet, og ja; dette er ugly

	person :=&konto{}
	db.Debug().First(&person, "Kunde=?", navn)
	fmt.Print(person.Kunde, person.Kontonummer, person.Penger)
	fmt.Print("skriv inn nytt navn:")
	navnRead, _ = reader.ReadString('\n')
	navn = strings.Trim(navnRead, "\n")
	person.Kunde = navn
	db.Debug().Save(&person)

}
func transferMOney(db *gorm.DB){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("hvem skal overføre penger?")
	donorRead,_:=reader.ReadString('\n')
	donor :=strings.Trim(donorRead,"\n")

	donorPerson :=&konto{}
	db.Debug().First(&donorPerson, "Kunde=?", donor)
	//her burde man sikkert ha hatt en sjekk på at kunden faktisk finnes, MEEEEN det driter jeg i

	fmt.Print("hvem skal få penger?")
	motakerRead,_:=reader.ReadString('\n')
	motaker := strings.Trim(motakerRead,"\n")
	motakerPerson :=&konto{}
	db.Debug().First(&motakerPerson, "Kunde=?", motaker)
	//igjen her mangler det en del if-sjekker, men de bryr vi oss ikke om

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

	//her får man lov til å overføre mer penger enn man har, null stress med minus på konto
	donorPerson.Penger = donorPerson.Penger - penger
	motakerPerson.Penger = motakerPerson.Penger + penger

	//så bare lagre dette i databasen og så er vi good
	db.Debug().Save(&donorPerson)
	db.Debug().Save(&motakerPerson)

}

