package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/lazytooo/lottery-draw/errors"
	"github.com/lazytooo/lottery-draw/models"
	"github.com/lazytooo/lottery-draw/usecase"
	"net/http"
	"strconv"
	"time"
)

type HttpHandler struct {
	ucase usecase.Usecase
}

func NewHttpHandler(e *echo.Echo, ucase usecase.Usecase) *HttpHandler {
	handler := HttpHandler{
		ucase: ucase,
	}

	baseRouter := e.Group("/lottery_draw")

	baseRouter.GET("/health_check", func(c echo.Context) error {
		msg := struct {
			Msg string `json:"msg"`
		}{"success"}
		return c.JSON(http.StatusOK, &msg)
	})

	managerRouter := baseRouter.Group("/api/manager")
	publicRouter := baseRouter.Group("/api/public")
	// 创建活动信息
	managerRouter.POST("/user/:user_id/playstation", handler.createPlaystation)
	// 在活动下创建抽奖项
	managerRouter.POST("/user/:user_id/playstation/:play_id/award", handler.createAward)
	// 开奖🎉
	managerRouter.POST("/user/:user_id/playstation/:play_id/award/:award_id/good_luck", handler.goodLuck)

	// 获取活动信息
	publicRouter.GET("/user/:user_id/playstations", handler.getPlaystations)
	// 获取活动下的抽奖项列表
	publicRouter.GET("/user/:user_id/playstation/:play_id/awards", handler.getAwardList)
	// 获取指定抽奖项信息
	publicRouter.GET("/user/:user_id/playstation/:play_id/award/:award_id", handler.getAwardInfo)
	// 获取活动下的抽奖用户列表
	publicRouter.GET("/user/:user_id/playstation/:play_id/visitors", handler.getPlayVisitors)
	// 参与抽奖
	publicRouter.POST("/user/:user_id/playstation/:play_id/visitor", handler.createPlayVisitor)
	return &handler
}

func (h HttpHandler) createPlaystation(c echo.Context) error {
	reqData := models.LotteryDrawPlay{}
	respData := errors.GetErrorBaseResponse(nil)
	data := struct {
		PlayID int64 `json:"play_id"`
	}{}
	userID, _ := strconv.Atoi(c.Param("user_id"))
	if userID <= 0 {
		return c.JSON(http.StatusBadRequest, errors.GetErrorBaseResponse(errors.InvalidParams))
	}
	if err := c.Bind(&reqData); err != nil {
		return c.JSON(http.StatusBadRequest, errors.GetErrorBaseResponse(errors.InvalidParams))
	}

	reqData.UserID = int64(userID)
	reqData.CreateTime = time.Now().Local().Format(models.TIME_FORMAT)

	playID, err := h.ucase.CreatePlaystation(reqData)
	if err != nil {
		fmt.Println(err)
	}

	data.PlayID = playID
	respData.Data = data

	return c.JSON(http.StatusOK, respData)
}

func (h HttpHandler) createAward(c echo.Context) error {
	playID, _ := strconv.Atoi(c.Param("play_id"))
	reqData := models.LotteryDrawAward{}
	respData := errors.GetErrorBaseResponse(nil)
	data := struct {
		AwardID int64 `json:"award_id"`
	}{}
	if err := c.Bind(&reqData); err != nil {
		c.JSON(http.StatusBadRequest, errors.GetErrorBaseResponse(errors.InvalidParams))
	}

	reqData.PlayID = int64(playID)
	reqData.AwardRestNum = reqData.AwardNum
	reqData.CreateTime = time.Now().Local().Format(models.TIME_FORMAT)
	reqData.UpdateTime = time.Now().Local().Format(models.TIME_FORMAT)

	awardID, err := h.ucase.CreateAward(reqData)
	if err != nil {
		fmt.Println(err)
	}

	data.AwardID = awardID
	respData.Data = data

	return c.JSON(http.StatusOK, respData)
}

func (h HttpHandler) goodLuck(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (h HttpHandler) getPlaystations(c echo.Context) error {
	respData := errors.GetErrorBaseResponse(nil)

	userID, _ := strconv.Atoi(c.Param("user_id"))

	list, err := h.ucase.GetPlaystations(int64(userID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.GetErrorBaseResponse(errors.InvalidParams))
	}

	respData.Data = list
	return c.JSON(http.StatusOK, respData)
}

func (h HttpHandler) getAwardList(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (h HttpHandler) getAwardInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (h HttpHandler) getPlayVisitors(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (h HttpHandler) createPlayVisitor(c echo.Context) error {
	playID, _ := strconv.Atoi(c.Param("play_id"))

	reqData := models.LotteryDrawVisitor{}
	if err := c.Bind(&reqData); err != nil {
		return c.JSON(http.StatusBadRequest, errors.GetErrorBaseResponse(errors.InvalidParams))
	}

	reqData.PlayID = int64(playID)
	reqData.CreateTime = time.Now().Local().Format(models.TIME_FORMAT)
	err := h.ucase.CreateVisitor(reqData)
	fmt.Println(err)
	return c.JSON(http.StatusOK, reqData)
}
