package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_  "github.com/go-sql-driver/mysql"
)

type  Book  struct{
	ID   int 					`orm:"column(id)"`
	Title  string				`orm:"column(title)"`
	Price  int 							`orm:"column(price)"`
}

func  init (){
	//1.连库
	err := orm.RegisterDataBase("default", "mysql",
		"root:123@tcp(127.0.0.1:3306)/my_book?charset=utf8", 50, 20)
	if err != nil{
		fmt.Println("err",err)
		return
	}

	//2.orm注册model

orm.RegisterModel(new(Book))

}
//指定表名,(指定你要操作的数据库的哪一个表)
func  (book *Book )  TableName ()string{
  return  "book"
}

//增删改查
func   (book *Book)Insert() error{
if 	_, err := orm.NewOrm().Insert(book);err !=nil{
	return  err
}
return nil
}

func   (book *Book)Delete() error{
	if 	_, err := orm.NewOrm().Delete(book);err !=nil{
		return  err
	}
	return nil
}

func   (book *Book)Update(fields ...string) error{
	if 	_, err := orm.NewOrm().Update(book,fields...);err !=nil{
		return  err
	}
	return nil
}

func   (book *Book)Read(fields ...string) error{
	if 	err := orm.NewOrm().Read(book,fields...);err !=nil{
		return  err
	}
	return nil
}