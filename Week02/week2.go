package Week02

import (
	"database/sql"
	"errors"
	"fmt"
)

// dao 层中是否应该 Wrap 这个 error抛给上层？  分情况：（1）如果查询 无关紧要的逻辑可以内部现场处理 （2）若是增 删 改错误级别 就需要wrap error到上层了
//sql.ErrNoRows 不需要抛异常  因为它不影响任何逻辑
type Dao struct {
	......
}


func New(......) *Dao {
	return &Dao{ ...... }
}

func (d *Dao) SelectOper(condition1 int) ([] *Result,  error) {
	var records []*Result
	results,err := DB.SelectOper(condition1)
	if errors.Is(err, sql.ErrNoRows) {
		err = ErrRecordNotFound
	}
	if err != nil {
		if err == sql.ErrNoRows{
			return nil,nil
		}
		return nil,err
	}
	for v := range results {
		temp := &Result{}
		temp.Record = v
		records = append(records,temp)
	}
	return records,nil
}


// Service 层

type Service struct {
	......
}


func (s *Service) OperUser(userID int){
	records, err  := dao.SelectOper(userID)
	if err != nil{
		dosomething
		return
	}
	if records == nil{
		doSomething
		return
	}
	dosomething
}