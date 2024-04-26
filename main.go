package main

import (
	"database/sql"
	"final-project/database"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "final-project"
)

var (
	DB  *sql.DB
	err error
)

// // func Auth(c *gin.Context) {
// // 	uname, pwd, ok := c.Request.BasicAuth()
// // 	if !ok {
// // 		c.String(http.StatusUnauthorized, "Username atau Password tidak boleh kosong")
// // 		c.Abort()
// // 		return
// // 	}

// // 	if (uname == "admin" && pwd == "password") || (uname == "editor" && pwd == "secret") {
// // 		c.Next()
// // 		return
// // 	}
// // 	c.String(http.StatusUnauthorized, "Username atau Password tidak sesuai")
// // 	c.Abort()
// // }

func main() {
	// err = godotenv.Load("config/.env")
	// if err != nil {
	// 	fmt.Println("failed load file environment")
	// } else {
	// 	fmt.Println("success read file environment")
	// }

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed: ", err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	router := gin.Default()

	// router.GET("/categories", controllers.GetAllCategories)
	// router.POST("/categories", Auth, controllers.InsertCategories)
	// router.PUT("/categories/:id", Auth, controllers.UpdateCategories)
	// router.DELETE("/categories/:id", Auth, controllers.DeleteCategories)
	// router.GET("/categories/:id/books", controllers.ShowCategoriesBook)

	// router.GET("/books", controllers.GetAllBook)
	// router.POST("/books", Auth, controllers.InsertBook)
	// router.PUT("/books/:id", Auth, controllers.UpdateBook)
	// router.DELETE("/books/:id", Auth, controllers.DeleteBook)

	router.Run("localhost:8080")
	// router.Run(":" + os.Getenv("PORT"))
}

// package main

// import (
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// )

// var jwtKey = []byte("your-secret-key")

// // User struct represents a user in the system
// type User struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// // Claims struct digunakan untuk menentukan payload token JWT
// type Claims struct {
// 	Username string `json:"username"`
// 	jwt.RegisteredClaims
// }

// // Database sementara untuk menyimpan pengguna yang terdaftar
// var users = map[string]string{
// 	"user1": "password1",
// 	"user2": "password2",
// }

// // RegisterHandler menangani pendaftaran pengguna baru
// func RegisterHandler(c *gin.Context) {
// 	var user User
// 	if err := c.BindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	users[user.Username] = user.Password
// 	c.Status(http.StatusCreated)
// }

// // LoginHandler menangani proses login pengguna
// func LoginHandler(c *gin.Context) {
// 	var user User
// 	if err := c.BindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	password, ok := users[user.Username]
// 	if !ok || password != user.Password {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
// 		return
// 	}

// 	token := GenerateToken(user.Username)
// 	c.Header("Authorization", "Bearer "+token)
// 	c.Status(http.StatusOK)
// }

// // GenerateToken digunakan untuk membuat token JWT
// func GenerateToken(username string) string {
// 	expirationTime := time.Now().Add(24 * time.Hour) // Token berlaku selama 24 jam
// 	claims := Claims{
// 		Username: username,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(expirationTime),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, _ := token.SignedString(jwtKey)
// 	return tokenString
// }

// func MiddlewareAuth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			c.Header("WWW-Authenticate", `Bearer realm="Restricted"`)
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 		// Mengambil token JWT dari header Authorization
// 		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

// 		// Memverifikasi token JWT
// 		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 			return jwtKey, nil
// 		})

// 		if err != nil || !token.Valid {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 		// Jika token valid, lanjutkan ke handler berikutnya
// 		c.Next()
// 	}
// }

// // CheckCredentials digunakan untuk memeriksa kredensial pengguna
// func CheckCredentials(username, password string) bool {
// 	storedPassword, ok := users[username]
// 	if !ok || storedPassword != password {
// 		return false
// 	}
// 	return true
// }

// func main() {
// 	r := gin.Default()

// 	// Endpoint untuk register dan login
// 	r.POST("/register", RegisterHandler)
// 	r.POST("/login", LoginHandler)

// 	// Endpoint yang memerlukan autentikasi
// 	r.GET("/secure", MiddlewareAuth(), func(c *gin.Context) {
// 		c.String(http.StatusOK, "Secure endpoint")
// 	})

// 	r.Run("localhost:8080")
// }
