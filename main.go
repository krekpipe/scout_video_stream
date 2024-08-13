package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aler9/goroslib"
	"github.com/aler9/goroslib/pkg/msg"
)

// Define the Frame message type from roller_eye
type Frame struct {
	msg.Package     `ros:"roller_eye"`
	msg.Definitions `ros:"int8 VIDEO_STREAM_H264=0,int8 VIDEO_STREAM_JPG=1,int8 AUDIO_STREAM_AAC=2"`
	Seq             uint32
	Stamp           uint64
	Session         uint32
	Type            int8
	Oseq            uint32
	Par1            int32
	Par2            int32
	Par3            int32
	Par4            int32
	Data            []uint8
}

var (
	cameraData         = make(chan []byte)
	ROSHostAddress     = "192.168.1.61:11311"
	localhostAddress   = "127.0.0.1"
	flagROSHostAddress = flag.String("h", ROSHostAddress, "ROS endpoint such as IP_ADDRESS:PORT")
	flaglocalhost      = flag.String("l", localhostAddress, "localhost address")
)

func main() {
	flag.Parse()

	// Create a new ROS node
	n, err := goroslib.NewNode(goroslib.NodeConf{
		Name:          "jpeg-stream-server",
		MasterAddress: *flagROSHostAddress,
		Host:          *flaglocalhost,
	})
	if err != nil {
		log.Fatalf("Failed to create node: %v", err)
	}
	defer n.Close()

	// Create a subscriber to the ROS topic
	sub, err := goroslib.NewSubscriber(goroslib.SubscriberConf{
		Node:      n,
		Topic:     "/CoreNode/jpg",
		Callback:  onMessageFrame,
		QueueSize: 0,
	})
	if err != nil {
		log.Fatalf("Failed to create subscriber: %v", err)
	}
	defer sub.Close()

	// Start the MJPEG server
	http.HandleFunc("/stream", streamHandler)
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Listen for interrupts
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

// onMessageFrame is called when a message is received from the ROS topic
func onMessageFrame(msg *Frame) {
	// Convert uint8 slice to a byte slice
	cameraData <- msg.Data
}

// streamHandler handles HTTP requests and serves MJPEG stream
func streamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")

	for {
		img, ok := <-cameraData
		if !ok {
			http.Error(w, "Failed to fetch image", http.StatusInternalServerError)
			return
		}

		// Write JPEG image in MJPEG format
		w.Write([]byte("--frame\r\n"))
		w.Write([]byte("Content-Type: image/jpeg\r\n\r\n"))
		w.Write(img)
		w.Write([]byte("\r\n"))

		// Simulate a frame rate (adjust as needed)
		time.Sleep(100 * time.Millisecond)
	}
}
