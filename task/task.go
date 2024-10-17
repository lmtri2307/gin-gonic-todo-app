package task

type Task struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Description string `json:"description"`
}
