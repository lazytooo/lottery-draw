package usecase

import (
	"github.com/lazytooo/lottery-draw/models"
	"github.com/lazytooo/lottery-draw/repo"
)

type usecase struct {
	repo repo.Repository
}

func NewUsecase(repo repo.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (u usecase) CreatePlaystation(play models.LotteryDrawPlay) (playID int64, err error) {
	return u.repo.CreatePlaystation(play)
}

func (u usecase) CreateAward(award models.LotteryDrawAward) (awardID int64, err error) {
	return u.repo.CreateAward(award)
}

func (u usecase) CreateVisitor(visitor models.LotteryDrawVisitor) (err error) {
	return u.repo.CreateVisitor(visitor)
}

func (u usecase) GetPlaystations(userID int64) (list models.LotteryDrawPlays, err error) {
	return u.repo.GetPlaystations(userID)
}

func (u usecase) GoodLuck() {

}
