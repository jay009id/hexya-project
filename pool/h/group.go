// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package h

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h/group"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

// ------- MODEL ---------

// GroupModel is a strongly typed model definition that is used
// to extend the Group model or to get a GroupSet through
// its NewSet() function.
//
// To get the unique instance of this type, call Group().
type GroupModel struct {
	*models.Model
}

// NewSet returns a new GroupSet instance in the given Environment
func (md GroupModel) NewSet(env models.Environment) m.GroupSet {
	return group.GroupSet{
		RecordCollection: env.Pool("Group"),
	}
}

// Create creates a new Group record and returns the newly created
// GroupSet instance.
func (md GroupModel) Create(env models.Environment, data m.GroupData) m.GroupSet {
	return group.GroupSet{
		RecordCollection: md.Model.Create(env, data),
	}
}

// Search searches the database and returns a new GroupSet instance
// with the records found.
func (md GroupModel) Search(env models.Environment, cond q.GroupCondition) m.GroupSet {
	return group.GroupSet{
		RecordCollection: md.Model.Search(env, cond),
	}
}

// Browse returns a new RecordSet with the records with the given ids.
// Note that this function is just a shorcut for Search on a list of ids.
func (md GroupModel) Browse(env models.Environment, ids []int64) m.GroupSet {
	return group.GroupSet{
		RecordCollection: md.Model.Browse(env, ids),
	}
}

// BrowseOne returns a new RecordSet with the record with the given id.
// Note that this function is just a shorcut for Search on the given id.
func (md GroupModel) BrowseOne(env models.Environment, id int64) m.GroupSet {
	return group.GroupSet{
		RecordCollection: md.Model.BrowseOne(env, id),
	}
}

// NewData returns a pointer to a new empty GroupData instance.
//
// Optional field maps if given will be used to populate the data.
func (md GroupModel) NewData(fm ...models.FieldMap) m.GroupData {
	return &group.GroupData{
		ModelData: models.NewModelData(Group(), fm...),
	}
}

// Fields returns the Field Collection of the Group Model
func (md GroupModel) Fields() group.FieldsCollection {
	return group.FieldsCollection{
		FieldsCollection: md.Model.Fields(),
	}
}

// Methods returns the Method Collection of the Group Model
func (md GroupModel) Methods() group.MethodsCollection {
	return group.MethodsCollection{
		MethodsCollection: md.Model.Methods(),
	}
}

// Underlying returns the underlying models.Model instance
func (md GroupModel) Underlying() *models.Model {
	return md.Model
}

var _ models.Modeler = GroupModel{}

// Coalesce takes a list of GroupSet and return the first non-empty one
// if every record set is empty, it will return the last given
func (md GroupModel) Coalesce(lst ...m.GroupSet) m.GroupSet {
	var last m.GroupSet
	for _, elem := range lst {
		if elem.Collection().IsNotEmpty() {
			return elem
		}
		last = elem
	}
	return last
}

// Group returns the unique instance of the GroupModel type
// which is used to extend the Group model or to get a GroupSet through
// its NewSet() function.
func Group() GroupModel {
	return GroupModel{
		Model: models.Registry.MustGet("Group"),
	}
}
