// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/99designs/gqlgen/plugin/federation/fedruntime"
)

func (ec *executionContext) __resolve__service(ctx context.Context) (fedruntime.GqlgenService, error) {
	if ec.DisableIntrospection {
		return fedruntime.GqlgenService{}, errors.New("federated introspection disabled")
	}

	var sdl []string

	for _, src := range sources {
		if src.BuiltIn {
			continue
		}
		sdl = append(sdl, src.Input)
	}

	return fedruntime.GqlgenService{
		SDL: strings.Join(sdl, "\n"),
	}, nil
}

func (ec *executionContext) __resolve_entities(ctx context.Context, representations []map[string]interface{}) []fedruntime.GqlgenEntity {
	list := make([]fedruntime.GqlgenEntity, len(representations))
	resolveEntity := func(ctx context.Context, i int, rep map[string]interface{}) (err error) {
		// we need to do our own panic handling, because we may be called in a
		// goroutine, where the usual panic handling can't catch us
		defer func() {
			if r := recover(); r != nil {
				err = ec.Recover(ctx, r)
			}
		}()

		typeName, ok := rep["__typename"].(string)
		if !ok {
			return errors.New("__typename must be an existing string")
		}
		switch typeName {

		case "EmailHost":
			id0, err := ec.unmarshalNString2string(ctx, rep["id"])
			if err != nil {
				return errors.New(fmt.Sprintf("Field %s undefined in schema.", "id"))
			}

			entity, err := ec.resolvers.Entity().FindEmailHostByID(ctx,
				id0)
			if err != nil {
				return err
			}

			list[i] = entity
			return nil

		case "User":
			id0, err := ec.unmarshalNID2string(ctx, rep["id"])
			if err != nil {
				return errors.New(fmt.Sprintf("Field %s undefined in schema.", "id"))
			}

			entity, err := ec.resolvers.Entity().FindUserByID(ctx,
				id0)
			if err != nil {
				return err
			}

			list[i] = entity
			return nil

		default:
			return errors.New("unknown type: " + typeName)
		}
	}

	// if there are multiple entities to resolve, parallelize (similar to
	// graphql.FieldSet.Dispatch)
	switch len(representations) {
	case 0:
		return list
	case 1:
		err := resolveEntity(ctx, 0, representations[0])
		if err != nil {
			ec.Error(ctx, err)
		}
		return list
	default:
		var g sync.WaitGroup
		g.Add(len(representations))
		for i, rep := range representations {
			go func(i int, rep map[string]interface{}) {
				err := resolveEntity(ctx, i, rep)
				if err != nil {
					ec.Error(ctx, err)
				}
				g.Done()
			}(i, rep)
		}
		g.Wait()
		return list
	}
}
