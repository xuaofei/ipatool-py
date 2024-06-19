package main

// 请求任务
type taskInfoRequest struct {
	AppleID     string `json:"apple_id"`
	ApplePwd    string `json:"apple_pwd"`
	AppBundleID string `json:"app_bundle_id"`
	Country     string `json:"country"`
	TaskID      string `json:"task_id"`
}

type taskInfoResponse struct {
	AppleID     string `json:"apple_id"`
	ApplePwd    string `json:"apple_pwd"`
	AppBundleID string `json:"app_bundle_id"`
	Country     string `json:"country"`
	TaskID      string `json:"task_id"`
}

// 请求二次验证码
type twoFAInfoRequest struct {
	AppleID string `json:"apple_id"`
	TaskID  string `json:"task_id"`
}

type twoFAInfoResponse struct {
	TwoFACode string `json:"two_fa_code"`
}

// 上报ipa版本信息
type ipaVersionsRequest struct {
	AppleID        string `json:"apple_id"`
	TaskID         string `json:"task_id"`
	VersionCount   int    `json:"version_count"`
	AllVersionList []struct {
		AppVer                   string `json:"app_ver"`
		AppVerID                 string `json:"app_ver_id"`
		BundleShortVersionString string `json:"bundle_short_version_string"`
	} `json:"all_version_list"`
}

// 要下载的ipa版本
type downloadIpaVersionsRequest struct {
	AppleID string `json:"apple_id"`
	TaskID  string `json:"task_id"`
}

type downloadIpaVersionsResponse struct {
	DownloadVersionList struct {
		AppVer                   string `json:"app_ver"`
		AppVerID                 string `json:"app_ver_id"`
		BundleShortVersionString string `json:"bundle_short_version_string"`
	} `json:"download_version_list"`
}
