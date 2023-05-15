// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rediscloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Database within a specified Active-Active Subscription in your Redis Enterprise Cloud Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/RedisLabs/pulumi-rediscloud/sdk/go/rediscloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			card, err := rediscloud.GetPaymentMethod(ctx, &rediscloud.GetPaymentMethodArgs{
//				CardType: pulumi.StringRef("Visa"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = rediscloud.NewActiveActiveSubscription(ctx, "subscription-resource", &rediscloud.ActiveActiveSubscriptionArgs{
//				PaymentMethodId: *pulumi.String(card.Id),
//				CloudProvider:   pulumi.String("AWS"),
//				CreationPlan: &rediscloud.ActiveActiveSubscriptionCreationPlanArgs{
//					MemoryLimitInGb: pulumi.Float64(1),
//					Quantity:        pulumi.Int(1),
//					Regions: rediscloud.ActiveActiveSubscriptionCreationPlanRegionArray{
//						&rediscloud.ActiveActiveSubscriptionCreationPlanRegionArgs{
//							Region:                   pulumi.String("us-east-1"),
//							NetworkingDeploymentCidr: pulumi.String("192.168.0.0/24"),
//							WriteOperationsPerSecond: pulumi.Int(1000),
//							ReadOperationsPerSecond:  pulumi.Int(1000),
//						},
//						&rediscloud.ActiveActiveSubscriptionCreationPlanRegionArgs{
//							Region:                   pulumi.String("us-east-2"),
//							NetworkingDeploymentCidr: pulumi.String("10.0.1.0/24"),
//							WriteOperationsPerSecond: pulumi.Int(1000),
//							ReadOperationsPerSecond:  pulumi.Int(2000),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rediscloud.NewActiveActiveSubscriptionDatabase(ctx, "database-resource", &rediscloud.ActiveActiveSubscriptionDatabaseArgs{
//				SubscriptionId:        subscription_resource.ID(),
//				MemoryLimitInGb:       pulumi.Float64(1),
//				GlobalDataPersistence: pulumi.String("aof-every-1-second"),
//				GlobalPassword:        pulumi.String("some-random-pass-2"),
//				GlobalSourceIps: pulumi.StringArray{
//					pulumi.String("192.168.0.0/16"),
//				},
//				GlobalAlerts: rediscloud.ActiveActiveSubscriptionDatabaseGlobalAlertArray{
//					&rediscloud.ActiveActiveSubscriptionDatabaseGlobalAlertArgs{
//						Name:  pulumi.String("dataset-size"),
//						Value: pulumi.Int(40),
//					},
//				},
//				OverrideRegions: rediscloud.ActiveActiveSubscriptionDatabaseOverrideRegionArray{
//					&rediscloud.ActiveActiveSubscriptionDatabaseOverrideRegionArgs{
//						Name: pulumi.String("us-east-2"),
//						OverrideGlobalSourceIps: pulumi.StringArray{
//							pulumi.String("192.10.0.0/16"),
//						},
//					},
//					&rediscloud.ActiveActiveSubscriptionDatabaseOverrideRegionArgs{
//						Name:                          pulumi.String("us-east-1"),
//						OverrideGlobalDataPersistence: pulumi.String("none"),
//						OverrideGlobalPassword:        pulumi.String("region-specific-password"),
//						OverrideGlobalAlerts: rediscloud.ActiveActiveSubscriptionDatabaseOverrideRegionOverrideGlobalAlertArray{
//							&rediscloud.ActiveActiveSubscriptionDatabaseOverrideRegionOverrideGlobalAlertArgs{
//								Name:  pulumi.String("dataset-size"),
//								Value: pulumi.Int(60),
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("us-east-1-public-endpoints", database_resource.PublicEndpoint.ApplyT(func(publicEndpoint map[string]string) (string, error) {
//				return publicEndpoint.UsEast1, nil
//			}).(pulumi.StringOutput))
//			ctx.Export("us-east-2-private-endpoints", database_resource.PrivateEndpoint.ApplyT(func(privateEndpoint map[string]string) (string, error) {
//				return privateEndpoint.UsEast1, nil
//			}).(pulumi.StringOutput))
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// `rediscloud_active_active_subscription_database` can be imported using the ID of the Active-Active subscription and the ID of the database in the format {subscription ID}/{database ID}, e.g.
//
// ```sh
//
//	$ pulumi import rediscloud:index/activeActiveSubscriptionDatabase:ActiveActiveSubscriptionDatabase database-resource 123456/12345678
//
// ```
//
//	NoteDue to constraints in the Redis Cloud API, the import process will not import global attributes or override region attributes. If you wish to use these attributes in your Terraform configuraton, you will need to manually add them to your Terraform configuration and run `terraform apply` to update the database.
type ActiveActiveSubscriptionDatabase struct {
	pulumi.CustomResourceState

	// SSL certificate to authenticate user connections.
	ClientSslCertificate pulumi.StringPtrOutput `pulumi:"clientSslCertificate"`
	// The data items eviction policy (either: 'allkeys-lru', 'allkeys-lfu', 'allkeys-random', 'volatile-lru', 'volatile-lfu', 'volatile-random', 'volatile-ttl' or 'noeviction'. Default: 'volatile-lru')
	DataEviction pulumi.StringPtrOutput `pulumi:"dataEviction"`
	// Identifier of the database created
	DbId pulumi.IntOutput `pulumi:"dbId"`
	// Use TLS for authentication. Default: ‘false’
	EnableTls pulumi.BoolPtrOutput `pulumi:"enableTls"`
	// Should use the external endpoint for open-source (OSS) Cluster API.
	// Can only be enabled if OSS Cluster API support is enabled. Default: 'false'
	ExternalEndpointForOssClusterApi pulumi.BoolPtrOutput `pulumi:"externalEndpointForOssClusterApi"`
	// A block defining Redis database alert of regions that dont override global settings, documented below, can be specified multiple times
	GlobalAlerts ActiveActiveSubscriptionDatabaseGlobalAlertArrayOutput `pulumi:"globalAlerts"`
	// Global rate of database data persistence (in persistent storage) of regions that dont override global settings. Default: 'none'
	GlobalDataPersistence pulumi.StringPtrOutput `pulumi:"globalDataPersistence"`
	// Password to access the database of regions that dont override global settings. If left empty, the password will be generated automatically
	GlobalPassword pulumi.StringOutput `pulumi:"globalPassword"`
	// List of source IP addresses or subnet masks of regions that dont override global settings. If specified, Redis clients will be able to connect to this database only from within the specified source IP addresses ranges (example: ['192.168.10.0/32', '192.168.12.0/24'])
	GlobalSourceIps pulumi.StringArrayOutput `pulumi:"globalSourceIps"`
	// Maximum memory usage for this specific database, including replication and other overhead
	MemoryLimitInGb pulumi.Float64Output `pulumi:"memoryLimitInGb"`
	// A meaningful name to identify the database
	Name pulumi.StringOutput `pulumi:"name"`
	// Override region specific configuration, documented below
	OverrideRegions ActiveActiveSubscriptionDatabaseOverrideRegionArrayOutput `pulumi:"overrideRegions"`
	// A map of which private endpoints can to access the database per region, uses region name as key.
	PrivateEndpoint pulumi.StringMapOutput `pulumi:"privateEndpoint"`
	// A map of which public endpoints can to access the database per region, uses region name as key.
	PublicEndpoint pulumi.StringMapOutput `pulumi:"publicEndpoint"`
	// The ID of the Active-Active subscription to create the database in
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// Support Redis open-source (OSS) Cluster API. Default: ‘false’
	SupportOssClusterApi pulumi.BoolPtrOutput `pulumi:"supportOssClusterApi"`
}

// NewActiveActiveSubscriptionDatabase registers a new resource with the given unique name, arguments, and options.
func NewActiveActiveSubscriptionDatabase(ctx *pulumi.Context,
	name string, args *ActiveActiveSubscriptionDatabaseArgs, opts ...pulumi.ResourceOption) (*ActiveActiveSubscriptionDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MemoryLimitInGb == nil {
		return nil, errors.New("invalid value for required argument 'MemoryLimitInGb'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	if args.GlobalPassword != nil {
		args.GlobalPassword = pulumi.ToSecret(args.GlobalPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"globalPassword",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource ActiveActiveSubscriptionDatabase
	err := ctx.RegisterResource("rediscloud:index/activeActiveSubscriptionDatabase:ActiveActiveSubscriptionDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActiveActiveSubscriptionDatabase gets an existing ActiveActiveSubscriptionDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActiveActiveSubscriptionDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActiveActiveSubscriptionDatabaseState, opts ...pulumi.ResourceOption) (*ActiveActiveSubscriptionDatabase, error) {
	var resource ActiveActiveSubscriptionDatabase
	err := ctx.ReadResource("rediscloud:index/activeActiveSubscriptionDatabase:ActiveActiveSubscriptionDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActiveActiveSubscriptionDatabase resources.
type activeActiveSubscriptionDatabaseState struct {
	// SSL certificate to authenticate user connections.
	ClientSslCertificate *string `pulumi:"clientSslCertificate"`
	// The data items eviction policy (either: 'allkeys-lru', 'allkeys-lfu', 'allkeys-random', 'volatile-lru', 'volatile-lfu', 'volatile-random', 'volatile-ttl' or 'noeviction'. Default: 'volatile-lru')
	DataEviction *string `pulumi:"dataEviction"`
	// Identifier of the database created
	DbId *int `pulumi:"dbId"`
	// Use TLS for authentication. Default: ‘false’
	EnableTls *bool `pulumi:"enableTls"`
	// Should use the external endpoint for open-source (OSS) Cluster API.
	// Can only be enabled if OSS Cluster API support is enabled. Default: 'false'
	ExternalEndpointForOssClusterApi *bool `pulumi:"externalEndpointForOssClusterApi"`
	// A block defining Redis database alert of regions that dont override global settings, documented below, can be specified multiple times
	GlobalAlerts []ActiveActiveSubscriptionDatabaseGlobalAlert `pulumi:"globalAlerts"`
	// Global rate of database data persistence (in persistent storage) of regions that dont override global settings. Default: 'none'
	GlobalDataPersistence *string `pulumi:"globalDataPersistence"`
	// Password to access the database of regions that dont override global settings. If left empty, the password will be generated automatically
	GlobalPassword *string `pulumi:"globalPassword"`
	// List of source IP addresses or subnet masks of regions that dont override global settings. If specified, Redis clients will be able to connect to this database only from within the specified source IP addresses ranges (example: ['192.168.10.0/32', '192.168.12.0/24'])
	GlobalSourceIps []string `pulumi:"globalSourceIps"`
	// Maximum memory usage for this specific database, including replication and other overhead
	MemoryLimitInGb *float64 `pulumi:"memoryLimitInGb"`
	// A meaningful name to identify the database
	Name *string `pulumi:"name"`
	// Override region specific configuration, documented below
	OverrideRegions []ActiveActiveSubscriptionDatabaseOverrideRegion `pulumi:"overrideRegions"`
	// A map of which private endpoints can to access the database per region, uses region name as key.
	PrivateEndpoint map[string]string `pulumi:"privateEndpoint"`
	// A map of which public endpoints can to access the database per region, uses region name as key.
	PublicEndpoint map[string]string `pulumi:"publicEndpoint"`
	// The ID of the Active-Active subscription to create the database in
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Support Redis open-source (OSS) Cluster API. Default: ‘false’
	SupportOssClusterApi *bool `pulumi:"supportOssClusterApi"`
}

type ActiveActiveSubscriptionDatabaseState struct {
	// SSL certificate to authenticate user connections.
	ClientSslCertificate pulumi.StringPtrInput
	// The data items eviction policy (either: 'allkeys-lru', 'allkeys-lfu', 'allkeys-random', 'volatile-lru', 'volatile-lfu', 'volatile-random', 'volatile-ttl' or 'noeviction'. Default: 'volatile-lru')
	DataEviction pulumi.StringPtrInput
	// Identifier of the database created
	DbId pulumi.IntPtrInput
	// Use TLS for authentication. Default: ‘false’
	EnableTls pulumi.BoolPtrInput
	// Should use the external endpoint for open-source (OSS) Cluster API.
	// Can only be enabled if OSS Cluster API support is enabled. Default: 'false'
	ExternalEndpointForOssClusterApi pulumi.BoolPtrInput
	// A block defining Redis database alert of regions that dont override global settings, documented below, can be specified multiple times
	GlobalAlerts ActiveActiveSubscriptionDatabaseGlobalAlertArrayInput
	// Global rate of database data persistence (in persistent storage) of regions that dont override global settings. Default: 'none'
	GlobalDataPersistence pulumi.StringPtrInput
	// Password to access the database of regions that dont override global settings. If left empty, the password will be generated automatically
	GlobalPassword pulumi.StringPtrInput
	// List of source IP addresses or subnet masks of regions that dont override global settings. If specified, Redis clients will be able to connect to this database only from within the specified source IP addresses ranges (example: ['192.168.10.0/32', '192.168.12.0/24'])
	GlobalSourceIps pulumi.StringArrayInput
	// Maximum memory usage for this specific database, including replication and other overhead
	MemoryLimitInGb pulumi.Float64PtrInput
	// A meaningful name to identify the database
	Name pulumi.StringPtrInput
	// Override region specific configuration, documented below
	OverrideRegions ActiveActiveSubscriptionDatabaseOverrideRegionArrayInput
	// A map of which private endpoints can to access the database per region, uses region name as key.
	PrivateEndpoint pulumi.StringMapInput
	// A map of which public endpoints can to access the database per region, uses region name as key.
	PublicEndpoint pulumi.StringMapInput
	// The ID of the Active-Active subscription to create the database in
	SubscriptionId pulumi.StringPtrInput
	// Support Redis open-source (OSS) Cluster API. Default: ‘false’
	SupportOssClusterApi pulumi.BoolPtrInput
}

func (ActiveActiveSubscriptionDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*activeActiveSubscriptionDatabaseState)(nil)).Elem()
}

type activeActiveSubscriptionDatabaseArgs struct {
	// SSL certificate to authenticate user connections.
	ClientSslCertificate *string `pulumi:"clientSslCertificate"`
	// The data items eviction policy (either: 'allkeys-lru', 'allkeys-lfu', 'allkeys-random', 'volatile-lru', 'volatile-lfu', 'volatile-random', 'volatile-ttl' or 'noeviction'. Default: 'volatile-lru')
	DataEviction *string `pulumi:"dataEviction"`
	// Use TLS for authentication. Default: ‘false’
	EnableTls *bool `pulumi:"enableTls"`
	// Should use the external endpoint for open-source (OSS) Cluster API.
	// Can only be enabled if OSS Cluster API support is enabled. Default: 'false'
	ExternalEndpointForOssClusterApi *bool `pulumi:"externalEndpointForOssClusterApi"`
	// A block defining Redis database alert of regions that dont override global settings, documented below, can be specified multiple times
	GlobalAlerts []ActiveActiveSubscriptionDatabaseGlobalAlert `pulumi:"globalAlerts"`
	// Global rate of database data persistence (in persistent storage) of regions that dont override global settings. Default: 'none'
	GlobalDataPersistence *string `pulumi:"globalDataPersistence"`
	// Password to access the database of regions that dont override global settings. If left empty, the password will be generated automatically
	GlobalPassword *string `pulumi:"globalPassword"`
	// List of source IP addresses or subnet masks of regions that dont override global settings. If specified, Redis clients will be able to connect to this database only from within the specified source IP addresses ranges (example: ['192.168.10.0/32', '192.168.12.0/24'])
	GlobalSourceIps []string `pulumi:"globalSourceIps"`
	// Maximum memory usage for this specific database, including replication and other overhead
	MemoryLimitInGb float64 `pulumi:"memoryLimitInGb"`
	// A meaningful name to identify the database
	Name *string `pulumi:"name"`
	// Override region specific configuration, documented below
	OverrideRegions []ActiveActiveSubscriptionDatabaseOverrideRegion `pulumi:"overrideRegions"`
	// The ID of the Active-Active subscription to create the database in
	SubscriptionId string `pulumi:"subscriptionId"`
	// Support Redis open-source (OSS) Cluster API. Default: ‘false’
	SupportOssClusterApi *bool `pulumi:"supportOssClusterApi"`
}

// The set of arguments for constructing a ActiveActiveSubscriptionDatabase resource.
type ActiveActiveSubscriptionDatabaseArgs struct {
	// SSL certificate to authenticate user connections.
	ClientSslCertificate pulumi.StringPtrInput
	// The data items eviction policy (either: 'allkeys-lru', 'allkeys-lfu', 'allkeys-random', 'volatile-lru', 'volatile-lfu', 'volatile-random', 'volatile-ttl' or 'noeviction'. Default: 'volatile-lru')
	DataEviction pulumi.StringPtrInput
	// Use TLS for authentication. Default: ‘false’
	EnableTls pulumi.BoolPtrInput
	// Should use the external endpoint for open-source (OSS) Cluster API.
	// Can only be enabled if OSS Cluster API support is enabled. Default: 'false'
	ExternalEndpointForOssClusterApi pulumi.BoolPtrInput
	// A block defining Redis database alert of regions that dont override global settings, documented below, can be specified multiple times
	GlobalAlerts ActiveActiveSubscriptionDatabaseGlobalAlertArrayInput
	// Global rate of database data persistence (in persistent storage) of regions that dont override global settings. Default: 'none'
	GlobalDataPersistence pulumi.StringPtrInput
	// Password to access the database of regions that dont override global settings. If left empty, the password will be generated automatically
	GlobalPassword pulumi.StringPtrInput
	// List of source IP addresses or subnet masks of regions that dont override global settings. If specified, Redis clients will be able to connect to this database only from within the specified source IP addresses ranges (example: ['192.168.10.0/32', '192.168.12.0/24'])
	GlobalSourceIps pulumi.StringArrayInput
	// Maximum memory usage for this specific database, including replication and other overhead
	MemoryLimitInGb pulumi.Float64Input
	// A meaningful name to identify the database
	Name pulumi.StringPtrInput
	// Override region specific configuration, documented below
	OverrideRegions ActiveActiveSubscriptionDatabaseOverrideRegionArrayInput
	// The ID of the Active-Active subscription to create the database in
	SubscriptionId pulumi.StringInput
	// Support Redis open-source (OSS) Cluster API. Default: ‘false’
	SupportOssClusterApi pulumi.BoolPtrInput
}

func (ActiveActiveSubscriptionDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activeActiveSubscriptionDatabaseArgs)(nil)).Elem()
}

type ActiveActiveSubscriptionDatabaseInput interface {
	pulumi.Input

	ToActiveActiveSubscriptionDatabaseOutput() ActiveActiveSubscriptionDatabaseOutput
	ToActiveActiveSubscriptionDatabaseOutputWithContext(ctx context.Context) ActiveActiveSubscriptionDatabaseOutput
}

func (*ActiveActiveSubscriptionDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveActiveSubscriptionDatabase)(nil)).Elem()
}

func (i *ActiveActiveSubscriptionDatabase) ToActiveActiveSubscriptionDatabaseOutput() ActiveActiveSubscriptionDatabaseOutput {
	return i.ToActiveActiveSubscriptionDatabaseOutputWithContext(context.Background())
}

func (i *ActiveActiveSubscriptionDatabase) ToActiveActiveSubscriptionDatabaseOutputWithContext(ctx context.Context) ActiveActiveSubscriptionDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveActiveSubscriptionDatabaseOutput)
}

// ActiveActiveSubscriptionDatabaseArrayInput is an input type that accepts ActiveActiveSubscriptionDatabaseArray and ActiveActiveSubscriptionDatabaseArrayOutput values.
// You can construct a concrete instance of `ActiveActiveSubscriptionDatabaseArrayInput` via:
//
//	ActiveActiveSubscriptionDatabaseArray{ ActiveActiveSubscriptionDatabaseArgs{...} }
type ActiveActiveSubscriptionDatabaseArrayInput interface {
	pulumi.Input

	ToActiveActiveSubscriptionDatabaseArrayOutput() ActiveActiveSubscriptionDatabaseArrayOutput
	ToActiveActiveSubscriptionDatabaseArrayOutputWithContext(context.Context) ActiveActiveSubscriptionDatabaseArrayOutput
}

type ActiveActiveSubscriptionDatabaseArray []ActiveActiveSubscriptionDatabaseInput

func (ActiveActiveSubscriptionDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActiveActiveSubscriptionDatabase)(nil)).Elem()
}

