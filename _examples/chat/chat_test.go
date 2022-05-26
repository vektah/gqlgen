package chat

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatSubscriptions(t *testing.T) {
	c := client.New(handler.NewDefaultServer(NewExecutableSchema(New())))

	const batchSize = 16
	var wg sync.WaitGroup
	for i := 0; i < 1024; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sub := c.Websocket(fmt.Sprintf(
				`subscription @user(username:"vektah") { messageAdded(roomName:"#gophers%d") { text createdBy } }`,
				i,
			))
			defer sub.Close()

			go func() {
				var resp interface{}
				time.Sleep(10 * time.Millisecond)
				err := c.Post(fmt.Sprintf(`mutation {
					a:post(text:"Hello!", roomName:"#gophers%d", username:"vektah") { id }
					b:post(text:"Hello Vektah!", roomName:"#gophers%d", username:"andrey") { id }
					c:post(text:"Whats up?", roomName:"#gophers%d", username:"vektah") { id }
				}`, i, i, i), &resp)
				assert.NoError(t, err)
			}()

			var msg struct {
				resp struct {
					MessageAdded struct {
						Text      string
						CreatedBy string
					}
				}
				err error
			}

			msg.err = sub.Next(&msg.resp)
			require.NoError(t, msg.err, "sub.Next")
			require.Equal(t, "Hello!", msg.resp.MessageAdded.Text)
			require.Equal(t, "vektah", msg.resp.MessageAdded.CreatedBy)

			msg.err = sub.Next(&msg.resp)
			require.NoError(t, msg.err, "sub.Next")
			require.Equal(t, "Whats up?", msg.resp.MessageAdded.Text)
			require.Equal(t, "vektah", msg.resp.MessageAdded.CreatedBy)
		}(i)
		// wait for goroutines to finish every N tests to not get unhandled keepalives while starving on CPU
		// TODO: add proper handling of keepalives to properly test concurrency?
		if i%batchSize == 0 {
			wg.Wait()
		}
	}
	wg.Wait()

	// 1 for the main thread, 1 for the testing package and remainder is reserved for the HTTP server threads
	// TODO: use something like runtime.Stack to filter out HTTP server threads,
	// TODO: which is required for proper concurrency and leaks testing
	require.Less(t, runtime.NumGoroutine(), 1+1+batchSize*2, "goroutine leak")
}
