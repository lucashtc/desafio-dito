package coletora
import (
	"github.com/jinzhu/gorm"
)

type Colector struct {
	gorm.Model
	Event     string `json:"event"`
	Timestamp string `json:"timestamp"`
}