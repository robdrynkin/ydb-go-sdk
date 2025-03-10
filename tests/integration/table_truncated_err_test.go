//go:build integration
// +build integration

package integration

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/retry"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/options"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

// https://github.com/ydb-platform/ydb-go-sdk/issues/798
func TestIssue798TruncatedError(t *testing.T) {
	const rowsLimit = 1000

	var (
		scope     = newScope(t)
		driver    = scope.Driver(ydb.WithExecuteDataQueryOverQueryClient(false))
		db        = scope.SQLDriver(ydb.WithQueryService(false))
		tablePath = scope.TablePath()
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// clear table
	{
		driver.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
			_, _, err := s.Execute(ctx,
				table.DefaultTxControl(),
				fmt.Sprintf("DELETE FROM `%s`;", tablePath),
				nil,
			)
			if err != nil {
				return err
			}

			return err
		}, table.WithIdempotent())
	}

	// upsert rows
	{
		rows := make([]types.Value, rowsLimit)
		for i := range rows {
			rows[i] = types.StructValue(
				types.StructFieldValue("id", types.Int64Value(int64(i))),
				types.StructFieldValue("val", types.TextValue(strconv.Itoa(i))),
			)
		}
		err := driver.Table().BulkUpsert(ctx, tablePath,
			table.BulkUpsertDataRows(types.ListValue(rows...)), table.WithIdempotent(),
		)
		scope.Require.NoError(err)
	}

	// select rows without truncated error
	{
		err := driver.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
			_, results, err := s.Execute(ctx,
				table.DefaultTxControl(),
				fmt.Sprintf("SELECT * FROM `%s`;", tablePath),
				nil,
			)
			if err != nil {
				return err
			}
			if err = results.NextResultSetErr(ctx); err != nil {
				return fmt.Errorf("no result sets: %w", err)
			}
			if results.CurrentResultSet().RowCount() != rowsLimit {
				return fmt.Errorf("unexpected rows count: %d", results.CurrentResultSet().RowCount())
			}

			return results.Err()
		}, table.WithIdempotent())
		scope.Require.NoError(err)

		err = retry.Do(ctx, db, func(ctx context.Context, cc *sql.Conn) error {
			rows, err := cc.QueryContext(ctx, fmt.Sprintf("SELECT * FROM `%s`;", tablePath))
			if err != nil {
				return err
			}
			defer func() {
				_ = rows.Close()
			}()
			count := 0
			for rows.Next() {
				count++
			}
			if count != rowsLimit {
				return fmt.Errorf("unexpected rows count: %d", count)
			}

			return rows.Err()
		}, retry.WithIdempotent(true))
		scope.Require.NoError(err)
	}

	// upsert 1 row for get 1001 rows and truncated error
	{
		err := driver.Table().BulkUpsert(ctx, tablePath,
			table.BulkUpsertDataRows(types.ListValue(types.StructValue(
				types.StructFieldValue("id", types.Int64Value(rowsLimit)),
				types.StructFieldValue("val", types.TextValue(strconv.Itoa(rowsLimit))),
			))),
			table.WithIdempotent(),
		)
		scope.Require.NoError(err)
	}

	// select all rows with truncated result error
	{
		{
			err := driver.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
				_, results, err := s.Execute(ctx,
					table.DefaultTxControl(),
					fmt.Sprintf("SELECT * FROM `%s`;", tablePath),
					nil,
				)
				if err != nil {
					return err
				}
				if err = results.NextResultSetErr(ctx); err != nil {
					return fmt.Errorf("no result sets: %w", err)
				}

				rowsCount := results.CurrentResultSet().RowCount()

				if rowsCount != rowsLimit {
					return fmt.Errorf("unexpected rows count: %d", rowsCount)
				}

				return results.Err() // expected truncated error for execute data query using table client
			}, table.WithIdempotent())
			scope.Require.ErrorIs(err, result.ErrTruncated)
		}

		{
			var rowsCount int
			err := retry.Do(ctx, db, func(ctx context.Context, cc *sql.Conn) error {
				rowsCount = 0
				rows, err := cc.QueryContext(ctx, fmt.Sprintf("SELECT * FROM `%s`;", tablePath))
				if err != nil {
					return err
				}
				defer func() {
					_ = rows.Close()
				}()
				for rows.Next() {
					rowsCount++
				}

				if err := rows.Err(); err != nil {
					return err
				}

				if rowsCount != rowsLimit {
					return fmt.Errorf("unexpected rows count: %d", rowsCount)
				}

				return nil
			}, retry.WithIdempotent(true))
			scope.Require.ErrorIs(err, result.ErrTruncated)
		}
	}

	// select all rows without truncated result error
	{
		err := driver.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
			_, results, err := s.Execute(ctx,
				table.DefaultTxControl(),
				fmt.Sprintf("SELECT * FROM `%s`;", tablePath),
				nil,
				options.WithIgnoreTruncated(),
			)
			if err != nil {
				return err
			}
			if err = results.NextResultSetErr(ctx); err != nil {
				return fmt.Errorf("no result sets: %w", err)
			}
			if rowsCount := results.CurrentResultSet().RowCount(); rowsCount != rowsLimit {
				return fmt.Errorf("unexpected rows count: %d", rowsCount)
			}

			return nil
		}, table.WithIdempotent())
		scope.Require.NoError(err)
	}

	// connect with default option ignore truncated without truncated result error
	{
		driver, err := driver.With(ctx, ydb.WithIgnoreTruncated())
		scope.Require.NoError(err)

		err = driver.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
			_, results, err := s.Execute(ctx,
				table.DefaultTxControl(),
				fmt.Sprintf("SELECT * FROM `%s`;", tablePath),
				nil,
			)
			if err != nil {
				return err
			}
			if err = results.NextResultSetErr(ctx); err != nil {
				return fmt.Errorf("no result sets: %w", err)
			}
			if results.CurrentResultSet().RowCount() != rowsLimit {
				return fmt.Errorf("unexpected rows count: %d", results.CurrentResultSet().RowCount())
			}
			return results.Err() // expected nil
		}, table.WithIdempotent())
		scope.Require.NoError(err)

		db = sql.OpenDB(ydb.MustConnector(driver))
		err = retry.Do(ctx, db, func(ctx context.Context, cc *sql.Conn) error {
			rows, err := cc.QueryContext(ctx, fmt.Sprintf("SELECT * FROM `%s`;", tablePath))
			if err != nil {
				return err
			}
			defer func() {
				_ = rows.Close()
			}()
			count := 0
			for rows.Next() {
				count++
			}
			return rows.Err()
		}, retry.WithIdempotent(true))
		scope.Require.NoError(err)
	}
}

