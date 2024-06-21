# Concurrency controll in postgres

Some PostgreSQL data types and functions have special rules regarding transactional behavior. In particular, changes made to a sequence (and therefore the counter of a column declared using serial) are immediately visible to all other transactions and are not rolled back if the transaction that made the changes aborts.

### Various Phenomena prohibited at various levels are

#### 1. dirty read 
  - data written by any concurrent non commited transaction
#### 2. nonrepeatable read
  - transaction re reads data and founds it changed from previous data.
#### 3. phantom read
  - TODO

#### 4. serialization anamoly
  - The result of successfully committing a group of transactions is inconsistent with all possible orderings of running those transactions one at a time.

## Isolation Levels in Postgresql

#### 1. Read Commited
  - Select query sees only commited data.
  - Neither it sees concurrent commited data. only sees snapshot of database at the time of query begin.
  - Two successive select statements in same transaction can see different data.
  - UPDATE, DELETE, SELECT FOR UPDATE, and SELECT FOR SHARE commands behave the same as SELECT in terms of searching for target rows.
  - esa ho sakta he ki koi transaction parallelly chal raha ho agar vo rollback hota he then dusra transaction start kar dega update karna but if not then ham new version pe dekhenge ki kya abhi bhi hamari condition match ho rahi he and then we procceed.
  - In the case of SELECT FOR UPDATE and SELECT FOR SHARE, this means it is the updated version of the row that is locked and returned to the client.
  - see INSERT WITH ON CONFLICT DO NOTHING and DO UPDATE case
  - The partial transaction isolation provided by Read Committed mode is adequate for many applications, and this mode is fast and simple to use


commit vale data ko read karenge.. 2 select statements can produce different output in same transaction if other tx changes the data.

#### For Update
Another common parameter is NOWAIT, which returns an error immediately if a transaction is not able to immediately lock a row. In SQL syntax, NOWAIT appears directly after FOR UPDATE, like so:


Reference : https://www.cockroachlabs.com/blog/select-for-update/

#### 2. Repeatable Read Isolation Level

NOTE : Applications using this level must be prepared to retry transactions due to serialization failures.

- In this case two successive select statements will not result in different data because
- They do not see changes made by other transactions that committed after their own transaction started.
- This is a stronger guarantee than is required by the SQL standard for this isolation level except serialization failures.
- It never sees uncommited changes nor commited changes by other transaction after its transaction begin. hence bad me vo khud ko hi dekhega ha ha.
- UPDATE, DELETE, MERGE, SELECT FOR UPDATE and SELECT FOR SHARE behaves same as SELECT, but it will only find records that were commited as of transaction start time. if target row updated by concurrent transaction too, in this case repeatable read transaction will wait for first transaction to completeted. if it is rollback then second transaction starts but if it is commited then second transaction will get rollbacked and it will get message :
- ERROR:  could not serialize access due to concurrent update
- because a repeatable read transaction cannot modify or lock rows changed by other transactions after the repeatable read transaction began. [ ek bar chalu hone ke bad ye kisiki nahi sunata ha ha ]

- NOTE : Note that only updating transactions might need to be retried; read-only transactions will never have serialization conflicts.

#### 3. Serializable Isolation Level

- Provides stricted isolation level.
- emulates the serial execution for all commited transactions.
- It is same as the repeatable read but it also monitors for the condition where there is a serialization anamoly between concurrent transactions.
- This monitoring does not introduce any blocking beyond that present in repeatable read, but there is some overhead to the monitoring, and detection of the conditions which could cause a serialization anomaly will trigger a serialization failure.


## Explicit Loacking

- Prostgresql provides various lock modes to control concurrent access to data in tables.
- To examine currently outstanding locks in database server use `pg_locks` system view.

1. Table level locks
2. Row level locks
3. Page level locks
4. Deadlocks
5. Advisory locks


- Locks automatically applied by postgresql.
- we can also acquire locks explicitly with the command LOCK.
- Two transactions cannot hold locks of conflicting modes on the same table at the same time. (However, a transaction never conflicts with itself. For example, it might acquire ACCESS EXCLUSIVE lock and later acquire ACCESS SHARE lock on the same table.)

### 1. Table Level Lock Modes

#### ACCESS SHARE
    - conflicts with ACCESS EXCLUSIVE only.
    - any query that reads only will acquire this lock. ex. select command

#### ROW SHARE
    - conflicts with EXCLUSIVE and ACCESS  EXCLUSIVE.
    - select command will acquire this locks when any of (FOR UPDATE, FOR NO KEY UPDATE, FOR SHARE and FOR KEY SHARE) option is specified.

#### ROW EXCLUSIVE
    - Conflicts with the SHARE, SHARE ROW EXCLUSIVE, EXCLUSIVE, and ACCESS EXCLUSIVE lock modes.
    -   UPDATE, DELETE, INSERT and MERGE acquire this lock mode on target table.

