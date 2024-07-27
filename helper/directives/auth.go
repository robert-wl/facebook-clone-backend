package directives

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/yahkerobertkertasnya/TPAWebBack/database/postgresql"
	"github.com/yahkerobertkertasnya/TPAWebBack/database/redis"
	"github.com/yahkerobertkertasnya/TPAWebBack/graph/model"
	"github.com/yahkerobertkertasnya/TPAWebBack/helper"
	"time"
)

func AuthDirectives(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	token := ctx.Value("token")

	if token == nil {
		return nil, &gqlerror.Error{
			Message: "Permission denied",
		}
	}

	claims, err := helper.ParseJWT(token.(string))

	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
		}
	}

	userId := claims

	client := redis.GetInstance()

	if _, err := client.Get(ctx, fmt.Sprintf(`user:%s`, userId)).Result(); err != nil {
		var user model.User
		DB := postgresql.GetInstance()

		if err := DB.First(&user, "id = ?", userId).Error; err != nil {
			return nil, err
		}

		serializedUser, err := json.Marshal(user)

		if err != nil {
			return nil, err
		}

		client.Set(ctx, fmt.Sprintf(`user:%s`, userId), serializedUser, 10*60*time.Minute)
	}

	ctx = context.WithValue(ctx, "UserID", userId)

	return next(ctx)
}
