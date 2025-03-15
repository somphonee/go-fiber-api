package migrations
import (
	"gorm.io/gorm"
	"github.com/somphonee/go-fiber-api/internal/models"
)
// Migrate performs database migrations
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Product{},
	)
}