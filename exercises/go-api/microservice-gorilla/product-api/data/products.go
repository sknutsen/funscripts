package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	protos "../../grpc/currency/protos/currency"
	"github.com/go-playground/validator"
	"github.com/hashicorp/go-hclog"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// The id for this product
	//
	// required: true
	// min: 1
	ID int `json:"id"`
	// The name for this product
	//
	// required: true
	Name string `json:"name" validate:"required"`
	// The description for this product
	//
	// required: false
	Description string `json:"description"`
	// The price for this product
	//
	// required: true
	// min: 0.01
	Price float64 `json:"price" validate:"gt=0"`
	// The sku for this product
	//
	// required: false
	SKU       string `json:"sku" validate:"required,sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

// Products is a collection of product
type Products []*Product

// ProductsDB ...
type ProductsDB struct {
	currency protos.CurrencyClient
	log      hclog.Logger
}

// NewProductsDB ...
func NewProductsDB(c protos.CurrencyClient, l hclog.Logger) *ProductsDB {
	return &ProductsDB{c, l}
}

// GetProducts returns the productlist
func (p *ProductsDB) GetProducts(currency string) (Products, error) {
	if currency == "" {
		return productList, nil
	}

	rate, err := p.getRate(currency)
	if err != nil {
		p.log.Error("Error getting new rate", "currency", currency, "error", err)
		return nil, err
	}

	pr := Products{}
	for _, p := range productList {
		np := *p
		np.Price = np.Price * rate
		pr = append(pr, &np)
	}

	return nil, nil
}

// GetProduct returns the product with the matching id
func (p *ProductsDB) GetProduct(id int, currency string) (*Product, error) {
	i, err := findProduct(id)
	if err != nil {
		return nil, err
	}

	if currency == "" {
		return productList[i], nil
	}

	rate, err := p.getRate(currency)
	if err != nil {
		p.log.Error("Error getting new rate", "currency", currency, "error", err)
		return nil, err
	}

	np := *productList[i]
	np.Price = productList[i].Price * rate

	return np, nil
}

// FromJSON deserializes json to product
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// Validate struct
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

// ToJSON serializes productlist to json
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// ToJSON serializes productlist to json
func (p *Product) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// AddProduct adds a new product to collection
func AddProduct(p *Product) {
	p.ID = getNextID()
	p.CreatedOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()
	productList = append(productList, p)
}

// UpdateProduct updates the product with the given id
func UpdateProduct(id int, p *Product) error {
	pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

// DeleteProduct deletes the product with the given id
func DeleteProduct(id int) error {
	_, err := findProduct(id)
	if err != nil {
		return err
	}

	return nil
}

// ErrProductNotFound is an error which is returned when a product is not found
var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}

	return -1, ErrProductNotFound
}

func (p *ProductsDB) getRate(dest string) (float64, error) {
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value[dest]),
	}

	resp, err := p.currency.GetRate(context.Background(), rr)

	return resp, err
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "dc332",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
