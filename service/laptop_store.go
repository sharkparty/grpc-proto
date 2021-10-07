package service

import (
	"errors"
	"fmt"
	"grpc-proto/pb"
	"sync"

	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("Record already exists")

type LaptopStore interface {
	Save (laptop *pb.Laptop) error
	Find (id string) (*pb.Laptop, error)
}

type InMemoryStore struct {
	mutex sync.RWMutex
	data map[string]*pb.Laptop
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryStore) Save (laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("Cannot copy laptop data: %w", err)
	}

	store.data[other.Id] = other

	return nil
}

func (store *InMemoryStore) Find (id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("Cannot copy laptop data: %w", err)
	}

	return other, nil
}