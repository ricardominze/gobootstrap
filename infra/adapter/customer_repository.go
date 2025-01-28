package adapter

import (
	"context"
	"database/sql"

	customer_entity "github.com/ricardominze/gobootstrap/core/domain/customer/entity"
	"github.com/ricardominze/gobootstrap/core/valueobject"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (re *CustomerRepository) Get(ctx context.Context, id int) (*customer_entity.Customer, error) {

	_, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	customer := &customer_entity.Customer{Address: &valueobject.Address{}}

	result := re.db.QueryRowContext(ctx, `SELECT id, name, city, street, zipcode FROM customer WHERE id = $1`, id)
	result.Scan(&customer.Id, &customer.Name, &customer.Address.City, &customer.Address.Street, &customer.Address.Zipcode)

	if result.Err() != nil {
		return nil, result.Err()
	}

	return customer, nil
}

func (re *CustomerRepository) Save(ctx context.Context, customer *customer_entity.Customer) (*customer_entity.Customer, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	var err error
	var newCustomer *customer_entity.Customer

	if customer.Id == 0 {
		newCustomer, err = re.create(ctx, customer)
	} else {
		newCustomer, err = re.update(ctx, customer)
	}

	if err != nil {
		return nil, err
	}

	return newCustomer, nil
}

func (re *CustomerRepository) create(ctx context.Context, customer *customer_entity.Customer) (*customer_entity.Customer, error) {

	_, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	result := re.db.QueryRowContext(ctx, `INSERT INTO customer (name, city, street, zipcode) VALUES ($1, $2, $3, $4) RETURNING id`, customer.Name, customer.Address.City, customer.Address.Street, customer.Address.Zipcode)
	result.Scan(&customer.Id)

	return customer, nil
}

func (re *CustomerRepository) update(ctx context.Context, customer *customer_entity.Customer) (*customer_entity.Customer, error) {

	_, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	_, err := re.db.ExecContext(ctx, `UPDATE customer SET name = $1, city = $2, street = $3, zipcode = $4 WHERE id = $5`, customer.Name, customer.Address.City, customer.Address.Street, customer.Address.Zipcode, customer.Id)

	if err != nil {
		return nil, err
	}

	return customer, nil
}