func (i ActiveActiveSubscriptionDatabaseArray) ToActiveActiveSubscriptionDatabaseArrayOutput() ActiveActiveSubscriptionDatabaseArrayOutput {
	return i.ToActiveActiveSubscriptionDatabaseArrayOutputWithContext(context.Background())
}

func (i ActiveActiveSubscriptionDatabaseArray) ToActiveActiveSubscriptionDatabaseArrayOutputWithContext(ctx context.Context) ActiveActiveSubscriptionDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveActiveSubscriptionDatabaseArrayOutput)
}

// ActiveActiveSubscriptionDatabaseMapInput is an input type that accepts ActiveActiveSubscriptionDatabaseMap and ActiveActiveSubscriptionDatabaseMapOutput values.
// You can construct a concrete instance of `ActiveActiveSubscriptionDatabaseMapInput` via:
//
//	ActiveActiveSubscriptionDatabaseMap{ "key": ActiveActiveSubscriptionDatabaseArgs{...} }
type ActiveActiveSubscriptionDatabaseMapInput interface {
	pulumi.Input

	ToActiveActiveSubscriptionDatabaseMapOutput() ActiveActiveSubscriptionDatabaseMapOutput
	ToActiveActiveSubscriptionDatabaseMapOutputWithContext(context.Context) ActiveActiveSubscriptionDatabaseMapOutput
}

