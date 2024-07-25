// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/integrationattachment"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/referrer"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/robotaccount"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflow"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflowcontract"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflowrun"
	"github.com/google/uuid"
)

// WorkflowQuery is the builder for querying Workflow entities.
type WorkflowQuery struct {
	config
	ctx                        *QueryContext
	order                      []workflow.OrderOption
	inters                     []Interceptor
	predicates                 []predicate.Workflow
	withRobotaccounts          *RobotAccountQuery
	withWorkflowruns           *WorkflowRunQuery
	withOrganization           *OrganizationQuery
	withContract               *WorkflowContractQuery
	withIntegrationAttachments *IntegrationAttachmentQuery
	withReferrers              *ReferrerQuery
	withFKs                    bool
	modifiers                  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkflowQuery builder.
func (wq *WorkflowQuery) Where(ps ...predicate.Workflow) *WorkflowQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit the number of records to be returned by this query.
func (wq *WorkflowQuery) Limit(limit int) *WorkflowQuery {
	wq.ctx.Limit = &limit
	return wq
}

// Offset to start from.
func (wq *WorkflowQuery) Offset(offset int) *WorkflowQuery {
	wq.ctx.Offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WorkflowQuery) Unique(unique bool) *WorkflowQuery {
	wq.ctx.Unique = &unique
	return wq
}

// Order specifies how the records should be ordered.
func (wq *WorkflowQuery) Order(o ...workflow.OrderOption) *WorkflowQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryRobotaccounts chains the current query on the "robotaccounts" edge.
func (wq *WorkflowQuery) QueryRobotaccounts() *RobotAccountQuery {
	query := (&RobotAccountClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workflow.Table, workflow.FieldID, selector),
			sqlgraph.To(robotaccount.Table, robotaccount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workflow.RobotaccountsTable, workflow.RobotaccountsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkflowruns chains the current query on the "workflowruns" edge.
func (wq *WorkflowQuery) QueryWorkflowruns() *WorkflowRunQuery {
	query := (&WorkflowRunClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workflow.Table, workflow.FieldID, selector),
			sqlgraph.To(workflowrun.Table, workflowrun.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workflow.WorkflowrunsTable, workflow.WorkflowrunsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (wq *WorkflowQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workflow.Table, workflow.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workflow.OrganizationTable, workflow.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryContract chains the current query on the "contract" edge.
func (wq *WorkflowQuery) QueryContract() *WorkflowContractQuery {
	query := (&WorkflowContractClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workflow.Table, workflow.FieldID, selector),
			sqlgraph.To(workflowcontract.Table, workflowcontract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, workflow.ContractTable, workflow.ContractColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIntegrationAttachments chains the current query on the "integration_attachments" edge.
func (wq *WorkflowQuery) QueryIntegrationAttachments() *IntegrationAttachmentQuery {
	query := (&IntegrationAttachmentClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workflow.Table, workflow.FieldID, selector),
			sqlgraph.To(integrationattachment.Table, integrationattachment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, workflow.IntegrationAttachmentsTable, workflow.IntegrationAttachmentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReferrers chains the current query on the "referrers" edge.
func (wq *WorkflowQuery) QueryReferrers() *ReferrerQuery {
	query := (&ReferrerClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workflow.Table, workflow.FieldID, selector),
			sqlgraph.To(referrer.Table, referrer.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, workflow.ReferrersTable, workflow.ReferrersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Workflow entity from the query.
// Returns a *NotFoundError when no Workflow was found.
func (wq *WorkflowQuery) First(ctx context.Context) (*Workflow, error) {
	nodes, err := wq.Limit(1).All(setContextOp(ctx, wq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workflow.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WorkflowQuery) FirstX(ctx context.Context) *Workflow {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Workflow ID from the query.
// Returns a *NotFoundError when no Workflow ID was found.
func (wq *WorkflowQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wq.Limit(1).IDs(setContextOp(ctx, wq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workflow.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WorkflowQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Workflow entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Workflow entity is found.
// Returns a *NotFoundError when no Workflow entities are found.
func (wq *WorkflowQuery) Only(ctx context.Context) (*Workflow, error) {
	nodes, err := wq.Limit(2).All(setContextOp(ctx, wq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workflow.Label}
	default:
		return nil, &NotSingularError{workflow.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WorkflowQuery) OnlyX(ctx context.Context) *Workflow {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Workflow ID in the query.
// Returns a *NotSingularError when more than one Workflow ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WorkflowQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wq.Limit(2).IDs(setContextOp(ctx, wq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workflow.Label}
	default:
		err = &NotSingularError{workflow.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WorkflowQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Workflows.
func (wq *WorkflowQuery) All(ctx context.Context) ([]*Workflow, error) {
	ctx = setContextOp(ctx, wq.ctx, "All")
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Workflow, *WorkflowQuery]()
	return withInterceptors[[]*Workflow](ctx, wq, qr, wq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wq *WorkflowQuery) AllX(ctx context.Context) []*Workflow {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Workflow IDs.
func (wq *WorkflowQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if wq.ctx.Unique == nil && wq.path != nil {
		wq.Unique(true)
	}
	ctx = setContextOp(ctx, wq.ctx, "IDs")
	if err = wq.Select(workflow.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WorkflowQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WorkflowQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wq.ctx, "Count")
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wq, querierCount[*WorkflowQuery](), wq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WorkflowQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WorkflowQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wq.ctx, "Exist")
	switch _, err := wq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WorkflowQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkflowQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WorkflowQuery) Clone() *WorkflowQuery {
	if wq == nil {
		return nil
	}
	return &WorkflowQuery{
		config:                     wq.config,
		ctx:                        wq.ctx.Clone(),
		order:                      append([]workflow.OrderOption{}, wq.order...),
		inters:                     append([]Interceptor{}, wq.inters...),
		predicates:                 append([]predicate.Workflow{}, wq.predicates...),
		withRobotaccounts:          wq.withRobotaccounts.Clone(),
		withWorkflowruns:           wq.withWorkflowruns.Clone(),
		withOrganization:           wq.withOrganization.Clone(),
		withContract:               wq.withContract.Clone(),
		withIntegrationAttachments: wq.withIntegrationAttachments.Clone(),
		withReferrers:              wq.withReferrers.Clone(),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

// WithRobotaccounts tells the query-builder to eager-load the nodes that are connected to
// the "robotaccounts" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkflowQuery) WithRobotaccounts(opts ...func(*RobotAccountQuery)) *WorkflowQuery {
	query := (&RobotAccountClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withRobotaccounts = query
	return wq
}

// WithWorkflowruns tells the query-builder to eager-load the nodes that are connected to
// the "workflowruns" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkflowQuery) WithWorkflowruns(opts ...func(*WorkflowRunQuery)) *WorkflowQuery {
	query := (&WorkflowRunClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withWorkflowruns = query
	return wq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkflowQuery) WithOrganization(opts ...func(*OrganizationQuery)) *WorkflowQuery {
	query := (&OrganizationClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withOrganization = query
	return wq
}

// WithContract tells the query-builder to eager-load the nodes that are connected to
// the "contract" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkflowQuery) WithContract(opts ...func(*WorkflowContractQuery)) *WorkflowQuery {
	query := (&WorkflowContractClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withContract = query
	return wq
}

// WithIntegrationAttachments tells the query-builder to eager-load the nodes that are connected to
// the "integration_attachments" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkflowQuery) WithIntegrationAttachments(opts ...func(*IntegrationAttachmentQuery)) *WorkflowQuery {
	query := (&IntegrationAttachmentClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withIntegrationAttachments = query
	return wq
}

// WithReferrers tells the query-builder to eager-load the nodes that are connected to
// the "referrers" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkflowQuery) WithReferrers(opts ...func(*ReferrerQuery)) *WorkflowQuery {
	query := (&ReferrerClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withReferrers = query
	return wq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Workflow.Query().
//		GroupBy(workflow.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wq *WorkflowQuery) GroupBy(field string, fields ...string) *WorkflowGroupBy {
	wq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WorkflowGroupBy{build: wq}
	grbuild.flds = &wq.ctx.Fields
	grbuild.label = workflow.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Workflow.Query().
//		Select(workflow.FieldName).
//		Scan(ctx, &v)
func (wq *WorkflowQuery) Select(fields ...string) *WorkflowSelect {
	wq.ctx.Fields = append(wq.ctx.Fields, fields...)
	sbuild := &WorkflowSelect{WorkflowQuery: wq}
	sbuild.label = workflow.Label
	sbuild.flds, sbuild.scan = &wq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WorkflowSelect configured with the given aggregations.
func (wq *WorkflowQuery) Aggregate(fns ...AggregateFunc) *WorkflowSelect {
	return wq.Select().Aggregate(fns...)
}

func (wq *WorkflowQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wq); err != nil {
				return err
			}
		}
	}
	for _, f := range wq.ctx.Fields {
		if !workflow.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WorkflowQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Workflow, error) {
	var (
		nodes       = []*Workflow{}
		withFKs     = wq.withFKs
		_spec       = wq.querySpec()
		loadedTypes = [6]bool{
			wq.withRobotaccounts != nil,
			wq.withWorkflowruns != nil,
			wq.withOrganization != nil,
			wq.withContract != nil,
			wq.withIntegrationAttachments != nil,
			wq.withReferrers != nil,
		}
	)
	if wq.withContract != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, workflow.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Workflow).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Workflow{config: wq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wq.withRobotaccounts; query != nil {
		if err := wq.loadRobotaccounts(ctx, query, nodes,
			func(n *Workflow) { n.Edges.Robotaccounts = []*RobotAccount{} },
			func(n *Workflow, e *RobotAccount) { n.Edges.Robotaccounts = append(n.Edges.Robotaccounts, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withWorkflowruns; query != nil {
		if err := wq.loadWorkflowruns(ctx, query, nodes,
			func(n *Workflow) { n.Edges.Workflowruns = []*WorkflowRun{} },
			func(n *Workflow, e *WorkflowRun) { n.Edges.Workflowruns = append(n.Edges.Workflowruns, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withOrganization; query != nil {
		if err := wq.loadOrganization(ctx, query, nodes, nil,
			func(n *Workflow, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withContract; query != nil {
		if err := wq.loadContract(ctx, query, nodes, nil,
			func(n *Workflow, e *WorkflowContract) { n.Edges.Contract = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withIntegrationAttachments; query != nil {
		if err := wq.loadIntegrationAttachments(ctx, query, nodes,
			func(n *Workflow) { n.Edges.IntegrationAttachments = []*IntegrationAttachment{} },
			func(n *Workflow, e *IntegrationAttachment) {
				n.Edges.IntegrationAttachments = append(n.Edges.IntegrationAttachments, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := wq.withReferrers; query != nil {
		if err := wq.loadReferrers(ctx, query, nodes,
			func(n *Workflow) { n.Edges.Referrers = []*Referrer{} },
			func(n *Workflow, e *Referrer) { n.Edges.Referrers = append(n.Edges.Referrers, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wq *WorkflowQuery) loadRobotaccounts(ctx context.Context, query *RobotAccountQuery, nodes []*Workflow, init func(*Workflow), assign func(*Workflow, *RobotAccount)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Workflow)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.RobotAccount(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(workflow.RobotaccountsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.workflow_robotaccounts
		if fk == nil {
			return fmt.Errorf(`foreign-key "workflow_robotaccounts" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "workflow_robotaccounts" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wq *WorkflowQuery) loadWorkflowruns(ctx context.Context, query *WorkflowRunQuery, nodes []*Workflow, init func(*Workflow), assign func(*Workflow, *WorkflowRun)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Workflow)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.WorkflowRun(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(workflow.WorkflowrunsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.workflow_workflowruns
		if fk == nil {
			return fmt.Errorf(`foreign-key "workflow_workflowruns" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "workflow_workflowruns" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wq *WorkflowQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*Workflow, init func(*Workflow), assign func(*Workflow, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Workflow)
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
func (wq *WorkflowQuery) loadContract(ctx context.Context, query *WorkflowContractQuery, nodes []*Workflow, init func(*Workflow), assign func(*Workflow, *WorkflowContract)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Workflow)
	for i := range nodes {
		if nodes[i].workflow_contract == nil {
			continue
		}
		fk := *nodes[i].workflow_contract
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(workflowcontract.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "workflow_contract" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WorkflowQuery) loadIntegrationAttachments(ctx context.Context, query *IntegrationAttachmentQuery, nodes []*Workflow, init func(*Workflow), assign func(*Workflow, *IntegrationAttachment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Workflow)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.IntegrationAttachment(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(workflow.IntegrationAttachmentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.integration_attachment_workflow
		if fk == nil {
			return fmt.Errorf(`foreign-key "integration_attachment_workflow" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "integration_attachment_workflow" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wq *WorkflowQuery) loadReferrers(ctx context.Context, query *ReferrerQuery, nodes []*Workflow, init func(*Workflow), assign func(*Workflow, *Referrer)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Workflow)
	nids := make(map[uuid.UUID]map[*Workflow]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(workflow.ReferrersTable)
		s.Join(joinT).On(s.C(referrer.FieldID), joinT.C(workflow.ReferrersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(workflow.ReferrersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(workflow.ReferrersPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Workflow]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Referrer](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "referrers" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (wq *WorkflowQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	_spec.Node.Columns = wq.ctx.Fields
	if len(wq.ctx.Fields) > 0 {
		_spec.Unique = wq.ctx.Unique != nil && *wq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WorkflowQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(workflow.Table, workflow.Columns, sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID))
	_spec.From = wq.sql
	if unique := wq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wq.path != nil {
		_spec.Unique = true
	}
	if fields := wq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflow.FieldID)
		for i := range fields {
			if fields[i] != workflow.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if wq.withOrganization != nil {
			_spec.Node.AddColumnOnce(workflow.FieldOrganizationID)
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WorkflowQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(workflow.Table)
	columns := wq.ctx.Fields
	if len(columns) == 0 {
		columns = workflow.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.ctx.Unique != nil && *wq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range wq.modifiers {
		m(selector)
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (wq *WorkflowQuery) ForUpdate(opts ...sql.LockOption) *WorkflowQuery {
	if wq.driver.Dialect() == dialect.Postgres {
		wq.Unique(false)
	}
	wq.modifiers = append(wq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return wq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (wq *WorkflowQuery) ForShare(opts ...sql.LockOption) *WorkflowQuery {
	if wq.driver.Dialect() == dialect.Postgres {
		wq.Unique(false)
	}
	wq.modifiers = append(wq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return wq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wq *WorkflowQuery) Modify(modifiers ...func(s *sql.Selector)) *WorkflowSelect {
	wq.modifiers = append(wq.modifiers, modifiers...)
	return wq.Select()
}

// WorkflowGroupBy is the group-by builder for Workflow entities.
type WorkflowGroupBy struct {
	selector
	build *WorkflowQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WorkflowGroupBy) Aggregate(fns ...AggregateFunc) *WorkflowGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the selector query and scans the result into the given value.
func (wgb *WorkflowGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wgb.build.ctx, "GroupBy")
	if err := wgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkflowQuery, *WorkflowGroupBy](ctx, wgb.build, wgb, wgb.build.inters, v)
}

func (wgb *WorkflowGroupBy) sqlScan(ctx context.Context, root *WorkflowQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wgb.flds)+len(wgb.fns))
		for _, f := range *wgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WorkflowSelect is the builder for selecting fields of Workflow entities.
type WorkflowSelect struct {
	*WorkflowQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ws *WorkflowSelect) Aggregate(fns ...AggregateFunc) *WorkflowSelect {
	ws.fns = append(ws.fns, fns...)
	return ws
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WorkflowSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ws.ctx, "Select")
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkflowQuery, *WorkflowSelect](ctx, ws.WorkflowQuery, ws, ws.inters, v)
}

func (ws *WorkflowSelect) sqlScan(ctx context.Context, root *WorkflowQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ws.fns))
	for _, fn := range ws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ws *WorkflowSelect) Modify(modifiers ...func(s *sql.Selector)) *WorkflowSelect {
	ws.modifiers = append(ws.modifiers, modifiers...)
	return ws
}
