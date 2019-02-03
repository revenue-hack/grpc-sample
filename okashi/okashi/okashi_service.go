package okashi

import (
	"errors"
	"math/rand"
	"time"

	context "golang.org/x/net/context"
)

type OkashiService struct {
}

func (s *OkashiService) GetFavoriteOkashi(ctx context.Context, message *GetOkashiMessage) (*OkashiResponse, error) {
	name := message.Name
	likeFreq := message.LikeFreq
	rand.Seed(time.Now().UnixNano())
	if name == "" || likeFreq == 0 {
		return nil, errors.New("invalid parameter")
	}
	return &OkashiResponse{
		Name: name,
		Eval: likeFreq * int32(rand.Intn(10)),
	}, nil
}
