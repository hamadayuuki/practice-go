package repository

import (
	"fmt"
	
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	// CRUD
	CreateTask(task *model.Task) error
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

// MARK: - Task Repository の実体
// DBへの操作 を実装する
// DBへの接続 は db.go で行う

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	/*
	Joins("User"):
		他のテーブルとの結合を指定します。ここでは、"User"という名前のテーブルとの結合を行います。
		これは、データベースのリレーショナルな特性を活用して、複数のテーブルから関連するデータを一度に取得するために使用されます。

	Where("user_id=?", userId): 
		クエリに条件を追加します。この条件は、user_idが指定されたuserIdに等しいレコードを選択するために使用されます。
		このようにして、特定のユーザーに関連するレコードだけをフィルタリングします。

	Order("created_at"): 
		結果のソート順を指定します。ここでは、レコードをcreated_atフィールドの値に基づいて並べ替えます。
		通常、これはレコードが作成された順序によるソートを意味します。

	Find(tasks): 
		上記の条件に合致するレコードをデータベースから取得し、それをtasks変数に格納します。
	*/
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	/*
	Clauses(clause.Returning{})
		Model(〜) 〜がさすポインタへ結果を代入する
	*/
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Update("title", task.Title)
	if result.Error != nil { 
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("[UpdateTask] object does not exist")
	}
	return nil
}

func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("[DeleteTask] object does not exist ")
	}
	return nil
}
