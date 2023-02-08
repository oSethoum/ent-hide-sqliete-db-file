package ent

import (
	"app/ent/predicate"
	"app/ent/user"
	"errors"
	"fmt"
	"strings"
)

type OrderField struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type UserQueryInput struct {
	Select []string        `json:"select"`
	Offset *int            `json:"offset"`
	Limit  *int            `json:"limit"`
	Where  *UserWhereInput `json:"where"`
	Order  []OrderField    `json:"order"`
}

type UserWhereInput struct {
	Not *UserWhereInput   `json:"not"`
	Or  []*UserWhereInput `json:"or"`
	And []*UserWhereInput `json:"and"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`
	NameEqualFold    *string  `json:"nameEqualsFold,omitempty"`
}

var ErrEmptyUserWhereInput = errors.New("ent: empty UserWhereInput")

func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	// "id" field predicates.
	if v := i.ID; v != nil {
		predicates = append(predicates, user.IDEQ(*v))
	}
	if v := i.IDNEQ; v != nil {
		predicates = append(predicates, user.IDNEQ(*v))
	}
	if v := i.IDIn; len(v) > 0 {
		predicates = append(predicates, user.IDIn(v...))
	}
	if v := i.IDNotIn; len(v) > 0 {
		predicates = append(predicates, user.IDNotIn(v...))
	}
	if v := i.IDGT; v != nil {
		predicates = append(predicates, user.IDGT(*v))
	}
	if v := i.IDGTE; v != nil {
		predicates = append(predicates, user.IDGTE(*v))
	}
	if v := i.IDLT; v != nil {
		predicates = append(predicates, user.IDLT(*v))
	}
	if v := i.IDLTE; v != nil {
		predicates = append(predicates, user.IDLTE(*v))
	}

	// "name" field predicates.
	if v := i.Name; v != nil {
		predicates = append(predicates, user.NameEQ(*v))
	}
	if v := i.NameNEQ; v != nil {
		predicates = append(predicates, user.NameNEQ(*v))
	}
	if v := i.NameIn; len(v) > 0 {
		predicates = append(predicates, user.NameIn(v...))
	}
	if v := i.NameNotIn; len(v) > 0 {
		predicates = append(predicates, user.NameNotIn(v...))
	}
	if v := i.NameGT; v != nil {
		predicates = append(predicates, user.NameGT(*v))
	}
	if v := i.NameGTE; v != nil {
		predicates = append(predicates, user.NameGTE(*v))
	}
	if v := i.NameLT; v != nil {
		predicates = append(predicates, user.NameLT(*v))
	}
	if v := i.NameLTE; v != nil {
		predicates = append(predicates, user.NameLTE(*v))
	}
	if v := i.NameContains; v != nil {
		predicates = append(predicates, user.NameContains(*v))
	}
	if v := i.NameHasPrefix; v != nil {
		predicates = append(predicates, user.NameHasPrefix(*v))
	}
	if v := i.NameHasSuffix; v != nil {
		predicates = append(predicates, user.NameHasSuffix(*v))
	}
	if v := i.NameHasSuffix; v != nil {
		predicates = append(predicates, user.NameHasSuffix(*v))
	}
	if v := i.NameContainsFold; v != nil {
		predicates = append(predicates, user.NameContainsFold(*v))
	}
	if v := i.NameEqualFold; v != nil {
		predicates = append(predicates, user.NameEqualFold(*v))
	}

	switch len(predicates) {
	case 0:
		return nil, ErrEmptyUserWhereInput
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}

func (q *UserQuery) SetInput(i *UserQueryInput) *UserQuery {
	if i == nil {
		return q
	}
	if i.Where != nil {
		ps, err := i.Where.P()
		if err == nil {
			q.Where(ps)
		}
	}
	if v := i.Order; len(v) > 0 {
		for _, o := range v {
			if strings.ToUpper(o.Direction) == "DESC" {
				q.Order(Desc(o.Field))
			} else {
				q.Order(Asc(o.Field))
			}
		}
	}
	if len(i.Select) > 0 {
		q.Select(i.Select...)
	}
	if v := i.Offset; v != nil {
		q.Offset(*v)
	}
	if v := i.Limit; v != nil {
		q.Limit(*v)
	}
	return q
}
