package ftpsync

import (
	ftp "github.com/jlaffaye/ftp"
	"os"
	"strings"
	"time"
)

type Config struct {
	Host       string
	Port       string
	Username   string
	Password   string
	LocalPath  string
	RemotePath string
}

type Session struct {
	Config     *Config
	Connection *ftp.ServerConn
}

func Connect(config *Config) (*Session, error) {
	if config.Port == "" {
		config.Port = "21"
	}
	c, err := ftp.Dial(config.Host+":"+config.Port, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	err = c.Login(config.Username, config.Password)
	if err != nil {
		return nil, err
	}

	if !strings.HasSuffix(config.LocalPath, string(os.PathSeparator)) {
		config.LocalPath += string(os.PathSeparator)
	}
	if !strings.HasSuffix(config.RemotePath, string(os.PathSeparator)) {
		config.RemotePath += string(os.PathSeparator)
	}

	return &Session{config, c}, nil
}

func (session *Session) UploadIfNewer(relpath, src, dest string) error {
	srcinfo, err := os.Lstat(session.Config.LocalPath + relpath + string(os.PathSeparator) + src)
	if err != nil {
		return err
	}

	return nil
}

func (session *Session) SyncUp() error {

	return nil
}

func (session *Session) Close() error {
	return session.Connection.Quit()
}
