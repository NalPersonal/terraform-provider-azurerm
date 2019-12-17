package client

import (
	appconf "github.com/Azure/azure-sdk-for-go/services/appconfiguration/mgmt/2019-10-01/appconfiguration"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	AppConfigurationsClient *appconf.ConfigurationStoresClient
}

func NewClient(o *common.ClientOptions) *Client {
	appConfigurationsClient := appconf.NewConfigurationStoresClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&appConfigurationsClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		AppConfigurationsClient: &appConfigurationsClient,
	}
}
