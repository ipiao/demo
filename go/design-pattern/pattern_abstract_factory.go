package designPattern

// 抽象工厂相对于简单工厂将工厂类也进行了抽象

// interface Sender
// struct SmsSender
// struct MailSender

// Provider 抽象制造器
type Provider interface {
	ProduceSender() Sender
}

// SmsFactory sms工厂
type SmsFactory struct {
	// Provider
}

// ProduceSender 接口实现
func (SmsFactory) ProduceSender() Sender {
	return SmsSender{}
}

// MailFactory mail工厂
type MailFactory struct {
	// Provider
}

// ProduceSender 接口实现
func (MailFactory) ProduceSender() Sender {
	return MailSender{}
}
