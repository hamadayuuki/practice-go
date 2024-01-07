package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUseCase interface {
	SignUp(user model.User) (model.UserResponse, error)
	// JWT token を返すため、string を返すようにしている
	Login(user model.User) (string, error)
}

// MARK: - UseCase の実体

type userUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) IUserUseCase {
	return &userUsecase{userRepository}
}

func (userUsecase *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	// パスワードをハッシュ化
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if error != nil { return model.UserResponse{}, error }

	// DB へ書きこむUserを作成
	newUser := model.User{Email: user.Email, Password: string(hashedPassword)}
	if err := userUsecase.userRepository.CreateUser(&newUser); err != nil {// TODO: CreateUser() は今後実装する
		return model.UserResponse{}, err 
	} 

	// repository? へのレスポンス用のUserを作成
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (userUsecase *userUsecase) Login(user model.User) (string, error) {
	// storedUser は初期化しているだけではないのか？ データが格納されているのか？ されていない場合CompareHashAndPassword() で何を比較しているのか
	storedUser := model.User{}   // 初期化
	if err := userUsecase.userRepository.GetUserByEmail(&storedUser, user.Email); err != nil {   // TODO: GetUserByEmail() は今後実装する
		return "", err
	}

	// パスワード の確認
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil { return "", err }

	// JWT を発行
	// JWT : JSON をやり取りする際のデジタル認証情報
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"user_id": storedUser.ID,
		"exp"    : time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))   // SECRET=uu5pveql
	if err != nil { return "", err }

	return tokenString, nil
}