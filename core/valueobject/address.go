package valueobject

type Address struct {
	City    string
	Street  string
	Zipcode string
}

func NewAddress(city string, street string, zipcode string) *Address {

	return &Address{
		City:    city,
		Street:  street,
		Zipcode: zipcode,
	}
}

func (o *Address) IsValid() error {

	return nil
}

func (o *Address) ChangeCity(city string) {

	o.City = city
}

func (o *Address) ChangeStreet(street string) {

	o.Street = street
}

func (o *Address) ChangeZipcode(zipcode string) {

	o.Zipcode = zipcode
}
