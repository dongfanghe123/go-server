package handler

import (
	"context"
	"crypto/rand"
	"fmt"
	"go-server/dao"
	"go-server/entity"
	"go-server/logger"
	"go-server/manager"
	"go-server/service"
	"go-server/util"
	"math/big"

	"github.com/gin-gonic/gin"
)

func GetPhoneCode(c *gin.Context) {

	phone := c.Query("phone")
	if !util.IsValidPhone(phone) {
		c.JSON(400, gin.H{
			"msg": "手机号格式错误",
		})
		return
	}

	code, err := GenCode()
	if err != nil {
		logger.Log.Error().Msg(("生成验证码错误:" + err.Error()))
		c.JSON(500, gin.H{
			"msg": "生成验证码错误",
		})
		return
	}
	fmt.Println("验证码：", code)

	//client.Set(c.Request.Context(), "phone:code:"+phone, code, 2*time.Minute)

	manager.Client.Set(c.Request.Context(), "phone:code:"+phone, code, 0)

	c.JSON(200, gin.H{
		"msg": "生成验证码成功",
	})

}

func UserLogin(c *gin.Context) {
	var req entity.LoginRequestV2

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	if !util.IsValidPhone(*req.Phone) {
		c.JSON(200, gin.H{
			"success": false,
			"msg":     "手机号格式错误",
		})
	}

	//校验用户名和密码
	code := manager.Client.Get(c.Request.Context(), "phone:code:"+*req.Phone)

	if *req.Code != code.Val() {
		c.JSON(500, gin.H{
			"msg": "验证码错误",
		})
	} else {

		//判断当前用户是否注册，如果未注册，自动注册
		params := make(map[string]interface{})
		params["phone"] = *req.Phone
		user, err := dao.DefaultUserDao.SelectUserByPhone(c.Request.Context(), params)
		if err != nil {
			fmt.Println("查询用户信息失败")
			c.JSON(500, gin.H{
				"msg": "查询用户信息失败",
			})
		}

		if user == nil {
			//自动注册
			ctx := context.Background()
			tk, err := service.UserRegisterByPhone(ctx, *req.Phone)
			if err != nil {
				fmt.Println("自动注册失败")
				c.JSON(500, gin.H{
					"msg": err.Error(),
				})
				return
			} else {
				c.JSON(200, gin.H{
					"msg": "",
					"data": gin.H{
						"token": tk,
					},
				})
				return
			}

		} else {
			id := user.ID
			nickName := user.NickName
			token, err := util.GenerateToken(*id, *nickName)
			if err != nil {
				c.JSON(500, gin.H{
					"msg":  "生成token错误:" + err.Error(),
					"data": "",
				})
				return
			} else {
				c.JSON(200, gin.H{
					"msg": "",
					"data": gin.H{
						"token": token,
					},
				})
				return
			}
		}

	}
}

func GenCode() (string, error) {
	max := big.NewInt(1000000) // \[0, 999999\]
	num, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	// 不足6位前面补0
	return fmt.Sprintf("%06d", num), nil
}
