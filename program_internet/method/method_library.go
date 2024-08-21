package method 

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
)

type User struct{
	Username string
	Password string
	Gender string
	Age int32
	BookName string
	Books string
	Date time.Time
}

func Begin_html(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/login.html")
}

func Login_html(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.FormValue("username"))
	// fmt.Println(r.FormValue("password"))

	// 检查请求方法是否为 POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// 获取html发送的数据格式
	fmt.Println("r.Header \"Content-Type\" = ", r.Header.Get("Content-Type"))

	// html网页提交的表单数据格式是application/x-www-form-urlencoded 或 multipart/form-data 格式，不可使用json解析
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&user)
	// if err != nil {
	// 	http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
	// 	fmt.Println("Error decoding JSON:", err)
	// 	return
	// }

	// 解析表单数据
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	user := User{ 
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
		Gender: r.FormValue("gender"),
		Age: func()int32{
			age, _ := strconv.Atoi(r.FormValue("age"))
			return int32(age)
		}(),
		BookName: r.FormValue("bookname"),
		Books: r.FormValue("books"),
		Date: func()time.Time{
			date, _ := time.Parse("2006-01-02", r.FormValue("date"))
			return date
		}(),
	}

	fmt.Println("user = ", user)

	// 调用函数进行插入
	InsertData(&user)
}