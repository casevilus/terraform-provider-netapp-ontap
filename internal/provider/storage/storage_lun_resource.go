package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/netapp/terraform-provider-netapp-ontap/internal/provider/connection"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netapp/terraform-provider-netapp-ontap/internal/interfaces"
	"github.com/netapp/terraform-provider-netapp-ontap/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &StorageLunResource{}
var _ resource.ResourceWithImportState = &StorageLunResource{}

// NewStorageLunResource is a helper function to simplify the provider implementation.
func NewStorageLunResource() resource.Resource {
	return &StorageLunResource{
		config: connection.ResourceOrDataSourceConfig{
			Name: "lun",
		},
	}
}

// NewStorageLunResourceAlias is a helper function to simplify the provider implementation.
func NewStorageLunResourceAlias() resource.Resource {
	return &StorageLunResource{
		config: connection.ResourceOrDataSourceConfig{
			Name: "storage_lun_resource",
		},
	}
}

// StorageLunResource defines the resource implementation.
type StorageLunResource struct {
	config connection.ResourceOrDataSourceConfig
}

// StorageLunResourceModel describes the resource data model.
type StorageLunResourceModel struct {
	CxProfileName types.String `tfsdk:"cx_profile_name"`
	Name          types.String `tfsdk:"name"`
	SVMName       types.String `tfsdk:"svm_name"`
	VolumeName    types.String `tfsdk:"volume_name"`
	OSType        types.String `tfsdk:"os_type"`
	Size          types.Int64  `tfsdk:"size"`
	SizeUnit      types.String `tfsdk:"size_unit"`
	QoSPolicyName types.String `tfsdk:"qos_policy_name"`
	SerialNumber  types.String `tfsdk:"serial_number"`
	LogicalUnit   types.String `tfsdk:"logical_unit"`
	ID            types.String `tfsdk:"id"`
}

// Metadata returns the resource type name.
func (r *StorageLunResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.config.Name
}

