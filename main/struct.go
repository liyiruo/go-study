package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	var a = Books{"111", "222", "444", 333}
	fmt.Println("这个是一种定义方式", a)
	var b = Books{
		title:   "string",
		author:  "string",
		subject: "string",
		bookId:  123,
	}
	fmt.Println("这个是一种定义方式", b)

	var c = Books{
		author:  "wo shi liyiruo",
		subject: "string",
		bookId:  123,
	}
	fmt.Println("这个是一种定义方式", c)
	fmt.Println("这个是一种定义方式", c.author)

	var Books1 Books
	var Books2 Books

	Books1.author = "张三"
	Books1.bookId = 1234
	Books1.subject = "主题"
	Books1.title = "标题"

	Books2.author = "李四"
	Books2.bookId = 12345
	Books2.subject = "主题1"
	Books2.title = "标题1"

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Books1.title)
	fmt.Printf("Book 1 author : %s\n", Books1.author)
	fmt.Printf("Book 1 subject : %s\n", Books1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Books1.bookId)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Books2.title)
	fmt.Printf("Book 2 author : %s\n", Books2.author)
	fmt.Printf("Book 2 subject : %s\n", Books2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Books2.bookId)
}
