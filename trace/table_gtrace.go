// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
	"time"
)

// Compose returns a new Table which has functional fields composed
// both from t and x.
func (t Table) Compose(x Table) (ret Table) {
	switch {
	case t.OnCreateSession == nil:
		ret.OnCreateSession = x.OnCreateSession
	case x.OnCreateSession == nil:
		ret.OnCreateSession = t.OnCreateSession
	default:
		h1 := t.OnCreateSession
		h2 := x.OnCreateSession
		ret.OnCreateSession = func(c CreateSessionStartInfo) func(CreateSessionDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c CreateSessionDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnKeepAlive == nil:
		ret.OnKeepAlive = x.OnKeepAlive
	case x.OnKeepAlive == nil:
		ret.OnKeepAlive = t.OnKeepAlive
	default:
		h1 := t.OnKeepAlive
		h2 := x.OnKeepAlive
		ret.OnKeepAlive = func(k KeepAliveStartInfo) func(KeepAliveDoneInfo) {
			r1 := h1(k)
			r2 := h2(k)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(k KeepAliveDoneInfo) {
					r1(k)
					r2(k)
				}
			}
		}
	}
	switch {
	case t.OnDeleteSession == nil:
		ret.OnDeleteSession = x.OnDeleteSession
	case x.OnDeleteSession == nil:
		ret.OnDeleteSession = t.OnDeleteSession
	default:
		h1 := t.OnDeleteSession
		h2 := x.OnDeleteSession
		ret.OnDeleteSession = func(d DeleteSessionStartInfo) func(DeleteSessionDoneInfo) {
			r1 := h1(d)
			r2 := h2(d)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(d DeleteSessionDoneInfo) {
					r1(d)
					r2(d)
				}
			}
		}
	}
	switch {
	case t.OnPrepareDataQuery == nil:
		ret.OnPrepareDataQuery = x.OnPrepareDataQuery
	case x.OnPrepareDataQuery == nil:
		ret.OnPrepareDataQuery = t.OnPrepareDataQuery
	default:
		h1 := t.OnPrepareDataQuery
		h2 := x.OnPrepareDataQuery
		ret.OnPrepareDataQuery = func(p PrepareDataQueryStartInfo) func(PrepareDataQueryDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PrepareDataQueryDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnExecuteDataQuery == nil:
		ret.OnExecuteDataQuery = x.OnExecuteDataQuery
	case x.OnExecuteDataQuery == nil:
		ret.OnExecuteDataQuery = t.OnExecuteDataQuery
	default:
		h1 := t.OnExecuteDataQuery
		h2 := x.OnExecuteDataQuery
		ret.OnExecuteDataQuery = func(e ExecuteDataQueryStartInfo) func(ExecuteDataQueryDoneInfo) {
			r1 := h1(e)
			r2 := h2(e)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(e ExecuteDataQueryDoneInfo) {
					r1(e)
					r2(e)
				}
			}
		}
	}
	switch {
	case t.OnStreamReadTable == nil:
		ret.OnStreamReadTable = x.OnStreamReadTable
	case x.OnStreamReadTable == nil:
		ret.OnStreamReadTable = t.OnStreamReadTable
	default:
		h1 := t.OnStreamReadTable
		h2 := x.OnStreamReadTable
		ret.OnStreamReadTable = func(s StreamReadTableStartInfo) func(StreamReadTableDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s StreamReadTableDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnStreamExecuteScanQuery == nil:
		ret.OnStreamExecuteScanQuery = x.OnStreamExecuteScanQuery
	case x.OnStreamExecuteScanQuery == nil:
		ret.OnStreamExecuteScanQuery = t.OnStreamExecuteScanQuery
	default:
		h1 := t.OnStreamExecuteScanQuery
		h2 := x.OnStreamExecuteScanQuery
		ret.OnStreamExecuteScanQuery = func(s StreamExecuteScanQueryStartInfo) func(StreamExecuteScanQueryDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s StreamExecuteScanQueryDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnBeginTransaction == nil:
		ret.OnBeginTransaction = x.OnBeginTransaction
	case x.OnBeginTransaction == nil:
		ret.OnBeginTransaction = t.OnBeginTransaction
	default:
		h1 := t.OnBeginTransaction
		h2 := x.OnBeginTransaction
		ret.OnBeginTransaction = func(b BeginTransactionStartInfo) func(BeginTransactionDoneInfo) {
			r1 := h1(b)
			r2 := h2(b)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(b BeginTransactionDoneInfo) {
					r1(b)
					r2(b)
				}
			}
		}
	}
	switch {
	case t.OnCommitTransaction == nil:
		ret.OnCommitTransaction = x.OnCommitTransaction
	case x.OnCommitTransaction == nil:
		ret.OnCommitTransaction = t.OnCommitTransaction
	default:
		h1 := t.OnCommitTransaction
		h2 := x.OnCommitTransaction
		ret.OnCommitTransaction = func(c CommitTransactionStartInfo) func(CommitTransactionDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c CommitTransactionDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnRollbackTransaction == nil:
		ret.OnRollbackTransaction = x.OnRollbackTransaction
	case x.OnRollbackTransaction == nil:
		ret.OnRollbackTransaction = t.OnRollbackTransaction
	default:
		h1 := t.OnRollbackTransaction
		h2 := x.OnRollbackTransaction
		ret.OnRollbackTransaction = func(r RollbackTransactionStartInfo) func(RollbackTransactionDoneInfo) {
			r1 := h1(r)
			r2 := h2(r)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(r RollbackTransactionDoneInfo) {
					r1(r)
					r2(r)
				}
			}
		}
	}
	switch {
	case t.OnPoolCreate == nil:
		ret.OnPoolCreate = x.OnPoolCreate
	case x.OnPoolCreate == nil:
		ret.OnPoolCreate = t.OnPoolCreate
	default:
		h1 := t.OnPoolCreate
		h2 := x.OnPoolCreate
		ret.OnPoolCreate = func(p PoolCreateStartInfo) func(PoolCreateDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolCreateDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnPoolGet == nil:
		ret.OnPoolGet = x.OnPoolGet
	case x.OnPoolGet == nil:
		ret.OnPoolGet = t.OnPoolGet
	default:
		h1 := t.OnPoolGet
		h2 := x.OnPoolGet
		ret.OnPoolGet = func(p PoolGetStartInfo) func(PoolGetDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolGetDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnPoolWait == nil:
		ret.OnPoolWait = x.OnPoolWait
	case x.OnPoolWait == nil:
		ret.OnPoolWait = t.OnPoolWait
	default:
		h1 := t.OnPoolWait
		h2 := x.OnPoolWait
		ret.OnPoolWait = func(p PoolWaitStartInfo) func(PoolWaitDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolWaitDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnPoolTake == nil:
		ret.OnPoolTake = x.OnPoolTake
	case x.OnPoolTake == nil:
		ret.OnPoolTake = t.OnPoolTake
	default:
		h1 := t.OnPoolTake
		h2 := x.OnPoolTake
		ret.OnPoolTake = func(p PoolTakeStartInfo) func(PoolTakeWaitInfo) func(PoolTakeDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolTakeWaitInfo) func(PoolTakeDoneInfo) {
					r11 := r1(p)
					r21 := r2(p)
					switch {
					case r11 == nil:
						return r21
					case r21 == nil:
						return r11
					default:
						return func(p PoolTakeDoneInfo) {
							r11(p)
							r21(p)
						}
					}
				}
			}
		}
	}
	switch {
	case t.OnPoolPut == nil:
		ret.OnPoolPut = x.OnPoolPut
	case x.OnPoolPut == nil:
		ret.OnPoolPut = t.OnPoolPut
	default:
		h1 := t.OnPoolPut
		h2 := x.OnPoolPut
		ret.OnPoolPut = func(p PoolPutStartInfo) func(PoolPutDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolPutDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnPoolCloseSession == nil:
		ret.OnPoolCloseSession = x.OnPoolCloseSession
	case x.OnPoolCloseSession == nil:
		ret.OnPoolCloseSession = t.OnPoolCloseSession
	default:
		h1 := t.OnPoolCloseSession
		h2 := x.OnPoolCloseSession
		ret.OnPoolCloseSession = func(p PoolCloseSessionStartInfo) func(PoolCloseSessionDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolCloseSessionDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnPoolClose == nil:
		ret.OnPoolClose = x.OnPoolClose
	case x.OnPoolClose == nil:
		ret.OnPoolClose = t.OnPoolClose
	default:
		h1 := t.OnPoolClose
		h2 := x.OnPoolClose
		ret.OnPoolClose = func(p PoolCloseStartInfo) func(PoolCloseDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolCloseDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	return ret
}
func (t Table) onCreateSession(c1 CreateSessionStartInfo) func(CreateSessionDoneInfo) {
	fn := t.OnCreateSession
	if fn == nil {
		return func(CreateSessionDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(CreateSessionDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onKeepAlive(k KeepAliveStartInfo) func(KeepAliveDoneInfo) {
	fn := t.OnKeepAlive
	if fn == nil {
		return func(KeepAliveDoneInfo) {
			return
		}
	}
	res := fn(k)
	if res == nil {
		return func(KeepAliveDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onDeleteSession(d DeleteSessionStartInfo) func(DeleteSessionDoneInfo) {
	fn := t.OnDeleteSession
	if fn == nil {
		return func(DeleteSessionDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DeleteSessionDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPrepareDataQuery(p PrepareDataQueryStartInfo) func(PrepareDataQueryDoneInfo) {
	fn := t.OnPrepareDataQuery
	if fn == nil {
		return func(PrepareDataQueryDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PrepareDataQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onExecuteDataQuery(e ExecuteDataQueryStartInfo) func(ExecuteDataQueryDoneInfo) {
	fn := t.OnExecuteDataQuery
	if fn == nil {
		return func(ExecuteDataQueryDoneInfo) {
			return
		}
	}
	res := fn(e)
	if res == nil {
		return func(ExecuteDataQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onStreamReadTable(s StreamReadTableStartInfo) func(StreamReadTableDoneInfo) {
	fn := t.OnStreamReadTable
	if fn == nil {
		return func(StreamReadTableDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(StreamReadTableDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onStreamExecuteScanQuery(s StreamExecuteScanQueryStartInfo) func(StreamExecuteScanQueryDoneInfo) {
	fn := t.OnStreamExecuteScanQuery
	if fn == nil {
		return func(StreamExecuteScanQueryDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(StreamExecuteScanQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onBeginTransaction(b BeginTransactionStartInfo) func(BeginTransactionDoneInfo) {
	fn := t.OnBeginTransaction
	if fn == nil {
		return func(BeginTransactionDoneInfo) {
			return
		}
	}
	res := fn(b)
	if res == nil {
		return func(BeginTransactionDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onCommitTransaction(c1 CommitTransactionStartInfo) func(CommitTransactionDoneInfo) {
	fn := t.OnCommitTransaction
	if fn == nil {
		return func(CommitTransactionDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(CommitTransactionDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onRollbackTransaction(r RollbackTransactionStartInfo) func(RollbackTransactionDoneInfo) {
	fn := t.OnRollbackTransaction
	if fn == nil {
		return func(RollbackTransactionDoneInfo) {
			return
		}
	}
	res := fn(r)
	if res == nil {
		return func(RollbackTransactionDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPoolCreate(p PoolCreateStartInfo) func(PoolCreateDoneInfo) {
	fn := t.OnPoolCreate
	if fn == nil {
		return func(PoolCreateDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolCreateDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPoolGet(p PoolGetStartInfo) func(PoolGetDoneInfo) {
	fn := t.OnPoolGet
	if fn == nil {
		return func(PoolGetDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolGetDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPoolWait(p PoolWaitStartInfo) func(PoolWaitDoneInfo) {
	fn := t.OnPoolWait
	if fn == nil {
		return func(PoolWaitDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolWaitDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPoolTake(p PoolTakeStartInfo) func(PoolTakeWaitInfo) func(PoolTakeDoneInfo) {
	fn := t.OnPoolTake
	if fn == nil {
		return func(PoolTakeWaitInfo) func(PoolTakeDoneInfo) {
			return func(PoolTakeDoneInfo) {
				return
			}
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolTakeWaitInfo) func(PoolTakeDoneInfo) {
			return func(PoolTakeDoneInfo) {
				return
			}
		}
	}
	return func(p PoolTakeWaitInfo) func(PoolTakeDoneInfo) {
		res := res(p)
		if res == nil {
			return func(PoolTakeDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t Table) onPoolPut(p PoolPutStartInfo) func(PoolPutDoneInfo) {
	fn := t.OnPoolPut
	if fn == nil {
		return func(PoolPutDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolPutDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPoolCloseSession(p PoolCloseSessionStartInfo) func(PoolCloseSessionDoneInfo) {
	fn := t.OnPoolCloseSession
	if fn == nil {
		return func(PoolCloseSessionDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolCloseSessionDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPoolClose(p PoolCloseStartInfo) func(PoolCloseDoneInfo) {
	fn := t.OnPoolClose
	if fn == nil {
		return func(PoolCloseDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolCloseDoneInfo) {
			return
		}
	}
	return res
}
func TableOnCreateSession(t Table, c context.Context) func(_ context.Context, sessionID string, endpoint string, latency time.Duration, _ error) {
	var p CreateSessionStartInfo
	p.Context = c
	res := t.onCreateSession(p)
	return func(c context.Context, sessionID string, endpoint string, latency time.Duration, e error) {
		var p CreateSessionDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.Endpoint = endpoint
		p.Latency = latency
		p.Error = e
		res(p)
	}
}
func TableOnKeepAlive(t Table, c context.Context, sessionID string) func(_ context.Context, sessionID string, _ SessionInfo, _ error) {
	var p KeepAliveStartInfo
	p.Context = c
	p.SessionID = sessionID
	res := t.onKeepAlive(p)
	return func(c context.Context, sessionID string, s SessionInfo, e error) {
		var p KeepAliveDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.SessionInfo = s
		p.Error = e
		res(p)
	}
}
func TableOnDeleteSession(t Table, c context.Context, sessionID string) func(_ context.Context, sessionID string, latency time.Duration, _ error) {
	var p DeleteSessionStartInfo
	p.Context = c
	p.SessionID = sessionID
	res := t.onDeleteSession(p)
	return func(c context.Context, sessionID string, latency time.Duration, e error) {
		var p DeleteSessionDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.Latency = latency
		p.Error = e
		res(p)
	}
}
func TableOnPrepareDataQuery(t Table, c context.Context, sessionID string, query string) func(_ context.Context, sessionID string, query string, result DataQuery, cached bool, _ error) {
	var p PrepareDataQueryStartInfo
	p.Context = c
	p.SessionID = sessionID
	p.Query = query
	res := t.onPrepareDataQuery(p)
	return func(c context.Context, sessionID string, query string, result DataQuery, cached bool, e error) {
		var p PrepareDataQueryDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.Query = query
		p.Result = result
		p.Cached = cached
		p.Error = e
		res(p)
	}
}
func TableOnExecuteDataQuery(t Table, c context.Context, sessionID string, txID string, query DataQuery, parameters QueryParameters) func(_ context.Context, sessionID string, txID string, query DataQuery, parameters QueryParameters, prepared bool, _ Result, _ error) {
	var p ExecuteDataQueryStartInfo
	p.Context = c
	p.SessionID = sessionID
	p.TxID = txID
	p.Query = query
	p.Parameters = parameters
	res := t.onExecuteDataQuery(p)
	return func(c context.Context, sessionID string, txID string, query DataQuery, parameters QueryParameters, prepared bool, r Result, e error) {
		var p ExecuteDataQueryDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.TxID = txID
		p.Query = query
		p.Parameters = parameters
		p.Prepared = prepared
		p.Result = r
		p.Error = e
		res(p)
	}
}
func TableOnStreamReadTable(t Table, c context.Context, sessionID string) func(_ context.Context, sessionID string, _ Result, _ error) {
	var p StreamReadTableStartInfo
	p.Context = c
	p.SessionID = sessionID
	res := t.onStreamReadTable(p)
	return func(c context.Context, sessionID string, r Result, e error) {
		var p StreamReadTableDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.Result = r
		p.Error = e
		res(p)
	}
}
func TableOnStreamExecuteScanQuery(t Table, c context.Context, sessionID string, query DataQuery, parameters QueryParameters) func(_ context.Context, sessionID string, query DataQuery, parameters QueryParameters, _ Result, _ error) {
	var p StreamExecuteScanQueryStartInfo
	p.Context = c
	p.SessionID = sessionID
	p.Query = query
	p.Parameters = parameters
	res := t.onStreamExecuteScanQuery(p)
	return func(c context.Context, sessionID string, query DataQuery, parameters QueryParameters, r Result, e error) {
		var p StreamExecuteScanQueryDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.Query = query
		p.Parameters = parameters
		p.Result = r
		p.Error = e
		res(p)
	}
}
func TableOnBeginTransaction(t Table, c context.Context, sessionID string) func(_ context.Context, sessionID string, txID string, _ error) {
	var p BeginTransactionStartInfo
	p.Context = c
	p.SessionID = sessionID
	res := t.onBeginTransaction(p)
	return func(c context.Context, sessionID string, txID string, e error) {
		var p BeginTransactionDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.TxID = txID
		p.Error = e
		res(p)
	}
}
func TableOnCommitTransaction(t Table, c context.Context, sessionID string, txID string) func(_ context.Context, sessionID string, txID string, _ error) {
	var p CommitTransactionStartInfo
	p.Context = c
	p.SessionID = sessionID
	p.TxID = txID
	res := t.onCommitTransaction(p)
	return func(c context.Context, sessionID string, txID string, e error) {
		var p CommitTransactionDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.TxID = txID
		p.Error = e
		res(p)
	}
}
func TableOnRollbackTransaction(t Table, c context.Context, sessionID string, txID string) func(_ context.Context, sessionID string, txID string, _ error) {
	var p RollbackTransactionStartInfo
	p.Context = c
	p.SessionID = sessionID
	p.TxID = txID
	res := t.onRollbackTransaction(p)
	return func(c context.Context, sessionID string, txID string, e error) {
		var p RollbackTransactionDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.TxID = txID
		p.Error = e
		res(p)
	}
}
func TableOnPoolCreate(t Table, c context.Context) func(_ context.Context, sessionID string, _ error) {
	var p PoolCreateStartInfo
	p.Context = c
	res := t.onPoolCreate(p)
	return func(c context.Context, sessionID string, e error) {
		var p PoolCreateDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.Error = e
		res(p)
	}
}
func TableOnPoolGet(t Table, c context.Context) func(_ context.Context, sessionID string, latency time.Duration, retryAttempts int, _ error) {
	var p PoolGetStartInfo
	p.Context = c
	res := t.onPoolGet(p)
	return func(c context.Context, sessionID string, latency time.Duration, retryAttempts int, e error) {
		var p PoolGetDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.Latency = latency
		p.RetryAttempts = retryAttempts
		p.Error = e
		res(p)
	}
}
func TableOnPoolWait(t Table, c context.Context) func(_ context.Context, sessionID string, _ error) {
	var p PoolWaitStartInfo
	p.Context = c
	res := t.onPoolWait(p)
	return func(c context.Context, sessionID string, e error) {
		var p PoolWaitDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.Error = e
		res(p)
	}
}
func TableOnPoolTake(t Table, c context.Context, sessionID string) func(_ context.Context, sessionID string) func(_ context.Context, sessionID string, took bool, _ error) {
	var p PoolTakeStartInfo
	p.Context = c
	p.SessionID = sessionID
	res := t.onPoolTake(p)
	return func(c context.Context, sessionID string) func(context.Context, string, bool, error) {
		var p PoolTakeWaitInfo
		p.Context = c
		p.SessionID = sessionID
		res := res(p)
		return func(c context.Context, sessionID string, took bool, e error) {
			var p PoolTakeDoneInfo
			p.Context = c
			p.SessionID = sessionID
			p.Took = took
			p.Error = e
			res(p)
		}
	}
}
func TableOnPoolPut(t Table, c context.Context, sessionID string) func(_ context.Context, sessionID string, _ error) {
	var p PoolPutStartInfo
	p.Context = c
	p.SessionID = sessionID
	res := t.onPoolPut(p)
	return func(c context.Context, sessionID string, e error) {
		var p PoolPutDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.Error = e
		res(p)
	}
}
func TableOnPoolCloseSession(t Table, c context.Context, sessionID string) func(_ context.Context, sessionID string, _ error) {
	var p PoolCloseSessionStartInfo
	p.Context = c
	p.SessionID = sessionID
	res := t.onPoolCloseSession(p)
	return func(c context.Context, sessionID string, e error) {
		var p PoolCloseSessionDoneInfo
		p.Context = c
		p.SessionID = sessionID
		p.Error = e
		res(p)
	}
}
func TableOnPoolClose(t Table, c context.Context) func(context.Context, error) {
	var p PoolCloseStartInfo
	p.Context = c
	res := t.onPoolClose(p)
	return func(c context.Context, e error) {
		var p PoolCloseDoneInfo
		p.Context = c
		p.Error = e
		res(p)
	}
}
