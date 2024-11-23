package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	filebased *filebased.Data
}

func NewTaskRepo(filebasedDb *filebased.Data) *taskRepository {
	return &taskRepository{
		filebased: filebasedDb,
	}
}

func (t *taskRepository) Store(task *model.Task) error {
	t.filebased.StoreTask(*task)

	return nil
}

func (t *taskRepository) Update(task *model.Task) error {
	t.filebased.UpdateTask(task.ID, *task)

	
	return nil // TODO: replace this
}

func (t *taskRepository) Delete(id int) error {
	t.filebased.DeleteTask(id)
	return nil // TODO: replace this
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	TaskById,err := t.filebased.GetTaskByID(id)
	return TaskById, err // TODO: replace this
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	TaskList, err := t.filebased.GetTasks()
	return TaskList, err // TODO: replace this
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	TaskCtgr, err := t.filebased.GetTaskListByCategory(id)
	return TaskCtgr, err // TODO: replace this
}
