package sts

//https://cloud.tencent.com/document/product/598/33416
//使用子账户模式，请先通过主账户授权开通 QcloudTCBFullAccess(云开发全读写访问), QcloudAccessForTCBRole(云开发对云资源的访问权限)
//https://cloud.tencent.com/document/product/598/36256
type Config struct {
	SecretId        string //访问管理密钥ID
	SecretKey       string //访问管理密钥KEY
	Region          string //区域 https://cloud.tencent.com/document/product/614/18940
	Name            string //您可以自定义调用方英文名称，由字母组成。
	Policy          string //授予该临时证书权限的CAM策略  https://cloud.tencent.com/document/product/598/10603
	DurationSeconds uint64 //指定临时证书的有效期，单位：秒，默认1800秒，主账号最长可设定有效期为7200秒，子账号最长可设定有效期为129600秒。
	Debug           bool   //debug
}