type ActiveActiveSubscriptionDatabaseMap map[string]ActiveActiveSubscriptionDatabaseInput

func (ActiveActiveSubscriptionDatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActiveActiveSubscriptionDatabase)(nil)).Elem()
}

func (i ActiveActiveSubscriptionDatabaseMap) ToActiveActiveSubscriptionDatabaseMapOutput() ActiveActiveSubscriptionDatabaseMapOutput {
	return i.ToActiveActiveSubscriptionDatabaseMapOutputWithContext(context.Background())
}

func (i ActiveActiveSubscriptionDatabaseMap) ToActiveActiveSubscriptionDatabaseMapOutputWithContext(ctx context.Context) ActiveActiveSubscriptionDatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveActiveSubscriptionDatabaseMapOutput)
}

type ActiveActiveSubscriptionDatabaseOutput struct{ *pulumi.OutputState }

func (ActiveActiveSubscriptionDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveActiveSubscriptionDatabase)(nil)).Elem()
}

func (o ActiveActiveSubscriptionDatabaseOutput) ToActiveActiveSubscriptionDatabaseOutput() ActiveActiveSubscriptionDatabaseOutput {
	return o
}

func (o ActiveActiveSubscriptionDatabaseOutput) ToActiveActiveSubscriptionDatabaseOutputWithContext(ctx context.Context) ActiveActiveSubscriptionDatabaseOutput {
	return o
}

