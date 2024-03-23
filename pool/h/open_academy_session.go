// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package h

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h/open_academy_session"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

// ------- MODEL ---------

// OpenAcademySessionModel is a strongly typed model definition that is used
// to extend the OpenAcademySession model or to get a OpenAcademySessionSet through
// its NewSet() function.
//
// To get the unique instance of this type, call OpenAcademySession().
type OpenAcademySessionModel struct {
	*models.Model
}

// NewSet returns a new OpenAcademySessionSet instance in the given Environment
func (md OpenAcademySessionModel) NewSet(env models.Environment) m.OpenAcademySessionSet {
	return open_academy_session.OpenAcademySessionSet{
		RecordCollection: env.Pool("OpenAcademySession"),
	}
}

// Create creates a new OpenAcademySession record and returns the newly created
// OpenAcademySessionSet instance.
func (md OpenAcademySessionModel) Create(env models.Environment, data m.OpenAcademySessionData) m.OpenAcademySessionSet {
	return open_academy_session.OpenAcademySessionSet{
		RecordCollection: md.Model.Create(env, data),
	}
}

// Search searches the database and returns a new OpenAcademySessionSet instance
// with the records found.
func (md OpenAcademySessionModel) Search(env models.Environment, cond q.OpenAcademySessionCondition) m.OpenAcademySessionSet {
	return open_academy_session.OpenAcademySessionSet{
		RecordCollection: md.Model.Search(env, cond),
	}
}

// Browse returns a new RecordSet with the records with the given ids.
// Note that this function is just a shorcut for Search on a list of ids.
func (md OpenAcademySessionModel) Browse(env models.Environment, ids []int64) m.OpenAcademySessionSet {
	return open_academy_session.OpenAcademySessionSet{
		RecordCollection: md.Model.Browse(env, ids),
	}
}

// BrowseOne returns a new RecordSet with the record with the given id.
// Note that this function is just a shorcut for Search on the given id.
func (md OpenAcademySessionModel) BrowseOne(env models.Environment, id int64) m.OpenAcademySessionSet {
	return open_academy_session.OpenAcademySessionSet{
		RecordCollection: md.Model.BrowseOne(env, id),
	}
}

// NewData returns a pointer to a new empty OpenAcademySessionData instance.
//
// Optional field maps if given will be used to populate the data.
func (md OpenAcademySessionModel) NewData(fm ...models.FieldMap) m.OpenAcademySessionData {
	return &open_academy_session.OpenAcademySessionData{
		ModelData: models.NewModelData(OpenAcademySession(), fm...),
	}
}

// Fields returns the Field Collection of the OpenAcademySession Model
func (md OpenAcademySessionModel) Fields() open_academy_session.FieldsCollection {
	return open_academy_session.FieldsCollection{
		FieldsCollection: md.Model.Fields(),
	}
}

// Methods returns the Method Collection of the OpenAcademySession Model
func (md OpenAcademySessionModel) Methods() open_academy_session.MethodsCollection {
	return open_academy_session.MethodsCollection{
		MethodsCollection: md.Model.Methods(),
	}
}

// Underlying returns the underlying models.Model instance
func (md OpenAcademySessionModel) Underlying() *models.Model {
	return md.Model
}

var _ models.Modeler = OpenAcademySessionModel{}

// Coalesce takes a list of OpenAcademySessionSet and return the first non-empty one
// if every record set is empty, it will return the last given
func (md OpenAcademySessionModel) Coalesce(lst ...m.OpenAcademySessionSet) m.OpenAcademySessionSet {
	var last m.OpenAcademySessionSet
	for _, elem := range lst {
		if elem.Collection().IsNotEmpty() {
			return elem
		}
		last = elem
	}
	return last
}

// OpenAcademySession returns the unique instance of the OpenAcademySessionModel type
// which is used to extend the OpenAcademySession model or to get a OpenAcademySessionSet through
// its NewSet() function.
func OpenAcademySession() OpenAcademySessionModel {
	return OpenAcademySessionModel{
		Model: models.Registry.MustGet("OpenAcademySession"),
	}
}