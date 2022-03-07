package adapter

import "testing"

func TestAwsClientAdapter_CreateServer(t *testing.T) {
	var a ICreateServer = &AwsClientAdapter{
		Client: AwsClient{},
	}
	a.CreateServer(1.0, 2.0)
}

func TestAliyunClientAdapter_CreateServer(t *testing.T) {
	var a ICreateServer = &AliyunClientAdapter{
		Client: AliyunClient{},
	}
	a.CreateServer(1.0, 2.0)
}
