package main

import (
	"database/sql"
	"final-project/controllers"
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

// var jwtKey = []byte("SECRET")

// type Claims struct {
// 	Username string `json:"username"`
// 	jwt.RegisteredClaims
// }

// // RegisterHandler menangani pendaftaran pengguna baru
// func RegisterHandler(c *gin.Context) {
// 	var user structs.User
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

	router.GET("/task-status", controllers.GetAllTaskStatus)
	router.POST("/task-status", controllers.InsertTaskStatus)
	router.PUT("/task-status/:id", controllers.UpdateTaskStatus)
	router.DELETE("/task-status/:id", controllers.DeleteTaskStatus)

	router.GET("/task", controllers.GetAllTask)
	router.POST("/task", controllers.InsertTask)
	router.PUT("/task/:id", controllers.UpdateTask)
	router.DELETE("/task/:id", controllers.DeleteTask)

	router.GET("/project", controllers.GetAllProject)
	router.POST("/project", controllers.InsertProject)
	router.PUT("/project/:id", controllers.UpdateProject)
	router.DELETE("/project/:id", controllers.DeleteProject)

	router.Run("localhost:8080")
	// router.Run(":" + os.Getenv("PORT"))
}

// // User struct represents a user in the system
// type User struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// // Claims struct digunakan untuk menentukan payload token JWT

// // Database sementara untuk menyimpan pengguna yang terdaftar
// var users = map[string]string{
// 	"user1": "password1",
// 	"user2": "password2",
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
