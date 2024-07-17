// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/99designs/gqlgen/plugin/federation/fedruntime"
	"github.com/99designs/gqlgen/plugin/federation/testdata/allthethings/model"
)

var (
	ErrUnknownType  = errors.New("unknown type")
	ErrTypeNotFound = errors.New("type not found")
)

func (ec *executionContext) __resolve__service(ctx context.Context) (fedruntime.Service, error) {
	if ec.DisableIntrospection {
		return fedruntime.Service{}, errors.New("federated introspection disabled")
	}

	var sdl []string

	for _, src := range sources {
		if src.BuiltIn {
			continue
		}
		sdl = append(sdl, src.Input)
	}

	return fedruntime.Service{
		SDL: strings.Join(sdl, "\n"),
	}, nil
}

func (ec *executionContext) __resolve_entities(ctx context.Context, representations []map[string]interface{}) []fedruntime.Entity {
	list := make([]fedruntime.Entity, len(representations))

	repsMap := ec.buildRepresentationGroups(ctx, representations)

	switch len(repsMap) {
	case 0:
		return list
	case 1:
		for typeName, reps := range repsMap {
			ec.resolveEntityGroup(ctx, typeName, reps, list)
		}
		return list
	default:
		var g sync.WaitGroup
		g.Add(len(repsMap))
		for typeName, reps := range repsMap {
			go func(typeName string, reps []EntityWithIndex) {
				ec.resolveEntityGroup(ctx, typeName, reps, list)
				g.Done()
			}(typeName, reps)
		}
		g.Wait()
		return list
	}
}

type EntityWithIndex struct {
	// The index in the original representation array
	index  int
	entity EntityRepresentation
}

// EntityRepresentation is the JSON representation of an entity sent by the Router
// used as the inputs for us to resolve.
//
// We make it a map because we know the top level JSON is always an object.
type EntityRepresentation map[string]any

// We group entities by typename so that we can parallelize their resolution.
// This is particularly helpful when there are entity groups in multi mode.
func (ec *executionContext) buildRepresentationGroups(
	ctx context.Context,
	representations []map[string]any,
) map[string][]EntityWithIndex {
	repsMap := make(map[string][]EntityWithIndex)
	for i, rep := range representations {
		typeName, ok := rep["__typename"].(string)
		if !ok {
			// If there is no __typename, we just skip the representation;
			// we just won't be resolving these unknown types.
			ec.Error(ctx, errors.New("__typename must be an existing string"))
			continue
		}

		repsMap[typeName] = append(repsMap[typeName], EntityWithIndex{
			index:  i,
			entity: rep,
		})
	}

	return repsMap
}

func (ec *executionContext) resolveEntityGroup(
	ctx context.Context,
	typeName string,
	reps []EntityWithIndex,
	list []fedruntime.Entity,
) {
	if isMulti(typeName) {
		err := ec.resolveManyEntities(ctx, typeName, reps, list)
		if err != nil {
			ec.Error(ctx, err)
		}
	} else {
		// if there are multiple entities to resolve, parallelize (similar to
		// graphql.FieldSet.Dispatch)
		var e sync.WaitGroup
		e.Add(len(reps))
		for i, rep := range reps {
			i, rep := i, rep
			go func(i int, rep EntityWithIndex) {
				entity, err := ec.resolveEntity(ctx, typeName, rep.entity)
				if err != nil {
					ec.Error(ctx, err)
				} else {
					list[rep.index] = entity
				}
				e.Done()
			}(i, rep)
		}
		e.Wait()
	}
}

func isMulti(typeName string) bool {
	switch typeName {
	case "MultiHelloMultiKey":
		return true
	default:
		return false
	}
}

