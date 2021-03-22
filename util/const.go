package util

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"strconv"
)

type NumberInt int

func (n *NumberInt) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.BigEndian, n); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (n *NumberInt) UnmarshalJSON(data []byte) (err error) {
	wrapper := struct {
		NumberInt string `json:"$numberInt"`
	}{}
	if err = json.Unmarshal(data, &wrapper); err != nil {
		return
	}
	v, err := strconv.ParseInt(wrapper.NumberInt, 10, 64)
	if err != nil {
		return
	}
	*n = NumberInt(v)
	return
}

type NumberDouble float64

func (n *NumberDouble) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.BigEndian, n); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (n *NumberDouble) UnmarshalJSON(data []byte) (err error) {
	wrapper := struct {
		NumberDouble string `json:"$numberDouble"`
	}{}
	if err = json.Unmarshal(data, &wrapper); err != nil {
		return
	}
	v, err := strconv.ParseFloat(wrapper.NumberDouble, 64)
	if err != nil {
		return
	}
	*n = NumberDouble(v)
	return
}

type NumberLong int64

func (n *NumberLong) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.BigEndian, n); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (n *NumberLong) UnmarshalJSON(data []byte) (err error) {
	wrapper := struct {
		NumberLong string `json:"$numberLong"`
	}{}
	if err = json.Unmarshal(data, &wrapper); err != nil {
		return
	}
	v, err := strconv.ParseInt(wrapper.NumberLong, 10, 64)
	if err != nil {
		return
	}
	*n = NumberLong(v)
	return
}
