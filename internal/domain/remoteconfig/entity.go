package remoteconfig

type RemoteConfigQuery struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type APKVersionQuery struct {
	ApkVersion string `json:"apk-version"`
}
