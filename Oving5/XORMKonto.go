package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Account struct {
	//gorm.Model
	Acccount int64 `xorm:"pk autoincr"`
	Customer string `xorm:"varchar(150)"`
	Money int64
}

func main(){
	engine, err := CreateConnection()
	if err!= nil{
		return
	}
	ResetDatabase(engine)
	AddUser(engine, "ole", 900)
}

func AddUser(engine *xorm.Engine, Customer string, Money int64){
	newAccount := Account{Customer:Customer, Money:Money}
	affected, err := engine.Insert(newAccount)
	if err!=nil{
		fmt.Print("error occured while adding new customer " ,err, "\n")
	}else if affected !=1{
		fmt.Print("more or less than one customer was added: ", affected, "\n")
	}else{
		fmt.Print("user added\n")
	}
}

func ChangeUserName(engine *xorm.Engine, oldName string, newName string){
	
}

func ResetDatabase(engine *xorm.Engine){
	engine.DropTables(Account{})
	engine.CreateTables(Account{})
}

func CreateConnection() (*xorm.Engine, error) {
	mysql := "anettosi:LN8iIcr6@tcp(mysql.stud.iie.ntnu.no:3306)/anettosi?charset=utf8&parseTime=True&loc=Local"
	engine, err := xorm.NewEngine("mysql", mysql)
	if err != nil{
		fmt.Print("error connecting: ", err, "\n")
		return engine, err
	}
	fmt.Print("connection established\n")
	return engine, err
}
