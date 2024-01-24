package encrypt

import "golang.org/x/crypto/bcrypt"

func EncryptPwd(pwd string) (string, error) {
	// 生成密码的哈希值
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// 将哈希值转换为字符串并返回
	return string(hashedPwd), nil
}

func ComparePwd(hashedPwd string, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}