#### SHARE UPDATE EXCLUSIVE
    - This mode protects a table against concurrent schema changes and vaccum runs.
    - conflicts with SHARE UPDATE EXCLUSIVE, SHARE, SHARE ROW EXCLUSIVE, EXCLUSIVE, and ACCESS EXCLUSIVE
    - Acquired by VACUUM (without FULL), ANALYZE, CREATE INDEX CONCURRENTLY, CREATE STATISTICS, COMMENT ON, REINDEX CONCURRENTLY, and certain ALTER INDEX and ALTER TABLE variants

#### SHARE


#### SHARE ROW EXCLUSIVE


#### EXCLUSIVE


#### ACCESS EXCLUSIVE



### 2. Row Level Lock Modes

- Row-level locks do not affect data querying; they block only writers and lockers to the same row. Row-level locks are released at transaction end or during savepoint rollback, just like table-level locks.

#### FOR UPDATE
    - causes row retrieved by the select statement to be locked as thogugh for update. this prevents any changes to them until current transaction completes.
    - blocks other transaction till commit. if data updated then repeatable and serializable transaction will get error message.

#### FOR NO KEY UPDATE
    - similar to FOR UPDATE but lock is weaker lock.
    -  this lock will not block SELECT FOR KEY SHARE commands that attempt to acquire a lock on the same rows. This lock mode is also acquired by any UPDATE that does not acquire a FOR UPDATE lock.


#### FOR SHARE
    - similar to FOR NO KEY UPDATE, except that it acquires shared lock rather then exclusive lock on each retrived rows.
    - A shared lock blocks other transactions from performing UPDATE, DELETE, SELECT FOR UPDATE or SELECT FOR NO KEY UPDATE on these rows, but it does not prevent them from performing SELECT FOR SHARE or SELECT FOR KEY SHARE.

#### FOR KEY SHARE ???
    - similar to FOR SHARE, except that this lock is weaker.
    - SELECT FOR UPDATE is blocked, but not SELECT FOR NO KEY UPDATE. A key-shared lock blocks other transactions from performing DELETE or any UPDATE that changes the key values, but not other UPDATE, and neither does it prevent SELECT FOR NO KEY UPDATE, SELECT FOR SHARE, or SELECT FOR KEY SHARE.

#### NOTE
PostgreSQL doesn't remember any information about modified rows in memory, so there is no limit on the number of rows locked at one time. However, locking a row might cause a disk write, e.g., SELECT FOR UPDATE modifies selected rows to mark them locked, and so will result in disk writes.


### 3. Page Level Locks

- Application developers need not to be concerned about this. database handles it.


### 4. Deadlocks

- The use of explicit locking can increase the likelihood of deadlocks, wherein two (or more) transactions each hold locks that the other wants. For example, if transaction 1 acquires an exclusive lock on table A and then tries to acquire an exclusive lock on table B, while transaction 2 has already exclusive-locked table B and now wants an exclusive lock on table A, then neither one can proceed. PostgreSQL automatically detects deadlock situations and resolves them by aborting one of the transactions involved, allowing the other(s) to complete. (Exactly which transaction will be aborted is difficult to predict and should not be relied upon.)

How to avoid deadlocks effectively ?? see documentations..

### 5. Advisory locks

- Imagine a scenario where you have a distributed system talking to a central database. You are supposed to control the access to the data for all the systems or nodes that are talking to the database and no two nodes are supposed to alter the data at the same time. 

- Imagine it this way. The advisory lock is a gateway for certain piece of code in your application while the other pieces of code in your application is free to do whatever it wants to do with the data.

- Other use cases for advisory locks back be when you want to run a background process that should be executed by only one worker or a node. Advisory locks can be used to make sure you are not wasting your compute power, executing the process more than once.

#### Types of advisory locks in postgresql

- You can acquire an advisory lock on a `session level or a transaction level`. Just like the other locks, if the lock is acquired on a transaction level, the lock is released when the transaction is complete. Similarly, a session level lock is released when the session ends or when you manually release it.

- you can acquire it by blocking or non blocking function. ( wait or not wait )

- Example

```
// to acquire lock using non blocking function
SELECT pg_try_advisory_lock(100);

// blocking function
SELECT pg_advisory_lock(100);

// get all the advisory locks
SELECT mode, classid, objid FROM pg_locks WHERE locktype = 'advisory';
```

### Lock Cluase

```
FOR lock_strength [ OF table_name [, ...] ] [ NOWAIT | SKIP LOCKED ]
```

- To prevent the operation from waiting for other transactions to commit, use either the NOWAIT or SKIP LOCKED option. With NOWAIT, the statement reports an error, rather than waiting, if a selected row cannot be locked immediately. With SKIP LOCKED, any selected rows that cannot be immediately locked are skipped. Skipping locked rows provides an inconsistent view of the data, so this is not suitable for general purpose work, but can be used to avoid lock contention with multiple consumers accessing a queue-like table. Note that NOWAIT and SKIP LOCKED apply only to the row-level lock(s) — the required ROW SHARE table-level lock is still taken in the ordinary way (see Chapter 13). You can use LOCK with the NOWAIT option first, if you need to acquire the table-level lock without waiting.

### SQL 

## References :

- https://www.postgresql.org/docs/current/applevel-consistency.html
- 