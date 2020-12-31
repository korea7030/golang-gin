package service

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gitlab.com/pragmaticreviews/gin-poc/entity"
	"gitlab.com/pragmaticreviews/gin-poc/repository"
)

const (
	TITLE       = "Video Title"
	DESCRIPTION = "Video Description"
	URL         = "https://youtu.be/JgW-i2QjgHQ"
	FIRSTNAME   = "John"
	LASTNAME    = "Doe"
	EMAIL       = "jdoe@mail.com"
)

var testVideo entity.Video = entity.Video{
	Title:       TITLE,
	Description: DESCRIPTION,
	URL:         URL,
	Author: entity.Person{
		FirstName: FIRSTNAME,
		LastName:  LASTNAME,
		Email:     EMAIL,
	},
}

var _ = Describe("Video Service", func() {
	var (
		videoRepository repository.VideoRepository
		videoService    VideoService
	)

	// test 시작전 사전작업
	BeforeSuite(func() {
		videoRepository = repository.NewVideoRepository()
		videoService = New(videoRepository)
	})

	Describe("Fetching all existing videos", func() {
		// 조건(if와 같은 역할)
		Context("If there is a video in the database", func() {
			BeforeEach(func() {
				videoService.Save(testVideo)
			})

			// test1
			It("should return at least one element", func() {
				videoList := videoService.FindAll()

				Ω(videoList).ShouldNot(BeEmpty())
			})

			// test2
			It("should map the fields correctly", func() {
				firstVideo := videoService.FindAll()[0]

				Ω(firstVideo.Title).Should(Equal(TITLE))
				Ω(firstVideo.Description).Should(Equal(DESCRIPTION))
				Ω(firstVideo.URL).Should(Equal(URL))
				Ω(firstVideo.Author.FirstName).Should(Equal(FIRSTNAME))
				Ω(firstVideo.Author.LastName).Should(Equal(LASTNAME))
				Ω(firstVideo.Author.Email).Should(Equal(EMAIL))
			})

			AfterEach(func() {
				video := videoService.FindAll()[0]
				videoService.Delete(video)
			})
		})

		// 조건(if와 같은 역할)
		Context("If there are no videos in the database", func() {
			// test3
			It("should return an empty list", func() {
				videos := videoService.FindAll()
				Ω(videos).Should(BeEmpty())
			})
		})
	})

	// test 끝난 후 작업
	AfterSuite(func() {
		videoRepository.CloseDB()
	})
})