func TestIssue798NoTruncatedErrorOverQueryService(t *testing.T) {
	const rowsLimit = 1000
	var (
		scope     = newScope(t)
		driver    = scope.Driver(ydb.WithExecuteDataQueryOverQueryClient(true))
		db        = scope.SQLDriver(ydb.WithQueryService(true))
		tablePath = scope.TablePath()
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// clear table
	{
		driver.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
			_, _, err := s.Execute(ctx,
				table.DefaultTxControl(),
				fmt.Sprintf("DELETE FROM `%s`;", tablePath),
				nil,
			)
			if err != nil {
				return err
			}

			return err
		}, table.WithIdempotent())
	}

	// upsert rows
	{
		rows := make([]types.Value, rowsLimit)
		for i := range rows {
			rows[i] = types.StructValue(
				types.StructFieldValue("id", types.Int64Value(int64(i))),
				types.StructFieldValue("val", types.TextValue(strconv.Itoa(i))),
			)
		}
		err := driver.Table().BulkUpsert(ctx, tablePath,
			table.BulkUpsertDataRows(types.ListValue(rows...)), table.WithIdempotent(),
		)
		scope.Require.NoError(err)
	}

	// select rows without truncated error
	{
		err := driver.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
			_, results, err := s.Execute(ctx,
				table.DefaultTxControl(),
				fmt.Sprintf("SELECT * FROM `%s`;", tablePath),
				nil,
			)
			if err != nil {
				return err
			}
			if err = results.NextResultSetErr(ctx); err != nil {
				return fmt.Errorf("no result sets: %w", err)
			}
			if results.CurrentResultSet().RowCount() != rowsLimit {
				return fmt.Errorf("unexpected rows count: %d", results.CurrentResultSet().RowCount())
			}

			return results.Err()
		}, table.WithIdempotent())
		scope.Require.NoError(err)

		err = retry.Do(ctx, db, func(ctx context.Context, cc *sql.Conn) error {
			rows, err := cc.QueryContext(ctx, fmt.Sprintf("SELECT * FROM `%s`;", tablePath))
			if err != nil {
				return err
			}
			defer func() {
				_ = rows.Close()
			}()
			count := 0
			for rows.Next() {
				count++
			}
			if count != rowsLimit {
				return fmt.Errorf("unexpected rows count: %d", count)
			}

			return rows.Err()
		}, retry.WithIdempotent(true))
		scope.Require.NoError(err)
	}

	// upsert 1 row for get 1001 rows and truncated error
	{
		err := driver.Table().BulkUpsert(ctx, tablePath,
			table.BulkUpsertDataRows(types.ListValue(types.StructValue(
				types.StructFieldValue("id", types.Int64Value(rowsLimit)),
				types.StructFieldValue("val", types.TextValue(strconv.Itoa(rowsLimit))),
			))),
			table.WithIdempotent(),
		)
		scope.Require.NoError(err)
	}

	// select all rows with truncated result error
	{
		{
			err := driver.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
				_, results, err := s.Execute(ctx,
					table.DefaultTxControl(),
					fmt.Sprintf("SELECT * FROM `%s`;", tablePath),
					nil,
				)
				if err != nil {
					return err
				}
				if err = results.NextResultSetErr(ctx); err != nil {
					return fmt.Errorf("no result sets: %w", err)
				}

				rowsCount := results.CurrentResultSet().RowCount()
				if rowsCount != rowsLimit+1 {
					return fmt.Errorf("unexpected rows count: %d", rowsCount)
				}

				return results.Err() // expected truncated error for execute data query using table client
			}, table.WithIdempotent())
			scope.Require.NoError(err)
		}

		{
			var rowsCount int
			err := retry.Do(ctx, db, func(ctx context.Context, cc *sql.Conn) error {
				rowsCount = 0
				rows, err := cc.QueryContext(ctx, fmt.Sprintf("SELECT * FROM `%s`;", tablePath))
				if err != nil {
					return err
				}
				defer func() {
					_ = rows.Close()
				}()
				for rows.Next() {
					rowsCount++
				}
				if rowsCount != rowsLimit+1 {
					return fmt.Errorf("unexpected rows count: %d", rowsCount)
				}

				return rows.Err()
			}, retry.WithIdempotent(true))
			scope.Require.NoError(err)
		}
	}

	// select all rows without truncated result error
	{
		err := driver.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
			_, results, err := s.Execute(ctx,
				table.DefaultTxControl(),
				fmt.Sprintf("SELECT * FROM `%s`;", tablePath),
				nil,
			)
			if err != nil {
				return err
			}
			if err = results.NextResultSetErr(ctx); err != nil {
				return fmt.Errorf("no result sets: %w", err)
			}
			if results.CurrentResultSet().RowCount() != rowsLimit+1 {
				return fmt.Errorf("unexpected rows count: %d", results.CurrentResultSet().RowCount())
			}
			return results.Err() // expected nil
		}, table.WithIdempotent())
		scope.Require.NoError(err)
	}

	// connect with default option ignore truncated without truncated result error
	{
		driver, err := driver.With(ctx, ydb.WithIgnoreTruncated())
		scope.Require.NoError(err)

		err = driver.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
			_, results, err := s.Execute(ctx,
				table.DefaultTxControl(),
				fmt.Sprintf("SELECT * FROM `%s`;", tablePath),
				nil,
			)
			if err != nil {
				return err
			}
			if err = results.NextResultSetErr(ctx); err != nil {
				return fmt.Errorf("no result sets: %w", err)
			}
			if results.CurrentResultSet().RowCount() != rowsLimit+1 {
				return fmt.Errorf("unexpected rows count: %d", results.CurrentResultSet().RowCount())
			}
			return results.Err() // expected nil
		}, table.WithIdempotent())
		scope.Require.NoError(err)

		db = sql.OpenDB(ydb.MustConnector(driver))
		err = retry.Do(ctx, db, func(ctx context.Context, cc *sql.Conn) error {
			rows, err := cc.QueryContext(ctx, fmt.Sprintf("SELECT * FROM `%s`;", tablePath))
			if err != nil {
				return err
			}
			defer func() {
				_ = rows.Close()
			}()
			count := 0
			for rows.Next() {
				count++
			}
			return rows.Err()
		}, retry.WithIdempotent(true))
		scope.Require.NoError(err)
	}
}
