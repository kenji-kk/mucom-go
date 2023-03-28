package handler

import (
	"log"
	"net/http"

	"github.com/kenji-kk/mucom-go/internal/models"
	"github.com/labstack/echo/v4"
)

type ReviewHandler interface {
	// CheckToken(echo.Context) error
	GetReviews(echo.Context) error
}
type reviewHnadler struct {
	// usReview usecase.ReveiwUsecase
}

func NewReviewHandler() ReviewHandler {
	return &reviewHnadler{}
}

func (haReview *reviewHnadler) GetReviews(c echo.Context) error {

	// func WriteCookie(c echo.Context, cookieName, cookievalue string, expiresTime time.Duration) error {
	// 	cookie := &http.Cookie{
	// 		Name:     cookieName,
	// 		Value:    cookievalue,
	// 		Expires:  time.Now().Add(expiresTime * time.Hour),
	// 		HttpOnly: true,
	// 	}
	// cookie := new(http.Cookie)
	// cookie.Name = "username"
	// cookie.Value = "jon"
	// cookie.HttpOnly = false
	// cookie.Expires = time.Now().Add(1 * time.Hour)
	// c.SetCookie(cookie)

	// cookie, err := c.Cookie(cookie.Value)
	// if err != nil {
	// 	return err
	// }
	// log.Print("-- -- cookie取得 -- --")
	// fmt.Println(cookie.String())

	// token, ok := c.Get(jwt).(*jwt.Token) // by default token is stored under `user` key
	// if !ok {
	// 	return errors.New("JWT token missing or invalid")
	// }
	// claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	// if !ok {
	// 	return errors.New("failed to cast claims as jwt.MapClaims")
	// }
	// log.Print(token, claims)

	log.Print("-- --jwt解析処理-- --")
	// JWTの検証に使う署名キーを定義
	// var mySigningKey = []byte(jwt)

	// JWTをパースしてClaimsに格納
	log.Print("-- --reviews-- --")

	// //データの受け取り
	reviewResult := []*models.ReviewResultEntity{
		{
			Id:      1,
			Title:   "レビュータイトル",
			Content: "レビュー内容",
		},
		{
			Id:      2,
			Title:   "レビュータイトル2",
			Content: "レビュー内容2",
		},
	}
	return c.JSON(http.StatusOK, reviewResult)
}
