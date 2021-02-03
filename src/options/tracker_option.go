package options

// TrackerOption defines a set of options for tracker uses
type TrackerOption struct {
	NodeID *int64

	GRPCOption *GRPCConfigs
}

// MergeTrackerOptions should be used for merging opts for tracker
func MergeTrackerOptions(opts ...*TrackerOption) *TrackerOption {
	var o = new(TrackerOption)

	if len(opts) == 0 {
		return o
	}
	for _, opt := range opts {
		if opt == nil {
			continue
		}

		if opt.NodeID != nil {
			o.NodeID = opt.NodeID
		}

		if opt.GRPCOption != nil {
			o.GRPCOption = opt.GRPCOption
		}
	}

	return o
}
