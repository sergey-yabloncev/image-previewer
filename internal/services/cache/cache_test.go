package cache

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(3)
		var wasInCache bool
		var key Key

		wasInCache, key = c.Set("1", 100)
		fmt.Printf("%T: %#v \n", wasInCache, wasInCache)
		fmt.Printf("%T: %#v \n", key, key)
		fmt.Println("--------------------")
		wasInCache, key = c.Set("2", 200)
		fmt.Printf("%T: %#v \n", wasInCache, wasInCache)
		fmt.Printf("%T: %#v \n", key, key)
		fmt.Println("--------------------")
		wasInCache, key = c.Set("2", 300)
		fmt.Printf("%T: %#v \n", wasInCache, wasInCache)
		fmt.Printf("%T: %#v \n", key, key)
		wasInCache, key = c.Set("3", 300)
		fmt.Println("--------------------")
		fmt.Printf("%T: %#v \n", wasInCache, wasInCache)
		fmt.Printf("%T: %#v \n", key, key)
		wasInCache, key = c.Set("4", 400)
		fmt.Printf("%T: %#v \n", wasInCache, wasInCache)
		fmt.Printf("%T: %#v \n", key, key)
		fmt.Println("--------------------")
		wasInCache, key = c.Set("5", 500)
		fmt.Printf("%T: %#v \n", wasInCache, wasInCache)
		fmt.Printf("%T: %#v \n", key, key)
		fmt.Println("--------------------")
		wasInCache, key = c.Set("6", 600)
		fmt.Printf("%T: %#v \n", wasInCache, wasInCache)
		fmt.Printf("%T: %#v \n", key, key)
		fmt.Println("--------------------")
		wasInCache, key = c.Set("7", 700)
		fmt.Printf("%T: %#v \n", wasInCache, wasInCache)
		fmt.Printf("%T: %#v \n", key, key)
		fmt.Println("--------------------")
		fmt.Printf("%T: %#v \n", c, c)
		fmt.Println("--------------------")

		//c.Set("1", "")
		//fmt.Printf("%T: %#v \n", c, c)
		//fmt.Println("--------------------")
		//c.Set("2", "")
		//fmt.Printf("%T: %#v \n", c, c)
		//fmt.Println("--------------------")
		//c.Set("3", "")
		//fmt.Printf("%T: %#v \n", c, c)
		//fmt.Println("--------------------")
		//c.Set("4", "")
		//fmt.Printf("%T: %#v \n", c, c)
		//fmt.Println("--------------------")
		//c.Set("5", "")
		//fmt.Printf("%T: %#v \n", c, c)
		//fmt.Println("--------------------")
		//c.Set("6", "")
		//fmt.Printf("%T: %#v \n", c, c)
		//fmt.Println("--------------------")

	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache, _ := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache, _ = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache, _ = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
