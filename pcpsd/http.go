package pcpsd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pcps/internal/conf"
	"runtime"
	"time"

	"github.com/dgrijalva/jwt-go"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//StartHTTPServer 启动http服务
func StartHTTPServer(w io.Writer) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		//格式化日志 暂时记录全部内容
		// Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: w,
	}))

	//注册路由
	e.GET("/", accessible)
	e.POST("/login", login)
	e.GET("/ws", startWebsocketServer)
	e.GET("/list", list)
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login" || c.Path() == "/list" {
				return true
			}
			return false
		},
	}
	e.Use(middleware.JWTWithConfig(config))

	e.Logger.Fatal(e.Start(":" + conf.GetString("port")))
}

type WyCamerainfo struct {
	Mac              sql.NullString `db:"mac"`
	ParentMac        sql.NullString `db:"parent_mac"`        //父mac
	CmdCode          sql.NullString `db:"cmd_code"`          //指令code
	Location         sql.NullString `db:"location"`          //车位
	Etime            sql.NullString `db:"e_time"`            //执行时间
	Status           sql.NullInt64  `db:"status"`            //状态
	Code             sql.NullInt64  `db:"code"`              // codes
	Version          sql.NullString `db:"version"`           // 版本
	LedStatus        sql.NullInt64  `db:"led_status"`        // led状态
	CarNumber        sql.NullString `db:"car_number"`        // 车牌信息
	NumberURL        sql.NullString `db:"number_url"`        // 图片url
	OppositeLocation sql.NullString `db:"opposite_location"` // 对面车位
	IdentifyNumber   sql.NullString `db:"identify_number"`   // 确认编号
	Modal            sql.NullInt64  `db:"modal"`             // 启动模式
	CreateTime       sql.NullString `db:"create_time"`       // 创建时间
	UpdateTime       sql.NullString `db:"update_time"`       // 更新时间
}

var (
	//数据库操作对象
	Db *sqlx.DB
)

func initDb() {
	// database, err := sqlx.Open("mysql", "root:root@tcp(10.10.83.162:3306)/etm-m-om")
	sqlSession, err := sqlx.Open("mysql", "root:chenle07.@tcp(127.0.0.1:3306)/camera_log")
	SimplePanic(err)
	if err != nil {
		panic("open database error!")
	}

	if err := sqlSession.Ping(); err != nil {
		panic("connection database fail!")
	}

	Db = sqlSession
}

func SimplePanic(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(file, line, err)
		runtime.Goexit()
	}

}

func list(c echo.Context) error {
	initDb()
	s := c.FormValue("start_time")
	e := c.FormValue("end_time")
	var w []WyCamerainfo
	// db, err := sql.Open("mysql", "root:chenle07.@/local?charset=utf8")
	selectSql := "SELECT * FROM wy_camerainfo WHERE create_time BETWEEN ? AND ?"
	// selectSql := "SELECT * FROM wy_camerainfo WHERE create_time BETWEEN ? AND ? LIMIT 10"
	if err := Db.Select(&w, selectSql, s, e); err != nil {
		fmt.Println(err)
		return err
	}

	j, _ := json.Marshal(w)

	return c.String(http.StatusOK, string(j))
	// selectSql := "select * from t_user where id = ?"
	// if err := db.Get(&WyCamerainfo, selectSql, uid); err != nil {

	// }

	// requested_id := c.Param("id")
	// fmt.Println(requested_id)
	// db, err := sql.Open("mysql", "root:wangshubo@/test?charset=utf8")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	response := Excuse{Id: "", Error: "true", Quote: ""}
	// 	return c.JSON(http.StatusInternalServerError, response)
	// }
	// defer db.Close()

	// var quote string
	// var id string
	// err = db.QueryRow("", requested_id).Scan(&id, &quote)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// response := Excuse{Id: id, Error: "false", Quote: quote}
	// return c.JSON(http.StatusOK, response)

	// return c.String(http.StatusOK, "dev_list")
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!xx" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}
