package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CabRides", func() {
	It("should calculate area of rectangle of length 5 and breadth 3 as 15", func() {
		rectangle := NewRectangle(5, 3)
		area := rectangle.Area()

		Expect(area).To(Equal(15))
	})

	It("should calculate area of rectangle of length 6 and breadth 9 as 54", func() {
		rectangle := NewRectangle(6, 9)
		Expect(rectangle.Area()).To(Equal(54))
	})

	It("should calculate the perimeter of a rectangle of length 5 and breadth 3 as 16", func() {
		rectangle := NewRectangle(5, 3)
		Expect(rectangle.Perimeter()).To(Equal(16))
	})

	It("should calculate the perimeter of a rectangle of length 6 and breadth 9 as 30", func() {
		rectangle := NewRectangle(6, 9)
		Expect(rectangle.Perimeter()).To(Equal(30))
	})

	It("should calculate the area of a square of side 10 as 100", func() {
		square := NewSquare(10)
		Expect(square.Area()).To(Equal(100))
	})

	It("should calculate the perimeter of a square of side 10 as 40", func() {
		square := NewSquare(10)
		Expect(square.Perimeter()).To(Equal(40))
	})

})
