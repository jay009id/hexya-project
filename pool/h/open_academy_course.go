// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package h

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h/open_academy_course"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

// ------- MODEL ---------

// OpenAcademyCourseModel is a strongly typed model definition that is used
// to extend the OpenAcademyCourse model or to get a OpenAcademyCourseSet through
// its NewSet() function.
//
// To get the unique instance of this type, call OpenAcademyCourse().
type OpenAcademyCourseModel struct {
	*models.Model
}

// NewSet returns a new OpenAcademyCourseSet instance in the given Environment
func (md OpenAcademyCourseModel) NewSet(env models.Environment) m.OpenAcademyCourseSet {
	return open_academy_course.OpenAcademyCourseSet{
		RecordCollection: env.Pool("OpenAcademyCourse"),
	}
}

// Create creates a new OpenAcademyCourse record and returns the newly created
// OpenAcademyCourseSet instance.
func (md OpenAcademyCourseModel) Create(env models.Environment, data m.OpenAcademyCourseData) m.OpenAcademyCourseSet {
	return open_academy_course.OpenAcademyCourseSet{
		RecordCollection: md.Model.Create(env, data),
	}
}

// Search searches the database and returns a new OpenAcademyCourseSet instance
// with the records found.
func (md OpenAcademyCourseModel) Search(env models.Environment, cond q.OpenAcademyCourseCondition) m.OpenAcademyCourseSet {
	return open_academy_course.OpenAcademyCourseSet{
		RecordCollection: md.Model.Search(env, cond),
	}
}

// Browse returns a new RecordSet with the records with the given ids.
// Note that this function is just a shorcut for Search on a list of ids.
func (md OpenAcademyCourseModel) Browse(env models.Environment, ids []int64) m.OpenAcademyCourseSet {
	return open_academy_course.OpenAcademyCourseSet{
		RecordCollection: md.Model.Browse(env, ids),
	}
}

// BrowseOne returns a new RecordSet with the record with the given id.
// Note that this function is just a shorcut for Search on the given id.
func (md OpenAcademyCourseModel) BrowseOne(env models.Environment, id int64) m.OpenAcademyCourseSet {
	return open_academy_course.OpenAcademyCourseSet{
		RecordCollection: md.Model.BrowseOne(env, id),
	}
}

// NewData returns a pointer to a new empty OpenAcademyCourseData instance.
//
// Optional field maps if given will be used to populate the data.
func (md OpenAcademyCourseModel) NewData(fm ...models.FieldMap) m.OpenAcademyCourseData {
	return &open_academy_course.OpenAcademyCourseData{
		ModelData: models.NewModelData(OpenAcademyCourse(), fm...),
	}
}

// Fields returns the Field Collection of the OpenAcademyCourse Model
func (md OpenAcademyCourseModel) Fields() open_academy_course.FieldsCollection {
	return open_academy_course.FieldsCollection{
		FieldsCollection: md.Model.Fields(),
	}
}

// Methods returns the Method Collection of the OpenAcademyCourse Model
func (md OpenAcademyCourseModel) Methods() open_academy_course.MethodsCollection {
	return open_academy_course.MethodsCollection{
		MethodsCollection: md.Model.Methods(),
	}
}

// Underlying returns the underlying models.Model instance
func (md OpenAcademyCourseModel) Underlying() *models.Model {
	return md.Model
}

var _ models.Modeler = OpenAcademyCourseModel{}

// Coalesce takes a list of OpenAcademyCourseSet and return the first non-empty one
// if every record set is empty, it will return the last given
func (md OpenAcademyCourseModel) Coalesce(lst ...m.OpenAcademyCourseSet) m.OpenAcademyCourseSet {
	var last m.OpenAcademyCourseSet
	for _, elem := range lst {
		if elem.Collection().IsNotEmpty() {
			return elem
		}
		last = elem
	}
	return last
}

// OpenAcademyCourse returns the unique instance of the OpenAcademyCourseModel type
// which is used to extend the OpenAcademyCourse model or to get a OpenAcademyCourseSet through
// its NewSet() function.
func OpenAcademyCourse() OpenAcademyCourseModel {
	return OpenAcademyCourseModel{
		Model: models.Registry.MustGet("OpenAcademyCourse"),
	}
}
