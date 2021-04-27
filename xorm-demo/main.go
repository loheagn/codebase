package xorm_demo

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var engine *xorm.Engine

type User struct {
	ID     int64
	Name   string
	Age    int
	Number int
}

func initEngine() (*xorm.Engine, error) {
	return xorm.NewEngine("sqlite3", "./test.db")
}

func init() {
	var err error
	err = os.Remove("./test.db")
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	engine.SetMapper(names.GonicMapper{})

	err = engine.Sync2(new(User))
	if err != nil {
		panic(err)
	}
}

func OrderDemo(order string) *User {
	users := []*User{
		{
			Name:   "Li",
			Age:    18,
			Number: 20,
		},
		{
			Name:   "Nan",
			Age:    18,
			Number: 22,
		},
		{
			Name:   "Zhao",
			Age:    18,
			Number: 21,
		},
	}
	_, err := engine.Insert(&users)
	if err != nil {
		panic(err)
	}

	user := &User{
		Age: 18,
	}

	has, err := engine.OrderBy(fmt.Sprintf("number %s", order)).Get(user)
	if err != nil || !has {
		panic(err)
	}

	for _, user := range users {
		_, err = engine.Delete(user)
		if err != nil {
			panic(err)
		}
	}

	return user
}
