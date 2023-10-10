package scylladb

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type queryBuilder[T any] struct {
	model   ITable
	session ISessionx
}

func (queryBuilder *queryBuilder[T]) Insert(context context.Context, insertData *T) error {
	insertStatement, insertNames := queryBuilder.model.Insert()
	insertQuery := queryBuilder.session.Query(insertStatement, insertNames)
	err := insertQuery.BindStruct(insertData).ExecRelease()
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	return nil
}

func (queryBuilder *queryBuilder[T]) Select(context context.Context, dataToGet *T) ([]T, error) {
	selectStatement, selectNames := queryBuilder.model.Select()
	selectQuery := queryBuilder.session.Query(selectStatement, selectNames)

	var results []T
	err := selectQuery.BindStruct(dataToGet).SelectRelease(&results)
	if err != nil {
		logrus.Error("Select error:", err.Error())
		return nil, err
	}

	return results, nil

}

func (queryBuilder *queryBuilder[T]) Get(ctx context.Context, dataToGet *T) (*T, error) {
	selectStatement, selectNames := queryBuilder.model.Get()
	selectQuery := queryBuilder.session.Query(selectStatement, selectNames)

	var result []T
	err := selectQuery.BindStruct(dataToGet).WithContext(ctx).SelectRelease(&result)
	if err != nil {
		logrus.Error("Get error", err.Error())
		return nil, err
	}

	if len(result) > 0 {
		return &result[0], nil
	}

	return nil, nil
}

func (queryBuilder *queryBuilder[T]) Delete(ctx context.Context, dataToBeDeleted *T)(error){
	deleteStatment, deleteName := queryBuilder.model.Delete()
	fmt.Print(deleteName, deleteStatment)
	deleteQuery := queryBuilder.session.Query(deleteStatment, deleteName)

	if err := deleteQuery.BindStruct(dataToBeDeleted).WithContext(ctx).ExecRelease(); err != nil{
		logrus.Error("Delete error", err.Error())
		return err

	}

	return nil

}

/*
	It will everything from table.

SELECT * FROM table;
*/
func (queryBuilder *queryBuilder[T]) SelectAll(ctx context.Context) ([]T, error) {
	selectAllStatement, selectAllNames := queryBuilder.model.SelectAll()
	fmt.Println(selectAllNames, selectAllNames)
	selectAllQuery := queryBuilder.session.Query(selectAllStatement, selectAllNames)

	var results []T
	err := selectAllQuery.WithContext(ctx).SelectRelease(&results)
	if err != nil {
		logrus.Error("SelectAll error", err.Error())
		return nil, err
	}

	return results, nil
}

func NewQueryBuilder[T any](model ITable, session ISessionx) *queryBuilder[T] {
	return &queryBuilder[T]{
		model:   model,
		session: session,
	}

}
