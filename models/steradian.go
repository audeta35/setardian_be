package models

type Article struct {
	ID    string `json:id`
	Title string `json:title`
	Body  string `json:body`
}

// User Admin Request & Response

type UserAdminLogin struct {
	Email    string `json:email`
	Password string `json:password`
}

type UserAdminData struct {
	UserID   string `json:userId`
	Email    string `json:email`
	Password string `json:password`
}

type AdminRegister struct {
	ID       string `json:id`
	Email    string `json:email`
	Password string `json:password`
}

type UserRegister struct {
	ID          string `json:userId`
	Email       string `json:email`
	Password    string `json:password`
	PhoneNumber string `json:phoneNumber`
	City        string `json:city`
	Zip         string `json:zip`
	Message     string `json:message`
	UserName    string `json:username`
	Address     string `json:address`
}

// Order Request & Response

type Orders struct {
	ID          string `json:orderId`
	PickUpLoc   string `json:pickUpLoc`
	DropOffLoc  string `json:dropOffLoc`
	PickUpDate  string `json:pickUpDate`
	DropOffDate string `json:dropOffDate`
	PickUpTime  string `json:pickUpTime`
	CarId       string `json:carId`
	UserId      string `json:userId`
	AdminId     string `json:adminId`
}

// Cars Request & Response

type Cars struct {
	ID        string `json:carId`
	Name      string `json:name`
	CarId     string `json:carId`
	CarType   string `json:carType`
	Rating    string `json:rating`
	Fuel      string `json:fuel`
	Image     string `json:image`
	HourRate  string `json:hourRate`
	DayRate   string `json:dayRate`
	MonthRate string `json:monthRate`
}
