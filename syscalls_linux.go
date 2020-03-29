// Copyright (c) 2020 Peter Hagelund
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// +build linux,arm

package v4l2

import "syscall"

// Ioctl performs a low-level ioctl operation.
func Ioctl(fd int, op uint32, arg uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(op), arg)
	if err == 0 {
		return nil
	}
	return err
}

// SetFd sets the the bit for the fd in the fd set.
func SetFd(fd int, s *syscall.FdSet) {
	s.Bits[fd/64] |= 1 << (uint(fd) % 64)
}

// WaitFd waits for data to be ready for the specified fd.
func WaitFd(fd int) error {
	fdSet := &syscall.FdSet{}
	SetFd(fd, fdSet)
	timeout := &syscall.Timeval{
		Sec:  2,
		Usec: 0,
	}
	_, err := syscall.Select(fd+1, fdSet, nil, nil, timeout)
	return err
}