// SSL certificate to authenticate user connections.
func (o ActiveActiveSubscriptionDatabaseOutput) ClientSslCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) pulumi.StringPtrOutput { return v.ClientSslCertificate }).(pulumi.StringPtrOutput)
}

// The data items eviction policy (either: 'allkeys-lru', 'allkeys-lfu', 'allkeys-random', 'volatile-lru', 'volatile-lfu', 'volatile-random', 'volatile-ttl' or 'noeviction'. Default: 'volatile-lru')
func (o ActiveActiveSubscriptionDatabaseOutput) DataEviction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) pulumi.StringPtrOutput { return v.DataEviction }).(pulumi.StringPtrOutput)
}

// Identifier of the database created
func (o ActiveActiveSubscriptionDatabaseOutput) DbId() pulumi.IntOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) pulumi.IntOutput { return v.DbId }).(pulumi.IntOutput)
}

// Use TLS for authentication. Default: ‘false’
func (o ActiveActiveSubscriptionDatabaseOutput) EnableTls() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) pulumi.BoolPtrOutput { return v.EnableTls }).(pulumi.BoolPtrOutput)
}

// Should use the external endpoint for open-source (OSS) Cluster API.
// Can only be enabled if OSS Cluster API support is enabled. Default: 'false'
func (o ActiveActiveSubscriptionDatabaseOutput) ExternalEndpointForOssClusterApi() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) pulumi.BoolPtrOutput {
		return v.ExternalEndpointForOssClusterApi
	}).(pulumi.BoolPtrOutput)
}

