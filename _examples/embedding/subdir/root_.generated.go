// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package subdir

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Query struct {
		InSchemadir        func(childComplexity int) int
		Parentdir          func(childComplexity int) int
		Subdir             func(childComplexity int) int
		__resolve__service func(childComplexity int) int
	}

	_Service struct {
		SDL func(childComplexity int) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]any) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "Query.inSchemadir":
		if e.complexity.Query.InSchemadir == nil {
			break
		}

		return e.complexity.Query.InSchemadir(childComplexity), true

	case "Query.parentdir":
		if e.complexity.Query.Parentdir == nil {
			break
		}

		return e.complexity.Query.Parentdir(childComplexity), true

	case "Query.subdir":
		if e.complexity.Query.Subdir == nil {
			break
		}

		return e.complexity.Query.Subdir(childComplexity), true

	case "Query._service":
		if e.complexity.Query.__resolve__service == nil {
			break
		}

		return e.complexity.Query.__resolve__service(childComplexity), true

	case "_Service.sdl":
		if e.complexity._Service.SDL == nil {
			break
		}

		return e.complexity._Service.SDL(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	opCtx := graphql.GetOperationContext(ctx)
	ec := executionContext{opCtx, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap()
	first := true

	switch opCtx.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, opCtx.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx, ec.OperationContext)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

//go:embed "schemadir/root.graphqls" "subdir.graphqls"
var sourcesFS embed.FS

func sourceData(filename string) string {
	data, err := sourcesFS.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("codegen problem: %s not available", filename))
	}
	return string(data)
}

var sources = []*ast.Source{
	{Name: "schemadir/root.graphqls", Input: sourceData("schemadir/root.graphqls"), BuiltIn: false},
	{Name: "../parent.graphqls", Input: `extend type Query {
    parentdir: String!
}
`, BuiltIn: false},
	{Name: "subdir.graphqls", Input: sourceData("subdir.graphqls"), BuiltIn: false},
	{Name: "federation/directives.graphql", Input: `
	directive @key(fields: _FieldSet!) repeatable on OBJECT | INTERFACE
	directive @requires(fields: _FieldSet!) on FIELD_DEFINITION
	directive @provides(fields: _FieldSet!) on FIELD_DEFINITION
	directive @extends on OBJECT | INTERFACE
	directive @external on FIELD_DEFINITION
	scalar _Any
	scalar _FieldSet
`, BuiltIn: true},
	{Name: "federation/entity.graphql", Input: `
type _Service {
  sdl: String
}

extend type Query {
  _service: _Service!
}
`, BuiltIn: true},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
