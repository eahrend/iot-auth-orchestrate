// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testWorkspaces(t *testing.T) {
	t.Parallel()

	query := Workspaces()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testWorkspacesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Workspace{}
	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Workspaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWorkspacesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Workspace{}
	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Workspaces().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Workspaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWorkspacesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Workspace{}
	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WorkspaceSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Workspaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWorkspacesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Workspace{}
	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := WorkspaceExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Workspace exists: %s", err)
	}
	if !e {
		t.Errorf("Expected WorkspaceExists to return true, but got false.")
	}
}

func testWorkspacesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Workspace{}
	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	workspaceFound, err := FindWorkspace(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if workspaceFound == nil {
		t.Error("want a record, got nil")
	}
}

func testWorkspacesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Workspace{}
	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Workspaces().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testWorkspacesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Workspace{}
	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Workspaces().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testWorkspacesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	workspaceOne := &Workspace{}
	workspaceTwo := &Workspace{}
	if err = randomize.Struct(seed, workspaceOne, workspaceDBTypes, false, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}
	if err = randomize.Struct(seed, workspaceTwo, workspaceDBTypes, false, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = workspaceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = workspaceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Workspaces().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testWorkspacesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	workspaceOne := &Workspace{}
	workspaceTwo := &Workspace{}
	if err = randomize.Struct(seed, workspaceOne, workspaceDBTypes, false, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}
	if err = randomize.Struct(seed, workspaceTwo, workspaceDBTypes, false, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = workspaceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = workspaceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Workspaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func workspaceBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Workspace) error {
	*o = Workspace{}
	return nil
}

func workspaceAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Workspace) error {
	*o = Workspace{}
	return nil
}

func workspaceAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Workspace) error {
	*o = Workspace{}
	return nil
}

func workspaceBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Workspace) error {
	*o = Workspace{}
	return nil
}

func workspaceAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Workspace) error {
	*o = Workspace{}
	return nil
}

func workspaceBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Workspace) error {
	*o = Workspace{}
	return nil
}

func workspaceAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Workspace) error {
	*o = Workspace{}
	return nil
}

func workspaceBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Workspace) error {
	*o = Workspace{}
	return nil
}

func workspaceAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Workspace) error {
	*o = Workspace{}
	return nil
}

func testWorkspacesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Workspace{}
	o := &Workspace{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, workspaceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Workspace object: %s", err)
	}

	AddWorkspaceHook(boil.BeforeInsertHook, workspaceBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	workspaceBeforeInsertHooks = []WorkspaceHook{}

	AddWorkspaceHook(boil.AfterInsertHook, workspaceAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	workspaceAfterInsertHooks = []WorkspaceHook{}

	AddWorkspaceHook(boil.AfterSelectHook, workspaceAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	workspaceAfterSelectHooks = []WorkspaceHook{}

	AddWorkspaceHook(boil.BeforeUpdateHook, workspaceBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	workspaceBeforeUpdateHooks = []WorkspaceHook{}

	AddWorkspaceHook(boil.AfterUpdateHook, workspaceAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	workspaceAfterUpdateHooks = []WorkspaceHook{}

	AddWorkspaceHook(boil.BeforeDeleteHook, workspaceBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	workspaceBeforeDeleteHooks = []WorkspaceHook{}

	AddWorkspaceHook(boil.AfterDeleteHook, workspaceAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	workspaceAfterDeleteHooks = []WorkspaceHook{}

	AddWorkspaceHook(boil.BeforeUpsertHook, workspaceBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	workspaceBeforeUpsertHooks = []WorkspaceHook{}

	AddWorkspaceHook(boil.AfterUpsertHook, workspaceAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	workspaceAfterUpsertHooks = []WorkspaceHook{}
}

func testWorkspacesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Workspace{}
	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Workspaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWorkspacesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Workspace{}
	if err = randomize.Struct(seed, o, workspaceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(workspaceColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Workspaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWorkspaceToManyRoles(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Workspace
	var b, c Role

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.WorkspaceID = a.ID
	c.WorkspaceID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Roles().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.WorkspaceID == b.WorkspaceID {
			bFound = true
		}
		if v.WorkspaceID == c.WorkspaceID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := WorkspaceSlice{&a}
	if err = a.L.LoadRoles(ctx, tx, false, (*[]*Workspace)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Roles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Roles = nil
	if err = a.L.LoadRoles(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Roles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testWorkspaceToManyUserWorkspaces(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Workspace
	var b, c UserWorkspace

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userWorkspaceDBTypes, false, userWorkspaceColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userWorkspaceDBTypes, false, userWorkspaceColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.WorkspaceID = a.ID
	c.WorkspaceID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserWorkspaces().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.WorkspaceID == b.WorkspaceID {
			bFound = true
		}
		if v.WorkspaceID == c.WorkspaceID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := WorkspaceSlice{&a}
	if err = a.L.LoadUserWorkspaces(ctx, tx, false, (*[]*Workspace)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserWorkspaces); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserWorkspaces = nil
	if err = a.L.LoadUserWorkspaces(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserWorkspaces); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testWorkspaceToManyWorkspaceDevices(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Workspace
	var b, c WorkspaceDevice

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, workspaceDeviceDBTypes, false, workspaceDeviceColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, workspaceDeviceDBTypes, false, workspaceDeviceColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.WorkspaceID = a.ID
	c.WorkspaceID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.WorkspaceDevices().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.WorkspaceID == b.WorkspaceID {
			bFound = true
		}
		if v.WorkspaceID == c.WorkspaceID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := WorkspaceSlice{&a}
	if err = a.L.LoadWorkspaceDevices(ctx, tx, false, (*[]*Workspace)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.WorkspaceDevices); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.WorkspaceDevices = nil
	if err = a.L.LoadWorkspaceDevices(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.WorkspaceDevices); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testWorkspaceToManyAddOpRoles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Workspace
	var b, c, d, e Role

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, workspaceDBTypes, false, strmangle.SetComplement(workspacePrimaryKeyColumns, workspaceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Role{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Role{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRoles(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.WorkspaceID {
			t.Error("foreign key was wrong value", a.ID, first.WorkspaceID)
		}
		if a.ID != second.WorkspaceID {
			t.Error("foreign key was wrong value", a.ID, second.WorkspaceID)
		}

		if first.R.Workspace != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Workspace != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Roles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Roles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Roles().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testWorkspaceToManyAddOpUserWorkspaces(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Workspace
	var b, c, d, e UserWorkspace

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, workspaceDBTypes, false, strmangle.SetComplement(workspacePrimaryKeyColumns, workspaceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserWorkspace{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userWorkspaceDBTypes, false, strmangle.SetComplement(userWorkspacePrimaryKeyColumns, userWorkspaceColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*UserWorkspace{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserWorkspaces(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.WorkspaceID {
			t.Error("foreign key was wrong value", a.ID, first.WorkspaceID)
		}
		if a.ID != second.WorkspaceID {
			t.Error("foreign key was wrong value", a.ID, second.WorkspaceID)
		}

		if first.R.Workspace != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Workspace != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserWorkspaces[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserWorkspaces[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserWorkspaces().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testWorkspaceToManyAddOpWorkspaceDevices(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Workspace
	var b, c, d, e WorkspaceDevice

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, workspaceDBTypes, false, strmangle.SetComplement(workspacePrimaryKeyColumns, workspaceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*WorkspaceDevice{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, workspaceDeviceDBTypes, false, strmangle.SetComplement(workspaceDevicePrimaryKeyColumns, workspaceDeviceColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*WorkspaceDevice{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddWorkspaceDevices(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.WorkspaceID {
			t.Error("foreign key was wrong value", a.ID, first.WorkspaceID)
		}
		if a.ID != second.WorkspaceID {
			t.Error("foreign key was wrong value", a.ID, second.WorkspaceID)
		}

		if first.R.Workspace != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Workspace != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.WorkspaceDevices[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.WorkspaceDevices[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.WorkspaceDevices().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testWorkspacesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Workspace{}
	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWorkspacesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Workspace{}
	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WorkspaceSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWorkspacesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Workspace{}
	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Workspaces().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	workspaceDBTypes = map[string]string{`ID`: `int`, `WorkspaceName`: `varchar`}
	_                = bytes.MinRead
)

func testWorkspacesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(workspacePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(workspaceAllColumns) == len(workspacePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Workspace{}
	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Workspaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspacePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testWorkspacesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(workspaceAllColumns) == len(workspacePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Workspace{}
	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Workspaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, workspaceDBTypes, true, workspacePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(workspaceAllColumns, workspacePrimaryKeyColumns) {
		fields = workspaceAllColumns
	} else {
		fields = strmangle.SetComplement(
			workspaceAllColumns,
			workspacePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := WorkspaceSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testWorkspacesUpsert(t *testing.T) {
	t.Parallel()

	if len(workspaceAllColumns) == len(workspacePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLWorkspaceUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Workspace{}
	if err = randomize.Struct(seed, &o, workspaceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Workspace: %s", err)
	}

	count, err := Workspaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, workspaceDBTypes, false, workspacePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Workspace struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Workspace: %s", err)
	}

	count, err = Workspaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
