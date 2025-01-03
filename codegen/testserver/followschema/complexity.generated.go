// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"errors"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type OverlappingFieldsResolver interface {
	OldFoo(ctx context.Context, obj *OverlappingFields) (int, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _OverlappingFields_oneFoo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_OverlappingFields_oneFoo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Foo, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_OverlappingFields_oneFoo(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "OverlappingFields",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _OverlappingFields_twoFoo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_OverlappingFields_twoFoo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Foo, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_OverlappingFields_twoFoo(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "OverlappingFields",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _OverlappingFields_oldFoo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_OverlappingFields_oldFoo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.OverlappingFields().OldFoo(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_OverlappingFields_oldFoo(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "OverlappingFields",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _OverlappingFields_newFoo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_OverlappingFields_newFoo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.NewFoo, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_OverlappingFields_newFoo(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "OverlappingFields",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _OverlappingFields_new_foo(ctx context.Context, field graphql.CollectedField, obj *OverlappingFields) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_OverlappingFields_new_foo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.NewFoo, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_OverlappingFields_new_foo(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "OverlappingFields",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var overlappingFieldsImplementors = []string{"OverlappingFields"}

func (ec *executionContext) _OverlappingFields(ctx context.Context, sel ast.SelectionSet, obj *OverlappingFields) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, overlappingFieldsImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("OverlappingFields")
		case "oneFoo":
			out.Values[i] = ec._OverlappingFields_oneFoo(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&out.Invalids, 1)
			}
		case "twoFoo":
			out.Values[i] = ec._OverlappingFields_twoFoo(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&out.Invalids, 1)
			}
		case "oldFoo":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._OverlappingFields_oldFoo(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		case "newFoo":
			out.Values[i] = ec._OverlappingFields_newFoo(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&out.Invalids, 1)
			}
		case "new_foo":
			out.Values[i] = ec._OverlappingFields_new_foo(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&out.Invalids, 1)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx, ec.OperationContext)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalOOverlappingFields2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐOverlappingFields(ctx context.Context, sel ast.SelectionSet, v *OverlappingFields) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._OverlappingFields(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
