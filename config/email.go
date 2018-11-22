package config

import "time"

// EmailServer is email server address
var EmailServer = "smtp.gmail.com" //smtp.live.com

// EmailPort is 465 by default 587
var EmailPort = 587 // 587

// EmailUsername is username
var EmailUsername = "skyClean906@gmail.com" // hantig1986@outlook.com

// EmailPassword is password
var EmailPassword = "wyjxiyhqhcymcvfd" // tscj3490han919

// Constants for Email service
const (
	EmailTimeout = 20 * time.Second
)
