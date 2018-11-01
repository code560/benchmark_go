package main

import (
	"fmt"
	"time"
)

const TestCount = 100000

func main() {

	{
		buffForTest := float64(1)
		Profile("float64+", func() {
			for i := buffForTest; i < TestCount; i++ {
				buffForTest += i
			}
		})
	}
	{
		buffForTest := float64(1)
		Profile("float64*", func() {
			for i := buffForTest; i < TestCount; i++ {
				buffForTest *= i
			}
		})
	}
	{
		buffForTest := float64(1)
		testData := float64(10000)
		Profile("float64/", func() {
			for i := buffForTest; i < TestCount; i++ {
				buffForTest = testData / i
			}
		})
	}
	{
		buffForTest := int64(1)
		Profile("int64+", func() {
			for i := buffForTest; i < TestCount; i++ {
				buffForTest += i
			}
		})
	}
	{

		buffForTest := int64(1)
		Profile("int64*", func() {
			for i := buffForTest; i < TestCount; i++ {
				buffForTest *= i
			}
		})
	}
	{

		buffForTest := int64(1)
		testData := int64(10000)
		Profile("int64/", func() {
			for i := buffForTest; i < TestCount; i++ {
				buffForTest = testData / i
			}
		})
	}
	{

		buffForTest := int(1)
		Profile("int*", func() {
			for i := buffForTest; i < TestCount; i++ {
				buffForTest *= i
			}
		})
	}
	{

		buffForTest := int(1)
		testData := int(10000)
		Profile("int/", func() {
			for i := buffForTest; i < TestCount; i++ {
				buffForTest = testData / i
			}
		})
	}
	{

		buffForTest := int32(1)
		Profile("int32*", func() {
			for i := buffForTest; i < TestCount; i++ {
				buffForTest *= i
			}
		})
	}
	{

		buffForTest := int32(1)
		testData := int32(10000)
		Profile("int32/", func() {
			for i := buffForTest; i < TestCount; i++ {
				buffForTest = testData / i
			}
		})
	}
	{

		buffForTest := float32(1)
		Profile("float32*", func() {
			for i := buffForTest; i < TestCount; i++ {
				buffForTest *= i
			}
		})
	}
	{
		buffForTest := float32(1)
		testData := float32(10000)
		Profile("float32/", func() {
			for i := buffForTest; i < TestCount; i++ {
				buffForTest = testData / i
			}
		})
	}
}

func Profile(label string, target func()) {
	start := time.Now()

	target()

	duration := time.Now().Sub(start)

	fmt.Printf("time(%s): %0.4f\n", label, float64(duration)/float64(time.Millisecond))

}
