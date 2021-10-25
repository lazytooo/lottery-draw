package repo

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/lazytooo/lottery-draw/models"
)

type repository struct {
	DB  *sqlx.DB
	RDS *redis.Client
}

func NewRepository(db *sqlx.DB, rds *redis.Client) Repository {
	return &repository{
		DB:  db,
		RDS: rds,
	}
}

func (r repository) CreatePlaystation(play models.LotteryDrawPlay) (playID int64, err error) {
	sqlStr := `INSERT INTO lottery_draw_play (play_name, user_id, create_time) VALUES (:play_name, :user_id, :create_time);`

	rows, err := r.DB.NamedExec(sqlStr, &play)
	if err != nil {
		return 0, err
	}

	playID, _ = rows.LastInsertId()
	return playID, nil
}

func (r repository) CreateAward(award models.LotteryDrawAward) (awardID int64, err error) {
	sqlStr := `INSERT INTO lottery_draw_award (play_id, award_name, award_num, award_rest_num, create_time, update_time) 
			VALUES (:play_id, :award_name, :award_num, :award_rest_num, :create_time, :update_time);`

	rows, err := r.DB.NamedExec(sqlStr, &award)
	if err != nil {
		return 0, err
	}

	awardID, _ = rows.LastInsertId()
	return awardID, nil
}

func (r repository) CreateVisitor(visitor models.LotteryDrawVisitor) (err error) {
	sqlStr := `INSERT INTO lottery_draw_visitor (play_id, visitor_name, visitor_phone, ticket, create_time)
			VALUES (:play_id, :visitor_name, :visitor_phone, :ticket, :create_time)`

	_, err = r.DB.NamedExec(sqlStr, &visitor)
	if err != nil {
		return err
	}

	return nil

}

func (r repository) GetPlaystations(userID int64) (list models.LotteryDrawPlays, err error) {
	sqlStr := `SELECT id, play_name FROM lottery_draw_play WHERE user_id = ?`
	err = r.DB.Select(&list, sqlStr, userID)
	if err != nil {
		return
	}
	return
}
