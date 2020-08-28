package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"good"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	println("Hello World")
	var a int = 2
	i := 1
	const c int = 10
	const (
		A = iota
		B
		C
	)

	println("a : ", a)
	println("i : ", i)
	i++
	println("i : ", i)

	println(A)
	println(B)
	println(C)

	rawLiteral := `good \n good`
	interLiteral := "good\ngood"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)

	k := 10
	p := &k
	println(&k)
	println(*p)
	println(p)

	if k > 5 {
		println("Good")
	} else if k < 5 {
		println("good!")
	} else {
		println("good")
	}

	if val := k * 2; val > 15 {
		println(val)
	}

	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	println(sum)

	names := []string{"good", "good!", "Good!"}

	for index, name := range names {
		println(index, name)
	}

	add := func(i int, j int) int {
		return i + j
	}

	r1 := calc(add, 10, 20)
	println(r1)

	r2 := calc(func(x int, y int) int { return x * y }, 2, 3)
	println(r2)

	next := nextValue()

	println(next())
	println(next())
	println(next())

	anotherNext := nextValue()
	println(anotherNext())
	println(anotherNext())

	//array
	var a1 = [3]int{1, 2, 3}
	var a2 = [...]int{4, 5, 6}
	a3 := [...]int{7, 8, 9}

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	//slice
	var s1 = []int{1, 2, 3, 4}
	s2 := []int{5, 6, 7, 8}
	s3 := make([]int, 5, 10)
	s4 := s1[1:3]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	println(len(s3), cap(s3))
	fmt.Println(s4)

	s4 = append(s4, 4, 5, 6)
	fmt.Println(s4)

	s5 := append(s1, s2...)
	fmt.Println(s5)

	var m map[int]string //Nil map
	m = make(map[int]string)
	m[1] = "Good"
	m[244] = "Great"

	car := map[string]string{
		"KIA":     "K3",
		"HYUNDAI": "AVANTE",
		"BMW":     "X3",
	}
	fmt.Println(m)
	str := m[244]
	println(str)

	_, exists := m[123]
	if !exists {
		println("No 123 Key value")
	}

	fmt.Println(car)

	for key, val := range car {
		fmt.Println(key, val)
	}

	p1 := person{}
	p1.name = "Lee"
	p1.age = 10

	var p2 person
	p2 = person{"Good", 50}

	p3 := person{name: "Great", age: 20}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	rect := Rect{10, 20}
	area := rect.area()
	area2 := rect.area2()
	println(area)
	println(area2)

	say("Sync")

	go say("Async1")
	go say("Async2")
	go say("Async3")

	time.Sleep(time.Second * 1)

	jsonBytes, err := json.Marshal(car)
	if err != nil {
		panic(err)
	}
	jsonString := string(jsonBytes)

	fmt.Println(jsonString)

	good.Good()

	db, err := sql.Open("mysql", "good:123456qwer!@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var id int
	var name string
	var value int
	err = db.QueryRow("SELECT name FROM test WHERE id = 2").Scan(&name)
	fmt.Println(name)
	rows, err := db.Query("SELECT * FROM test")
	defer rows.Close()

	fmt.Println("BEFORE ============")

	for rows.Next() {
		err := rows.Scan(&id, &name, &value)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, value)
	}

	/* == Insertion
	result, err := db.Exec("INSERT INTO test VALUES (?, ?, ?)", 7, "Un", 1111)
	if err != nil {
		log.Fatal(err)
	}
	n, err := result.RowsAffected()
	fmt.Println(n, "row inserted.")

	stmt, err := db.Prepare("INSERT INTO test VALUES (?, ?, ?)")
	checkError(err)
	defer stmt.Close()

	_, err = stmt.Exec(8, "Duex", 2222)
	checkError(err)
	_, err = stmt.Exec(9, "Trois", 3333)
	checkError(err) */

	fmt.Println("UPDATE ============")

	var first string
	err = db.QueryRow("SELECT name FROM test WHERE id = 7").Scan(&first)
	checkError(err)

	stmt, err := db.Prepare("UPDATE test SET name = ? WHERE id = ?")
	checkError(err)
	defer stmt.Close()

	if first == "Uno" {
		_, err = stmt.Exec("Un", 7)
		checkError(err)
		_, err = stmt.Exec("Duex", 8)
		checkError(err)
		_, err = stmt.Exec("Trois", 9)
		checkError(err)
	} else {
		_, err = stmt.Exec("Uno", 7)
		checkError(err)
		_, err = stmt.Exec("Dos", 8)
		checkError(err)
		_, err = stmt.Exec("Tres", 9)
		checkError(err)
	}

	newrows, err := db.Query("SELECT * FROM test")
	checkError(err)
	defer newrows.Close()

	for newrows.Next() {
		err := newrows.Scan(&id, &name, &value)
		checkError(err)
		fmt.Println(id, name, value)
	}

	tx, err := db.Begin()
	checkError(err)
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE test2 SET balance = balance - 500 WHERE account = 1101")
	checkError(err)
	_, err = tx.Exec("UPDATE test2 SET balance = balance + 500 WHERE account = 2101")
	checkError(err)

	err = tx.Commit()
	checkError(err)

	fmt.Println("TRANSACTION EXECUTED ====")

	var account int
	var balance int

	t2rows, err := db.Query("SELECT * FROM test2")
	checkError(err)
	defer t2rows.Close()

	for t2rows.Next() {
		err := t2rows.Scan(&id, &account, &balance)
		checkError(err)
		fmt.Println(id, account, balance)
	}
}

type calculator func(int, int) int

func calc(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type person struct {
	name string
	age  int
}

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
