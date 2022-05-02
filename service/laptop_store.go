package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/shuntagami/pcbook-go/pb"
)

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	// Save saves the laptop to the store
	Save(laptop *pb.Laptop) error
}

// InMemoryLaptopStore stores laptop in memory
// if we want to save laptops to the database, we can always implement another DBLaptopStore to do so.
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	// key is the laptop ID, and the value is the laptop object
	data map[string]*pb.Laptop
}

// NewInMemoryLaptopStore returns a new InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

// ErrAlreadyExists is returned when a record with the same ID already exists in the store
var ErrAlreadyExists = errors.New("record already exists")

// Save saves the laptop to the store
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	other, err := deepCopy(laptop)
	if err != nil {
		return err
	}

	store.data[other.Id] = other
	return nil
}

func deepCopy(laptop *pb.Laptop) (*pb.Laptop, error) {
	other := &pb.Laptop{}

	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil
}
