package opensds

import (
	"log"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"golang.org/x/net/context"
)

////////////////////////////////////////////////////////////////////////////////
//                            Controller Service                              //
////////////////////////////////////////////////////////////////////////////////

// CreateVolume implementation
func (p *Plugin) CreateVolume(
	ctx context.Context,
	req *csi.CreateVolumeRequest) (
	*csi.CreateVolumeResponse, error) {
	// TODO
	return nil, nil
}

// DeleteVolume implementation
func (p *Plugin) DeleteVolume(
	ctx context.Context,
	req *csi.DeleteVolumeRequest) (
	*csi.DeleteVolumeResponse, error) {
	// TODO
	return nil, nil
}

// ControllerPublishVolume implementation
func (p *Plugin) ControllerPublishVolume(
	ctx context.Context,
	req *csi.ControllerPublishVolumeRequest) (
	*csi.ControllerPublishVolumeResponse, error) {
	// TODO
	return nil, nil
}

// ControllerUnpublishVolume implementation
func (p *Plugin) ControllerUnpublishVolume(
	ctx context.Context,
	req *csi.ControllerUnpublishVolumeRequest) (
	*csi.ControllerUnpublishVolumeResponse, error) {
	// TODO
	return nil, nil
}

// ValidateVolumeCapabilities implementation
func (p *Plugin) ValidateVolumeCapabilities(
	ctx context.Context,
	req *csi.ValidateVolumeCapabilitiesRequest) (
	*csi.ValidateVolumeCapabilitiesResponse, error) {

	log.Println("start to ValidateVolumeCapabilities")
	defer log.Println("end to ValidateVolumeCapabilities")

	volumeid := req.VolumeInfo.Handle.Id
	if strings.TrimSpace(volumeid) == "" {
		return &csi.ValidateVolumeCapabilitiesResponse{
			Reply: &csi.ValidateVolumeCapabilitiesResponse_Error{
				Error: &csi.Error{
					Value: &csi.Error_ValidateVolumeCapabilitiesError_{
						ValidateVolumeCapabilitiesError: &csi.Error_ValidateVolumeCapabilitiesError{
							ErrorCode:        csi.Error_ValidateVolumeCapabilitiesError_INVALID_VOLUME_INFO,
							ErrorDescription: "invalid volume id",
						},
					},
				},
			},
		}, nil
	}

	for _, capabilities := range req.VolumeCapabilities {
		if capabilities.GetMount() != nil {
			return &csi.ValidateVolumeCapabilitiesResponse{
				Reply: &csi.ValidateVolumeCapabilitiesResponse_Result_{
					Result: &csi.ValidateVolumeCapabilitiesResponse_Result{
						Supported: false,
						Message:   "opensds does not support mounted volume",
					},
				},
			}, nil
		}
	}

	return &csi.ValidateVolumeCapabilitiesResponse{
		Reply: &csi.ValidateVolumeCapabilitiesResponse_Result_{
			Result: &csi.ValidateVolumeCapabilitiesResponse_Result{
				Supported: true,
				Message:   "supported",
			},
		},
	}, nil
}

// ListVolumes implementation
func (p *Plugin) ListVolumes(
	ctx context.Context,
	req *csi.ListVolumesRequest) (
	*csi.ListVolumesResponse, error) {
	// TODO
	return nil, nil
}

// GetCapacity implementation
func (p *Plugin) GetCapacity(
	ctx context.Context,
	req *csi.GetCapacityRequest) (
	*csi.GetCapacityResponse, error) {
	// TODO
	return nil, nil
}

// ControllerGetCapabilities implementation
func (p *Plugin) ControllerGetCapabilities(
	ctx context.Context,
	req *csi.ControllerGetCapabilitiesRequest) (
	*csi.ControllerGetCapabilitiesResponse, error) {

	log.Println("start to ControllerGetCapabilities")
	defer log.Println("end to ControllerGetCapabilities")

	return &csi.ControllerGetCapabilitiesResponse{
		Reply: &csi.ControllerGetCapabilitiesResponse_Result_{
			Result: &csi.ControllerGetCapabilitiesResponse_Result{
				Capabilities: []*csi.ControllerServiceCapability{
					&csi.ControllerServiceCapability{
						Type: &csi.ControllerServiceCapability_Rpc{
							Rpc: &csi.ControllerServiceCapability_RPC{
								Type: csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
							},
						},
					},
					&csi.ControllerServiceCapability{
						Type: &csi.ControllerServiceCapability_Rpc{
							Rpc: &csi.ControllerServiceCapability_RPC{
								Type: csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
							},
						},
					},
					&csi.ControllerServiceCapability{
						Type: &csi.ControllerServiceCapability_Rpc{
							Rpc: &csi.ControllerServiceCapability_RPC{
								Type: csi.ControllerServiceCapability_RPC_LIST_VOLUMES,
							},
						},
					},
					&csi.ControllerServiceCapability{
						Type: &csi.ControllerServiceCapability_Rpc{
							Rpc: &csi.ControllerServiceCapability_RPC{
								Type: csi.ControllerServiceCapability_RPC_GET_CAPACITY,
							},
						},
					},
				},
			},
		},
	}, nil
}
