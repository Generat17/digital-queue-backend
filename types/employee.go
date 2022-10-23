package types

// Employee структура, которая соответствует таблице Employee в БД
type Employee struct {
	EmployeeId   int    `json:"employee_id" db:"employee_id"`
	Login        string `json:"login" db:"login"`
	Password     string `json:"password" db:"password"`
	FirstName    string `json:"first_name" db:"first_name"`
	SecondName   string `json:"second_name" db:"second_name"`
	Position     int    `json:"position" db:"position"`
	SessionState bool   `json:"session_state" db:"session_state"`
	Status       int    `json:"status" db:"status"`
}

type GetNewClientResponse struct {
	NumberTicket  int    `json:"number_ticket"`
	ServiceTicket string `json:"service_ticket"`
}
