// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/usememos/memos/ent/memo"
	"github.com/usememos/memos/ent/memorelation"
)

// MemoRelationCreate is the builder for creating a MemoRelation entity.
type MemoRelationCreate struct {
	config
	mutation *MemoRelationMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (mrc *MemoRelationCreate) SetType(s string) *MemoRelationCreate {
	mrc.mutation.SetType(s)
	return mrc
}

// SetMemoID sets the "memo_id" field.
func (mrc *MemoRelationCreate) SetMemoID(i int) *MemoRelationCreate {
	mrc.mutation.SetMemoID(i)
	return mrc
}

// SetRelatedMemoID sets the "related_memo_id" field.
func (mrc *MemoRelationCreate) SetRelatedMemoID(i int) *MemoRelationCreate {
	mrc.mutation.SetRelatedMemoID(i)
	return mrc
}

// SetMemo sets the "memo" edge to the Memo entity.
func (mrc *MemoRelationCreate) SetMemo(m *Memo) *MemoRelationCreate {
	return mrc.SetMemoID(m.ID)
}

// SetRelatedMemo sets the "related_memo" edge to the Memo entity.
func (mrc *MemoRelationCreate) SetRelatedMemo(m *Memo) *MemoRelationCreate {
	return mrc.SetRelatedMemoID(m.ID)
}

// Mutation returns the MemoRelationMutation object of the builder.
func (mrc *MemoRelationCreate) Mutation() *MemoRelationMutation {
	return mrc.mutation
}

// Save creates the MemoRelation in the database.
func (mrc *MemoRelationCreate) Save(ctx context.Context) (*MemoRelation, error) {
	return withHooks(ctx, mrc.sqlSave, mrc.mutation, mrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mrc *MemoRelationCreate) SaveX(ctx context.Context) *MemoRelation {
	v, err := mrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mrc *MemoRelationCreate) Exec(ctx context.Context) error {
	_, err := mrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mrc *MemoRelationCreate) ExecX(ctx context.Context) {
	if err := mrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mrc *MemoRelationCreate) check() error {
	if _, ok := mrc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "MemoRelation.type"`)}
	}
	if _, ok := mrc.mutation.MemoID(); !ok {
		return &ValidationError{Name: "memo_id", err: errors.New(`ent: missing required field "MemoRelation.memo_id"`)}
	}
	if _, ok := mrc.mutation.RelatedMemoID(); !ok {
		return &ValidationError{Name: "related_memo_id", err: errors.New(`ent: missing required field "MemoRelation.related_memo_id"`)}
	}
	if _, ok := mrc.mutation.MemoID(); !ok {
		return &ValidationError{Name: "memo", err: errors.New(`ent: missing required edge "MemoRelation.memo"`)}
	}
	if _, ok := mrc.mutation.RelatedMemoID(); !ok {
		return &ValidationError{Name: "related_memo", err: errors.New(`ent: missing required edge "MemoRelation.related_memo"`)}
	}
	return nil
}

func (mrc *MemoRelationCreate) sqlSave(ctx context.Context) (*MemoRelation, error) {
	if err := mrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mrc.mutation.id = &_node.ID
	mrc.mutation.done = true
	return _node, nil
}

func (mrc *MemoRelationCreate) createSpec() (*MemoRelation, *sqlgraph.CreateSpec) {
	var (
		_node = &MemoRelation{config: mrc.config}
		_spec = sqlgraph.NewCreateSpec(memorelation.Table, sqlgraph.NewFieldSpec(memorelation.FieldID, field.TypeInt))
	)
	if value, ok := mrc.mutation.GetType(); ok {
		_spec.SetField(memorelation.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if nodes := mrc.mutation.MemoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   memorelation.MemoTable,
			Columns: []string{memorelation.MemoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MemoID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mrc.mutation.RelatedMemoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   memorelation.RelatedMemoTable,
			Columns: []string{memorelation.RelatedMemoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RelatedMemoID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MemoRelationCreateBulk is the builder for creating many MemoRelation entities in bulk.
type MemoRelationCreateBulk struct {
	config
	err      error
	builders []*MemoRelationCreate
}

// Save creates the MemoRelation entities in the database.
func (mrcb *MemoRelationCreateBulk) Save(ctx context.Context) ([]*MemoRelation, error) {
	if mrcb.err != nil {
		return nil, mrcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mrcb.builders))
	nodes := make([]*MemoRelation, len(mrcb.builders))
	mutators := make([]Mutator, len(mrcb.builders))
	for i := range mrcb.builders {
		func(i int, root context.Context) {
			builder := mrcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemoRelationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mrcb *MemoRelationCreateBulk) SaveX(ctx context.Context) []*MemoRelation {
	v, err := mrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mrcb *MemoRelationCreateBulk) Exec(ctx context.Context) error {
	_, err := mrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mrcb *MemoRelationCreateBulk) ExecX(ctx context.Context) {
	if err := mrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
