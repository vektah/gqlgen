// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputNestedInput(ctx context.Context, obj interface{}) (NestedInput, error) {
	var it NestedInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"field"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "field":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("field"))
			it.Field, err = ec.unmarshalNEmail2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐEmail(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputSpecialInput(ctx context.Context, obj interface{}) (SpecialInput, error) {
	var it SpecialInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"nesting"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "nesting":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("nesting"))
			it.Nesting, err = ec.unmarshalNNestedInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐNestedInput(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNEmail2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐEmail(ctx context.Context, v interface{}) (Email, error) {
	var res Email
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNEmail2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐEmail(ctx context.Context, sel ast.SelectionSet, v Email) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalNNestedInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐNestedInput(ctx context.Context, v interface{}) (*NestedInput, error) {
	res, err := ec.unmarshalInputNestedInput(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNSpecialInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐSpecialInput(ctx context.Context, v interface{}) (SpecialInput, error) {
	res, err := ec.unmarshalInputSpecialInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
