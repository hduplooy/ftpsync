# ftpsync

ftpsync for golang

This is a work in progress and can at the moment just sync files upward to the ftp server. It won't currently delete any files that are on the server and not locally.

File times are checked and only newer ones will be uploaded.

Folders will be created when they are not on the server.

See the basic_upload.go file for an example.
