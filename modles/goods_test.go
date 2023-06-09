package models

import (
	"testing"
	"time"
)

func TestSaveGoods(t *testing.T) {
	goods := Goods{
		Title:      "苹果",
		Price:      20,
		Stock:      100,
		Type:       0,
		CreateTime: time.Now(),
	}
	SvaeGoods(goods)

}

func TestUpdateGoods(t *testing.T) {
	UpdateGoods()
}

func TestDeleteGoods(t *testing.T) {
	DeleteGoods()
}

func TestFindGoods(t *testing.T) {
	FindGoods()
}

func TestSessionContext(t *testing.T) {
	SessionContext()
}

func TestTranscation(t *testing.T) {
	Transcation()
}

func TestFindUserAndGoods(t *testing.T) {
	FindUserAndGoods()
}
