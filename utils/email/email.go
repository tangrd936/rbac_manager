package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"rbac_manager/global"
)

func SendMail(mailTo []string, subject string, body string) error {
	mailConf := global.Conf.Email
	// 设置smtp服务器配置
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(mailConf.User, "测试")) //这种方式可以添加别名，即“测试”
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)    //收件人邮箱，可以发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConf.Host, mailConf.Port, mailConf.User, mailConf.Pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // 忽略证书校验，仅用于测试环境

	err := d.DialAndSend(m)
	if err != nil {
		global.Log.Error("发送邮件失败：" + err.Error())
		return err
	}
	global.Log.Info("send email successfully")
	return nil
}
