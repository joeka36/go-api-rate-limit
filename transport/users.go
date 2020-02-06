package transport

// User represents the user calling our API
type User struct {
	Key         string
	AccountType string
}

// APIUsers is our non persistent user table 
var APIUsers = map[string]User{
	"test@gmail.com": User{
		Key: "A1B2C3",
		AccountType: "Basic",
	},
	"test@yahoo.com": User{
		Key: "D4E5F6",
		AccountType: "Premium",
	},
	"test@hotmail.com": User{
		Key: "123456",
		AccountType: "Basic",
	},
}
