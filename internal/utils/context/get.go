package context

import (
	"context"
	"errors"

	"github.com/poin4003/eCommerce_golang_api/internal/consts"
	"github.com/poin4003/eCommerce_golang_api/internal/utils/cache"
)

type InfoUserUUID struct {
	UserID      uint64
	UserAccount string
}

func GetSubjectUUID(ctx context.Context) (string, error) {
	sUUID, ok := ctx.Value(consts.SUBJECT_UUID_KEY).(string)
	if !ok {
		return "", errors.New("failed to get subject UUID")
	}
	return sUUID, nil
}

func GetUserIdFromUUID(ctx context.Context) (uint64, error) {
	sUUID, err := GetSubjectUUID(ctx)
	if err != nil {
		return 0, err
	}
	// Get infoUser Redis from uuid
	var inforUser InfoUserUUID
	if err := cache.GetCache(ctx, sUUID, &inforUser); err != nil {
		return 0, err
	}
	return inforUser.UserID, nil
}
