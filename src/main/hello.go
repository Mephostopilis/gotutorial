package main

import (
	"fmt"
	// "model"
	"github.com/astaxie/beego/orm"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int
	Name string `orm:"size(100)"`
}

var client *redis.Client

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

func main() {
	 o := orm.NewOrm()

    user := User{Name: "slene"}

    // insert
    id, err := o.Insert(&user)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    // update
    user.Name = "astaxie"
    num, err := o.Update(&user)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)

    // read one
    u := User{Id: user.Id}
    err = o.Read(&u)
    fmt.Printf("ERR: %v, Name: %v\n", err, u.Name)

    // // delete
    // num, err = o.Delete(&u)
    // fmt.Printf("NUM: %d, ERR: %v\n", num, err)
    ExampleNewClient()
    ExampleClient()

}


func ExampleNewClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}