package implement

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/poin4003/eCommerce_golang_api/global"
	"github.com/poin4003/eCommerce_golang_api/internal/consts"
	"github.com/poin4003/eCommerce_golang_api/internal/database"
	"github.com/poin4003/eCommerce_golang_api/internal/model"
	"github.com/poin4003/eCommerce_golang_api/internal/utils"
	"github.com/poin4003/eCommerce_golang_api/internal/utils/crypto"
	"github.com/poin4003/eCommerce_golang_api/internal/utils/random"
	"github.com/poin4003/eCommerce_golang_api/internal/utils/sendto"
	"github.com/poin4003/eCommerce_golang_api/pkg/response"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"strings"
	"time"
)

type sUserLogin struct {
	r *database.Queries
}

func NewUserLoginImplement(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}

func (s *sUserLogin) Login(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	//. 1 hash email
	fmt.Printf("VerifyKey: %s\n", in.VerifyKey)
	fmt.Printf("VerifyType: %d\n", in.VerifyType)
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("hashKey: %s\n", hashKey)

	//. 2 check user exists in user base
	userFound, err := s.r.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeUserHasExists, err
	}

	if userFound > 0 {
		return response.ErrCodeUserHasExists, fmt.Errorf("user %s already exists", in.VerifyType)
	}

	// 3. Create OTP
	userKey := utils.GetUserKey(hashKey)
	otpFound, err := global.Rdb.Get(ctx, userKey).Result()

	switch {
	case err == redis.Nil:
		fmt.Println("Key dose not exists")
	case err != nil:
		fmt.Println("Get failed::", err)
		return response.ErrInvalidOTP, err
	case otpFound != "":
		return response.ErrCodeOtpNotExists, fmt.Errorf("otp %s exists but not registered", otpFound)
	}

	// 4. Generate OTP
	otpNew := random.GenerateSixDigitOtp()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Printf("Otp is:::%d\n", otpNew)

	// 5. save OTP into Redis with expiration time
	err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrInvalidOTP, err
	}

	// 6. send OTP
	switch in.VerifyType {
	case consts.EMAIL:
		err = sendto.SendTextEmailOtp([]string{in.VerifyKey}, consts.HOST_EMAIL, strconv.Itoa(otpNew))
		if err != nil {
			return response.ErrSendEmailOTP, err
		}
		// 7. save OTP to MySql
		result, err := s.r.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
			VerifyOtp:     strconv.Itoa(otpNew),
			VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
			VerifyKey:     in.VerifyKey,
			VerifyKeyHash: hashKey,
		})

		if err != nil {
			return response.ErrSendEmailOTP, err
		}

		// 8. getlastId
		lastIdVerifyUser, err := result.LastInsertId()
		if err != nil {
			return response.ErrSendEmailOTP, err
		}
		log.Println("LastIdVerifyUser:", lastIdVerifyUser)

		return response.ErrCodeSuccess, nil
	case consts.MOBILE:
		return response.ErrCodeSuccess, nil
	}

	return response.ErrCodeSuccess, nil
}

func (s *sUserLogin) VerifyOTP(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}
