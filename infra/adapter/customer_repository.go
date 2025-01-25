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

	result := re.db.QueryRowContext(ctx, "SELECT id, name, city, street, zipcode FROM customer WHERE id = ?", id)
	result.Scan(&customer.Id, &customer.Name, &customer.Address.City, &customer.Address.Street, &customer.Address.Zipcode)

	if result.Err() != nil {
		return nil, result.Err()
	}

	return customer, nil
}

func (re *CustomerRepository) Save(ctx context.Context, customer *customer_entity.Customer) (*customer_entity.Customer, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	customerOri := &customer_entity.Customer{Address: &valueobject.Address{}}

	result := re.db.QueryRowContext(ctx, "SELECT id, name, city, street, zipcode FROM customer WHERE id = ?", customer.Id)
	result.Scan(&customerOri.Id, &customerOri.Name, &customerOri.Address.City, &customerOri.Address.Street, &customerOri.Address.Zipcode)

	if customerOri.Id == 0 {

		_, err := re.create(ctx, customer)

		if err != nil {
			return nil, err
		}

	} else {

		_, err := re.update(ctx, customer)

		if err != nil {
			return nil, err
		}
	}

	return customer, nil
}

func (re *CustomerRepository) create(ctx context.Context, customer *customer_entity.Customer) (*customer_entity.Customer, error) {

	_, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	result, err := re.db.ExecContext(ctx, "INSERT INTO customer (name, city, street, zipcode) VALUES (?, ?, ?, ?)", customer.Name, customer.Address.City, customer.Address.Street, customer.Address.Zipcode)

	if err != nil {
		return nil, err
	}

	lastInsertId, _ := result.LastInsertId()
	customer.Id = int(lastInsertId)

	return customer, nil
}

func (re *CustomerRepository) update(ctx context.Context, customer *customer_entity.Customer) (*customer_entity.Customer, error) {

	_, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	_, err := re.db.ExecContext(ctx, "UPDATE customer SET name = ?, city = ?, street = ?, zipcode = ? WHERE id = ?", customer.Name, customer.Address.City, customer.Address.Street, customer.Address.Zipcode, customer.Id)

	if err != nil {
		return nil, err
	}

	return customer, nil
}
