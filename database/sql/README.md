# sql包

sql包实现了基于sql的通用接口，里面有查询、执行、事务等接口

* func Register(name string, driver driver.Driver)
* type DB
    * func Open(driverName, dataSourceName string) (*DB, error)
    * func (db *DB) Begin() (*Tx, error)
    * func (db *DB) Close() error
    * func (db *DB) Driver() driver.Driver
    * func (db *DB) Exec(query string, args ...interface{}) (Result, error)
    * func (db *DB) Prepare(query string) (*Stmt, error)
    * func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
    * func (db *DB) QueryRow(query string, args ...interface{}) *Row
* type NullBool
    * func (n *NullBool) Scan(value interface{}) error
    * func (n NullBool) Value() (driver.Value, error)
* type NullFloat64
    * func (n *NullFloat64) Scan(value interface{}) error
    * func (n NullFloat64) Value() (driver.Value, error)
* type NullInt64
    * func (n *NullInt64) Scan(value interface{}) error
    * func (n NullInt64) Value() (driver.Value, error)
* type NullString
    * func (ns *NullString) Scan(value interface{}) error
    * func (ns NullString) Value() (driver.Value, error)
* type RawBytes
* type Result
* type Row
    * func (r *Row) Scan(dest ...interface{}) error
* type Rows
    * func (rs *Rows) Close() error
    * func (rs *Rows) Columns() ([]string, error)
    * func (rs *Rows) Err() error
    * func (rs *Rows) Next() bool
    * func (rs *Rows) Scan(dest ...interface{}) error
* type Scanner
* type Stmt
    * func (s *Stmt) Close() error
    * func (s *Stmt) Exec(args ...interface{}) (Result, error)
    * func (s *Stmt) Query(args ...interface{}) (*Rows, error)
    * func (s *Stmt) QueryRow(args ...interface{}) *Row
* type Tx
    * func (tx *Tx) Commit() error
    * func (tx *Tx) Exec(query string, args ...interface{}) (Result, error)
    * func (tx *Tx) Prepare(query string) (*Stmt, error)
    * func (tx *Tx) Query(query string, args ...interface{}) (*Rows, error)
    * func (tx *Tx) QueryRow(query string, args ...interface{}) *Row
    * func (tx *Tx) Rollback() error
    * func (tx *Tx) Stmt(stmt *Stmt) *Stmt