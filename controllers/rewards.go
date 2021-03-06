package controllers

import (
	"github.com/ewagmig/contract-connector/server"
	"github.com/op/go-logging"
)

var rewardslogger = logging.MustGetLogger("rewards.collection.controller")

func init() {
	err := server.RegisterController(&rewardsCol{})
	if err != nil {
		rewardslogger.Errorf("user can not be registered")
	}
}


type rewardsCol struct {
}

func (rc *rewardsCol) Name() string {
	return "rewards"
}


func (rc *rewardsCol) Routes() []*server.Router {
//	//jwt := admin.CreateJWTFactory(false, models.Login, models.VerifyUserPermission)
	return []*server.Router{
//		{
//			Path:         "/getValidatorsRewards",
//			Method:       "GET",
//			//AuthType: utils.BasicAuth,
//			Handler:      rc.getRewards,
//		},
//		//{
//		//	Path:         "/put",
//		//	Method:       "POST",
//		//	//AuthType: utils.BasicAuth,
//		//	Handler:      rc.putState,
//		//},
	}
}

//func (rc *rewardsCol) getRewards(ctx *gin.Context)  {
//	req := &models.CallParams{}
//	if err := utils.GetJSONBody(ctx, req); err != nil {
//		errors.BadRequestError(errors.InvalidJSONBody, err.Error()).Write(ctx)
//		return
//	}
//
//	resp, err := models.GetRewards(req)
//	if err != nil {
//		errors.BadRequestError(errors.IDNotFound, err.Error()).Write(ctx)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, resp)
//}




