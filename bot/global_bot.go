package bot

import (
	"os"
	"time"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/spf13/viper"
	"github.com/sslime336/awbot/logging"
)

type GlobalBot struct {
	*client.QQClient

	CronList map[string]func()
}

var Bot = &GlobalBot{
	QQClient: client.NewClient(viper.GetInt64("bot.uin"), viper.GetString("bot.password")),
	CronList: make(map[string]func(), 3),
}

const _deviceInfo = "device.json"
const _sessionToken = "session.token"
const _qrcode = "qrcode.png"

func Setup() {
	// 生成设备文件
	client.GenRandomDevice()
	err := os.WriteFile(_deviceInfo, client.SystemDeviceInfo.ToJson(), 0o644)
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}

	// 读取设备文件中的配置信息
	b, err := os.ReadFile(_deviceInfo)
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}
	err = client.SystemDeviceInfo.ReadJson(b)
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}

	loginbot()
}

func loginbot() {
	// 若存在合法的登录 token
	token, err := os.ReadFile(_sessionToken)
	if err == nil && token != nil {
		err = Bot.TokenLogin(token)
		// token 登录失败，转为二维码登录
		if err != nil {
			_ = os.Remove(_sessionToken)
			Bot.Disconnect()
			Bot.Release()
			logging.Logger.Error(err.Error())

			goto qrLogin
		}

		// 保存登录后的 token: 不知道登录后会不会刷新 token 所以干脆保存一次
		storeToken()
	}

	// 如果 Bot 已经成功上线
	if Bot.Online.Load() {
		return
	}

	logging.Logger.Info("未检测到 token 将使用二维码登录")

	// 采用二维码登录，二维码会生成在项目根目录
qrLogin:
	qrcodeLogin()
}

func storeToken() {
	tokenContent := Bot.GenToken()
	_ = os.WriteFile(_sessionToken, tokenContent, 0o644)
}

func qrcodeLogin() {
	resp, err := Bot.FetchQRCode()
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}

	// 保存验证码
	_ = os.WriteFile(_qrcode, resp.ImageData, 0o644)
	defer func() {
		_ = os.Remove(_qrcode)
	}()
	logging.Logger.Info("已生成登录所用二维码，请及时扫码")

	s, err := Bot.QueryQRCodeStatus(resp.Sig)
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}
	prevState := s.State
	// 等待用户扫码登录
success:
	for {
		time.Sleep(time.Second)
		s, _ = Bot.QueryQRCodeStatus(resp.Sig)
		if s == nil {
			continue
		}
		if prevState == s.State {
			continue
		}
		prevState = s.State
		switch s.State {
		case client.QRCodeCanceled:
			logging.Logger.Fatal("扫码被用户取消")
		case client.QRCodeTimeout:
			logging.Logger.Fatal("二维码过期")
		case client.QRCodeWaitingForConfirm:
			logging.Logger.Info("扫码成功, 请在手机端确认登录")
		case client.QRCodeConfirmed:
			res, err := Bot.QRCodeLogin(s.LoginInfo)
			if err != nil {
				logging.Logger.Fatal(err.Error())
			}
			if !res.Success {
				logging.Logger.Fatal(res.ErrorMessage)
			}
			break success
		default: // client.QRCodeImageFetch, client.QRCodeWaitingForScan:
		}
	}

	storeToken()
}