// Schema defines the schema for the resource.
func (r *StorageLunResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "StorageLun resource",

		Attributes: map[string]schema.Attribute{
			"cx_profile_name": schema.StringAttribute{
				MarkdownDescription: "Connection profile name",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Lun name",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"logical_unit": schema.StringAttribute{
				MarkdownDescription: "Logical unit for lun",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"svm_name": schema.StringAttribute{
				MarkdownDescription: "SVM name",
				Required:            true,
			},
			"volume_name": schema.StringAttribute{
				MarkdownDescription: "Volume name",
				Required:            true,
			},
			"os_type": schema.StringAttribute{
				MarkdownDescription: "OS type",
				Required:            true,
			},
			"size": schema.Int64Attribute{
				MarkdownDescription: "Size of the lun in byte if size_unit is not provided, otherwise size in the specified unit",
				Required:            true,
			},
			"size_unit": schema.StringAttribute{
				MarkdownDescription: "The unit used to interpret the size parameter",
				Optional:            true,
			},
			"qos_policy_name": schema.StringAttribute{
				MarkdownDescription: "QoS policy name",
				Optional:            true,
			},
			"serial_number": schema.StringAttribute{
				MarkdownDescription: "Serial number for lun",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "StorageLun UUID",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *StorageLunResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	config, ok := req.ProviderData.(connection.Config)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected Config, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
	}
	r.config.ProviderConfig = config
}

// Read refreshes the Terraform state with the latest data.
func (r *StorageLunResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data StorageLunResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	errorHandler := utils.NewErrorHandler(ctx, &resp.Diagnostics)
	// we need to defer setting the client until we can read the connection profile name
	client, err := connection.GetRestClient(errorHandler, r.config, data.CxProfileName)
	if err != nil {
		// error reporting done inside NewClient
		return
	}

	var restInfo *interfaces.StorageLunGetDataModelONTAP
	if data.ID.ValueString() != "" {
		restInfo, err = interfaces.GetStorageLunByUUID(errorHandler, *client, data.ID.ValueString())
		if err != nil {
			// error reporting done inside GetStorageLunByUUID
			return
		}
	} else {
		restInfo, err = interfaces.GetStorageLunByName(errorHandler, *client, data.Name.ValueString(), data.SVMName.ValueString(), data.VolumeName.ValueString())
		if err != nil {
			// error reporting done inside GetStorageLunByName
			return
		}
	}

	if restInfo == nil {
		errorHandler.MakeAndReportError("error reading info", "No Lun found")
		return
	}

	data.Name = types.StringValue(restInfo.Name)
	data.LogicalUnit = types.StringValue(restInfo.Location.LogicalUnit)
	data.ID = types.StringValue(restInfo.UUID)
	data.SVMName = types.StringValue(restInfo.SVM.Name)
	data.VolumeName = types.StringValue(restInfo.Location.Volume.Name)
	data.OSType = types.StringValue(restInfo.OSType)
	data.SerialNumber = types.StringValue(restInfo.SerialNumber)
	if !data.SizeUnit.IsNull() {
		var sizeUnit string
		var size int64
		size, sizeUnit = interfaces.ByteFormat(int64(restInfo.Space.Size))
		data.Size = types.Int64Value(size)
		data.SizeUnit = types.StringValue(sizeUnit)
	} else {
		data.Size = types.Int64Value(restInfo.Space.Size)
	}
	if restInfo.QoSPolicy.Name != "" {
		data.QoSPolicyName = types.StringValue(restInfo.QoSPolicy.Name)
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Debug(ctx, fmt.Sprintf("read a resource: %#v", data))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Create a resource and retrieve UUID
func (r *StorageLunResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *StorageLunResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	var body interfaces.StorageLunResourceBodyDataModelONTAP
	errorHandler := utils.NewErrorHandler(ctx, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	if !data.Name.IsNull() {
		body.Name = data.Name.ValueString()
	}
	if !data.LogicalUnit.IsNull() {
		body.Locations.LogicalUnit = data.LogicalUnit.ValueString()
	}
	body.Locations.Volume.Name = data.VolumeName.ValueString()
	body.SVM.Name = data.SVMName.ValueString()
	body.OsType = data.OSType.ValueString()
	if !data.SizeUnit.IsNull() {
		if _, ok := interfaces.POW2BYTEMAP[data.SizeUnit.ValueString()]; !ok {
			errorHandler.MakeAndReportError("error creating flexcache", fmt.Sprintf("invalid input for size_unit: %s, required one of: bytes, b, kb, mb, gb, tb, pb, eb, zb, yb", data.SizeUnit.ValueString()))
			return
		}
		body.Space.Size = data.Size.ValueInt64() * int64(interfaces.POW2BYTEMAP[data.SizeUnit.ValueString()])
	} else {
		body.Space.Size = data.Size.ValueInt64()
	}

	if !data.QoSPolicyName.IsNull() {
		body.QosPolicy = data.QoSPolicyName.ValueString()
	}

	client, err := connection.GetRestClient(errorHandler, r.config, data.CxProfileName)
	if err != nil {
		// error reporting done inside NewClient
		return
	}

	resource, err := interfaces.CreateStorageLun(errorHandler, *client, body)
	if err != nil {
		return
	}

	data.ID = types.StringValue(resource.UUID)
	data.SerialNumber = types.StringValue(resource.SerialNumber)
	data.LogicalUnit = types.StringValue(resource.Location.LogicalUnit)
	data.Name = types.StringValue(resource.Name)

	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *StorageLunResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state *StorageLunResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	errorHandler := utils.NewErrorHandler(ctx, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}
	client, err := connection.GetRestClient(utils.NewErrorHandler(ctx, &resp.Diagnostics), r.config, plan.CxProfileName)
	if err != nil {
		// error reporting done inside NewClient
		return
	}

	var request interfaces.StorageLunResourceBodyDataModelONTAP
	if !plan.Name.Equal(state.Name) {
		request.Name = plan.Name.ValueString()
	}
	if !plan.LogicalUnit.Equal(state.LogicalUnit) {
		request.Locations.LogicalUnit = plan.LogicalUnit.ValueString()
	}
	if !plan.VolumeName.Equal(state.VolumeName) {
		request.Locations.Volume.Name = plan.VolumeName.ValueString()
	}
	if !plan.SVMName.Equal(state.SVMName) {
		request.SVM.Name = plan.SVMName.ValueString()
	}
	if !plan.OSType.Equal(state.OSType) {
		request.OsType = plan.OSType.ValueString()
	}
	baseUnit := int64(1)
	if !plan.SizeUnit.IsNull() {
		if _, ok := interfaces.POW2BYTEMAP[plan.SizeUnit.ValueString()]; !ok {
			errorHandler.MakeAndReportError("error updating lun", fmt.Sprintf("invalid input for size_unit: %s, required one of: bytes, b, kb, mb, gb, tb, pb, eb, zb, yb", plan.SizeUnit.ValueString()))
			return
		}
		baseUnit = int64(interfaces.POW2BYTEMAP[plan.SizeUnit.ValueString()])
	}
	request.Space.Size = plan.Size.ValueInt64() * baseUnit

	if !plan.QoSPolicyName.Equal(state.QoSPolicyName) {
		request.QosPolicy = plan.QoSPolicyName.ValueString()
	}
	err = interfaces.UpdateStorageLun(errorHandler, *client, state.ID.ValueString(), request)
	if err != nil {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *StorageLunResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *StorageLunResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	errorHandler := utils.NewErrorHandler(ctx, &resp.Diagnostics)
	client, err := connection.GetRestClient(errorHandler, r.config, data.CxProfileName)
	if err != nil {
		// error reporting done inside NewClient
		return
	}

	if data.ID.IsNull() {
		errorHandler.MakeAndReportError("UUID is null", "storage_lun UUID is null")
		return
	}

	err = interfaces.DeleteStorageLun(errorHandler, *client, data.ID.ValueString())
	if err != nil {
		return
	}

}

// ImportState imports a resource using ID from terraform import command by calling the Read method.
func (r *StorageLunResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, fmt.Sprintf("import req an lun resource: %#v", req))
	idParts := strings.Split(req.ID, ",")
	if len(idParts) != 4 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" || idParts[3] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprint("Expected ID in the format 'name,volume_name,svm_name,cx_profile_name', got: ", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("volume_name"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("svm_name"), idParts[2])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cx_profile_name"), idParts[3])...)

}
