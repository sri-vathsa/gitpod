package lift

// Lift can execute commands in a namespace context and return the stdin/out/err FDs
// to the caller. This allows us to lift commands into ring1, akin to `docker exec`.

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"unsafe"

	"github.com/gitpod-io/gitpod/common-go/log"
	"golang.org/x/sys/unix"
)

const (
	DefaultSocketPath = "/tmp/workspacekit-lift.socket"

	// must not clash with the socket control message types:
	//    https://github.com/torvalds/linux/blob/master/include/linux/socket.h#L163
	msgTypeLiftRequest = 0xCA01
)

type LiftRequest struct {
	Command []string `json:"command"`
}

func ServeLift(socket string) error {
	skt, err := net.Listen("unix", socket)
	if err != nil {
		return err
	}
	defer skt.Close()

	for {
		conn, err := skt.Accept()
		if err != nil {
			log.WithError(err).Error("cannot accept lift connection")
			continue
		}

		go func() {
			err := serveLiftClient(conn)
			if err != nil {
				log.WithError(err).Error("cannot serve lift connection")
			}
		}()
	}
}

func serveLiftClient(conn net.Conn) error {
	unixConn := conn.(*net.UnixConn)

	buf := make([]byte, unix.CmsgSpace(256))
	f, err := unixConn.File()
	if err != nil {
		return err
	}
	defer f.Close()
	connfd := int(f.Fd())

	_, _, _, _, err = unix.Recvmsg(connfd, nil, buf, 0)
	if err != nil {
		return err
	}

	msgs, err := unix.ParseSocketControlMessage(buf)
	if err != nil {
		return err
	}
	if len(msgs) != 1 {
		return fmt.Errorf("expected a single socket control message")
	}
	if msgs[0].Header.Type != msgTypeLiftRequest {
		return fmt.Errorf("expected lift control message")
	}

	var msg LiftRequest
	err = json.Unmarshal(msgs[0].Data, &msg)
	if err != nil {
		return err
	}
	if len(msg.Command) == 0 {
		return fmt.Errorf("expected non-empty command")
	}

	soutR, soutW, err := os.Pipe()
	if err != nil {
		return err
	}
	defer soutW.Close()

	serrR, serrW, err := os.Pipe()
	if err != nil {
		return err
	}
	defer serrW.Close()

	sinR, sinW, err := os.Pipe()
	if err != nil {
		return err
	}
	defer sinR.Close()

	cmd := exec.Command(msg.Command[0], msg.Command[1:]...)
	cmd.Stdout = soutW
	cmd.Stderr = serrW
	cmd.Stdin = sinR

	err = unix.Sendmsg(connfd, nil, unix.UnixRights(int(soutR.Fd()), int(serrR.Fd()), int(sinW.Fd())), nil, 0)
	if err != nil {
		return err
	}

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func RunCommand(socket string, command []string) error {
	rconn, err := net.Dial("unix", socket)
	if err != nil {
		return err
	}
	conn := rconn.(*net.UnixConn)
	f, err := conn.File()
	if err != nil {
		return err
	}
	defer f.Close()
	connfd := int(f.Fd())

	msg, err := json.Marshal(LiftRequest{Command: command})
	if err != nil {
		return err
	}
	datalen := len(msg)
	b := make([]byte, unix.CmsgSpace(datalen))
	h := (*unix.Cmsghdr)(unsafe.Pointer(&b[0]))
	h.Level = unix.SOL_SOCKET
	h.Type = msgTypeLiftRequest
	h.SetLen(unix.CmsgLen(datalen))
	copy(b[unix.CmsgLen(0):], msg)

	err = unix.Sendmsg(connfd, nil, b, nil, 0)
	if err != nil {
		return err
	}

	buf := make([]byte, unix.CmsgSpace(3*4)) // we expect 3 FDs
	_, _, _, _, err = unix.Recvmsg(connfd, nil, buf, 0)
	if err != nil {
		return err
	}

	msgs, err := unix.ParseSocketControlMessage(buf)
	if err != nil {
		return err
	}
	if len(msgs) != 1 {
		return fmt.Errorf("expected a single socket control message")
	}
	fds, err := unix.ParseUnixRights(&msgs[0])
	if err != nil {
		return err
	}
	if len(fds) != 3 {
		return fmt.Errorf("expected three file descriptors")
	}

	go io.Copy(os.Stdout, os.NewFile(uintptr(fds[0]), "stdout"))
	go io.Copy(os.Stderr, os.NewFile(uintptr(fds[1]), "stderr"))
	io.Copy(os.NewFile(uintptr(fds[2]), "stdin"), os.Stdin)

	return nil
}