// A block defining Redis database alert of regions that dont override global settings, documented below, can be specified multiple times
func (o ActiveActiveSubscriptionDatabaseOutput) GlobalAlerts() ActiveActiveSubscriptionDatabaseGlobalAlertArrayOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) ActiveActiveSubscriptionDatabaseGlobalAlertArrayOutput {
		return v.GlobalAlerts
	}).(ActiveActiveSubscriptionDatabaseGlobalAlertArrayOutput)
}

// Global rate of database data persistence (in persistent storage) of regions that dont override global settings. Default: 'none'
func (o ActiveActiveSubscriptionDatabaseOutput) GlobalDataPersistence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) pulumi.StringPtrOutput { return v.GlobalDataPersistence }).(pulumi.StringPtrOutput)
}

// Password to access the database of regions that dont override global settings. If left empty, the password will be generated automatically
func (o ActiveActiveSubscriptionDatabaseOutput) GlobalPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) pulumi.StringOutput { return v.GlobalPassword }).(pulumi.StringOutput)
}

// List of source IP addresses or subnet masks of regions that dont override global settings. If specified, Redis clients will be able to connect to this database only from within the specified source IP addresses ranges (example: ['192.168.10.0/32', '192.168.12.0/24'])
func (o ActiveActiveSubscriptionDatabaseOutput) GlobalSourceIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) pulumi.StringArrayOutput { return v.GlobalSourceIps }).(pulumi.StringArrayOutput)
}

