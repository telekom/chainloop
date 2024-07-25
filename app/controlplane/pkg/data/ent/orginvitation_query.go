// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/orginvitation"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/user"
	"github.com/google/uuid"
)

// OrgInvitationQuery is the builder for querying OrgInvitation entities.
type OrgInvitationQuery struct {
	config
	ctx              *QueryContext
	order            []orginvitation.OrderOption
	inters           []Interceptor
	predicates       []predicate.OrgInvitation
	withOrganization *OrganizationQuery
	withSender       *UserQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrgInvitationQuery builder.
func (oiq *OrgInvitationQuery) Where(ps ...predicate.OrgInvitation) *OrgInvitationQuery {
	oiq.predicates = append(oiq.predicates, ps...)
	return oiq
}

// Limit the number of records to be returned by this query.
func (oiq *OrgInvitationQuery) Limit(limit int) *OrgInvitationQuery {
	oiq.ctx.Limit = &limit
	return oiq
}

// Offset to start from.
func (oiq *OrgInvitationQuery) Offset(offset int) *OrgInvitationQuery {
	oiq.ctx.Offset = &offset
	return oiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oiq *OrgInvitationQuery) Unique(unique bool) *OrgInvitationQuery {
	oiq.ctx.Unique = &unique
	return oiq
}

// Order specifies how the records should be ordered.
func (oiq *OrgInvitationQuery) Order(o ...orginvitation.OrderOption) *OrgInvitationQuery {
	oiq.order = append(oiq.order, o...)
	return oiq
}

// QueryOrganization chains the current query on the "organization" edge.
func (oiq *OrgInvitationQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: oiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orginvitation.Table, orginvitation.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, orginvitation.OrganizationTable, orginvitation.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(oiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySender chains the current query on the "sender" edge.
func (oiq *OrgInvitationQuery) QuerySender() *UserQuery {
	query := (&UserClient{config: oiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orginvitation.Table, orginvitation.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, orginvitation.SenderTable, orginvitation.SenderColumn),
		)
		fromU = sqlgraph.SetNeighbors(oiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrgInvitation entity from the query.
// Returns a *NotFoundError when no OrgInvitation was found.
func (oiq *OrgInvitationQuery) First(ctx context.Context) (*OrgInvitation, error) {
	nodes, err := oiq.Limit(1).All(setContextOp(ctx, oiq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orginvitation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oiq *OrgInvitationQuery) FirstX(ctx context.Context) *OrgInvitation {
	node, err := oiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrgInvitation ID from the query.
// Returns a *NotFoundError when no OrgInvitation ID was found.
func (oiq *OrgInvitationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oiq.Limit(1).IDs(setContextOp(ctx, oiq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orginvitation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oiq *OrgInvitationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := oiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrgInvitation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrgInvitation entity is found.
// Returns a *NotFoundError when no OrgInvitation entities are found.
func (oiq *OrgInvitationQuery) Only(ctx context.Context) (*OrgInvitation, error) {
	nodes, err := oiq.Limit(2).All(setContextOp(ctx, oiq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orginvitation.Label}
	default:
		return nil, &NotSingularError{orginvitation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oiq *OrgInvitationQuery) OnlyX(ctx context.Context) *OrgInvitation {
	node, err := oiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrgInvitation ID in the query.
// Returns a *NotSingularError when more than one OrgInvitation ID is found.
// Returns a *NotFoundError when no entities are found.
func (oiq *OrgInvitationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oiq.Limit(2).IDs(setContextOp(ctx, oiq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orginvitation.Label}
	default:
		err = &NotSingularError{orginvitation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oiq *OrgInvitationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := oiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrgInvitations.
func (oiq *OrgInvitationQuery) All(ctx context.Context) ([]*OrgInvitation, error) {
	ctx = setContextOp(ctx, oiq.ctx, "All")
	if err := oiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrgInvitation, *OrgInvitationQuery]()
	return withInterceptors[[]*OrgInvitation](ctx, oiq, qr, oiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oiq *OrgInvitationQuery) AllX(ctx context.Context) []*OrgInvitation {
	nodes, err := oiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrgInvitation IDs.
func (oiq *OrgInvitationQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if oiq.ctx.Unique == nil && oiq.path != nil {
		oiq.Unique(true)
	}
	ctx = setContextOp(ctx, oiq.ctx, "IDs")
	if err = oiq.Select(orginvitation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oiq *OrgInvitationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := oiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oiq *OrgInvitationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oiq.ctx, "Count")
	if err := oiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oiq, querierCount[*OrgInvitationQuery](), oiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oiq *OrgInvitationQuery) CountX(ctx context.Context) int {
	count, err := oiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oiq *OrgInvitationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oiq.ctx, "Exist")
	switch _, err := oiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oiq *OrgInvitationQuery) ExistX(ctx context.Context) bool {
	exist, err := oiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrgInvitationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oiq *OrgInvitationQuery) Clone() *OrgInvitationQuery {
	if oiq == nil {
		return nil
	}
	return &OrgInvitationQuery{
		config:           oiq.config,
		ctx:              oiq.ctx.Clone(),
		order:            append([]orginvitation.OrderOption{}, oiq.order...),
		inters:           append([]Interceptor{}, oiq.inters...),
		predicates:       append([]predicate.OrgInvitation{}, oiq.predicates...),
		withOrganization: oiq.withOrganization.Clone(),
		withSender:       oiq.withSender.Clone(),
		// clone intermediate query.
		sql:  oiq.sql.Clone(),
		path: oiq.path,
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (oiq *OrgInvitationQuery) WithOrganization(opts ...func(*OrganizationQuery)) *OrgInvitationQuery {
	query := (&OrganizationClient{config: oiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oiq.withOrganization = query
	return oiq
}

// WithSender tells the query-builder to eager-load the nodes that are connected to
// the "sender" edge. The optional arguments are used to configure the query builder of the edge.
func (oiq *OrgInvitationQuery) WithSender(opts ...func(*UserQuery)) *OrgInvitationQuery {
	query := (&UserClient{config: oiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oiq.withSender = query
	return oiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ReceiverEmail string `json:"receiver_email,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrgInvitation.Query().
//		GroupBy(orginvitation.FieldReceiverEmail).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oiq *OrgInvitationQuery) GroupBy(field string, fields ...string) *OrgInvitationGroupBy {
	oiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrgInvitationGroupBy{build: oiq}
	grbuild.flds = &oiq.ctx.Fields
	grbuild.label = orginvitation.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ReceiverEmail string `json:"receiver_email,omitempty"`
//	}
//
//	client.OrgInvitation.Query().
//		Select(orginvitation.FieldReceiverEmail).
//		Scan(ctx, &v)
func (oiq *OrgInvitationQuery) Select(fields ...string) *OrgInvitationSelect {
	oiq.ctx.Fields = append(oiq.ctx.Fields, fields...)
	sbuild := &OrgInvitationSelect{OrgInvitationQuery: oiq}
	sbuild.label = orginvitation.Label
	sbuild.flds, sbuild.scan = &oiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrgInvitationSelect configured with the given aggregations.
func (oiq *OrgInvitationQuery) Aggregate(fns ...AggregateFunc) *OrgInvitationSelect {
	return oiq.Select().Aggregate(fns...)
}

func (oiq *OrgInvitationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oiq); err != nil {
				return err
			}
		}
	}
	for _, f := range oiq.ctx.Fields {
		if !orginvitation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oiq.path != nil {
		prev, err := oiq.path(ctx)
		if err != nil {
			return err
		}
		oiq.sql = prev
	}
	return nil
}

func (oiq *OrgInvitationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrgInvitation, error) {
	var (
		nodes       = []*OrgInvitation{}
		_spec       = oiq.querySpec()
		loadedTypes = [2]bool{
			oiq.withOrganization != nil,
			oiq.withSender != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrgInvitation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrgInvitation{config: oiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(oiq.modifiers) > 0 {
		_spec.Modifiers = oiq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oiq.withOrganization; query != nil {
		if err := oiq.loadOrganization(ctx, query, nodes, nil,
			func(n *OrgInvitation, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := oiq.withSender; query != nil {
		if err := oiq.loadSender(ctx, query, nodes, nil,
			func(n *OrgInvitation, e *User) { n.Edges.Sender = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oiq *OrgInvitationQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*OrgInvitation, init func(*OrgInvitation), assign func(*OrgInvitation, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*OrgInvitation)
	for i := range nodes {
		fk := nodes[i].OrganizationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (oiq *OrgInvitationQuery) loadSender(ctx context.Context, query *UserQuery, nodes []*OrgInvitation, init func(*OrgInvitation), assign func(*OrgInvitation, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*OrgInvitation)
	for i := range nodes {
		fk := nodes[i].SenderID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "sender_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (oiq *OrgInvitationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oiq.querySpec()
	if len(oiq.modifiers) > 0 {
		_spec.Modifiers = oiq.modifiers
	}
	_spec.Node.Columns = oiq.ctx.Fields
	if len(oiq.ctx.Fields) > 0 {
		_spec.Unique = oiq.ctx.Unique != nil && *oiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oiq.driver, _spec)
}

func (oiq *OrgInvitationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orginvitation.Table, orginvitation.Columns, sqlgraph.NewFieldSpec(orginvitation.FieldID, field.TypeUUID))
	_spec.From = oiq.sql
	if unique := oiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oiq.path != nil {
		_spec.Unique = true
	}
	if fields := oiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orginvitation.FieldID)
		for i := range fields {
			if fields[i] != orginvitation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if oiq.withOrganization != nil {
			_spec.Node.AddColumnOnce(orginvitation.FieldOrganizationID)
		}
		if oiq.withSender != nil {
			_spec.Node.AddColumnOnce(orginvitation.FieldSenderID)
		}
	}
	if ps := oiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oiq *OrgInvitationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oiq.driver.Dialect())
	t1 := builder.Table(orginvitation.Table)
	columns := oiq.ctx.Fields
	if len(columns) == 0 {
		columns = orginvitation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oiq.sql != nil {
		selector = oiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oiq.ctx.Unique != nil && *oiq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range oiq.modifiers {
		m(selector)
	}
	for _, p := range oiq.predicates {
		p(selector)
	}
	for _, p := range oiq.order {
		p(selector)
	}
	if offset := oiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (oiq *OrgInvitationQuery) ForUpdate(opts ...sql.LockOption) *OrgInvitationQuery {
	if oiq.driver.Dialect() == dialect.Postgres {
		oiq.Unique(false)
	}
	oiq.modifiers = append(oiq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return oiq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (oiq *OrgInvitationQuery) ForShare(opts ...sql.LockOption) *OrgInvitationQuery {
	if oiq.driver.Dialect() == dialect.Postgres {
		oiq.Unique(false)
	}
	oiq.modifiers = append(oiq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return oiq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (oiq *OrgInvitationQuery) Modify(modifiers ...func(s *sql.Selector)) *OrgInvitationSelect {
	oiq.modifiers = append(oiq.modifiers, modifiers...)
	return oiq.Select()
}

// OrgInvitationGroupBy is the group-by builder for OrgInvitation entities.
type OrgInvitationGroupBy struct {
	selector
	build *OrgInvitationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oigb *OrgInvitationGroupBy) Aggregate(fns ...AggregateFunc) *OrgInvitationGroupBy {
	oigb.fns = append(oigb.fns, fns...)
	return oigb
}

// Scan applies the selector query and scans the result into the given value.
func (oigb *OrgInvitationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oigb.build.ctx, "GroupBy")
	if err := oigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrgInvitationQuery, *OrgInvitationGroupBy](ctx, oigb.build, oigb, oigb.build.inters, v)
}

func (oigb *OrgInvitationGroupBy) sqlScan(ctx context.Context, root *OrgInvitationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(oigb.fns))
	for _, fn := range oigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*oigb.flds)+len(oigb.fns))
		for _, f := range *oigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*oigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrgInvitationSelect is the builder for selecting fields of OrgInvitation entities.
type OrgInvitationSelect struct {
	*OrgInvitationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ois *OrgInvitationSelect) Aggregate(fns ...AggregateFunc) *OrgInvitationSelect {
	ois.fns = append(ois.fns, fns...)
	return ois
}

// Scan applies the selector query and scans the result into the given value.
func (ois *OrgInvitationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ois.ctx, "Select")
	if err := ois.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrgInvitationQuery, *OrgInvitationSelect](ctx, ois.OrgInvitationQuery, ois, ois.inters, v)
}

func (ois *OrgInvitationSelect) sqlScan(ctx context.Context, root *OrgInvitationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ois.fns))
	for _, fn := range ois.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ois.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ois.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ois *OrgInvitationSelect) Modify(modifiers ...func(s *sql.Selector)) *OrgInvitationSelect {
	ois.modifiers = append(ois.modifiers, modifiers...)
	return ois
}
