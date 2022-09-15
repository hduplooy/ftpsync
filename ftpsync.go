package ftpsync

import (
	"github.com/jlaffaye/ftp"
	"io/fs"
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

type session struct {
	Config     *Config
	Connection *ftp.ServerConn
}

func connect(config *Config) (*session, error) {
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

	return &session{config, c}, nil
}

func (session *session) close() error {
	return session.Connection.Quit()
}

func SyncFolders(config *Config) error {
	session, err := connect(config)
	if err != nil {
		return err
	}
	defer session.close()

	con := session.Connection
	err = con.ChangeDir(config.RemotePath)
	if err != nil {
		return err
	}

	return fs.WalkDir(os.DirFS(session.Config.LocalPath), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || path == "." {
			return nil
		}
		localInfo, _ := d.Info()
		remoteInfo, err := con.GetEntry(path)
		if d.IsDir() {
			if err != nil {
				err = con.MakeDir(path)
			}
		} else {
			if err != nil || localInfo.ModTime().Unix() > remoteInfo.Time.Unix() {
				rdr, err := os.Open(session.Config.LocalPath + path)

				if err == nil {
					err = con.Stor(path, rdr)
				}
				return err
			} else {
				err = nil
			}
		}
		return err
	})
}
