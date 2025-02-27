package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func TeamAppPermissions() *schema.Table {
	return &schema.Table{
		Name:        "heroku_team_app_permissions",
		Description: `https://devcenter.heroku.com/articles/platform-api-reference#team-app-permission`,
		Resolver:    fetchTeamAppPermissions,
		Transform:   transformers.TransformWithStruct(&heroku.TeamAppPermission{}),
		Columns:     []schema.Column{},
	}
}

func fetchTeamAppPermissions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange) // nolint:revive,staticcheck
		v, err := c.Heroku.TeamAppPermissionList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}
