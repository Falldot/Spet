package db

import (
	model "github.com/Survereignty/spet-rest-api/private/core/api/models"
)

type GroupRep struct {
	store *Store
}

// Создать Группу
func (r *GroupRep) Create(m *model.Group) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO groups (name) VALUES ($1) RETURNING id",
		m.Name,
	).Scan(&m.Id); err != nil {
		return err
	}

	return nil
}

// Удалить Группу
func (r *GroupRep) Delete(m *model.Group) error {
	_, err := r.store.db.Exec(
		"DELETE from groups where name = $1",
		m.Name,
	)
	if err != nil {
		return err
	}

	return nil
}

// Получить Группы
func (r *GroupRep) Get() ([]*model.Group, error) {

	query := "SELECT * FROM groups"

	rows, err := r.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sts := make([]*model.Group, 0)
	for rows.Next() {
		st := new(model.Group)
		err := rows.Scan(&st.Id, &st.Name)
		if err != nil {
			return nil, err
		}
		sts = append(sts, st)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sts, nil
}