// Maximum memory usage for this specific database, including replication and other overhead
func (o ActiveActiveSubscriptionDatabaseOutput) MemoryLimitInGb() pulumi.Float64Output {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) pulumi.Float64Output { return v.MemoryLimitInGb }).(pulumi.Float64Output)
}

// A meaningful name to identify the database
func (o ActiveActiveSubscriptionDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Override region specific configuration, documented below
func (o ActiveActiveSubscriptionDatabaseOutput) OverrideRegions() ActiveActiveSubscriptionDatabaseOverrideRegionArrayOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) ActiveActiveSubscriptionDatabaseOverrideRegionArrayOutput {
		return v.OverrideRegions
	}).(ActiveActiveSubscriptionDatabaseOverrideRegionArrayOutput)
}

// A map of which private endpoints can to access the database per region, uses region name as key.
func (o ActiveActiveSubscriptionDatabaseOutput) PrivateEndpoint() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) pulumi.StringMapOutput { return v.PrivateEndpoint }).(pulumi.StringMapOutput)
}

// A map of which public endpoints can to access the database per region, uses region name as key.
func (o ActiveActiveSubscriptionDatabaseOutput) PublicEndpoint() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) pulumi.StringMapOutput { return v.PublicEndpoint }).(pulumi.StringMapOutput)
}

