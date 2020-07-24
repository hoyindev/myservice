package smtp

import (
	"bytes"
	"fmt"
	netsmtp "net/smtp"
	"strings"
	"time"
)

//Auth Plain Auth: sender account user name and password required
func (mail *MailConfig) Auth() {
	mail.Authorised = netsmtp.PlainAuth("", mail.User, mail.Password, mail.Host)
}

//Send send email
func (mail MailConfig) Send(message Message) error {
	mail.Auth()
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	Header := make(map[string]string)
	Header["From"] = message.From
	Header["To"] = strings.Join(message.To, ";")
	Header["Cc"] = strings.Join(message.Cc, ";")
	Header["Bcc"] = strings.Join(message.Bcc, ";")
	Header["Subject"] = message.Subject
	Header["Content-Type"] = "multipart/related;boundary=" + boundary
	Header["Date"] = time.Now().String()
	mail.writeHeader(buffer, Header)

	var template = `
    <html>
        <body>
            %s<br>
                     
        </body>
    </html>
    `
	var content = fmt.Sprintf(template, message.Body)
	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type: text/html; charset=UTF-8 \r\n"
	body += content
	buffer.WriteString(body)

	buffer.WriteString("\r\n--" + boundary + "--")
	fmt.Println(buffer.String())
	// smtp SendMail() requires all email addresses; otherwise the email will still display cc and bcc in "to" reciever inbox but other receivers don't get the email at all.
	var allMailAddressSendTo []string
	allMailAddressSendTo = append(allMailAddressSendTo, message.To...)
	allMailAddressSendTo = append(allMailAddressSendTo, message.Cc...)
	allMailAddressSendTo = append(allMailAddressSendTo, message.Bcc...)

	err := netsmtp.SendMail(mail.Host+":"+mail.Port, mail.Authorised, message.From, allMailAddressSendTo, buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (mail MailConfig) writeHeader(buffer *bytes.Buffer, Header map[string]string) string {
	header := ""
	for key, value := range Header {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)
	return header
}
