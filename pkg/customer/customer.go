package customer

type Customer struct { // Customer es la estructura de un cliente
	name    string
	address string
	phone   string
}

// función New retorna un nuevo cliente
func New(name, address, phone string) Customer {
	return Customer{name, address, phone}
}