// The ID of the Active-Active subscription to create the database in
func (o ActiveActiveSubscriptionDatabaseOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

// Support Redis open-source (OSS) Cluster API. Default: ‘false’
func (o ActiveActiveSubscriptionDatabaseOutput) SupportOssClusterApi() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ActiveActiveSubscriptionDatabase) pulumi.BoolPtrOutput { return v.SupportOssClusterApi }).(pulumi.BoolPtrOutput)
}

type ActiveActiveSubscriptionDatabaseArrayOutput struct{ *pulumi.OutputState }

func (ActiveActiveSubscriptionDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActiveActiveSubscriptionDatabase)(nil)).Elem()
}

func (o ActiveActiveSubscriptionDatabaseArrayOutput) ToActiveActiveSubscriptionDatabaseArrayOutput() ActiveActiveSubscriptionDatabaseArrayOutput {
	return o
}

func (o ActiveActiveSubscriptionDatabaseArrayOutput) ToActiveActiveSubscriptionDatabaseArrayOutputWithContext(ctx context.Context) ActiveActiveSubscriptionDatabaseArrayOutput {
	return o
}

func (o ActiveActiveSubscriptionDatabaseArrayOutput) Index(i pulumi.IntInput) ActiveActiveSubscriptionDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ActiveActiveSubscriptionDatabase {
		return vs[0].([]*ActiveActiveSubscriptionDatabase)[vs[1].(int)]
	}).(ActiveActiveSubscriptionDatabaseOutput)
}

