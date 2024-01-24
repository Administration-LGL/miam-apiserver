package util

import (
	v1 "miam-apiserver/api/common/v1"
	secretpb "miam-apiserver/api/secret/v1"
	userpb "miam-apiserver/api/user/v1"

	"miam-apiserver/internal/biz"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateUserReq2BizUser(req *userpb.CreateUserRequest) *biz.User {
	if req == nil {
		return &biz.User{}
	}
	return &biz.User{
		ObjectMeta: biz.ObjectMeta{
			Name: req.Meta.Name,
		},
		Status:      int(req.Status),
		Nickname:    req.Nickname,
		Password:    req.Password,
		Email:       req.Email,
		Phone:       req.Phone,
		IsAdmin:     int(req.IsAdmin),
		TotalPolicy: req.TotalPolicy,
	}
}

func GetUserReq2BizUser(req *userpb.GetUserRequest) *biz.User {
	if req == nil {
		return &biz.User{}
	}
	return &biz.User{
		Status:   int(req.Status),
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
		ObjectMeta: biz.ObjectMeta{
			ID: req.Id,
		},
	}
}

func BizUser2UserReply(user *biz.User) *userpb.UserReply {
	if user == nil {
		return &userpb.UserReply{}
	}
	return &userpb.UserReply{
		// Meta: &v1.ObjectMeta{
		// 	Id:         user.ID,
		// 	Name:       user.Name,
		// 	InstanceID: user.InstanceID,
		// 	// Extend: user.Extend,
		// 	CreateAt: timestamppb.New(user.CreatedAt),
		// 	UpdateAt: timestamppb.New(user.UpdatedAt),
		// },
		Meta:        objectMeta2ObjectMeta(&user.ObjectMeta),
		Status:      int32(user.Status),
		Nickname:    user.Nickname,
		Email:       user.Email,
		Phone:       user.Phone,
		IsAdmin:     int32(user.IsAdmin),
		TotalPolicy: user.TotalPolicy,
		LoginAt:     timestamppb.New(user.LoginedAt),
	}
}

func SecretList2SecretListReply(list *biz.SecretList) *secretpb.ListSecretReply {
	rlist := make([]*secretpb.Secret, 0, len(list.Items))
	for _, v := range list.Items {
		rlist = append(rlist, Secret2SecretReply(v))
	}
	return &secretpb.ListSecretReply{
		ListMeta: &v1.ListMeta{TotalCount: list.TotalCount},
		Secrets:  rlist,
	}
}

func Secret2SecretReply(secret *biz.Secret) *secretpb.Secret {
	return &secretpb.Secret{
		Meta:      objectMeta2ObjectMeta(&secret.ObjectMeta),
		Username:  secret.Username,
		SecretID:  secret.SecretID,
		SecretKey: secret.SecretKey,
		Expires:   secret.Expires,
	}
}

func objectMeta2ObjectMeta(meta *biz.ObjectMeta) *v1.ObjectMeta {
	return &v1.ObjectMeta{
		Id:         meta.ID,
		Name:       meta.Name,
		InstanceID: meta.InstanceID,
		// Extend: meta.Extend,
		CreateAt: timestamppb.New(meta.CreatedAt),
		UpdateAt: timestamppb.New(meta.UpdatedAt),
	}
}
