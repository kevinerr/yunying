package repository

import "yy/model"

type RecordRepository struct {
}

//insert one comment
func (c RecordRepository) CreatComment(record *model.Reocrd) error {
	err := model.DB.Create(record).Error
	return err
}