func (ec *executionContext) resolveEntity(
	ctx context.Context,
	typeName string,
	rep EntityRepresentation,
) (e fedruntime.Entity, err error) {
	// we need to do our own panic handling, because we may be called in a
	// goroutine, where the usual panic handling can't catch us
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
		}
	}()

	switch typeName {
	case "ExternalExtension":
		resolverName, err := entityResolverNameForExternalExtension(ctx, rep)
		if err != nil {
			return nil, fmt.Errorf(`finding resolver for Entity "ExternalExtension": %w`, err)
		}
		switch resolverName {

		case "findExternalExtensionByUpc":
			id0, err := ec.unmarshalNString2string(ctx, rep["upc"])
			if err != nil {
				return nil, fmt.Errorf(`unmarshalling param 0 for findExternalExtensionByUpc(): %w`, err)
			}
			entity, err := ec.resolvers.Entity().FindExternalExtensionByUpc(ctx, id0)
			if err != nil {
				return nil, fmt.Errorf(`resolving Entity "ExternalExtension": %w`, err)
			}

			return entity, nil
		}
	case "Hello":
		resolverName, err := entityResolverNameForHello(ctx, rep)
		if err != nil {
			return nil, fmt.Errorf(`finding resolver for Entity "Hello": %w`, err)
		}
		switch resolverName {

		case "findHelloByName":
			id0, err := ec.unmarshalNString2string(ctx, rep["name"])
			if err != nil {
				return nil, fmt.Errorf(`unmarshalling param 0 for findHelloByName(): %w`, err)
			}
			entity, err := ec.resolvers.Entity().FindHelloByName(ctx, id0)
			if err != nil {
				return nil, fmt.Errorf(`resolving Entity "Hello": %w`, err)
			}

			return entity, nil
		}
	case "NestedKey":
		resolverName, err := entityResolverNameForNestedKey(ctx, rep)
		if err != nil {
			return nil, fmt.Errorf(`finding resolver for Entity "NestedKey": %w`, err)
		}
		switch resolverName {

		case "findNestedKeyByIDAndHelloName":
			id0, err := ec.unmarshalNString2string(ctx, rep["id"])
			if err != nil {
				return nil, fmt.Errorf(`unmarshalling param 0 for findNestedKeyByIDAndHelloName(): %w`, err)
			}
			id1, err := ec.unmarshalNString2string(ctx, rep["hello"].(map[string]interface{})["name"])
			if err != nil {
				return nil, fmt.Errorf(`unmarshalling param 1 for findNestedKeyByIDAndHelloName(): %w`, err)
			}
			entity, err := ec.resolvers.Entity().FindNestedKeyByIDAndHelloName(ctx, id0, id1)
			if err != nil {
				return nil, fmt.Errorf(`resolving Entity "NestedKey": %w`, err)
			}

			return entity, nil
		}
	case "VeryNestedKey":
		resolverName, err := entityResolverNameForVeryNestedKey(ctx, rep)
		if err != nil {
			return nil, fmt.Errorf(`finding resolver for Entity "VeryNestedKey": %w`, err)
		}
		switch resolverName {

		case "findVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo":
			id0, err := ec.unmarshalNString2string(ctx, rep["id"])
			if err != nil {
				return nil, fmt.Errorf(`unmarshalling param 0 for findVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo(): %w`, err)
			}
			id1, err := ec.unmarshalNString2string(ctx, rep["hello"].(map[string]interface{})["name"])
			if err != nil {
				return nil, fmt.Errorf(`unmarshalling param 1 for findVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo(): %w`, err)
			}
			id2, err := ec.unmarshalNString2string(ctx, rep["world"].(map[string]interface{})["foo"])
			if err != nil {
				return nil, fmt.Errorf(`unmarshalling param 2 for findVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo(): %w`, err)
			}
			id3, err := ec.unmarshalNInt2int(ctx, rep["world"].(map[string]interface{})["bar"])
			if err != nil {
				return nil, fmt.Errorf(`unmarshalling param 3 for findVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo(): %w`, err)
			}
			id4, err := ec.unmarshalNString2string(ctx, rep["more"].(map[string]interface{})["world"].(map[string]interface{})["foo"])
			if err != nil {
				return nil, fmt.Errorf(`unmarshalling param 4 for findVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo(): %w`, err)
			}
			entity, err := ec.resolvers.Entity().FindVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo(ctx, id0, id1, id2, id3, id4)
			if err != nil {
				return nil, fmt.Errorf(`resolving Entity "VeryNestedKey": %w`, err)
			}

			entity.ID, err = ec.unmarshalNString2string(ctx, rep["id"])
			if err != nil {
				return nil, err
			}
			entity.Hello.Secondary, err = ec.unmarshalNString2string(ctx, rep["hello"].(map[string]interface{})["secondary"])
			if err != nil {
				return nil, err
			}
			return entity, nil
		}
	case "World":
		resolverName, err := entityResolverNameForWorld(ctx, rep)
		if err != nil {
			return nil, fmt.Errorf(`finding resolver for Entity "World": %w`, err)
		}
		switch resolverName {

		case "findWorldByFoo":
			id0, err := ec.unmarshalNString2string(ctx, rep["foo"])
			if err != nil {
				return nil, fmt.Errorf(`unmarshalling param 0 for findWorldByFoo(): %w`, err)
			}
			entity, err := ec.resolvers.Entity().FindWorldByFoo(ctx, id0)
			if err != nil {
				return nil, fmt.Errorf(`resolving Entity "World": %w`, err)
			}

			return entity, nil
		case "findWorldByBar":
			id0, err := ec.unmarshalNInt2int(ctx, rep["bar"])
			if err != nil {
				return nil, fmt.Errorf(`unmarshalling param 0 for findWorldByBar(): %w`, err)
			}
			entity, err := ec.resolvers.Entity().FindWorldByBar(ctx, id0)
			if err != nil {
				return nil, fmt.Errorf(`resolving Entity "World": %w`, err)
			}

			return entity, nil
		}

	}
	return nil, fmt.Errorf("%w: %s", ErrUnknownType, typeName)
}

