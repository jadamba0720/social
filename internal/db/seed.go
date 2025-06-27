package db

import (
	"context"
	"log"
	"math/rand"

	"github.com/sikozonpc/social/internal/store"
)

var usernames = []string{
	"test", "admin", "superadmin", "user", "client",
}

var titles = []string{
	"Why Simplicity Wins",
	"The Power of Saying No",
	"How I Organize My Day",
	"5 Tools I Use Daily",
	"What I Wish I Knew Sooner",
	"A Beginner’s Guide to Focus",
	"Why Speed Isn’t Everything",
	"Lessons from a Side Project",
	"Stop Multitasking, Start Finishing",
	"Your First 100 Lines of Code",
	"Why I Switched to Markdown",
	"The Magic of Small Habits",
	"How to Avoid Burnout",
	"Notes on Learning in Public",
	"How I Write Without Distractions",
	"The Case for Boring Tech",
	"Thinking in Systems",
	"Don’t Over-Engineer It",
	"Tiny Projects, Big Wins",
	"You Don’t Need a Framework Yet",
}

var contents = []string{
	"Simplicity makes things easier to understand, maintain, and scale. In a world full of complexity, choose clarity.",
	"Learning to say no helps you focus on what truly matters. Don’t overcommit your time or energy.",
	"I use time blocks, a prioritized to-do list, and breaks to stay productive without burning out.",
	"From VS Code to Notion, here are five tools that help me stay productive every day.",
	"Some lessons only come from experience. Here’s what I wish I had known at the beginning of my journey.",
	"Focus is a skill. Start with small sessions, reduce distractions, and train your brain like a muscle.",
	"Fast results can be fragile. Sustainable growth comes from building steady, thoughtful systems.",
	"This side project taught me more than any tutorial: real problems, real feedback, real growth.",
	"Multitasking kills focus. Switch to single-tasking and watch your productivity and clarity soar.",
	"The first 100 lines are often messy—and that’s okay. Just start coding and iterate from there.",
	"Markdown is fast, portable, and minimal. Once I switched, I never looked back.",
	"Big changes start small. Tiny, consistent habits compound into massive results over time.",
	"Burnout doesn’t just come from doing too much—it comes from doing what doesn’t align with you.",
	"Sharing what you're learning helps others and reinforces your understanding. Be public, be brave.",
	"Turn off notifications, clear your desk, and use a writing ritual to enter deep focus mode.",
	"Sometimes, the best tech is the one that just works—no need for the shiny and new.",
	"Systems thinking helps you understand how things connect. It’s a superpower for problem-solvers.",
	"Over-engineering wastes time. Start simple, solve the problem, then improve when needed.",
	"Quick wins build momentum. Start small, build confidence, and grow over time.",
	"Learn the language and solve real problems first—frameworks can come later.",
}

var tags = []string{
	"simplicity",
	"productivity",
	"time-management",
	"tools",
	"life-lessons",
	"focus",
	"mindset",
	"side-projects",
	"deep-work",
	"programming",
	"markdown",
	"habits",
	"burnout",
	"learning",
	"writing",
	"tech-philosophy",
	"systems-thinking",
	"minimalism",
	"project-management",
	"beginner-friendly",
}

var comments = []string{
	"I totally agree—simplicity often leads to better results.",
	"Saying no changed my life. Great reminder!",
	"This planning method looks helpful. I’ll give it a try tomorrow.",
	"Thanks for the tool list! I’m already using two of them.",
	"This hit me hard—needed to hear it today.",
	"Short and practical guide. Helped me refocus instantly.",
	"Such a good point about speed vs. sustainability.",
	"Your side project story is inspiring. Thanks for sharing!",
	"I didn’t realize how much multitasking was hurting me. This was eye-opening.",
	"This makes starting less scary. Great encouragement!",
	"Markdown is so underrated. I’m using it for everything now.",
	"Love the habit-building angle. Small steps really do matter.",
	"This explains burnout better than most articles I’ve read.",
	"Learning in public is scary but rewarding. Well said.",
	"Perfect timing! I’ve been struggling to write with focus.",
	"I love boring tech—it keeps my systems stable and sane.",
	"Systems thinking is powerful. More people need to learn this!",
	"This saved me hours. Simplicity first, always.",
	"Exactly the push I needed to start something small today.",
	"I’ve been stuck picking frameworks—this clarified everything. Thanks!",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(5)

	for _, user := range users {
		if err := store.User.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
		}
	}

	posts := generatePosts(20, users)
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i],
			Email:    usernames[i] + "@gmail.com",
			Password: "1",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[i],
			Content: contents[i],
			Tags: []string{
				titles[rand.Intn(len(titles))],
				titles[rand.Intn(len(titles))],
				titles[rand.Intn(len(titles))],
			},
			// Comments: ,
		}
	}

	return posts
}
