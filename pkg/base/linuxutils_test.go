package base

import (
	"errors"
	"testing"

	"github.com/sirupsen/logrus"

	"eos2git.cec.lab.emc.com/ECS/baremetal-csi-plugin.git/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

var luLogger = logrus.New()

func TestLinuxUtils_LsblkSuccess(t *testing.T) {

	e := &mocks.GoMockExecutor{}
	e.On("RunCmd", LsblkCmd).Return(mocks.LsblkTwoDevicesStr, "", nil)
	l := NewLinuxUtils(e, luLogger)

	out, err := l.Lsblk(DriveTypeDisk)
	assert.Nil(t, err)
	assert.NotNil(t, out)
	assert.Equal(t, 2, len(*out))

}

func TestLinuxUtils_LsblkFail(t *testing.T) {
	e := &mocks.GoMockExecutor{}
	l := NewLinuxUtils(e, luLogger)

	e.On(mocks.RunCmd, LsblkCmd).Return("not a json", "", nil).Times(1)
	out, err := l.Lsblk(DriveTypeDisk)
	assert.Nil(t, out)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unable to unmarshal output to LsblkOutput instance")

	expectedError := errors.New("lsblk failed")
	e.On(mocks.RunCmd, LsblkCmd).Return("", "", expectedError).Times(1)
	out, err = l.Lsblk(DriveTypeDisk)
	assert.Nil(t, out)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)

	e.On(mocks.RunCmd, LsblkCmd).Return(mocks.NoLsblkKeyStr, "", nil).Times(1)
	out, err = l.Lsblk(DriveTypeDisk)
	assert.Nil(t, out)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unexpected lsblk output format")
}

func Test_GetBmcIP(t *testing.T) {
	e := &mocks.GoMockExecutor{}
	l := NewLinuxUtils(e, luLogger)

	strOut := "IP Address Source       : DHCP Address \n IP Address              : 10.245.137.136"
	e.On(mocks.RunCmd, IpmitoolCmd).Return(strOut, "", nil).Times(1)
	ip := l.GetBmcIP()
	assert.Equal(t, "10.245.137.136", ip)

	strOut = "IP Address Source       : DHCP Address \n"
	e.On(mocks.RunCmd, IpmitoolCmd).Return(strOut, "", nil).Times(1)
	ip = l.GetBmcIP()
	assert.Equal(t, "", ip)

	expectedError := errors.New("ipmitool failed")
	e.On(mocks.RunCmd, IpmitoolCmd).Return("", "", expectedError).Times(1)
	ip = l.GetBmcIP()
	assert.Equal(t, "", ip)
}