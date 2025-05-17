package database

import(
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
	"fmt"
)

func ConnectDB() *gorm.DB {
	
	//Dsn For Connect
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",			
	//Get From my env 
	 os.Getenv("DB_HOST"),
	 os.Getenv("DB_PORT"),
	 os.Getenv("DB_USER"),
	 os.Getenv("DB_PASSWORD"),
	 os.Getenv("DB_NAME"),
)

	//Connect Database With GORM
	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err !=nil{
		panic("Failed to Connect Database !!")
	}

	return db

		// New logger for detailed SQL logging
  /*newLogger := logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
    logger.Config{
      SlowThreshold: time.Second, // Slow SQL threshold
      LogLevel:      logger.Info, // Log level
      Colorful:      true,        // Enable color
    },
  )*/
	
}