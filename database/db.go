package database

import (
	"jabar-nearby-places/models"
	"jabar-nearby-places/utils"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	var db *gorm.DB
	var err error

	dsn := utils.Getenv("DATABASE_URL", "root:@tcp(127.0.0.1:3306)/db_nearby?charset=utf8mb4&parseTime=True&loc=Local")
	if utils.Getenv("ENV", "") == "production" {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(
		&models.Place{},
		&models.Category{},
	)

	if db.Migrator().HasTable(&models.Category{}) {
		cs := []models.Category{
			{Name: models.CATEGORY_KANTOR_PEM_KABKOTA},
			{Name: models.CATEGORY_RUMAH_SAKIT},
			{Name: models.CATEGORY_SMA},
			{Name: models.CATEGORY_PUSKESMAS},
			{Name: models.CATEGORY_SMP},
			{Name: models.CATEGORY_KANTOR_PEM_KECAMATAN},
			{Name: models.CATEGORY_SD},
			{Name: models.CATEGORY_TEPAT_IBADAH},
			{Name: models.CATEGORY_KANTOR_PEM_KELURAHANDESA},
		}
		db.Create(&cs)
	}

	return db
}