type ActiveActiveSubscriptionDatabaseMapOutput struct{ *pulumi.OutputState }

func (ActiveActiveSubscriptionDatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActiveActiveSubscriptionDatabase)(nil)).Elem()
}

func (o ActiveActiveSubscriptionDatabaseMapOutput) ToActiveActiveSubscriptionDatabaseMapOutput() ActiveActiveSubscriptionDatabaseMapOutput {
	return o
}

func (o ActiveActiveSubscriptionDatabaseMapOutput) ToActiveActiveSubscriptionDatabaseMapOutputWithContext(ctx context.Context) ActiveActiveSubscriptionDatabaseMapOutput {
	return o
}

func (o ActiveActiveSubscriptionDatabaseMapOutput) MapIndex(k pulumi.StringInput) ActiveActiveSubscriptionDatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ActiveActiveSubscriptionDatabase {
		return vs[0].(map[string]*ActiveActiveSubscriptionDatabase)[vs[1].(string)]
	}).(ActiveActiveSubscriptionDatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActiveActiveSubscriptionDatabaseInput)(nil)).Elem(), &ActiveActiveSubscriptionDatabase{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActiveActiveSubscriptionDatabaseArrayInput)(nil)).Elem(), ActiveActiveSubscriptionDatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActiveActiveSubscriptionDatabaseMapInput)(nil)).Elem(), ActiveActiveSubscriptionDatabaseMap{})
	pulumi.RegisterOutputType(ActiveActiveSubscriptionDatabaseOutput{})
	pulumi.RegisterOutputType(ActiveActiveSubscriptionDatabaseArrayOutput{})
	pulumi.RegisterOutputType(ActiveActiveSubscriptionDatabaseMapOutput{})
}
