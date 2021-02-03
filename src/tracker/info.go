package tracker

// TrackingInfo defines the collection of the tracking information.
// `Tracks` is an object-array saving the calling information one by one.
type TrackingInfo struct {
	Tracks []*info
}

type info struct {
	Bucket bucket

	FilePath  string
	FuncName  string
	Line      int
	Remarks   string
	Timestamp int64
	Duration  int64
}

type bucket struct {
	Name string
	Hash string
}
