package migrator

type MigrationCondition interface {
	SQL(dialect Dialect) (string, []interface{})
	IsFulfilled(results []map[string][]byte) bool
}

type ExistsMigrationCondition struct {}

func (c *ExistsMigrationCondition) IsFulfilled(results []map[string][]byte) bool {
	return len(results) >= 1
}

type NotExistsMigrationCondition struct{}

func (c *NotExistsMigrationCondition) IsFulfilled(results []map[string][]byte) bool {
	return len(results) == 0
}

type IfIndexExistsCondition struct {
	ExistsMigrationCondition
	TableName string
	IndexName string
}

func (c *IfIndexExistsCondition) SQL(dialect Dialect) (string, []interface{}) {
	return dialect.
}
