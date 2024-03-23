// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package h

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h/lang"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

// ------- MODEL ---------

// LangModel is a strongly typed model definition that is used
// to extend the Lang model or to get a LangSet through
// its NewSet() function.
//
// To get the unique instance of this type, call Lang().
type LangModel struct {
	*models.Model
}

// NewSet returns a new LangSet instance in the given Environment
func (md LangModel) NewSet(env models.Environment) m.LangSet {
	return lang.LangSet{
		RecordCollection: env.Pool("Lang"),
	}
}

// Create creates a new Lang record and returns the newly created
// LangSet instance.
func (md LangModel) Create(env models.Environment, data m.LangData) m.LangSet {
	return lang.LangSet{
		RecordCollection: md.Model.Create(env, data),
	}
}

// Search searches the database and returns a new LangSet instance
// with the records found.
func (md LangModel) Search(env models.Environment, cond q.LangCondition) m.LangSet {
	return lang.LangSet{
		RecordCollection: md.Model.Search(env, cond),
	}
}

// Browse returns a new RecordSet with the records with the given ids.
// Note that this function is just a shorcut for Search on a list of ids.
func (md LangModel) Browse(env models.Environment, ids []int64) m.LangSet {
	return lang.LangSet{
		RecordCollection: md.Model.Browse(env, ids),
	}
}

// BrowseOne returns a new RecordSet with the record with the given id.
// Note that this function is just a shorcut for Search on the given id.
func (md LangModel) BrowseOne(env models.Environment, id int64) m.LangSet {
	return lang.LangSet{
		RecordCollection: md.Model.BrowseOne(env, id),
	}
}

// NewData returns a pointer to a new empty LangData instance.
//
// Optional field maps if given will be used to populate the data.
func (md LangModel) NewData(fm ...models.FieldMap) m.LangData {
	return &lang.LangData{
		ModelData: models.NewModelData(Lang(), fm...),
	}
}

// Fields returns the Field Collection of the Lang Model
func (md LangModel) Fields() lang.FieldsCollection {
	return lang.FieldsCollection{
		FieldsCollection: md.Model.Fields(),
	}
}

// Methods returns the Method Collection of the Lang Model
func (md LangModel) Methods() lang.MethodsCollection {
	return lang.MethodsCollection{
		MethodsCollection: md.Model.Methods(),
	}
}

// Underlying returns the underlying models.Model instance
func (md LangModel) Underlying() *models.Model {
	return md.Model
}

var _ models.Modeler = LangModel{}

// Coalesce takes a list of LangSet and return the first non-empty one
// if every record set is empty, it will return the last given
func (md LangModel) Coalesce(lst ...m.LangSet) m.LangSet {
	var last m.LangSet
	for _, elem := range lst {
		if elem.Collection().IsNotEmpty() {
			return elem
		}
		last = elem
	}
	return last
}

// Lang returns the unique instance of the LangModel type
// which is used to extend the Lang model or to get a LangSet through
// its NewSet() function.
func Lang() LangModel {
	return LangModel{
		Model: models.Registry.MustGet("Lang"),
	}
}