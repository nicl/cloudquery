package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Fleets() *schema.Table {
	tableName := "aws_appstream_fleets"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Fleet.html`,
		Resolver:    fetchAppstreamFleets,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appstream2"),
		Transform:   transformers.TransformWithStruct(&types.Fleet{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchAppstreamFleets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input appstream.DescribeFleetsInput
	c := meta.(*client.Client)
	svc := c.Services().Appstream
	// No paginator available
	for {
		response, err := svc.DescribeFleets(ctx, &input, func(options *appstream.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Fleets

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
