package storage

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func fetchBucketPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	bkt := parent.Get("name").(*schema.Text).Str
	storageClient, err := storage.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	output, err := storageClient.Bucket(bkt).IAM().V3().Policy(ctx)
	if err != nil {
		return err
	}
	res <- output
	return nil
}
