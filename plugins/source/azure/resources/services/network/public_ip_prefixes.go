package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func PublicIpPrefixes() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_public_ip_prefixes",
		Resolver:    fetchPublicIpPrefixes,
		Description: "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/public-ip-prefixes/list?tabs=HTTP#publicipprefix",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_network_public_ip_prefixes", client.Namespacemicrosoft_network),
		Transform:   transformers.TransformWithStruct(&armnetwork.PublicIPPrefix{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchPublicIpPrefixes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewPublicIPPrefixesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListAllPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
