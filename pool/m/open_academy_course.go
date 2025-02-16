// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package m

import (
	"github.com/beevik/etree"
	"github.com/hexya-addons/web/domains"
	"github.com/hexya-addons/web/webtypes"
	"github.com/hexya-erp/hexya/src/actions"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/operator"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/hexya/src/models/types/dates"
	"github.com/hexya-erp/pool/q"
)

// OpenAcademyCourseSet is an autogenerated type to handle OpenAcademyCourse objects.
type OpenAcademyCourseSet interface {
	models.RecordSet
	// OpenAcademyCourseSetHexyaFunc is a dummy function to uniquely match interfaces.
	OpenAcademyCourseSetHexyaFunc()
	// ForceLoad reloads the cache for the given fields and updates the ids of this OpenAcademyCourseSet.
	//
	// If no fields are given, all DB columns of the OpenAcademyCourse model are retrieved.
	//
	// It also returns this OpenAcademyCourseSet.
	ForceLoad(fields ...models.FieldName) OpenAcademyCourseSet
	// CreateDate is a getter for the value of the "CreateDate" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	CreateDate() dates.DateTime
	// SetCreateDate is a setter for the value of the "CreateDate" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetCreateDate panics if the RecordSet is empty.
	SetCreateDate(value dates.DateTime)
	// CreateUID is a getter for the value of the "CreateUID" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	CreateUID() int64
	// SetCreateUID is a setter for the value of the "CreateUID" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetCreateUID panics if the RecordSet is empty.
	SetCreateUID(value int64)
	// Description is a getter for the value of the "Description" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Description() string
	// SetDescription is a setter for the value of the "Description" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetDescription panics if the RecordSet is empty.
	SetDescription(value string)
	// DisplayName is a getter for the value of the "DisplayName" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	DisplayName() string
	// SetDisplayName is a setter for the value of the "DisplayName" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetDisplayName panics if the RecordSet is empty.
	SetDisplayName(value string)
	// HexyaExternalID is a getter for the value of the "HexyaExternalID" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	HexyaExternalID() string
	// SetHexyaExternalID is a setter for the value of the "HexyaExternalID" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetHexyaExternalID panics if the RecordSet is empty.
	SetHexyaExternalID(value string)
	// HexyaVersion is a getter for the value of the "HexyaVersion" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	HexyaVersion() int
	// SetHexyaVersion is a setter for the value of the "HexyaVersion" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetHexyaVersion panics if the RecordSet is empty.
	SetHexyaVersion(value int)
	// ID is a getter for the value of the "ID" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	ID() int64
	// SetID is a setter for the value of the "ID" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetID panics if the RecordSet is empty.
	SetID(value int64)
	// LastUpdate is a getter for the value of the "LastUpdate" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	LastUpdate() dates.DateTime
	// SetLastUpdate is a setter for the value of the "LastUpdate" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetLastUpdate panics if the RecordSet is empty.
	SetLastUpdate(value dates.DateTime)
	// Name is a getter for the value of the "Name" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Name() string
	// SetName is a setter for the value of the "Name" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetName panics if the RecordSet is empty.
	SetName(value string)
	// Responsible is a getter for the value of the "Responsible" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Responsible() UserSet
	// SetResponsible is a setter for the value of the "Responsible" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetResponsible panics if the RecordSet is empty.
	SetResponsible(value UserSet)
	// Sessions is a getter for the value of the "Sessions" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Sessions() OpenAcademySessionSet
	// SetSessions is a setter for the value of the "Sessions" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetSessions panics if the RecordSet is empty.
	SetSessions(value OpenAcademySessionSet)
	// State is a getter for the value of the "State" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	State() string
	// SetState is a setter for the value of the "State" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetState panics if the RecordSet is empty.
	SetState(value string)
	// WriteDate is a getter for the value of the "WriteDate" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	WriteDate() dates.DateTime
	// SetWriteDate is a setter for the value of the "WriteDate" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetWriteDate panics if the RecordSet is empty.
	SetWriteDate(value dates.DateTime)
	// WriteUID is a getter for the value of the "WriteUID" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	WriteUID() int64
	// SetWriteUID is a setter for the value of the "WriteUID" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetWriteUID panics if the RecordSet is empty.
	SetWriteUID(value int64)
	// ActionArchive sets Active=false on a recordset, by calling ToggleActive to take the
	// corresponding actions according to the model
	ActionArchive()
	// ActionUnarchive sets Active=true on a recordset, by calling ToggleActive to take the
	// corresponding actions according to the model
	ActionUnarchive()
	// AddDomainLimitOffset adds the given domain, limit, offset
	// and order to the current RecordSet query.
	AddDomainLimitOffset(domain domains.Domain, limit int, offset int, order string) OpenAcademyCourseSet
	// AddModifiers adds the modifiers attribute nodes to given xml doc.
	AddModifiers(doc *etree.Document, fieldInfos map[string]*models.FieldInfo)
	// AddNameToRelations returns the given RecordData after getting the name of all 2one relation ids
	AddNamesToRelations(data models.RecordData, fInfos map[string]*models.FieldInfo) models.RecordData

	Aggregates(fieldNames ...models.FieldName) []OpenAcademyCourseGroupAggregateRow
	// Browse returns a new RecordSet with only the records with the given ids.
	// Note that this function is just a shorcut for Search on a list of ids.
	Browse(ids []int64) OpenAcademyCourseSet
	// BrowseOne returns a new RecordSet with only the record with the given id.
	// Note that this function is just a shorcut for Search on a given id.
	BrowseOne(id int64) OpenAcademyCourseSet

	CartesianProduct(others ...OpenAcademyCourseSet) []OpenAcademyCourseSet
	// CheckAccessRights verifies that the operation given by "operation" is allowed for
	// the current user according to the access rights.
	//
	// operation must be one of "read", "create", "unlink", "write".
	CheckAccessRights(args webtypes.CheckAccessRightsArgs) bool
	// CheckExecutionPermission panics if the current user is not allowed to execute the given method.
	//
	// If dontPanic is false, this function will panic, otherwise it returns true
	// if the user has the execution permission and false otherwise.
	CheckExecutionPermission(method *models.Method, dontPanic ...bool) bool
	// CheckRecursion verifies that there is no loop in a hierarchical structure of records,
	// by following the parent relationship using the 'Parent' field until a loop is detected or
	// until a top-level record is found.
	//
	// It returns true if no loop was found, false otherwise`,
	CheckRecursion() bool
	// ComputeDisplayName updates the DisplayName field with the result of NameGet
	ComputeDisplayName() *models.ModelData
	// ComputeLastUpdate returns the last datetime at which the record has been updated.
	ComputeLastUpdate() *models.ModelData

	Copy(overrides OpenAcademyCourseData) OpenAcademyCourseSet

	CopyData(overrides OpenAcademyCourseData) OpenAcademyCourseData

	Create(data OpenAcademyCourseData) OpenAcademyCourseSet

	DefaultGet() OpenAcademyCourseData
	// Enqueue queues the execution of the given method with the given arguments on this recordset.
	// description will be the name given to the job.
	Enqueue(description string, method models.Methoder, arguments ...interface{}) QueueJobSet
	// Equals returns true if this RecordSet is the same as other
	// i.e. they are of the same model and have the same ids
	Equals(other OpenAcademyCourseSet) bool
	// ExecuteO2MActions executes the actions on one2many fields given by
	// the list of triplets received from the client
	ExecuteO2MActions(fieldName models.FieldName, info *models.FieldInfo, value interface{}) interface{}
	// Fetch query the database with the current filter and returns a RecordSet
	// with the queries ids.
	//
	// Fetch is lazy and only return ids. Use Load() instead if you want to fetch all fields.
	Fetch() OpenAcademyCourseSet
	// FieldGet returns the definition of the given field.
	// The string, help, and selection (if present) attributes are translated.
	FieldGet(field models.FieldName) *models.FieldInfo
	// FieldsGet returns the definition of each field.
	// The embedded fields are included.
	// The string, help, and selection (if present) attributes are translated.
	//
	// The result map is indexed by the fields JSON names.
	FieldsGet(args models.FieldsGetArgs) map[string]*models.FieldInfo
	// FieldsViewGet is the base implementation of the 'FieldsViewGet' method which
	// gets the detailed composition of the requested view like fields, mixin,
	// view architecture.
	FieldsViewGet(args webtypes.FieldsViewGetParams) *webtypes.FieldsViewData

	Filtered(test func(OpenAcademyCourseSet) bool) OpenAcademyCourseSet
	// FormatRelationFields returns the given data with all relation fields converted to int64 or []int64
	FormatRelationFields(data models.RecordData, fInfos map[string]*models.FieldInfo) models.RecordData
	// GetFormviewAction returns an action to open the document.
	// This method is meant to be overridden in addons that want
	// to give specific view ids for example.`,
	GetFormviewAction() *actions.Action
	// GetFormviewID returns an view id to open the document with.
	// This method is meant to be overridden in addons that want
	// to give specific view ids for example.
	GetFormviewId() string
	// GetRecord returns the Recordset with the given externalID. It panics if the externalID does not exist.
	GetRecord(externalID string) OpenAcademyCourseSet
	// GetToolbar returns a toolbar populated with the actions linked to this model
	GetToolbar() webtypes.Toolbar
	// GroupBy returns a new RecordSet grouped with the given GROUP BY expressions.
	GroupBy(exprs ...models.FieldName) OpenAcademyCourseSet
	// Intersect returns a new RecordCollection with only the records that are both
	// in this RecordCollection and in the other RecordSet.
	Intersect(other OpenAcademyCourseSet) OpenAcademyCourseSet
	// Limit returns a new RecordSet with only the first 'limit' records.
	Limit(limit int) OpenAcademyCourseSet
	// Load looks up cache for fields of the RecordCollection and
	// query database for missing values.
	// fields are the fields to retrieve in the expression format,
	// i.e. "User.Profile.Age" or "user_id.profile_id.age".
	// If no fields are given, all DB columns of the RecordCollection's
	// model are retrieved.
	Load(fields ...models.FieldName) OpenAcademyCourseSet
	// LoadViews returns the data for all the views and filters required in the parameters.
	LoadViews(args webtypes.LoadViewsArgs) *webtypes.LoadViewsData
	// ManageGroupsOnFields adds the invisible attribute to fields nodes if the current
	// user does not belong to one of the groups of the 'groups' attribute
	ManageGroupsOnFields(doc *etree.Document, fieldInfos map[string]*models.FieldInfo)
	// NameGet retrieves the human readable name of this record.`,
	NameGet() string
	// NameSearch searches for records that have a display name matching the given
	// "name" pattern when compared with the given "operator", while also
	// matching the optional search domain ("args").
	//
	// This is used for example to provide suggestions based on a partial
	// value for a relational field. Sometimes be seen as the inverse
	// function of NameGet but it is not guaranteed to be.
	NameSearch(params webtypes.NameSearchParams) []webtypes.RecordIDWithName

	New(data OpenAcademyCourseData) OpenAcademyCourseSet
	// NormalizeM2MData converts the list of triplets received from the client into the final list of ids
	// to keep in the Many2Many relationship of this model through the given field.
	NormalizeM2MData(fieldName models.FieldName, info *models.FieldInfo, value interface{}) interface{}
	// Offset returns a new RecordSet with only the records starting at offset
	Offset(offset int) OpenAcademyCourseSet
	// Onchange returns the values that must be modified according to each field's Onchange
	// method in the pseudo-record given as params.Values`,
	Onchange(params models.OnchangeParams) models.OnchangeResult
	// OrderBy returns a new RecordSet ordered by the given ORDER BY expressions.
	// Each expression contains a field name and optionally one of "asc" or "desc", such as:
	//
	// rs.OrderBy("Company", "Name desc")
	OrderBy(exprs ...string) OpenAcademyCourseSet
	// PostProcessCreateValues updates FK of related records created at the same time.
	//
	// This method is meant to be called with the second returned value of ProcessCreateValues
	// after record creation.
	PostProcessCreateValues(data models.RecordData)
	// PostProcessFilters transforms a map[models.FieldName]models.Conditioner
	// in a map[string][]interface{} which acts as a map of domains.
	PostProcessFilters(in map[models.FieldName]models.Conditioner) map[string][]interface{}
	// ProcessCreateValues updates the given data values for Create method to be
	// compatible with the ORM, in particular for relation fields.
	//
	// It returns a first FieldMap to be used as argument to the Create method, and
	// a second map to be used with a subsequent call to PostProcessCreateValues (for
	// updating FKs pointing to the newly created record).
	ProcessCreateValues(data models.RecordData) (models.RecordData, models.RecordData)
	// ProcessElementAttrs returns a modifiers map according to the domain
	// in attrs of the given element
	ProcessElementAttrs(element *etree.Element, fieldInfos map[string]*models.FieldInfo) map[string]interface{}
	// ProcessFieldElementModifiers modifies the given modifiers map by taking into account:
	// - 'invisible', 'readonly' and 'required' attributes in field tags
	// - 'ReadOnly' and 'Required' parameters of the model's field'
	// It returns the modified map.
	ProcessFieldElementModifiers(element *etree.Element, fieldInfos map[string]*models.FieldInfo, modifiers map[string]interface{}) map[string]interface{}
	// ProcessView makes all the necessary modifications to the view
	// arch and returns the new xml string.`,
	ProcessView(arch *etree.Document, fieldInfos map[string]*models.FieldInfo) string
	// ProcessWriteValues updates the given data values for Write method to be
	// compatible with the ORM, in particular for relation fields
	ProcessWriteValues(data models.RecordData) models.RecordData
	// Read reads the database and returns a slice of FieldMap of the given model.
	Read(fields models.FieldNames) []models.RecordData
	// ReadGroup gets a list of record aggregates according to the given parameters.
	ReadGroup(params webtypes.ReadGroupParams) []models.FieldMap
	// SQLFromCondition returns the WHERE clause sql and arguments corresponding to
	// the given condition.`,
	SQLFromCondition(c *models.Condition) (string, models.SQLParams)

	Search(condition q.OpenAcademyCourseCondition) OpenAcademyCourseSet
	// SearchAll returns a RecordSet with all items of the table, regardless of the
	// current RecordSet query. It is mainly meant to be used on an empty RecordSet.
	SearchAll() OpenAcademyCourseSet

	SearchByName(name string, op operator.Operator, additionalCond q.OpenAcademyCourseCondition, limit int) OpenAcademyCourseSet
	// SearchCount fetch from the database the number of records that match the RecordSet conditions.
	SearchCount() int
	// SearchDomain execute a search on the given domain.
	SearchDomain(domain domains.Domain) CommonMixinSet
	// SearchRead retrieves database records according to the filters defined in params.
	SearchRead(params webtypes.SearchParams) []models.RecordData

	Sorted(less func(OpenAcademyCourseSet, OpenAcademyCourseSet) bool) OpenAcademyCourseSet
	// SortedByField returns a new record set with the same records as rc but sorted by the given field.
	// If reverse is true, the sort is done in reversed order
	SortedByField(namer models.FieldName, reverse bool) OpenAcademyCourseSet
	// SortedDefault returns a new record set with the same records as rc but sorted according
	// to the default order of this model
	SortedDefault() OpenAcademyCourseSet
	// Subtract returns a RecordSet with the Records that are in this
	// RecordCollection but not in the given 'other' one.
	// The result is guaranteed to be a set of unique records.
	Subtract(other OpenAcademyCourseSet) OpenAcademyCourseSet
	// Sudo returns a new RecordSet with the given userID
	// or the superuser ID if not specified
	Sudo(userID ...int64) OpenAcademyCourseSet
	// ToggleActive toggles the Active field of this object if it exists.
	ToggleActive()
	// Union returns a new RecordSet that is the union of this RecordSet and the given
	// "other" RecordSet. The result is guaranteed to be a set of unique records.
	Union(other OpenAcademyCourseSet) OpenAcademyCourseSet
	// Unlink deletes the given records in the database.
	Unlink() int64
	// WebReadGroup returns the result of a read_group (and optionally search for and read records inside each
	// group), and the total number of groups matching the search domain.
	WebReadGroup(params webtypes.WebReadGroupParams) webtypes.WebReadGroupResult
	// WebReadGroupPrivate performs a read_group and optionally a web_search_read for each group.
	WebReadGroupPrivate(params webtypes.WebReadGroupParams) []models.FieldMap
	// WebSearchRead performs a search_read and a search_count.
	WebSearchRead(params webtypes.SearchParams) webtypes.SearchReadResult
	// WithContext returns a copy of the current RecordSet with
	// its context extended by the given key and value.
	WithContext(key string, value interface{}) OpenAcademyCourseSet
	// WithEnv returns a copy of the current RecordSet with the given Environment.
	WithEnv(env models.Environment) OpenAcademyCourseSet
	// WithNewContext returns a copy of the current RecordSet with its context
	// replaced by the given one.
	WithNewContext(context *types.Context) OpenAcademyCourseSet

	Write(data OpenAcademyCourseData) bool
	// Super returns a RecordSet with a modified callstack so that call to the current
	// method will execute the next method layer.
	//
	// This method is meant to be used inside a method layer function to call its parent,
	// such as:
	//
	//    func (rs h.MyRecordSet) MyMethod() string {
	//        res := rs.Super().MyMethod()
	//        res += " ok!"
	//        return res
	//    }
	//
	// Calls to a different method than the current method will call its next layer only
	// if the current method has been called from a layer of the other method. Otherwise,
	// it will be the same as calling the other method directly.
	Super() OpenAcademyCourseSet
	// ModelData returns a new OpenAcademyCourseData object populated with the values
	// of the given FieldMap.
	ModelData(fMap models.FieldMap) OpenAcademyCourseData
	// Records returns a slice with all the records of this RecordSet, as singleton RecordSets
	Records() []OpenAcademyCourseSet
	// First returns the values of the first Record of the RecordSet as a pointer to a OpenAcademyCourseData.
	//
	// If this RecordSet is empty, it returns an empty OpenAcademyCourseData.
	First() OpenAcademyCourseData
	// All returns the values of all Records of the RecordCollection as a slice of OpenAcademyCourseData pointers.
	All() []OpenAcademyCourseData
}

