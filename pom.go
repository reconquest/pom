package pom

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/clbanning/mxj"
)

type ErrMissingField struct {
	Field string
}

func (field ErrMissingField) Error() string {
	return "missing field: " + field.Field
}

func IsMissingField(err error) bool {
	_, ok := err.(ErrMissingField)
	return ok
}

var reInterpret = regexp.MustCompile(`\$\{([^\}]+)\}`)

type Model struct {
	xml        mxj.Map
	properties map[string]string
}

func Unmarshal(data []byte) (*Model, error) {
	xml, err := mxj.NewMapXml(data)
	if err != nil {
		return nil, err
	}

	model := &Model{xml: xml, properties: map[string]string{}}

	return model, nil
}

func (model *Model) SetProperty(key, value string) {
	model.properties[key] = value
}

func (model *Model) Get(key string) (string, error) {
	value, err := model.xml.ValueForPathString("project." + key)
	if err != nil {
		return "", ErrMissingField{key}
	}

	return model.interpret(fmt.Sprint(value))
}

func (model *Model) GetProperty(key string) (string, error) {
	known, ok := model.properties[key]
	if ok {
		return model.interpret(known)
	}

	project, ok := model.xml["project"].(map[string]interface{})
	if !ok {
		return "", ErrMissingField{"project"}
	}

	properties, ok := project["properties"].(map[string]interface{})
	if !ok {
		return "", ErrMissingField{"project.properties"}
	}

	property, ok := properties[key]
	if !ok {
		return "", ErrMissingField{key}
	}

	return model.interpret(fmt.Sprint(property))
}

func (model *Model) interpret(value string) (string, error) {
	matches := reInterpret.FindAllStringSubmatch(value, -1)

	oldnew := []string{}
	for _, group := range matches {
		if len(group) < 1 {
			panic(fmt.Sprintf("group doesn't have submatch: %q", value))
		}

		key := group[1]

		property, err := model.GetProperty(key)
		if err != nil {
			return "", err
		}

		oldnew = append(oldnew, "${"+key+"}", property)
	}

	return strings.NewReplacer(oldnew...).Replace(value), nil
}
