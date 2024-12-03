package models
import(
	"time"
)
type Order struct {
	OrderUID        string    `gorm:"primaryKey"`
	TrackNumber     string    `gorm:"not null"`
	Entry           string    `gorm:"not null"`
	Locale          string    `gorm:"not null"`
	InternalSignature string  `gorm:"default:null"`
	CustomerID      string    `gorm:"not null"`
	DeliveryService string    `gorm:"not null"`
	ShardKey        string    `gorm:"not null"`
	SmID            int       `gorm:"not null"`
	DateCreated     time.Time `gorm:"not null"`
	OofShard        string    `gorm:"not null"`
}

// Модель для таблицы delivery
type Delivery struct {
	DeliveryID uint   `gorm:"primaryKey"`
	OrderUID   string `gorm:"not null" gorm:"index"`
	Name       string `gorm:"not null"`
	Phone      string `gorm:"not null"`
	Zip        string `gorm:"not null"`
	City       string `gorm:"not null"`
	Address    string `gorm:"not null"`
	Region     string `gorm:"not null"`
	Email      string `gorm:"not null"`
}

// Модель для таблицы payment
type Payment struct {
	PaymentID     uint   `gorm:"primaryKey"`
	OrderUID      string `gorm:"not null" gorm:"index"`
	Transaction   string `gorm:"not null"`
	RequestID     string
	Currency      string `gorm:"not null"`
	Provider      string `gorm:"not null"`
	Amount        int    `gorm:"not null"`
	PaymentDt     time.Time
	Bank          string `gorm:"not null"`
	DeliveryCost  int    `gorm:"not null"`
	GoodsTotal    int    `gorm:"not null"`
	CustomFee     int    `gorm:"not null"`
}

// Модель для таблицы items
type Item struct {
	ItemID      uint   `gorm:"primaryKey"`
	OrderUID    string `gorm:"not null" gorm:"index"`
	ChrtID      int    `gorm:"not null"`
	TrackNumber string `gorm:"not null"`
	Price       int    `gorm:"not null"`
	RID         string `gorm:"not null"`
	Name        string `gorm:"not null"`
	Sale        int    `gorm:"not null"`
	Size        string `gorm:"not null"`
	TotalPrice  int    `gorm:"not null"`
	NMID        int    `gorm:"not null"`
	Brand       string `gorm:"not null"`
	Status      int    `gorm:"not null"`
}