// OpenAcademyCourseData is used to hold values of an OpenAcademyCourse object instance
// when creating or updating a OpenAcademyCourseSet.
type OpenAcademyCourseData interface {
	models.RecordData
	// Get returns the value of the given field.
	//
	// The field can be either its name or is JSON name.
	Get(field models.FieldName) interface{}
	// Has returns true if a value is set for the given field.
	//
	// The field can be either its name or is JSON name.
	Has(field models.FieldName) bool
	// Set sets the given field with the given value.
	// If the field already exists, then it is updated with value.
	// Otherwise, a new entry is inserted.
	//
	// It returns the given OpenAcademyCourseData so that calls can be chained
	Set(field models.FieldName, value interface{}) OpenAcademyCourseData
	// Unset removes the value of the given field if it exists.
	//
	// It returns the given ModelData so that calls can be chained
	Unset(field models.FieldName) OpenAcademyCourseData
	// Copy returns a copy of this OpenAcademyCourseData
	Copy() OpenAcademyCourseData
	// MergeWith updates this OpenAcademyCourseData with the given other OpenAcademyCourseData
	// If a field of the other OpenAcademyCourseData already exists here, the value is overridden,
	// otherwise, the field is inserted.
	MergeWith(other OpenAcademyCourseData)
	// Keys returns the OpenAcademyCourseData keys as a slice of strings
	Keys() (res []string)
	// OrderedKeys returns the keys of this OpenAcademyCourseData ordered.
	//
	// This has the convenient side effect of having shorter paths come before longer paths,
	// which is particularly useful when creating or updating related records.
	OrderedKeys() []string
	// FieldNames returns the OpenAcademyCourseData keys as a slice of FieldNames.
	FieldNames() models.FieldNames
	// CreateDate returns the value of the CreateDate field.
	// If this CreateDate is not set in this OpenAcademyCourseData, then
	// the Go zero value for the type is returned.
	CreateDate() dates.DateTime
	// HasCreateDate returns true if CreateDate is set in this OpenAcademyCourseData
	HasCreateDate() bool
	// SetCreateDate sets the CreateDate field with the given value.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	SetCreateDate(value dates.DateTime) OpenAcademyCourseData
	// UnsetCreateDate removes the value of the CreateDate field if it exists.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	UnsetCreateDate() OpenAcademyCourseData

	// CreateUID returns the value of the CreateUID field.
	// If this CreateUID is not set in this OpenAcademyCourseData, then
	// the Go zero value for the type is returned.
	CreateUID() int64
	// HasCreateUID returns true if CreateUID is set in this OpenAcademyCourseData
	HasCreateUID() bool
	// SetCreateUID sets the CreateUID field with the given value.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	SetCreateUID(value int64) OpenAcademyCourseData
	// UnsetCreateUID removes the value of the CreateUID field if it exists.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	UnsetCreateUID() OpenAcademyCourseData

	// Description returns the value of the Description field.
	// If this Description is not set in this OpenAcademyCourseData, then
	// the Go zero value for the type is returned.
	Description() string
	// HasDescription returns true if Description is set in this OpenAcademyCourseData
	HasDescription() bool
	// SetDescription sets the Description field with the given value.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	SetDescription(value string) OpenAcademyCourseData
	// UnsetDescription removes the value of the Description field if it exists.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	UnsetDescription() OpenAcademyCourseData

	// DisplayName returns the value of the DisplayName field.
	// If this DisplayName is not set in this OpenAcademyCourseData, then
	// the Go zero value for the type is returned.
	DisplayName() string
	// HasDisplayName returns true if DisplayName is set in this OpenAcademyCourseData
	HasDisplayName() bool
	// SetDisplayName sets the DisplayName field with the given value.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	SetDisplayName(value string) OpenAcademyCourseData
	// UnsetDisplayName removes the value of the DisplayName field if it exists.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	UnsetDisplayName() OpenAcademyCourseData

	// HexyaExternalID returns the value of the HexyaExternalID field.
	// If this HexyaExternalID is not set in this OpenAcademyCourseData, then
	// the Go zero value for the type is returned.
	HexyaExternalID() string
	// HasHexyaExternalID returns true if HexyaExternalID is set in this OpenAcademyCourseData
	HasHexyaExternalID() bool
	// SetHexyaExternalID sets the HexyaExternalID field with the given value.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	SetHexyaExternalID(value string) OpenAcademyCourseData
	// UnsetHexyaExternalID removes the value of the HexyaExternalID field if it exists.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	UnsetHexyaExternalID() OpenAcademyCourseData

	// HexyaVersion returns the value of the HexyaVersion field.
	// If this HexyaVersion is not set in this OpenAcademyCourseData, then
	// the Go zero value for the type is returned.
	HexyaVersion() int
	// HasHexyaVersion returns true if HexyaVersion is set in this OpenAcademyCourseData
	HasHexyaVersion() bool
	// SetHexyaVersion sets the HexyaVersion field with the given value.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	SetHexyaVersion(value int) OpenAcademyCourseData
	// UnsetHexyaVersion removes the value of the HexyaVersion field if it exists.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	UnsetHexyaVersion() OpenAcademyCourseData

	// ID returns the value of the ID field.
	// If this ID is not set in this OpenAcademyCourseData, then
	// the Go zero value for the type is returned.
	ID() int64
	// HasID returns true if ID is set in this OpenAcademyCourseData
	HasID() bool
	// SetID sets the ID field with the given value.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	SetID(value int64) OpenAcademyCourseData
	// UnsetID removes the value of the ID field if it exists.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	UnsetID() OpenAcademyCourseData

	// LastUpdate returns the value of the LastUpdate field.
	// If this LastUpdate is not set in this OpenAcademyCourseData, then
	// the Go zero value for the type is returned.
	LastUpdate() dates.DateTime
	// HasLastUpdate returns true if LastUpdate is set in this OpenAcademyCourseData
	HasLastUpdate() bool
	// SetLastUpdate sets the LastUpdate field with the given value.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	SetLastUpdate(value dates.DateTime) OpenAcademyCourseData
	// UnsetLastUpdate removes the value of the LastUpdate field if it exists.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	UnsetLastUpdate() OpenAcademyCourseData

	// Name returns the value of the Name field.
	// If this Name is not set in this OpenAcademyCourseData, then
	// the Go zero value for the type is returned.
	Name() string
	// HasName returns true if Name is set in this OpenAcademyCourseData
	HasName() bool
	// SetName sets the Name field with the given value.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	SetName(value string) OpenAcademyCourseData
	// UnsetName removes the value of the Name field if it exists.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	UnsetName() OpenAcademyCourseData

	// Responsible returns the value of the Responsible field.
	// If this Responsible is not set in this OpenAcademyCourseData, then
	// the Go zero value for the type is returned.
	Responsible() UserSet
	// HasResponsible returns true if Responsible is set in this OpenAcademyCourseData
	HasResponsible() bool
	// SetResponsible sets the Responsible field with the given value.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	SetResponsible(value UserSet) OpenAcademyCourseData
	// UnsetResponsible removes the value of the Responsible field if it exists.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	UnsetResponsible() OpenAcademyCourseData

	// CreateResponsible stores the related UserData to be used to create
	// a related record on the fly for Responsible.
	//
	// This method can be called multiple times to create multiple records
	CreateResponsible(related UserData) OpenAcademyCourseData
	// Sessions returns the value of the Sessions field.
	// If this Sessions is not set in this OpenAcademyCourseData, then
	// the Go zero value for the type is returned.
	Sessions() OpenAcademySessionSet
	// HasSessions returns true if Sessions is set in this OpenAcademyCourseData
	HasSessions() bool
	// SetSessions sets the Sessions field with the given value.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	SetSessions(value OpenAcademySessionSet) OpenAcademyCourseData
	// UnsetSessions removes the value of the Sessions field if it exists.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	UnsetSessions() OpenAcademyCourseData

	// CreateSessions stores the related OpenAcademySessionData to be used to create
	// a related record on the fly for Sessions.
	//
	// This method can be called multiple times to create multiple records
	CreateSessions(related OpenAcademySessionData) OpenAcademyCourseData
	// State returns the value of the State field.
	// If this State is not set in this OpenAcademyCourseData, then
	// the Go zero value for the type is returned.
	State() string
	// HasState returns true if State is set in this OpenAcademyCourseData
	HasState() bool
	// SetState sets the State field with the given value.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	SetState(value string) OpenAcademyCourseData
	// UnsetState removes the value of the State field if it exists.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	UnsetState() OpenAcademyCourseData

	// WriteDate returns the value of the WriteDate field.
	// If this WriteDate is not set in this OpenAcademyCourseData, then
	// the Go zero value for the type is returned.
	WriteDate() dates.DateTime
	// HasWriteDate returns true if WriteDate is set in this OpenAcademyCourseData
	HasWriteDate() bool
	// SetWriteDate sets the WriteDate field with the given value.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	SetWriteDate(value dates.DateTime) OpenAcademyCourseData
	// UnsetWriteDate removes the value of the WriteDate field if it exists.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	UnsetWriteDate() OpenAcademyCourseData

	// WriteUID returns the value of the WriteUID field.
	// If this WriteUID is not set in this OpenAcademyCourseData, then
	// the Go zero value for the type is returned.
	WriteUID() int64
	// HasWriteUID returns true if WriteUID is set in this OpenAcademyCourseData
	HasWriteUID() bool
	// SetWriteUID sets the WriteUID field with the given value.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	SetWriteUID(value int64) OpenAcademyCourseData
	// UnsetWriteUID removes the value of the WriteUID field if it exists.
	// It returns this OpenAcademyCourseData so that calls can be chained.
	UnsetWriteUID() OpenAcademyCourseData
}

// A OpenAcademyCourseGroupAggregateRow holds a row of results of a query with a group by clause
type OpenAcademyCourseGroupAggregateRow interface {
	// Values() returns the values of the actual query
	Values() OpenAcademyCourseData
	// Count is the number of lines aggregated into this one
	Count() int
	// Condition can be used to query the aggregated rows separately if needed
	Condition() q.OpenAcademyCourseCondition
}
