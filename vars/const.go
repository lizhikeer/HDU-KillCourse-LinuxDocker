package vars

import "github.com/corpix/uarand"

var (
	// 不需要debug的url
	NoDebugUrl = map[string]bool{
		"https://sso.hdu.edu.cn/login":                                                                   true,
		"https://newjw.hdu.edu.cn/sso/driot4login":                                                       true,
		"https://sso.hdu.edu.cn/login?service=http://newjw.hdu.edu.cn/sso/driot4login":                   true,
		"https://newjw.hdu.edu.cn/jwglxt/xtgl/login_slogin.html":                                         true,
		"https://newjw.hdu.edu.cn/jwglxt/rwlscx/rwlscx_cxRwlsIndex.html?doType=query&gnmkdm=N1548":       true, // 获取课程
		"https://newjw.hdu.edu.cn/jwglxt/xsxk/zzxkyzb_cxZzxkYzbIndex.html?gnmkdm=N253512&layout=default": true, // 选课配置
		"https://api.github.com/repos/cr4n5/HDU-KillCourse/releases/latest":                              true, // 仓库
	}
)

// Version 当前版本
const Version = "v1.4.9"

// DefaultUserAgent returns a randomly selected browser User-Agent.
func DefaultUserAgent() string {
	return uarand.GetRandom()
}
