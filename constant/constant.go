package constant

import "time"

type FeatureName string

const (
	FeatureNameSF      FeatureName = "SF"
	FeatureNameExistor FeatureName = "EXISTOR"

	EventSkuFeaturesUpsert string = "SKU_FEATURES_UPSERT"
	EventWatermark         string = "WATERMARK"
	EventExistUpdate       string = "EXISTOR_UPDATE"

	ExpiryInterval time.Duration = 10 * time.Minute

	ORDER_BOOL_SUFFIX = "_order"
	IMPR_BOOL_SUFFIX  = "_impr"
)

var (
	EXCLUDE_METHOD_BLOCK_GRPC = map[string]bool{"HeartBeat": true, "Lock": true, "Release": true}
	EXCLUDE_METHOD_BLOCK_HTTP = map[string]bool{"/ping/live": true, "/ping/ready": true, "/metrics": true}
)
