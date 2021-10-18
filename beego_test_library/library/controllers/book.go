package controllers


import (

	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"library/models"
)

type BookController struct {
	beego.Controller

}

func  (c *BookController)  List(){
	var  list  []*models.Book
	_, _ = orm.NewOrm().QueryTable(new(models.Book)).OrderBy("id").All(&list)
	//fmt.Println(list)
	c.Data["list"] = list
	c.TplName = "book_list.html"
	//c.TplName = "/book_html/book_list.html"
}


func (c *BookController)  Delete (){
	id, err := c.GetInt("id")
	if  err  !=nil{
		fmt.Println(err)
		return
	}
	book := &models.Book{ID:id}
if  err  =	book.Read();err ==nil{
	err := book.Delete()
	if err !=nil{
		fmt.Println(err)
		return
	}
	c.Redirect("/book/list",302)
}
}

func (c *BookController)  Add (){
//判断请求方式，进行跳页面和真正添加
	if c.Ctx.Request.Method == "GET"{
		c.TplName = "add_book.html"
	}else if c.Ctx.Request.Method == "POST"{
		//接收页面数据添加
		title := c.GetString("title")
		price, err := c.GetInt("price")
		if  err !=nil{
			fmt.Println("Add获取price数据出错",err)
			return
		}
		var  book =   &models.Book{Title : title,Price: price}
	if  err =	book.Insert(); err !=nil{
		fmt.Println("Insert添加数据出错",err)
		return
	}
		c.Redirect("/book/list",302)
	}else {
		c.Redirect("/book/list",302)
	}
}



func (c *BookController)  Edit (){
//接受更新数据的id
	id, err := c.GetInt("id")
	if  err  !=nil{
		fmt.Println(err)
		return
	}
	//查到该数据
	book := &models.Book{ID:id}
	if  err  := book.Read(); err !=nil{
		fmt.Println(err)
		return
	}
	//将根据id查到的数据，反到edit_book.html上
	c.Data["ID"] =book.ID
	c.Data["Title"] =book.Title
	c.Data["Price"] =book.Price
	c.TplName = "edit_book.html"
	//若是post请求，则更新
	if c.Ctx.Request.Method  == "POST"{
		//获取数据，进行更新
		id, err := c.GetInt("id")
		if  err  !=nil{
			fmt.Println(err)
			return
		}
		title := c.GetString("title")
		price ,err := c.GetInt("price")
		if  err  !=nil{
			fmt.Println(err)
			return
		}
		
		var   book =  &models.Book{ID: id,Title: title,Price: price}
		err = book.Update()
		if  err  !=nil{
			fmt.Println(err)
			return
		}

		c.Redirect("/book/list",302)
	}
}
