package reflectvisitor

import (
	"reflect"
	"strings"

	"github.com/rs/zerolog/log"
)

type (
	Walker[T any] struct {
		T1             T
		T2             string
		OptionalSetter OptSetter
	}

	Visitor struct {
		root reflect.Type
	}

	OptSetter struct {
		Value *reflect.Value
		IsSet reflect.Value
	}
)

func NewVisitor(root reflect.Type) *Visitor {
	return &Visitor{
		root: root,
	}
}

// Visit traverses the structure of the provided rootInstance using reflection,
// starting from the type stored in the Visitor. It applies the user-provided function f
// to each struct type encountered, allowing custom logic to generate additional types to visit.
// The traversal handles optional types (structs with names prefixed by "Opt") by mapping their
// "Value" field to the underlying type and supports indirect pointer dereferencing.
// The function f receives a Walker containing the current type, a string prefix, and an OptSetter
// representing the value being visited. The traversal continues as long as f returns additional
// Walker instances to process.
func (v *Visitor) Visit(rootInstance *reflect.Value, f func(t Walker[reflect.Type]) []Walker[reflect.Type]) {
	t := v.root
	// Ensure t is a struct type (dereference pointer if needed)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}

	indirect := func(i OptSetter) OptSetter {
		if i.Value != nil && i.Value.IsValid() {
			if i.Value.Kind() == reflect.Ptr {
				val := i.Value.Elem()
				if !val.IsValid() {
					return i
				}
				i.Value = &val
			}

			if i.Value.Kind() != reflect.Struct {
				return i
			}

			boolSetter := i.Value.FieldByName("Set")
			if !boolSetter.IsValid() {
				return i
			}

			newInstance := i.Value.FieldByName("Value")
			if !newInstance.IsValid() {
				return i
			}

			return OptSetter{
				Value: &newInstance,
				IsSet: boolSetter,
			}
		}
		return i
	}

	instance := indirect(OptSetter{
		Value: rootInstance,
	})

	if strings.HasPrefix(t.Name(), "Opt") {
		if mappingType, found := t.FieldByName("Value"); found {
			//overriding optional type Value to be mapped to flags
			t = mappingType.Type
			instance = indirect(instance)
		}
	}

	toVisit := []Walker[reflect.Type]{{t, "", instance}}
	for len(toVisit) > 0 {
		var p Walker[reflect.Type]
		p, toVisit = toVisit[0], toVisit[1:]
		log.Trace().Msgf("Visiting type: %s, prefix: `%s`", p.T1.Name(), p.T2)

		t = p.T1
		localInstance := p.OptionalSetter
		if strings.HasPrefix(p.T1.Name(), "Opt") {
			if mappingType, found := p.T1.FieldByName("Value"); found {
				//overriding optional type Value to be mapped to flags
				t = mappingType.Type
				localInstance = indirect(p.OptionalSetter)
			}
		}

		log.Trace().Msgf("- processing type: %s, prefix: `%s`", t.Name(), p.T2)

		if t.Kind() != reflect.Struct {
			log.Trace().Msgf("Skipping non-struct type: %s", t.Name())
			continue
		}
		adds := f(Walker[reflect.Type]{t, p.T2, localInstance})

		if len(adds) > 0 {
			toVisit = append(toVisit, adds...)
		}
	}
}
