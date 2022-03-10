package gorm

//contoh 1
//type User struct {
//	gorm.Model
//	Name       string     `gorm:"uniqueIndex"`
//	CreditCard CreditCard `gorm:"foreignKey:UserName;references:Name"`
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number   string
//	UserName string
//}

//contoh 2
// User has one CreditCard, CreditCardID is the foreign key
//type User struct {
//	gorm.Model
//	CreditCard CreditCard
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number string
//	UserID uint
//}
