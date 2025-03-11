package seeder

import (
	"fmt"
	"math/rand"
	"message/app"
	"message/types/message"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Seeder struct {
	app app.App
}

func NewSeeder(app app.App) *Seeder {
	return &Seeder{
		app: app,
	}
}

func (s *Seeder) Seed(mMessages int, batchSize int) error {
	gofakeit.Seed(time.Now().UnixNano())

	// Fake users IDs
	users := []string{
		"user_1",
		"user_2",
		"user_3",
		"user_4",
		"user_5",
	}

	// Fake servers IDs
	servers := []string{
		"server_1",
		"server_2",
		"server_3",
		"server_4",
		"server_5",
	}

	// Fake channels IDs
	channels := []string{
		"general",
		"random",
		"help",
		"announcements",
		"off-topic",
	}

	for {
		messages := make([]*message.Message, 0, batchSize)

		for range batchSize {
			messageID := gofakeit.UUID()
			userID := users[rand.Intn(len(users))]

			// 70% chance of being in a server
			var serverID *string
			if rand.Float32() < 0.7 {
				randomServer := servers[rand.Intn(len(servers))]
				serverID = &randomServer
			}

			channelID := channels[rand.Intn(len(channels))]

			// Generate a random time within the last month
			createdAt := gofakeit.DateRange(
				time.Now().AddDate(0, -1, 0),
				time.Now(),
			)

			msg := &message.Message{
				ID:          &messageID,
				OwnerID:     &userID,
				Content:     generateRealisticMessage(),
				Attachments: generateAttachments(),
				ServerID:    serverID,
				ChannelID:   channelID,
				CreatedAt:   createdAt,
				UpdatedAt:   createdAt,
			}

			messages = append(messages, msg)
		}

		err := s.app.MessageService.MCreate(messages)
		if err != nil {
			return err
		}

		if mMessages < batchSize {
			batchSize = mMessages
			mMessages = 0
		} else {
			mMessages -= batchSize
		}

		if mMessages <= 0 {
			break
		}
	}

	return nil
}

func generateRealisticMessage() string {
	messageTypes := []string{"greeting", "question", "statement", "reaction", "sharing"}

	switch messageTypes[rand.Intn(len(messageTypes))] {
	case "greeting":
		return generateGreeting()
	case "question":
		return generateQuestion()
	case "reaction":
		return generateReaction()
	case "sharing":
		return generateSharing()
	default:
		return gofakeit.Sentence(10)
	}
}

func generateAttachments() []string {
	// 30% chance of having attachments
	if rand.Float32() > 0.3 {
		return []string{}
	}

	numAttachments := rand.Intn(3) + 1 // 1-3 attachments
	attachments := make([]string, 0, numAttachments)

	fileTypes := []string{
		"jpg", "png", "gif", "pdf", "doc", "mp4", "mp3",
	}

	for range numAttachments {
		fileType := fileTypes[rand.Intn(len(fileTypes))]
		fileName := fmt.Sprintf("%s.%s", gofakeit.UUID(), fileType)
		attachments = append(attachments, fileName)
	}

	return attachments
}

func generateGreeting() string {
	greetings := []string{
		"Hey everyone!",
		fmt.Sprintf("Hi! %s", gofakeit.Sentence(5)),
		fmt.Sprintf("Good %s! %s", getDayTime(), gofakeit.Sentence(5)),
		"Hello! Anyone around?",
		fmt.Sprintf("Hey! Has anyone %s?", gofakeit.Verb()),
	}
	return greetings[rand.Intn(len(greetings))]
}

func generateQuestion() string {
	templates := []string{
		"Has anyone tried %s?",
		"What do you think about %s?",
		"Does anyone know how to %s?",
		"I'm wondering if %s?",
		"Can someone help me understand %s?",
		"What's the best way to %s?",
	}

	subjects := []string{
		gofakeit.Company(),
		gofakeit.MovieName(),
		gofakeit.AppName(),
		gofakeit.ProgrammingLanguage(),
		gofakeit.HackerPhrase(),
	}

	return fmt.Sprintf(templates[rand.Intn(len(templates))], subjects[rand.Intn(len(subjects))])
}

func generateReaction() string {
	reactions := []string{
		"That's awesome! 🎉",
		"Wow, I didn't know that! 😮",
		"Thanks for sharing! 👍",
		"This is incredible! ⭐",
		"No way! 😱",
		"That's interesting 🤔",
		fmt.Sprintf("That's %s! 🎉", gofakeit.AdjectiveDescriptive()),
	}
	return reactions[rand.Intn(len(reactions))]
}

func generateSharing() string {
	templates := []string{
		"Check out this %s I found: %s",
		"Just discovered this amazing %s: %s",
		"Thought you might be interested in this %s: %s",
		"Found this cool %s: %s",
	}

	items := []string{
		"article",
		"website",
		"tool",
		"app",
		"resource",
	}

	return fmt.Sprintf(
		templates[rand.Intn(len(templates))],
		items[rand.Intn(len(items))],
		gofakeit.URL(),
	)
}

func getDayTime() string {
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		return "morning"
	case hour < 17:
		return "afternoon"
	default:
		return "evening"
	}
}
