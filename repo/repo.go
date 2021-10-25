package repo

import "github.com/lazytooo/lottery-draw/models"

type Repository interface {
	CreatePlaystation(play models.LotteryDrawPlay) (playID int64, err error)
	CreateAward(award models.LotteryDrawAward) (awardID int64, err error)
	CreateVisitor(visitor models.LotteryDrawVisitor) (err error)
	GetPlaystations(userID int64) (list models.LotteryDrawPlays, err error)
}
