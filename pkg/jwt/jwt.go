package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"time"
)

const TokenExpireDuration = time.Hour * 24 * 365

// CustomSecret 用于加盐的字符串
var CustomSecret = []byte("secret")

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID               int64  `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己声明的数据
	claims := CustomClaims{
		userID,
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), //过期时间
			Issuer:    "bluebell",                                              // 签发人
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, err := token.SignedString(CustomSecret)

	// 记录生成的token
	zap.L().Info("Generated token", zap.String("token", tokenString))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	var mc = new(CustomClaims)
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	//token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
	// 直接使用标准的Claim则可以直接使用Parse方法
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
