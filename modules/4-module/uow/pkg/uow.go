package uow

import (
	"context"
	"database/sql"
	"fmt"
)

/*
   This package provides a Unit of Work (UoW) pattern implementation for managing transactions

   Problem:
   When there are multiple repositories operating with a database, there might occur inconsistency in datas.
   - CreateCategoryRepo
   - CreateCourseRepo
   Insertion for Category succeeds, but fails at Course insertion.

   Unit Of Work pattern instantiates repositories with transaction capability.
*/

type RepositoryFactory func(tx *sql.Tx) any

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (any, error)
	Do(ctx context.Context, fn func(uow *Uow) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}

type Uow struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(ctx context.Context, db *sql.DB) *Uow {
	return &Uow{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *Uow) Register(name string, fc RepositoryFactory) {
	u.Repositories[name] = fc
}

func (u *Uow) UnRegister(name string) {
	delete(u.Repositories, name)
}

func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.Tx = tx
	}

	repo := u.Repositories[name](u.Tx)
	return repo, nil
}

func (u *Uow) Do(ctx context.Context, fn func(uow *Uow) error) error {
	// Initialize transaction
	if u.Tx != nil {
		return fmt.Errorf("transaction already started")
	}
	tx, err := u.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.Tx = tx

	// Execute Unit of Work
	err = fn(u)
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return fmt.Errorf("error of uow: %s Error of rollback: %s", err, errRb)
		}
		return err
	}

	return u.CommitOrRollback()
}

func (u *Uow) CommitOrRollback() error {
	err := u.Tx.Commit()
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err, errRb)
		}
	}

	u.Tx = nil
	return nil
}

func (u *Uow) Rollback() error {
	if u.Tx == nil {
		return fmt.Errorf("no transaction to rollback")
	}

	err := u.Tx.Rollback()
	if err != nil {
		return err
	}

	u.Tx = nil
	return nil
}
