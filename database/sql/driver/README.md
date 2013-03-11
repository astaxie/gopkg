# driver包

这个包里面定义很多需要实现的接口，以便/database/sql来调用

* func IsScanValue(v interface{}) bool
* func IsValue(v interface{}) bool
* type ColumnConverter
* type Conn
* type Driver
* type Execer
* type NotNull
    * func (n NotNull) ConvertValue(v interface{}) (Value, error)
* type Null
    * func (n Null) ConvertValue(v interface{}) (Value, error)
* type Result
* type Rows
* type RowsAffected
    * func (RowsAffected) LastInsertId() (int64, error)
    * func (v RowsAffected) RowsAffected() (int64, error)
* type Stmt
* type Tx
* type Value
* type ValueConverter
* type Valuer