package migration

import (
	"Project_PBO/database"
	"Project_PBO/models/entity"
	"fmt"
)

func StartMigration() {

	err := database.DB.AutoMigrate(&entity.Siswa{}, &entity.Berita{}, &entity.Prestasi{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Database Migration Completed!")
}