func (ec *executionContext) resolveManyEntities(
	ctx context.Context,
	typeName string,
	reps []EntityWithIndex,
	list []fedruntime.Entity,
) (err error) {
	// we need to do our own panic handling, because we may be called in a
	// goroutine, where the usual panic handling can't catch us
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
		}
	}()

	switch typeName {

	case "MultiHelloMultiKey":
		resolverName, err := entityResolverNameForMultiHelloMultiKey(ctx, reps[0].entity)
		if err != nil {
			return fmt.Errorf(`finding resolver for Entity "MultiHelloMultiKey": %w`, err)
		}
		switch resolverName {

		case "findManyMultiHelloMultiKeyByNames":
			typedReps := make([]*model.MultiHelloMultiKeyByNamesInput, len(reps))

			for i, rep := range reps {
				id0, err := ec.unmarshalNString2string(ctx, rep.entity["name"])
				if err != nil {
					return errors.New(fmt.Sprintf("Field %s undefined in schema.", "name"))
				}

				typedReps[i] = &model.MultiHelloMultiKeyByNamesInput{
					Name: id0,
				}
			}

			entities, err := ec.resolvers.Entity().FindManyMultiHelloMultiKeyByNames(ctx, typedReps)
			if err != nil {
				return err
			}

			for i, entity := range entities {
				list[reps[i].index] = entity
			}
			return nil

		case "findManyMultiHelloMultiKeyByKey2s":
			typedReps := make([]*model.MultiHelloMultiKeyByKey2sInput, len(reps))

			for i, rep := range reps {
				id0, err := ec.unmarshalNString2string(ctx, rep.entity["key2"])
				if err != nil {
					return errors.New(fmt.Sprintf("Field %s undefined in schema.", "key2"))
				}

				typedReps[i] = &model.MultiHelloMultiKeyByKey2sInput{
					Key2: id0,
				}
			}

			entities, err := ec.resolvers.Entity().FindManyMultiHelloMultiKeyByKey2s(ctx, typedReps)
			if err != nil {
				return err
			}

			for i, entity := range entities {
				list[reps[i].index] = entity
			}
			return nil

		default:
			return fmt.Errorf("unknown resolver: %s", resolverName)
		}

	default:
		return errors.New("unknown type: " + typeName)
	}
}

func entityResolverNameForExternalExtension(ctx context.Context, rep EntityRepresentation) (string, error) {
	for {
		var (
			m   EntityRepresentation
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["upc"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findExternalExtensionByUpc", nil
	}
	return "", fmt.Errorf("%w for ExternalExtension", ErrTypeNotFound)
}

func entityResolverNameForHello(ctx context.Context, rep EntityRepresentation) (string, error) {
	for {
		var (
			m   EntityRepresentation
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findHelloByName", nil
	}
	return "", fmt.Errorf("%w for Hello", ErrTypeNotFound)
}

func entityResolverNameForMultiHelloMultiKey(ctx context.Context, rep EntityRepresentation) (string, error) {
	for {
		var (
			m   EntityRepresentation
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findManyMultiHelloMultiKeyByNames", nil
	}
	for {
		var (
			m   EntityRepresentation
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["key2"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findManyMultiHelloMultiKeyByKey2s", nil
	}
	return "", fmt.Errorf("%w for MultiHelloMultiKey", ErrTypeNotFound)
}

func entityResolverNameForNestedKey(ctx context.Context, rep EntityRepresentation) (string, error) {
	for {
		var (
			m   EntityRepresentation
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["id"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		m = rep
		val, ok = m["hello"]
		if !ok {
			break
		}
		if m, ok = val.(map[string]interface{}); !ok {
			break
		}
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findNestedKeyByIDAndHelloName", nil
	}
	return "", fmt.Errorf("%w for NestedKey", ErrTypeNotFound)
}

func entityResolverNameForVeryNestedKey(ctx context.Context, rep EntityRepresentation) (string, error) {
	for {
		var (
			m   EntityRepresentation
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["id"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		m = rep
		val, ok = m["hello"]
		if !ok {
			break
		}
		if m, ok = val.(map[string]interface{}); !ok {
			break
		}
		val, ok = m["name"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		m = rep
		val, ok = m["world"]
		if !ok {
			break
		}
		if m, ok = val.(map[string]interface{}); !ok {
			break
		}
		val, ok = m["foo"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		m = rep
		val, ok = m["world"]
		if !ok {
			break
		}
		if m, ok = val.(map[string]interface{}); !ok {
			break
		}
		val, ok = m["bar"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		m = rep
		val, ok = m["more"]
		if !ok {
			break
		}
		if m, ok = val.(map[string]interface{}); !ok {
			break
		}
		val, ok = m["world"]
		if !ok {
			break
		}
		if m, ok = val.(map[string]interface{}); !ok {
			break
		}
		val, ok = m["foo"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo", nil
	}
	return "", fmt.Errorf("%w for VeryNestedKey", ErrTypeNotFound)
}

func entityResolverNameForWorld(ctx context.Context, rep EntityRepresentation) (string, error) {
	for {
		var (
			m   EntityRepresentation
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["foo"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findWorldByFoo", nil
	}
	for {
		var (
			m   EntityRepresentation
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["bar"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findWorldByBar", nil
	}
	return "", fmt.Errorf("%w for World", ErrTypeNotFound)
}
