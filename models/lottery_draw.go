package models

import (
	"errors"
	"math/rand"
	"time"
)

const (
	TIME_FORMAT = "2006-01-02 13:04:05"
)

type LotteryDrawPlay struct {
	ID         int64  `json:"play_id" db:"id"`
	UserID     int64  `json:"user_id,omitempty" db:"user_id"`
	PlayName   string `json:"play_name" db:"play_name"`
	CreateTime string `json:"create_time,omitempty" db:"create_time"`
}

type LotteryDrawPlays []LotteryDrawPlay

type LotteryDrawAward struct {
	ID           int64  `json:"award_id" db:"id"`
	PlayID       int64  `json:"play_id" db:"play_id"`
	AwardName    string `json:"award_name" db:"award_name"`
	AwardNum     int64  `json:"award_num" db:"award_num"`
	AwardRestNum int64  `json:"award_rest_num" db:"award_rest_num"`
	IsOpen       bool   `json:"is_open" db:"is_open"`
	CreateTime   string `json:"create_time" db:"create_time"`
	UpdateTime   string `json:"update_time" db:"update_time"`
}

type LotteryDrawVisitor struct {
	ID           int64  `json:"visitor_id" db:"id"`
	PlayID       int64  `json:"play_id" db:"play_id"`
	VisitorName  string `json:"visitor_name" db:"visitor_name"`
	VisitorPhone string `json:"visitor_phone" db:"visitor_phone"`
	Ticket       string `json:"ticket" db:"ticket"`
	CreateTime   string `json:"create_time" db:"create_time"`
}

type LotteryDrawShot struct {
	ID          int64  `json:"shot_id" db:"id"`
	AwardID     int64  `json:"award_id" db:"award_id"`
	AwardName   string `json:"award_name" db:"award_name"`
	VisitorID   int64  `json:"visitor_id" db:"visitor_id"`
	VisitorName string `json:"visitor_name" db:"visitor_name"`
	Ticket      string `json:"ticket" db:"ticket"`
}

type LotteryDrawVisitorList []LotteryDrawVisitor

func (us LotteryDrawVisitorList) GetLotteryDrawUser(num int) (err error) {
	us, err = shuffle(num, us)
	if err != nil {
		return err
	}
	return nil
}

func shuffle(num int, slices LotteryDrawVisitorList) (LotteryDrawVisitorList, error) {
	if num >= len(slices) {
		return slices, errors.New("invalid params")
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slices) > 0 {
		n := len(slices)
		randIndex := r.Intn(n)
		slices[n-1], slices[randIndex] = slices[randIndex], slices[n-1]
		slices = slices[:n-1]
	}

	return slices[0 : num-1], nil

}
