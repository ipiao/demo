package designPattern

import (
	"fmt"
)

// Sender 发送
type Sender interface {
	Send()
}

// MailSender 邮件发送器
type MailSender struct {
	// Sender
}

// Send 接口实现
func (MailSender) Send() {
	fmt.Println("this is mail sender")
}

// SmsSender 短信发送器
type SmsSender struct {
	// Sender
}

// Send 接口实现
func (SmsSender) Send() {
	fmt.Println("this is sms sender")
}

// SendFactory 工厂
type SendFactory struct {
}

//*********普通工厂模式

// ProduceSender 工厂制造发送器(普通工厂模式)
func (SendFactory) ProduceSender(t string) Sender {
	if t == "mail" {
		return MailSender{}
	} else if t == "sms" {
		return SmsSender{}
	} else {
		fmt.Println("请输入正确的类型")
		return nil
	}
}

//*********多个工厂模式

// ProduceMailSender 制造mail发送器
func (SendFactory) ProduceMailSender() Sender {
	return MailSender{}
}

// ProduceSmsSender 制造sms发送器
func (SendFactory) ProduceSmsSender() Sender {
	return SmsSender{}
}

// 简单工厂模式
