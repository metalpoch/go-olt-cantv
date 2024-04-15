package model

type Config struct {
	ProxyUser       string `json:"proxy_user"`
	ProxyHost       string `json:"proxy_host"`
	SSHPrivateKey   string `json:"ssh_private_key"`
	SSHPrivatePassw string `json:"ssh_private_passw"`
	DirDB           string `json:"dir_db"`
}
