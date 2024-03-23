// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package h

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h/partner"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

// ------- MODEL ---------

// PartnerModel is a strongly typed model definition that is used
// to extend the Partner model or to get a PartnerSet through
// its NewSet() function.
//
// To get the unique instance of this type, call Partner().
type PartnerModel struct {
	*models.Model
}

// NewSet returns a new PartnerSet instance in the given Environment
func (md PartnerModel) NewSet(env models.Environment) m.PartnerSet {
	return partner.PartnerSet{
		RecordCollection: env.Pool("Partner"),
	}
}

// Create creates a new Partner record and returns the newly created
// PartnerSet instance.
func (md PartnerModel) Create(env models.Environment, data m.PartnerData) m.PartnerSet {
	return partner.PartnerSet{
		RecordCollection: md.Model.Create(env, data),
	}
}

// Search searches the database and returns a new PartnerSet instance
// with the records found.
func (md PartnerModel) Search(env models.Environment, cond q.PartnerCondition) m.PartnerSet {
	return partner.PartnerSet{
		RecordCollection: md.Model.Search(env, cond),
	}
}

// Browse returns a new RecordSet with the records with the given ids.
// Note that this function is just a shorcut for Search on a list of ids.
func (md PartnerModel) Browse(env models.Environment, ids []int64) m.PartnerSet {
	return partner.PartnerSet{
		RecordCollection: md.Model.Browse(env, ids),
	}
}

// BrowseOne returns a new RecordSet with the record with the given id.
// Note that this function is just a shorcut for Search on the given id.
func (md PartnerModel) BrowseOne(env models.Environment, id int64) m.PartnerSet {
	return partner.PartnerSet{
		RecordCollection: md.Model.BrowseOne(env, id),
	}
}

// NewData returns a pointer to a new empty PartnerData instance.
//
// Optional field maps if given will be used to populate the data.
func (md PartnerModel) NewData(fm ...models.FieldMap) m.PartnerData {
	return &partner.PartnerData{
		ModelData: models.NewModelData(Partner(), fm...),
	}
}

// Fields returns the Field Collection of the Partner Model
func (md PartnerModel) Fields() partner.FieldsCollection {
	return partner.FieldsCollection{
		FieldsCollection: md.Model.Fields(),
	}
}

// Methods returns the Method Collection of the Partner Model
func (md PartnerModel) Methods() partner.MethodsCollection {
	return partner.MethodsCollection{
		MethodsCollection: md.Model.Methods(),
	}
}

// Underlying returns the underlying models.Model instance
func (md PartnerModel) Underlying() *models.Model {
	return md.Model
}

var _ models.Modeler = PartnerModel{}

// Coalesce takes a list of PartnerSet and return the first non-empty one
// if every record set is empty, it will return the last given
func (md PartnerModel) Coalesce(lst ...m.PartnerSet) m.PartnerSet {
	var last m.PartnerSet
	for _, elem := range lst {
		if elem.Collection().IsNotEmpty() {
			return elem
		}
		last = elem
	}
	return last
}

// Partner returns the unique instance of the PartnerModel type
// which is used to extend the Partner model or to get a PartnerSet through
// its NewSet() function.
func Partner() PartnerModel {
	return PartnerModel{
		Model: models.Registry.MustGet("Partner"),
	}
}
