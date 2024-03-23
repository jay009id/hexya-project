// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package h

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h/currency"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

// ------- MODEL ---------

// CurrencyModel is a strongly typed model definition that is used
// to extend the Currency model or to get a CurrencySet through
// its NewSet() function.
//
// To get the unique instance of this type, call Currency().
type CurrencyModel struct {
	*models.Model
}

// NewSet returns a new CurrencySet instance in the given Environment
func (md CurrencyModel) NewSet(env models.Environment) m.CurrencySet {
	return currency.CurrencySet{
		RecordCollection: env.Pool("Currency"),
	}
}

// Create creates a new Currency record and returns the newly created
// CurrencySet instance.
func (md CurrencyModel) Create(env models.Environment, data m.CurrencyData) m.CurrencySet {
	return currency.CurrencySet{
		RecordCollection: md.Model.Create(env, data),
	}
}

// Search searches the database and returns a new CurrencySet instance
// with the records found.
func (md CurrencyModel) Search(env models.Environment, cond q.CurrencyCondition) m.CurrencySet {
	return currency.CurrencySet{
		RecordCollection: md.Model.Search(env, cond),
	}
}

// Browse returns a new RecordSet with the records with the given ids.
// Note that this function is just a shorcut for Search on a list of ids.
func (md CurrencyModel) Browse(env models.Environment, ids []int64) m.CurrencySet {
	return currency.CurrencySet{
		RecordCollection: md.Model.Browse(env, ids),
	}
}

// BrowseOne returns a new RecordSet with the record with the given id.
// Note that this function is just a shorcut for Search on the given id.
func (md CurrencyModel) BrowseOne(env models.Environment, id int64) m.CurrencySet {
	return currency.CurrencySet{
		RecordCollection: md.Model.BrowseOne(env, id),
	}
}

// NewData returns a pointer to a new empty CurrencyData instance.
//
// Optional field maps if given will be used to populate the data.
func (md CurrencyModel) NewData(fm ...models.FieldMap) m.CurrencyData {
	return &currency.CurrencyData{
		ModelData: models.NewModelData(Currency(), fm...),
	}
}

// Fields returns the Field Collection of the Currency Model
func (md CurrencyModel) Fields() currency.FieldsCollection {
	return currency.FieldsCollection{
		FieldsCollection: md.Model.Fields(),
	}
}

// Methods returns the Method Collection of the Currency Model
func (md CurrencyModel) Methods() currency.MethodsCollection {
	return currency.MethodsCollection{
		MethodsCollection: md.Model.Methods(),
	}
}

// Underlying returns the underlying models.Model instance
func (md CurrencyModel) Underlying() *models.Model {
	return md.Model
}

var _ models.Modeler = CurrencyModel{}

// Coalesce takes a list of CurrencySet and return the first non-empty one
// if every record set is empty, it will return the last given
func (md CurrencyModel) Coalesce(lst ...m.CurrencySet) m.CurrencySet {
	var last m.CurrencySet
	for _, elem := range lst {
		if elem.Collection().IsNotEmpty() {
			return elem
		}
		last = elem
	}
	return last
}

// Currency returns the unique instance of the CurrencyModel type
// which is used to extend the Currency model or to get a CurrencySet through
// its NewSet() function.
func Currency() CurrencyModel {
	return CurrencyModel{
		Model: models.Registry.MustGet("Currency"),
	}
}
