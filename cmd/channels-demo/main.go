package main

// import "concurrency"
import "github.com/mywyau/go_lang/concurrency"
	


func main() {
    concurrency.UnbufferedChannelExample()
    concurrency.BasicChannelDemo()
    concurrency.SelectChannelExample()
    concurrency.ChannelCloseDetection()
    concurrency.FanOutFanIn()
    concurrency.ChannelWithContext()
    concurrency.DirectionalChannelsExample()
    concurrency.NonBlockingChannel()
}
