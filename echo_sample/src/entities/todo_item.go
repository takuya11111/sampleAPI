package entities

type TodoItem struct {
	ID       string `db:"id" json:"id"`
	TodoItem string `db:"todo" json:"todo"`
}
