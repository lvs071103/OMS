package logic

import (
	"fmt"
	"strconv"
	"time"

	"github.com/oms/dao/mysql"
	"github.com/oms/models"
	"github.com/oms/pkg/jwt"
	"github.com/oms/pkg/snowflake"
	"golang.org/x/exp/rand"
)

var verifyCodes = make(map[string]string)

// GenerateAndSendVerifyCode - 生成并发送验证码
func GenerateAndSendVerifyCode(email string) (err error) {
	// 生成验证码
	rand.Seed(uint64(time.Now().UnixNano()))
	code := fmt.Sprintf("%06d", rand.Int31n(1000000))
	// 将验证存储到内存中(可以使用持久化的存储，如Redis)
	verifyCodes[email] = code

	// 发送验证码到用户的邮箱（这里省略实际发送邮件的代码）
	fmt.Printf("Sending verification code %s to email %s\n", code, email)
	return nil
}

// VerifyCode - 验证验证码是否正确
func VerifyCode(email, code string) bool {
	if storedCode, exists := verifyCodes[email]; exists && storedCode == code {
		delete(verifyCodes, email) // 验证成功后删除验证码
		return true
	}
	return false
}

// SignUp - 注册业务逻辑的代码
func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	if err := mysql.CheckUserIsExists(p.UserName); err != nil {
		return err
	}
	// 2.生成uid
	userID := snowflake.GenID()
	// 构造一个User实例
	user := models.AuthUser{
		ID:       userID,
		UserName: p.UserName,
		Password: p.Password,
		Email:    p.Email,
	}
	// 3.保存数据进数据库
	return mysql.InsertUser(&user)
}

func Login(p *models.ParamsLogin) (user *models.ResponseUser, err error) {
	user = &models.ResponseUser{
		UserName: p.UserName,
		Password: p.Password,
	}
	//传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT的token
	userID, err := strconv.ParseInt(user.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	token, err := jwt.GenToken(userID, user.UserName)
	if err != nil {
		return
	}
	user.Token = token
	return
}
