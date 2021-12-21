package testDemo

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var logs *log.Logger

func doClean(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logs.Println("doClean:收到Cancel，做好收尾工作后马上退出。")
			return
		default:
			logs.Println("doClean:每隔1秒观察信号，继续观察...")
		}
	}
}

func doNothing(ctx context.Context) {
	for {
		time.Sleep(3 * time.Second)
		select {
		case <-ctx.Done():
			logs.Println("doNothing:收到Cancel，但不退出......")
			//return
		default:
			logs.Println("doNothing:每隔3秒观察信号，继续观察...")
		}
	}
}

func TestContext(t *testing.T) {
	logs = log.New(os.Stdout, "", log.Ltime)
	ctx, cancel := context.WithCancel(context.Background())

	// 传递ctx
	go doClean(ctx)
	go doNothing(ctx)

	// 主程序阻塞20秒，留给协程来演示
	time.Sleep(20 * time.Second)
	logs.Println("cancel")

	// 调用cancel：context.WithCancel 返回的CancelFunc
	cancel()

	// 发出cancel 命令后，主程序阻塞10秒，再看协程的运行情况
	time.Sleep(10 * time.Second)
}

func A(ctx context.Context) int {
	ctx = context.WithValue(ctx, "AFunction", "Great")

	go B(ctx)

	select {
	// 监测自己上层的ctx ...
	case <-ctx.Done():
		fmt.Println("A Done")
		return -1
	}
	return 1
}

func B(ctx context.Context) int {
	fmt.Println("A value in B:", ctx.Value("AFunction"))
	ctx = context.WithValue(ctx, "BFunction", 999)

	go C(ctx)

	select {
	// 监测自己上层的ctx ...
	case <-ctx.Done():
		fmt.Println("B Done")
		return -2
	}
	return 2
}

func C(ctx context.Context) int {
	fmt.Println("B value in C:", ctx.Value("AFunction"))
	fmt.Println("B value in C:", ctx.Value("BFunction"))
	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		fmt.Println("C Done")
		return -3
	}
	return 3
}

func TestContext2(t *testing.T) {
	// 自动取消(定时取消)
	{
		timeout := 10 * time.Second
		ctx, _ := context.WithTimeout(context.Background(), timeout)

		fmt.Println("A 执行完成，返回：", A(ctx))
		select {
		case <-ctx.Done():
			fmt.Println("context Done")
			break
		}
	}
	time.Sleep(20 * time.Second)
}

func TestContextUseInHttp(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/chk", checkHandler)

	middle := ContextMiddle(mux)
	http.ListenAndServe(":8080", middle)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(1 * time.Minute)
	cookie := http.Cookie{Name: "Check", Value: "42", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if chk := r.Context().Value("check"); chk == "42" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("welcome my friend!\n"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("sorry page not found"))
	}
}

// 先读取cookie再通过context传递值
func ContextMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("Check")
		if cookie != nil {
			ctx := context.WithValue(r.Context(), "check", cookie.Value)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
