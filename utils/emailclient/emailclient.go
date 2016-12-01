package emailclient

import (
	cfg "ManageCenter/config"
	"fmt"

	log "github.com/inconshreveable/log15"
	"gopkg.in/gomail.v2"
)

var (
	mailUsername = cfg.Cfg.Mail.Username
	mailPassword = cfg.Cfg.Mail.Password
	mailHost     = cfg.Cfg.Mail.Host
	mailPort     = cfg.Cfg.Mail.Port
	S            gomail.SendCloser
)

func init() {

	var d = gomail.NewDialer(mailHost, mailPort, mailUsername, mailPassword)
	var err error
	S, err = d.Dial()
	if err != nil {
		log.Error(fmt.Sprintf("mail err%v", err))
	}
}

func SendPlainMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "消息推广!")
	m.SetBody("text/plain", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send plain mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}

/*
	verify_mail_body := "<p style=\" font-size:16px\">尊敬的开发者，您好:</p>" +
				"<p>感谢您使用深图的服务。</p>" +
				"<p>您之所以收到邮件，是因为您填写了邮件地址，我们需要验证您邮箱地址是真实有效的，<a href=\"" + cfg.Cfg.WebsiteDomain + "check_email?name=" + username + "&checkStr=" + checkStr + "\">请点击这里进行有效性验证</a>。验证之后的邮箱是可以用来当用户名登录的。</p>" +
				"<p>或者复制以下地址，粘贴到浏览器的地址栏来进行验证：</p>" +
				"<a href=\"" + cfg.Cfg.WebsiteDomain + "check_email?name=" + username + "&checkStr=" + checkStr + "\">" + cfg.Cfg.WebsiteDomain + "check_email?name=" + username + "&checkStr=" + checkStr + "</a>" +
				"<p>该链接有效期为24个小时，请及时验证）如果该操作不是您本人进行，请忽略这封邮件。</p>" +
				"<p>如果您有其他需求，请随时与我们联系：bd@deepir.com</p>" +
				"<p>深图团队</p>"
*/
func SendHtmlMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "消息推广!")
	m.SetBody("text/html", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send html mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}

func SendBillMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "账单生成通知!")
	m.SetBody("text/plain", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send generate the bill mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}

//email_body := "尊敬的" + username + "，您的账户已经成功变更成" + utype + "，有任何问题，请与我们联系"
func SendPackageApplyMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "用户类型变更通知!")
	m.SetBody("text/plain", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send package apply mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}
