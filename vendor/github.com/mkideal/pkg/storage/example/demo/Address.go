// NOTE: AUTO-GENERATED by midc, DON'T edit!!

package demo

import (
	"fmt"

	"github.com/mkideal/pkg/storage"
	"github.com/mkideal/pkg/typeconv"
	"gopkg.in/redis.v5"
)

var (
	_ = fmt.Printf
	_ = storage.Unused
	_ = typeconv.Unused
	_ = redis.Nil
)

// Table

type Address struct {
	Id   int64  `xorm:"pk BIGINT(20) autoincr" json:"id"`
	Addr string `xorm:"VARCHAR(1024) not null" json:"addr"`
}

func NewAddress() *Address {
	return &Address{}
}

func (Address) Meta() AddressMeta            { return addressMetaVar }
func (Address) TableMeta() storage.TableMeta { return addressMetaVar }
func (x Address) Key() interface{}           { return x.Id }
func (x *Address) SetKey(value string) error {
	return typeconv.String2Int64(&x.Id, value)
}

func (x Address) GetField(field string) (interface{}, bool) {
	switch field {
	case addressMetaVar.F_addr:
		return x.Addr, true
	}
	return nil, false
}

func (x *Address) SetField(field, value string) error {
	switch field {
	case addressMetaVar.F_addr:
		x.Addr = value
	}
	return nil
}

// Meta
type AddressMeta struct {
	F_addr string
}

func (AddressMeta) Name() string     { return "address" }
func (AddressMeta) Key() string      { return "id" }
func (AddressMeta) Fields() []string { return _address_fields }

var addressMetaVar = AddressMeta{
	F_addr: "addr",
}

var _address_fields = []string{
	addressMetaVar.F_addr,
}

// Slice
type AddressSlice []Address

func NewAddressSlice(cap int) *AddressSlice {
	s := AddressSlice(make([]Address, 0, cap))
	return &s
}

func (s AddressSlice) TableMeta() storage.TableMeta { return addressMetaVar }
func (s AddressSlice) Len() int                     { return len(s) }
func (s *AddressSlice) Slice() []Address            { return []Address(*s) }

func (s *AddressSlice) New(table string, index int, key string) (storage.Table, error) {
	for len(*s) <= index {
		*s = append(*s, Address{})
	}
	x := &((*s)[index])
	err := x.SetKey(key)
	return x, err
}

// View
type AddressView struct {
	Address
}

type AddressViewSlice []AddressView

func NewAddressViewSlice(cap int) *AddressViewSlice {
	s := AddressViewSlice(make([]AddressView, 0, cap))
	return &s
}

func (s AddressViewSlice) TableMeta() storage.TableMeta { return addressMetaVar }
func (s AddressViewSlice) Len() int                     { return len(s) }
func (s *AddressViewSlice) Slice() []AddressView        { return []AddressView(*s) }

func (s *AddressViewSlice) New(table string, index int, key string) (storage.Table, error) {
	if table == "address" {
		for len(*s) <= index {
			x := Address{}
			*s = append(*s, AddressView{Address: x})
		}
		x := &((*s)[index].Address)
		err := x.SetKey(key)
		return x, err
	}
	v := &((*s)[index])
	for t, x := range v.tables() {
		if t == table {
			err := x.SetKey(key)
			return x, err
		}
	}
	return nil, storage.ErrTableNotFoundInView
}

var (
	AddressViewVar  = AddressView{}
	addressViewRefs = map[string]storage.View{}
)

func (AddressView) TableMeta() storage.TableMeta  { return addressMetaVar }
func (AddressView) Fields() storage.FieldList     { return storage.FieldSlice(addressMetaVar.Fields()) }
func (AddressView) Refs() map[string]storage.View { return addressViewRefs }
func (view *AddressView) tables() map[string]storage.Table {
	m := make(map[string]storage.Table)
	m["address"] = &view.Address
	return m
}
