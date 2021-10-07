package sample

import (
	"grpc-proto/pb"
	"github.com/golang/protobuf/ptypes"
)

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBoolean(),
	}

	return keyboard
}

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberOfCores := randomInt(2, 8)
	numberOfThreads := randomInt(numberOfCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumerCores:    uint32(numberOfCores),
		NumberThreads: uint32(numberOfThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 2.0)
	maxGhz := randomFloat64(minGhz, 3.5)

	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return gpu
}

func NewRam() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}

	return ram
}

func NewSSD() *pb.Storage {
	memory := &pb.Memory{
		Value: uint64(randomFloat64(128, 1024)),
		Unit:  pb.Memory_GIGABYTE,
	}

	storage := &pb.Storage{
		Driver: pb.Storage_SDD,
		Memory: memory,
	}

	return storage
}

func NewHDD() *pb.Storage {
	memory := &pb.Memory{
		Value: uint64(randomFloat64(1, 6)),
		Unit:  pb.Memory_TERABYTE,
	}

	storage := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: memory,
	}

	return storage
}

func NewScreen() *pb.Screen {
	size := randomFloat32(13, 30)

	screen := &pb.Screen{
		SizeInch:   size,
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBoolean(),
	}

	return screen
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()

	laptop := &pb.Laptop{
		Id:       randomUUID(),
		Brand:    brand,
		Name:     randomLaptopName(brand),
		Cpu:      NewCPU(),
		Ram:      NewRam(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewHDD(), NewSSD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightLbs{
			WeightLbs: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1000.00, 5000.00),
		ReleaseYear: uint32(randomInt(2018, 2021)),
		Timestamp:   ptypes.TimestampNow(),
	}

	return laptop
}
