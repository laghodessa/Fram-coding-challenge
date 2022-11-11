package sqlite

import (
	"context"
	"database/sql"
	"personia/domain/hr"

	"github.com/jmoiron/sqlx"
)

type Employee struct {
	Name       string `db:"name"`
	Supervisor string `db:"supervisor"`
}

func NewHierarchyRepo(db *sql.DB) *HierarchyRepo {
	return &HierarchyRepo{
		db: sqlx.NewDb(db, "sqlite3"),
	}
}

type HierarchyRepo struct {
	db *sqlx.DB
}

// Get implements hr.HierarchyRepo
func (repo *HierarchyRepo) Get(ctx context.Context) (hr.Hierarchy, error) {
	q := `SELECT name, supervisor FROM employee`
	var rows []Employee
	if err := repo.db.SelectContext(ctx, &rows, q); err != nil {
		return nil, err
	}

	hier := make(hr.Hierarchy, len(rows))
	for _, r := range rows {
		hier[r.Name] = r.Supervisor
	}
	return hier, nil
}

// Update implements hr.HierarchyRepo
func (repo *HierarchyRepo) Update(ctx context.Context, hier hr.Hierarchy) error {
	q := `INSERT INTO employee (name, supervisor) VALUES (:name, :supervisor)`
	rows := make([]Employee, 0, len(hier))
	for empl, sup := range hier {
		rows = append(rows, Employee{
			Name:       empl,
			Supervisor: sup,
		})
	}

	if _, err := repo.db.NamedExecContext(ctx, q, rows); err != nil {
		return err
	}
	return nil
}

var _ hr.HierarchyRepo = (*HierarchyRepo)(nil)
