package utils

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Cast[T interface{}](v T) *T { return &v }

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Takes the result of flatmap.Expand for an array of strings
// and returns a []*string
func ExpandStringList(configured []interface{}) []*string {
	vs := make([]*string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, String(val))
		}
	}
	return vs
}

type KeyValue struct {
	Name  string
	Value interface{}
}

func GetFields[T models.AllTheThings](obj T) []KeyValue {
	s := reflect.ValueOf(&obj).Elem()
	typeOfT := s.Type()

	values := make([]KeyValue, 0)

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fieldName := ToSnakeCase(typeOfT.Field(i).Name)
		fieldValue := f.Interface()
		values = append(values, KeyValue{Name: fieldName, Value: fieldValue})
	}

	return values
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func SetResourceDataWithDiagnostic(d *schema.ResourceData, name string, data interface{}, diags *diag.Diagnostics) {
	if data != nil {
		if err := d.Set(name, data); err != nil {
			*diags = append(*diags, diag.FromErr(err)...)
		}
	}
}

func ResourceToDataSource(in *schema.Resource) *schema.Resource {
	for _, schema := range in.Schema {
		schema.Required = false
		schema.Optional = true
		schema.Computed = true
	}

	return in
}

func ExpandList[T any](in []map[string]interface{}) []T {
	items := make([]T, 0)

	for _, item := range in {
		var target T
		if err := mapstructure.Decode(item, &target); err != nil {
			continue
		}

		items = append(items, target)
	}

	return items
}

func ExpandSet[T any](in []interface{}) []T {
	items := make([]T, 0)

	for _, item := range in {
		var target T
		if err := mapstructure.Decode(item, &target); err != nil {
			continue
		}

		items = append(items, target)
	}

	return items
}
