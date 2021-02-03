package tracker

// GetUUID returns the tracker's UUID
func (t *Tracker) GetUUID() int64 {
	return t.uuid
}

// GetNodeID returns the current node id
func (t *Tracker) GetNodeID() int64 {
	return t.nodeID
}

// GetDepth returns how many records does the tracker have
func (t *Tracker) GetDepth() int {
	return t.depth
}

// // GetBuckets will be used to return an array of all buckets
// func (t *Tracker) GetBuckets() (buckets []bucket, err error) {
// 	if v, ok := t.ctx.Value(TrackCalling).(*TrackingInfo); ok {
// 		for _, trackInfo := range v.Tracks {
// 			buckets = append(buckets, trackInfo.Bucket)
// 		}
// 		return
// 	}

// 	if len(buckets) == 0 {
// 		err = ErrNoBuckets
// 		return
// 	}

// 	err = ErrGeneral
// 	return
// }
