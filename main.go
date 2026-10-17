package main

import (
	"fmt"
	"my/http_server"
)

func main() {
	// connString := os.Getenv("CONN_STRING")
	// if connString != "" {
	// 	fmt.Println("val:", connString)
	// } else {
	// 	fmt.Println("val пустой")
	// }

	// ctx := context.TODO()
	// conn, err := simpleconnection.ConnectDB(ctx, connString)
	// if err != nil {
	// 	log.Fatal("cannot connect DB:", err)
	//}

	// if err := simplesql.CreateTable(ctx, conn); err != nil {
	// 	log.Fatal("cannot create table:", err)
	// }

	// // if err = simplesql.InsertRow(ctx, conn, "Обед", "Покушат", false, time.Now()); err != nil {
	// // 	log.Fatal("cannot insert values", err)
	// // }

	// // if err := simplesql.UpdateRow(ctx, conn, 3); err != nil {
	// // 	log.Fatal("cannot updatee row:", err)
	// // }

	// // if err := simplesql.DeleteRow(ctx, conn, 1); err != nil {
	// // 	log.Fatal("cannot delete row:", err)
	// // }

	// tasks, err := simplesql.SelectRows(ctx, conn)
	// if err != nil {
	// 	log.Fatal("cannot select rows:", err)
	// }

	// // for _, task := range tasks {
	// // 	if task.Id == 3 {
	// // 		task.Title = "Бородааа"
	// // 		task.Description = "аааааааааааааааааааааааа"
	// // 		task.Completed = true
	// // 		now := time.Now()
	// // 		task.CompletedAt = &now

	// // 		if err := simplesql.UpdateTask(ctx, conn, task); err != nil {
	// // 			panic(err)
	// // 		}
	// // 		break
	// // 	}
	// // }

	// fmt.Println(tasks)

	// ------- Docker -------

	// fmt.Println("Запуск сервера 2")

	if err := http_server.StartHttpServer(); err != nil {
		panic(err)
	} else {
		fmt.Println("Сервер завершился корректно")
	}

	// _, err := os.Create("out/newfile.txt")
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Done")
}
