package core

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"user/model"
	"user/services"
)

//登录
func (u *UserService) UserLogin(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	var user model.User
	resp.Code = 200
	if err := model.DB.Where("user_name=?", req.UserName).First(&user).Error; err != nil { //查询不到username
		if gorm.IsRecordNotFoundError(err) { //	未找到记录错误
			resp.Code = 400 // 参数错误
			return nil
		}
		resp.Code = 500 //其他错误
		return nil
	}
	// 查询到username，就检验密码
	if user.CheckPassWord(req.Password) == false { //密码未通过
		resp.Code = 400 // 参数错误
		return nil
	}
	resp.UserDetail = BuildUser(user) //将得到的数据进序列化
	return nil
}

//将得到的数据进序列化
func BuildUser(item model.User) *services.UserModel {
	userModel := services.UserModel{
		ID:       uint32(item.ID),
		UserName: item.UserName,
		CreateAt: item.CreatedAt.Unix(),
		UpdateAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}

//UserRegister 用户注册
func (*UserService) UserRegister(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	//判断输入密码和输入的第二次密码是否是相等
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}
	count := 0
	if err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("用户名已经存在") // 用户名唯一
		return err
	}
	user := model.User{
		UserName: req.UserName,
	}
	//加密密码
	if err := user.SetPassWord(req.Password); err != nil {
		return err
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}
	resp.UserDetail = BuildUser(user)
	return nil
}
