package app

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var calls atomic.Int32

func HandleWebSocketTimestamp(ctx context.Context, app *App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			app.Logger(ctx).Error("vivian: socket: [error] handshake failure", "err", "connection refused")
		} else {
			app.Logger(ctx).Info("vivian: socket: [ok] handshake success", "remote", conn.RemoteAddr(), "local", conn.LocalAddr())
		}
		defer conn.Close()

		reconnectChannel := make(chan int)
		defer close(reconnectChannel)

		go func() {
			for {
				select {
				case <-reconnectChannel:
					if err := app.socket.Get().LoggerSocket(ctx, "vivian: socket: [ok] reconnected"); err != nil {
						app.Logger(ctx).Error("vivian: socket: [error]", "err", err)
					}
					return
				case <-ctx.Done():
					app.Logger(ctx).Error("vivian: socket: [error]", "err", "context lost")
					return
				}
			}
		}()

		for {
			select {
			case <-ctx.Done():
				app.Logger(ctx).Error("vivian: socket: [error]", "err", "context lost")
				return
			default:
				timestamp, _ := app.socket.Get().Time(ctx)
				calls.Add(1)
				err := conn.WriteMessage(websocket.TextMessage, timestamp)
				if err != nil {
					if err := app.socket.Get().LoggerSocket(ctx, "vivian: socket: [error] disconnected <- broken pipe ?"); err != nil {
						app.Logger(ctx).Error("vivian: socket: [error]", "err", err)
					}
					reconnectChannel <- 1
					return
				}
				time.Sleep(time.Second)
			}
		}
	})
}

var liveConn *websocket.Conn
var socketSync sync.Mutex

//type liveData struct {
//	X uint32 `json:"success"`
//	Y uint32 `json:"failure"`
//}

func SocketCalls(ctx context.Context, app *App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			app.Logger(ctx).Error("vivian: socket: [error] handshake failure", "err", websocket.HandshakeError{})
			return
		}
		defer conn.Close()

		app.Logger(ctx).Info("vivian: socket: [ok] handshake success", "remote", conn.RemoteAddr(), "local", conn.LocalAddr())

		socketSync.Lock()
		if liveConn != nil {
			liveConn.Close()
		}
		socketSync.Unlock()
		liveConn = conn

		reconnectChannel := make(chan int)
		defer close(reconnectChannel)

		go func() {
			for {
				select {
				case <-reconnectChannel:
					if err := app.socket.Get().LoggerSocket(ctx, "vivian: [ok] socket: reconnected"); err != nil {
						app.Logger(ctx).Error("vivian: socket: [error]", "err", err)
					}
					return
				case <-ctx.Done():
					app.Logger(ctx).Error("vivian: socket: [error]", "err", "context lost")
					return
				}
			}
		}()

		//var once sync.Once
		for {
			select {
			case <-ctx.Done():
				app.Logger(ctx).Error("vivian: socket: [error]", "err", "context lost")
				return
			default:
				//data := liveData{
				//	X: uint32(login.LoginSuccess.Load()),
				//	Y: uint32(login.LoginFailure.Load()),
				//}
				marshal_data, err := json.Marshal(uint32(calls.Load()))
				if err != nil {
					app.Logger(ctx).Error("vivian: socket: [error]", "err", "unable to marshalize data")
				}
				//log current count per refresh
				//once.Do(func(){
				//	app.Logger(ctx).Debug("vivian: socket: [ok] timestamp.calls", "amt", data)
				//})
				err = liveConn.WriteMessage(websocket.TextMessage, marshal_data)
				if err != nil {
					if err := app.socket.Get().LoggerSocket(ctx, "vivian: socket: [error] disconnected <- broken pipe ?"); err != nil {
						app.Logger(ctx).Error("vivian: socket: [error]", "err", err)
					}
					reconnectChannel <- 1
					return
				}
				time.Sleep(time.Second)
			}
		}
	})
}
