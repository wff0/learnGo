package golangProgrammingMode

import (
	"encoding/binary"
	"io"
)

type Point struct {
	Longitude     interface{}
	Latitude      interface{}
	Distance      interface{}
	ElevationGain interface{}
	ElevationLoss interface{}
}

// 正常处理方式
func handle1(r io.Reader) (*Point, error) {
	var p Point
	if err := binary.Read(r, binary.BigEndian, &p.Longitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Latitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Distance); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationGain); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationLoss); err != nil {
		return nil, err
	}
	return nil, nil
}

// 函数式编程
func handle2(r io.Reader) (*Point, error) {
	var p Point
	var err error
	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}

	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)
	if err != nil {
		return &p, err
	}
	return &p, nil
}

type reader struct {
	r   io.Reader
	err error
}

func (r *reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

// fluent接口
func handle3(input io.Reader) (*Point, error) {
	var p Point
	r := reader{r: input}

	r.read(&p.Longitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)
	if r.err != nil {
		return nil, r.err
	}
	return &p, nil
}
